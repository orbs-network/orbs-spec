// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package gossiptopics

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionRelay

type TransactionRelay interface {
	BroadcastForwardedTransactions(input *ForwardedTransactionsInput) (*TransactionRelayOutput, error)
	RegisterTransactionRelayHandler(handler TransactionRelayHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service TransactionRelayHandler

type TransactionRelayHandler interface {
	HandleForwardedTransactions(input *ForwardedTransactionsInput) (*TransactionRelayOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message ForwardedTransactionsInput (non serializable)

type ForwardedTransactionsInput struct {
	Header *gossipmessages.ForwardedTransactionsHeader
	Transactions []*protocol.SignedTransaction
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionRelayOutput (non serializable)

type TransactionRelayOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// enums

