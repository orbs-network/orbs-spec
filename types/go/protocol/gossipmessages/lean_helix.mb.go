// AUTO GENERATED FILE (by membufc proto compiler v0.0.14)
package gossipmessages

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrePrepareHeader

// reader

type LeanHelixPrePrepareHeader struct {
	// SenderPublicKey primitives.Ed25519Pkey
	// Signature primitives.Ed25519Sig

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixPrePrepareHeader_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var _LeanHelixPrePrepareHeader_Unions = [][]membuffers.FieldType{}

func LeanHelixPrePrepareHeaderReader(buf []byte) *LeanHelixPrePrepareHeader {
	x := &LeanHelixPrePrepareHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixPrePrepareHeader_Scheme, _LeanHelixPrePrepareHeader_Unions)
	return x
}

func (x *LeanHelixPrePrepareHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixPrePrepareHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixPrePrepareHeader) SenderPublicKey() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(x._message.GetBytes(0))
}

func (x *LeanHelixPrePrepareHeader) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixPrePrepareHeader) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *LeanHelixPrePrepareHeader) Signature() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(x._message.GetBytes(1))
}

func (x *LeanHelixPrePrepareHeader) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixPrePrepareHeader) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

// builder

type LeanHelixPrePrepareHeaderBuilder struct {
	SenderPublicKey primitives.Ed25519Pkey
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixPrePrepareHeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteBytes(buf, []byte(w.SenderPublicKey))
	w._builder.WriteBytes(buf, []byte(w.Signature))
	return nil
}

func (w *LeanHelixPrePrepareHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixPrePrepareHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixPrePrepareHeaderBuilder) Build() *LeanHelixPrePrepareHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixPrePrepareHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrepareHeader

// reader

type LeanHelixPrepareHeader struct {
	// SenderPublicKey primitives.Ed25519Pkey
	// Signature primitives.Ed25519Sig

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixPrepareHeader_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var _LeanHelixPrepareHeader_Unions = [][]membuffers.FieldType{}

func LeanHelixPrepareHeaderReader(buf []byte) *LeanHelixPrepareHeader {
	x := &LeanHelixPrepareHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixPrepareHeader_Scheme, _LeanHelixPrepareHeader_Unions)
	return x
}

func (x *LeanHelixPrepareHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixPrepareHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixPrepareHeader) SenderPublicKey() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(x._message.GetBytes(0))
}

func (x *LeanHelixPrepareHeader) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixPrepareHeader) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *LeanHelixPrepareHeader) Signature() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(x._message.GetBytes(1))
}

func (x *LeanHelixPrepareHeader) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixPrepareHeader) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

// builder

type LeanHelixPrepareHeaderBuilder struct {
	SenderPublicKey primitives.Ed25519Pkey
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixPrepareHeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteBytes(buf, []byte(w.SenderPublicKey))
	w._builder.WriteBytes(buf, []byte(w.Signature))
	return nil
}

func (w *LeanHelixPrepareHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixPrepareHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixPrepareHeaderBuilder) Build() *LeanHelixPrepareHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixPrepareHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixCommitHeader

// reader

type LeanHelixCommitHeader struct {
	// SenderPublicKey primitives.Ed25519Pkey
	// Signature primitives.Ed25519Sig
	// RandomSeedShare primitives.Bls1Sig

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixCommitHeader_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytes,}
var _LeanHelixCommitHeader_Unions = [][]membuffers.FieldType{}

func LeanHelixCommitHeaderReader(buf []byte) *LeanHelixCommitHeader {
	x := &LeanHelixCommitHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixCommitHeader_Scheme, _LeanHelixCommitHeader_Unions)
	return x
}

func (x *LeanHelixCommitHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixCommitHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixCommitHeader) SenderPublicKey() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(x._message.GetBytes(0))
}

func (x *LeanHelixCommitHeader) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixCommitHeader) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *LeanHelixCommitHeader) Signature() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(x._message.GetBytes(1))
}

func (x *LeanHelixCommitHeader) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixCommitHeader) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *LeanHelixCommitHeader) RandomSeedShare() primitives.Bls1Sig {
	return primitives.Bls1Sig(x._message.GetBytes(2))
}

