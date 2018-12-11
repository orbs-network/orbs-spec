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
	RunQuery(ctx context.Context, input *RunQueryInput) (*RunQueryOutput, error)
	GetTransactionStatus(ctx context.Context, input *GetTransactionStatusInput) (*GetTransactionStatusOutput, error)
	GetTransactionProof(ctx context.Context, input *GetTransactionProofInput) (*GetTransactionProofOutput, error)
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
// message RunQueryInput (non serializable)

type RunQueryInput struct {
	ClientRequest *client.RunQueryRequest
}

func (x *RunQueryInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ClientRequest:%s,}", x.StringClientRequest())
}

func (x *RunQueryInput) StringClientRequest() (res string) {
	res = x.ClientRequest.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message RunQueryOutput (non serializable)

type RunQueryOutput struct {
	ClientResponse *client.RunQueryResponse
}

func (x *RunQueryOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ClientResponse:%s,}", x.StringClientResponse())
}

func (x *RunQueryOutput) StringClientResponse() (res string) {
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
// message GetTransactionProofInput (non serializable)

type GetTransactionProofInput struct {
	ClientRequest *client.GetTransactionProofRequest
}

func (x *GetTransactionProofInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ClientRequest:%s,}", x.StringClientRequest())
}

func (x *GetTransactionProofInput) StringClientRequest() (res string) {
	res = x.ClientRequest.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionProofOutput (non serializable)

type GetTransactionProofOutput struct {
	ClientResponse *client.GetTransactionProofResponse
}

func (x *GetTransactionProofOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ClientResponse:%s,}", x.StringClientResponse())
}

func (x *GetTransactionProofOutput) StringClientResponse() (res string) {
	res = x.ClientResponse.String()
	return
}
