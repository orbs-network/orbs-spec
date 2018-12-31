// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package handlers

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ContractSdkCallHandler

type ContractSdkCallHandler interface {
	HandleSdkCall(ctx context.Context, input *HandleSdkCallInput) (*HandleSdkCallOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleSdkCallInput (non serializable)

type HandleSdkCallInput struct {
	ContextId       primitives.ExecutionContextId
	OperationName   primitives.ContractName
	MethodName      primitives.MethodName
	InputArguments  []*protocol.Argument
	PermissionScope protocol.ExecutionPermissionScope
}

func (x *HandleSdkCallInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContextId:%s,OperationName:%s,MethodName:%s,InputArguments:%s,PermissionScope:%s,}", x.StringContextId(), x.StringOperationName(), x.StringMethodName(), x.StringInputArguments(), x.StringPermissionScope())
}

func (x *HandleSdkCallInput) StringContextId() (res string) {
	res = fmt.Sprintf("%s", x.ContextId)
	return
}

func (x *HandleSdkCallInput) StringOperationName() (res string) {
	res = fmt.Sprintf("%s", x.OperationName)
	return
}

func (x *HandleSdkCallInput) StringMethodName() (res string) {
	res = fmt.Sprintf("%s", x.MethodName)
	return
}

func (x *HandleSdkCallInput) StringInputArguments() (res string) {
	res = "["
	for _, v := range x.InputArguments {
		res += v.String() + ","
	}
	res += "]"
	return
}

func (x *HandleSdkCallInput) StringPermissionScope() (res string) {
	res = fmt.Sprintf("%x", x.PermissionScope)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HandleSdkCallOutput (non serializable)

type HandleSdkCallOutput struct {
	OutputArguments []*protocol.Argument
}

func (x *HandleSdkCallOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{OutputArguments:%s,}", x.StringOutputArguments())
}

func (x *HandleSdkCallOutput) StringOutputArguments() (res string) {
	res = "["
	for _, v := range x.OutputArguments {
		res += v.String() + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
