// AUTO GENERATED FILE (by membufc proto compiler)
package leanhelixconsensus

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
var m_Message_Unions = [][]membuffers.FieldType{{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}}

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

type MessageLeanHelixMessage uint16

const (
	MessageLeanHelixMessagePrePrepare MessageLeanHelixMessage = 0
	MessageLeanHelixMessagePrepare MessageLeanHelixMessage = 1
	MessageLeanHelixMessageCommit MessageLeanHelixMessage = 2
	MessageLeanHelixMessageViewChange MessageLeanHelixMessage = 3
	MessageLeanHelixMessageNewView MessageLeanHelixMessage = 4
)

func (x *Message) LeanHelixMessage() MessageLeanHelixMessage {
	return MessageLeanHelixMessage(x.message.GetUint16(0))
}

func (x *Message) IsLeanHelixMessagePrePrepare() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 0)
	return is
}

func (x *Message) LeanHelixMessagePrePrepare() PrePrepareMessage {
	_, off := x.message.IsUnionIndex(0, 0, 0)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) IsLeanHelixMessagePrepare() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 1)
	return is
}

func (x *Message) LeanHelixMessagePrepare() PrepareMessage {
	_, off := x.message.IsUnionIndex(0, 0, 1)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) IsLeanHelixMessageCommit() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 2)
	return is
}

func (x *Message) LeanHelixMessageCommit() CommitMessage {
	_, off := x.message.IsUnionIndex(0, 0, 2)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) IsLeanHelixMessageViewChange() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 3)
	return is
}

func (x *Message) LeanHelixMessageViewChange() ViewChangeMessage {
	_, off := x.message.IsUnionIndex(0, 0, 3)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) IsLeanHelixMessageNewView() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 4)
	return is
}

func (x *Message) LeanHelixMessageNewView() NewViewMessage {
	_, off := x.message.IsUnionIndex(0, 0, 4)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) RawLeanHelixMessage() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type MessageBuilder struct {
	builder membuffers.Builder
	LeanHelixMessage MessageLeanHelixMessage
	PrePrepare PrePrepareMessage
	Prepare PrepareMessage
	Commit CommitMessage
	ViewChange ViewChangeMessage
	NewView NewViewMessage
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
	w.builder.WriteUnionIndex(buf, uint16(w.LeanHelixMessage))
	switch w.LeanHelixMessage {
	case MessageLeanHelixMessagePrePrepare:
		w.builder.WriteMessage(buf, w.PrePrepare)
	case MessageLeanHelixMessagePrepare:
		w.builder.WriteMessage(buf, w.Prepare)
	case MessageLeanHelixMessageCommit:
		w.builder.WriteMessage(buf, w.Commit)
	case MessageLeanHelixMessageViewChange:
		w.builder.WriteMessage(buf, w.ViewChange)
	case MessageLeanHelixMessageNewView:
		w.builder.WriteMessage(buf, w.NewView)
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
// message PrePrepareMessage

// reader

type PrePrepareMessage struct {
	message membuffers.Message
}

var m_PrePrepareMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytes,}
var m_PrePrepareMessage_Unions = [][]membuffers.FieldType{}

func PrePrepareMessageReader(buf []byte) *PrePrepareMessage {
	x := &PrePrepareMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_PrePrepareMessage_Scheme, m_PrePrepareMessage_Unions)
	return x
}

func (x *PrePrepareMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *PrePrepareMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *PrePrepareMessage) PrePrepareSignedFields() *CommitPreparePPMessageSignedFields {
	b, s := x.message.GetMessage(0)
	return CommitPreparePPMessageSignedFieldsReader(b[:s])
}

func (x *PrePrepareMessage) RawPrePrepareSignedFields() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *PrePrepareMessage) Signature() []byte {
	return x.message.GetBytes(1)
}

func (x *PrePrepareMessage) RawSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *PrePrepareMessage) MutateSignature(v []byte) error {
	return x.message.SetBytes(1, v)
}

func (x *PrePrepareMessage) BlockPairData() []byte {
	return x.message.GetBytes(2)
}

func (x *PrePrepareMessage) RawBlockPairData() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *PrePrepareMessage) MutateBlockPairData(v []byte) error {
	return x.message.SetBytes(2, v)
}

// builder

type PrePrepareMessageBuilder struct {
	builder membuffers.Builder
	PrePrepareSignedFields *CommitPreparePPMessageSignedFieldsBuilder
	Signature []byte
	BlockPairData []byte
}

