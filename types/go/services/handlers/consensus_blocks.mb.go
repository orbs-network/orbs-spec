// AUTO GENERATED FILE (by membufc proto compiler v0.0.14)
package handlers

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusBlocksHandler

type ConsensusBlocksHandler interface {
	HandleTransactionsBlock(input *HandleTransactionsBlockInput) (*HandleTransactionsBlockOutput, error)
	HandleResultsBlock(input *HandleResultsBlockInput) (*HandleResultsBlockOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionsBlockInput (non serializable)

type HandleTransactionsBlockInput struct {
	TransactionsBlockHeader *protocol.TransactionsBlockHeader
	TransactionsBlockMetadata *protocol.TransactionsBlockMetadata
	TransactionsBlockProof *protocol.TransactionsBlockProof
	PrevCommittedTransactionsBlockHeader *protocol.TransactionsBlockHeader
	PrevCommittedTransactionsBlockMetadata *protocol.TransactionsBlockMetadata
	PrevCommittedTransactionsBlockProof *protocol.TransactionsBlockProof
	PrevCommittedResultsBlockProof *protocol.ResultsBlockProof
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionsBlockOutput (non serializable)

type HandleTransactionsBlockOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message HandleResultsBlockInput (non serializable)

type HandleResultsBlockInput struct {
	ResultsBlockHeader *protocol.ResultsBlockHeader
	ResultsBlockProof *protocol.ResultsBlockProof
	PrevResultsBlockHeader *protocol.ResultsBlockHeader
	PrevCommittedTransactionsBlockMetadata *protocol.TransactionsBlockMetadata
	PrevCommittedTransactionsBlockProof *protocol.TransactionsBlockProof
	PrevCommittedResultsBlockProof *protocol.ResultsBlockProof
}

/////////////////////////////////////////////////////////////////////////////
// message HandleResultsBlockOutput (non serializable)

type HandleResultsBlockOutput struct {
}

