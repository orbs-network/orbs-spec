// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
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
	BlockHeight primitives.BlockHeight
	MaxBlockSizeKb uint32
	MaxNumberOfTransactions uint32
	PrevBlockHash primitives.Sha256
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewTransactionsBlockOutput (non serializable)

type RequestNewTransactionsBlockOutput struct {
	TransactionsBlock *protocol.TransactionsBlockContainer
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewResultsBlockInput (non serializable)

type RequestNewResultsBlockInput struct {
	BlockHeight primitives.BlockHeight
	PrevBlockHash primitives.Sha256
	TransactionsBlock *protocol.TransactionsBlockContainer
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewResultsBlockOutput (non serializable)

type RequestNewResultsBlockOutput struct {
	ResultsBlock *protocol.ResultsBlockContainer
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsBlockInput (non serializable)

type ValidateTransactionsBlockInput struct {
	TransactionsBlock *protocol.TransactionsBlockContainer
	PrevBlockHash primitives.Sha256
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsBlockOutput (non serializable)

type ValidateTransactionsBlockOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateResultsBlockInput (non serializable)

type ValidateResultsBlockInput struct {
	ResultsBlock *protocol.ResultsBlockContainer
	PrevBlockHash primitives.Sha256
	TransactionsBlock *protocol.TransactionsBlockContainer
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateResultsBlockOutput (non serializable)

type ValidateResultsBlockOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message RequestCommitteeInput (non serializable)

type RequestCommitteeInput struct {
	BlockHeight primitives.BlockHeight
	RandomSeed uint64
	MaxCommitteeSize uint32
}

/////////////////////////////////////////////////////////////////////////////
// message RequestCommitteeOutput (non serializable)

type RequestCommitteeOutput struct {
	NodePublicKeys []primitives.Ed25519Pkey
	NodeRandomSeedPublicKeys []primitives.Bls1Pkey
}

/////////////////////////////////////////////////////////////////////////////
// enums

