package eventloop

import (
	"context"
	"github.com/onflow/flow-go/consensus/hotstuff/helper"
	"io/ioutil"
	"sync"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/atomic"

	"github.com/onflow/flow-go/consensus/hotstuff/mocks"
	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/irrecoverable"
	"github.com/onflow/flow-go/module/metrics"
	"github.com/onflow/flow-go/utils/unittest"
)

// TestEventLoop performs unit testing of event loop, checks if submitted events are propagated
// to event handler as well as handling of timeouts.
func TestEventLoop(t *testing.T) {
	suite.Run(t, new(EventLoopTestSuite))
}

type EventLoopTestSuite struct {
	suite.Suite

	eh     *mocks.EventHandler
	cancel context.CancelFunc

	eventLoop *EventLoop
}

func (s *EventLoopTestSuite) SetupTest() {
	s.eh = mocks.NewEventHandler(s.T())
	s.eh.On("Start").Return(nil).Maybe()
	s.eh.On("TimeoutChannel").Return(time.NewTimer(10 * time.Second).C).Maybe()
	s.eh.On("OnLocalTimeout").Return(nil).Maybe()

	log := zerolog.New(ioutil.Discard)

	eventLoop, err := NewEventLoop(log, metrics.NewNoopCollector(), s.eh, time.Time{})
	require.NoError(s.T(), err)
	s.eventLoop = eventLoop

	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	signalerCtx, _ := irrecoverable.WithSignaler(ctx)

	s.eventLoop.Start(signalerCtx)
	unittest.RequireCloseBefore(s.T(), s.eventLoop.Ready(), 100*time.Millisecond, "event loop not started")
}

func (s *EventLoopTestSuite) TearDownTest() {
	s.cancel()
	unittest.RequireCloseBefore(s.T(), s.eventLoop.Done(), 100*time.Millisecond, "event loop not stopped")
}

// TestReadyDone tests if event loop stops internal worker thread
func (s *EventLoopTestSuite) TestReadyDone() {
	time.Sleep(1 * time.Second)
	go func() {
		s.cancel()
	}()

	unittest.RequireCloseBefore(s.T(), s.eventLoop.Done(), 100*time.Millisecond, "event loop not stopped")
}

// Test_SubmitQC tests that submitted proposal is eventually sent to event handler for processing
func (s *EventLoopTestSuite) Test_SubmitProposal() {
	proposal := unittest.BlockHeaderFixture()
	expectedProposal := model.ProposalFromFlow(proposal, proposal.View-1)
	processed := atomic.NewBool(false)
	s.eh.On("OnReceiveProposal", expectedProposal).Run(func(args mock.Arguments) {
		processed.Store(true)
	}).Return(nil).Once()
	s.eventLoop.SubmitProposal(proposal, proposal.View-1)
	require.Eventually(s.T(), processed.Load, time.Millisecond*100, time.Millisecond*10)
}

// Test_SubmitQC tests that submitted QC is eventually sent to `EventHandler.OnReceiveQc` for processing
func (s *EventLoopTestSuite) Test_SubmitQC() {
	// qcIngestionFunction is the archetype for EventLoop.OnQcConstructedFromVotes and EventLoop.OnNewQcDiscovered
	type qcIngestionFunction func(*flow.QuorumCertificate)

	testQCIngestionFunction := func(f qcIngestionFunction, qcView uint64) {
		qc := helper.MakeQC(helper.WithQCView(qcView))
		processed := atomic.NewBool(false)
		s.eh.On("OnReceiveQc", qc).Run(func(args mock.Arguments) {
			processed.Store(true)
		}).Return(nil).Once()
		f(qc)
		require.Eventually(s.T(), processed.Load, time.Millisecond*100, time.Millisecond*10)
	}

	s.Run("QCs handed to EventLoop.OnQcConstructedFromVotes are forwarded to EventHandler", func() {
		testQCIngestionFunction(s.eventLoop.OnQcConstructedFromVotes, 100)
	})

	s.Run("QCs handed to EventLoop.OnNewQcDiscovered are forwarded to EventHandler", func() {
		testQCIngestionFunction(s.eventLoop.OnNewQcDiscovered, 101)
	})
}

