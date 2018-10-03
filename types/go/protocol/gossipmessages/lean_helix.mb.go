// AUTO GENERATED FILE (by membufc proto compiler v0.0.19)
package gossipmessages

import (
	"github.com/orbs-network/membuffers/go"
	"bytes"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/consensus"
)

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrePrepareMessage (non serializable)

type LeanHelixPrePrepareMessage struct {
	SignedHeader *consensus.LeanHelixBlockRef
	Sender *consensus.LeanHelixSenderSignature
	BlockPair *protocol.BlockPairContainer
}

func (x *LeanHelixPrePrepareMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedHeader:%s,Sender:%s,BlockPair:%s,}", x.StringSignedHeader(), x.StringSender(), x.StringBlockPair())
}

func (x *LeanHelixPrePrepareMessage) StringSignedHeader() (res string) {
	res = x.SignedHeader.String()
	return
}

func (x *LeanHelixPrePrepareMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

func (x *LeanHelixPrePrepareMessage) StringBlockPair() (res string) {
	res = x.BlockPair.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrepareMessage (non serializable)

type LeanHelixPrepareMessage struct {
	SignedHeader *consensus.LeanHelixBlockRef
	Sender *consensus.LeanHelixSenderSignature
}

func (x *LeanHelixPrepareMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedHeader:%s,Sender:%s,}", x.StringSignedHeader(), x.StringSender())
}

func (x *LeanHelixPrepareMessage) StringSignedHeader() (res string) {
	res = x.SignedHeader.String()
	return
}

func (x *LeanHelixPrepareMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixCommitMessage (non serializable)

type LeanHelixCommitMessage struct {
	SignedHeader *consensus.LeanHelixBlockRef
	Sender *consensus.LeanHelixSenderSignature
	Share *LeanHelixRandomSeedShare
}

func (x *LeanHelixCommitMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedHeader:%s,Sender:%s,Share:%s,}", x.StringSignedHeader(), x.StringSender(), x.StringShare())
}

func (x *LeanHelixCommitMessage) StringSignedHeader() (res string) {
	res = x.SignedHeader.String()
	return
}

func (x *LeanHelixCommitMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

func (x *LeanHelixCommitMessage) StringShare() (res string) {
	res = x.Share.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixViewChangeMessage (non serializable)

type LeanHelixViewChangeMessage struct {
	SignedHeader *LeanHelixViewChangeHeader
	Sender *consensus.LeanHelixSenderSignature
	BlockPair *protocol.BlockPairContainer
}

func (x *LeanHelixViewChangeMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedHeader:%s,Sender:%s,BlockPair:%s,}", x.StringSignedHeader(), x.StringSender(), x.StringBlockPair())
}

func (x *LeanHelixViewChangeMessage) StringSignedHeader() (res string) {
	res = x.SignedHeader.String()
	return
}

func (x *LeanHelixViewChangeMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

func (x *LeanHelixViewChangeMessage) StringBlockPair() (res string) {
	res = x.BlockPair.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewViewMessage (non serializable)

type LeanHelixNewViewMessage struct {
	SignedHeader *LeanHelixNewViewHeader
	Sender *consensus.LeanHelixSenderSignature
	PrepareMessage *LeanHelixPrePrepareMessage
}

func (x *LeanHelixNewViewMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedHeader:%s,Sender:%s,PrepareMessage:%s,}", x.StringSignedHeader(), x.StringSender(), x.StringPrepareMessage())
}

func (x *LeanHelixNewViewMessage) StringSignedHeader() (res string) {
	res = x.SignedHeader.String()
	return
}

func (x *LeanHelixNewViewMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

func (x *LeanHelixNewViewMessage) StringPrepareMessage() (res string) {
	res = x.PrepareMessage.String()
	return
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
	if x == nil {
		return "<nil>"
	}
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
  if x == nil && y == nil {
    return true
  }
  if x == nil || y == nil {
    return false
  }
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
	return fmt.Sprintf("%s", x.RandomSeedShare())
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
	// MessageType consensus.LeanHelixMessageType
	// BlockHeight primitives.BlockHeight
	// View uint32
	// PreparedProof LeanHelixPreparedProof

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixViewChangeHeader) String() string {
	if x == nil {
		return "<nil>"
	}
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
  if x == nil && y == nil {
    return true
  }
  if x == nil || y == nil {
    return false
  }
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixViewChangeHeader) MessageType() consensus.LeanHelixMessageType {
	return consensus.LeanHelixMessageType(x._message.GetUint16(0))
}

func (x *LeanHelixViewChangeHeader) RawMessageType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixViewChangeHeader) MutateMessageType(v consensus.LeanHelixMessageType) error {
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
	return fmt.Sprintf("%s", x.BlockHeight())
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

func (x *LeanHelixViewChangeHeader) RawPreparedProofWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *LeanHelixViewChangeHeader) StringPreparedProof() string {
	return x.PreparedProof().String()
}

// builder

type LeanHelixViewChangeHeaderBuilder struct {
	MessageType consensus.LeanHelixMessageType
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
	// BlockRef consensus.LeanHelixBlockRef
	// Senders []consensus.LeanHelixSenderSignature

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixPreparedProof) String() string {
	if x == nil {
		return "<nil>"
	}
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
  if x == nil && y == nil {
    return true
  }
  if x == nil || y == nil {
    return false
  }
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixPreparedProof) BlockRef() *consensus.LeanHelixBlockRef {
	b, s := x._message.GetMessage(0)
	return consensus.LeanHelixBlockRefReader(b[:s])
}

func (x *LeanHelixPreparedProof) RawBlockRef() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixPreparedProof) RawBlockRefWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
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

func (i *LeanHelixPreparedProofSendersIterator) NextSenders() *consensus.LeanHelixSenderSignature {
	b, s := i.iterator.NextMessage()
	return consensus.LeanHelixSenderSignatureReader(b[:s])
}

func (x *LeanHelixPreparedProof) RawSendersArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixPreparedProof) RawSendersArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
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
	BlockRef *consensus.LeanHelixBlockRefBuilder
	Senders []*consensus.LeanHelixSenderSignatureBuilder

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
	// MessageType consensus.LeanHelixMessageType
	// BlockHeight primitives.BlockHeight
	// View uint32
	// ViewChangeSignedHeaders []LeanHelixViewChangeHeader
	// ViewChangeSenders []consensus.LeanHelixSenderSignature

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixNewViewHeader) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{MessageType:%s,BlockHeight:%s,View:%s,ViewChangeSignedHeaders:%s,ViewChangeSenders:%s,}", x.StringMessageType(), x.StringBlockHeight(), x.StringView(), x.StringViewChangeSignedHeaders(), x.StringViewChangeSenders())
}

