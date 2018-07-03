// AUTO GENERATED FILE (by membufc proto compiler)
package blocksync

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message Message

// reader

type Message struct {
	message membuffers.Message
}

var m_Message_Scheme = []membuffers.FieldType{membuffers.TypeUnion,}
var m_Message_Unions = [][]membuffers.FieldType{{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}}

func MessageReader(buf []byte) *Message {
	x := &Message{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_Message_Scheme, m_Message_Unions)
	return x
}

func (x *Message) IsValid() bool {
	return x.message.IsValid()
}

func (x *Message) Raw() []byte {
	return x.message.RawBuffer()
}

type MessageBlockSyncMessage uint16

const (
	MessageBlockSyncMessageBlockAvailabilityRequest MessageBlockSyncMessage = 0
	MessageBlockSyncMessageBlockAvailabilityResponse MessageBlockSyncMessage = 1
	MessageBlockSyncMessageBlockSyncRequest MessageBlockSyncMessage = 2
	MessageBlockSyncMessageBlockSyncResponse MessageBlockSyncMessage = 3
)

func (x *Message) BlockSyncMessage() MessageBlockSyncMessage {
	return MessageBlockSyncMessage(x.message.GetUint16(0))
}

func (x *Message) IsBlockSyncMessageBlockAvailabilityRequest() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 0)
	return is
}

func (x *Message) BlockSyncMessageBlockAvailabilityRequest() BlockAvailabilityRequestMessage {
	_, off := x.message.IsUnionIndex(0, 0, 0)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) IsBlockSyncMessageBlockAvailabilityResponse() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 1)
	return is
}

func (x *Message) BlockSyncMessageBlockAvailabilityResponse() BlockAvailabilityRequestMessage {
	_, off := x.message.IsUnionIndex(0, 0, 1)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) IsBlockSyncMessageBlockSyncRequest() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 2)
	return is
}

func (x *Message) BlockSyncMessageBlockSyncRequest() BlockSyncRequestMessage {
	_, off := x.message.IsUnionIndex(0, 0, 2)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) IsBlockSyncMessageBlockSyncResponse() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 3)
	return is
}

func (x *Message) BlockSyncMessageBlockSyncResponse() BlockSyncResponseMessage {
	_, off := x.message.IsUnionIndex(0, 0, 3)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) RawBlockSyncMessage() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type MessageBuilder struct {
	builder membuffers.Builder
	BlockSyncMessage MessageBlockSyncMessage
	BlockAvailabilityRequest BlockAvailabilityRequestMessage
	BlockAvailabilityResponse BlockAvailabilityRequestMessage
	BlockSyncRequest BlockSyncRequestMessage
	BlockSyncResponse BlockSyncResponseMessage
}

func (w *MessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUnionIndex(buf, uint16(w.BlockSyncMessage))
	switch w.BlockSyncMessage {
	case MessageBlockSyncMessageBlockAvailabilityRequest:
		w.builder.WriteMessage(buf, w.BlockAvailabilityRequest)
	case MessageBlockSyncMessageBlockAvailabilityResponse:
		w.builder.WriteMessage(buf, w.BlockAvailabilityResponse)
	case MessageBlockSyncMessageBlockSyncRequest:
		w.builder.WriteMessage(buf, w.BlockSyncRequest)
	case MessageBlockSyncMessageBlockSyncResponse:
		w.builder.WriteMessage(buf, w.BlockSyncResponse)
	}
	return nil
}

func (w *MessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *MessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *MessageBuilder) Build() *Message {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return MessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestMessage

// reader

type BlockSyncRequestMessage struct {
	message membuffers.Message
}

var m_BlockSyncRequestMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var m_BlockSyncRequestMessage_Unions = [][]membuffers.FieldType{}

func BlockSyncRequestMessageReader(buf []byte) *BlockSyncRequestMessage {
	x := &BlockSyncRequestMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_BlockSyncRequestMessage_Scheme, m_BlockSyncRequestMessage_Unions)
	return x
}

func (x *BlockSyncRequestMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *BlockSyncRequestMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *BlockSyncRequestMessage) BlockSyncRequestSignedFields() *BlockRequestSignedFields {
	b, s := x.message.GetMessage(0)
	return BlockRequestSignedFieldsReader(b[:s])
}

func (x *BlockSyncRequestMessage) RawBlockSyncRequestSignedFields() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *BlockSyncRequestMessage) Signature() []byte {
	return x.message.GetBytes(1)
}

