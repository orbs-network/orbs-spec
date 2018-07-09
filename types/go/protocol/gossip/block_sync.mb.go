// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package gossip

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncAvailabilityRequestMessage

// reader

type BlockSyncAvailabilityRequestMessage struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _BlockSyncAvailabilityRequestMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var _BlockSyncAvailabilityRequestMessage_Unions = [][]membuffers.FieldType{}

func BlockSyncAvailabilityRequestMessageReader(buf []byte) *BlockSyncAvailabilityRequestMessage {
	x := &BlockSyncAvailabilityRequestMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncAvailabilityRequestMessage_Scheme, _BlockSyncAvailabilityRequestMessage_Unions)
	return x
}

func (x *BlockSyncAvailabilityRequestMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncAvailabilityRequestMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncAvailabilityRequestMessage) Body() *BlockSyncAvailabilitySignedFields {
	b, s := x._message.GetMessage(0)
	return BlockSyncAvailabilitySignedFieldsReader(b[:s])
}

func (x *BlockSyncAvailabilityRequestMessage) RawBody() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncAvailabilityRequestMessage) Signature() primitives.Ed25519Sig {
	return x._message.GetBytes(1)
}

func (x *BlockSyncAvailabilityRequestMessage) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncAvailabilityRequestMessage) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, v)
}

// builder

type BlockSyncAvailabilityRequestMessageBuilder struct {
	Body *BlockSyncAvailabilitySignedFieldsBuilder
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncAvailabilityRequestMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.Body)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *BlockSyncAvailabilityRequestMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncAvailabilityRequestMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncAvailabilityRequestMessageBuilder) Build() *BlockSyncAvailabilityRequestMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncAvailabilityRequestMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncAvailabilityResponseMessage

// reader

type BlockSyncAvailabilityResponseMessage struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _BlockSyncAvailabilityResponseMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var _BlockSyncAvailabilityResponseMessage_Unions = [][]membuffers.FieldType{}

func BlockSyncAvailabilityResponseMessageReader(buf []byte) *BlockSyncAvailabilityResponseMessage {
	x := &BlockSyncAvailabilityResponseMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncAvailabilityResponseMessage_Scheme, _BlockSyncAvailabilityResponseMessage_Unions)
	return x
}

func (x *BlockSyncAvailabilityResponseMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncAvailabilityResponseMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncAvailabilityResponseMessage) Body() *BlockSyncAvailabilitySignedFields {
	b, s := x._message.GetMessage(0)
	return BlockSyncAvailabilitySignedFieldsReader(b[:s])
}

func (x *BlockSyncAvailabilityResponseMessage) RawBody() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncAvailabilityResponseMessage) Signature() primitives.Ed25519Sig {
	return x._message.GetBytes(1)
}

func (x *BlockSyncAvailabilityResponseMessage) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncAvailabilityResponseMessage) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, v)
}

// builder

type BlockSyncAvailabilityResponseMessageBuilder struct {
	Body *BlockSyncAvailabilitySignedFieldsBuilder
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncAvailabilityResponseMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.Body)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *BlockSyncAvailabilityResponseMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncAvailabilityResponseMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncAvailabilityResponseMessageBuilder) Build() *BlockSyncAvailabilityResponseMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncAvailabilityResponseMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestMessage

// reader

type BlockSyncRequestMessage struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _BlockSyncRequestMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var _BlockSyncRequestMessage_Unions = [][]membuffers.FieldType{}

func BlockSyncRequestMessageReader(buf []byte) *BlockSyncRequestMessage {
	x := &BlockSyncRequestMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncRequestMessage_Scheme, _BlockSyncRequestMessage_Unions)
	return x
}

func (x *BlockSyncRequestMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncRequestMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncRequestMessage) Body() *BlockSyncRequestSignedFields {
	b, s := x._message.GetMessage(0)
	return BlockSyncRequestSignedFieldsReader(b[:s])
}

func (x *BlockSyncRequestMessage) RawBody() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncRequestMessage) Signature() primitives.Ed25519Sig {
	return x._message.GetBytes(1)
}

func (x *BlockSyncRequestMessage) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncRequestMessage) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, v)
}

// builder

type BlockSyncRequestMessageBuilder struct {
	Body *BlockSyncRequestSignedFieldsBuilder
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncRequestMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.Body)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *BlockSyncRequestMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncRequestMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncRequestMessageBuilder) Build() *BlockSyncRequestMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncRequestMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseMessage

