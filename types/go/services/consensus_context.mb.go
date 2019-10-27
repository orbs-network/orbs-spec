// AUTO GENERATED FILE (by membufc proto compiler v0.0.32)
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
	CurrentBlockHeight      primitives.BlockHeight
	MaxBlockSizeKb          uint32
	MaxNumberOfTransactions uint32
	PrevBlockHash           primitives.Sha256
	PrevBlockTimestamp      primitives.TimestampNano
	BlockProposerAddress    primitives.NodeAddress
}

func (x *RequestNewTransactionsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,MaxBlockSizeKb:%s,MaxNumberOfTransactions:%s,PrevBlockHash:%s,PrevBlockTimestamp:%s,BlockProposerAddress:%s,}", x.StringCurrentBlockHeight(), x.StringMaxBlockSizeKb(), x.StringMaxNumberOfTransactions(), x.StringPrevBlockHash(), x.StringPrevBlockTimestamp(), x.StringBlockProposerAddress())
}

func (x *RequestNewTransactionsBlockInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
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

func (x *RequestNewTransactionsBlockInput) StringPrevBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockTimestamp)
	return
}

func (x *RequestNewTransactionsBlockInput) StringBlockProposerAddress() (res string) {
	res = fmt.Sprintf("%s", x.BlockProposerAddress)
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
	CurrentBlockHeight   primitives.BlockHeight
	PrevBlockHash        primitives.Sha256
	TransactionsBlock    *protocol.TransactionsBlockContainer
	PrevBlockTimestamp   primitives.TimestampNano
	BlockProposerAddress primitives.NodeAddress
}

func (x *RequestNewResultsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,PrevBlockHash:%s,TransactionsBlock:%s,PrevBlockTimestamp:%s,BlockProposerAddress:%s,}", x.StringCurrentBlockHeight(), x.StringPrevBlockHash(), x.StringTransactionsBlock(), x.StringPrevBlockTimestamp(), x.StringBlockProposerAddress())
}

func (x *RequestNewResultsBlockInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
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

func (x *RequestNewResultsBlockInput) StringPrevBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockTimestamp)
	return
}

func (x *RequestNewResultsBlockInput) StringBlockProposerAddress() (res string) {
	res = fmt.Sprintf("%s", x.BlockProposerAddress)
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
	CurrentBlockHeight   primitives.BlockHeight
	TransactionsBlock    *protocol.TransactionsBlockContainer
	PrevBlockHash        primitives.Sha256
	PrevBlockTimestamp   primitives.TimestampNano
	BlockProposerAddress primitives.NodeAddress
}

func (x *ValidateTransactionsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,TransactionsBlock:%s,PrevBlockHash:%s,PrevBlockTimestamp:%s,BlockProposerAddress:%s,}", x.StringCurrentBlockHeight(), x.StringTransactionsBlock(), x.StringPrevBlockHash(), x.StringPrevBlockTimestamp(), x.StringBlockProposerAddress())
}

func (x *ValidateTransactionsBlockInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *ValidateTransactionsBlockInput) StringTransactionsBlock() (res string) {
	res = x.TransactionsBlock.String()
	return
}

func (x *ValidateTransactionsBlockInput) StringPrevBlockHash() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockHash)
	return
}

func (x *ValidateTransactionsBlockInput) StringPrevBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockTimestamp)
	return
}

func (x *ValidateTransactionsBlockInput) StringBlockProposerAddress() (res string) {
	res = fmt.Sprintf("%s", x.BlockProposerAddress)
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
	CurrentBlockHeight   primitives.BlockHeight
	ResultsBlock         *protocol.ResultsBlockContainer
	PrevBlockHash        primitives.Sha256
	TransactionsBlock    *protocol.TransactionsBlockContainer
	PrevBlockTimestamp   primitives.TimestampNano
	BlockProposerAddress primitives.NodeAddress
}

func (x *ValidateResultsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,ResultsBlock:%s,PrevBlockHash:%s,TransactionsBlock:%s,PrevBlockTimestamp:%s,BlockProposerAddress:%s,}", x.StringCurrentBlockHeight(), x.StringResultsBlock(), x.StringPrevBlockHash(), x.StringTransactionsBlock(), x.StringPrevBlockTimestamp(), x.StringBlockProposerAddress())
}

func (x *ValidateResultsBlockInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
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

func (x *ValidateResultsBlockInput) StringPrevBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockTimestamp)
	return
}

func (x *ValidateResultsBlockInput) StringBlockProposerAddress() (res string) {
	res = fmt.Sprintf("%s", x.BlockProposerAddress)
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
	CurrentBlockHeight primitives.BlockHeight
	RandomSeed         uint64
	MaxCommitteeSize   uint32
}

func (x *RequestCommitteeInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,RandomSeed:%s,MaxCommitteeSize:%s,}", x.StringCurrentBlockHeight(), x.StringRandomSeed(), x.StringMaxCommitteeSize())
}

func (x *RequestCommitteeInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
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
	NodeAddresses            []primitives.NodeAddress
	NodeRandomSeedPublicKeys []primitives.Bls1PublicKey
}

func (x *RequestCommitteeOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{NodeAddresses:%s,NodeRandomSeedPublicKeys:%s,}", x.StringNodeAddresses(), x.StringNodeRandomSeedPublicKeys())
}

func (x *RequestCommitteeOutput) StringNodeAddresses() (res string) {
	res = "["
	for _, v := range x.NodeAddresses {
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
