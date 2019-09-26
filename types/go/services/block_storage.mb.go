// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package services

import (
	"context"
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
	CommitBlock(ctx context.Context, input *CommitBlockInput) (*CommitBlockOutput, error)
	GetTransactionsBlockHeader(ctx context.Context, input *GetTransactionsBlockHeaderInput) (*GetTransactionsBlockHeaderOutput, error)
	GetResultsBlockHeader(ctx context.Context, input *GetResultsBlockHeaderInput) (*GetResultsBlockHeaderOutput, error)
	GetTransactionReceipt(ctx context.Context, input *GetTransactionReceiptInput) (*GetTransactionReceiptOutput, error)
	GetLastCommittedBlockHeight(ctx context.Context, input *GetLastCommittedBlockHeightInput) (*GetLastCommittedBlockHeightOutput, error)
	GenerateReceiptProof(ctx context.Context, input *GenerateReceiptProofInput) (*GenerateReceiptProofOutput, error)
	GetBlockPair(ctx context.Context, input *GetBlockPairInput) (*GetBlockPairOutput, error)
	ValidateBlockForCommit(ctx context.Context, input *ValidateBlockForCommitInput) (*ValidateBlockForCommitOutput, error)
	GetCommittedBlockInfoByTime(ctx context.Context, input *GetCommittedBlockInfoByTimeInput) (*BlockInfoOutput, error)
	GetCommittedBlockInfoByHeight(ctx context.Context, input *GetCommittedBlockInfoByHeightInput) (*BlockInfoOutput, error)
	RegisterConsensusBlocksHandler(handler handlers.ConsensusBlocksHandler)
}

/////////////////////////////////////////////////////////////////////////////
// message GetCommittedBlockInfoByTimeInput (non serializable)

type GetCommittedBlockInfoByTimeInput struct {
	RefTime primitives.TimestampNano
}

func (x *GetCommittedBlockInfoByTimeInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RefTime:%s,}", x.StringRefTime())
}

func (x *GetCommittedBlockInfoByTimeInput) StringRefTime() (res string) {
	res = fmt.Sprintf("%s", x.RefTime)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetCommittedBlockInfoByHeightInput (non serializable)

type GetCommittedBlockInfoByHeightInput struct {
	BlockHeight primitives.BlockHeight
}

func (x *GetCommittedBlockInfoByHeightInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,}", x.StringBlockHeight())
}

func (x *GetCommittedBlockInfoByHeightInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BlockInfoOutput (non serializable)

type BlockInfoOutput struct {
	BlockHeight    primitives.BlockHeight
	BlockTimestamp primitives.TimestampNano
}

func (x *BlockInfoOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,BlockTimestamp:%s,}", x.StringBlockHeight(), x.StringBlockTimestamp())
}

func (x *BlockInfoOutput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *BlockInfoOutput) StringBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.BlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CommitBlockInput (non serializable)

type CommitBlockInput struct {
	BlockPair *protocol.BlockPairContainer
}

func (x *CommitBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
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
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsBlockHeaderInput (non serializable)

type GetTransactionsBlockHeaderInput struct {
	BlockHeight primitives.BlockHeight
}

func (x *GetTransactionsBlockHeaderInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,}", x.StringBlockHeight())
}

func (x *GetTransactionsBlockHeaderInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsBlockHeaderOutput (non serializable)

type GetTransactionsBlockHeaderOutput struct {
	TransactionsBlockHeader   *protocol.TransactionsBlockHeader
	TransactionsBlockMetadata *protocol.TransactionsBlockMetadata
	TransactionsBlockProof    *protocol.TransactionsBlockProof
}

func (x *GetTransactionsBlockHeaderOutput) String() string {
	if x == nil {
		return "<nil>"
	}
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
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,}", x.StringBlockHeight())
}

func (x *GetResultsBlockHeaderInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetResultsBlockHeaderOutput (non serializable)

type GetResultsBlockHeaderOutput struct {
	ResultsBlockHeader *protocol.ResultsBlockHeader
	ResultsBlockProof  *protocol.ResultsBlockProof
}

func (x *GetResultsBlockHeaderOutput) String() string {
	if x == nil {
		return "<nil>"
	}
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
	Txhash               primitives.Sha256
	TransactionTimestamp primitives.TimestampNano
}

func (x *GetTransactionReceiptInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Txhash:%s,TransactionTimestamp:%s,}", x.StringTxhash(), x.StringTransactionTimestamp())
}

func (x *GetTransactionReceiptInput) StringTxhash() (res string) {
	res = fmt.Sprintf("%s", x.Txhash)
	return
}

func (x *GetTransactionReceiptInput) StringTransactionTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.TransactionTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionReceiptOutput (non serializable)

type GetTransactionReceiptOutput struct {
	TransactionReceipt *protocol.TransactionReceipt
	BlockHeight        primitives.BlockHeight
	BlockTimestamp     primitives.TimestampNano
}

func (x *GetTransactionReceiptOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionReceipt:%s,BlockHeight:%s,BlockTimestamp:%s,}", x.StringTransactionReceipt(), x.StringBlockHeight(), x.StringBlockTimestamp())
}

func (x *GetTransactionReceiptOutput) StringTransactionReceipt() (res string) {
	res = x.TransactionReceipt.String()
	return
}

func (x *GetTransactionReceiptOutput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *GetTransactionReceiptOutput) StringBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.BlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetLastCommittedBlockHeightInput (non serializable)

type GetLastCommittedBlockHeightInput struct {
}

func (x *GetLastCommittedBlockHeightInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message GetLastCommittedBlockHeightOutput (non serializable)

type GetLastCommittedBlockHeightOutput struct {
	LastCommittedBlockHeight    primitives.BlockHeight
	LastCommittedBlockTimestamp primitives.TimestampNano
}

func (x *GetLastCommittedBlockHeightOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{LastCommittedBlockHeight:%s,LastCommittedBlockTimestamp:%s,}", x.StringLastCommittedBlockHeight(), x.StringLastCommittedBlockTimestamp())
}

func (x *GetLastCommittedBlockHeightOutput) StringLastCommittedBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.LastCommittedBlockHeight)
	return
}

func (x *GetLastCommittedBlockHeightOutput) StringLastCommittedBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.LastCommittedBlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateBlockForCommitInput (non serializable)

type ValidateBlockForCommitInput struct {
	BlockPair *protocol.BlockPairContainer
}

func (x *ValidateBlockForCommitInput) String() string {
	if x == nil {
		return "<nil>"
	}
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
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message GenerateReceiptProofInput (non serializable)

type GenerateReceiptProofInput struct {
	Txhash      primitives.Sha256
	BlockHeight primitives.BlockHeight
}

func (x *GenerateReceiptProofInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Txhash:%s,BlockHeight:%s,}", x.StringTxhash(), x.StringBlockHeight())
}

func (x *GenerateReceiptProofInput) StringTxhash() (res string) {
	res = fmt.Sprintf("%s", x.Txhash)
	return
}

func (x *GenerateReceiptProofInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GenerateReceiptProofOutput (non serializable)

type GenerateReceiptProofOutput struct {
	Proof *protocol.ReceiptProof
}

func (x *GenerateReceiptProofOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Proof:%s,}", x.StringProof())
}

func (x *GenerateReceiptProofOutput) StringProof() (res string) {
	res = x.Proof.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetBlockPairInput (non serializable)

type GetBlockPairInput struct {
	BlockHeight primitives.BlockHeight
}

func (x *GetBlockPairInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,}", x.StringBlockHeight())
}

func (x *GetBlockPairInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetBlockPairOutput (non serializable)

type GetBlockPairOutput struct {
	BlockPair *protocol.BlockPairContainer
}

func (x *GetBlockPairOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockPair:%s,}", x.StringBlockPair())
}

func (x *GetBlockPairOutput) StringBlockPair() (res string) {
	res = x.BlockPair.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
