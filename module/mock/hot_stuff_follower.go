// Code generated by mockery v2.13.1. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// HotStuffFollower is an autogenerated mock type for the HotStuffFollower type
type HotStuffFollower struct {
	mock.Mock
}

// Done provides a mock function with given fields:
func (_m *HotStuffFollower) Done() <-chan struct{} {
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

// Ready provides a mock function with given fields:
func (_m *HotStuffFollower) Ready() <-chan struct{} {
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

// SubmitProposal provides a mock function with given fields: proposal, parentView
func (_m *HotStuffFollower) SubmitProposal(proposal *flow.Header, parentView uint64) {
	_m.Called(proposal, parentView)
}

type mockConstructorTestingTNewHotStuffFollower interface {
	mock.TestingT
	Cleanup(func())
}

// NewHotStuffFollower creates a new instance of HotStuffFollower. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHotStuffFollower(t mockConstructorTestingTNewHotStuffFollower) *HotStuffFollower {
	mock := &HotStuffFollower{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
