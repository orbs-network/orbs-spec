// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package handlers

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusBlocksHandler

type ConsensusBlocksHandler interface {
	HandleBlockConsensus(ctx context.Context, input *HandleBlockConsensusInput) (*HandleBlockConsensusOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleBlockConsensusInput (non serializable)

type HandleBlockConsensusInput struct {
	Mode          HandleBlockConsensusMode
	BlockType     protocol.BlockType
	BlockPair     *protocol.BlockPairContainer
	PrevBlockPair *protocol.BlockPairContainer
}

func (x *HandleBlockConsensusInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Mode:%s,BlockType:%s,BlockPair:%s,PrevBlockPair:%s,}", x.StringMode(), x.StringBlockType(), x.StringBlockPair(), x.StringPrevBlockPair())
}

func (x *HandleBlockConsensusInput) StringMode() (res string) {
	res = fmt.Sprintf("%x", x.Mode)
	return
}

func (x *HandleBlockConsensusInput) StringBlockType() (res string) {
	res = fmt.Sprintf("%x", x.BlockType)
	return
}

func (x *HandleBlockConsensusInput) StringBlockPair() (res string) {
	res = x.BlockPair.String()
	return
}

func (x *HandleBlockConsensusInput) StringPrevBlockPair() (res string) {
	res = x.PrevBlockPair.String()
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

type HandleBlockConsensusMode uint16

const (
	HANDLE_BLOCK_CONSENSUS_MODE_RESERVED          HandleBlockConsensusMode = 0
	HANDLE_BLOCK_CONSENSUS_MODE_VERIFY_AND_UPDATE HandleBlockConsensusMode = 1
	HANDLE_BLOCK_CONSENSUS_MODE_UPDATE_ONLY       HandleBlockConsensusMode = 2
	HANDLE_BLOCK_CONSENSUS_MODE_VERIFY_ONLY       HandleBlockConsensusMode = 3
	HANDLE_BLOCK_CONSENSUS_MODE_VERIFY_CHAIN_TIP  HandleBlockConsensusMode = 4
)

func (n HandleBlockConsensusMode) String() string {
	switch n {
	case HANDLE_BLOCK_CONSENSUS_MODE_RESERVED:
		return "HANDLE_BLOCK_CONSENSUS_MODE_RESERVED"
	case HANDLE_BLOCK_CONSENSUS_MODE_VERIFY_AND_UPDATE:
		return "HANDLE_BLOCK_CONSENSUS_MODE_VERIFY_AND_UPDATE"
	case HANDLE_BLOCK_CONSENSUS_MODE_UPDATE_ONLY:
		return "HANDLE_BLOCK_CONSENSUS_MODE_UPDATE_ONLY"
	case HANDLE_BLOCK_CONSENSUS_MODE_VERIFY_ONLY:
		return "HANDLE_BLOCK_CONSENSUS_MODE_VERIFY_ONLY"
	case HANDLE_BLOCK_CONSENSUS_MODE_VERIFY_CHAIN_TIP:
		return "HANDLE_BLOCK_CONSENSUS_MODE_VERIFY_CHAIN_TIP"
	}
	return "UNKNOWN"
}
