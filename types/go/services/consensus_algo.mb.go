// AUTO GENERATED FILE (by membufc proto compiler v0.0.11)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusAlgoLeanHelix

type ConsensusAlgoLeanHelix interface {
	handlers.ConsensusBlocksHandler
	handlers.LeanHelixConsensusGossipHandler
	OnNewConsensusRound(input *OnNewConsensusRoundInput) (*OnNewConsensusRoundOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message OnNewConsensusRoundInput (non serializable)

type OnNewConsensusRoundInput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message OnNewConsensusRoundOutput (non serializable)

type OnNewConsensusRoundOutput struct {
}

