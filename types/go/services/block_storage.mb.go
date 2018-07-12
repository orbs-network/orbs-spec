// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package services

import (
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

/////////////////////////////////////////////////////////////////////////////
// message CommitBlockOutput (non serializable)

type CommitBlockOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsBlockHeaderInput (non serializable)

type GetTransactionsBlockHeaderInput struct {
	BlockHeight primitives.BlockHeight
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsBlockHeaderOutput (non serializable)

type GetTransactionsBlockHeaderOutput struct {
	TransactionsBlockHeader *protocol.TransactionsBlockHeader
	TransactionsBlockMetadata *protocol.TransactionsBlockMetadata
	TransactionsBlockProof *protocol.TransactionsBlockProof
}

/////////////////////////////////////////////////////////////////////////////
// message GetResultsBlockHeaderInput (non serializable)

type GetResultsBlockHeaderInput struct {
	BlockHeight primitives.BlockHeight
}

/////////////////////////////////////////////////////////////////////////////
// message GetResultsBlockHeaderOutput (non serializable)

type GetResultsBlockHeaderOutput struct {
	ResultsBlockHeader *protocol.ResultsBlockHeader
	ResultsBlockProof *protocol.ResultsBlockProof
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionReceiptInput (non serializable)

type GetTransactionReceiptInput struct {
	Txhash primitives.Sha256
	TransactionTimestamp primitives.Timestamp
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionReceiptOutput (non serializable)

type GetTransactionReceiptOutput struct {
	TransactionReceipt *protocol.TransactionReceipt
	BlockHeight primitives.BlockHeight
	BlockTimestamp primitives.Timestamp
}

/////////////////////////////////////////////////////////////////////////////
// message GetLastCommittedBlockHeightInput (non serializable)

type GetLastCommittedBlockHeightInput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message GetLastCommittedBlockHeightOutput (non serializable)

type GetLastCommittedBlockHeightOutput struct {
	LastCommittedBlockHeight primitives.BlockHeight
	LastCommittedBlockTimestamp primitives.Timestamp
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateBlockForCommitInput (non serializable)

type ValidateBlockForCommitInput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateBlockForCommitOutput (non serializable)

type ValidateBlockForCommitOutput struct {
}