// reader

type BlockSyncResponseMessage struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _BlockSyncResponseMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytesArray,}
var _BlockSyncResponseMessage_Unions = [][]membuffers.FieldType{}

func BlockSyncResponseMessageReader(buf []byte) *BlockSyncResponseMessage {
	x := &BlockSyncResponseMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncResponseMessage_Scheme, _BlockSyncResponseMessage_Unions)
	return x
}

func (x *BlockSyncResponseMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncResponseMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncResponseMessage) Body() *BlockSyncResponseSignedFields {
	b, s := x._message.GetMessage(0)
	return BlockSyncResponseSignedFieldsReader(b[:s])
}

func (x *BlockSyncResponseMessage) RawBody() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncResponseMessage) Signature() primitives.Ed25519Sig {
	return x._message.GetBytes(1)
}

func (x *BlockSyncResponseMessage) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncResponseMessage) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, v)
}

func (x *BlockSyncResponseMessage) BlockDataIterator() *BlockSyncResponseMessageBlockDataIterator {
	return &BlockSyncResponseMessageBlockDataIterator{iterator: x._message.GetBytesArrayIterator(2)}
}

type BlockSyncResponseMessageBlockDataIterator struct {
	iterator *membuffers.Iterator
}

func (i *BlockSyncResponseMessageBlockDataIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *BlockSyncResponseMessageBlockDataIterator) NextBlockData() []byte {
	return i.iterator.NextBytes()
}

func (x *BlockSyncResponseMessage) RawBlockDataArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

// builder

type BlockSyncResponseMessageBuilder struct {
	Body *BlockSyncResponseSignedFieldsBuilder
	Signature primitives.Ed25519Sig
	BlockData [][]byte

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncResponseMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.Body)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, w.Signature)
	w._builder.WriteBytesArray(buf, w.BlockData)
	return nil
}

func (w *BlockSyncResponseMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncResponseMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncResponseMessageBuilder) Build() *BlockSyncResponseMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncResponseMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncAvailabilitySignedFields

// reader

type BlockSyncAvailabilitySignedFields struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _BlockSyncAvailabilitySignedFields_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeBytes,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,}
var _BlockSyncAvailabilitySignedFields_Unions = [][]membuffers.FieldType{}

func BlockSyncAvailabilitySignedFieldsReader(buf []byte) *BlockSyncAvailabilitySignedFields {
	x := &BlockSyncAvailabilitySignedFields{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncAvailabilitySignedFields_Scheme, _BlockSyncAvailabilitySignedFields_Unions)
	return x
}

func (x *BlockSyncAvailabilitySignedFields) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncAvailabilitySignedFields) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncAvailabilitySignedFields) BlockType() BlockType {
	return BlockType(x._message.GetUint16(0))
}

func (x *BlockSyncAvailabilitySignedFields) RawBlockType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncAvailabilitySignedFields) MutateBlockType(v BlockType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *BlockSyncAvailabilitySignedFields) SenderPublicKey() primitives.Ed25519Pkey {
	return x._message.GetBytes(1)
}

func (x *BlockSyncAvailabilitySignedFields) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncAvailabilitySignedFields) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(1, v)
}

func (x *BlockSyncAvailabilitySignedFields) FirstAvailableBlockHeight() uint64 {
	return x._message.GetUint64(2)
}

func (x *BlockSyncAvailabilitySignedFields) RawFirstAvailableBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BlockSyncAvailabilitySignedFields) MutateFirstAvailableBlockHeight(v uint64) error {
	return x._message.SetUint64(2, v)
}

func (x *BlockSyncAvailabilitySignedFields) LastAvailableBlockHeight() uint64 {
	return x._message.GetUint64(3)
}

func (x *BlockSyncAvailabilitySignedFields) RawLastAvailableBlockHeight() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *BlockSyncAvailabilitySignedFields) MutateLastAvailableBlockHeight(v uint64) error {
	return x._message.SetUint64(3, v)
}

func (x *BlockSyncAvailabilitySignedFields) LastCommittedBlockHeight() uint64 {
	return x._message.GetUint64(4)
}

func (x *BlockSyncAvailabilitySignedFields) RawLastCommittedBlockHeight() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *BlockSyncAvailabilitySignedFields) MutateLastCommittedBlockHeight(v uint64) error {
	return x._message.SetUint64(4, v)
}

