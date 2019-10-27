// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ManagementConnector

type ManagementConnector interface {
	GetBlockInfoByTime(ctx context.Context, input *GetBlockInfoByTimeInput) (*GetBlockInfoByTimeOutput, error)
	CallContract(ctx context.Context, input *CallContractInput) (*CallContractOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message GetBlockInfoByTimeInput (non serializable)

type GetBlockInfoByTimeInput struct {
	ReferenceTimestamp primitives.TimestampNano
}

func (x *GetBlockInfoByTimeInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ReferenceTimestamp:%s,}", x.StringReferenceTimestamp())
}

func (x *GetBlockInfoByTimeInput) StringReferenceTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetBlockInfoByTimeOutput (non serializable)

type GetBlockInfoByTimeOutput struct {
	CallResult     protocol.CrossChainCallResult
	BlockHeight    primitives.BlockHeight
	BlockTimestamp primitives.TimestampNano
}

func (x *GetBlockInfoByTimeOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CallResult:%s,BlockHeight:%s,BlockTimestamp:%s,}", x.StringCallResult(), x.StringBlockHeight(), x.StringBlockTimestamp())
}

func (x *GetBlockInfoByTimeOutput) StringCallResult() (res string) {
	res = fmt.Sprintf("%x", x.CallResult)
	return
}

func (x *GetBlockInfoByTimeOutput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *GetBlockInfoByTimeOutput) StringBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.BlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CallContractInput (non serializable)

type CallContractInput struct {
	BlockHeight        primitives.BlockHeight
	ContractName       primitives.ContractName
	MethodName         primitives.MethodName
	InputArgumentArray primitives.PackedArgumentArray
}

func (x *CallContractInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,ContractName:%s,MethodName:%s,InputArgumentArray:%s,}", x.StringBlockHeight(), x.StringContractName(), x.StringMethodName(), x.StringInputArgumentArray())
}

func (x *CallContractInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *CallContractInput) StringContractName() (res string) {
	res = fmt.Sprintf("%s", x.ContractName)
	return
}

func (x *CallContractInput) StringMethodName() (res string) {
	res = fmt.Sprintf("%s", x.MethodName)
	return
}

func (x *CallContractInput) StringInputArgumentArray() (res string) {
	res = fmt.Sprintf("%s", x.InputArgumentArray)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CallContractOutput (non serializable)

type CallContractOutput struct {
	CallResult          protocol.CrossChainCallResult
	OutputArgumentArray primitives.PackedArgumentArray
}

func (x *CallContractOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CallResult:%s,OutputArgumentArray:%s,}", x.StringCallResult(), x.StringOutputArgumentArray())
}

func (x *CallContractOutput) StringCallResult() (res string) {
	res = fmt.Sprintf("%x", x.CallResult)
	return
}

func (x *CallContractOutput) StringOutputArgumentArray() (res string) {
	res = fmt.Sprintf("%s", x.OutputArgumentArray)
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
