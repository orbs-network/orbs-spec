// AUTO GENERATED FILE (by membufc proto compiler v0.0.11)
package services

import (
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
	BlockHeight uint64
	Contract *protocol.ContractAddress
	Key [][]byte
}

/////////////////////////////////////////////////////////////////////////////
// message ReadKeysOutput (non serializable)

type ReadKeysOutput struct {
	Status protocol.RequestStatus
	Blob [][]byte
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateStorageBlockHeightInput (non serializable)

type GetStateStorageBlockHeightInput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateStorageBlockHeightOutput (non serializable)

type GetStateStorageBlockHeightOutput struct {
	LastCommittedBlockHeight uint64
	LastCommittedBlockTimestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// message CommitStateDiffInput (non serializable)

type CommitStateDiffInput struct {
	ResultsBlockHeader *protocol.ResultsBlockHeader
	ContractStateDiff []*protocol.ContractStateDiff
	LastCommittedBlockHeight uint64
}

/////////////////////////////////////////////////////////////////////////////
// message CommitStateDiffOutput (non serializable)

type CommitStateDiffOutput struct {
	NextDesiredBlockHeight uint64
	LastCommittedBlockHeight uint64
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateHashInput (non serializable)

type GetStateHashInput struct {
	BlockHeight uint64
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateHashOutput (non serializable)

type GetStateHashOutput struct {
	Status protocol.RequestStatus
	StateRootHash []byte
}

/////////////////////////////////////////////////////////////////////////////
// enums

