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
	BlockAvailabilityRequestSignedHeader *BlockAvailabilitySignedHeader
	Signer *MessageSigner
}

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilityResponseMessage (non serializable)

type BlockAvailabilityResponseMessage struct {
	BlockAvailabilityResponseSignedHeader *BlockAvailabilitySignedHeader
	Signer *MessageSigner
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestMessage (non serializable)

type BlockSyncRequestMessage struct {
	BlockSyncRequestSignedHeader *BlockSyncRequestSignedHeader
	Signer *MessageSigner
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseMessage (non serializable)

type BlockSyncResponseMessage struct {
	BlockSyncResponseSignedHeader *BlockSyncResponseSignedHeader
	Signer *MessageSigner
	BlockPairs []*protocol.BlockPair
}

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilitySignedHeader

// reader

type BlockAvailabilitySignedHeader struct {
	// BlockType BlockType
	// FirstAvailableBlockHeight primitives.BlockHeight
	// LastAvailableBlockHeight primitives.BlockHeight
	// LastCommittedBlockHeight primitives.BlockHeight

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *BlockAvailabilitySignedHeader) String() string {
	return fmt.Sprintf("{BlockType:%s,FirstAvailableBlockHeight:%s,LastAvailableBlockHeight:%s,LastCommittedBlockHeight:%s,}", x.StringBlockType(), x.StringFirstAvailableBlockHeight(), x.StringLastAvailableBlockHeight(), x.StringLastCommittedBlockHeight())
}

var _BlockAvailabilitySignedHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,}
var _BlockAvailabilitySignedHeader_Unions = [][]membuffers.FieldType{}

func BlockAvailabilitySignedHeaderReader(buf []byte) *BlockAvailabilitySignedHeader {
	x := &BlockAvailabilitySignedHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockAvailabilitySignedHeader_Scheme, _BlockAvailabilitySignedHeader_Unions)
	return x
}

func (x *BlockAvailabilitySignedHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockAvailabilitySignedHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockAvailabilitySignedHeader) BlockType() BlockType {
	return BlockType(x._message.GetUint16(0))
}

func (x *BlockAvailabilitySignedHeader) RawBlockType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockAvailabilitySignedHeader) MutateBlockType(v BlockType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *BlockAvailabilitySignedHeader) StringBlockType() string {
	return x.BlockType().String()
}

func (x *BlockAvailabilitySignedHeader) FirstAvailableBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *BlockAvailabilitySignedHeader) RawFirstAvailableBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockAvailabilitySignedHeader) MutateFirstAvailableBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *BlockAvailabilitySignedHeader) StringFirstAvailableBlockHeight() string {
	return fmt.Sprintf("%x", x.FirstAvailableBlockHeight())
}

func (x *BlockAvailabilitySignedHeader) LastAvailableBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *BlockAvailabilitySignedHeader) RawLastAvailableBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BlockAvailabilitySignedHeader) MutateLastAvailableBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *BlockAvailabilitySignedHeader) StringLastAvailableBlockHeight() string {
	return fmt.Sprintf("%x", x.LastAvailableBlockHeight())
}

func (x *BlockAvailabilitySignedHeader) LastCommittedBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(3))
}

func (x *BlockAvailabilitySignedHeader) RawLastCommittedBlockHeight() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *BlockAvailabilitySignedHeader) MutateLastCommittedBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *BlockAvailabilitySignedHeader) StringLastCommittedBlockHeight() string {
	return fmt.Sprintf("%x", x.LastCommittedBlockHeight())
}

// builder

