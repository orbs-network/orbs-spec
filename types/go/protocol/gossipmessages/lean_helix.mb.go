// AUTO GENERATED FILE (by membufc proto compiler v0.0.16)
package gossipmessages

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"bytes"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrePrepareMessage (non serializable)

type LeanHelixPrePrepareMessage struct {
	SignedHeader *LeanHelixBlockRef
	Sender *SenderSignature
	BlockPair *protocol.BlockPairContainer
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrepareMessage (non serializable)

type LeanHelixPrepareMessage struct {
	SignedHeader *LeanHelixBlockRef
	Sender *SenderSignature
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixCommitMessage (non serializable)

type LeanHelixCommitMessage struct {
	SignedHeader *LeanHelixBlockRef
	Sender *SenderSignature
	Share *LeanHelixRandomSeedShare
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixViewChangeMessage (non serializable)

type LeanHelixViewChangeMessage struct {
	SignedHeader *LeanHelixViewChangeHeader
	Sender *SenderSignature
	BlockPair *protocol.BlockPairContainer
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewViewMessage (non serializable)

type LeanHelixNewViewMessage struct {
	SignedHeader *LeanHelixNewViewHeader
	Sender *SenderSignature
	BlockPair *protocol.BlockPairContainer
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixBlockRef

// reader

type LeanHelixBlockRef struct {
	// MessageType LeanHelixMessageType
	// BlockHeight primitives.BlockHeight
	// View uint32
	// BlockHash primitives.Uint256

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixBlockRef) String() string {
	return fmt.Sprintf("{MessageType:%s,BlockHeight:%s,View:%s,BlockHash:%s,}", x.StringMessageType(), x.StringBlockHeight(), x.StringView(), x.StringBlockHash())
}

var _LeanHelixBlockRef_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeBytes,}
var _LeanHelixBlockRef_Unions = [][]membuffers.FieldType{}

func LeanHelixBlockRefReader(buf []byte) *LeanHelixBlockRef {
	x := &LeanHelixBlockRef{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixBlockRef_Scheme, _LeanHelixBlockRef_Unions)
	return x
}

func (x *LeanHelixBlockRef) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixBlockRef) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixBlockRef) Equal(y *LeanHelixBlockRef) bool {
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixBlockRef) MessageType() LeanHelixMessageType {
	return LeanHelixMessageType(x._message.GetUint16(0))
}

func (x *LeanHelixBlockRef) RawMessageType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixBlockRef) MutateMessageType(v LeanHelixMessageType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *LeanHelixBlockRef) StringMessageType() string {
	return x.MessageType().String()
}

func (x *LeanHelixBlockRef) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *LeanHelixBlockRef) RawBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixBlockRef) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *LeanHelixBlockRef) StringBlockHeight() string {
	return fmt.Sprintf("%x", x.BlockHeight())
}

func (x *LeanHelixBlockRef) View() uint32 {
	return x._message.GetUint32(2)
}

func (x *LeanHelixBlockRef) RawView() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixBlockRef) MutateView(v uint32) error {
	return x._message.SetUint32(2, v)
}

func (x *LeanHelixBlockRef) StringView() string {
	return fmt.Sprintf("%x", x.View())
}

func (x *LeanHelixBlockRef) BlockHash() primitives.Uint256 {
	return primitives.Uint256(x._message.GetBytes(3))
}

func (x *LeanHelixBlockRef) RawBlockHash() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *LeanHelixBlockRef) MutateBlockHash(v primitives.Uint256) error {
	return x._message.SetBytes(3, []byte(v))
}

func (x *LeanHelixBlockRef) StringBlockHash() string {
	return fmt.Sprintf("%x", x.BlockHash())
}

// builder

type LeanHelixBlockRefBuilder struct {
	MessageType LeanHelixMessageType
	BlockHeight primitives.BlockHeight
	View uint32
	BlockHash primitives.Uint256

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixBlockRefBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint16(buf, uint16(w.MessageType))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteUint32(buf, w.View)
	w._builder.WriteBytes(buf, []byte(w.BlockHash))
	return nil
}

