// AUTO GENERATED FILE (by membufc proto compiler v0.0.11)
package services

import (
	"github.com/maraino/go-mock"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusAlgoLeanHelix

type MockConsensusAlgoLeanHelix struct {
	mock.Mock
	handlers.MockConsensusBlocksHandler
	handlers.MockLeanHelixConsensusGossipHandler
}

func (s *MockConsensusAlgoLeanHelix) OnNewConsensusRound(input *OnNewConsensusRoundInput) (*OnNewConsensusRoundOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*OnNewConsensusRoundOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

