// AUTO GENERATED FILE (by membufc proto compiler v0.0.14)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service VirtualMachine

type VirtualMachine interface {
	handlers.ContractSdkCallHandler
	ProcessTransactionSet(input *ProcessTransactionSetInput) (*ProcessTransactionSetOutput, error)
	RunLocalMethod(input *RunLocalMethodInput) (*RunLocalMethodOutput, error)
	TransactionSetPreOrder(input *TransactionSetPreOrderInput) (*TransactionSetPreOrderOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessTransactionSetInput (non serializable)

type ProcessTransactionSetInput struct {
	BlockHeight primitives.BlockHeight
	SignedTransactions []*protocol.SignedTransaction
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessTransactionSetOutput (non serializable)

type ProcessTransactionSetOutput struct {
	TransactionReceipts []*protocol.TransactionReceipt
	ContractStateDiffs []*protocol.ContractStateDiff
}

/////////////////////////////////////////////////////////////////////////////
// message RunLocalMethodInput (non serializable)

type RunLocalMethodInput struct {
	BlockHeight primitives.BlockHeight
	Transaction *protocol.Transaction
}

/////////////////////////////////////////////////////////////////////////////
// message RunLocalMethodOutput (non serializable)

type RunLocalMethodOutput struct {
	CallResult protocol.ExecutionResult
	OutputArguments []*protocol.MethodArgument
	ReferenceBlockHeight primitives.BlockHeight
	ReferenceBlockTimestamp primitives.Timestamp
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSetPreOrderInput (non serializable)

type TransactionSetPreOrderInput struct {
	BlockHeight primitives.BlockHeight
	SignedTransactions []*protocol.SignedTransaction
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSetPreOrderOutput (non serializable)

type TransactionSetPreOrderOutput struct {
	PreOrderResults []protocol.TransactionStatus
}

/////////////////////////////////////////////////////////////////////////////
// enums