func (w *LeanHelixBlockRefBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixBlockRefBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixBlockRefBuilder) Build() *LeanHelixBlockRef {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixBlockRefReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixRandomSeedShare

// reader

type LeanHelixRandomSeedShare struct {
	// RandomSeedShare primitives.Bls1Sig

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixRandomSeedShare) String() string {
	return fmt.Sprintf("{RandomSeedShare:%s,}", x.StringRandomSeedShare())
}

var _LeanHelixRandomSeedShare_Scheme = []membuffers.FieldType{membuffers.TypeBytes,}
var _LeanHelixRandomSeedShare_Unions = [][]membuffers.FieldType{}

func LeanHelixRandomSeedShareReader(buf []byte) *LeanHelixRandomSeedShare {
	x := &LeanHelixRandomSeedShare{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixRandomSeedShare_Scheme, _LeanHelixRandomSeedShare_Unions)
	return x
}

func (x *LeanHelixRandomSeedShare) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixRandomSeedShare) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixRandomSeedShare) Equal(y *LeanHelixRandomSeedShare) bool {
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixRandomSeedShare) RandomSeedShare() primitives.Bls1Sig {
	return primitives.Bls1Sig(x._message.GetBytes(0))
}

func (x *LeanHelixRandomSeedShare) RawRandomSeedShare() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixRandomSeedShare) MutateRandomSeedShare(v primitives.Bls1Sig) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *LeanHelixRandomSeedShare) StringRandomSeedShare() string {
	return fmt.Sprintf("%x", x.RandomSeedShare())
}

// builder

type LeanHelixRandomSeedShareBuilder struct {
	RandomSeedShare primitives.Bls1Sig

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixRandomSeedShareBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteBytes(buf, []byte(w.RandomSeedShare))
	return nil
}

func (w *LeanHelixRandomSeedShareBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixRandomSeedShareBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixRandomSeedShareBuilder) Build() *LeanHelixRandomSeedShare {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixRandomSeedShareReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixViewChangeHeader

// reader

type LeanHelixViewChangeHeader struct {
	// MessageType LeanHelixMessageType
	// BlockHeight primitives.BlockHeight
	// View uint32
	// PreparedProof LeanHelixPreparedProof

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixViewChangeHeader) String() string {
	return fmt.Sprintf("{MessageType:%s,BlockHeight:%s,View:%s,PreparedProof:%s,}", x.StringMessageType(), x.StringBlockHeight(), x.StringView(), x.StringPreparedProof())
}

var _LeanHelixViewChangeHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeMessage,}
var _LeanHelixViewChangeHeader_Unions = [][]membuffers.FieldType{}

func LeanHelixViewChangeHeaderReader(buf []byte) *LeanHelixViewChangeHeader {
	x := &LeanHelixViewChangeHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixViewChangeHeader_Scheme, _LeanHelixViewChangeHeader_Unions)
	return x
}

func (x *LeanHelixViewChangeHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixViewChangeHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixViewChangeHeader) Equal(y *LeanHelixViewChangeHeader) bool {
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixViewChangeHeader) MessageType() LeanHelixMessageType {
	return LeanHelixMessageType(x._message.GetUint16(0))
}

func (x *LeanHelixViewChangeHeader) RawMessageType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixViewChangeHeader) MutateMessageType(v LeanHelixMessageType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *LeanHelixViewChangeHeader) StringMessageType() string {
	return x.MessageType().String()
}

func (x *LeanHelixViewChangeHeader) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *LeanHelixViewChangeHeader) RawBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixViewChangeHeader) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *LeanHelixViewChangeHeader) StringBlockHeight() string {
	return fmt.Sprintf("%x", x.BlockHeight())
}

func (x *LeanHelixViewChangeHeader) View() uint32 {
	return x._message.GetUint32(2)
}

func (x *LeanHelixViewChangeHeader) RawView() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixViewChangeHeader) MutateView(v uint32) error {
	return x._message.SetUint32(2, v)
}

func (x *LeanHelixViewChangeHeader) StringView() string {
	return fmt.Sprintf("%x", x.View())
}

func (x *LeanHelixViewChangeHeader) PreparedProof() *LeanHelixPreparedProof {
	b, s := x._message.GetMessage(3)
	return LeanHelixPreparedProofReader(b[:s])
}

func (x *LeanHelixViewChangeHeader) RawPreparedProof() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *LeanHelixViewChangeHeader) StringPreparedProof() string {
	return x.PreparedProof().String()
}

