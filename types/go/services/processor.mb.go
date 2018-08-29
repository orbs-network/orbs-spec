// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package services

import (
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service Processor

type Processor interface {
	ProcessCall(input *ProcessCallInput) (*ProcessCallOutput, error)
	GetContractInfo(input *GetContractInfoInput) (*GetContractInfoOutput, error)
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

func (x *ProcessCallInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContextId:%s,ContractName:%s,MethodName:%s,InputArguments:%s,AccessScope:%s,PermissionScope:%s,CallingService:%s,TransactionSigner:%s,}", x.StringContextId(), x.StringContractName(), x.StringMethodName(), x.StringInputArguments(), x.StringAccessScope(), x.StringPermissionScope(), x.StringCallingService(), x.StringTransactionSigner())
}

func (x *ProcessCallInput) StringContextId() (res string) {
	res = fmt.Sprintf("%s", x.ContextId)
	return
}

func (x *ProcessCallInput) StringContractName() (res string) {
	res = fmt.Sprintf("%s", x.ContractName)
	return
}

func (x *ProcessCallInput) StringMethodName() (res string) {
	res = fmt.Sprintf("%s", x.MethodName)
	return
}

func (x *ProcessCallInput) StringInputArguments() (res string) {
	res = "["
		for _, v := range x.InputArguments {
		res += v.String() + ","
  }
	res += "]"
	return
}

func (x *ProcessCallInput) StringAccessScope() (res string) {
	res = fmt.Sprintf("%x", x.AccessScope)
	return
}

func (x *ProcessCallInput) StringPermissionScope() (res string) {
	res = fmt.Sprintf("%x", x.PermissionScope)
	return
}

func (x *ProcessCallInput) StringCallingService() (res string) {
	res = fmt.Sprintf("%s", x.CallingService)
	return
}

func (x *ProcessCallInput) StringTransactionSigner() (res string) {
	res = x.TransactionSigner.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessCallOutput (non serializable)

type ProcessCallOutput struct {
	OutputArguments []*protocol.MethodArgument
	CallResult protocol.ExecutionResult
}

func (x *ProcessCallOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{OutputArguments:%s,CallResult:%s,}", x.StringOutputArguments(), x.StringCallResult())
}

func (x *ProcessCallOutput) StringOutputArguments() (res string) {
	res = "["
		for _, v := range x.OutputArguments {
		res += v.String() + ","
  }
	res += "]"
	return
}

func (x *ProcessCallOutput) StringCallResult() (res string) {
	res = fmt.Sprintf("%x", x.CallResult)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetContractInfoInput (non serializable)

type GetContractInfoInput struct {
	ContractName primitives.ContractName
}

func (x *GetContractInfoInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContractName:%s,}", x.StringContractName())
}

func (x *GetContractInfoInput) StringContractName() (res string) {
	res = fmt.Sprintf("%s", x.ContractName)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetContractInfoOutput (non serializable)

type GetContractInfoOutput struct {
	PermissionScope protocol.ExecutionPermissionScope
}

func (x *GetContractInfoOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{PermissionScope:%s,}", x.StringPermissionScope())
}

func (x *GetContractInfoOutput) StringPermissionScope() (res string) {
	res = fmt.Sprintf("%x", x.PermissionScope)
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums

