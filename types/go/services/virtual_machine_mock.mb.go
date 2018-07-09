// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package services

import (
	"github.com/maraino/go-mock"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service VirtualMachine

type MockVirtualMachine struct {
	mock.Mock
	handlers.MockContractSdkCallHandler
}

func (s *MockVirtualMachine) ProcessTransactionSet(input *ProcessTransactionSetInput) (*ProcessTransactionSetOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ProcessTransactionSetOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockVirtualMachine) RunLocalMethod(input *RunLocalMethodInput) (*RunLocalMethodOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*RunLocalMethodOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockVirtualMachine) TransactionSetPreOrder(input *TransactionSetPreOrderInput) (*TransactionSetPreOrderOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*TransactionSetPreOrderOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

