// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusContext

type ConsensusContext interface {
	RequestNewTransactionsBlock(input *RequestNewTransactionsBlockInput) (*RequestNewTransactionsBlockOutput, error)
	RequestNewResultsBlock(input *RequestNewResultsBlockInput) (*RequestNewResultsBlockOutput, error)
	ValidateTransactionsBlock(input *ValidateTransactionsBlockInput) (*ValidateTransactionsBlockOutput, error)
	ValidateResultsBlock(input *ValidateResultsBlockInput) (*ValidateResultsBlockOutput, error)
	RequestOrderingCommittee(input *RequestCommitteeInput) (*RequestCommitteeOutput, error)
	RequestValidationCommittee(input *RequestCommitteeInput) (*RequestCommitteeOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewTransactionsBlockInput (non serializable)

type RequestNewTransactionsBlockInput struct {
	BlockHeight uint64
	MaxBlockSizeKb uint32
	MaxNumberOfTransactions uint32
	PrevBlockHash []byte
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewTransactionsBlockOutput (non serializable)

type RequestNewTransactionsBlockOutput struct {
	Status protocol.RequestStatus
	TransactionsBlock *protocol.TransactionsBlock
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewResultsBlockInput (non serializable)

type RequestNewResultsBlockInput struct {
	BlockHeight uint64
	PrevBlockHash []byte
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewResultsBlockOutput (non serializable)

type RequestNewResultsBlockOutput struct {
	Status protocol.RequestStatus
	ResultsBlock *protocol.ResultsBlock
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsBlockInput (non serializable)

type ValidateTransactionsBlockInput struct {
	TransactionsBlock *protocol.TransactionsBlock
	PrevBlockHash []byte
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsBlockOutput (non serializable)

type ValidateTransactionsBlockOutput struct {
	Status protocol.RequestStatus
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateResultsBlockInput (non serializable)

type ValidateResultsBlockInput struct {
	ResultsBlock *protocol.ResultsBlock
	PrevBlockHash []byte
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateResultsBlockOutput (non serializable)

type ValidateResultsBlockOutput struct {
	Status protocol.RequestStatus
}

/////////////////////////////////////////////////////////////////////////////
// message RequestCommitteeInput (non serializable)

type RequestCommitteeInput struct {
	BlockHeight uint64
	RandomSeed uint64
	MaxCommitteeSize uint32
}

/////////////////////////////////////////////////////////////////////////////
// message RequestCommitteeOutput (non serializable)

type RequestCommitteeOutput struct {
	Status protocol.RequestStatus
	NodeData []*protocol.NodeData
}

/////////////////////////////////////////////////////////////////////////////
// enums

