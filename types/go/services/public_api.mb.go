// AUTO GENERATED FILE (by membufc proto compiler v0.0.11)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol/publicapi"
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
	ClientInput *publicapi.SendTransactionInput
}

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionOutput (non serializable)

type SendTransactionOutput struct {
	ClientOutput *publicapi.SendTransactionOutput
}

/////////////////////////////////////////////////////////////////////////////
// message CallMethodInput (non serializable)

type CallMethodInput struct {
	ClientInput *publicapi.CallMethodInput
}

/////////////////////////////////////////////////////////////////////////////
// message CallMethodOutput (non serializable)

type CallMethodOutput struct {
	ClientOutput *publicapi.CallMethodOutput
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusInput (non serializable)

type GetTransactionStatusInput struct {
	ClientInput *publicapi.GetTransactionStatusInput
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusOutput (non serializable)

type GetTransactionStatusOutput struct {
	ClientOutput *publicapi.GetTransactionStatusOutput
}