func (w *PrePrepareMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.PrePrepareSignedFields)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.Signature)
	w.builder.WriteBytes(buf, w.BlockPairData)
	return nil
}

func (w *PrePrepareMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *PrePrepareMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *PrePrepareMessageBuilder) Build() *PrePrepareMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return PrePrepareMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message PrepareMessage

// reader

type PrepareMessage struct {
	message membuffers.Message
}

var m_PrepareMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var m_PrepareMessage_Unions = [][]membuffers.FieldType{}

func PrepareMessageReader(buf []byte) *PrepareMessage {
	x := &PrepareMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_PrepareMessage_Scheme, m_PrepareMessage_Unions)
	return x
}

func (x *PrepareMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *PrepareMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *PrepareMessage) PrepareSignedFields() *CommitPreparePPMessageSignedFields {
	b, s := x.message.GetMessage(0)
	return CommitPreparePPMessageSignedFieldsReader(b[:s])
}

func (x *PrepareMessage) RawPrepareSignedFields() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *PrepareMessage) Signature() []byte {
	return x.message.GetBytes(1)
}

func (x *PrepareMessage) RawSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *PrepareMessage) MutateSignature(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type PrepareMessageBuilder struct {
	builder membuffers.Builder
	PrepareSignedFields *CommitPreparePPMessageSignedFieldsBuilder
	Signature []byte
}

func (w *PrepareMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.PrepareSignedFields)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *PrepareMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *PrepareMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *PrepareMessageBuilder) Build() *PrepareMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return PrepareMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message CommitMessage

// reader

type CommitMessage struct {
	message membuffers.Message
}

var m_CommitMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytes,}
var m_CommitMessage_Unions = [][]membuffers.FieldType{}

func CommitMessageReader(buf []byte) *CommitMessage {
	x := &CommitMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_CommitMessage_Scheme, m_CommitMessage_Unions)
	return x
}

func (x *CommitMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *CommitMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *CommitMessage) CommitSignedFields() *CommitPreparePPMessageSignedFields {
	b, s := x.message.GetMessage(0)
	return CommitPreparePPMessageSignedFieldsReader(b[:s])
}

func (x *CommitMessage) RawCommitSignedFields() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *CommitMessage) Signature() []byte {
	return x.message.GetBytes(1)
}

func (x *CommitMessage) RawSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *CommitMessage) MutateSignature(v []byte) error {
	return x.message.SetBytes(1, v)
}

func (x *CommitMessage) RandomSeedShare() []byte {
	return x.message.GetBytes(2)
}

func (x *CommitMessage) RawRandomSeedShare() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *CommitMessage) MutateRandomSeedShare(v []byte) error {
	return x.message.SetBytes(2, v)
}

// builder

type CommitMessageBuilder struct {
	builder membuffers.Builder
	CommitSignedFields *CommitPreparePPMessageSignedFieldsBuilder
	Signature []byte
	RandomSeedShare []byte
}

func (w *CommitMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.CommitSignedFields)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.Signature)
	w.builder.WriteBytes(buf, w.RandomSeedShare)
	return nil
}

func (w *CommitMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *CommitMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *CommitMessageBuilder) Build() *CommitMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return CommitMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ViewChangeMessage

// reader

type ViewChangeMessage struct {
	message membuffers.Message
}

var m_ViewChangeMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytes,}
var m_ViewChangeMessage_Unions = [][]membuffers.FieldType{}

func ViewChangeMessageReader(buf []byte) *ViewChangeMessage {
	x := &ViewChangeMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ViewChangeMessage_Scheme, m_ViewChangeMessage_Unions)
	return x
}

func (x *ViewChangeMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *ViewChangeMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ViewChangeMessage) ViewChangeSignedFields() *ViewChangeSignedFields {
	b, s := x.message.GetMessage(0)
	return ViewChangeSignedFieldsReader(b[:s])
}

func (x *ViewChangeMessage) RawViewChangeSignedFields() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ViewChangeMessage) Signature() []byte {
	return x.message.GetBytes(1)
}

func (x *ViewChangeMessage) RawSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ViewChangeMessage) MutateSignature(v []byte) error {
	return x.message.SetBytes(1, v)
}

func (x *ViewChangeMessage) BlockPairData() []byte {
	return x.message.GetBytes(2)
}

