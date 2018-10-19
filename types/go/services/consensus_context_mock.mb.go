// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package services

import (
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusContext

type MockConsensusContext struct {
	mock.Mock
}

func (s *MockConsensusContext) RequestNewTransactionsBlock(input *RequestNewTransactionsBlockInput) (*RequestNewTransactionsBlockOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*RequestNewTransactionsBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusContext) RequestNewResultsBlock(input *RequestNewResultsBlockInput) (*RequestNewResultsBlockOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*RequestNewResultsBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusContext) ValidateTransactionsBlock(input *ValidateTransactionsBlockInput) (*ValidateTransactionsBlockOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ValidateTransactionsBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusContext) ValidateResultsBlock(input *ValidateResultsBlockInput) (*ValidateResultsBlockOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ValidateResultsBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusContext) RequestOrderingCommittee(input *RequestCommitteeInput) (*RequestCommitteeOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*RequestCommitteeOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusContext) RequestValidationCommittee(input *RequestCommitteeInput) (*RequestCommitteeOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*RequestCommitteeOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

