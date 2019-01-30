// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import multichannel "github.com/oktfolio/hyperledger-fabric-gm/orderer/common/multichannel"

// ChainGetter is an autogenerated mock type for the ChainGetter type
type ChainGetter struct {
	mock.Mock
}

// GetChain provides a mock function with given fields: chainID
func (_m *ChainGetter) GetChain(chainID string) *multichannel.ChainSupport {
	ret := _m.Called(chainID)

	var r0 *multichannel.ChainSupport
	if rf, ok := ret.Get(0).(func(string) *multichannel.ChainSupport); ok {
		r0 = rf(chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*multichannel.ChainSupport)
		}
	}

	return r0
}
