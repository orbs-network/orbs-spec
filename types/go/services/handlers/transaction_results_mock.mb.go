// AUTO GENERATED FILE (by membufc proto compiler v0.0.32)
package handlers

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionResultsHandler

type MockTransactionResultsHandler struct {
	mock.Mock
}

func (s *MockTransactionResultsHandler) HandleTransactionResults(ctx context.Context, input *HandleTransactionResultsInput) (*HandleTransactionResultsOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*HandleTransactionResultsOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockTransactionResultsHandler) HandleTransactionError(ctx context.Context, input *HandleTransactionErrorInput) (*HandleTransactionErrorOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*HandleTransactionErrorOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
