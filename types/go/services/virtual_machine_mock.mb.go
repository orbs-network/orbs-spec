// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package services

import (
	"github.com/orbs-network/go-mock"
	"context"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service VirtualMachine

type MockVirtualMachine struct {
	mock.Mock
	handlers.MockContractSdkCallHandler
}

func (s *MockVirtualMachine) ProcessTransactionSet(ctx context.Context, input *ProcessTransactionSetInput) (*ProcessTransactionSetOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*ProcessTransactionSetOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockVirtualMachine) RunLocalMethod(ctx context.Context, input *RunLocalMethodInput) (*RunLocalMethodOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*RunLocalMethodOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockVirtualMachine) TransactionSetPreOrder(ctx context.Context, input *TransactionSetPreOrderInput) (*TransactionSetPreOrderOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*TransactionSetPreOrderOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