// builder

type BlockSyncAvailabilitySignedFieldsBuilder struct {
	BlockType BlockType
	SenderPublicKey primitives.Ed25519Pkey
	FirstAvailableBlockHeight uint64
	LastAvailableBlockHeight uint64
	LastCommittedBlockHeight uint64

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncAvailabilitySignedFieldsBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, w.SenderPublicKey)
	w._builder.WriteUint64(buf, w.FirstAvailableBlockHeight)
	w._builder.WriteUint64(buf, w.LastAvailableBlockHeight)
	w._builder.WriteUint64(buf, w.LastCommittedBlockHeight)
	return nil
}

func (w *BlockSyncAvailabilitySignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncAvailabilitySignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncAvailabilitySignedFieldsBuilder) Build() *BlockSyncAvailabilitySignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncAvailabilitySignedFieldsReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestSignedFields

// reader

type BlockSyncRequestSignedFields struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _BlockSyncRequestSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeBytes,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,}
var _BlockSyncRequestSignedFields_Unions = [][]membuffers.FieldType{}

func BlockSyncRequestSignedFieldsReader(buf []byte) *BlockSyncRequestSignedFields {
	x := &BlockSyncRequestSignedFields{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncRequestSignedFields_Scheme, _BlockSyncRequestSignedFields_Unions)
	return x
}

func (x *BlockSyncRequestSignedFields) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncRequestSignedFields) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncRequestSignedFields) BlockType() BlockType {
	return BlockType(x._message.GetUint16(0))
}

func (x *BlockSyncRequestSignedFields) RawBlockType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncRequestSignedFields) MutateBlockType(v BlockType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *BlockSyncRequestSignedFields) SenderPublicKey() primitives.Ed25519Pkey {
	return x._message.GetBytes(1)
}

func (x *BlockSyncRequestSignedFields) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncRequestSignedFields) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(1, v)
}

func (x *BlockSyncRequestSignedFields) FirstRequestedBlockHeight() uint64 {
	return x._message.GetUint64(2)
}

func (x *BlockSyncRequestSignedFields) RawFirstRequestedBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BlockSyncRequestSignedFields) MutateFirstRequestedBlockHeight(v uint64) error {
	return x._message.SetUint64(2, v)
}

func (x *BlockSyncRequestSignedFields) LastRequestedBlockHeight() uint64 {
	return x._message.GetUint64(3)
}

func (x *BlockSyncRequestSignedFields) RawLastRequestedBlockHeight() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *BlockSyncRequestSignedFields) MutateLastRequestedBlockHeight(v uint64) error {
	return x._message.SetUint64(3, v)
}

func (x *BlockSyncRequestSignedFields) LastCommittedBlockHeight() uint64 {
	return x._message.GetUint64(4)
}

func (x *BlockSyncRequestSignedFields) RawLastCommittedBlockHeight() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *BlockSyncRequestSignedFields) MutateLastCommittedBlockHeight(v uint64) error {
	return x._message.SetUint64(4, v)
}

// builder

type BlockSyncRequestSignedFieldsBuilder struct {
	BlockType BlockType
	SenderPublicKey primitives.Ed25519Pkey
	FirstRequestedBlockHeight uint64
	LastRequestedBlockHeight uint64
	LastCommittedBlockHeight uint64

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncRequestSignedFieldsBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, w.SenderPublicKey)
	w._builder.WriteUint64(buf, w.FirstRequestedBlockHeight)
	w._builder.WriteUint64(buf, w.LastRequestedBlockHeight)
	w._builder.WriteUint64(buf, w.LastCommittedBlockHeight)
	return nil
}

func (w *BlockSyncRequestSignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncRequestSignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncRequestSignedFieldsBuilder) Build() *BlockSyncRequestSignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncRequestSignedFieldsReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseSignedFields

// reader

type BlockSyncResponseSignedFields struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _BlockSyncResponseSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeBytes,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeBytesArray,}
var _BlockSyncResponseSignedFields_Unions = [][]membuffers.FieldType{}

func BlockSyncResponseSignedFieldsReader(buf []byte) *BlockSyncResponseSignedFields {
	x := &BlockSyncResponseSignedFields{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncResponseSignedFields_Scheme, _BlockSyncResponseSignedFields_Unions)
	return x
}

