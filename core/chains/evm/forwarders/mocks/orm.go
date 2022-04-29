// Code generated by mockery v2.12.1. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	forwarders "github.com/smartcontractkit/chainlink/core/chains/evm/forwarders"
	mock "github.com/stretchr/testify/mock"

	testing "testing"

	utils "github.com/smartcontractkit/chainlink/core/utils"
)

// ORM is an autogenerated mock type for the ORM type
type ORM struct {
	mock.Mock
}

// CreateForwarder provides a mock function with given fields: addr, evmChainId
func (_m *ORM) CreateForwarder(addr common.Address, evmChainId utils.Big) (forwarders.Forwarder, error) {
	ret := _m.Called(addr, evmChainId)

	var r0 forwarders.Forwarder
	if rf, ok := ret.Get(0).(func(common.Address, utils.Big) forwarders.Forwarder); ok {
		r0 = rf(addr, evmChainId)
	} else {
		r0 = ret.Get(0).(forwarders.Forwarder)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, utils.Big) error); ok {
		r1 = rf(addr, evmChainId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteForwarder provides a mock function with given fields: id
func (_m *ORM) DeleteForwarder(id int32) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindForwarders provides a mock function with given fields: offset, limit
func (_m *ORM) FindForwarders(offset int, limit int) ([]forwarders.Forwarder, int, error) {
	ret := _m.Called(offset, limit)

	var r0 []forwarders.Forwarder
	if rf, ok := ret.Get(0).(func(int, int) []forwarders.Forwarder); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]forwarders.Forwarder)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(offset, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewORM creates a new instance of ORM. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewORM(t testing.TB) *ORM {
	mock := &ORM{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
