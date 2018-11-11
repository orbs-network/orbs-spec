// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package services

import (
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusAlgo

type ConsensusAlgo interface {
	handlers.ConsensusBlocksHandler
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

func (x *OnNewConsensusRoundInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message OnNewConsensusRoundOutput (non serializable)

type OnNewConsensusRoundOutput struct {
}

func (x *OnNewConsensusRoundOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// enums
