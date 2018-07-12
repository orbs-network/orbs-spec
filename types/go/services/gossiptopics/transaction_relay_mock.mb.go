// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package gossiptopics

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionRelay

type MockTransactionRelay struct {
	mock.Mock
}

func (s *MockTransactionRelay) BroadcastForwardedTransactions(input *ForwardedTransactionsInput) (*TransactionRelayOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*TransactionRelayOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockTransactionRelay) RegisterTransactionRelayHandler(handler TransactionRelayHandler) {
	s.Called(handler)
}

/////////////////////////////////////////////////////////////////////////////
// service TransactionRelayHandler

type MockTransactionRelayHandler struct {
	mock.Mock
}

func (s *MockTransactionRelayHandler) HandleForwardedTransactions(input *ForwardedTransactionsInput) (*TransactionRelayOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*TransactionRelayOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

