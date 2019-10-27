// AUTO GENERATED FILE (by membufc proto compiler v0.0.32)
package gossiptopics

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionRelay

type MockTransactionRelay struct {
	mock.Mock
}

func (s *MockTransactionRelay) BroadcastForwardedTransactions(ctx context.Context, input *ForwardedTransactionsInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
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

func (s *MockTransactionRelayHandler) HandleForwardedTransactions(ctx context.Context, input *ForwardedTransactionsInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