// builder

type LeanHelixViewChangeHeaderBuilder struct {
	MessageType LeanHelixMessageType
	BlockHeight primitives.BlockHeight
	View uint32
	PreparedProof *LeanHelixPreparedProofBuilder

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixViewChangeHeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint16(buf, uint16(w.MessageType))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteUint32(buf, w.View)
	err = w._builder.WriteMessage(buf, w.PreparedProof)
	if err != nil {
		return
	}
	return nil
}

func (w *LeanHelixViewChangeHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixViewChangeHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixViewChangeHeaderBuilder) Build() *LeanHelixViewChangeHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixViewChangeHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPreparedProof

// reader

type LeanHelixPreparedProof struct {
	// BlockRef LeanHelixBlockRef
	// Senders []SenderSignature

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixPreparedProof) String() string {
	return fmt.Sprintf("{BlockRef:%s,Senders:%s,}", x.StringBlockRef(), x.StringSenders())
}

var _LeanHelixPreparedProof_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessageArray,}
var _LeanHelixPreparedProof_Unions = [][]membuffers.FieldType{}

func LeanHelixPreparedProofReader(buf []byte) *LeanHelixPreparedProof {
	x := &LeanHelixPreparedProof{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixPreparedProof_Scheme, _LeanHelixPreparedProof_Unions)
	return x
}

func (x *LeanHelixPreparedProof) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixPreparedProof) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixPreparedProof) Equal(y *LeanHelixPreparedProof) bool {
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixPreparedProof) BlockRef() *LeanHelixBlockRef {
	b, s := x._message.GetMessage(0)
	return LeanHelixBlockRefReader(b[:s])
}

func (x *LeanHelixPreparedProof) RawBlockRef() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixPreparedProof) StringBlockRef() string {
	return x.BlockRef().String()
}

func (x *LeanHelixPreparedProof) SendersIterator() *LeanHelixPreparedProofSendersIterator {
	return &LeanHelixPreparedProofSendersIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type LeanHelixPreparedProofSendersIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixPreparedProofSendersIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixPreparedProofSendersIterator) NextSenders() *SenderSignature {
	b, s := i.iterator.NextMessage()
	return SenderSignatureReader(b[:s])
}

func (x *LeanHelixPreparedProof) RawSendersArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixPreparedProof) StringSenders() (res string) {
	res = "["
	for i := x.SendersIterator(); i.HasNext(); {
		res += i.NextSenders().String() + ","
	}
	res += "]"
	return
}

// builder

type LeanHelixPreparedProofBuilder struct {
	BlockRef *LeanHelixBlockRefBuilder
	Senders []*SenderSignatureBuilder

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixPreparedProofBuilder) arrayOfSenders() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Senders))
	for i, v := range w.Senders {
		res[i] = v
	}
	return res
}

func (w *LeanHelixPreparedProofBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.BlockRef)
	if err != nil {
		return
	}
	err = w._builder.WriteMessageArray(buf, w.arrayOfSenders())
	if err != nil {
		return
	}
	return nil
}

func (w *LeanHelixPreparedProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixPreparedProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixPreparedProofBuilder) Build() *LeanHelixPreparedProof {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixPreparedProofReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewViewHeader

// reader

type LeanHelixNewViewHeader struct {
	// MessageType LeanHelixMessageType
	// BlockHeight primitives.BlockHeight
	// View uint32
	// NewViewProof LeanHelixNewViewProof
	// NewViewPrePrepareSignedHeader LeanHelixBlockRef
	// NewViewPrePrepareSender SenderSignature

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixNewViewHeader) String() string {
	return fmt.Sprintf("{MessageType:%s,BlockHeight:%s,View:%s,NewViewProof:%s,NewViewPrePrepareSignedHeader:%s,NewViewPrePrepareSender:%s,}", x.StringMessageType(), x.StringBlockHeight(), x.StringView(), x.StringNewViewProof(), x.StringNewViewPrePrepareSignedHeader(), x.StringNewViewPrePrepareSender())
}

var _LeanHelixNewViewHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}
var _LeanHelixNewViewHeader_Unions = [][]membuffers.FieldType{}

