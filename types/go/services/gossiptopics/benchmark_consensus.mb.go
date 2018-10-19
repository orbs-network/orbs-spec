// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package gossiptopics

import (
	"fmt"
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

func (x *BenchmarkConsensusCommitInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Message:%s,}", x.StringMessage())
}

func (x *BenchmarkConsensusCommitInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BenchmarkConsensusCommittedInput (non serializable)

type BenchmarkConsensusCommittedInput struct {
	RecipientPublicKey primitives.Ed25519PublicKey
	Message *gossipmessages.BenchmarkConsensusCommittedMessage
}

func (x *BenchmarkConsensusCommittedInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientPublicKey:%s,Message:%s,}", x.StringRecipientPublicKey(), x.StringMessage())
}

func (x *BenchmarkConsensusCommittedInput) StringRecipientPublicKey() (res string) {
	res = fmt.Sprintf("%s", x.RecipientPublicKey)
	return
}

func (x *BenchmarkConsensusCommittedInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

