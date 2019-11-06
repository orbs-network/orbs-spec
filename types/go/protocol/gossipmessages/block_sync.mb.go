// AUTO GENERATED FILE (by membufc proto compiler v0.3.6)
package gossipmessages

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilityRequestMessage (non serializable)

type BlockAvailabilityRequestMessage struct {
	SignedBatchRange *BlockSyncRange
	Sender           *SenderSignature
}

func (x *BlockAvailabilityRequestMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedBatchRange:%s,Sender:%s,}", x.StringSignedBatchRange(), x.StringSender())
}

func (x *BlockAvailabilityRequestMessage) StringSignedBatchRange() (res string) {
	res = x.SignedBatchRange.String()
	return
}

func (x *BlockAvailabilityRequestMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilityResponseMessage (non serializable)

type BlockAvailabilityResponseMessage struct {
	SignedBatchRange *BlockSyncRange
	Sender           *SenderSignature
}

func (x *BlockAvailabilityResponseMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedBatchRange:%s,Sender:%s,}", x.StringSignedBatchRange(), x.StringSender())
}

func (x *BlockAvailabilityResponseMessage) StringSignedBatchRange() (res string) {
	res = x.SignedBatchRange.String()
	return
}

func (x *BlockAvailabilityResponseMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestMessage (non serializable)

type BlockSyncRequestMessage struct {
	SignedChunkRange *BlockSyncRange
	Sender           *SenderSignature
}

func (x *BlockSyncRequestMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedChunkRange:%s,Sender:%s,}", x.StringSignedChunkRange(), x.StringSender())
}

func (x *BlockSyncRequestMessage) StringSignedChunkRange() (res string) {
	res = x.SignedChunkRange.String()
	return
}

func (x *BlockSyncRequestMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseMessage (non serializable)

type BlockSyncResponseMessage struct {
	SignedChunkRange *BlockSyncRange
	Sender           *SenderSignature
	BlockPairs       []*protocol.BlockPairContainer
}

func (x *BlockSyncResponseMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedChunkRange:%s,Sender:%s,BlockPairs:%s,}", x.StringSignedChunkRange(), x.StringSender(), x.StringBlockPairs())
}

func (x *BlockSyncResponseMessage) StringSignedChunkRange() (res string) {
	res = x.SignedChunkRange.String()
	return
}

func (x *BlockSyncResponseMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

func (x *BlockSyncResponseMessage) StringBlockPairs() (res string) {
	res = "["
	for _, v := range x.BlockPairs {
		res += v.String() + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRange

// reader

type BlockSyncRange struct {
	// BlockType BlockType
	// FirstBlockHeight primitives.BlockHeight
	// LastBlockHeight primitives.BlockHeight
	// LastCommittedBlockHeight primitives.BlockHeight

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *BlockSyncRange) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockType:%s,FirstBlockHeight:%s,LastBlockHeight:%s,LastCommittedBlockHeight:%s,}", x.StringBlockType(), x.StringFirstBlockHeight(), x.StringLastBlockHeight(), x.StringLastCommittedBlockHeight())
}

var _BlockSyncRange_Scheme = []membuffers.FieldType{membuffers.TypeUint16, membuffers.TypeUint64, membuffers.TypeUint64, membuffers.TypeUint64}
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

func (x *BlockSyncRange) Equal(y *BlockSyncRange) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
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

func (x *BlockSyncRange) FirstBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *BlockSyncRange) RawFirstBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncRange) MutateFirstBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *BlockSyncRange) StringFirstBlockHeight() string {
	return fmt.Sprintf("%s", x.FirstBlockHeight())
}

func (x *BlockSyncRange) LastBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *BlockSyncRange) RawLastBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BlockSyncRange) MutateLastBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *BlockSyncRange) StringLastBlockHeight() string {
	return fmt.Sprintf("%s", x.LastBlockHeight())
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
	return fmt.Sprintf("%s", x.LastCommittedBlockHeight())
}

// builder

type BlockSyncRangeBuilder struct {
	BlockType                BlockType
	FirstBlockHeight         primitives.BlockHeight
	LastBlockHeight          primitives.BlockHeight
	LastCommittedBlockHeight primitives.BlockHeight

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *BlockSyncRangeBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteUint16(buf, uint16(w.BlockType))
	w._builder.WriteUint64(buf, uint64(w.FirstBlockHeight))
	w._builder.WriteUint64(buf, uint64(w.LastBlockHeight))
	w._builder.WriteUint64(buf, uint64(w.LastCommittedBlockHeight))
	return nil
}

func (w *BlockSyncRangeBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint16(prefix, offsetFromStart, "BlockSyncRange.BlockType", uint16(w.BlockType))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "BlockSyncRange.FirstBlockHeight", uint64(w.FirstBlockHeight))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "BlockSyncRange.LastBlockHeight", uint64(w.LastBlockHeight))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "BlockSyncRange.LastCommittedBlockHeight", uint64(w.LastCommittedBlockHeight))
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

func BlockSyncRangeBuilderFromRaw(raw []byte) *BlockSyncRangeBuilder {
	return &BlockSyncRangeBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// enums

type BlockSyncMessageType uint16

const (
	BLOCK_SYNC_RESERVED              BlockSyncMessageType = 0
	BLOCK_SYNC_AVAILABILITY_REQUEST  BlockSyncMessageType = 1
	BLOCK_SYNC_AVAILABILITY_RESPONSE BlockSyncMessageType = 2
	BLOCK_SYNC_REQUEST               BlockSyncMessageType = 3
	BLOCK_SYNC_RESPONSE              BlockSyncMessageType = 4
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
	BLOCK_TYPE_RESERVED           BlockType = 0
	BLOCK_TYPE_BLOCK_PAIR         BlockType = 1
	BLOCK_TYPE_TRANSACTIONS_BLOCK BlockType = 2
	BLOCK_TYPE_RESULTS_BLOCK      BlockType = 3
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
