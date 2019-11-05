// AUTO GENERATED FILE (by membufc proto compiler v0.3.6)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service VirtualMachine

type VirtualMachine interface {
	handlers.ContractSdkCallHandler
	ProcessTransactionSet(ctx context.Context, input *ProcessTransactionSetInput) (*ProcessTransactionSetOutput, error)
	CallSystemContract(ctx context.Context, input *CallSystemContractInput) (*CallSystemContractOutput, error)
	ProcessQuery(ctx context.Context, input *ProcessQueryInput) (*ProcessQueryOutput, error)
	TransactionSetPreOrder(ctx context.Context, input *TransactionSetPreOrderInput) (*TransactionSetPreOrderOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessTransactionSetInput (non serializable)

type ProcessTransactionSetInput struct {
	CurrentBlockHeight    primitives.BlockHeight
	CurrentBlockTimestamp primitives.TimestampNano
	SignedTransactions    []*protocol.SignedTransaction
	BlockProposerAddress  primitives.NodeAddress
}

func (x *ProcessTransactionSetInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,CurrentBlockTimestamp:%s,SignedTransactions:%s,BlockProposerAddress:%s,}", x.StringCurrentBlockHeight(), x.StringCurrentBlockTimestamp(), x.StringSignedTransactions(), x.StringBlockProposerAddress())
}

func (x *ProcessTransactionSetInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *ProcessTransactionSetInput) StringCurrentBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockTimestamp)
	return
}

func (x *ProcessTransactionSetInput) StringSignedTransactions() (res string) {
	res = "["
	for _, v := range x.SignedTransactions {
		res += v.String() + ","
	}
	res += "]"
	return
}

func (x *ProcessTransactionSetInput) StringBlockProposerAddress() (res string) {
	res = fmt.Sprintf("%s", x.BlockProposerAddress)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessTransactionSetOutput (non serializable)

type ProcessTransactionSetOutput struct {
	TransactionReceipts []*protocol.TransactionReceipt
	ContractStateDiffs  []*protocol.ContractStateDiff
}

func (x *ProcessTransactionSetOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionReceipts:%s,ContractStateDiffs:%s,}", x.StringTransactionReceipts(), x.StringContractStateDiffs())
}

func (x *ProcessTransactionSetOutput) StringTransactionReceipts() (res string) {
	res = "["
	for _, v := range x.TransactionReceipts {
		res += v.String() + ","
	}
	res += "]"
	return
}

func (x *ProcessTransactionSetOutput) StringContractStateDiffs() (res string) {
	res = "["
	for _, v := range x.ContractStateDiffs {
		res += v.String() + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessQueryInput (non serializable)

type ProcessQueryInput struct {
	BlockHeight primitives.BlockHeight
	SignedQuery *protocol.SignedQuery
}

func (x *ProcessQueryInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,SignedQuery:%s,}", x.StringBlockHeight(), x.StringSignedQuery())
}

func (x *ProcessQueryInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *ProcessQueryInput) StringSignedQuery() (res string) {
	res = x.SignedQuery.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessQueryOutput (non serializable)

type ProcessQueryOutput struct {
	CallResult              protocol.ExecutionResult
	OutputArgumentArray     primitives.PackedArgumentArray
	OutputEventsArray       primitives.PackedEventsArray
	ReferenceBlockHeight    primitives.BlockHeight
	ReferenceBlockTimestamp primitives.TimestampNano
}

func (x *ProcessQueryOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CallResult:%s,OutputArgumentArray:%s,OutputEventsArray:%s,ReferenceBlockHeight:%s,ReferenceBlockTimestamp:%s,}", x.StringCallResult(), x.StringOutputArgumentArray(), x.StringOutputEventsArray(), x.StringReferenceBlockHeight(), x.StringReferenceBlockTimestamp())
}

func (x *ProcessQueryOutput) StringCallResult() (res string) {
	res = fmt.Sprintf("%x", x.CallResult)
	return
}

func (x *ProcessQueryOutput) StringOutputArgumentArray() (res string) {
	res = fmt.Sprintf("%s", x.OutputArgumentArray)
	return
}

func (x *ProcessQueryOutput) StringOutputEventsArray() (res string) {
	res = fmt.Sprintf("%s", x.OutputEventsArray)
	return
}

func (x *ProcessQueryOutput) StringReferenceBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceBlockHeight)
	return
}

func (x *ProcessQueryOutput) StringReferenceBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceBlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSetPreOrderInput (non serializable)

type TransactionSetPreOrderInput struct {
	CurrentBlockHeight    primitives.BlockHeight
	CurrentBlockTimestamp primitives.TimestampNano
	SignedTransactions    []*protocol.SignedTransaction
}

func (x *TransactionSetPreOrderInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,CurrentBlockTimestamp:%s,SignedTransactions:%s,}", x.StringCurrentBlockHeight(), x.StringCurrentBlockTimestamp(), x.StringSignedTransactions())
}

func (x *TransactionSetPreOrderInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *TransactionSetPreOrderInput) StringCurrentBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockTimestamp)
	return
}

func (x *TransactionSetPreOrderInput) StringSignedTransactions() (res string) {
	res = "["
	for _, v := range x.SignedTransactions {
		res += v.String() + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSetPreOrderOutput (non serializable)

type TransactionSetPreOrderOutput struct {
	PreOrderResults []protocol.TransactionStatus
}

func (x *TransactionSetPreOrderOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{PreOrderResults:%s,}", x.StringPreOrderResults())
}

func (x *TransactionSetPreOrderOutput) StringPreOrderResults() (res string) {
	res = "["
	for _, v := range x.PreOrderResults {
		res += fmt.Sprintf("%x", v) + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CallSystemContractInput (non serializable)

type CallSystemContractInput struct {
	BlockHeight        primitives.BlockHeight
	BlockTimestamp     primitives.TimestampNano
	ContractName       primitives.ContractName
	MethodName         primitives.MethodName
	InputArgumentArray *protocol.ArgumentArray
}

func (x *CallSystemContractInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,BlockTimestamp:%s,ContractName:%s,MethodName:%s,InputArgumentArray:%s,}", x.StringBlockHeight(), x.StringBlockTimestamp(), x.StringContractName(), x.StringMethodName(), x.StringInputArgumentArray())
}

func (x *CallSystemContractInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *CallSystemContractInput) StringBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.BlockTimestamp)
	return
}

func (x *CallSystemContractInput) StringContractName() (res string) {
	res = fmt.Sprintf("%s", x.ContractName)
	return
}

func (x *CallSystemContractInput) StringMethodName() (res string) {
	res = fmt.Sprintf("%s", x.MethodName)
	return
}

func (x *CallSystemContractInput) StringInputArgumentArray() (res string) {
	res = x.InputArgumentArray.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CallSystemContractOutput (non serializable)

type CallSystemContractOutput struct {
	OutputArgumentArray *protocol.ArgumentArray
	CallResult          protocol.ExecutionResult
}

func (x *CallSystemContractOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{OutputArgumentArray:%s,CallResult:%s,}", x.StringOutputArgumentArray(), x.StringCallResult())
}

func (x *CallSystemContractOutput) StringOutputArgumentArray() (res string) {
	res = x.OutputArgumentArray.String()
	return
}

func (x *CallSystemContractOutput) StringCallResult() (res string) {
	res = fmt.Sprintf("%x", x.CallResult)
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