func (x *BlockSyncRequestMessage) RawSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *BlockSyncRequestMessage) MutateSignature(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type BlockSyncRequestMessageBuilder struct {
	builder membuffers.Builder
	BlockSyncRequestSignedFields *BlockRequestSignedFieldsBuilder
	Signature []byte
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
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.BlockSyncRequestSignedFields)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *BlockSyncRequestMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *BlockSyncRequestMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *BlockSyncRequestMessageBuilder) Build() *BlockSyncRequestMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncRequestMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockRequestSignedFields

// reader

type BlockRequestSignedFields struct {
	message membuffers.Message
}

var m_BlockRequestSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeBytes,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,}
var m_BlockRequestSignedFields_Unions = [][]membuffers.FieldType{}

func BlockRequestSignedFieldsReader(buf []byte) *BlockRequestSignedFields {
	x := &BlockRequestSignedFields{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_BlockRequestSignedFields_Scheme, m_BlockRequestSignedFields_Unions)
	return x
}

func (x *BlockRequestSignedFields) IsValid() bool {
	return x.message.IsValid()
}

func (x *BlockRequestSignedFields) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *BlockRequestSignedFields) MessageType() BlockSyncMessageType {
	return BlockSyncMessageType(x.message.GetUint16(0))
}

func (x *BlockRequestSignedFields) RawMessageType() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *BlockRequestSignedFields) MutateMessageType(v BlockSyncMessageType) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *BlockRequestSignedFields) BlockType() BlockType {
	return BlockType(x.message.GetUint16(1))
}

func (x *BlockRequestSignedFields) RawBlockType() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *BlockRequestSignedFields) MutateBlockType(v BlockType) error {
	return x.message.SetUint16(1, uint16(v))
}

func (x *BlockRequestSignedFields) SenderPublicKey() []byte {
	return x.message.GetBytes(2)
}

func (x *BlockRequestSignedFields) RawSenderPublicKey() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *BlockRequestSignedFields) MutateSenderPublicKey(v []byte) error {
	return x.message.SetBytes(2, v)
}

func (x *BlockRequestSignedFields) FirstRequestedBlockHeight() uint64 {
	return x.message.GetUint64(3)
}

func (x *BlockRequestSignedFields) RawFirstRequestedBlockHeight() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *BlockRequestSignedFields) MutateFirstRequestedBlockHeight(v uint64) error {
	return x.message.SetUint64(3, v)
}

func (x *BlockRequestSignedFields) LastRequestedBlockHeight() uint64 {
	return x.message.GetUint64(4)
}

func (x *BlockRequestSignedFields) RawLastRequestedBlockHeight() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *BlockRequestSignedFields) MutateLastRequestedBlockHeight(v uint64) error {
	return x.message.SetUint64(4, v)
}

func (x *BlockRequestSignedFields) LastCommittedBlockHeight() uint64 {
	return x.message.GetUint64(5)
}

func (x *BlockRequestSignedFields) RawLastCommittedBlockHeight() []byte {
	return x.message.RawBufferForField(5, 0)
}

func (x *BlockRequestSignedFields) MutateLastCommittedBlockHeight(v uint64) error {
	return x.message.SetUint64(5, v)
}

// builder

type BlockRequestSignedFieldsBuilder struct {
	builder membuffers.Builder
	MessageType BlockSyncMessageType
	BlockType BlockType
	SenderPublicKey []byte
	FirstRequestedBlockHeight uint64
	LastRequestedBlockHeight uint64
	LastCommittedBlockHeight uint64
}

func (w *BlockRequestSignedFieldsBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.MessageType))
	w.builder.WriteUint16(buf, uint16(w.BlockType))
	w.builder.WriteBytes(buf, w.SenderPublicKey)
	w.builder.WriteUint64(buf, w.FirstRequestedBlockHeight)
	w.builder.WriteUint64(buf, w.LastRequestedBlockHeight)
	w.builder.WriteUint64(buf, w.LastCommittedBlockHeight)
	return nil
}

func (w *BlockRequestSignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *BlockRequestSignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *BlockRequestSignedFieldsBuilder) Build() *BlockRequestSignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockRequestSignedFieldsReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseMessage

// reader

type BlockSyncResponseMessage struct {
	message membuffers.Message
}

var m_BlockSyncResponseMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytesArray,}
var m_BlockSyncResponseMessage_Unions = [][]membuffers.FieldType{}

func BlockSyncResponseMessageReader(buf []byte) *BlockSyncResponseMessage {
	x := &BlockSyncResponseMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_BlockSyncResponseMessage_Scheme, m_BlockSyncResponseMessage_Unions)
	return x
}

