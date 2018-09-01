// AUTO GENERATED FILE (by membufc proto compiler v0.0.19)
package consensus

import (
)

/////////////////////////////////////////////////////////////////////////////
// enums

type ConsensusAlgoType uint16

const (
	CONSENSUS_ALGO_TYPE_RESERVED ConsensusAlgoType = 0
	CONSENSUS_ALGO_TYPE_BENCHMARK_CONSENSUS ConsensusAlgoType = 1
	CONSENSUS_ALGO_TYPE_LEAN_HELIX ConsensusAlgoType = 2
)

func (n ConsensusAlgoType) String() string {
	switch n {
	case CONSENSUS_ALGO_TYPE_RESERVED:
		return "CONSENSUS_ALGO_TYPE_RESERVED"
	case CONSENSUS_ALGO_TYPE_BENCHMARK_CONSENSUS:
		return "CONSENSUS_ALGO_TYPE_BENCHMARK_CONSENSUS"
	case CONSENSUS_ALGO_TYPE_LEAN_HELIX:
		return "CONSENSUS_ALGO_TYPE_LEAN_HELIX"
	}
	return "UNKNOWN"
}

