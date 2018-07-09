// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
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
	ClientRequest *client.SendTransactionRequest
}

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionOutput (non serializable)

type SendTransactionOutput struct {
	ClientResponse *client.SendTransactionResponse
}

/////////////////////////////////////////////////////////////////////////////
// message CallMethodInput (non serializable)

type CallMethodInput struct {
	ClientRequest *client.CallMethodRequest
}

/////////////////////////////////////////////////////////////////////////////
// message CallMethodOutput (non serializable)

type CallMethodOutput struct {
	ClientResponse *client.CallMethodResponse
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusInput (non serializable)

type GetTransactionStatusInput struct {
	ClientRequest *client.GetTransactionStatusRequest
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusOutput (non serializable)

type GetTransactionStatusOutput struct {
	ClientResponse *client.GetTransactionStatusResponse
}

