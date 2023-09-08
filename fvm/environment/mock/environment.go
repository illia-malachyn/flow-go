// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	atree "github.com/onflow/atree"
	ast "github.com/onflow/cadence/runtime/ast"

	attribute "go.opentelemetry.io/otel/attribute"

	cadence "github.com/onflow/cadence"

	common "github.com/onflow/cadence/runtime/common"

	environment "github.com/onflow/flow-go/fvm/environment"

	flow "github.com/onflow/flow-go/model/flow"

	interpreter "github.com/onflow/cadence/runtime/interpreter"

	meter "github.com/onflow/flow-go/fvm/meter"

	mock "github.com/stretchr/testify/mock"

	oteltrace "go.opentelemetry.io/otel/trace"

	runtime "github.com/onflow/flow-go/fvm/runtime"

	sema "github.com/onflow/cadence/runtime/sema"

	stdlib "github.com/onflow/cadence/runtime/stdlib"

	time "time"

	trace "github.com/onflow/flow-go/module/trace"

	tracing "github.com/onflow/flow-go/fvm/tracing"

	zerolog "github.com/rs/zerolog"
)

// Environment is an autogenerated mock type for the Environment type
type Environment struct {
	mock.Mock
}

// AccountKeysCount provides a mock function with given fields: address
func (_m *Environment) AccountKeysCount(address common.Address) (uint64, error) {
	ret := _m.Called(address)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address) (uint64, error)); ok {
		return rf(address)
	}
	if rf, ok := ret.Get(0).(func(common.Address) uint64); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountsStorageCapacity provides a mock function with given fields: addresses, payer, maxTxFees