func (x *BlockSyncResponseSignedFields) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncResponseSignedFields) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncResponseSignedFields) BlockType() BlockType {
	return BlockType(x._message.GetUint16(0))
}

func (x *BlockSyncResponseSignedFields) RawBlockType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncResponseSignedFields) MutateBlockType(v BlockType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *BlockSyncResponseSignedFields) SenderPublicKey() primitives.Ed25519Pkey {
	return x._message.GetBytes(1)
}

func (x *BlockSyncResponseSignedFields) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncResponseSignedFields) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(1, v)
}

func (x *BlockSyncResponseSignedFields) FirstBlockHeight() uint64 {
	return x._message.GetUint64(2)
}

func (x *BlockSyncResponseSignedFields) RawFirstBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BlockSyncResponseSignedFields) MutateFirstBlockHeight(v uint64) error {
	return x._message.SetUint64(2, v)
}

func (x *BlockSyncResponseSignedFields) LastBlockHeight() uint64 {
	return x._message.GetUint64(3)
}

func (x *BlockSyncResponseSignedFields) RawLastBlockHeight() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *BlockSyncResponseSignedFields) MutateLastBlockHeight(v uint64) error {
	return x._message.SetUint64(3, v)
}

func (x *BlockSyncResponseSignedFields) LastCommittedBlockHeight() uint64 {
	return x._message.GetUint64(4)
}

func (x *BlockSyncResponseSignedFields) RawLastCommittedBlockHeight() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *BlockSyncResponseSignedFields) MutateLastCommittedBlockHeight(v uint64) error {
	return x._message.SetUint64(4, v)
}

func (x *BlockSyncResponseSignedFields) BlockHashIterator() *BlockSyncResponseSignedFieldsBlockHashIterator {
	return &BlockSyncResponseSignedFieldsBlockHashIterator{iterator: x._message.GetBytesArrayIterator(5)}
}

type BlockSyncResponseSignedFieldsBlockHashIterator struct {
	iterator *membuffers.Iterator
}

func (i *BlockSyncResponseSignedFieldsBlockHashIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *BlockSyncResponseSignedFieldsBlockHashIterator) NextBlockHash() primitives.Sha256 {
	return i.iterator.NextBytes()
}

func (x *BlockSyncResponseSignedFields) RawBlockHashArray() []byte {
	return x._message.RawBufferForField(5, 0)
}

// builder

type BlockSyncResponseSignedFieldsBuilder struct {
	BlockType BlockType
	SenderPublicKey primitives.Ed25519Pkey
	FirstBlockHeight uint64
	LastBlockHeight uint64
	LastCommittedBlockHeight uint64
	BlockHash []primitives.Sha256

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncResponseSignedFieldsBuilder) arrayOfBlockHash() [][]byte {
	res := make([][]byte, len(w.BlockHash))
	for i, v := range w.BlockHash {
		res[i] = v
	}
	return res
}

func (w *BlockSyncResponseSignedFieldsBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, w.SenderPublicKey)
	w._builder.WriteUint64(buf, w.FirstBlockHeight)
	w._builder.WriteUint64(buf, w.LastBlockHeight)
	w._builder.WriteUint64(buf, w.LastCommittedBlockHeight)
	w._builder.WriteBytesArray(buf, w.arrayOfBlockHash())
	return nil
}

func (w *BlockSyncResponseSignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncResponseSignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncResponseSignedFieldsBuilder) Build() *BlockSyncResponseSignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncResponseSignedFieldsReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type BlockSyncMessageType uint16

const (
	BLOCK_SYNC_RESERVED BlockSyncMessageType = 0
	BLOCK_SYNC_REQUEST BlockSyncMessageType = 1
	BLOCK_SYNC_RESPONSE BlockSyncMessageType = 2
	BLOCK_SYNC_BLOCK_AVAILABILITY_REQUEST BlockSyncMessageType = 3
	BLOCK_SYNC_BLOCK_AVAILABILITY_RESPONSE BlockSyncMessageType = 4
)

type BlockType uint16

const (
	BLOCK_TYPE_RESERVED BlockType = 0
	BLOCK_TYPE_BLOCK_PAIR BlockType = 1
	BLOCK_TYPE_TRANSACTIONS_BLOCK BlockType = 2
	BLOCK_TYPE_RESULTS_BLOCK BlockType = 3
)

