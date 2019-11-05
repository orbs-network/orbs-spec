// AUTO GENERATED FILE (by membufc proto compiler v0.3.6)
package services

import (
	"context"
	"github.com/orbs-network/go-mock"
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

func (s *MockVirtualMachine) CallSystemContract(ctx context.Context, input *CallSystemContractInput) (*CallSystemContractOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*CallSystemContractOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockVirtualMachine) ProcessQuery(ctx context.Context, input *ProcessQueryInput) (*ProcessQueryOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*ProcessQueryOutput), ret.Error(1)
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
