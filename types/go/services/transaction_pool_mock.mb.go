// AUTO GENERATED FILE (by membufc proto compiler v0.0.14)
package services

import (
	"github.com/maraino/go-mock"
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionPool

type MockTransactionPool struct {
	mock.Mock
	gossiptopics.MockTransactionRelayHandler
}

func (s *MockTransactionPool) AddNewTransaction(input *AddNewTransactionInput) (*AddNewTransactionOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*AddNewTransactionOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockTransactionPool) GetCommittedTransactionReceipt(input *GetCommittedTransactionReceiptInput) (*GetCommittedTransactionReceiptOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GetCommittedTransactionReceiptOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockTransactionPool) GetTransactionsForOrdering(input *GetTransactionsForOrderingInput) (*GetTransactionsForOrderingOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GetTransactionsForOrderingOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockTransactionPool) ValidateTransactionsForOrdering(input *ValidateTransactionsForOrderingInput) (*ValidateTransactionsForOrderingOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ValidateTransactionsForOrderingOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockTransactionPool) CommitTransactionReceipts(input *CommitTransactionReceiptsInput) (*CommitTransactionReceiptsOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*CommitTransactionReceiptsOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockTransactionPool) RegisterTransactionResultsHandler(handler handlers.TransactionResultsHandler) {
	s.Called(handler)
}

