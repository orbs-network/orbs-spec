// AUTO GENERATED FILE (by membufc proto compiler)
package handlers

import (
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
	ContextId uint32
	Contract *protocol.ContractAddress
	Method *protocol.MethodAddress
	InputArgument []*protocol.MethodArgument
}

/////////////////////////////////////////////////////////////////////////////
// message HandleSdkCallOutput (non serializable)

type HandleSdkCallOutput struct {
	OutputArgument []*protocol.MethodArgument
}

