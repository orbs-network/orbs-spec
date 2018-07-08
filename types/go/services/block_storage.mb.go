// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockStorage

type BlockStorage interface {
	handlers.BlockSyncGossipHandler
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
	BlockPair *protocol.BlockPair
}

/////////////////////////////////////////////////////////////////////////////
// message CommitBlockOutput (non serializable)

type CommitBlockOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsBlockHeaderInput (non serializable)

type GetTransactionsBlockHeaderInput struct {
	BlockHeight uint64
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsBlockHeaderOutput (non serializable)

type GetTransactionsBlockHeaderOutput struct {
	TransactionsBlockHeader *protocol.TransactionsBlockHeader
	TransactionsBlockMetaData *protocol.TransactionsBlockMetaData
	TransactionsBlockProof *protocol.TransactionsBlockProof
}

/////////////////////////////////////////////////////////////////////////////
// message GetResultsBlockHeaderInput (non serializable)

type GetResultsBlockHeaderInput struct {
	BlockHeight uint64
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
	Txid []byte
	Timestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionReceiptOutput (non serializable)

type GetTransactionReceiptOutput struct {
	Receipt *protocol.TransactionReceipt
	BlockHeight uint64
	BlockTimestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// message GetLastCommittedBlockHeightInput (non serializable)

type GetLastCommittedBlockHeightInput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message GetLastCommittedBlockHeightOutput (non serializable)

type GetLastCommittedBlockHeightOutput struct {
	LastCommittedBlockHeight uint64
	LastCommittedBlockTimestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateBlockForCommitInput (non serializable)

type ValidateBlockForCommitInput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateBlockForCommitOutput (non serializable)

type ValidateBlockForCommitOutput struct {
}

