// AUTO GENERATED FILE (by membufc proto compiler v0.0.32)
package gossiptopics

import (
	"context"
	"fmt"
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
	RecipientNodeAddress primitives.NodeAddress
	Message              *gossipmessages.BenchmarkConsensusCommittedMessage
}

func (x *BenchmarkConsensusCommittedInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientNodeAddress:%s,Message:%s,}", x.StringRecipientNodeAddress(), x.StringMessage())
}

func (x *BenchmarkConsensusCommittedInput) StringRecipientNodeAddress() (res string) {
	res = fmt.Sprintf("%s", x.RecipientNodeAddress)
	return
}

func (x *BenchmarkConsensusCommittedInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}
