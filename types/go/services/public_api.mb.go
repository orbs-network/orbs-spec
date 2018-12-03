// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service PublicApi

type PublicApi interface {
	handlers.TransactionResultsHandler
	SendTransaction(ctx context.Context, input *SendTransactionInput) (*SendTransactionOutput, error)
	CallMethod(ctx context.Context, input *CallMethodInput) (*CallMethodOutput, error)
	GetTransactionStatus(ctx context.Context, input *GetTransactionStatusInput) (*GetTransactionStatusOutput, error)
	GetTransactionReceiptProof(ctx context.Context, input *GetTransactionReceiptProofInput) (*GetTransactionReceiptProofOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionInput (non serializable)

type SendTransactionInput struct {
	ClientRequest     *client.SendTransactionRequest
	ReturnImmediately uint32
}

func (x *SendTransactionInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ClientRequest:%s,ReturnImmediately:%s,}", x.StringClientRequest(), x.StringReturnImmediately())
}

func (x *SendTransactionInput) StringClientRequest() (res string) {
	res = x.ClientRequest.String()
	return
}

func (x *SendTransactionInput) StringReturnImmediately() (res string) {
	res = fmt.Sprintf("%x", x.ReturnImmediately)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionOutput (non serializable)

type SendTransactionOutput struct {
	ClientResponse *client.SendTransactionResponse
}

func (x *SendTransactionOutput) String() string {
	if x == nil {
		return "<nil>"
	}
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
	if x == nil {
		return "<nil>"
	}
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
	if x == nil {
		return "<nil>"
	}
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
	if x == nil {
		return "<nil>"
	}
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
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ClientResponse:%s,}", x.StringClientResponse())
}

func (x *GetTransactionStatusOutput) StringClientResponse() (res string) {
	res = x.ClientResponse.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionReceiptProofInput (non serializable)

type GetTransactionReceiptProofInput struct {
	ClientRequest *client.GetTransactionStatusRequest
}

func (x *GetTransactionReceiptProofInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ClientRequest:%s,}", x.StringClientRequest())
}

func (x *GetTransactionReceiptProofInput) StringClientRequest() (res string) {
	res = x.ClientRequest.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionReceiptProofOutput (non serializable)

type GetTransactionReceiptProofOutput struct {
	ClientResponse *client.GetTransactionReceiptProofResponse
}

func (x *GetTransactionReceiptProofOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ClientResponse:%s,}", x.StringClientResponse())
}

func (x *GetTransactionReceiptProofOutput) StringClientResponse() (res string) {
	res = x.ClientResponse.String()
	return
}
