// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
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
	BlockProposerAddress  primitives.NodeAddress
	BlockReferenceTime    primitives.TimestampSeconds
	SignedTransactions    []*protocol.SignedTransaction
}

func (x *ProcessTransactionSetInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,CurrentBlockTimestamp:%s,BlockProposerAddress:%s,BlockReferenceTime:%s,SignedTransactions:%s,}", x.StringCurrentBlockHeight(), x.StringCurrentBlockTimestamp(), x.StringBlockProposerAddress(), x.StringBlockReferenceTime(), x.StringSignedTransactions())
}

func (x *ProcessTransactionSetInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *ProcessTransactionSetInput) StringCurrentBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockTimestamp)
	return
}

func (x *ProcessTransactionSetInput) StringBlockProposerAddress() (res string) {
	res = fmt.Sprintf("%s", x.BlockProposerAddress)
	return
}

func (x *ProcessTransactionSetInput) StringBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.BlockReferenceTime)
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
	CallResult           protocol.ExecutionResult
	OutputArgumentArray  primitives.PackedArgumentArray
	OutputEventsArray    primitives.PackedEventsArray
	ReferenceBlockHeight primitives.BlockHeight
	BlockReferenceTime   primitives.TimestampSeconds
}

func (x *ProcessQueryOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CallResult:%s,OutputArgumentArray:%s,OutputEventsArray:%s,ReferenceBlockHeight:%s,BlockReferenceTime:%s,}", x.StringCallResult(), x.StringOutputArgumentArray(), x.StringOutputEventsArray(), x.StringReferenceBlockHeight(), x.StringBlockReferenceTime())
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

func (x *ProcessQueryOutput) StringBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.BlockReferenceTime)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSetPreOrderInput (non serializable)

type TransactionSetPreOrderInput struct {
	CurrentBlockHeight        primitives.BlockHeight
	CurrentBlockTimestamp     primitives.TimestampNano
	CurrentBlockReferenceTime primitives.TimestampSeconds
	SignedTransactions        []*protocol.SignedTransaction
}

func (x *TransactionSetPreOrderInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,CurrentBlockTimestamp:%s,CurrentBlockReferenceTime:%s,SignedTransactions:%s,}", x.StringCurrentBlockHeight(), x.StringCurrentBlockTimestamp(), x.StringCurrentBlockReferenceTime(), x.StringSignedTransactions())
}

func (x *TransactionSetPreOrderInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *TransactionSetPreOrderInput) StringCurrentBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockTimestamp)
	return
}

func (x *TransactionSetPreOrderInput) StringCurrentBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockReferenceTime)
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
	BlockReferenceTime primitives.TimestampSeconds
	BlockReferenceTime primitives.TimestampSeconds
	ContractName       primitives.ContractName
	MethodName         primitives.MethodName
	InputArgumentArray *protocol.ArgumentArray
}

func (x *CallSystemContractInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,BlockTimestamp:%s,BlockReferenceTime:%s,BlockReferenceTime:%s,ContractName:%s,MethodName:%s,InputArgumentArray:%s,}", x.StringBlockHeight(), x.StringBlockTimestamp(), x.StringBlockReferenceTime(), x.StringBlockReferenceTime(), x.StringContractName(), x.StringMethodName(), x.StringInputArgumentArray())
}

func (x *CallSystemContractInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *CallSystemContractInput) StringBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.BlockTimestamp)
	return
}

func (x *CallSystemContractInput) StringBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.BlockReferenceTime)
	return
}

func (x *CallSystemContractInput) StringBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.BlockReferenceTime)
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
