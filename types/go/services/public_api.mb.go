// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service PublicApi

type PublicApi interface {
	handlers.TransactionResultsHandler
	SendTransaction(input *SendTransactionInput) (*SendTransactionOutput, error)
	CallMethod(input *CallMethodInput) (*CallMethodOutput, error)
	GetTransactionStatus(input *GetTransactionStatusInput) (*GetTransactionStatusOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionInput (non serializable)

type SendTransactionInput struct {
	SignedTransaction *protocol.SignedTransaction
}

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionOutput (non serializable)

type SendTransactionOutput struct {
	Receipt *protocol.TransactionReceipt
	TransactionStatus protocol.TransactionStatus
	BlockHeight uint64
	BlockTimestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// message CallMethodInput (non serializable)

type CallMethodInput struct {
	Transaction *protocol.Transaction
}

/////////////////////////////////////////////////////////////////////////////
// message CallMethodOutput (non serializable)

type CallMethodOutput struct {
	OutputArgument []*protocol.MethodArgument
	CallStatus protocol.CallMethodStatus
	BlockHeight uint64
	BlockTimestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusInput (non serializable)

type GetTransactionStatusInput struct {
	Txid []byte
	Timestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusOutput (non serializable)

type GetTransactionStatusOutput struct {
	Receipt *protocol.TransactionReceipt
	Status protocol.TransactionStatus
	BlockHeight uint64
	BlockTimestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// enums

