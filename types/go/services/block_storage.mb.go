// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package services

import (
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockStorage

type BlockStorage interface {
	gossiptopics.BlockSyncHandler
	CommitBlock(input *CommitBlockInput) (*CommitBlockOutput, error)
	GetTransactionsBlockHeader(input *GetTransactionsBlockHeaderInput) (*GetTransactionsBlockHeaderOutput, error)
	GetResultsBlockHeader(input *GetResultsBlockHeaderInput) (*GetResultsBlockHeaderOutput, error)
	GetTransactionReceipt(input *GetTransactionReceiptInput) (*GetTransactionReceiptOutput, error)
	GetLastCommittedBlockHeight(input *GetLastCommittedBlockHeightInput) (*GetLastCommittedBlockHeightOutput, error)
	ValidateBlockForCommit(input *ValidateBlockForCommitInput) (*ValidateBlockForCommitOutput, error)
	RegisterConsensusBlocksHandler(handler handlers.ConsensusBlocksHandler)
}

/////////////////////////////////////////////////////////////////////////////
// message CommitBlockInput (non serializable)

type CommitBlockInput struct {
	BlockPair *protocol.BlockPairContainer
}

func (x *CommitBlockInput) String() string {
	return fmt.Sprintf("{BlockPair:%s,}", x.StringBlockPair())
}

func (x *CommitBlockInput) StringBlockPair() (res string) {
	res = x.BlockPair.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CommitBlockOutput (non serializable)

type CommitBlockOutput struct {
}

func (x *CommitBlockOutput) String() string {
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsBlockHeaderInput (non serializable)

type GetTransactionsBlockHeaderInput struct {
	BlockHeight primitives.BlockHeight
}

func (x *GetTransactionsBlockHeaderInput) String() string {
	return fmt.Sprintf("{BlockHeight:%s,}", x.StringBlockHeight())
}

func (x *GetTransactionsBlockHeaderInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%x", x.BlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsBlockHeaderOutput (non serializable)

type GetTransactionsBlockHeaderOutput struct {
	TransactionsBlockHeader *protocol.TransactionsBlockHeader
	TransactionsBlockMetadata *protocol.TransactionsBlockMetadata
	TransactionsBlockProof *protocol.TransactionsBlockProof
}

func (x *GetTransactionsBlockHeaderOutput) String() string {
	return fmt.Sprintf("{TransactionsBlockHeader:%s,TransactionsBlockMetadata:%s,TransactionsBlockProof:%s,}", x.StringTransactionsBlockHeader(), x.StringTransactionsBlockMetadata(), x.StringTransactionsBlockProof())
}

func (x *GetTransactionsBlockHeaderOutput) StringTransactionsBlockHeader() (res string) {
	res = x.TransactionsBlockHeader.String()
	return
}

func (x *GetTransactionsBlockHeaderOutput) StringTransactionsBlockMetadata() (res string) {
	res = x.TransactionsBlockMetadata.String()
	return
}

func (x *GetTransactionsBlockHeaderOutput) StringTransactionsBlockProof() (res string) {
	res = x.TransactionsBlockProof.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetResultsBlockHeaderInput (non serializable)

type GetResultsBlockHeaderInput struct {
	BlockHeight primitives.BlockHeight
}

func (x *GetResultsBlockHeaderInput) String() string {
	return fmt.Sprintf("{BlockHeight:%s,}", x.StringBlockHeight())
}

func (x *GetResultsBlockHeaderInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%x", x.BlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetResultsBlockHeaderOutput (non serializable)

type GetResultsBlockHeaderOutput struct {
	ResultsBlockHeader *protocol.ResultsBlockHeader
	ResultsBlockProof *protocol.ResultsBlockProof
}

func (x *GetResultsBlockHeaderOutput) String() string {
	return fmt.Sprintf("{ResultsBlockHeader:%s,ResultsBlockProof:%s,}", x.StringResultsBlockHeader(), x.StringResultsBlockProof())
}

func (x *GetResultsBlockHeaderOutput) StringResultsBlockHeader() (res string) {
	res = x.ResultsBlockHeader.String()
	return
}

func (x *GetResultsBlockHeaderOutput) StringResultsBlockProof() (res string) {
	res = x.ResultsBlockProof.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionReceiptInput (non serializable)

type GetTransactionReceiptInput struct {
	Txhash primitives.Sha256
	TransactionTimestamp primitives.TimestampNano
}

func (x *GetTransactionReceiptInput) String() string {
	return fmt.Sprintf("{Txhash:%s,TransactionTimestamp:%s,}", x.StringTxhash(), x.StringTransactionTimestamp())
}

func (x *GetTransactionReceiptInput) StringTxhash() (res string) {
	res = fmt.Sprintf("%x", x.Txhash)
	return
}

func (x *GetTransactionReceiptInput) StringTransactionTimestamp() (res string) {
	res = fmt.Sprintf("%x", x.TransactionTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionReceiptOutput (non serializable)

type GetTransactionReceiptOutput struct {
	TransactionReceipt *protocol.TransactionReceipt
	BlockHeight primitives.BlockHeight
	BlockTimestamp primitives.TimestampNano
}

func (x *GetTransactionReceiptOutput) String() string {
	return fmt.Sprintf("{TransactionReceipt:%s,BlockHeight:%s,BlockTimestamp:%s,}", x.StringTransactionReceipt(), x.StringBlockHeight(), x.StringBlockTimestamp())
}

func (x *GetTransactionReceiptOutput) StringTransactionReceipt() (res string) {
	res = x.TransactionReceipt.String()
	return
}

func (x *GetTransactionReceiptOutput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%x", x.BlockHeight)
	return
}

func (x *GetTransactionReceiptOutput) StringBlockTimestamp() (res string) {
	res = fmt.Sprintf("%x", x.BlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetLastCommittedBlockHeightInput (non serializable)

type GetLastCommittedBlockHeightInput struct {
}

func (x *GetLastCommittedBlockHeightInput) String() string {
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message GetLastCommittedBlockHeightOutput (non serializable)

type GetLastCommittedBlockHeightOutput struct {
	LastCommittedBlockHeight primitives.BlockHeight
	LastCommittedBlockTimestamp primitives.TimestampNano
}

func (x *GetLastCommittedBlockHeightOutput) String() string {
	return fmt.Sprintf("{LastCommittedBlockHeight:%s,LastCommittedBlockTimestamp:%s,}", x.StringLastCommittedBlockHeight(), x.StringLastCommittedBlockTimestamp())
}

func (x *GetLastCommittedBlockHeightOutput) StringLastCommittedBlockHeight() (res string) {
	res = fmt.Sprintf("%x", x.LastCommittedBlockHeight)
	return
}

func (x *GetLastCommittedBlockHeightOutput) StringLastCommittedBlockTimestamp() (res string) {
	res = fmt.Sprintf("%x", x.LastCommittedBlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateBlockForCommitInput (non serializable)

type ValidateBlockForCommitInput struct {
	BlockPair *protocol.BlockPairContainer
}

func (x *ValidateBlockForCommitInput) String() string {
	return fmt.Sprintf("{BlockPair:%s,}", x.StringBlockPair())
}

func (x *ValidateBlockForCommitInput) StringBlockPair() (res string) {
	res = x.BlockPair.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateBlockForCommitOutput (non serializable)

type ValidateBlockForCommitOutput struct {
}

func (x *ValidateBlockForCommitOutput) String() string {
	return fmt.Sprintf("{}")
}

