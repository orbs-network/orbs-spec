// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package gossipmessages

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrePrepareMessage (non serializable)

type LeanHelixPrePrepareMessage struct {
	PrePrepareSignedHeader *LeanHelixBlockRef
	Signer *MessageSigner
	BlockPair *protocol.BlockPair
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrepareMessage (non serializable)

type LeanHelixPrepareMessage struct {
	PrepareSignedHeader *LeanHelixBlockRef
	Signer *MessageSigner
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixCommitMessage (non serializable)

type LeanHelixCommitMessage struct {
	CommitSignedHeader *LeanHelixBlockRef
	Signer *MessageSigner
	LeanHelixRandomSeedShare *LeanHelixRandomSeedShare
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixViewChangeMessage (non serializable)

type LeanHelixViewChangeMessage struct {
	ViewChangeSignedHeader *LeanHelixViewChangeSignedHeader
	Signer *MessageSigner
	BlockPair *protocol.BlockPair
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewView (non serializable)

type LeanHelixNewView struct {
	NewViewSignedHeader *LeanHelixNewViewSignedHeader
	Signer *MessageSigner
	BlockPair *protocol.BlockPair
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
	membuffers.Message // interface
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
	membuffers.Builder // interface
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
	membuffers.Message // interface
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
	membuffers.Builder // interface
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
// message LeanHelixViewChangeSignedHeader

// reader

type LeanHelixViewChangeSignedHeader struct {
	// MessageType LeanHelixMessageType
	// BlockHeight primitives.BlockHeight
	// View uint32
	// PreparedProof LeanHelixPreparedProof

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *LeanHelixViewChangeSignedHeader) String() string {
	return fmt.Sprintf("{MessageType:%s,BlockHeight:%s,View:%s,PreparedProof:%s,}", x.StringMessageType(), x.StringBlockHeight(), x.StringView(), x.StringPreparedProof())
}

var _LeanHelixViewChangeSignedHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeMessage,}
var _LeanHelixViewChangeSignedHeader_Unions = [][]membuffers.FieldType{}

func LeanHelixViewChangeSignedHeaderReader(buf []byte) *LeanHelixViewChangeSignedHeader {
	x := &LeanHelixViewChangeSignedHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixViewChangeSignedHeader_Scheme, _LeanHelixViewChangeSignedHeader_Unions)
	return x
}

func (x *LeanHelixViewChangeSignedHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixViewChangeSignedHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixViewChangeSignedHeader) MessageType() LeanHelixMessageType {
	return LeanHelixMessageType(x._message.GetUint16(0))
}

func (x *LeanHelixViewChangeSignedHeader) RawMessageType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixViewChangeSignedHeader) MutateMessageType(v LeanHelixMessageType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *LeanHelixViewChangeSignedHeader) StringMessageType() string {
	return x.MessageType().String()
}

func (x *LeanHelixViewChangeSignedHeader) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *LeanHelixViewChangeSignedHeader) RawBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixViewChangeSignedHeader) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *LeanHelixViewChangeSignedHeader) StringBlockHeight() string {
	return fmt.Sprintf("%x", x.BlockHeight())
}

func (x *LeanHelixViewChangeSignedHeader) View() uint32 {
	return x._message.GetUint32(2)
}

func (x *LeanHelixViewChangeSignedHeader) RawView() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixViewChangeSignedHeader) MutateView(v uint32) error {
	return x._message.SetUint32(2, v)
}

func (x *LeanHelixViewChangeSignedHeader) StringView() string {
	return fmt.Sprintf("%x", x.View())
}

func (x *LeanHelixViewChangeSignedHeader) PreparedProof() *LeanHelixPreparedProof {
	b, s := x._message.GetMessage(3)
	return LeanHelixPreparedProofReader(b[:s])
}

func (x *LeanHelixViewChangeSignedHeader) RawPreparedProof() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *LeanHelixViewChangeSignedHeader) StringPreparedProof() string {
	return x.PreparedProof().String()
}

// builder