func (_m *Environment) AccountsStorageCapacity(addresses []flow.Address, payer flow.Address, maxTxFees uint64) (cadence.Value, error) {
	ret := _m.Called(addresses, payer, maxTxFees)

	var r0 cadence.Value
	var r1 error
	if rf, ok := ret.Get(0).(func([]flow.Address, flow.Address, uint64) (cadence.Value, error)); ok {
		return rf(addresses, payer, maxTxFees)
	}
	if rf, ok := ret.Get(0).(func([]flow.Address, flow.Address, uint64) cadence.Value); ok {
		r0 = rf(addresses, payer, maxTxFees)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cadence.Value)
		}
	}

	if rf, ok := ret.Get(1).(func([]flow.Address, flow.Address, uint64) error); ok {
		r1 = rf(addresses, payer, maxTxFees)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddAccountKey provides a mock function with given fields: address, publicKey, hashAlgo, weight
func (_m *Environment) AddAccountKey(address common.Address, publicKey *stdlib.PublicKey, hashAlgo sema.HashAlgorithm, weight int) (*stdlib.AccountKey, error) {
	ret := _m.Called(address, publicKey, hashAlgo, weight)

	var r0 *stdlib.AccountKey
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, *stdlib.PublicKey, sema.HashAlgorithm, int) (*stdlib.AccountKey, error)); ok {
		return rf(address, publicKey, hashAlgo, weight)
	}
	if rf, ok := ret.Get(0).(func(common.Address, *stdlib.PublicKey, sema.HashAlgorithm, int) *stdlib.AccountKey); ok {
		r0 = rf(address, publicKey, hashAlgo, weight)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stdlib.AccountKey)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, *stdlib.PublicKey, sema.HashAlgorithm, int) error); ok {
		r1 = rf(address, publicKey, hashAlgo, weight)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllocateStorageIndex provides a mock function with given fields: owner
func (_m *Environment) AllocateStorageIndex(owner []byte) (atree.StorageIndex, error) {
	ret := _m.Called(owner)

	var r0 atree.StorageIndex
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (atree.StorageIndex, error)); ok {
		return rf(owner)
	}
	if rf, ok := ret.Get(0).(func([]byte) atree.StorageIndex); ok {
		r0 = rf(owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(atree.StorageIndex)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BLSAggregatePublicKeys provides a mock function with given fields: publicKeys
func (_m *Environment) BLSAggregatePublicKeys(publicKeys []*stdlib.PublicKey) (*stdlib.PublicKey, error) {
	ret := _m.Called(publicKeys)

	var r0 *stdlib.PublicKey
	var r1 error
	if rf, ok := ret.Get(0).(func([]*stdlib.PublicKey) (*stdlib.PublicKey, error)); ok {
		return rf(publicKeys)
	}
	if rf, ok := ret.Get(0).(func([]*stdlib.PublicKey) *stdlib.PublicKey); ok {
		r0 = rf(publicKeys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stdlib.PublicKey)
		}
	}

	if rf, ok := ret.Get(1).(func([]*stdlib.PublicKey) error); ok {
		r1 = rf(publicKeys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BLSAggregateSignatures provides a mock function with given fields: signatures
func (_m *Environment) BLSAggregateSignatures(signatures [][]byte) ([]byte, error) {
	ret := _m.Called(signatures)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([][]byte) ([]byte, error)); ok {
		return rf(signatures)
	}
	if rf, ok := ret.Get(0).(func([][]byte) []byte); ok {
		r0 = rf(signatures)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([][]byte) error); ok {
		r1 = rf(signatures)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BLSVerifyPOP provides a mock function with given fields: publicKey, signature
func (_m *Environment) BLSVerifyPOP(publicKey *stdlib.PublicKey, signature []byte) (bool, error) {
	ret := _m.Called(publicKey, signature)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*stdlib.PublicKey, []byte) (bool, error)); ok {
		return rf(publicKey, signature)
	}
	if rf, ok := ret.Get(0).(func(*stdlib.PublicKey, []byte) bool); ok {
		r0 = rf(publicKey, signature)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*stdlib.PublicKey, []byte) error); ok {
		r1 = rf(publicKey, signature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BorrowCadenceRuntime provides a mock function with given fields:
func (_m *Environment) BorrowCadenceRuntime() *runtime.ReusableCadenceRuntime {
	ret := _m.Called()

	var r0 *runtime.ReusableCadenceRuntime
	if rf, ok := ret.Get(0).(func() *runtime.ReusableCadenceRuntime); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*runtime.ReusableCadenceRuntime)
		}
	}

	return r0
}

// CheckPayerBalanceAndGetMaxTxFees provides a mock function with given fields: payer, inclusionEffort, executionEffort
func (_m *Environment) CheckPayerBalanceAndGetMaxTxFees(payer flow.Address, inclusionEffort uint64, executionEffort uint64) (cadence.Value, error) {
	ret := _m.Called(payer, inclusionEffort, executionEffort)

	var r0 cadence.Value
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Address, uint64, uint64) (cadence.Value, error)); ok {
		return rf(payer, inclusionEffort, executionEffort)
	}
	if rf, ok := ret.Get(0).(func(flow.Address, uint64, uint64) cadence.Value); ok {
		r0 = rf(payer, inclusionEffort, executionEffort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cadence.Value)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Address, uint64, uint64) error); ok {
		r1 = rf(payer, inclusionEffort, executionEffort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ComputationIntensities provides a mock function with given fields:
func (_m *Environment) ComputationIntensities() meter.MeteredComputationIntensities {
	ret := _m.Called()

	var r0 meter.MeteredComputationIntensities
	if rf, ok := ret.Get(0).(func() meter.MeteredComputationIntensities); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(meter.MeteredComputationIntensities)
		}
	}

	return r0
}

// ComputationUsed provides a mock function with given fields:
func (_m *Environment) ComputationUsed() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConvertedServiceEvents provides a mock function with given fields:
func (_m *Environment) ConvertedServiceEvents() flow.ServiceEventList {
	ret := _m.Called()

	var r0 flow.ServiceEventList
	if rf, ok := ret.Get(0).(func() flow.ServiceEventList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.ServiceEventList)
		}
	}

	return r0
}

