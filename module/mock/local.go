// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	crypto "github.com/onflow/flow-go/crypto"
	flow "github.com/onflow/flow-go/model/flow"

	hash "github.com/onflow/flow-go/crypto/hash"

	mock "github.com/stretchr/testify/mock"
)

// Local is an autogenerated mock type for the Local type
type Local struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *Local) Address() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NodeID provides a mock function with given fields:
func (_m *Local) NodeID() flow.Identifier {
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

// NotMeFilter provides a mock function with given fields:
func (_m *Local) NotMeFilter() flow.IdentityFilter[flow.Identity] {
	ret := _m.Called()

	var r0 flow.IdentityFilter[flow.Identity]
	if rf, ok := ret.Get(0).(func() flow.IdentityFilter[flow.Identity]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.IdentityFilter[flow.Identity])
		}
	}

	return r0
}

// Sign provides a mock function with given fields: _a0, _a1
func (_m *Local) Sign(_a0 []byte, _a1 hash.Hasher) (crypto.Signature, error) {
	ret := _m.Called(_a0, _a1)

	var r0 crypto.Signature
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, hash.Hasher) (crypto.Signature, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func([]byte, hash.Hasher) crypto.Signature); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.Signature)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, hash.Hasher) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignFunc provides a mock function with given fields: _a0, _a1, _a2
func (_m *Local) SignFunc(_a0 []byte, _a1 hash.Hasher, _a2 func(crypto.PrivateKey, []byte, hash.Hasher) (crypto.Signature, error)) (crypto.Signature, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 crypto.Signature
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, hash.Hasher, func(crypto.PrivateKey, []byte, hash.Hasher) (crypto.Signature, error)) (crypto.Signature, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func([]byte, hash.Hasher, func(crypto.PrivateKey, []byte, hash.Hasher) (crypto.Signature, error)) crypto.Signature); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.Signature)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, hash.Hasher, func(crypto.PrivateKey, []byte, hash.Hasher) (crypto.Signature, error)) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLocal interface {
	mock.TestingT
	Cleanup(func())
}

// NewLocal creates a new instance of Local. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLocal(t mockConstructorTestingTNewLocal) *Local {
	mock := &Local{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
