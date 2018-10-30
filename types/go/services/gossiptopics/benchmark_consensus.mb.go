// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package gossiptopics

import (
	"fmt"
	"context"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service BenchmarkConsensus

type BenchmarkConsensus interface {
	BroadcastBenchmarkConsensusCommit(ctx context.Context, input *BenchmarkConsensusCommitInput) (*EmptyOutput, error)
	SendBenchmarkConsensusCommitted(ctx context.Context, input *BenchmarkConsensusCommittedInput) (*EmptyOutput, error)
	RegisterBenchmarkConsensusHandler(handler BenchmarkConsensusHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service BenchmarkConsensusHandler

type BenchmarkConsensusHandler interface {
	HandleBenchmarkConsensusCommit(ctx context.Context, input *BenchmarkConsensusCommitInput) (*EmptyOutput, error)
	HandleBenchmarkConsensusCommitted(ctx context.Context, input *BenchmarkConsensusCommittedInput) (*EmptyOutput, error)
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

