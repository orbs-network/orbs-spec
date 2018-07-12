// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package gossiptopics

import (
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

/////////////////////////////////////////////////////////////////////////////
// enums