func LeanHelixNewViewHeaderReader(buf []byte) *LeanHelixNewViewHeader {
	x := &LeanHelixNewViewHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixNewViewHeader_Scheme, _LeanHelixNewViewHeader_Unions)
	return x
}

func (x *LeanHelixNewViewHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixNewViewHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixNewViewHeader) Equal(y *LeanHelixNewViewHeader) bool {
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixNewViewHeader) MessageType() LeanHelixMessageType {
	return LeanHelixMessageType(x._message.GetUint16(0))
}

func (x *LeanHelixNewViewHeader) RawMessageType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixNewViewHeader) MutateMessageType(v LeanHelixMessageType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *LeanHelixNewViewHeader) StringMessageType() string {
	return x.MessageType().String()
}

func (x *LeanHelixNewViewHeader) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *LeanHelixNewViewHeader) RawBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixNewViewHeader) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *LeanHelixNewViewHeader) StringBlockHeight() string {
	return fmt.Sprintf("%x", x.BlockHeight())
}

func (x *LeanHelixNewViewHeader) View() uint32 {
	return x._message.GetUint32(2)
}

func (x *LeanHelixNewViewHeader) RawView() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixNewViewHeader) MutateView(v uint32) error {
	return x._message.SetUint32(2, v)
}

func (x *LeanHelixNewViewHeader) StringView() string {
	return fmt.Sprintf("%x", x.View())
}

func (x *LeanHelixNewViewHeader) NewViewProof() *LeanHelixNewViewProof {
	b, s := x._message.GetMessage(3)
	return LeanHelixNewViewProofReader(b[:s])
}

func (x *LeanHelixNewViewHeader) RawNewViewProof() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *LeanHelixNewViewHeader) StringNewViewProof() string {
	return x.NewViewProof().String()
}

func (x *LeanHelixNewViewHeader) NewViewPrePrepareSignedHeader() *LeanHelixBlockRef {
	b, s := x._message.GetMessage(4)
	return LeanHelixBlockRefReader(b[:s])
}

func (x *LeanHelixNewViewHeader) RawNewViewPrePrepareSignedHeader() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *LeanHelixNewViewHeader) StringNewViewPrePrepareSignedHeader() string {
	return x.NewViewPrePrepareSignedHeader().String()
}

func (x *LeanHelixNewViewHeader) NewViewPrePrepareSender() *SenderSignature {
	b, s := x._message.GetMessage(5)
	return SenderSignatureReader(b[:s])
}

func (x *LeanHelixNewViewHeader) RawNewViewPrePrepareSender() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *LeanHelixNewViewHeader) StringNewViewPrePrepareSender() string {
	return x.NewViewPrePrepareSender().String()
}

// builder

type LeanHelixNewViewHeaderBuilder struct {
	MessageType LeanHelixMessageType
	BlockHeight primitives.BlockHeight
	View uint32
	NewViewProof *LeanHelixNewViewProofBuilder
	NewViewPrePrepareSignedHeader *LeanHelixBlockRefBuilder
	NewViewPrePrepareSender *SenderSignatureBuilder

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixNewViewHeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint16(buf, uint16(w.MessageType))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteUint32(buf, w.View)
	err = w._builder.WriteMessage(buf, w.NewViewProof)
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.NewViewPrePrepareSignedHeader)
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.NewViewPrePrepareSender)
	if err != nil {
		return
	}
	return nil
}

func (w *LeanHelixNewViewHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixNewViewHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixNewViewHeaderBuilder) Build() *LeanHelixNewViewHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixNewViewHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewViewProof

// reader

type LeanHelixNewViewProof struct {
	// ViewChangeSignedHeaders []LeanHelixViewChangeHeader
	// ViewChangeSenders []SenderSignature

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixNewViewProof) String() string {
	return fmt.Sprintf("{ViewChangeSignedHeaders:%s,ViewChangeSenders:%s,}", x.StringViewChangeSignedHeaders(), x.StringViewChangeSenders())
}

var _LeanHelixNewViewProof_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeMessageArray,}
var _LeanHelixNewViewProof_Unions = [][]membuffers.FieldType{}

func LeanHelixNewViewProofReader(buf []byte) *LeanHelixNewViewProof {
	x := &LeanHelixNewViewProof{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixNewViewProof_Scheme, _LeanHelixNewViewProof_Unions)
	return x
}