// Test_SubmitTC tests that submitted TC is eventually sent to `EventHandler.OnReceiveTc` for processing
func (s *EventLoopTestSuite) Test_SubmitTC() {
	// tcIngestionFunction is the archetype for EventLoop.OnTcConstructedFromTimeouts and EventLoop.OnNewTcDiscovered
	type tcIngestionFunction func(*flow.TimeoutCertificate)

	testTCIngestionFunction := func(f tcIngestionFunction) {
		tc := &flow.TimeoutCertificate{}
		processed := atomic.NewBool(false)
		s.eh.On("OnReceiveTc", tc).Run(func(args mock.Arguments) {
			processed.Store(true)
		}).Return(nil).Once()
		f(tc)
		require.Eventually(s.T(), processed.Load, time.Millisecond*100, time.Millisecond*10)
	}

	s.Run("QCs handed to EventLoop.OnTcConstructedFromTimeouts are forwarded to EventHandler", func() {
		testTCIngestionFunction(s.eventLoop.OnTcConstructedFromTimeouts)
	})

	s.Run("QCs handed to EventLoop.OnNewTcDiscovered are forwarded to EventHandler", func() {
		testTCIngestionFunction(s.eventLoop.OnNewTcDiscovered)
	})
}

// TestEventLoop_Timeout tests that event loop delivers timeout events to event handler under pressure
func TestEventLoop_Timeout(t *testing.T) {
	eh := &mocks.EventHandler{}
	processed := atomic.NewBool(false)
	eh.On("Start").Return(nil).Once()
	eh.On("TimeoutChannel").Return(time.NewTimer(100 * time.Millisecond).C)
	eh.On("OnReceiveQc", mock.Anything).Return(nil).Maybe()
	eh.On("OnReceiveProposal", mock.Anything).Return(nil).Maybe()
	eh.On("OnLocalTimeout").Run(func(args mock.Arguments) {
		processed.Store(true)
	}).Return(nil).Once()

	log := zerolog.New(ioutil.Discard)

	eventLoop, err := NewEventLoop(log, metrics.NewNoopCollector(), eh, time.Time{})
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	signalerCtx, _ := irrecoverable.WithSignaler(ctx)
	eventLoop.Start(signalerCtx)

	unittest.RequireCloseBefore(t, eventLoop.Ready(), 100*time.Millisecond, "event loop not stopped")

	time.Sleep(10 * time.Millisecond)

	var wg sync.WaitGroup
	wg.Add(2)

	// spam with proposals and QCs
	go func() {
		defer wg.Done()
		for !processed.Load() {
			qc := unittest.QuorumCertificateFixture()
			eventLoop.OnQcConstructedFromVotes(qc)
		}
	}()

	go func() {
		defer wg.Done()
		for !processed.Load() {
			proposal := unittest.BlockHeaderFixture()
			eventLoop.SubmitProposal(proposal, proposal.View-1)
		}
	}()

	require.Eventually(t, processed.Load, time.Millisecond*200, time.Millisecond*10)
	wg.Wait()

	cancel()
	unittest.RequireCloseBefore(t, eventLoop.Done(), 100*time.Millisecond, "event loop not stopped")
}

// TestReadyDoneWithStartTime tests that event loop correctly starts and schedules start of processing
// when startTime argument is used
func TestReadyDoneWithStartTime(t *testing.T) {
	eh := &mocks.EventHandler{}
	eh.On("Start").Return(nil)
	eh.On("TimeoutChannel").Return(time.NewTimer(10 * time.Second).C)
	eh.On("OnLocalTimeout").Return(nil)

	metrics := metrics.NewNoopCollector()

	log := zerolog.New(ioutil.Discard)

	startTimeDuration := 2 * time.Second
	startTime := time.Now().Add(startTimeDuration)
	eventLoop, err := NewEventLoop(log, metrics, eh, startTime)
	require.NoError(t, err)

	done := make(chan struct{})
	eh.On("OnReceiveProposal", mock.AnythingOfType("*model.Proposal")).Run(func(args mock.Arguments) {
		require.True(t, time.Now().After(startTime))
		close(done)
	}).Return(nil).Once()

	ctx, cancel := context.WithCancel(context.Background())
	signalerCtx, _ := irrecoverable.WithSignaler(ctx)
	eventLoop.Start(signalerCtx)

	unittest.RequireCloseBefore(t, eventLoop.Ready(), 100*time.Millisecond, "event loop not started")

	parentBlock := unittest.BlockHeaderFixture()
	block := unittest.BlockHeaderWithParentFixture(parentBlock)
	eventLoop.SubmitProposal(block, parentBlock.View)

	unittest.RequireCloseBefore(t, done, startTimeDuration+100*time.Millisecond, "proposal wasn't received")
	cancel()
	unittest.RequireCloseBefore(t, eventLoop.Done(), 100*time.Millisecond, "event loop not stopped")
}
