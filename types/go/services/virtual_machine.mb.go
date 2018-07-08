// AUTO GENERATED FILE (by membufc proto compiler v0.0.11)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
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
	BlockHeight uint64
	SignedTransaction []*protocol.SignedTransaction
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessTransactionSetOutput (non serializable)

type ProcessTransactionSetOutput struct {
	Status protocol.RequestStatus
	TransactionReceipt []*protocol.TransactionReceipt
	ContractStateDiff []*protocol.ContractStateDiff
}

/////////////////////////////////////////////////////////////////////////////
// message RunLocalMethodInput (non serializable)

type RunLocalMethodInput struct {
	BlockHeight uint64
	Transaction *protocol.Transaction
}

/////////////////////////////////////////////////////////////////////////////
// message RunLocalMethodOutput (non serializable)

type RunLocalMethodOutput struct {
	Status protocol.CallMethodStatus
	OutputArgument []*protocol.MethodArgument
	ReferenceBlockHeight uint64
	ReferenceBlockTimestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSetPreOrderInput (non serializable)

type TransactionSetPreOrderInput struct {
	BlockHeight uint64
	SignedTransaction []*protocol.SignedTransaction
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSetPreOrderOutput (non serializable)

type TransactionSetPreOrderOutput struct {
	Status protocol.RequestStatus
	PreOrderResult []*protocol.PreOrderResult
}

/////////////////////////////////////////////////////////////////////////////
// enums

