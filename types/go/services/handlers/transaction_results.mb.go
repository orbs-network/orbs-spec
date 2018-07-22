// AUTO GENERATED FILE (by membufc proto compiler v0.0.17)
package handlers

import (
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

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionResultsOutput (non serializable)

type HandleTransactionResultsOutput struct {
}

