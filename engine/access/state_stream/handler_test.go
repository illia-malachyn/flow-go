package state_stream

import (
	"context"
	"fmt"
	"github.com/onflow/flow-go/engine/common/rpc/convert"
	"github.com/onflow/flow-go/utils/unittest"
	"time"

	"github.com/onflow/flow-go/model/flow"
	access "github.com/onflow/flow/protobuf/go/flow/executiondata"
	"github.com/stretchr/testify/require"
	pb "google.golang.org/genproto/googleapis/bytestream"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestHeartbeatResponseSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

type HandlerTestSuite struct {
	BackendExecutionDataSuite
	handler *Handler
}

type fakeReadServerImpl struct {
	pb.ByteStream_ReadServer
	ctx    context.Context
	//sender func(*access.SubscribeEventsResponse) error
	received chan *access.SubscribeEventsResponse
}

var _ access.ExecutionDataAPI_SubscribeEventsServer = (*fakeReadServerImpl)(nil)

func (fake *fakeReadServerImpl) Context() context.Context {
	return fake.ctx
}

func (fake *fakeReadServerImpl) Send(response *access.SubscribeEventsResponse) error {
	//return fake.sender(response)
	fake.received <- response
	return nil
}

func (s *HandlerTestSuite) SetupTest() {
	s.BackendExecutionDataSuite.SetupTest()
	conf := DefaultEventFilterConfig
	chain := flow.MonotonicEmulator.Chain()
	s.handler = NewHandler(s.backend, chain, conf, 5)
}

func (s *HandlerTestSuite) TestHeartbeatResponse() {
	reader := &fakeReadServerImpl{
		ctx:                   context.Background(),
		//sender: func(data *access.SubscribeEventsResponse) error {
		//	fmt.Printf("got block at height: %v\n", data.BlockHeight)
		//	return nil
		//},
		received: make(chan *access.SubscribeEventsResponse, 100),
	}

	s.backend.setHighestHeight(s.blocks[len(s.blocks) - 1].Header.Height)


	s.Run("Empty event filter", func() {

		filter := &access.EventFilter{}
		req := &access.SubscribeEventsRequest{
			//StartBlockId:         s.blocks[0].ID()[:],
			StartBlockHeight:     0,
			Filter:               filter,
			HeartbeatInterval:    1,
		}

		go func() {
			err := s.handler.SubscribeEvents(req, reader)
			require.NoError(s.T(), err)
		}()

		for _, b := range s.blocks {
			// consume execution data from subscription
			unittest.RequireReturnsBefore(s.T(), func() {
				resp, ok := <-reader.received
				require.True(s.T(), ok, "channel closed while waiting for exec data for block %d %v", b.Header.Height, b.ID())

				blockID, err := convert.BlockID(resp.BlockId)
				require.NoError(s.T(), err)
				require.Equal(s.T(), b.Header.ID(), blockID)
				require.Equal(s.T(), b.Header.Height, resp.BlockHeight)
			}, time.Second, fmt.Sprintf("timed out waiting for exec data for block %d %v", b.Header.Height, b.ID()))
		}
	})

	s.Run("Some events filter", func() {
		pbFilter := &access.EventFilter{
			EventType: []string{string(testEventTypes[0])},
			Contract: nil,
			Address: nil,
		}

		req := &access.SubscribeEventsRequest{
			//StartBlockId:         s.blocks[0].ID()[:],
			StartBlockHeight:     0,
			Filter:               pbFilter,
			HeartbeatInterval:    1,
		}

		go func() {
			err := s.handler.SubscribeEvents(req, reader)
			require.NoError(s.T(), err)
		}()


		for _, b := range s.blocks {

			// consume execution data from subscription
			unittest.RequireReturnsBefore(s.T(), func() {
				resp, ok := <-reader.received
				require.True(s.T(), ok, "channel closed while waiting for exec data for block %d %v", b.Header.Height, b.ID())

				blockID, err := convert.BlockID(resp.BlockId)
				require.NoError(s.T(), err)
				require.Equal(s.T(), b.Header.ID(), blockID)
				require.Equal(s.T(), b.Header.Height, resp.BlockHeight)
				//s.T().Logf(resp.Events[0].Type)
				//s.T().Logf("resp.Events len %d",(len(resp.Events)))
				//require.Equal(s.T(), expectedEvents, resp.Events)
			}, time.Second, fmt.Sprintf("timed out waiting for exec data for block %d %v", b.Header.Height, b.ID()))
		}
	})


	s.Run("No events", func() {
		pbFilter := &access.EventFilter{
			EventType: []string{"A.0x1.NonExistent.Event"},
			Contract: nil,
			Address: nil,
		}

		req := &access.SubscribeEventsRequest{
			//StartBlockId:         s.blocks[0].ID()[:],
			StartBlockHeight:     0,
			Filter:               pbFilter,
			HeartbeatInterval:    2,
		}

		go func() {
			err := s.handler.SubscribeEvents(req, reader)
			require.NoError(s.T(), err)
		}()

		expectedBlocks := make([]*flow.Block, 0)
		for i, block := range s.blocks {
			if (i+1) % int(req.HeartbeatInterval) == 0 {
				expectedBlocks = append(expectedBlocks, block)
			}
		}

		require.Len(s.T(), expectedBlocks, len(s.blocks) / int(req.HeartbeatInterval))

		for _, b := range expectedBlocks {
			// consume execution data from subscription
			unittest.RequireReturnsBefore(s.T(), func() {
				resp, ok := <-reader.received
				require.True(s.T(), ok, "channel closed while waiting for exec data for block %d %v", b.Header.Height, b.ID())

				blockID, err := convert.BlockID(resp.BlockId)
				require.NoError(s.T(), err)
				require.Equal(s.T(), b.Header.Height, resp.BlockHeight)
				require.Equal(s.T(), b.Header.ID(), blockID)
				require.Empty(s.T(), resp.Events)
			}, time.Second, fmt.Sprintf("timed out waiting for exec data for block %d %v", b.Header.Height, b.ID()))
		}
	})

}
