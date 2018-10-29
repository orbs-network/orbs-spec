// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package handlers

import (
	"fmt"
	"context"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionResultsHandler

type TransactionResultsHandler interface {
	HandleTransactionResults(ctx context.Context, input *HandleTransactionResultsInput) (*HandleTransactionResultsOutput, error)
	HandleTransactionError(ctx context.Context, input *HandleTransactionErrorInput) (*HandleTransactionErrorOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionResultsInput (non serializable)

type HandleTransactionResultsInput struct {
	TransactionReceipts []*protocol.TransactionReceipt
	BlockHeight primitives.BlockHeight
	Timestamp primitives.TimestampNano
}

func (x *HandleTransactionResultsInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionReceipts:%s,BlockHeight:%s,Timestamp:%s,}", x.StringTransactionReceipts(), x.StringBlockHeight(), x.StringTimestamp())
}

func (x *HandleTransactionResultsInput) StringTransactionReceipts() (res string) {
	res = "["
		for _, v := range x.TransactionReceipts {
		res += v.String() + ","
  }
	res += "]"
	return
}

func (x *HandleTransactionResultsInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *HandleTransactionResultsInput) StringTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.Timestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionResultsOutput (non serializable)

type HandleTransactionResultsOutput struct {
}

func (x *HandleTransactionResultsOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionErrorInput (non serializable)

type HandleTransactionErrorInput struct {
	Txhash primitives.Sha256
	TransactionStatus protocol.TransactionStatus
	BlockHeight primitives.BlockHeight
	BlockTimestamp primitives.TimestampNano
}

func (x *HandleTransactionErrorInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Txhash:%s,TransactionStatus:%s,BlockHeight:%s,BlockTimestamp:%s,}", x.StringTxhash(), x.StringTransactionStatus(), x.StringBlockHeight(), x.StringBlockTimestamp())
}

func (x *HandleTransactionErrorInput) StringTxhash() (res string) {
	res = fmt.Sprintf("%s", x.Txhash)
	return
}

func (x *HandleTransactionErrorInput) StringTransactionStatus() (res string) {
	res = fmt.Sprintf("%x", x.TransactionStatus)
	return
}

func (x *HandleTransactionErrorInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *HandleTransactionErrorInput) StringBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.BlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionErrorOutput (non serializable)

type HandleTransactionErrorOutput struct {
}

func (x *HandleTransactionErrorOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// enums

