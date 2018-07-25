// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package handlers

import (
	"fmt"
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

func (x *HandleTransactionsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionsBlockHeader:%s,TransactionsBlockMetadata:%s,TransactionsBlockProof:%s,PrevCommittedTransactionsBlockHeader:%s,PrevCommittedTransactionsBlockMetadata:%s,PrevCommittedTransactionsBlockProof:%s,PrevCommittedResultsBlockProof:%s,}", x.StringTransactionsBlockHeader(), x.StringTransactionsBlockMetadata(), x.StringTransactionsBlockProof(), x.StringPrevCommittedTransactionsBlockHeader(), x.StringPrevCommittedTransactionsBlockMetadata(), x.StringPrevCommittedTransactionsBlockProof(), x.StringPrevCommittedResultsBlockProof())
}

func (x *HandleTransactionsBlockInput) StringTransactionsBlockHeader() (res string) {
	res = x.TransactionsBlockHeader.String()
	return
}

func (x *HandleTransactionsBlockInput) StringTransactionsBlockMetadata() (res string) {
	res = x.TransactionsBlockMetadata.String()
	return
}

func (x *HandleTransactionsBlockInput) StringTransactionsBlockProof() (res string) {
	res = x.TransactionsBlockProof.String()
	return
}

func (x *HandleTransactionsBlockInput) StringPrevCommittedTransactionsBlockHeader() (res string) {
	res = x.PrevCommittedTransactionsBlockHeader.String()
	return
}

func (x *HandleTransactionsBlockInput) StringPrevCommittedTransactionsBlockMetadata() (res string) {
	res = x.PrevCommittedTransactionsBlockMetadata.String()
	return
}

func (x *HandleTransactionsBlockInput) StringPrevCommittedTransactionsBlockProof() (res string) {
	res = x.PrevCommittedTransactionsBlockProof.String()
	return
}

func (x *HandleTransactionsBlockInput) StringPrevCommittedResultsBlockProof() (res string) {
	res = x.PrevCommittedResultsBlockProof.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionsBlockOutput (non serializable)

type HandleTransactionsBlockOutput struct {
}

func (x *HandleTransactionsBlockOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
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

func (x *HandleResultsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ResultsBlockHeader:%s,ResultsBlockProof:%s,PrevResultsBlockHeader:%s,PrevCommittedTransactionsBlockMetadata:%s,PrevCommittedTransactionsBlockProof:%s,PrevCommittedResultsBlockProof:%s,}", x.StringResultsBlockHeader(), x.StringResultsBlockProof(), x.StringPrevResultsBlockHeader(), x.StringPrevCommittedTransactionsBlockMetadata(), x.StringPrevCommittedTransactionsBlockProof(), x.StringPrevCommittedResultsBlockProof())
}

func (x *HandleResultsBlockInput) StringResultsBlockHeader() (res string) {
	res = x.ResultsBlockHeader.String()
	return
}

func (x *HandleResultsBlockInput) StringResultsBlockProof() (res string) {
	res = x.ResultsBlockProof.String()
	return
}

func (x *HandleResultsBlockInput) StringPrevResultsBlockHeader() (res string) {
	res = x.PrevResultsBlockHeader.String()
	return
}

func (x *HandleResultsBlockInput) StringPrevCommittedTransactionsBlockMetadata() (res string) {
	res = x.PrevCommittedTransactionsBlockMetadata.String()
	return
}

func (x *HandleResultsBlockInput) StringPrevCommittedTransactionsBlockProof() (res string) {
	res = x.PrevCommittedTransactionsBlockProof.String()
	return
}

func (x *HandleResultsBlockInput) StringPrevCommittedResultsBlockProof() (res string) {
	res = x.PrevCommittedResultsBlockProof.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HandleResultsBlockOutput (non serializable)

type HandleResultsBlockOutput struct {
}

func (x *HandleResultsBlockOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