func (x *ViewChangeMessage) RawBlockPairData() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *ViewChangeMessage) MutateBlockPairData(v []byte) error {
	return x.message.SetBytes(2, v)
}

// builder

type ViewChangeMessageBuilder struct {
	builder membuffers.Builder
	ViewChangeSignedFields *ViewChangeSignedFieldsBuilder
	Signature []byte
	BlockPairData []byte
}

func (w *ViewChangeMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.ViewChangeSignedFields)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.Signature)
	w.builder.WriteBytes(buf, w.BlockPairData)
	return nil
}

func (w *ViewChangeMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ViewChangeMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ViewChangeMessageBuilder) Build() *ViewChangeMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ViewChangeMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message NewViewMessage

// reader

type NewViewMessage struct {
	message membuffers.Message
}

var m_NewViewMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytes,}
var m_NewViewMessage_Unions = [][]membuffers.FieldType{}

func NewViewMessageReader(buf []byte) *NewViewMessage {
	x := &NewViewMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_NewViewMessage_Scheme, m_NewViewMessage_Unions)
	return x
}

func (x *NewViewMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *NewViewMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *NewViewMessage) NewViewSignedFields() *NewViewProof {
	b, s := x.message.GetMessage(0)
	return NewViewProofReader(b[:s])
}

func (x *NewViewMessage) RawNewViewSignedFields() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *NewViewMessage) HeaderSignature() []byte {
	return x.message.GetBytes(1)
}

func (x *NewViewMessage) RawHeaderSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *NewViewMessage) MutateHeaderSignature(v []byte) error {
	return x.message.SetBytes(1, v)
}

func (x *NewViewMessage) BlockPairData() []byte {
	return x.message.GetBytes(2)
}

func (x *NewViewMessage) RawBlockPairData() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *NewViewMessage) MutateBlockPairData(v []byte) error {
	return x.message.SetBytes(2, v)
}

// builder

type NewViewMessageBuilder struct {
	builder membuffers.Builder
	NewViewSignedFields *NewViewProofBuilder
	HeaderSignature []byte
	BlockPairData []byte
}

func (w *NewViewMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.NewViewSignedFields)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.HeaderSignature)
	w.builder.WriteBytes(buf, w.BlockPairData)
	return nil
}

func (w *NewViewMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *NewViewMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *NewViewMessageBuilder) Build() *NewViewMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return NewViewMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message CommitPreparePPMessageSignedFields

// reader

type CommitPreparePPMessageSignedFields struct {
	message membuffers.Message
}

var m_CommitPreparePPMessageSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeBytes,membuffers.TypeBytes,}
var m_CommitPreparePPMessageSignedFields_Unions = [][]membuffers.FieldType{}

func CommitPreparePPMessageSignedFieldsReader(buf []byte) *CommitPreparePPMessageSignedFields {
	x := &CommitPreparePPMessageSignedFields{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_CommitPreparePPMessageSignedFields_Scheme, m_CommitPreparePPMessageSignedFields_Unions)
	return x
}

func (x *CommitPreparePPMessageSignedFields) IsValid() bool {
	return x.message.IsValid()
}

func (x *CommitPreparePPMessageSignedFields) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *CommitPreparePPMessageSignedFields) MessageType() LeanHelixMessageType {
	return LeanHelixMessageType(x.message.GetUint16(0))
}

func (x *CommitPreparePPMessageSignedFields) RawMessageType() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *CommitPreparePPMessageSignedFields) MutateMessageType(v LeanHelixMessageType) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *CommitPreparePPMessageSignedFields) BlockHeight() uint64 {
	return x.message.GetUint64(1)
}

func (x *CommitPreparePPMessageSignedFields) RawBlockHeight() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *CommitPreparePPMessageSignedFields) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(1, v)
}

func (x *CommitPreparePPMessageSignedFields) View() uint32 {
	return x.message.GetUint32(2)
}

func (x *CommitPreparePPMessageSignedFields) RawView() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *CommitPreparePPMessageSignedFields) MutateView(v uint32) error {
	return x.message.SetUint32(2, v)
}

func (x *CommitPreparePPMessageSignedFields) BlockHash() []byte {
	return x.message.GetBytes(3)
}

func (x *CommitPreparePPMessageSignedFields) RawBlockHash() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *CommitPreparePPMessageSignedFields) MutateBlockHash(v []byte) error {
	return x.message.SetBytes(3, v)
}

func (x *CommitPreparePPMessageSignedFields) SenderPublicKey() []byte {
	return x.message.GetBytes(4)
}

