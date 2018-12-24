// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
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
	EthereumGetTransactionLogs(ctx context.Context, input *EthereumGetTransactionLogsInput) (*EthereumGetTransactionLogsOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractInput (non serializable)

type EthereumCallContractInput struct {
	ReferenceTimestamp              primitives.TimestampNano
	EthereumContractAddress         string
	EthereumFunctionName            string
	EthereumJsonAbi                 string
	EthereumAbiPackedInputArguments []byte
}

func (x *EthereumCallContractInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ReferenceTimestamp:%s,EthereumContractAddress:%s,EthereumFunctionName:%s,EthereumJsonAbi:%s,EthereumAbiPackedInputArguments:%s,}", x.StringReferenceTimestamp(), x.StringEthereumContractAddress(), x.StringEthereumFunctionName(), x.StringEthereumJsonAbi(), x.StringEthereumAbiPackedInputArguments())
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

func (x *EthereumCallContractInput) StringEthereumJsonAbi() (res string) {
	res = fmt.Sprintf(x.EthereumJsonAbi)
	return
}

func (x *EthereumCallContractInput) StringEthereumAbiPackedInputArguments() (res string) {
	res = fmt.Sprintf("%x", x.EthereumAbiPackedInputArguments)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractOutput (non serializable)

type EthereumCallContractOutput struct {
	EthereumAbiPackedOutput []byte
}

func (x *EthereumCallContractOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EthereumAbiPackedOutput:%s,}", x.StringEthereumAbiPackedOutput())
}

func (x *EthereumCallContractOutput) StringEthereumAbiPackedOutput() (res string) {
	res = fmt.Sprintf("%x", x.EthereumAbiPackedOutput)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetTransactionLogsInput (non serializable)

type EthereumGetTransactionLogsInput struct {
	ReferenceTimestamp      primitives.TimestampNano
	EthereumContractAddress string
	EthereumEventName       string
	EthereumJsonAbi         string
	EthereumTxhash          primitives.Uint256
}

func (x *EthereumGetTransactionLogsInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ReferenceTimestamp:%s,EthereumContractAddress:%s,EthereumEventName:%s,EthereumJsonAbi:%s,EthereumTxhash:%s,}", x.StringReferenceTimestamp(), x.StringEthereumContractAddress(), x.StringEthereumEventName(), x.StringEthereumJsonAbi(), x.StringEthereumTxhash())
}

func (x *EthereumGetTransactionLogsInput) StringReferenceTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceTimestamp)
	return
}

func (x *EthereumGetTransactionLogsInput) StringEthereumContractAddress() (res string) {
	res = fmt.Sprintf(x.EthereumContractAddress)
	return
}

func (x *EthereumGetTransactionLogsInput) StringEthereumEventName() (res string) {
	res = fmt.Sprintf(x.EthereumEventName)
	return
}

func (x *EthereumGetTransactionLogsInput) StringEthereumJsonAbi() (res string) {
	res = fmt.Sprintf(x.EthereumJsonAbi)
	return
}

func (x *EthereumGetTransactionLogsInput) StringEthereumTxhash() (res string) {
	res = fmt.Sprintf("%s", x.EthereumTxhash)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetTransactionLogsOutput (non serializable)

type EthereumGetTransactionLogsOutput struct {
	EthereumAbiPackedOutput []byte
}

func (x *EthereumGetTransactionLogsOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EthereumAbiPackedOutput:%s,}", x.StringEthereumAbiPackedOutput())
}

func (x *EthereumGetTransactionLogsOutput) StringEthereumAbiPackedOutput() (res string) {
	res = fmt.Sprintf("%x", x.EthereumAbiPackedOutput)
	return
}
