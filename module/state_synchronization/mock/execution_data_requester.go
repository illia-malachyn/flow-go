// Code generated by mockery v2.13.1. DO NOT EDIT.

package state_synchronization

import (
	irrecoverable "github.com/onflow/flow-go/module/irrecoverable"
	mock "github.com/stretchr/testify/mock"

	model "github.com/onflow/flow-go/consensus/hotstuff/model"

	state_synchronization "github.com/onflow/flow-go/module/state_synchronization"
)

// ExecutionDataRequester is an autogenerated mock type for the ExecutionDataRequester type
type ExecutionDataRequester struct {
	mock.Mock
}

// AddOnExecutionDataFetchedConsumer provides a mock function with given fields: fn
func (_m *ExecutionDataRequester) AddOnExecutionDataFetchedConsumer(fn state_synchronization.ExecutionDataReceivedCallback) {
	_m.Called(fn)
}

// Done provides a mock function with given fields:
func (_m *ExecutionDataRequester) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// OnBlockFinalized provides a mock function with given fields: _a0
func (_m *ExecutionDataRequester) OnBlockFinalized(_a0 *model.Block) {
	_m.Called(_a0)
}

// Ready provides a mock function with given fields:
func (_m *ExecutionDataRequester) Ready() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *ExecutionDataRequester) Start(_a0 irrecoverable.SignalerContext) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewExecutionDataRequester interface {
	mock.TestingT
	Cleanup(func())
}

// NewExecutionDataRequester creates a new instance of ExecutionDataRequester. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExecutionDataRequester(t mockConstructorTestingTNewExecutionDataRequester) *ExecutionDataRequester {
	mock := &ExecutionDataRequester{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
