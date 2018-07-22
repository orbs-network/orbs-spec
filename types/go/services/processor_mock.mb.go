// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package services

import (
	"github.com/orbs-network/go-mock"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service Processor

type MockProcessor struct {
	mock.Mock
}

func (s *MockProcessor) ProcessCall(input *ProcessCallInput) (*ProcessCallOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ProcessCallOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockProcessor) DeployNativeService(input *DeployNativeServiceInput) (*DeployNativeServiceOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*DeployNativeServiceOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockProcessor) RegisterContractSdkCallHandler(handler handlers.ContractSdkCallHandler) {
	s.Called(handler)
}