type LeanHelixViewChangeSignedHeaderBuilder struct {
	MessageType LeanHelixMessageType
	BlockHeight primitives.BlockHeight
	View uint32
	PreparedProof *LeanHelixPreparedProofBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixViewChangeSignedHeaderBuilder) Write(buf []byte) (err error) {
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

func (w *LeanHelixViewChangeSignedHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixViewChangeSignedHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixViewChangeSignedHeaderBuilder) Build() *LeanHelixViewChangeSignedHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixViewChangeSignedHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPreparedProof

// reader

type LeanHelixPreparedProof struct {
	// BlockRefs LeanHelixBlockRef
	// Signers []MessageSigner

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *LeanHelixPreparedProof) String() string {
	return fmt.Sprintf("{BlockRefs:%s,Signers:%s,}", x.StringBlockRefs(), x.StringSigners())
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

func (x *LeanHelixPreparedProof) BlockRefs() *LeanHelixBlockRef {
	b, s := x._message.GetMessage(0)
	return LeanHelixBlockRefReader(b[:s])
}

func (x *LeanHelixPreparedProof) RawBlockRefs() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixPreparedProof) StringBlockRefs() string {
	return x.BlockRefs().String()
}

func (x *LeanHelixPreparedProof) SignersIterator() *LeanHelixPreparedProofSignersIterator {
	return &LeanHelixPreparedProofSignersIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type LeanHelixPreparedProofSignersIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixPreparedProofSignersIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixPreparedProofSignersIterator) NextSigners() *MessageSigner {
	b, s := i.iterator.NextMessage()
	return MessageSignerReader(b[:s])
}

func (x *LeanHelixPreparedProof) RawSignersArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixPreparedProof) StringSigners() (res string) {
	res = "["
	for i := x.SignersIterator(); i.HasNext(); {
		res += i.NextSigners().String() + ","
	}
	res += "]"
	return
}

// builder

type LeanHelixPreparedProofBuilder struct {
	BlockRefs *LeanHelixBlockRefBuilder
	Signers []*MessageSignerBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixPreparedProofBuilder) arrayOfSigners() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Signers))
	for i, v := range w.Signers {
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
	err = w._builder.WriteMessage(buf, w.BlockRefs)
	if err != nil {
		return
	}
	err = w._builder.WriteMessageArray(buf, w.arrayOfSigners())
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
// message LeanHelixNewViewSignedHeader

// reader

type LeanHelixNewViewSignedHeader struct {
	// MessageType LeanHelixMessageType
	// BlockHeight primitives.BlockHeight
	// View uint32
	// NewViewProof LeanHelixNewViewProof
	// NvPpSignedHeader LeanHelixBlockRef
	// NvPpSigner MessageSigner

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *LeanHelixNewViewSignedHeader) String() string {
	return fmt.Sprintf("{MessageType:%s,BlockHeight:%s,View:%s,NewViewProof:%s,NvPpSignedHeader:%s,NvPpSigner:%s,}", x.StringMessageType(), x.StringBlockHeight(), x.StringView(), x.StringNewViewProof(), x.StringNvPpSignedHeader(), x.StringNvPpSigner())
}

var _LeanHelixNewViewSignedHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}
var _LeanHelixNewViewSignedHeader_Unions = [][]membuffers.FieldType{}

func LeanHelixNewViewSignedHeaderReader(buf []byte) *LeanHelixNewViewSignedHeader {
	x := &LeanHelixNewViewSignedHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixNewViewSignedHeader_Scheme, _LeanHelixNewViewSignedHeader_Unions)
	return x
}

func (x *LeanHelixNewViewSignedHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixNewViewSignedHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixNewViewSignedHeader) MessageType() LeanHelixMessageType {
	return LeanHelixMessageType(x._message.GetUint16(0))
}

func (x *LeanHelixNewViewSignedHeader) RawMessageType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixNewViewSignedHeader) MutateMessageType(v LeanHelixMessageType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *LeanHelixNewViewSignedHeader) StringMessageType() string {
	return x.MessageType().String()
}

func (x *LeanHelixNewViewSignedHeader) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *LeanHelixNewViewSignedHeader) RawBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixNewViewSignedHeader) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *LeanHelixNewViewSignedHeader) StringBlockHeight() string {
	return fmt.Sprintf("%x", x.BlockHeight())
}

func (x *LeanHelixNewViewSignedHeader) View() uint32 {
	return x._message.GetUint32(2)
}

func (x *LeanHelixNewViewSignedHeader) RawView() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixNewViewSignedHeader) MutateView(v uint32) error {
	return x._message.SetUint32(2, v)
}

func (x *LeanHelixNewViewSignedHeader) StringView() string {
	return fmt.Sprintf("%x", x.View())
}

func (x *LeanHelixNewViewSignedHeader) NewViewProof() *LeanHelixNewViewProof {
	b, s := x._message.GetMessage(3)
	return LeanHelixNewViewProofReader(b[:s])
}

func (x *LeanHelixNewViewSignedHeader) RawNewViewProof() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *LeanHelixNewViewSignedHeader) StringNewViewProof() string {
	return x.NewViewProof().String()
}

