// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package handlers

import (
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusBlocksHandler

type MockConsensusBlocksHandler struct {
	mock.Mock
}

func (s *MockConsensusBlocksHandler) HandleBlockConsensus(input *HandleBlockConsensusInput) (*HandleBlockConsensusOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*HandleBlockConsensusOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

