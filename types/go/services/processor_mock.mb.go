// AUTO GENERATED FILE (by membufc proto compiler v0.0.32)
package services

import (
	"context"
	"github.com/orbs-network/go-mock"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service Processor

type MockProcessor struct {
	mock.Mock
}

func (s *MockProcessor) ProcessCall(ctx context.Context, input *ProcessCallInput) (*ProcessCallOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*ProcessCallOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockProcessor) GetContractInfo(ctx context.Context, input *GetContractInfoInput) (*GetContractInfoOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetContractInfoOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockProcessor) RegisterContractSdkCallHandler(handler handlers.ContractSdkCallHandler) {
	s.Called(handler)
}
