// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package handlers

import (
	"fmt"
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

func (x *HandleSdkCallInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContextId:%s,ContractName:%s,MethodName:%s,InputArguments:%s,}", x.StringContextId(), x.StringContractName(), x.StringMethodName(), x.StringInputArguments())
}

func (x *HandleSdkCallInput) StringContextId() (res string) {
	res = fmt.Sprintf("%x", x.ContextId)
	return
}

func (x *HandleSdkCallInput) StringContractName() (res string) {
	res = fmt.Sprintf("%x", x.ContractName)
	return
}

func (x *HandleSdkCallInput) StringMethodName() (res string) {
	res = fmt.Sprintf("%x", x.MethodName)
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

/////////////////////////////////////////////////////////////////////////////
// message HandleSdkCallOutput (non serializable)

type HandleSdkCallOutput struct {
	OutputArguments []*protocol.MethodArgument
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

