// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package handlers

import (
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusBlocksHandler

type ConsensusBlocksHandler interface {
	HandleBlockConsensus(input *HandleBlockConsensusInput) (*HandleBlockConsensusOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleBlockConsensusInput (non serializable)

type HandleBlockConsensusInput struct {
	BlockType protocol.BlockType
	BlockPair *protocol.BlockPairContainer
	PrevCommittedBlockPair *protocol.BlockPairContainer
}

func (x *HandleBlockConsensusInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockType:%s,BlockPair:%s,PrevCommittedBlockPair:%s,}", x.StringBlockType(), x.StringBlockPair(), x.StringPrevCommittedBlockPair())
}

func (x *HandleBlockConsensusInput) StringBlockType() (res string) {
	res = fmt.Sprintf("%x", x.BlockType)
	return
}

func (x *HandleBlockConsensusInput) StringBlockPair() (res string) {
	res = x.BlockPair.String()
	return
}

func (x *HandleBlockConsensusInput) StringPrevCommittedBlockPair() (res string) {
	res = x.PrevCommittedBlockPair.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HandleBlockConsensusOutput (non serializable)

type HandleBlockConsensusOutput struct {
}

func (x *HandleBlockConsensusOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// enums