type BlockAvailabilitySignedHeaderBuilder struct {
	BlockType BlockType
	FirstAvailableBlockHeight primitives.BlockHeight
	LastAvailableBlockHeight primitives.BlockHeight
	LastCommittedBlockHeight primitives.BlockHeight

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockAvailabilitySignedHeaderBuilder) Write(buf []byte) (err error) {
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

func (w *BlockAvailabilitySignedHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockAvailabilitySignedHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockAvailabilitySignedHeaderBuilder) Build() *BlockAvailabilitySignedHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockAvailabilitySignedHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestSignedHeader

// reader

type BlockSyncRequestSignedHeader struct {
	// BlockType BlockType
	// FirstAvailableBlockHeight primitives.BlockHeight
	// LastAvailableBlockHeight primitives.BlockHeight
	// LastCommittedBlockHeight primitives.BlockHeight

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *BlockSyncRequestSignedHeader) String() string {
	return fmt.Sprintf("{BlockType:%s,FirstAvailableBlockHeight:%s,LastAvailableBlockHeight:%s,LastCommittedBlockHeight:%s,}", x.StringBlockType(), x.StringFirstAvailableBlockHeight(), x.StringLastAvailableBlockHeight(), x.StringLastCommittedBlockHeight())
}

var _BlockSyncRequestSignedHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,}
var _BlockSyncRequestSignedHeader_Unions = [][]membuffers.FieldType{}

func BlockSyncRequestSignedHeaderReader(buf []byte) *BlockSyncRequestSignedHeader {
	x := &BlockSyncRequestSignedHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncRequestSignedHeader_Scheme, _BlockSyncRequestSignedHeader_Unions)
	return x
}

func (x *BlockSyncRequestSignedHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncRequestSignedHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncRequestSignedHeader) BlockType() BlockType {
	return BlockType(x._message.GetUint16(0))
}

func (x *BlockSyncRequestSignedHeader) RawBlockType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncRequestSignedHeader) MutateBlockType(v BlockType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *BlockSyncRequestSignedHeader) StringBlockType() string {
	return x.BlockType().String()
}

func (x *BlockSyncRequestSignedHeader) FirstAvailableBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *BlockSyncRequestSignedHeader) RawFirstAvailableBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncRequestSignedHeader) MutateFirstAvailableBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *BlockSyncRequestSignedHeader) StringFirstAvailableBlockHeight() string {
	return fmt.Sprintf("%x", x.FirstAvailableBlockHeight())
}

func (x *BlockSyncRequestSignedHeader) LastAvailableBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *BlockSyncRequestSignedHeader) RawLastAvailableBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BlockSyncRequestSignedHeader) MutateLastAvailableBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *BlockSyncRequestSignedHeader) StringLastAvailableBlockHeight() string {
	return fmt.Sprintf("%x", x.LastAvailableBlockHeight())
}

func (x *BlockSyncRequestSignedHeader) LastCommittedBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(3))
}

func (x *BlockSyncRequestSignedHeader) RawLastCommittedBlockHeight() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *BlockSyncRequestSignedHeader) MutateLastCommittedBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *BlockSyncRequestSignedHeader) StringLastCommittedBlockHeight() string {
	return fmt.Sprintf("%x", x.LastCommittedBlockHeight())
}

// builder

type BlockSyncRequestSignedHeaderBuilder struct {
	BlockType BlockType
	FirstAvailableBlockHeight primitives.BlockHeight
	LastAvailableBlockHeight primitives.BlockHeight
	LastCommittedBlockHeight primitives.BlockHeight

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncRequestSignedHeaderBuilder) Write(buf []byte) (err error) {
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

func (w *BlockSyncRequestSignedHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncRequestSignedHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncRequestSignedHeaderBuilder) Build() *BlockSyncRequestSignedHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncRequestSignedHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseSignedHeader

// reader

type BlockSyncResponseSignedHeader struct {
	// BlockType BlockType
	// FirstBlockHeight primitives.BlockHeight
	// LastBlockHeight primitives.BlockHeight
	// LastCommittedBlockHeight primitives.BlockHeight
	// BlockHashes []primitives.Sha256

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *BlockSyncResponseSignedHeader) String() string {
	return fmt.Sprintf("{BlockType:%s,FirstBlockHeight:%s,LastBlockHeight:%s,LastCommittedBlockHeight:%s,BlockHashes:%s,}", x.StringBlockType(), x.StringFirstBlockHeight(), x.StringLastBlockHeight(), x.StringLastCommittedBlockHeight(), x.StringBlockHashes())
}

var _BlockSyncResponseSignedHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeBytesArray,}
var _BlockSyncResponseSignedHeader_Unions = [][]membuffers.FieldType{}

