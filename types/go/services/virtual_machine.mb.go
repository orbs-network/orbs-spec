// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
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
	RunLocalMethod(ctx context.Context, input *RunLocalMethodInput) (*RunLocalMethodOutput, error)
	TransactionSetPreOrder(ctx context.Context, input *TransactionSetPreOrderInput) (*TransactionSetPreOrderOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessTransactionSetInput (non serializable)

type ProcessTransactionSetInput struct {
	BlockHeight        primitives.BlockHeight
	BlockTimestamp     primitives.TimestampNano
	SignedTransactions []*protocol.SignedTransaction
}

func (x *ProcessTransactionSetInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,BlockTimestamp:%s,SignedTransactions:%s,}", x.StringBlockHeight(), x.StringBlockTimestamp(), x.StringSignedTransactions())
}

func (x *ProcessTransactionSetInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *ProcessTransactionSetInput) StringBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.BlockTimestamp)
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
// message RunLocalMethodInput (non serializable)

type RunLocalMethodInput struct {
	BlockHeight primitives.BlockHeight
	Transaction *protocol.Transaction
}

func (x *RunLocalMethodInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,Transaction:%s,}", x.StringBlockHeight(), x.StringTransaction())
}

func (x *RunLocalMethodInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *RunLocalMethodInput) StringTransaction() (res string) {
	res = x.Transaction.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message RunLocalMethodOutput (non serializable)

type RunLocalMethodOutput struct {
	CallResult              protocol.ExecutionResult
	OutputArgumentArray     primitives.PackedArgumentArray
	OutputEventsArray       primitives.PackedEventsArray
	ReferenceBlockHeight    primitives.BlockHeight
	ReferenceBlockTimestamp primitives.TimestampNano
}

func (x *RunLocalMethodOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CallResult:%s,OutputArgumentArray:%s,OutputEventsArray:%s,ReferenceBlockHeight:%s,ReferenceBlockTimestamp:%s,}", x.StringCallResult(), x.StringOutputArgumentArray(), x.StringOutputEventsArray(), x.StringReferenceBlockHeight(), x.StringReferenceBlockTimestamp())
}

func (x *RunLocalMethodOutput) StringCallResult() (res string) {
	res = fmt.Sprintf("%x", x.CallResult)
	return
}

func (x *RunLocalMethodOutput) StringOutputArgumentArray() (res string) {
	res = fmt.Sprintf("%s", x.OutputArgumentArray)
	return
}

func (x *RunLocalMethodOutput) StringOutputEventsArray() (res string) {
	res = fmt.Sprintf("%s", x.OutputEventsArray)
	return
}

func (x *RunLocalMethodOutput) StringReferenceBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceBlockHeight)
	return
}

func (x *RunLocalMethodOutput) StringReferenceBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceBlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSetPreOrderInput (non serializable)

type TransactionSetPreOrderInput struct {
	BlockHeight        primitives.BlockHeight
	BlockTimestamp     primitives.TimestampNano
	SignedTransactions []*protocol.SignedTransaction
}

func (x *TransactionSetPreOrderInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,BlockTimestamp:%s,SignedTransactions:%s,}", x.StringBlockHeight(), x.StringBlockTimestamp(), x.StringSignedTransactions())
}

func (x *TransactionSetPreOrderInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *TransactionSetPreOrderInput) StringBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.BlockTimestamp)
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
// enums