func (x *LeanHelixNewViewProof) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixNewViewProof) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixNewViewProof) Equal(y *LeanHelixNewViewProof) bool {
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixNewViewProof) ViewChangeSignedHeadersIterator() *LeanHelixNewViewProofViewChangeSignedHeadersIterator {
	return &LeanHelixNewViewProofViewChangeSignedHeadersIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type LeanHelixNewViewProofViewChangeSignedHeadersIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixNewViewProofViewChangeSignedHeadersIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixNewViewProofViewChangeSignedHeadersIterator) NextViewChangeSignedHeaders() *LeanHelixViewChangeHeader {
	b, s := i.iterator.NextMessage()
	return LeanHelixViewChangeHeaderReader(b[:s])
}

func (x *LeanHelixNewViewProof) RawViewChangeSignedHeadersArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixNewViewProof) StringViewChangeSignedHeaders() (res string) {
	res = "["
	for i := x.ViewChangeSignedHeadersIterator(); i.HasNext(); {
		res += i.NextViewChangeSignedHeaders().String() + ","
	}
	res += "]"
	return
}

func (x *LeanHelixNewViewProof) ViewChangeSendersIterator() *LeanHelixNewViewProofViewChangeSendersIterator {
	return &LeanHelixNewViewProofViewChangeSendersIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type LeanHelixNewViewProofViewChangeSendersIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixNewViewProofViewChangeSendersIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixNewViewProofViewChangeSendersIterator) NextViewChangeSenders() *SenderSignature {
	b, s := i.iterator.NextMessage()
	return SenderSignatureReader(b[:s])
}

func (x *LeanHelixNewViewProof) RawViewChangeSendersArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixNewViewProof) StringViewChangeSenders() (res string) {
	res = "["
	for i := x.ViewChangeSendersIterator(); i.HasNext(); {
		res += i.NextViewChangeSenders().String() + ","
	}
	res += "]"
	return
}

// builder

type LeanHelixNewViewProofBuilder struct {
	ViewChangeSignedHeaders []*LeanHelixViewChangeHeaderBuilder
	ViewChangeSenders []*SenderSignatureBuilder

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixNewViewProofBuilder) arrayOfViewChangeSignedHeaders() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.ViewChangeSignedHeaders))
	for i, v := range w.ViewChangeSignedHeaders {
		res[i] = v
	}
	return res
}

func (w *LeanHelixNewViewProofBuilder) arrayOfViewChangeSenders() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.ViewChangeSenders))
	for i, v := range w.ViewChangeSenders {
		res[i] = v
	}
	return res
}

func (w *LeanHelixNewViewProofBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessageArray(buf, w.arrayOfViewChangeSignedHeaders())
	if err != nil {
		return
	}
	err = w._builder.WriteMessageArray(buf, w.arrayOfViewChangeSenders())
	if err != nil {
		return
	}
	return nil
}

func (w *LeanHelixNewViewProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixNewViewProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixNewViewProofBuilder) Build() *LeanHelixNewViewProof {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixNewViewProofReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type LeanHelixMessageType uint16

const (
	LEAN_HELIX_RESERVED LeanHelixMessageType = 0
	LEAN_HELIX_PRE_PREPARE LeanHelixMessageType = 1
	LEAN_HELIX_PREPARE LeanHelixMessageType = 2
	LEAN_HELIX_COMMIT LeanHelixMessageType = 3
	LEAN_HELIX_NEW_VIEW LeanHelixMessageType = 4
	LEAN_HELIX_VIEW_CHANGE LeanHelixMessageType = 5
)

func (n LeanHelixMessageType) String() string {
	switch n {
	case LEAN_HELIX_RESERVED:
		return "LEAN_HELIX_RESERVED"
	case LEAN_HELIX_PRE_PREPARE:
		return "LEAN_HELIX_PRE_PREPARE"
	case LEAN_HELIX_PREPARE:
		return "LEAN_HELIX_PREPARE"
	case LEAN_HELIX_COMMIT:
		return "LEAN_HELIX_COMMIT"
	case LEAN_HELIX_NEW_VIEW:
		return "LEAN_HELIX_NEW_VIEW"
	case LEAN_HELIX_VIEW_CHANGE:
		return "LEAN_HELIX_VIEW_CHANGE"
	}
	return "UNKNOWN"
}