func (x *LeanHelixCommitHeader) RawRandomSeedShare() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixCommitHeader) MutateRandomSeedShare(v primitives.Bls1Sig) error {
	return x._message.SetBytes(2, []byte(v))
}

// builder

type LeanHelixCommitHeaderBuilder struct {
	SenderPublicKey primitives.Ed25519Pkey
	Signature primitives.Ed25519Sig
	RandomSeedShare primitives.Bls1Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixCommitHeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteBytes(buf, []byte(w.SenderPublicKey))
	w._builder.WriteBytes(buf, []byte(w.Signature))
	w._builder.WriteBytes(buf, []byte(w.RandomSeedShare))
	return nil
}

func (w *LeanHelixCommitHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixCommitHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixCommitHeaderBuilder) Build() *LeanHelixCommitHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixCommitHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixViewChangeHeader

// reader

type LeanHelixViewChangeHeader struct {
	// SenderPublicKey primitives.Ed25519Pkey
	// Signature primitives.Ed25519Sig

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixViewChangeHeader_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
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

func (x *LeanHelixViewChangeHeader) SenderPublicKey() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(x._message.GetBytes(0))
}

func (x *LeanHelixViewChangeHeader) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixViewChangeHeader) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *LeanHelixViewChangeHeader) Signature() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(x._message.GetBytes(1))
}

func (x *LeanHelixViewChangeHeader) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixViewChangeHeader) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

// builder

type LeanHelixViewChangeHeaderBuilder struct {
	SenderPublicKey primitives.Ed25519Pkey
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
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
	w._builder.WriteBytes(buf, []byte(w.SenderPublicKey))
	w._builder.WriteBytes(buf, []byte(w.Signature))
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
// message LeanHelixNewViewHeader

// reader

type LeanHelixNewViewHeader struct {
	// SenderPublicKey primitives.Ed25519Pkey
	// Signature primitives.Ed25519Sig

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixNewViewHeader_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
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

func (x *LeanHelixNewViewHeader) SenderPublicKey() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(x._message.GetBytes(0))
}

func (x *LeanHelixNewViewHeader) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixNewViewHeader) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *LeanHelixNewViewHeader) Signature() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(x._message.GetBytes(1))
}

func (x *LeanHelixNewViewHeader) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixNewViewHeader) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

// builder

type LeanHelixNewViewHeaderBuilder struct {
	SenderPublicKey primitives.Ed25519Pkey
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
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
	w._builder.WriteBytes(buf, []byte(w.SenderPublicKey))
	w._builder.WriteBytes(buf, []byte(w.Signature))
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
// message LeanHelixBlockRef

// reader

type LeanHelixBlockRef struct {
	// BlockHeight primitives.BlockHeight
	// View uint32
	// BlockHash primitives.Sha256

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixBlockRef_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeBytes,}
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

func (x *LeanHelixBlockRef) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(0))
}

func (x *LeanHelixBlockRef) RawBlockHeight() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixBlockRef) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(0, uint64(v))
}

func (x *LeanHelixBlockRef) View() uint32 {
	return x._message.GetUint32(1)
}

func (x *LeanHelixBlockRef) RawView() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixBlockRef) MutateView(v uint32) error {
	return x._message.SetUint32(1, v)
}

func (x *LeanHelixBlockRef) BlockHash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(2))
}

func (x *LeanHelixBlockRef) RawBlockHash() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixBlockRef) MutateBlockHash(v primitives.Sha256) error {
	return x._message.SetBytes(2, []byte(v))
}

// builder

type LeanHelixBlockRefBuilder struct {
	BlockHeight primitives.BlockHeight
	View uint32
	BlockHash primitives.Sha256

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
// message LeanHelixViewChange

// reader

type LeanHelixViewChange struct {
	// BlockRefs []LeanHelixBlockRef
	// Signatures []primitives.Ed25519Sig

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixViewChange_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeBytesArray,}
var _LeanHelixViewChange_Unions = [][]membuffers.FieldType{}

func LeanHelixViewChangeReader(buf []byte) *LeanHelixViewChange {
	x := &LeanHelixViewChange{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixViewChange_Scheme, _LeanHelixViewChange_Unions)
	return x
}

func (x *LeanHelixViewChange) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixViewChange) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixViewChange) BlockRefsIterator() *LeanHelixViewChangeBlockRefsIterator {
	return &LeanHelixViewChangeBlockRefsIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type LeanHelixViewChangeBlockRefsIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixViewChangeBlockRefsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixViewChangeBlockRefsIterator) NextBlockRefs() *LeanHelixBlockRef {
	b, s := i.iterator.NextMessage()
	return LeanHelixBlockRefReader(b[:s])
}

