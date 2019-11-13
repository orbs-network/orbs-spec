// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusContext

type MockConsensusContext struct {
	mock.Mock
}

func (s *MockConsensusContext) RequestNewTransactionsBlock(ctx context.Context, input *RequestNewTransactionsBlockInput) (*RequestNewTransactionsBlockOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*RequestNewTransactionsBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusContext) RequestNewResultsBlock(ctx context.Context, input *RequestNewResultsBlockInput) (*RequestNewResultsBlockOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*RequestNewResultsBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusContext) ValidateTransactionsBlock(ctx context.Context, input *ValidateTransactionsBlockInput) (*ValidateTransactionsBlockOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*ValidateTransactionsBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusContext) ValidateResultsBlock(ctx context.Context, input *ValidateResultsBlockInput) (*ValidateResultsBlockOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*ValidateResultsBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusContext) RequestOrderingCommittee(ctx context.Context, input *RequestCommitteeInput) (*RequestCommitteeOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*RequestCommitteeOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusContext) RequestValidationCommittee(ctx context.Context, input *RequestCommitteeInput) (*RequestCommitteeOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*RequestCommitteeOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