func (x *CommitPreparePPMessageSignedFields) RawSenderPublicKey() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *CommitPreparePPMessageSignedFields) MutateSenderPublicKey(v []byte) error {
	return x.message.SetBytes(4, v)
}

// builder

type CommitPreparePPMessageSignedFieldsBuilder struct {
	builder membuffers.Builder
	MessageType LeanHelixMessageType
	BlockHeight uint64
	View uint32
	BlockHash []byte
	SenderPublicKey []byte
}

func (w *CommitPreparePPMessageSignedFieldsBuilder) Write(buf []byte) (err error) {
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
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteUint32(buf, w.View)
	w.builder.WriteBytes(buf, w.BlockHash)
	w.builder.WriteBytes(buf, w.SenderPublicKey)
	return nil
}

func (w *CommitPreparePPMessageSignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *CommitPreparePPMessageSignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *CommitPreparePPMessageSignedFieldsBuilder) Build() *CommitPreparePPMessageSignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return CommitPreparePPMessageSignedFieldsReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ViewChangeSignedFields

// reader

type ViewChangeSignedFields struct {
	message membuffers.Message
}

var m_ViewChangeSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeMessage,membuffers.TypeBytes,}
var m_ViewChangeSignedFields_Unions = [][]membuffers.FieldType{}

func ViewChangeSignedFieldsReader(buf []byte) *ViewChangeSignedFields {
	x := &ViewChangeSignedFields{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ViewChangeSignedFields_Scheme, m_ViewChangeSignedFields_Unions)
	return x
}

func (x *ViewChangeSignedFields) IsValid() bool {
	return x.message.IsValid()
}

func (x *ViewChangeSignedFields) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ViewChangeSignedFields) MessageType() LeanHelixMessageType {
	return LeanHelixMessageType(x.message.GetUint16(0))
}

func (x *ViewChangeSignedFields) RawMessageType() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ViewChangeSignedFields) MutateMessageType(v LeanHelixMessageType) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *ViewChangeSignedFields) BlockHeight() uint64 {
	return x.message.GetUint64(1)
}

func (x *ViewChangeSignedFields) RawBlockHeight() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ViewChangeSignedFields) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(1, v)
}

func (x *ViewChangeSignedFields) View() uint32 {
	return x.message.GetUint32(2)
}

func (x *ViewChangeSignedFields) RawView() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *ViewChangeSignedFields) MutateView(v uint32) error {
	return x.message.SetUint32(2, v)
}

func (x *ViewChangeSignedFields) PreparedProof() *PreparedProof {
	b, s := x.message.GetMessage(3)
	return PreparedProofReader(b[:s])
}

func (x *ViewChangeSignedFields) RawPreparedProof() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *ViewChangeSignedFields) SenderPublicKey() []byte {
	return x.message.GetBytes(4)
}

func (x *ViewChangeSignedFields) RawSenderPublicKey() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *ViewChangeSignedFields) MutateSenderPublicKey(v []byte) error {
	return x.message.SetBytes(4, v)
}

// builder

type ViewChangeSignedFieldsBuilder struct {
	builder membuffers.Builder
	MessageType LeanHelixMessageType
	BlockHeight uint64
	View uint32
	PreparedProof *PreparedProofBuilder
	SenderPublicKey []byte
}

func (w *ViewChangeSignedFieldsBuilder) Write(buf []byte) (err error) {
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
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteUint32(buf, w.View)
	err = w.builder.WriteMessage(buf, w.PreparedProof)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.SenderPublicKey)
	return nil
}

func (w *ViewChangeSignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ViewChangeSignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ViewChangeSignedFieldsBuilder) Build() *ViewChangeSignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ViewChangeSignedFieldsReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message PreparedProof

// reader

type PreparedProof struct {
	message membuffers.Message
}

var m_PreparedProof_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeBytesArray,}
var m_PreparedProof_Unions = [][]membuffers.FieldType{}

func PreparedProofReader(buf []byte) *PreparedProof {
	x := &PreparedProof{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_PreparedProof_Scheme, m_PreparedProof_Unions)
	return x
}

func (x *PreparedProof) IsValid() bool {
	return x.message.IsValid()
}

func (x *PreparedProof) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *PreparedProof) PpPrepareSignedFieldsIterator() *PreparedProofPpPrepareSignedFieldsIterator {
	return &PreparedProofPpPrepareSignedFieldsIterator{iterator: x.message.GetMessageArrayIterator(0)}
}

