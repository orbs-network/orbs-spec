// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package services

import (
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service CrosschainConnector

type CrosschainConnector interface {
	EthereumCallContract(input *EthereumCallContractInput) (*EthereumCallContractOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractInput (non serializable)

type EthereumCallContractInput struct {
	BlockHeight primitives.BlockHeight
	EthereumContractAddress string
	EthereumFunctionName string
	EthereumAbi string
	InputArguments []*protocol.MethodArgument
}

func (x *EthereumCallContractInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,EthereumContractAddress:%s,EthereumFunctionName:%s,EthereumAbi:%s,InputArguments:%s,}", x.StringBlockHeight(), x.StringEthereumContractAddress(), x.StringEthereumFunctionName(), x.StringEthereumAbi(), x.StringInputArguments())
}

func (x *EthereumCallContractInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *EthereumCallContractInput) StringEthereumContractAddress() (res string) {
	res = fmt.Sprintf(x.EthereumContractAddress)
	return
}

func (x *EthereumCallContractInput) StringEthereumFunctionName() (res string) {
	res = fmt.Sprintf(x.EthereumFunctionName)
	return
}

func (x *EthereumCallContractInput) StringEthereumAbi() (res string) {
	res = fmt.Sprintf(x.EthereumAbi)
	return
}

func (x *EthereumCallContractInput) StringInputArguments() (res string) {
	res = "["
		for _, v := range x.InputArguments {
		res += v.String() + ","
  }
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractOutput (non serializable)

type EthereumCallContractOutput struct {
	OutputArguments []*protocol.MethodArgument
	CallResult protocol.ExecutionResult
	EthereumBlockHeight uint64
	EthereumBlockTimestamp uint64
}

func (x *EthereumCallContractOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{OutputArguments:%s,CallResult:%s,EthereumBlockHeight:%s,EthereumBlockTimestamp:%s,}", x.StringOutputArguments(), x.StringCallResult(), x.StringEthereumBlockHeight(), x.StringEthereumBlockTimestamp())
}

func (x *EthereumCallContractOutput) StringOutputArguments() (res string) {
	res = "["
		for _, v := range x.OutputArguments {
		res += v.String() + ","
  }
	res += "]"
	return
}

func (x *EthereumCallContractOutput) StringCallResult() (res string) {
	res = fmt.Sprintf("%x", x.CallResult)
	return
}

func (x *EthereumCallContractOutput) StringEthereumBlockHeight() (res string) {
	res = fmt.Sprintf("%x", x.EthereumBlockHeight)
	return
}

func (x *EthereumCallContractOutput) StringEthereumBlockTimestamp() (res string) {
	res = fmt.Sprintf("%x", x.EthereumBlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums

