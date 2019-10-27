// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package services

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service ManagementConnector

type MockManagementConnector struct {
	mock.Mock
}

func (s *MockManagementConnector) GetBlockInfoByTime(ctx context.Context, input *GetBlockInfoByTimeInput) (*GetBlockInfoByTimeOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetBlockInfoByTimeOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockManagementConnector) CallContract(ctx context.Context, input *CallContractInput) (*CallContractOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*CallContractOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