// CreateAccount provides a mock function with given fields: payer
func (_m *Environment) CreateAccount(payer common.Address) (common.Address, error) {
	ret := _m.Called(payer)

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address) (common.Address, error)); ok {
		return rf(payer)
	}
	if rf, ok := ret.Get(0).(func(common.Address) common.Address); ok {
		r0 = rf(payer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(payer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecodeArgument provides a mock function with given fields: argument, argumentType
func (_m *Environment) DecodeArgument(argument []byte, argumentType cadence.Type) (cadence.Value, error) {
	ret := _m.Called(argument, argumentType)

	var r0 cadence.Value
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, cadence.Type) (cadence.Value, error)); ok {
		return rf(argument, argumentType)
	}
	if rf, ok := ret.Get(0).(func([]byte, cadence.Type) cadence.Value); ok {
		r0 = rf(argument, argumentType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cadence.Value)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, cadence.Type) error); ok {
		r1 = rf(argument, argumentType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeductTransactionFees provides a mock function with given fields: payer, inclusionEffort, executionEffort
func (_m *Environment) DeductTransactionFees(payer flow.Address, inclusionEffort uint64, executionEffort uint64) (cadence.Value, error) {
	ret := _m.Called(payer, inclusionEffort, executionEffort)

	var r0 cadence.Value
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Address, uint64, uint64) (cadence.Value, error)); ok {
		return rf(payer, inclusionEffort, executionEffort)
	}
	if rf, ok := ret.Get(0).(func(flow.Address, uint64, uint64) cadence.Value); ok {
		r0 = rf(payer, inclusionEffort, executionEffort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cadence.Value)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Address, uint64, uint64) error); ok {
		r1 = rf(payer, inclusionEffort, executionEffort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EmitEvent provides a mock function with given fields: _a0
func (_m *Environment) EmitEvent(_a0 cadence.Event) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(cadence.Event) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Events provides a mock function with given fields:
func (_m *Environment) Events() flow.EventsList {
	ret := _m.Called()

	var r0 flow.EventsList
	if rf, ok := ret.Get(0).(func() flow.EventsList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.EventsList)
		}
	}

	return r0
}

// FlushPendingUpdates provides a mock function with given fields:
func (_m *Environment) FlushPendingUpdates() (environment.ContractUpdates, error) {
	ret := _m.Called()

	var r0 environment.ContractUpdates
	var r1 error
	if rf, ok := ret.Get(0).(func() (environment.ContractUpdates, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() environment.ContractUpdates); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(environment.ContractUpdates)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateAccountID provides a mock function with given fields: address
func (_m *Environment) GenerateAccountID(address common.Address) (uint64, error) {
	ret := _m.Called(address)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address) (uint64, error)); ok {
		return rf(address)
	}
	if rf, ok := ret.Get(0).(func(common.Address) uint64); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateUUID provides a mock function with given fields:
func (_m *Environment) GenerateUUID() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: address
func (_m *Environment) GetAccount(address flow.Address) (*flow.Account, error) {
	ret := _m.Called(address)

	var r0 *flow.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Address) (*flow.Account, error)); ok {
		return rf(address)
	}
	if rf, ok := ret.Get(0).(func(flow.Address) *flow.Account); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountAvailableBalance provides a mock function with given fields: address
func (_m *Environment) GetAccountAvailableBalance(address common.Address) (uint64, error) {
	ret := _m.Called(address)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address) (uint64, error)); ok {
		return rf(address)
	}
	if rf, ok := ret.Get(0).(func(common.Address) uint64); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountBalance provides a mock function with given fields: address
func (_m *Environment) GetAccountBalance(address common.Address) (uint64, error) {
	ret := _m.Called(address)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address) (uint64, error)); ok {
		return rf(address)
	}
	if rf, ok := ret.Get(0).(func(common.Address) uint64); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountContractCode provides a mock function with given fields: location