func BlockSyncResponseSignedHeaderReader(buf []byte) *BlockSyncResponseSignedHeader {
	x := &BlockSyncResponseSignedHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncResponseSignedHeader_Scheme, _BlockSyncResponseSignedHeader_Unions)
	return x
}

func (x *BlockSyncResponseSignedHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncResponseSignedHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncResponseSignedHeader) BlockType() BlockType {
	return BlockType(x._message.GetUint16(0))
}

func (x *BlockSyncResponseSignedHeader) RawBlockType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncResponseSignedHeader) MutateBlockType(v BlockType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *BlockSyncResponseSignedHeader) StringBlockType() string {
	return x.BlockType().String()
}

func (x *BlockSyncResponseSignedHeader) FirstBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *BlockSyncResponseSignedHeader) RawFirstBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncResponseSignedHeader) MutateFirstBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *BlockSyncResponseSignedHeader) StringFirstBlockHeight() string {
	return fmt.Sprintf("%x", x.FirstBlockHeight())
}

func (x *BlockSyncResponseSignedHeader) LastBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *BlockSyncResponseSignedHeader) RawLastBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BlockSyncResponseSignedHeader) MutateLastBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *BlockSyncResponseSignedHeader) StringLastBlockHeight() string {
	return fmt.Sprintf("%x", x.LastBlockHeight())
}

func (x *BlockSyncResponseSignedHeader) LastCommittedBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(3))
}

func (x *BlockSyncResponseSignedHeader) RawLastCommittedBlockHeight() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *BlockSyncResponseSignedHeader) MutateLastCommittedBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *BlockSyncResponseSignedHeader) StringLastCommittedBlockHeight() string {
	return fmt.Sprintf("%x", x.LastCommittedBlockHeight())
}

func (x *BlockSyncResponseSignedHeader) BlockHashesIterator() *BlockSyncResponseSignedHeaderBlockHashesIterator {
	return &BlockSyncResponseSignedHeaderBlockHashesIterator{iterator: x._message.GetBytesArrayIterator(4)}
}

type BlockSyncResponseSignedHeaderBlockHashesIterator struct {
	iterator *membuffers.Iterator
}

func (i *BlockSyncResponseSignedHeaderBlockHashesIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *BlockSyncResponseSignedHeaderBlockHashesIterator) NextBlockHashes() primitives.Sha256 {
	return primitives.Sha256(i.iterator.NextBytes())
}

func (x *BlockSyncResponseSignedHeader) RawBlockHashesArray() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *BlockSyncResponseSignedHeader) StringBlockHashes() (res string) {
	res = "["
	for i := x.BlockHashesIterator(); i.HasNext(); {
		res += fmt.Sprintf("%x", i.NextBlockHashes()) + ","
	}
	res += "]"
	return
}

// builder

type BlockSyncResponseSignedHeaderBuilder struct {
	BlockType BlockType
	FirstBlockHeight primitives.BlockHeight
	LastBlockHeight primitives.BlockHeight
	LastCommittedBlockHeight primitives.BlockHeight
	BlockHashes []primitives.Sha256

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncResponseSignedHeaderBuilder) arrayOfBlockHashes() [][]byte {
	res := make([][]byte, len(w.BlockHashes))
	for i, v := range w.BlockHashes {
		res[i] = v
	}
	return res
}

func (w *BlockSyncResponseSignedHeaderBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUint64(buf, uint64(w.FirstBlockHeight))
	w._builder.WriteUint64(buf, uint64(w.LastBlockHeight))
	w._builder.WriteUint64(buf, uint64(w.LastCommittedBlockHeight))
	w._builder.WriteBytesArray(buf, w.arrayOfBlockHashes())
	return nil
}

func (w *BlockSyncResponseSignedHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncResponseSignedHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncResponseSignedHeaderBuilder) Build() *BlockSyncResponseSignedHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncResponseSignedHeaderReader(buf)
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

