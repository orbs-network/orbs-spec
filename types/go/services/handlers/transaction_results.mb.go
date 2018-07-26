// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package handlers

import (
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionResultsHandler

type TransactionResultsHandler interface {
	HandleTransactionResults(input *HandleTransactionResultsInput) (*HandleTransactionResultsOutput, error)
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

