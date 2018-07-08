// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
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
	ContextId uint32
	Contract *protocol.ContractAddress
	Method *protocol.MethodAddress
	InputArgument []*protocol.MethodArgument
	AccessScope protocol.AccessScope
	PermissionScope protocol.PermissionScope
	CallingService *protocol.ContractAddress
	TransactionSender *protocol.Sender
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessCallOutput (non serializable)

type ProcessCallOutput struct {
	OutputArgument []*protocol.MethodArgument
	CallStatus protocol.CallMethodStatus
}

/////////////////////////////////////////////////////////////////////////////
// message DeployNativeServiceInput (non serializable)

type DeployNativeServiceInput struct {
	ContextId uint32
	Contract *protocol.ContractAddress
	AccessScope protocol.AccessScope
	PermissionScope protocol.PermissionScope
	CallingService *protocol.ContractAddress
	TransactionSender *protocol.Sender
}

/////////////////////////////////////////////////////////////////////////////
// message DeployNativeServiceOutput (non serializable)

type DeployNativeServiceOutput struct {
	Status protocol.DeployServiceStatus
}

/////////////////////////////////////////////////////////////////////////////
// enums

