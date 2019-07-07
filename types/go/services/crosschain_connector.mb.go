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
	EthereumGetBlockTime(ctx context.Context, input *EthereumGetBlockTimeInput) (*EthereumGetBlockTimeOutput, error)
	EthereumGetBlockTimeByNumber(ctx context.Context, input *EthereumGetBlockTimeByNumberInput) (*EthereumGetBlockTimeByNumberOutput, error)
	EthereumGetBlockNumberByTime(ctx context.Context, input *EthereumGetBlockNumberByTimeInput) (*EthereumGetBlockNumberByTimeOutput, error)
	EthereumFilterLogs(ctx context.Context, input *EthereumFilterLogsInput) (*EthereumFilterLogsOutput, error)
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

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetBlockTimeByNumberInput (non serializable)

type EthereumGetBlockTimeByNumberInput struct {
	ReferenceTimestamp  primitives.TimestampNano
	EthereumBlockNumber uint64
}

func (x *EthereumGetBlockTimeByNumberInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ReferenceTimestamp:%s,EthereumBlockNumber:%s,}", x.StringReferenceTimestamp(), x.StringEthereumBlockNumber())
}

func (x *EthereumGetBlockTimeByNumberInput) StringReferenceTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceTimestamp)
	return
}

func (x *EthereumGetBlockTimeByNumberInput) StringEthereumBlockNumber() (res string) {
	res = fmt.Sprintf("%x", x.EthereumBlockNumber)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetBlockTimeByNumberOutput (non serializable)

type EthereumGetBlockTimeByNumberOutput struct {
	EthereumTimestamp primitives.TimestampNano
}

func (x *EthereumGetBlockTimeByNumberOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EthereumTimestamp:%s,}", x.StringEthereumTimestamp())
}

func (x *EthereumGetBlockTimeByNumberOutput) StringEthereumTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.EthereumTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetBlockTimeInput (non serializable)

type EthereumGetBlockTimeInput struct {
	ReferenceTimestamp primitives.TimestampNano
}

func (x *EthereumGetBlockTimeInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ReferenceTimestamp:%s,}", x.StringReferenceTimestamp())
}

func (x *EthereumGetBlockTimeInput) StringReferenceTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetBlockTimeOutput (non serializable)

type EthereumGetBlockTimeOutput struct {
	EthereumTimestamp primitives.TimestampNano
}

func (x *EthereumGetBlockTimeOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EthereumTimestamp:%s,}", x.StringEthereumTimestamp())
}

func (x *EthereumGetBlockTimeOutput) StringEthereumTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.EthereumTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetBlockNumberByTimeInput (non serializable)

type EthereumGetBlockNumberByTimeInput struct {
	ReferenceTimestamp primitives.TimestampNano
	EthereumTimestamp  primitives.TimestampNano
}

func (x *EthereumGetBlockNumberByTimeInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ReferenceTimestamp:%s,EthereumTimestamp:%s,}", x.StringReferenceTimestamp(), x.StringEthereumTimestamp())
}

func (x *EthereumGetBlockNumberByTimeInput) StringReferenceTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceTimestamp)
	return
}

func (x *EthereumGetBlockNumberByTimeInput) StringEthereumTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.EthereumTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetBlockNumberByTimeOutput (non serializable)

type EthereumGetBlockNumberByTimeOutput struct {
	EthereumBlockNumber uint64
}

func (x *EthereumGetBlockNumberByTimeOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EthereumBlockNumber:%s,}", x.StringEthereumBlockNumber())
}

func (x *EthereumGetBlockNumberByTimeOutput) StringEthereumBlockNumber() (res string) {
	res = fmt.Sprintf("%x", x.EthereumBlockNumber)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumFilterLogsInput (non serializable)

type EthereumFilterLogsInput struct {
	ReferenceTimestamp      primitives.TimestampNano
	EthereumContractAddress string
	EthereumEventName       string
	EthereumJsonAbi         string
	FromEthereumBlockNumber uint64
	ToEthereumBlockNumber   uint64
}

func (x *EthereumFilterLogsInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ReferenceTimestamp:%s,EthereumContractAddress:%s,EthereumEventName:%s,EthereumJsonAbi:%s,FromEthereumBlockNumber:%s,ToEthereumBlockNumber:%s,}", x.StringReferenceTimestamp(), x.StringEthereumContractAddress(), x.StringEthereumEventName(), x.StringEthereumJsonAbi(), x.StringFromEthereumBlockNumber(), x.StringToEthereumBlockNumber())
}

func (x *EthereumFilterLogsInput) StringReferenceTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceTimestamp)
	return
}

func (x *EthereumFilterLogsInput) StringEthereumContractAddress() (res string) {
	res = fmt.Sprintf(x.EthereumContractAddress)
	return
}

func (x *EthereumFilterLogsInput) StringEthereumEventName() (res string) {
	res = fmt.Sprintf(x.EthereumEventName)
	return
}

func (x *EthereumFilterLogsInput) StringEthereumJsonAbi() (res string) {
	res = fmt.Sprintf(x.EthereumJsonAbi)
	return
}

func (x *EthereumFilterLogsInput) StringFromEthereumBlockNumber() (res string) {
	res = fmt.Sprintf("%x", x.FromEthereumBlockNumber)
	return
}

func (x *EthereumFilterLogsInput) StringToEthereumBlockNumber() (res string) {
	res = fmt.Sprintf("%x", x.ToEthereumBlockNumber)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumFilterLogsOutput (non serializable)

type EthereumFilterLogsOutput struct {
	Event []*EthereumEventData
}

func (x *EthereumFilterLogsOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Event:%s,}", x.StringEvent())
}

func (x *EthereumFilterLogsOutput) StringEvent() (res string) {
	res = "["
	for _, v := range x.Event {
		res += v.String() + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumEventData (non serializable)

type EthereumEventData struct {
	EthereumAbiPackedOutputs []byte
	EthereumBlockNumber      uint64
	EthereumTxindex          uint32
	EthereumLogindex         uint32
	EthereumTxhash           uint32
}

func (x *EthereumEventData) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EthereumAbiPackedOutputs:%s,EthereumBlockNumber:%s,EthereumTxindex:%s,EthereumLogindex:%s,EthereumTxhash:%s,}", x.StringEthereumAbiPackedOutputs(), x.StringEthereumBlockNumber(), x.StringEthereumTxindex(), x.StringEthereumLogindex(), x.StringEthereumTxhash())
}

func (x *EthereumEventData) StringEthereumAbiPackedOutputs() (res string) {
	res = fmt.Sprintf("%x", x.EthereumAbiPackedOutputs)
	return
}

func (x *EthereumEventData) StringEthereumBlockNumber() (res string) {
	res = fmt.Sprintf("%x", x.EthereumBlockNumber)
	return
}

func (x *EthereumEventData) StringEthereumTxindex() (res string) {
	res = fmt.Sprintf("%x", x.EthereumTxindex)
	return
}

func (x *EthereumEventData) StringEthereumLogindex() (res string) {
	res = fmt.Sprintf("%x", x.EthereumLogindex)
	return
}

func (x *EthereumEventData) StringEthereumTxhash() (res string) {
	res = fmt.Sprintf("%x", x.EthereumTxhash)
	return
}
