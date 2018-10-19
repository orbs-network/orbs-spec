// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package services

import (
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service CrosschainConnector

type MockCrosschainConnector struct {
	mock.Mock
}

func (s *MockCrosschainConnector) EthereumCallContract(input *EthereumCallContractInput) (*EthereumCallContractOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EthereumCallContractOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