func (x *LeanHelixNewViewSignedHeader) NvPpSignedHeader() *LeanHelixBlockRef {
	b, s := x._message.GetMessage(4)
	return LeanHelixBlockRefReader(b[:s])
}

func (x *LeanHelixNewViewSignedHeader) RawNvPpSignedHeader() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *LeanHelixNewViewSignedHeader) StringNvPpSignedHeader() string {
	return x.NvPpSignedHeader().String()
}

func (x *LeanHelixNewViewSignedHeader) NvPpSigner() *MessageSigner {
	b, s := x._message.GetMessage(5)
	return MessageSignerReader(b[:s])
}

func (x *LeanHelixNewViewSignedHeader) RawNvPpSigner() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *LeanHelixNewViewSignedHeader) StringNvPpSigner() string {
	return x.NvPpSigner().String()
}

// builder

type LeanHelixNewViewSignedHeaderBuilder struct {
	MessageType LeanHelixMessageType
	BlockHeight primitives.BlockHeight
	View uint32
	NewViewProof *LeanHelixNewViewProofBuilder
	NvPpSignedHeader *LeanHelixBlockRefBuilder
	NvPpSigner *MessageSignerBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixNewViewSignedHeaderBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessage(buf, w.NvPpSignedHeader)
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.NvPpSigner)
	if err != nil {
		return
	}
	return nil
}

func (w *LeanHelixNewViewSignedHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixNewViewSignedHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixNewViewSignedHeaderBuilder) Build() *LeanHelixNewViewSignedHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixNewViewSignedHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewViewProof

// reader

type LeanHelixNewViewProof struct {
	// ViewChangeSignedHeader []LeanHelixViewChangeSignedHeader
	// ViewChangeSigners []MessageSigner

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *LeanHelixNewViewProof) String() string {
	return fmt.Sprintf("{ViewChangeSignedHeader:%s,ViewChangeSigners:%s,}", x.StringViewChangeSignedHeader(), x.StringViewChangeSigners())
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

func (x *LeanHelixNewViewProof) ViewChangeSignedHeaderIterator() *LeanHelixNewViewProofViewChangeSignedHeaderIterator {
	return &LeanHelixNewViewProofViewChangeSignedHeaderIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type LeanHelixNewViewProofViewChangeSignedHeaderIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixNewViewProofViewChangeSignedHeaderIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixNewViewProofViewChangeSignedHeaderIterator) NextViewChangeSignedHeader() *LeanHelixViewChangeSignedHeader {
	b, s := i.iterator.NextMessage()
	return LeanHelixViewChangeSignedHeaderReader(b[:s])
}

func (x *LeanHelixNewViewProof) RawViewChangeSignedHeaderArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixNewViewProof) StringViewChangeSignedHeader() (res string) {
	res = "["
	for i := x.ViewChangeSignedHeaderIterator(); i.HasNext(); {
		res += i.NextViewChangeSignedHeader().String() + ","
	}
	res += "]"
	return
}

func (x *LeanHelixNewViewProof) ViewChangeSignersIterator() *LeanHelixNewViewProofViewChangeSignersIterator {
	return &LeanHelixNewViewProofViewChangeSignersIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type LeanHelixNewViewProofViewChangeSignersIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixNewViewProofViewChangeSignersIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixNewViewProofViewChangeSignersIterator) NextViewChangeSigners() *MessageSigner {
	b, s := i.iterator.NextMessage()
	return MessageSignerReader(b[:s])
}

func (x *LeanHelixNewViewProof) RawViewChangeSignersArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixNewViewProof) StringViewChangeSigners() (res string) {
	res = "["
	for i := x.ViewChangeSignersIterator(); i.HasNext(); {
		res += i.NextViewChangeSigners().String() + ","
	}
	res += "]"
	return
}

// builder

type LeanHelixNewViewProofBuilder struct {
	ViewChangeSignedHeader []*LeanHelixViewChangeSignedHeaderBuilder
	ViewChangeSigners []*MessageSignerBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixNewViewProofBuilder) arrayOfViewChangeSignedHeader() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.ViewChangeSignedHeader))
	for i, v := range w.ViewChangeSignedHeader {
		res[i] = v
	}
	return res
}

func (w *LeanHelixNewViewProofBuilder) arrayOfViewChangeSigners() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.ViewChangeSigners))
	for i, v := range w.ViewChangeSigners {
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
	err = w._builder.WriteMessageArray(buf, w.arrayOfViewChangeSignedHeader())
	if err != nil {
		return
	}
	err = w._builder.WriteMessageArray(buf, w.arrayOfViewChangeSigners())
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