func (x *BlockSyncResponseMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *BlockSyncResponseMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *BlockSyncResponseMessage) BlockSyncResponseSignedFields() *BlockSyncResponseSignedFields {
	b, s := x.message.GetMessage(0)
	return BlockSyncResponseSignedFieldsReader(b[:s])
}

func (x *BlockSyncResponseMessage) RawBlockSyncResponseSignedFields() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *BlockSyncResponseMessage) Signature() []byte {
	return x.message.GetBytes(1)
}

func (x *BlockSyncResponseMessage) RawSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *BlockSyncResponseMessage) MutateSignature(v []byte) error {
	return x.message.SetBytes(1, v)
}

func (x *BlockSyncResponseMessage) BlockDataIterator() *BlockSyncResponseMessageBlockDataIterator {
	return &BlockSyncResponseMessageBlockDataIterator{iterator: x.message.GetBytesArrayIterator(2)}
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
	return x.message.RawBufferForField(2, 0)
}

// builder

type BlockSyncResponseMessageBuilder struct {
	builder membuffers.Builder
	BlockSyncResponseSignedFields *BlockSyncResponseSignedFieldsBuilder
	Signature []byte
	BlockData [][]byte
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
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.BlockSyncResponseSignedFields)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.Signature)
	w.builder.WriteBytesArray(buf, w.BlockData)
	return nil
}

func (w *BlockSyncResponseMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *BlockSyncResponseMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *BlockSyncResponseMessageBuilder) Build() *BlockSyncResponseMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncResponseMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseSignedFields

// reader

type BlockSyncResponseSignedFields struct {
	message membuffers.Message
}

var m_BlockSyncResponseSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeBytes,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeBytesArray,}
var m_BlockSyncResponseSignedFields_Unions = [][]membuffers.FieldType{}

func BlockSyncResponseSignedFieldsReader(buf []byte) *BlockSyncResponseSignedFields {
	x := &BlockSyncResponseSignedFields{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_BlockSyncResponseSignedFields_Scheme, m_BlockSyncResponseSignedFields_Unions)
	return x
}

func (x *BlockSyncResponseSignedFields) IsValid() bool {
	return x.message.IsValid()
}

func (x *BlockSyncResponseSignedFields) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *BlockSyncResponseSignedFields) MessageType() BlockSyncMessageType {
	return BlockSyncMessageType(x.message.GetUint16(0))
}

func (x *BlockSyncResponseSignedFields) RawMessageType() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *BlockSyncResponseSignedFields) MutateMessageType(v BlockSyncMessageType) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *BlockSyncResponseSignedFields) BlockType() BlockType {
	return BlockType(x.message.GetUint16(1))
}

func (x *BlockSyncResponseSignedFields) RawBlockType() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *BlockSyncResponseSignedFields) MutateBlockType(v BlockType) error {
	return x.message.SetUint16(1, uint16(v))
}

func (x *BlockSyncResponseSignedFields) SenderPublicKey() []byte {
	return x.message.GetBytes(2)
}

func (x *BlockSyncResponseSignedFields) RawSenderPublicKey() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *BlockSyncResponseSignedFields) MutateSenderPublicKey(v []byte) error {
	return x.message.SetBytes(2, v)
}

func (x *BlockSyncResponseSignedFields) FirstBlockHeight() uint64 {
	return x.message.GetUint64(3)
}

func (x *BlockSyncResponseSignedFields) RawFirstBlockHeight() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *BlockSyncResponseSignedFields) MutateFirstBlockHeight(v uint64) error {
	return x.message.SetUint64(3, v)
}

func (x *BlockSyncResponseSignedFields) LastBlockHeight() uint64 {
	return x.message.GetUint64(4)
}

func (x *BlockSyncResponseSignedFields) RawLastBlockHeight() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *BlockSyncResponseSignedFields) MutateLastBlockHeight(v uint64) error {
	return x.message.SetUint64(4, v)
}

func (x *BlockSyncResponseSignedFields) LastCommittedBlockHeight() uint64 {
	return x.message.GetUint64(5)
}

func (x *BlockSyncResponseSignedFields) RawLastCommittedBlockHeight() []byte {
	return x.message.RawBufferForField(5, 0)
}

func (x *BlockSyncResponseSignedFields) MutateLastCommittedBlockHeight(v uint64) error {
	return x.message.SetUint64(5, v)
}

func (x *BlockSyncResponseSignedFields) BlockHashIterator() *BlockSyncResponseSignedFieldsBlockHashIterator {
	return &BlockSyncResponseSignedFieldsBlockHashIterator{iterator: x.message.GetBytesArrayIterator(6)}
}

