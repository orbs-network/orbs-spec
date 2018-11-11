// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusContext

type ConsensusContext interface {
	RequestNewTransactionsBlock(ctx context.Context, input *RequestNewTransactionsBlockInput) (*RequestNewTransactionsBlockOutput, error)
	RequestNewResultsBlock(ctx context.Context, input *RequestNewResultsBlockInput) (*RequestNewResultsBlockOutput, error)
	ValidateTransactionsBlock(ctx context.Context, input *ValidateTransactionsBlockInput) (*ValidateTransactionsBlockOutput, error)
	ValidateResultsBlock(ctx context.Context, input *ValidateResultsBlockInput) (*ValidateResultsBlockOutput, error)
	RequestOrderingCommittee(ctx context.Context, input *RequestCommitteeInput) (*RequestCommitteeOutput, error)
	RequestValidationCommittee(ctx context.Context, input *RequestCommitteeInput) (*RequestCommitteeOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewTransactionsBlockInput (non serializable)

type RequestNewTransactionsBlockInput struct {
	BlockHeight             primitives.BlockHeight
	MaxBlockSizeKb          uint32
	MaxNumberOfTransactions uint32
	PrevBlockHash           primitives.Sha256
}

func (x *RequestNewTransactionsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,MaxBlockSizeKb:%s,MaxNumberOfTransactions:%s,PrevBlockHash:%s,}", x.StringBlockHeight(), x.StringMaxBlockSizeKb(), x.StringMaxNumberOfTransactions(), x.StringPrevBlockHash())
}

func (x *RequestNewTransactionsBlockInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *RequestNewTransactionsBlockInput) StringMaxBlockSizeKb() (res string) {
	res = fmt.Sprintf("%x", x.MaxBlockSizeKb)
	return
}

func (x *RequestNewTransactionsBlockInput) StringMaxNumberOfTransactions() (res string) {
	res = fmt.Sprintf("%x", x.MaxNumberOfTransactions)
	return
}

func (x *RequestNewTransactionsBlockInput) StringPrevBlockHash() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockHash)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewTransactionsBlockOutput (non serializable)

type RequestNewTransactionsBlockOutput struct {
	TransactionsBlock *protocol.TransactionsBlockContainer
}

func (x *RequestNewTransactionsBlockOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionsBlock:%s,}", x.StringTransactionsBlock())
}

func (x *RequestNewTransactionsBlockOutput) StringTransactionsBlock() (res string) {
	res = x.TransactionsBlock.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewResultsBlockInput (non serializable)

type RequestNewResultsBlockInput struct {
	BlockHeight       primitives.BlockHeight
	PrevBlockHash     primitives.Sha256
	TransactionsBlock *protocol.TransactionsBlockContainer
}

func (x *RequestNewResultsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,PrevBlockHash:%s,TransactionsBlock:%s,}", x.StringBlockHeight(), x.StringPrevBlockHash(), x.StringTransactionsBlock())
}

func (x *RequestNewResultsBlockInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *RequestNewResultsBlockInput) StringPrevBlockHash() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockHash)
	return
}

func (x *RequestNewResultsBlockInput) StringTransactionsBlock() (res string) {
	res = x.TransactionsBlock.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewResultsBlockOutput (non serializable)

type RequestNewResultsBlockOutput struct {
	ResultsBlock *protocol.ResultsBlockContainer
}

func (x *RequestNewResultsBlockOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ResultsBlock:%s,}", x.StringResultsBlock())
}

func (x *RequestNewResultsBlockOutput) StringResultsBlock() (res string) {
	res = x.ResultsBlock.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsBlockInput (non serializable)

type ValidateTransactionsBlockInput struct {
	TransactionsBlock *protocol.TransactionsBlockContainer
	PrevBlockHash     primitives.Sha256
}

func (x *ValidateTransactionsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionsBlock:%s,PrevBlockHash:%s,}", x.StringTransactionsBlock(), x.StringPrevBlockHash())
}

func (x *ValidateTransactionsBlockInput) StringTransactionsBlock() (res string) {
	res = x.TransactionsBlock.String()
	return
}

func (x *ValidateTransactionsBlockInput) StringPrevBlockHash() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockHash)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsBlockOutput (non serializable)

type ValidateTransactionsBlockOutput struct {
}

func (x *ValidateTransactionsBlockOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateResultsBlockInput (non serializable)

type ValidateResultsBlockInput struct {
	ResultsBlock      *protocol.ResultsBlockContainer
	PrevBlockHash     primitives.Sha256
	TransactionsBlock *protocol.TransactionsBlockContainer
}

func (x *ValidateResultsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ResultsBlock:%s,PrevBlockHash:%s,TransactionsBlock:%s,}", x.StringResultsBlock(), x.StringPrevBlockHash(), x.StringTransactionsBlock())
}

func (x *ValidateResultsBlockInput) StringResultsBlock() (res string) {
	res = x.ResultsBlock.String()
	return
}

func (x *ValidateResultsBlockInput) StringPrevBlockHash() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockHash)
	return
}

func (x *ValidateResultsBlockInput) StringTransactionsBlock() (res string) {
	res = x.TransactionsBlock.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateResultsBlockOutput (non serializable)

type ValidateResultsBlockOutput struct {
}

func (x *ValidateResultsBlockOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message RequestCommitteeInput (non serializable)

type RequestCommitteeInput struct {
	BlockHeight      primitives.BlockHeight
	RandomSeed       uint64
	MaxCommitteeSize uint32
}

func (x *RequestCommitteeInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,RandomSeed:%s,MaxCommitteeSize:%s,}", x.StringBlockHeight(), x.StringRandomSeed(), x.StringMaxCommitteeSize())
}

func (x *RequestCommitteeInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *RequestCommitteeInput) StringRandomSeed() (res string) {
	res = fmt.Sprintf("%x", x.RandomSeed)
	return
}

func (x *RequestCommitteeInput) StringMaxCommitteeSize() (res string) {
	res = fmt.Sprintf("%x", x.MaxCommitteeSize)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message RequestCommitteeOutput (non serializable)

type RequestCommitteeOutput struct {
	NodePublicKeys           []primitives.Ed25519PublicKey
	NodeRandomSeedPublicKeys []primitives.Bls1PublicKey
}

func (x *RequestCommitteeOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{NodePublicKeys:%s,NodeRandomSeedPublicKeys:%s,}", x.StringNodePublicKeys(), x.StringNodeRandomSeedPublicKeys())
}

func (x *RequestCommitteeOutput) StringNodePublicKeys() (res string) {
	res = "["
	for _, v := range x.NodePublicKeys {
		res += fmt.Sprintf("%s", v) + ","
	}
	res += "]"
	return
}

func (x *RequestCommitteeOutput) StringNodeRandomSeedPublicKeys() (res string) {
	res = "["
	for _, v := range x.NodeRandomSeedPublicKeys {
		res += fmt.Sprintf("%s", v) + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
