// AUTO GENERATED FILE (by membufc proto compiler v0.0.11)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionPool

type TransactionPool interface {
	handlers.TransactionRelayGossipHandler
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
	BlockHeight uint64
	BlockTimestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// message GetCommittedTransactionReceiptInput (non serializable)

type GetCommittedTransactionReceiptInput struct {
	Txid []byte
	Timestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// message GetCommittedTransactionReceiptOutput (non serializable)

type GetCommittedTransactionReceiptOutput struct {
	TransactionStatus protocol.TransactionStatus
	TransactionReceipt *protocol.TransactionReceipt
	BlockHeight uint64
	BlockTimestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsForOrderingInput (non serializable)

type GetTransactionsForOrderingInput struct {
	BlockHeight uint64
	MaxTransactionsSetSizeKb uint32
	MaxNumberOfTransactions uint32
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsForOrderingOutput (non serializable)

type GetTransactionsForOrderingOutput struct {
	Status protocol.RequestStatus
	SignedTransaction []*protocol.SignedTransaction
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsForOrderingInput (non serializable)

type ValidateTransactionsForOrderingInput struct {
	BlockHeight uint64
	Transaction []*protocol.SignedTransaction
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsForOrderingOutput (non serializable)

type ValidateTransactionsForOrderingOutput struct {
	Status protocol.RequestStatus
}

/////////////////////////////////////////////////////////////////////////////
// message CommitTransactionReceiptsInput (non serializable)

type CommitTransactionReceiptsInput struct {
	ResultsBlockHeader *protocol.ResultsBlockHeader
	TransactionReceipt []*protocol.TransactionReceipt
	LastCommittedBlockHeight uint64
}

/////////////////////////////////////////////////////////////////////////////
// message CommitTransactionReceiptsOutput (non serializable)

type CommitTransactionReceiptsOutput struct {
	NextDesiredBlockHeight uint64
	LastCommittedBlockHeight uint64
}

/////////////////////////////////////////////////////////////////////////////
// enums