type BlockSyncResponseSignedFieldsBlockHashIterator struct {
	iterator *membuffers.Iterator
}

func (i *BlockSyncResponseSignedFieldsBlockHashIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *BlockSyncResponseSignedFieldsBlockHashIterator) NextBlockHash() []byte {
	return i.iterator.NextBytes()
}

func (x *BlockSyncResponseSignedFields) RawBlockHashArray() []byte {
	return x.message.RawBufferForField(6, 0)
}

// builder

type BlockSyncResponseSignedFieldsBuilder struct {
	builder membuffers.Builder
	MessageType BlockSyncMessageType
	BlockType BlockType
	SenderPublicKey []byte
	FirstBlockHeight uint64
	LastBlockHeight uint64
	LastCommittedBlockHeight uint64
	BlockHash [][]byte
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
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.MessageType))
	w.builder.WriteUint16(buf, uint16(w.BlockType))
	w.builder.WriteBytes(buf, w.SenderPublicKey)
	w.builder.WriteUint64(buf, w.FirstBlockHeight)
	w.builder.WriteUint64(buf, w.LastBlockHeight)
	w.builder.WriteUint64(buf, w.LastCommittedBlockHeight)
	w.builder.WriteBytesArray(buf, w.BlockHash)
	return nil
}

func (w *BlockSyncResponseSignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *BlockSyncResponseSignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *BlockSyncResponseSignedFieldsBuilder) Build() *BlockSyncResponseSignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncResponseSignedFieldsReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilityRequestMessage

// reader

type BlockAvailabilityRequestMessage struct {
	message membuffers.Message
}

var m_BlockAvailabilityRequestMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var m_BlockAvailabilityRequestMessage_Unions = [][]membuffers.FieldType{}

func BlockAvailabilityRequestMessageReader(buf []byte) *BlockAvailabilityRequestMessage {
	x := &BlockAvailabilityRequestMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_BlockAvailabilityRequestMessage_Scheme, m_BlockAvailabilityRequestMessage_Unions)
	return x
}

func (x *BlockAvailabilityRequestMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *BlockAvailabilityRequestMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *BlockAvailabilityRequestMessage) BlockAvailabilityRequestSignedFields() *BlockRequestSignedFields {
	b, s := x.message.GetMessage(0)
	return BlockRequestSignedFieldsReader(b[:s])
}

func (x *BlockAvailabilityRequestMessage) RawBlockAvailabilityRequestSignedFields() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *BlockAvailabilityRequestMessage) Signature() []byte {
	return x.message.GetBytes(1)
}

func (x *BlockAvailabilityRequestMessage) RawSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *BlockAvailabilityRequestMessage) MutateSignature(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type BlockAvailabilityRequestMessageBuilder struct {
	builder membuffers.Builder
	BlockAvailabilityRequestSignedFields *BlockRequestSignedFieldsBuilder
	Signature []byte
}

func (w *BlockAvailabilityRequestMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.BlockAvailabilityRequestSignedFields)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *BlockAvailabilityRequestMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *BlockAvailabilityRequestMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *BlockAvailabilityRequestMessageBuilder) Build() *BlockAvailabilityRequestMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockAvailabilityRequestMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilityResponseMessage

// reader

type BlockAvailabilityResponseMessage struct {
	message membuffers.Message
}

var m_BlockAvailabilityResponseMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var m_BlockAvailabilityResponseMessage_Unions = [][]membuffers.FieldType{}

func BlockAvailabilityResponseMessageReader(buf []byte) *BlockAvailabilityResponseMessage {
	x := &BlockAvailabilityResponseMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_BlockAvailabilityResponseMessage_Scheme, m_BlockAvailabilityResponseMessage_Unions)
	return x
}

func (x *BlockAvailabilityResponseMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *BlockAvailabilityResponseMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *BlockAvailabilityResponseMessage) BlockAvailabilityResponseSignedFields() *BlockSyncAvailabilitySignedFields {
	b, s := x.message.GetMessage(0)
	return BlockSyncAvailabilitySignedFieldsReader(b[:s])
}

func (x *BlockAvailabilityResponseMessage) RawBlockAvailabilityResponseSignedFields() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *BlockAvailabilityResponseMessage) Signature() []byte {
	return x.message.GetBytes(1)
}

