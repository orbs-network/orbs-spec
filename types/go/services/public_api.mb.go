// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package services

import (
	"fmt"
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

func (x *SendTransactionInput) String() string {
	return fmt.Sprintf("{ClientRequest:%s,}", x.StringClientRequest())
}

func (x *SendTransactionInput) StringClientRequest() (res string) {
	res = x.ClientRequest.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionOutput (non serializable)

type SendTransactionOutput struct {
	ClientResponse *client.SendTransactionResponse
}

func (x *SendTransactionOutput) String() string {
	return fmt.Sprintf("{ClientResponse:%s,}", x.StringClientResponse())
}

func (x *SendTransactionOutput) StringClientResponse() (res string) {
	res = x.ClientResponse.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CallMethodInput (non serializable)

type CallMethodInput struct {
	ClientRequest *client.CallMethodRequest
}

func (x *CallMethodInput) String() string {
	return fmt.Sprintf("{ClientRequest:%s,}", x.StringClientRequest())
}

func (x *CallMethodInput) StringClientRequest() (res string) {
	res = x.ClientRequest.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CallMethodOutput (non serializable)

type CallMethodOutput struct {
	ClientResponse *client.CallMethodResponse
}

func (x *CallMethodOutput) String() string {
	return fmt.Sprintf("{ClientResponse:%s,}", x.StringClientResponse())
}

func (x *CallMethodOutput) StringClientResponse() (res string) {
	res = x.ClientResponse.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusInput (non serializable)

type GetTransactionStatusInput struct {
	ClientRequest *client.GetTransactionStatusRequest
}

func (x *GetTransactionStatusInput) String() string {
	return fmt.Sprintf("{ClientRequest:%s,}", x.StringClientRequest())
}

func (x *GetTransactionStatusInput) StringClientRequest() (res string) {
	res = x.ClientRequest.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusOutput (non serializable)

type GetTransactionStatusOutput struct {
	ClientResponse *client.GetTransactionStatusResponse
}

func (x *GetTransactionStatusOutput) String() string {
	return fmt.Sprintf("{ClientResponse:%s,}", x.StringClientResponse())
}

func (x *GetTransactionStatusOutput) StringClientResponse() (res string) {
	res = x.ClientResponse.String()
	return
}

