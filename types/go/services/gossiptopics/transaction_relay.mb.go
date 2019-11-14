// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package gossiptopics

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionRelay

type TransactionRelay interface {
	BroadcastForwardedTransactions(ctx context.Context, input *ForwardedTransactionsInput) (*EmptyOutput, error)
	RegisterTransactionRelayHandler(handler TransactionRelayHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service TransactionRelayHandler

type TransactionRelayHandler interface {
	HandleForwardedTransactions(ctx context.Context, input *ForwardedTransactionsInput) (*EmptyOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message ForwardedTransactionsInput (non serializable)

type ForwardedTransactionsInput struct {
	Message *gossipmessages.ForwardedTransactionsMessage
}

func (x *ForwardedTransactionsInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Message:%s,}", x.StringMessage())
}

func (x *ForwardedTransactionsInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
