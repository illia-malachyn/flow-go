// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import (
	common "github.com/onflow/cadence/common"

	mock "github.com/stretchr/testify/mock"
)

// AccountCreator is an autogenerated mock type for the AccountCreator type
type AccountCreator struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: runtimePayer
func (_m *AccountCreator) CreateAccount(runtimePayer common.Address) (common.Address, error) {
	ret := _m.Called(runtimePayer)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address) (common.Address, error)); ok {
		return rf(runtimePayer)
	}
	if rf, ok := ret.Get(0).(func(common.Address) common.Address); ok {
		r0 = rf(runtimePayer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(runtimePayer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccountCreator creates a new instance of AccountCreator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountCreator(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountCreator {
	mock := &AccountCreator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
