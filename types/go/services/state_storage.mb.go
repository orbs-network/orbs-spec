// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package services

import (
	"fmt"
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

func (x *ReadKeysInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,ContractName:%s,Keys:%s,}", x.StringBlockHeight(), x.StringContractName(), x.StringKeys())
}

func (x *ReadKeysInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *ReadKeysInput) StringContractName() (res string) {
	res = fmt.Sprintf("%s", x.ContractName)
	return
}

func (x *ReadKeysInput) StringKeys() (res string) {
	res = "["
		for _, v := range x.Keys {
		res += fmt.Sprintf("%s", v) + ","
  }
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ReadKeysOutput (non serializable)

type ReadKeysOutput struct {
	StateRecords []*protocol.StateRecord
}

func (x *ReadKeysOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{StateRecords:%s,}", x.StringStateRecords())
}

func (x *ReadKeysOutput) StringStateRecords() (res string) {
	res = "["
		for _, v := range x.StateRecords {
		res += v.String() + ","
  }
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateStorageBlockHeightInput (non serializable)

type GetStateStorageBlockHeightInput struct {
}

func (x *GetStateStorageBlockHeightInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateStorageBlockHeightOutput (non serializable)

type GetStateStorageBlockHeightOutput struct {
	LastCommittedBlockHeight primitives.BlockHeight
	LastCommittedBlockTimestamp primitives.TimestampNano
}

func (x *GetStateStorageBlockHeightOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{LastCommittedBlockHeight:%s,LastCommittedBlockTimestamp:%s,}", x.StringLastCommittedBlockHeight(), x.StringLastCommittedBlockTimestamp())
}

func (x *GetStateStorageBlockHeightOutput) StringLastCommittedBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.LastCommittedBlockHeight)
	return
}

func (x *GetStateStorageBlockHeightOutput) StringLastCommittedBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.LastCommittedBlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CommitStateDiffInput (non serializable)

type CommitStateDiffInput struct {
	ResultsBlockHeader *protocol.ResultsBlockHeader
	ContractStateDiffs []*protocol.ContractStateDiff
}

func (x *CommitStateDiffInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ResultsBlockHeader:%s,ContractStateDiffs:%s,}", x.StringResultsBlockHeader(), x.StringContractStateDiffs())
}

func (x *CommitStateDiffInput) StringResultsBlockHeader() (res string) {
	res = x.ResultsBlockHeader.String()
	return
}

func (x *CommitStateDiffInput) StringContractStateDiffs() (res string) {
	res = "["
		for _, v := range x.ContractStateDiffs {
		res += v.String() + ","
  }
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CommitStateDiffOutput (non serializable)

type CommitStateDiffOutput struct {
	NextDesiredBlockHeight primitives.BlockHeight
}

func (x *CommitStateDiffOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{NextDesiredBlockHeight:%s,}", x.StringNextDesiredBlockHeight())
}

func (x *CommitStateDiffOutput) StringNextDesiredBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.NextDesiredBlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateHashInput (non serializable)

type GetStateHashInput struct {
	BlockHeight primitives.BlockHeight
}

func (x *GetStateHashInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,}", x.StringBlockHeight())
}

func (x *GetStateHashInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateHashOutput (non serializable)

type GetStateHashOutput struct {
	StateRootHash primitives.MerkleSha256
}

func (x *GetStateHashOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{StateRootHash:%s,}", x.StringStateRootHash())
}

func (x *GetStateHashOutput) StringStateRootHash() (res string) {
	res = fmt.Sprintf("%s", x.StateRootHash)
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums

