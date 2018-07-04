// AUTO GENERATED FILE (by membufc proto compiler)
package handlers

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusBlocksHandler

type MockConsensusBlocksHandler struct {
	mock.Mock
}

func (s *MockConsensusBlocksHandler) HandleTransactionsBlock(input *HandleTransactionsBlockInput) (*HandleTransactionsBlockOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*HandleTransactionsBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusBlocksHandler) HandleResultsBlock(input *HandleResultsBlockInput) (*HandleResultsBlockOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*HandleResultsBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

