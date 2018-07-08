// AUTO GENERATED FILE (by membufc proto compiler)
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
	TransactionsBlockMetaData *protocol.TransactionsBlockMetaData
	TransactionsBlockProof *protocol.TransactionsBlockProof
	PreviousCommittedTransactionsBlockHeader *protocol.TransactionsBlockHeader
	PreviousCommittedTransactionsBlockMetaData *protocol.TransactionsBlockMetaData
	PreviousCommittedTransactionsBlockProof *protocol.TransactionsBlockProof
	PreviousCommittedResultsBlockProof *protocol.ResultsBlockProof
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionsBlockOutput (non serializable)

type HandleTransactionsBlockOutput struct {
	Status protocol.ValidateBlockConsensus
}

/////////////////////////////////////////////////////////////////////////////
// message HandleResultsBlockInput (non serializable)

type HandleResultsBlockInput struct {
	ResultsBlockHeader *protocol.ResultsBlockHeader
	ResultsBlockProof *protocol.ResultsBlockProof
	PreviousResultsBlockHeader *protocol.ResultsBlockHeader
	PreviousCommittedTransactionsBlockMetaData *protocol.TransactionsBlockMetaData
	PreviousCommittedTransactionsBlockProof *protocol.TransactionsBlockProof
	PreviousCommittedResultsBlockProof *protocol.ResultsBlockProof
}

/////////////////////////////////////////////////////////////////////////////
// message HandleResultsBlockOutput (non serializable)

type HandleResultsBlockOutput struct {
	Status protocol.ValidateBlockConsensus
}

/////////////////////////////////////////////////////////////////////////////
// enums