type PreparedProofPpPrepareSignedFieldsIterator struct {
	iterator *membuffers.Iterator
}

func (i *PreparedProofPpPrepareSignedFieldsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *PreparedProofPpPrepareSignedFieldsIterator) NextPpPrepareSignedFields() *CommitPreparePPMessageSignedFields {
	b, s := i.iterator.NextMessage()
	return CommitPreparePPMessageSignedFieldsReader(b[:s])
}

func (x *PreparedProof) RawPpPrepareSignedFieldsArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *PreparedProof) PpPrepareSignatureIterator() *PreparedProofPpPrepareSignatureIterator {
	return &PreparedProofPpPrepareSignatureIterator{iterator: x.message.GetBytesArrayIterator(1)}
}

type PreparedProofPpPrepareSignatureIterator struct {
	iterator *membuffers.Iterator
}

func (i *PreparedProofPpPrepareSignatureIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *PreparedProofPpPrepareSignatureIterator) NextPpPrepareSignature() []byte {
	return i.iterator.NextBytes()
}

func (x *PreparedProof) RawPpPrepareSignatureArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type PreparedProofBuilder struct {
	builder membuffers.Builder
	PpPrepareSignedFields []*CommitPreparePPMessageSignedFieldsBuilder
	PpPrepareSignature [][]byte
}

func (w *PreparedProofBuilder) arrayOfPpPrepareSignedFields() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.PpPrepareSignedFields))
	for i, v := range w.PpPrepareSignedFields {
		res[i] = v
	}
	return res
}

func (w *PreparedProofBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessageArray(buf, w.arrayOfPpPrepareSignedFields())
	if err != nil {
		return
	}
	w.builder.WriteBytesArray(buf, w.PpPrepareSignature)
	return nil
}

func (w *PreparedProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *PreparedProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *PreparedProofBuilder) Build() *PreparedProof {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return PreparedProofReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message NewViewSignedFields

// reader

type NewViewSignedFields struct {
	message membuffers.Message
}

var m_NewViewSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytes,}
var m_NewViewSignedFields_Unions = [][]membuffers.FieldType{}

func NewViewSignedFieldsReader(buf []byte) *NewViewSignedFields {
	x := &NewViewSignedFields{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_NewViewSignedFields_Scheme, m_NewViewSignedFields_Unions)
	return x
}

func (x *NewViewSignedFields) IsValid() bool {
	return x.message.IsValid()
}

func (x *NewViewSignedFields) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *NewViewSignedFields) MessageType() LeanHelixMessageType {
	return LeanHelixMessageType(x.message.GetUint16(0))
}

func (x *NewViewSignedFields) RawMessageType() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *NewViewSignedFields) MutateMessageType(v LeanHelixMessageType) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *NewViewSignedFields) BlockHeight() uint64 {
	return x.message.GetUint64(1)
}

func (x *NewViewSignedFields) RawBlockHeight() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *NewViewSignedFields) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(1, v)
}

func (x *NewViewSignedFields) View() uint32 {
	return x.message.GetUint32(2)
}

func (x *NewViewSignedFields) RawView() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *NewViewSignedFields) MutateView(v uint32) error {
	return x.message.SetUint32(2, v)
}

func (x *NewViewSignedFields) NewViewProof() *NewViewProof {
	b, s := x.message.GetMessage(3)
	return NewViewProofReader(b[:s])
}

func (x *NewViewSignedFields) RawNewViewProof() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *NewViewSignedFields) NvPpSignedFields() *CommitPreparePPMessageSignedFields {
	b, s := x.message.GetMessage(4)
	return CommitPreparePPMessageSignedFieldsReader(b[:s])
}

func (x *NewViewSignedFields) RawNvPpSignedFields() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *NewViewSignedFields) NvPpSignature() []byte {
	return x.message.GetBytes(5)
}

func (x *NewViewSignedFields) RawNvPpSignature() []byte {
	return x.message.RawBufferForField(5, 0)
}

func (x *NewViewSignedFields) MutateNvPpSignature(v []byte) error {
	return x.message.SetBytes(5, v)
}

func (x *NewViewSignedFields) SenderPublicKey() []byte {
	return x.message.GetBytes(6)
}

func (x *NewViewSignedFields) RawSenderPublicKey() []byte {
	return x.message.RawBufferForField(6, 0)
}

