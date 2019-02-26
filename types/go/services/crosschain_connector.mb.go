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
	EthereumGetBlockNumber(ctx context.Context, input *EthereumGetBlockNumberInput) (*EthereumGetBlockNumberOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractInput (non serializable)

type EthereumCallContractInput struct {
	ReferenceTimestamp              primitives.TimestampNano
	EthereumBlockNumber             uint64
	EthereumContractAddress         string
	EthereumFunctionName            string
	EthereumJsonAbi                 string
	EthereumAbiPackedInputArguments []byte
}

func (x *EthereumCallContractInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ReferenceTimestamp:%s,EthereumBlockNumber:%s,EthereumContractAddress:%s,EthereumFunctionName:%s,EthereumJsonAbi:%s,EthereumAbiPackedInputArguments:%s,}", x.StringReferenceTimestamp(), x.StringEthereumBlockNumber(), x.StringEthereumContractAddress(), x.StringEthereumFunctionName(), x.StringEthereumJsonAbi(), x.StringEthereumAbiPackedInputArguments())
}

func (x *EthereumCallContractInput) StringReferenceTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceTimestamp)
	return
}

func (x *EthereumCallContractInput) StringEthereumBlockNumber() (res string) {
	res = fmt.Sprintf("%x", x.EthereumBlockNumber)
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
	EthereumTxhash          string
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
	res = fmt.Sprintf(x.EthereumTxhash)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetTransactionLogsOutput (non serializable)

type EthereumGetTransactionLogsOutput struct {
	EthereumAbiPackedOutputs [][]byte
	EthereumBlockNumber      uint64
	EthereumTxindex          uint32
}

func (x *EthereumGetTransactionLogsOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EthereumAbiPackedOutputs:%s,EthereumBlockNumber:%s,EthereumTxindex:%s,}", x.StringEthereumAbiPackedOutputs(), x.StringEthereumBlockNumber(), x.StringEthereumTxindex())
}

func (x *EthereumGetTransactionLogsOutput) StringEthereumAbiPackedOutputs() (res string) {
	res = "["
	for _, v := range x.EthereumAbiPackedOutputs {
		res += fmt.Sprintf("%x", v) + ","
	}
	res += "]"
	return
}

func (x *EthereumGetTransactionLogsOutput) StringEthereumBlockNumber() (res string) {
	res = fmt.Sprintf("%x", x.EthereumBlockNumber)
	return
}

func (x *EthereumGetTransactionLogsOutput) StringEthereumTxindex() (res string) {
	res = fmt.Sprintf("%x", x.EthereumTxindex)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetBlockNumberInput (non serializable)

type EthereumGetBlockNumberInput struct {
	ReferenceTimestamp primitives.TimestampNano
}

func (x *EthereumGetBlockNumberInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ReferenceTimestamp:%s,}", x.StringReferenceTimestamp())
}

func (x *EthereumGetBlockNumberInput) StringReferenceTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetBlockNumberOutput (non serializable)

type EthereumGetBlockNumberOutput struct {
	EthereumBlockNumber uint64
}

func (x *EthereumGetBlockNumberOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EthereumBlockNumber:%s,}", x.StringEthereumBlockNumber())
}

func (x *EthereumGetBlockNumberOutput) StringEthereumBlockNumber() (res string) {
	res = fmt.Sprintf("%x", x.EthereumBlockNumber)
	return
}
