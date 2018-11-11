// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// service CrosschainConnector

type CrosschainConnector interface {
	EthereumCallContract(ctx context.Context, input *EthereumCallContractInput) (*EthereumCallContractOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractInput (non serializable)

type EthereumCallContractInput struct {
	ReferenceTimestamp           primitives.TimestampNano
	EthereumContractAddress      string
	EthereumFunctionName         string
	EthereumAbi                  string
	EthereumPackedInputArguments []byte
}

func (x *EthereumCallContractInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ReferenceTimestamp:%s,EthereumContractAddress:%s,EthereumFunctionName:%s,EthereumAbi:%s,EthereumPackedInputArguments:%s,}", x.StringReferenceTimestamp(), x.StringEthereumContractAddress(), x.StringEthereumFunctionName(), x.StringEthereumAbi(), x.StringEthereumPackedInputArguments())
}

func (x *EthereumCallContractInput) StringReferenceTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceTimestamp)
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

func (x *EthereumCallContractInput) StringEthereumPackedInputArguments() (res string) {
	res = fmt.Sprintf("%x", x.EthereumPackedInputArguments)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractOutput (non serializable)

type EthereumCallContractOutput struct {
	EthereumPackedOutput []byte
}

func (x *EthereumCallContractOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EthereumPackedOutput:%s,}", x.StringEthereumPackedOutput())
}

func (x *EthereumCallContractOutput) StringEthereumPackedOutput() (res string) {
	res = fmt.Sprintf("%x", x.EthereumPackedOutput)
	return
}