func (x *NewViewSignedFields) MutateSenderPublicKey(v []byte) error {
	return x.message.SetBytes(6, v)
}

// builder

type NewViewSignedFieldsBuilder struct {
	builder membuffers.Builder
	MessageType LeanHelixMessageType
	BlockHeight uint64
	View uint32
	NewViewProof *NewViewProofBuilder
	NvPpSignedFields *CommitPreparePPMessageSignedFieldsBuilder
	NvPpSignature []byte
	SenderPublicKey []byte
}

func (w *NewViewSignedFieldsBuilder) Write(buf []byte) (err error) {
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
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteUint32(buf, w.View)
	err = w.builder.WriteMessage(buf, w.NewViewProof)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.NvPpSignedFields)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.NvPpSignature)
	w.builder.WriteBytes(buf, w.SenderPublicKey)
	return nil
}

func (w *NewViewSignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *NewViewSignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *NewViewSignedFieldsBuilder) Build() *NewViewSignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return NewViewSignedFieldsReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message NewViewProof

// reader

type NewViewProof struct {
	message membuffers.Message
}

var m_NewViewProof_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeBytesArray,}
var m_NewViewProof_Unions = [][]membuffers.FieldType{}

func NewViewProofReader(buf []byte) *NewViewProof {
	x := &NewViewProof{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_NewViewProof_Scheme, m_NewViewProof_Unions)
	return x
}

func (x *NewViewProof) IsValid() bool {
	return x.message.IsValid()
}

func (x *NewViewProof) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *NewViewProof) ViewChangeHeaderIterator() *NewViewProofViewChangeHeaderIterator {
	return &NewViewProofViewChangeHeaderIterator{iterator: x.message.GetMessageArrayIterator(0)}
}

type NewViewProofViewChangeHeaderIterator struct {
	iterator *membuffers.Iterator
}

func (i *NewViewProofViewChangeHeaderIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *NewViewProofViewChangeHeaderIterator) NextViewChangeHeader() *ViewChangeSignedFields {
	b, s := i.iterator.NextMessage()
	return ViewChangeSignedFieldsReader(b[:s])
}

func (x *NewViewProof) RawViewChangeHeaderArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *NewViewProof) ViewChangeHeaderSignatureIterator() *NewViewProofViewChangeHeaderSignatureIterator {
	return &NewViewProofViewChangeHeaderSignatureIterator{iterator: x.message.GetBytesArrayIterator(1)}
}

type NewViewProofViewChangeHeaderSignatureIterator struct {
	iterator *membuffers.Iterator
}

func (i *NewViewProofViewChangeHeaderSignatureIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *NewViewProofViewChangeHeaderSignatureIterator) NextViewChangeHeaderSignature() []byte {
	return i.iterator.NextBytes()
}

func (x *NewViewProof) RawViewChangeHeaderSignatureArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type NewViewProofBuilder struct {
	builder membuffers.Builder
	ViewChangeHeader []*ViewChangeSignedFieldsBuilder
	ViewChangeHeaderSignature [][]byte
}

func (w *NewViewProofBuilder) arrayOfViewChangeHeader() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.ViewChangeHeader))
	for i, v := range w.ViewChangeHeader {
		res[i] = v
	}
	return res
}

func (w *NewViewProofBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessageArray(buf, w.arrayOfViewChangeHeader())
	if err != nil {
		return
	}
	w.builder.WriteBytesArray(buf, w.ViewChangeHeaderSignature)
	return nil
}

func (w *NewViewProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *NewViewProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *NewViewProofBuilder) Build() *NewViewProof {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return NewViewProofReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixBlockProof

// reader

type LeanHelixBlockProof struct {
	message membuffers.Message
}

var m_LeanHelixBlockProof_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytesArray,membuffers.TypeBytesArray,membuffers.TypeBytes,}
var m_LeanHelixBlockProof_Unions = [][]membuffers.FieldType{}

func LeanHelixBlockProofReader(buf []byte) *LeanHelixBlockProof {
	x := &LeanHelixBlockProof{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_LeanHelixBlockProof_Scheme, m_LeanHelixBlockProof_Unions)
	return x
}

func (x *LeanHelixBlockProof) IsValid() bool {
	return x.message.IsValid()
}

func (x *LeanHelixBlockProof) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *LeanHelixBlockProof) CommitMessageType() LeanHelixMessageType {
	return LeanHelixMessageType(x.message.GetUint16(0))
}

func (x *LeanHelixBlockProof) RawCommitMessageType() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *LeanHelixBlockProof) MutateCommitMessageType(v LeanHelixMessageType) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *LeanHelixBlockProof) BlockHeight() uint64 {
	return x.message.GetUint64(1)
}

func (x *LeanHelixBlockProof) RawBlockHeight() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *LeanHelixBlockProof) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(1, v)
}

func (x *LeanHelixBlockProof) View() uint32 {
	return x.message.GetUint32(2)
}

func (x *LeanHelixBlockProof) RawView() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *LeanHelixBlockProof) MutateView(v uint32) error {
	return x.message.SetUint32(2, v)
}

func (x *LeanHelixBlockProof) BlockHashMask() []byte {
	return x.message.GetBytes(3)
}

func (x *LeanHelixBlockProof) RawBlockHashMask() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *LeanHelixBlockProof) MutateBlockHashMask(v []byte) error {
	return x.message.SetBytes(3, v)
}

func (x *LeanHelixBlockProof) BlockHash() []byte {
	return x.message.GetBytes(4)
}

func (x *LeanHelixBlockProof) RawBlockHash() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *LeanHelixBlockProof) MutateBlockHash(v []byte) error {
	return x.message.SetBytes(4, v)
}

func (x *LeanHelixBlockProof) PkeyIterator() *LeanHelixBlockProofPkeyIterator {
	return &LeanHelixBlockProofPkeyIterator{iterator: x.message.GetBytesArrayIterator(5)}
}

type LeanHelixBlockProofPkeyIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixBlockProofPkeyIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixBlockProofPkeyIterator) NextPkey() []byte {
	return i.iterator.NextBytes()
}

func (x *LeanHelixBlockProof) RawPkeyArray() []byte {
	return x.message.RawBufferForField(5, 0)
}

func (x *LeanHelixBlockProof) SignatureIterator() *LeanHelixBlockProofSignatureIterator {
	return &LeanHelixBlockProofSignatureIterator{iterator: x.message.GetBytesArrayIterator(6)}
}

type LeanHelixBlockProofSignatureIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixBlockProofSignatureIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixBlockProofSignatureIterator) NextSignature() []byte {
	return i.iterator.NextBytes()
}

func (x *LeanHelixBlockProof) RawSignatureArray() []byte {
	return x.message.RawBufferForField(6, 0)
}

func (x *LeanHelixBlockProof) RandomSeedSignature() []byte {
	return x.message.GetBytes(7)
}

func (x *LeanHelixBlockProof) RawRandomSeedSignature() []byte {
	return x.message.RawBufferForField(7, 0)
}

func (x *LeanHelixBlockProof) MutateRandomSeedSignature(v []byte) error {
	return x.message.SetBytes(7, v)
}

// builder

type LeanHelixBlockProofBuilder struct {
	builder membuffers.Builder
	CommitMessageType LeanHelixMessageType
	BlockHeight uint64
	View uint32
	BlockHashMask []byte
	BlockHash []byte
	Pkey [][]byte
	Signature [][]byte
	RandomSeedSignature []byte
}

func (w *LeanHelixBlockProofBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.CommitMessageType))
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteUint32(buf, w.View)
	w.builder.WriteBytes(buf, w.BlockHashMask)
	w.builder.WriteBytes(buf, w.BlockHash)
	w.builder.WriteBytesArray(buf, w.Pkey)
	w.builder.WriteBytesArray(buf, w.Signature)
	w.builder.WriteBytes(buf, w.RandomSeedSignature)
	return nil
}

func (w *LeanHelixBlockProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *LeanHelixBlockProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *LeanHelixBlockProofBuilder) Build() *LeanHelixBlockProof {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixBlockProofReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type LeanHelixMessageType uint16

const (
	LEAN_HELIX_CONSENSUS_RESERVED LeanHelixMessageType = 0
	LEAN_HELIX_CONSENSUS_PRE_PREPARE LeanHelixMessageType = 1
	LEAN_HELIX_CONSENSUS_PREPARE LeanHelixMessageType = 2
	LEAN_HELIX_CONSENSUS_COMMIT LeanHelixMessageType = 3
	LEAN_HELIX_CONSENSUS_NEW_VIEW LeanHelixMessageType = 4
	LEAN_HELIX_CONSENSUS_VIEW_CHANGE LeanHelixMessageType = 5
)

