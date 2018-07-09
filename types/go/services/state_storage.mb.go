// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service StateStorage

type StateStorage interface {
	CommitStateDiff(input *CommitStateDiffInput) (*CommitStateDiffOutput, error)
	ReadKeys(input *ReadKeysInput) (*ReadKeysOutput, error)
	GetStateStorageBlockHeight(input *GetStateStorageBlockHeightInput) (*GetStateStorageBlockHeightOutput, error)
	GetStateHash(input *GetStateHashInput) (*GetStateHashOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message ReadKeysInput (non serializable)

type ReadKeysInput struct {
	BlockHeight primitives.BlockHeight
	ContractName primitives.ContractName
	Keys []primitives.Ripmd160Sha256
}

/////////////////////////////////////////////////////////////////////////////
// message ReadKeysOutput (non serializable)

type ReadKeysOutput struct {
	StateDiffs []*protocol.StateDiff
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateStorageBlockHeightInput (non serializable)

type GetStateStorageBlockHeightInput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateStorageBlockHeightOutput (non serializable)

type GetStateStorageBlockHeightOutput struct {
	LastCommittedBlockHeight primitives.BlockHeight
	LastCommittedBlockTimestamp primitives.Timestamp
}

/////////////////////////////////////////////////////////////////////////////
// message CommitStateDiffInput (non serializable)

type CommitStateDiffInput struct {
	ResultsBlockHeader *protocol.ResultsBlockHeader
	ContractStateDiffs []*protocol.ContractStateDiff
	LastCommittedBlockHeight primitives.BlockHeight
}

/////////////////////////////////////////////////////////////////////////////
// message CommitStateDiffOutput (non serializable)

type CommitStateDiffOutput struct {
	NextDesiredBlockHeight primitives.BlockHeight
	LastCommittedBlockHeight primitives.BlockHeight
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateHashInput (non serializable)

type GetStateHashInput struct {
	BlockHeight primitives.BlockHeight
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateHashOutput (non serializable)

type GetStateHashOutput struct {
	StateRootHash primitives.MerkleSha256
}

/////////////////////////////////////////////////////////////////////////////
// enums

