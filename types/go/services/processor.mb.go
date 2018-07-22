// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service Processor

type Processor interface {
	ProcessCall(input *ProcessCallInput) (*ProcessCallOutput, error)
	DeployNativeService(input *DeployNativeServiceInput) (*DeployNativeServiceOutput, error)
	RegisterContractSdkCallHandler(handler handlers.ContractSdkCallHandler)
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessCallInput (non serializable)

type ProcessCallInput struct {
	ContextId primitives.ExecutionContextId
	ContractName primitives.ContractName
	MethodName primitives.MethodName
	InputArguments []*protocol.MethodArgument
	AccessScope protocol.ExecutionAccessScope
	PermissionScope protocol.ExecutionPermissionScope
	CallingService primitives.ContractName
	TransactionSigner *protocol.Signer
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessCallOutput (non serializable)

type ProcessCallOutput struct {
	OutputArguments []*protocol.MethodArgument
	CallResult protocol.ExecutionResult
}

/////////////////////////////////////////////////////////////////////////////
// message DeployNativeServiceInput (non serializable)

type DeployNativeServiceInput struct {
	ContextId primitives.ExecutionContextId
	ContractName primitives.ContractName
	AccessScope protocol.ExecutionAccessScope
	PermissionScope protocol.ExecutionPermissionScope
	CallingService primitives.ContractName
	TransactionSigner *protocol.Signer
}

/////////////////////////////////////////////////////////////////////////////
// message DeployNativeServiceOutput (non serializable)

type DeployNativeServiceOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// enums

