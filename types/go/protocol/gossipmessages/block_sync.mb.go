// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package gossipmessages

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilityRequestMessage (non serializable)

type BlockAvailabilityRequestMessage struct {
	SignedRange *BlockSyncRange
	Sender *SenderSignature
}

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilityResponseMessage (non serializable)

type BlockAvailabilityResponseMessage struct {
	SignedRange *BlockSyncRange
	Sender *SenderSignature
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestMessage (non serializable)

type BlockSyncRequestMessage struct {
	SignedRange *BlockSyncRange
	Sender *SenderSignature
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseMessage (non serializable)

type BlockSyncResponseMessage struct {
	SignedRange *BlockSyncRange
	Sender *SenderSignature
	BlockPairs []*protocol.BlockPairContainer
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRange

// reader

type BlockSyncRange struct {
	// BlockType BlockType
	// FirstAvailableBlockHeight primitives.BlockHeight
	// LastAvailableBlockHeight primitives.BlockHeight
	// LastCommittedBlockHeight primitives.BlockHeight

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *BlockSyncRange) String() string {
	return fmt.Sprintf("{BlockType:%s,FirstAvailableBlockHeight:%s,LastAvailableBlockHeight:%s,LastCommittedBlockHeight:%s,}", x.StringBlockType(), x.StringFirstAvailableBlockHeight(), x.StringLastAvailableBlockHeight(), x.StringLastCommittedBlockHeight())
}

var _BlockSyncRange_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,}
var _BlockSyncRange_Unions = [][]membuffers.FieldType{}

func BlockSyncRangeReader(buf []byte) *BlockSyncRange {
	x := &BlockSyncRange{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncRange_Scheme, _BlockSyncRange_Unions)
	return x
}

func (x *BlockSyncRange) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncRange) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncRange) BlockType() BlockType {
	return BlockType(x._message.GetUint16(0))
}

func (x *BlockSyncRange) RawBlockType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncRange) MutateBlockType(v BlockType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *BlockSyncRange) StringBlockType() string {
	return x.BlockType().String()
}

func (x *BlockSyncRange) FirstAvailableBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *BlockSyncRange) RawFirstAvailableBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncRange) MutateFirstAvailableBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *BlockSyncRange) StringFirstAvailableBlockHeight() string {
	return fmt.Sprintf("%x", x.FirstAvailableBlockHeight())
}

func (x *BlockSyncRange) LastAvailableBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *BlockSyncRange) RawLastAvailableBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BlockSyncRange) MutateLastAvailableBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *BlockSyncRange) StringLastAvailableBlockHeight() string {
	return fmt.Sprintf("%x", x.LastAvailableBlockHeight())
}

func (x *BlockSyncRange) LastCommittedBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(3))
}

func (x *BlockSyncRange) RawLastCommittedBlockHeight() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *BlockSyncRange) MutateLastCommittedBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *BlockSyncRange) StringLastCommittedBlockHeight() string {
	return fmt.Sprintf("%x", x.LastCommittedBlockHeight())
}

// builder

type BlockSyncRangeBuilder struct {
	BlockType BlockType
	FirstAvailableBlockHeight primitives.BlockHeight
	LastAvailableBlockHeight primitives.BlockHeight
	LastCommittedBlockHeight primitives.BlockHeight

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncRangeBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint16(buf, uint16(w.BlockType))
	w._builder.WriteUint64(buf, uint64(w.FirstAvailableBlockHeight))
	w._builder.WriteUint64(buf, uint64(w.LastAvailableBlockHeight))
	w._builder.WriteUint64(buf, uint64(w.LastCommittedBlockHeight))
	return nil
}

func (w *BlockSyncRangeBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncRangeBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncRangeBuilder) Build() *BlockSyncRange {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncRangeReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type BlockSyncMessageType uint16

const (
	BLOCK_SYNC_RESERVED BlockSyncMessageType = 0
	BLOCK_SYNC_AVAILABILITY_REQUEST BlockSyncMessageType = 1
	BLOCK_SYNC_AVAILABILITY_RESPONSE BlockSyncMessageType = 2
	BLOCK_SYNC_REQUEST BlockSyncMessageType = 3
	BLOCK_SYNC_RESPONSE BlockSyncMessageType = 4
)

func (n BlockSyncMessageType) String() string {
	switch n {
	case BLOCK_SYNC_RESERVED:
		return "BLOCK_SYNC_RESERVED"
	case BLOCK_SYNC_AVAILABILITY_REQUEST:
		return "BLOCK_SYNC_AVAILABILITY_REQUEST"
	case BLOCK_SYNC_AVAILABILITY_RESPONSE:
		return "BLOCK_SYNC_AVAILABILITY_RESPONSE"
	case BLOCK_SYNC_REQUEST:
		return "BLOCK_SYNC_REQUEST"
	case BLOCK_SYNC_RESPONSE:
		return "BLOCK_SYNC_RESPONSE"
	}
	return "UNKNOWN"
}

type BlockType uint16

const (
	BLOCK_TYPE_RESERVED BlockType = 0
	BLOCK_TYPE_BLOCK_PAIR BlockType = 1
	BLOCK_TYPE_TRANSACTIONS_BLOCK BlockType = 2
	BLOCK_TYPE_RESULTS_BLOCK BlockType = 3
)

func (n BlockType) String() string {
	switch n {
	case BLOCK_TYPE_RESERVED:
		return "BLOCK_TYPE_RESERVED"
	case BLOCK_TYPE_BLOCK_PAIR:
		return "BLOCK_TYPE_BLOCK_PAIR"
	case BLOCK_TYPE_TRANSACTIONS_BLOCK:
		return "BLOCK_TYPE_TRANSACTIONS_BLOCK"
	case BLOCK_TYPE_RESULTS_BLOCK:
		return "BLOCK_TYPE_RESULTS_BLOCK"
	}
	return "UNKNOWN"
}