func (_m *Environment) GetAccountContractCode(location common.AddressLocation) ([]byte, error) {
	ret := _m.Called(location)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(common.AddressLocation) ([]byte, error)); ok {
		return rf(location)
	}
	if rf, ok := ret.Get(0).(func(common.AddressLocation) []byte); ok {
		r0 = rf(location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(common.AddressLocation) error); ok {
		r1 = rf(location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountContractNames provides a mock function with given fields: address
func (_m *Environment) GetAccountContractNames(address common.Address) ([]string, error) {
	ret := _m.Called(address)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address) ([]string, error)); ok {
		return rf(address)
	}
	if rf, ok := ret.Get(0).(func(common.Address) []string); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccountKey provides a mock function with given fields: address, index
func (_m *Environment) GetAccountKey(address common.Address, index int) (*stdlib.AccountKey, error) {
	ret := _m.Called(address, index)

	var r0 *stdlib.AccountKey
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, int) (*stdlib.AccountKey, error)); ok {
		return rf(address, index)
	}
	if rf, ok := ret.Get(0).(func(common.Address, int) *stdlib.AccountKey); ok {
		r0 = rf(address, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stdlib.AccountKey)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, int) error); ok {
		r1 = rf(address, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockAtHeight provides a mock function with given fields: height
func (_m *Environment) GetBlockAtHeight(height uint64) (stdlib.Block, bool, error) {
	ret := _m.Called(height)

	var r0 stdlib.Block
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(uint64) (stdlib.Block, bool, error)); ok {
		return rf(height)
	}
	if rf, ok := ret.Get(0).(func(uint64) stdlib.Block); ok {
		r0 = rf(height)
	} else {
		r0 = ret.Get(0).(stdlib.Block)
	}

	if rf, ok := ret.Get(1).(func(uint64) bool); ok {
		r1 = rf(height)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(uint64) error); ok {
		r2 = rf(height)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetCode provides a mock function with given fields: location
func (_m *Environment) GetCode(location common.Location) ([]byte, error) {
	ret := _m.Called(location)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Location) ([]byte, error)); ok {
		return rf(location)
	}
	if rf, ok := ret.Get(0).(func(common.Location) []byte); ok {
		r0 = rf(location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Location) error); ok {
		r1 = rf(location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCurrentBlockHeight provides a mock function with given fields:
func (_m *Environment) GetCurrentBlockHeight() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInterpreterSharedState provides a mock function with given fields:
func (_m *Environment) GetInterpreterSharedState() *interpreter.SharedState {
	ret := _m.Called()

	var r0 *interpreter.SharedState
	if rf, ok := ret.Get(0).(func() *interpreter.SharedState); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*interpreter.SharedState)
		}
	}

	return r0
}

// GetOrLoadProgram provides a mock function with given fields: location, load
func (_m *Environment) GetOrLoadProgram(location common.Location, load func() (*interpreter.Program, error)) (*interpreter.Program, error) {
	ret := _m.Called(location, load)

	var r0 *interpreter.Program
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Location, func() (*interpreter.Program, error)) (*interpreter.Program, error)); ok {
		return rf(location, load)
	}
	if rf, ok := ret.Get(0).(func(common.Location, func() (*interpreter.Program, error)) *interpreter.Program); ok {
		r0 = rf(location, load)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*interpreter.Program)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Location, func() (*interpreter.Program, error)) error); ok {
		r1 = rf(location, load)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSigningAccounts provides a mock function with given fields:
func (_m *Environment) GetSigningAccounts() ([]common.Address, error) {
	ret := _m.Called()

	var r0 []common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]common.Address, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageCapacity provides a mock function with given fields: address
func (_m *Environment) GetStorageCapacity(address common.Address) (uint64, error) {
	ret := _m.Called(address)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address) (uint64, error)); ok {
		return rf(address)
	}
	if rf, ok := ret.Get(0).(func(common.Address) uint64); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageUsed provides a mock function with given fields: address
func (_m *Environment) GetStorageUsed(address common.Address) (uint64, error) {
	ret := _m.Called(address)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address) (uint64, error)); ok {
		return rf(address)
	}
	if rf, ok := ret.Get(0).(func(common.Address) uint64); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetValue provides a mock function with given fields: owner, key
func (_m *Environment) GetValue(owner []byte, key []byte) ([]byte, error) {
	ret := _m.Called(owner, key)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) ([]byte, error)); ok {
		return rf(owner, key)
	}
	if rf, ok := ret.Get(0).(func([]byte, []byte) []byte); ok {
		r0 = rf(owner, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, []byte) error); ok {
		r1 = rf(owner, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Hash provides a mock function with given fields: data, tag, hashAlgorithm
func (_m *Environment) Hash(data []byte, tag string, hashAlgorithm sema.HashAlgorithm) ([]byte, error) {
	ret := _m.Called(data, tag, hashAlgorithm)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, string, sema.HashAlgorithm) ([]byte, error)); ok {
		return rf(data, tag, hashAlgorithm)
	}
	if rf, ok := ret.Get(0).(func([]byte, string, sema.HashAlgorithm) []byte); ok {
		r0 = rf(data, tag, hashAlgorithm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, string, sema.HashAlgorithm) error); ok {
		r1 = rf(data, tag, hashAlgorithm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImplementationDebugLog provides a mock function with given fields: message
func (_m *Environment) ImplementationDebugLog(message string) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InteractionUsed provides a mock function with given fields:
func (_m *Environment) InteractionUsed() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsServiceAccountAuthorizer provides a mock function with given fields:
func (_m *Environment) IsServiceAccountAuthorizer() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LimitAccountStorage provides a mock function with given fields:
func (_m *Environment) LimitAccountStorage() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Logger provides a mock function with given fields:
func (_m *Environment) Logger() *zerolog.Logger {
	ret := _m.Called()

	var r0 *zerolog.Logger
	if rf, ok := ret.Get(0).(func() *zerolog.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*zerolog.Logger)
		}
	}

	return r0
}

// Logs provides a mock function with given fields:
func (_m *Environment) Logs() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MemoryUsed provides a mock function with given fields:
func (_m *Environment) MemoryUsed() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func() (uint64, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MeterComputation provides a mock function with given fields: operationType, intensity
func (_m *Environment) MeterComputation(operationType common.ComputationKind, intensity uint) error {
	ret := _m.Called(operationType, intensity)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.ComputationKind, uint) error); ok {
		r0 = rf(operationType, intensity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MeterEmittedEvent provides a mock function with given fields: byteSize
func (_m *Environment) MeterEmittedEvent(byteSize uint64) error {
	ret := _m.Called(byteSize)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(byteSize)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MeterMemory provides a mock function with given fields: usage
func (_m *Environment) MeterMemory(usage common.MemoryUsage) error {
	ret := _m.Called(usage)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.MemoryUsage) error); ok {
		r0 = rf(usage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProgramLog provides a mock function with given fields: _a0
func (_m *Environment) ProgramLog(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadRandom provides a mock function with given fields: _a0
func (_m *Environment) ReadRandom(_a0 []byte) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RecordTrace provides a mock function with given fields: operation, location, duration, attrs
func (_m *Environment) RecordTrace(operation string, location common.Location, duration time.Duration, attrs []attribute.KeyValue) {
	_m.Called(operation, location, duration, attrs)
}

// RemoveAccountContractCode provides a mock function with given fields: location
func (_m *Environment) RemoveAccountContractCode(location common.AddressLocation) error {
	ret := _m.Called(location)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.AddressLocation) error); ok {
		r0 = rf(location)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reset provides a mock function with given fields:
func (_m *Environment) Reset() {
	_m.Called()
}

// ResolveLocation provides a mock function with given fields: identifiers, location
func (_m *Environment) ResolveLocation(identifiers []ast.Identifier, location common.Location) ([]sema.ResolvedLocation, error) {
	ret := _m.Called(identifiers, location)

	var r0 []sema.ResolvedLocation
	var r1 error
	if rf, ok := ret.Get(0).(func([]ast.Identifier, common.Location) ([]sema.ResolvedLocation, error)); ok {
		return rf(identifiers, location)
	}
	if rf, ok := ret.Get(0).(func([]ast.Identifier, common.Location) []sema.ResolvedLocation); ok {
		r0 = rf(identifiers, location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sema.ResolvedLocation)
		}
	}

	if rf, ok := ret.Get(1).(func([]ast.Identifier, common.Location) error); ok {
		r1 = rf(identifiers, location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceOwnerChanged provides a mock function with given fields: _a0, resource, oldOwner, newOwner
func (_m *Environment) ResourceOwnerChanged(_a0 *interpreter.Interpreter, resource *interpreter.CompositeValue, oldOwner common.Address, newOwner common.Address) {
	_m.Called(_a0, resource, oldOwner, newOwner)
}

// ReturnCadenceRuntime provides a mock function with given fields: _a0
func (_m *Environment) ReturnCadenceRuntime(_a0 *runtime.ReusableCadenceRuntime) {
	_m.Called(_a0)
}

// RevokeAccountKey provides a mock function with given fields: address, index
func (_m *Environment) RevokeAccountKey(address common.Address, index int) (*stdlib.AccountKey, error) {
	ret := _m.Called(address, index)

	var r0 *stdlib.AccountKey
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address, int) (*stdlib.AccountKey, error)); ok {
		return rf(address, index)
	}
	if rf, ok := ret.Get(0).(func(common.Address, int) *stdlib.AccountKey); ok {
		r0 = rf(address, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*stdlib.AccountKey)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address, int) error); ok {
		r1 = rf(address, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceEvents provides a mock function with given fields:
func (_m *Environment) ServiceEvents() flow.EventsList {
	ret := _m.Called()

	var r0 flow.EventsList
	if rf, ok := ret.Get(0).(func() flow.EventsList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.EventsList)
		}
	}

	return r0
}

// SetInterpreterSharedState provides a mock function with given fields: state
func (_m *Environment) SetInterpreterSharedState(state *interpreter.SharedState) {
	_m.Called(state)
}

// SetValue provides a mock function with given fields: owner, key, value
func (_m *Environment) SetValue(owner []byte, key []byte, value []byte) error {
	ret := _m.Called(owner, key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte, []byte) error); ok {
		r0 = rf(owner, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartChildSpan provides a mock function with given fields: name, options
func (_m *Environment) StartChildSpan(name trace.SpanName, options ...oteltrace.SpanStartOption) tracing.TracerSpan {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 tracing.TracerSpan
	if rf, ok := ret.Get(0).(func(trace.SpanName, ...oteltrace.SpanStartOption) tracing.TracerSpan); ok {
		r0 = rf(name, options...)
	} else {
		r0 = ret.Get(0).(tracing.TracerSpan)
	}

	return r0
}

// TotalEmittedEventBytes provides a mock function with given fields:
func (_m *Environment) TotalEmittedEventBytes() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// TransactionFeesEnabled provides a mock function with given fields:
func (_m *Environment) TransactionFeesEnabled() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// TxID provides a mock function with given fields:
func (_m *Environment) TxID() flow.Identifier {
	ret := _m.Called()

	var r0 flow.Identifier
	if rf, ok := ret.Get(0).(func() flow.Identifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Identifier)
		}
	}

	return r0
}

// TxIndex provides a mock function with given fields:
func (_m *Environment) TxIndex() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// UpdateAccountContractCode provides a mock function with given fields: location, code
func (_m *Environment) UpdateAccountContractCode(location common.AddressLocation, code []byte) error {
	ret := _m.Called(location, code)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.AddressLocation, []byte) error); ok {
		r0 = rf(location, code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidatePublicKey provides a mock function with given fields: key
func (_m *Environment) ValidatePublicKey(key *stdlib.PublicKey) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*stdlib.PublicKey) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValueExists provides a mock function with given fields: owner, key
func (_m *Environment) ValueExists(owner []byte, key []byte) (bool, error) {
	ret := _m.Called(owner, key)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) (bool, error)); ok {
		return rf(owner, key)
	}
	if rf, ok := ret.Get(0).(func([]byte, []byte) bool); ok {
		r0 = rf(owner, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func([]byte, []byte) error); ok {
		r1 = rf(owner, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifySignature provides a mock function with given fields: signature, tag, signedData, publicKey, signatureAlgorithm, hashAlgorithm
func (_m *Environment) VerifySignature(signature []byte, tag string, signedData []byte, publicKey []byte, signatureAlgorithm sema.SignatureAlgorithm, hashAlgorithm sema.HashAlgorithm) (bool, error) {
	ret := _m.Called(signature, tag, signedData, publicKey, signatureAlgorithm, hashAlgorithm)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, string, []byte, []byte, sema.SignatureAlgorithm, sema.HashAlgorithm) (bool, error)); ok {
		return rf(signature, tag, signedData, publicKey, signatureAlgorithm, hashAlgorithm)
	}
	if rf, ok := ret.Get(0).(func([]byte, string, []byte, []byte, sema.SignatureAlgorithm, sema.HashAlgorithm) bool); ok {
		r0 = rf(signature, tag, signedData, publicKey, signatureAlgorithm, hashAlgorithm)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func([]byte, string, []byte, []byte, sema.SignatureAlgorithm, sema.HashAlgorithm) error); ok {
		r1 = rf(signature, tag, signedData, publicKey, signatureAlgorithm, hashAlgorithm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEnvironment interface {
	mock.TestingT
	Cleanup(func())
}

// NewEnvironment creates a new instance of Environment. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEnvironment(t mockConstructorTestingTNewEnvironment) *Environment {
	mock := &Environment{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