func (x *BlockAvailabilityResponseMessage) RawSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *BlockAvailabilityResponseMessage) MutateSignature(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type BlockAvailabilityResponseMessageBuilder struct {
	builder membuffers.Builder
	BlockAvailabilityResponseSignedFields *BlockSyncAvailabilitySignedFieldsBuilder
	Signature []byte
}

func (w *BlockAvailabilityResponseMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.BlockAvailabilityResponseSignedFields)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *BlockAvailabilityResponseMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *BlockAvailabilityResponseMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *BlockAvailabilityResponseMessageBuilder) Build() *BlockAvailabilityResponseMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockAvailabilityResponseMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncAvailabilitySignedFields

// reader

type BlockSyncAvailabilitySignedFields struct {
	message membuffers.Message
}

var m_BlockSyncAvailabilitySignedFields_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeBytes,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,}
var m_BlockSyncAvailabilitySignedFields_Unions = [][]membuffers.FieldType{}

func BlockSyncAvailabilitySignedFieldsReader(buf []byte) *BlockSyncAvailabilitySignedFields {
	x := &BlockSyncAvailabilitySignedFields{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_BlockSyncAvailabilitySignedFields_Scheme, m_BlockSyncAvailabilitySignedFields_Unions)
	return x
}

func (x *BlockSyncAvailabilitySignedFields) IsValid() bool {
	return x.message.IsValid()
}

func (x *BlockSyncAvailabilitySignedFields) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *BlockSyncAvailabilitySignedFields) MessageType() BlockSyncMessageType {
	return BlockSyncMessageType(x.message.GetUint16(0))
}

func (x *BlockSyncAvailabilitySignedFields) RawMessageType() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *BlockSyncAvailabilitySignedFields) MutateMessageType(v BlockSyncMessageType) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *BlockSyncAvailabilitySignedFields) BlockType() BlockType {
	return BlockType(x.message.GetUint16(1))
}

func (x *BlockSyncAvailabilitySignedFields) RawBlockType() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *BlockSyncAvailabilitySignedFields) MutateBlockType(v BlockType) error {
	return x.message.SetUint16(1, uint16(v))
}

func (x *BlockSyncAvailabilitySignedFields) SenderPublicKey() []byte {
	return x.message.GetBytes(2)
}

func (x *BlockSyncAvailabilitySignedFields) RawSenderPublicKey() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *BlockSyncAvailabilitySignedFields) MutateSenderPublicKey(v []byte) error {
	return x.message.SetBytes(2, v)
}

func (x *BlockSyncAvailabilitySignedFields) FirstAvailableBlockHeight() uint64 {
	return x.message.GetUint64(3)
}

func (x *BlockSyncAvailabilitySignedFields) RawFirstAvailableBlockHeight() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *BlockSyncAvailabilitySignedFields) MutateFirstAvailableBlockHeight(v uint64) error {
	return x.message.SetUint64(3, v)
}

func (x *BlockSyncAvailabilitySignedFields) LastAvailableBlockHeight() uint64 {
	return x.message.GetUint64(4)
}

func (x *BlockSyncAvailabilitySignedFields) RawLastAvailableBlockHeight() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *BlockSyncAvailabilitySignedFields) MutateLastAvailableBlockHeight(v uint64) error {
	return x.message.SetUint64(4, v)
}

func (x *BlockSyncAvailabilitySignedFields) LastCommittedBlockHeight() uint64 {
	return x.message.GetUint64(5)
}

func (x *BlockSyncAvailabilitySignedFields) RawLastCommittedBlockHeight() []byte {
	return x.message.RawBufferForField(5, 0)
}

func (x *BlockSyncAvailabilitySignedFields) MutateLastCommittedBlockHeight(v uint64) error {
	return x.message.SetUint64(5, v)
}

// builder

type BlockSyncAvailabilitySignedFieldsBuilder struct {
	builder membuffers.Builder
	MessageType BlockSyncMessageType
	BlockType BlockType
	SenderPublicKey []byte
	FirstAvailableBlockHeight uint64
	LastAvailableBlockHeight uint64
	LastCommittedBlockHeight uint64
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
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.MessageType))
	w.builder.WriteUint16(buf, uint16(w.BlockType))
	w.builder.WriteBytes(buf, w.SenderPublicKey)
	w.builder.WriteUint64(buf, w.FirstAvailableBlockHeight)
	w.builder.WriteUint64(buf, w.LastAvailableBlockHeight)
	w.builder.WriteUint64(buf, w.LastCommittedBlockHeight)
	return nil
}

func (w *BlockSyncAvailabilitySignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *BlockSyncAvailabilitySignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *BlockSyncAvailabilitySignedFieldsBuilder) Build() *BlockSyncAvailabilitySignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncAvailabilitySignedFieldsReader(buf)
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

