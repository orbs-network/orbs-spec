// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package services

import (
	"context"
	"github.com/orbs-network/go-mock"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service PublicApi

type MockPublicApi struct {
	mock.Mock
	handlers.MockTransactionResultsHandler
}

func (s *MockPublicApi) SendTransaction(ctx context.Context, input *SendTransactionInput) (*SendTransactionOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*SendTransactionOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockPublicApi) SendTransactionAsync(ctx context.Context, input *SendTransactionInput) (*SendTransactionOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*SendTransactionOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockPublicApi) RunQuery(ctx context.Context, input *RunQueryInput) (*RunQueryOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*RunQueryOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockPublicApi) GetTransactionStatus(ctx context.Context, input *GetTransactionStatusInput) (*GetTransactionStatusOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetTransactionStatusOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockPublicApi) GetTransactionReceiptProof(ctx context.Context, input *GetTransactionReceiptProofInput) (*GetTransactionReceiptProofOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetTransactionReceiptProofOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockPublicApi) GetBlock(ctx context.Context, input *GetBlockInput) (*GetBlockOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
