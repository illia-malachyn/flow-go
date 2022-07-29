package insecure

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/model/libp2p/message"
	"github.com/onflow/flow-go/network"
	"github.com/onflow/flow-go/network/channels"
	"github.com/onflow/flow-go/utils/unittest"
)

// EgressMessageFixture creates and returns a randomly generated gRPC egress message that is sent between a corruptible conduit and the attack network.
// It also generates and returns the corresponding application-layer event of that message, which is sent between the attack network and the
// orchestrator.
func EgressMessageFixture(t *testing.T, codec network.Codec, protocol Protocol, content interface{}) (*Message, *Event, *flow.Identity) {
	// fixture for content of message
	originId := unittest.IdentifierFixture()

	var targetIds flow.IdentifierList
	targetNum := uint32(0)

	if protocol == Protocol_UNICAST {
		targetIds = unittest.IdentifierListFixture(1)
	} else {
		targetIds = unittest.IdentifierListFixture(10)
	}

	if protocol == Protocol_MULTICAST {
		targetNum = uint32(3)
	}

	channel := channels.Channel("test-channel")
	// encodes event to create payload
	payload, err := codec.Encode(content)
	require.NoError(t, err)

	// creates egress message that goes over gRPC.

	egressMsg := &EgressMessage{
		ChannelID: "test-channel",
		OriginID:  originId[:],
		TargetNum: targetNum,
		TargetIDs: flow.IdsToBytes(targetIds),
		Payload:   payload,
		Protocol:  protocol,
	}

	m := &Message{
		Egress: egressMsg,
	}

	// creates corresponding event of that message that
	// is sent by attack network to orchestrator.
	e := &Event{
		CorruptedNodeId:   originId,
		Channel:           channel,
		FlowProtocolEvent: content,
		Protocol:          protocol,
		TargetNum:         targetNum,
		TargetIds:         targetIds,
	}

	return m, e, unittest.IdentityFixture(unittest.WithNodeID(originId), unittest.WithAddress("localhost:0"))
}

// IngressMessageFixture creates and returns a randomly generated gRPC ingress message that is sent between a corruptible conduit and the attack network.
// It also generates and returns the corresponding application-layer event of that message, which is sent between the attack network and the
// orchestrator.
func IngressMessageFixture(t *testing.T, codec network.Codec, protocol Protocol, content interface{}) (*Message, *Event, *flow.Identity) {
	// fixture for content of message
	originId := unittest.IdentifierFixture()

	var targetIds flow.IdentifierList
	targetNum := uint32(0)

	if protocol == Protocol_UNICAST {
		targetIds = unittest.IdentifierListFixture(1)
	} else {
		targetIds = unittest.IdentifierListFixture(10)
	}

	if protocol == Protocol_MULTICAST {
		targetNum = uint32(3)
	}

	channel := channels.Channel("test-channel")
	// encodes event to create payload
	payload, err := codec.Encode(content)
	require.NoError(t, err)

	// creates ingress message that goes over gRPC.
	ingressMsg := &IngressMessage{
		ChannelID: "test-channel",
		OriginID:  originId[:],
		Payload:   payload,
	}

	m := &Message{
		Ingress: ingressMsg,
	}

	// creates corresponding event of that message that
	// is sent by attack network to orchestrator.
	e := &Event{
		CorruptedNodeId:   originId,
		Channel:           channel,
		FlowProtocolEvent: content,
		Protocol:          protocol,
		TargetNum:         targetNum,
		TargetIds:         targetIds,
	}

	return m, e, unittest.IdentityFixture(unittest.WithNodeID(originId), unittest.WithAddress("localhost:0"))
}

// EgressMessageFixtures creates and returns randomly generated gRCP messages and their corresponding protocol-level events.
// The messages are sent between a corruptible conduit and the attack network.
// The events are the corresponding protocol-level representation of messages.
func EgressMessageFixtures(t *testing.T, codec network.Codec, protocol Protocol, count int) ([]*Message, []*Event,
	flow.IdentityList) {
	msgs := make([]*Message, count)
	events := make([]*Event, count)
	identities := flow.IdentityList{}

	for i := 0; i < count; i++ {
		m, e, id := EgressMessageFixture(t, codec, protocol, &message.TestMessage{
			Text: fmt.Sprintf("this is a test message: %d", rand.Int()),
		})

		msgs[i] = m
		events[i] = e
		// created identity must be unique
		require.NotContains(t, identities, id)
		identities = append(identities, id)
	}

	return msgs, events, identities
}
