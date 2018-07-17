// AUTO GENERATED FILE (by membufc proto compiler v0.0.16)
package handlers

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ContractSdkCallHandler

type ContractSdkCallHandler interface {
	HandleSdkCall(input *HandleSdkCallInput) (*HandleSdkCallOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleSdkCallInput (non serializable)

type HandleSdkCallInput struct {
	ContextId primitives.ExecutionContextId
	ContractName primitives.ContractName
	MethodName primitives.MethodName
	InputArguments []*protocol.MethodArgument
}

/////////////////////////////////////////////////////////////////////////////
// message HandleSdkCallOutput (non serializable)

type HandleSdkCallOutput struct {
	OutputArguments []*protocol.MethodArgument
}

/////////////////////////////////////////////////////////////////////////////
// enums