func (x *LeanHelixViewChange) RawBlockRefsArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixViewChange) SignaturesIterator() *LeanHelixViewChangeSignaturesIterator {
	return &LeanHelixViewChangeSignaturesIterator{iterator: x._message.GetBytesArrayIterator(1)}
}

type LeanHelixViewChangeSignaturesIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixViewChangeSignaturesIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixViewChangeSignaturesIterator) NextSignatures() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(i.iterator.NextBytes())
}

func (x *LeanHelixViewChange) RawSignaturesArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

// builder

type LeanHelixViewChangeBuilder struct {
	BlockRefs []*LeanHelixBlockRefBuilder
	Signatures []primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixViewChangeBuilder) arrayOfBlockRefs() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.BlockRefs))
	for i, v := range w.BlockRefs {
		res[i] = v
	}
	return res
}

func (w *LeanHelixViewChangeBuilder) arrayOfSignatures() [][]byte {
	res := make([][]byte, len(w.Signatures))
	for i, v := range w.Signatures {
		res[i] = v
	}
	return res
}

func (w *LeanHelixViewChangeBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessageArray(buf, w.arrayOfBlockRefs())
	if err != nil {
		return
	}
	w._builder.WriteBytesArray(buf, w.arrayOfSignatures())
	return nil
}

func (w *LeanHelixViewChangeBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixViewChangeBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixViewChangeBuilder) Build() *LeanHelixViewChange {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixViewChangeReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewView

// reader

type LeanHelixNewView struct {
	// Signatures []primitives.Ed25519Sig
	// ViewChanges []LeanHelixViewChange

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixNewView_Scheme = []membuffers.FieldType{membuffers.TypeBytesArray,membuffers.TypeMessageArray,}
var _LeanHelixNewView_Unions = [][]membuffers.FieldType{}

func LeanHelixNewViewReader(buf []byte) *LeanHelixNewView {
	x := &LeanHelixNewView{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixNewView_Scheme, _LeanHelixNewView_Unions)
	return x
}

func (x *LeanHelixNewView) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixNewView) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixNewView) SignaturesIterator() *LeanHelixNewViewSignaturesIterator {
	return &LeanHelixNewViewSignaturesIterator{iterator: x._message.GetBytesArrayIterator(0)}
}

type LeanHelixNewViewSignaturesIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixNewViewSignaturesIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixNewViewSignaturesIterator) NextSignatures() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(i.iterator.NextBytes())
}

func (x *LeanHelixNewView) RawSignaturesArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixNewView) ViewChangesIterator() *LeanHelixNewViewViewChangesIterator {
	return &LeanHelixNewViewViewChangesIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type LeanHelixNewViewViewChangesIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixNewViewViewChangesIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixNewViewViewChangesIterator) NextViewChanges() *LeanHelixViewChange {
	b, s := i.iterator.NextMessage()
	return LeanHelixViewChangeReader(b[:s])
}

func (x *LeanHelixNewView) RawViewChangesArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

// builder

type LeanHelixNewViewBuilder struct {
	Signatures []primitives.Ed25519Sig
	ViewChanges []*LeanHelixViewChangeBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixNewViewBuilder) arrayOfViewChanges() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.ViewChanges))
	for i, v := range w.ViewChanges {
		res[i] = v
	}
	return res
}

func (w *LeanHelixNewViewBuilder) arrayOfSignatures() [][]byte {
	res := make([][]byte, len(w.Signatures))
	for i, v := range w.Signatures {
		res[i] = v
	}
	return res
}

func (w *LeanHelixNewViewBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteBytesArray(buf, w.arrayOfSignatures())
	err = w._builder.WriteMessageArray(buf, w.arrayOfViewChanges())
	if err != nil {
		return
	}
	return nil
}

func (w *LeanHelixNewViewBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixNewViewBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixNewViewBuilder) Build() *LeanHelixNewView {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixNewViewReader(buf)
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

