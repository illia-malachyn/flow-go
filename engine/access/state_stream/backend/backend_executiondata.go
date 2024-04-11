package backend

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/onflow/flow-go/engine/access/subscription"
	"github.com/onflow/flow-go/engine/common/rpc"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/executiondatasync/execution_data"
	"github.com/onflow/flow-go/storage"
)

type ExecutionDataResponse struct {
	Height        uint64
	ExecutionData *execution_data.BlockExecutionData
}

type ExecutionDataBackend struct {
	log     zerolog.Logger
	headers storage.Headers

	getExecutionData GetExecutionDataFunc

	subscriptionHandler  *subscription.SubscriptionHandler
	executionDataTracker subscription.ExecutionDataTracker
}

func (b *ExecutionDataBackend) GetExecutionDataByBlockID(ctx context.Context, blockID flow.Identifier) (*execution_data.BlockExecutionData, error) {
	header, err := b.headers.ByBlockID(blockID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return nil, fmt.Errorf("could not get block header for %s: %w", blockID, subscription.ErrBlockNotReady)
		}
		return nil, fmt.Errorf("could not get block header for %s: %w", blockID, err)
	}

	executionData, err := b.getExecutionData(ctx, header.Height)

	if err != nil {
		if !errors.Is(err, subscription.ErrBlockNotReady) {
			return nil, rpc.ConvertError(err, "could not get execution data", codes.Internal)
		}

		return nil, status.Errorf(codes.NotFound, "could not find execution data: %v", subscription.ErrBlockNotReady)
	}

	return executionData.BlockExecutionData, nil
}

func (b *ExecutionDataBackend) SubscribeExecutionData(ctx context.Context, startBlockID flow.Identifier, startHeight uint64) subscription.Subscription {
	nextHeight, err := b.executionDataTracker.GetStartHeight(ctx, startBlockID, startHeight)
	if err != nil {
		return subscription.NewFailedSubscription(err, "could not get start height")
	}

	return b.subscriptionHandler.Subscribe(ctx, nextHeight, b.getResponse)
}

func (b *ExecutionDataBackend) getResponse(ctx context.Context, height uint64) (interface{}, error) {
	executionData, err := b.getExecutionData(ctx, height)
	if err != nil {
		return nil, fmt.Errorf("could not get execution data for block %d: %w", height, err)
	}

	return &ExecutionDataResponse{
		Height:        height,
		ExecutionData: executionData.BlockExecutionData,
	}, nil
}
