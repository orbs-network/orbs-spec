// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package gossiptopics

import (
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionRelay

type TransactionRelay interface {
	BroadcastForwardedTransactions(input *ForwardedTransactionsInput) (*EmptyOutput, error)
	RegisterTransactionRelayHandler(handler TransactionRelayHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service TransactionRelayHandler

type TransactionRelayHandler interface {
	HandleForwardedTransactions(input *ForwardedTransactionsInput) (*EmptyOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message ForwardedTransactionsInput (non serializable)

type ForwardedTransactionsInput struct {
	Message *gossipmessages.ForwardedTransactionsMessage
}

func (x *ForwardedTransactionsInput) String() string {
	return fmt.Sprintf("{Message:%s,}", x.StringMessage())
}

func (x *ForwardedTransactionsInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums

