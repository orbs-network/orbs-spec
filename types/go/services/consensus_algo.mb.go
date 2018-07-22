// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusAlgo

type ConsensusAlgo interface {
	handlers.ConsensusBlocksHandler
	OnNewConsensusRound(input *OnNewConsensusRoundInput) (*OnNewConsensusRoundOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// service ConsensusAlgoLeanHelix

type ConsensusAlgoLeanHelix interface {
	ConsensusAlgo
	gossiptopics.LeanHelixHandler
}

/////////////////////////////////////////////////////////////////////////////
// service ConsensusAlgoBenchmark

type ConsensusAlgoBenchmark interface {
	ConsensusAlgo
	gossiptopics.BenchmarkConsensusHandler
}

/////////////////////////////////////////////////////////////////////////////
// message OnNewConsensusRoundInput (non serializable)

type OnNewConsensusRoundInput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message OnNewConsensusRoundOutput (non serializable)

type OnNewConsensusRoundOutput struct {
}

