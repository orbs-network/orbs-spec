// AUTO GENERATED FILE (by membufc proto compiler v0.0.17)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionPool

type TransactionPool interface {
	gossiptopics.TransactionRelayHandler
	AddNewTransaction(input *AddNewTransactionInput) (*AddNewTransactionOutput, error)
	GetCommittedTransactionReceipt(input *GetCommittedTransactionReceiptInput) (*GetCommittedTransactionReceiptOutput, error)
	GetTransactionsForOrdering(input *GetTransactionsForOrderingInput) (*GetTransactionsForOrderingOutput, error)
	ValidateTransactionsForOrdering(input *ValidateTransactionsForOrderingInput) (*ValidateTransactionsForOrderingOutput, error)
	CommitTransactionReceipts(input *CommitTransactionReceiptsInput) (*CommitTransactionReceiptsOutput, error)
	RegisterTransactionResultsHandler(handler handlers.TransactionResultsHandler)
}

/////////////////////////////////////////////////////////////////////////////
// message AddNewTransactionInput (non serializable)

type AddNewTransactionInput struct {
	SignedTransaction *protocol.SignedTransaction
}

/////////////////////////////////////////////////////////////////////////////
// message AddNewTransactionOutput (non serializable)

type AddNewTransactionOutput struct {
	TransactionStatus protocol.TransactionStatus
	TransactionReceipt *protocol.TransactionReceipt
	BlockHeight primitives.BlockHeight
	BlockTimestamp primitives.TimestampNano
}

/////////////////////////////////////////////////////////////////////////////
// message GetCommittedTransactionReceiptInput (non serializable)

type GetCommittedTransactionReceiptInput struct {
	Txhash primitives.Sha256
	TransactionTimestamp primitives.TimestampNano
}

/////////////////////////////////////////////////////////////////////////////
// message GetCommittedTransactionReceiptOutput (non serializable)

type GetCommittedTransactionReceiptOutput struct {
	TransactionStatus protocol.TransactionStatus
	TransactionReceipt *protocol.TransactionReceipt
	BlockHeight primitives.BlockHeight
	BlockTimestamp primitives.TimestampNano
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsForOrderingInput (non serializable)

type GetTransactionsForOrderingInput struct {
	BlockHeight primitives.BlockHeight
	MaxTransactionsSetSizeKb uint32
	MaxNumberOfTransactions uint32
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsForOrderingOutput (non serializable)

type GetTransactionsForOrderingOutput struct {
	SignedTransactions []*protocol.SignedTransaction
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsForOrderingInput (non serializable)

type ValidateTransactionsForOrderingInput struct {
	BlockHeight primitives.BlockHeight
	SignedTransactions []*protocol.SignedTransaction
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsForOrderingOutput (non serializable)

type ValidateTransactionsForOrderingOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message CommitTransactionReceiptsInput (non serializable)

type CommitTransactionReceiptsInput struct {
	ResultsBlockHeader *protocol.ResultsBlockHeader
	TransactionReceipts []*protocol.TransactionReceipt
	LastCommittedBlockHeight primitives.BlockHeight
}

/////////////////////////////////////////////////////////////////////////////
// message CommitTransactionReceiptsOutput (non serializable)

type CommitTransactionReceiptsOutput struct {
	NextDesiredBlockHeight primitives.BlockHeight
	LastCommittedBlockHeight primitives.BlockHeight
}

/////////////////////////////////////////////////////////////////////////////
// enums

