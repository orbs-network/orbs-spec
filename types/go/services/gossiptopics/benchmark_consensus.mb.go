// AUTO GENERATED FILE (by membufc proto compiler v0.0.17)
package gossiptopics

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service BenchmarkConsensus

type BenchmarkConsensus interface {
	BroadcastBenchmarkConsensusCommit(input *BenchmarkConsensusCommitInput) (*EmptyOutput, error)
	SendBenchmarkConsensusCommitted(input *BenchmarkConsensusCommittedInput) (*EmptyOutput, error)
	RegisterBenchmarkConsensusHandler(handler BenchmarkConsensusHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service BenchmarkConsensusHandler

type BenchmarkConsensusHandler interface {
	HandleBenchmarkConsensusCommit(input *BenchmarkConsensusCommitInput) (*EmptyOutput, error)
	HandleBenchmarkConsensusCommitted(input *BenchmarkConsensusCommittedInput) (*EmptyOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message BenchmarkConsensusCommitInput (non serializable)

type BenchmarkConsensusCommitInput struct {
	Message *gossipmessages.BenchmarkConsensusCommitMessage
}

/////////////////////////////////////////////////////////////////////////////
// message BenchmarkConsensusCommittedInput (non serializable)

type BenchmarkConsensusCommittedInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Message *gossipmessages.BenchmarkConsensusCommittedMessage
}

