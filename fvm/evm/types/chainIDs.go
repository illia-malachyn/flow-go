package types

import (
	"math/big"

	"github.com/onflow/flow-go/model/flow"
)

var (
	FlowEVMPreviewNetChainID = big.NewInt(646)
	FlowEVMTestNetChainID    = big.NewInt(545)
	FlowEVMMainNetChainID    = big.NewInt(747)

	FlowEVMPreviewNetChainIDInUInt64 = FlowEVMPreviewNetChainID.Uint64()
	FlowEVMTestNetChainIDInUInt64    = FlowEVMTestNetChainID.Uint64()
	FlowEVMMainNetChainIDInUInt64    = FlowEVMMainNetChainID.Uint64()
)

func EVMChainIDFromFlowChainID(flowChainID flow.ChainID) *big.Int {
	// default evm chain ID is previewNet
	switch flowChainID {
	case flow.Mainnet:
		return FlowEVMMainNetChainID
	case flow.Testnet:
		return FlowEVMTestNetChainID
	default:
		return FlowEVMPreviewNetChainID
	}
}