var _LeanHelixNewViewHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeMessageArray,membuffers.TypeMessageArray,}
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
  if x == nil && y == nil {
    return true
  }
  if x == nil || y == nil {
    return false
  }
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixNewViewHeader) MessageType() consensus.LeanHelixMessageType {
	return consensus.LeanHelixMessageType(x._message.GetUint16(0))
}

func (x *LeanHelixNewViewHeader) RawMessageType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixNewViewHeader) MutateMessageType(v consensus.LeanHelixMessageType) error {
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
	return fmt.Sprintf("%s", x.BlockHeight())
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

func (x *LeanHelixNewViewHeader) ViewChangeSignedHeadersIterator() *LeanHelixNewViewHeaderViewChangeSignedHeadersIterator {
	return &LeanHelixNewViewHeaderViewChangeSignedHeadersIterator{iterator: x._message.GetMessageArrayIterator(3)}
}

type LeanHelixNewViewHeaderViewChangeSignedHeadersIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixNewViewHeaderViewChangeSignedHeadersIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixNewViewHeaderViewChangeSignedHeadersIterator) NextViewChangeSignedHeaders() *LeanHelixViewChangeHeader {
	b, s := i.iterator.NextMessage()
	return LeanHelixViewChangeHeaderReader(b[:s])
}

func (x *LeanHelixNewViewHeader) RawViewChangeSignedHeadersArray() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *LeanHelixNewViewHeader) RawViewChangeSignedHeadersArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *LeanHelixNewViewHeader) StringViewChangeSignedHeaders() (res string) {
	res = "["
	for i := x.ViewChangeSignedHeadersIterator(); i.HasNext(); {
		res += i.NextViewChangeSignedHeaders().String() + ","
	}
	res += "]"
	return
}

func (x *LeanHelixNewViewHeader) ViewChangeSendersIterator() *LeanHelixNewViewHeaderViewChangeSendersIterator {
	return &LeanHelixNewViewHeaderViewChangeSendersIterator{iterator: x._message.GetMessageArrayIterator(4)}
}

type LeanHelixNewViewHeaderViewChangeSendersIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixNewViewHeaderViewChangeSendersIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixNewViewHeaderViewChangeSendersIterator) NextViewChangeSenders() *consensus.LeanHelixSenderSignature {
	b, s := i.iterator.NextMessage()
	return consensus.LeanHelixSenderSignatureReader(b[:s])
}

func (x *LeanHelixNewViewHeader) RawViewChangeSendersArray() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *LeanHelixNewViewHeader) RawViewChangeSendersArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(4, 0)
}

func (x *LeanHelixNewViewHeader) StringViewChangeSenders() (res string) {
	res = "["
	for i := x.ViewChangeSendersIterator(); i.HasNext(); {
		res += i.NextViewChangeSenders().String() + ","
	}
	res += "]"
	return
}

// builder

type LeanHelixNewViewHeaderBuilder struct {
	MessageType consensus.LeanHelixMessageType
	BlockHeight primitives.BlockHeight
	View uint32
	ViewChangeSignedHeaders []*LeanHelixViewChangeHeaderBuilder
	ViewChangeSenders []*consensus.LeanHelixSenderSignatureBuilder

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixNewViewHeaderBuilder) arrayOfViewChangeSignedHeaders() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.ViewChangeSignedHeaders))
	for i, v := range w.ViewChangeSignedHeaders {
		res[i] = v
	}
	return res
}

func (w *LeanHelixNewViewHeaderBuilder) arrayOfViewChangeSenders() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.ViewChangeSenders))
	for i, v := range w.ViewChangeSenders {
		res[i] = v
	}
	return res
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
// enums

