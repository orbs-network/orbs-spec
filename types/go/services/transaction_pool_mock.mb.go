// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"github.com/orbs-network/go-mock"
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionPool

type MockTransactionPool struct {
	mock.Mock
	gossiptopics.MockTransactionRelayHandler
}

func (s *MockTransactionPool) AddNewTransaction(ctx context.Context, input *AddNewTransactionInput) (*AddNewTransactionOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*AddNewTransactionOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockTransactionPool) GetCommittedTransactionReceipt(ctx context.Context, input *GetCommittedTransactionReceiptInput) (*GetCommittedTransactionReceiptOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetCommittedTransactionReceiptOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockTransactionPool) GetTransactionsForOrdering(ctx context.Context, input *GetTransactionsForOrderingInput) (*GetTransactionsForOrderingOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetTransactionsForOrderingOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockTransactionPool) ValidateTransactionsForOrdering(ctx context.Context, input *ValidateTransactionsForOrderingInput) (*ValidateTransactionsForOrderingOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*ValidateTransactionsForOrderingOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockTransactionPool) CommitTransactionReceipts(ctx context.Context, input *CommitTransactionReceiptsInput) (*CommitTransactionReceiptsOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*CommitTransactionReceiptsOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockTransactionPool) RegisterTransactionResultsHandler(handler handlers.TransactionResultsHandler) {
	s.Called(handler)
}
