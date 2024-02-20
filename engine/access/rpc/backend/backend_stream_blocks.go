package backend

import (
	"context"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/onflow/flow-go/engine"
	"github.com/onflow/flow-go/engine/access/subscription"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/state/protocol"
	"github.com/onflow/flow-go/storage"
	"github.com/onflow/flow-go/utils/logging"
)

// backendSubscribeBlocks is a struct representing a backend implementation for subscribing to blocks.
type backendSubscribeBlocks struct {
	log            zerolog.Logger
	state          protocol.State
	blocks         storage.Blocks
	Broadcaster    *engine.Broadcaster
	sendTimeout    time.Duration
	responseLimit  float64
	sendBufferSize int

	getStartHeight   subscription.GetStartHeightFunc
	getHighestHeight subscription.GetHighestHeight
}

// SubscribeBlocks subscribes to blocks starting from a specified block ID or height and with a given block status.
func (b *backendSubscribeBlocks) SubscribeBlocks(ctx context.Context, startBlockID flow.Identifier, startHeight uint64, blockStatus flow.BlockStatus) subscription.Subscription {
	nextHeight, err := b.getStartHeight(startBlockID, startHeight, blockStatus)
	if err != nil {
		return subscription.NewFailedSubscription(err, "could not get start height")
	}

	sub := subscription.NewHeightBasedSubscription(b.sendBufferSize, nextHeight, b.getBlockResponse(blockStatus))
	go subscription.NewStreamer(b.log, b.Broadcaster, b.sendTimeout, b.responseLimit, sub).Stream(ctx)

	return sub
}

// SubscribeBlockHeaders subscribes to block headers starting from a specified block ID or height and with a given block status.
func (b *backendSubscribeBlocks) SubscribeBlockHeaders(ctx context.Context, startBlockID flow.Identifier, startHeight uint64, blockStatus flow.BlockStatus) subscription.Subscription {
	nextHeight, err := b.getStartHeight(startBlockID, startHeight, blockStatus)
	if err != nil {
		return subscription.NewFailedSubscription(err, "could not get start height")
	}

	sub := subscription.NewHeightBasedSubscription(b.sendBufferSize, nextHeight, b.getBlockHeaderResponse(blockStatus))
	go subscription.NewStreamer(b.log, b.Broadcaster, b.sendTimeout, b.responseLimit, sub).Stream(ctx)

	return sub
}

// SubscribeBlockDigests subscribes to lightweight blocks starting from a specified block ID or height and with a given block status.
func (b *backendSubscribeBlocks) SubscribeBlockDigests(ctx context.Context, startBlockID flow.Identifier, startHeight uint64, blockStatus flow.BlockStatus) subscription.Subscription {
	nextHeight, err := b.getStartHeight(startBlockID, startHeight, blockStatus)
	if err != nil {
		return subscription.NewFailedSubscription(err, "could not get start height")
	}

	sub := subscription.NewHeightBasedSubscription(b.sendBufferSize, nextHeight, b.getBlockDigestResponse(blockStatus))
	go subscription.NewStreamer(b.log, b.Broadcaster, b.sendTimeout, b.responseLimit, sub).Stream(ctx)

	return sub
}

// getBlockResponse returns a GetDataByHeightFunc that retrieves block information for the specified height.
func (b *backendSubscribeBlocks) getBlockResponse(blockStatus flow.BlockStatus) subscription.GetDataByHeightFunc {
	return func(_ context.Context, height uint64) (interface{}, error) {
		block, err := b.getBlock(height, blockStatus)
		if err != nil {
			return nil, fmt.Errorf("could not get block by height %d: %w", height, err)
		}

		b.log.Trace().
			Hex("block_id", logging.ID(block.ID())).
			Uint64("height", height).
			Msgf("sending block info")

		return block, nil
	}
}

// getBlockHeaderResponse returns a GetDataByHeightFunc that retrieves block header information for the specified height.
func (b *backendSubscribeBlocks) getBlockHeaderResponse(blockStatus flow.BlockStatus) subscription.GetDataByHeightFunc {
	return func(_ context.Context, height uint64) (interface{}, error) {
		block, err := b.getBlock(height, blockStatus)
		if err != nil {
			return nil, fmt.Errorf("could not get block by height %d: %w", height, err)
		}

		b.log.Trace().
			Hex("block_id", logging.ID(block.ID())).
			Uint64("height", height).
			Msgf("sending block header info")

		return block.Header, nil
	}
}

// getBlockDigestResponse returns a GetDataByHeightFunc that retrieves lightweight block information for the specified height.
func (b *backendSubscribeBlocks) getBlockDigestResponse(blockStatus flow.BlockStatus) subscription.GetDataByHeightFunc {
	return func(_ context.Context, height uint64) (interface{}, error) {
		block, err := b.getBlock(height, blockStatus)
		if err != nil {
			return nil, fmt.Errorf("could not get block by height %d: %w", height, err)
		}

		b.log.Trace().
			Hex("block_id", logging.ID(block.ID())).
			Uint64("height", height).
			Msgf("sending lightweight block info")

		return &flow.LightweightBlock{
			ID:        block.ID(),
			Height:    block.Header.Height,
			Timestamp: block.Header.Timestamp,
		}, nil
	}
}

// getBlock returns the block for the given block height.
// Expected errors during normal operation:
// - storage.ErrNotFound or execution_data.BlobNotFoundError: block for the given block height is not available.
func (b *backendSubscribeBlocks) getBlock(height uint64, expectedBlockStatus flow.BlockStatus) (*flow.Block, error) {
	highestHeight, err := b.getHighestHeight(expectedBlockStatus)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get block by height: %v", err)
	}

	// fail early if no notification has been received for the given block height.
	// note: it's possible for the data to exist in the data store before the notification is
	// received. this ensures a consistent view is available to all streams.
	if height > highestHeight {
		return nil, fmt.Errorf("block %d is not available yet: %w", height, storage.ErrNotFound)
	}

	// since we are querying a finalized or sealed block, we can use the height index and save an ID computation
	block, err := b.blocks.ByHeight(height)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get block by height: %v", err)
	}

	return block, nil
}
