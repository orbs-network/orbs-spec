// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package gossip

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrePrepareMessage

// reader

type LeanHelixPrePrepareMessage struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixPrePrepareMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytes,}
var _LeanHelixPrePrepareMessage_Unions = [][]membuffers.FieldType{}

func LeanHelixPrePrepareMessageReader(buf []byte) *LeanHelixPrePrepareMessage {
	x := &LeanHelixPrePrepareMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixPrePrepareMessage_Scheme, _LeanHelixPrePrepareMessage_Unions)
	return x
}

func (x *LeanHelixPrePrepareMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixPrePrepareMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixPrePrepareMessage) Body() *LeanHelixCommitPreparePPMessageSignedFields {
	b, s := x._message.GetMessage(0)
	return LeanHelixCommitPreparePPMessageSignedFieldsReader(b[:s])
}

func (x *LeanHelixPrePrepareMessage) RawBody() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixPrePrepareMessage) Signature() primitives.Ed25519Sig {
	return x._message.GetBytes(1)
}

func (x *LeanHelixPrePrepareMessage) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixPrePrepareMessage) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, v)
}

func (x *LeanHelixPrePrepareMessage) BlockPairData() []byte {
	return x._message.GetBytes(2)
}

func (x *LeanHelixPrePrepareMessage) RawBlockPairData() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixPrePrepareMessage) MutateBlockPairData(v []byte) error {
	return x._message.SetBytes(2, v)
}

// builder

type LeanHelixPrePrepareMessageBuilder struct {
	Body *LeanHelixCommitPreparePPMessageSignedFieldsBuilder
	Signature primitives.Ed25519Sig
	BlockPairData []byte

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixPrePrepareMessageBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, w.BlockPairData)
	return nil
}

func (w *LeanHelixPrePrepareMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixPrePrepareMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixPrePrepareMessageBuilder) Build() *LeanHelixPrePrepareMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixPrePrepareMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrepareMessage

// reader

type LeanHelixPrepareMessage struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixPrepareMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var _LeanHelixPrepareMessage_Unions = [][]membuffers.FieldType{}

func LeanHelixPrepareMessageReader(buf []byte) *LeanHelixPrepareMessage {
	x := &LeanHelixPrepareMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixPrepareMessage_Scheme, _LeanHelixPrepareMessage_Unions)
	return x
}

func (x *LeanHelixPrepareMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixPrepareMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixPrepareMessage) Body() *LeanHelixCommitPreparePPMessageSignedFields {
	b, s := x._message.GetMessage(0)
	return LeanHelixCommitPreparePPMessageSignedFieldsReader(b[:s])
}

func (x *LeanHelixPrepareMessage) RawBody() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixPrepareMessage) Signature() primitives.Ed25519Sig {
	return x._message.GetBytes(1)
}

func (x *LeanHelixPrepareMessage) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixPrepareMessage) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, v)
}

// builder

type LeanHelixPrepareMessageBuilder struct {
	Body *LeanHelixCommitPreparePPMessageSignedFieldsBuilder
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixPrepareMessageBuilder) Write(buf []byte) (err error) {
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

func (w *LeanHelixPrepareMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixPrepareMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixPrepareMessageBuilder) Build() *LeanHelixPrepareMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixPrepareMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixCommitMessage

// reader

type LeanHelixCommitMessage struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixCommitMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytes,}
var _LeanHelixCommitMessage_Unions = [][]membuffers.FieldType{}

func LeanHelixCommitMessageReader(buf []byte) *LeanHelixCommitMessage {
	x := &LeanHelixCommitMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixCommitMessage_Scheme, _LeanHelixCommitMessage_Unions)
	return x
}

func (x *LeanHelixCommitMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixCommitMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixCommitMessage) Body() *LeanHelixCommitPreparePPMessageSignedFields {
	b, s := x._message.GetMessage(0)
	return LeanHelixCommitPreparePPMessageSignedFieldsReader(b[:s])
}

func (x *LeanHelixCommitMessage) RawBody() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixCommitMessage) Signature() primitives.Ed25519Sig {
	return x._message.GetBytes(1)
}

func (x *LeanHelixCommitMessage) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixCommitMessage) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, v)
}

func (x *LeanHelixCommitMessage) RandomSeedShare() primitives.Bls1Sig {
	return x._message.GetBytes(2)
}

func (x *LeanHelixCommitMessage) RawRandomSeedShare() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixCommitMessage) MutateRandomSeedShare(v primitives.Bls1Sig) error {
	return x._message.SetBytes(2, v)
}

// builder

type LeanHelixCommitMessageBuilder struct {
	Body *LeanHelixCommitPreparePPMessageSignedFieldsBuilder
	Signature primitives.Ed25519Sig
	RandomSeedShare primitives.Bls1Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixCommitMessageBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, w.RandomSeedShare)
	return nil
}

func (w *LeanHelixCommitMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixCommitMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixCommitMessageBuilder) Build() *LeanHelixCommitMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixCommitMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixViewChangeMessage

// reader

type LeanHelixViewChangeMessage struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixViewChangeMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytes,}
var _LeanHelixViewChangeMessage_Unions = [][]membuffers.FieldType{}

func LeanHelixViewChangeMessageReader(buf []byte) *LeanHelixViewChangeMessage {
	x := &LeanHelixViewChangeMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixViewChangeMessage_Scheme, _LeanHelixViewChangeMessage_Unions)
	return x
}

func (x *LeanHelixViewChangeMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixViewChangeMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixViewChangeMessage) Body() *LeanHelixViewChangeSignedFields {
	b, s := x._message.GetMessage(0)
	return LeanHelixViewChangeSignedFieldsReader(b[:s])
}

func (x *LeanHelixViewChangeMessage) RawBody() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixViewChangeMessage) Signature() primitives.Ed25519Sig {
	return x._message.GetBytes(1)
}

func (x *LeanHelixViewChangeMessage) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixViewChangeMessage) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, v)
}

func (x *LeanHelixViewChangeMessage) BlockPairData() []byte {
	return x._message.GetBytes(2)
}

func (x *LeanHelixViewChangeMessage) RawBlockPairData() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixViewChangeMessage) MutateBlockPairData(v []byte) error {
	return x._message.SetBytes(2, v)
}

// builder

type LeanHelixViewChangeMessageBuilder struct {
	Body *LeanHelixViewChangeSignedFieldsBuilder
	Signature primitives.Ed25519Sig
	BlockPairData []byte

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixViewChangeMessageBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, w.BlockPairData)
	return nil
}

func (w *LeanHelixViewChangeMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixViewChangeMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixViewChangeMessageBuilder) Build() *LeanHelixViewChangeMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixViewChangeMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewViewMessage

// reader

type LeanHelixNewViewMessage struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixNewViewMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytes,}
var _LeanHelixNewViewMessage_Unions = [][]membuffers.FieldType{}

func LeanHelixNewViewMessageReader(buf []byte) *LeanHelixNewViewMessage {
	x := &LeanHelixNewViewMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixNewViewMessage_Scheme, _LeanHelixNewViewMessage_Unions)
	return x
}

func (x *LeanHelixNewViewMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixNewViewMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixNewViewMessage) Body() *LeanHelixNewViewProof {
	b, s := x._message.GetMessage(0)
	return LeanHelixNewViewProofReader(b[:s])
}

func (x *LeanHelixNewViewMessage) RawBody() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixNewViewMessage) HeaderSignature() primitives.Ed25519Sig {
	return x._message.GetBytes(1)
}

func (x *LeanHelixNewViewMessage) RawHeaderSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixNewViewMessage) MutateHeaderSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, v)
}

func (x *LeanHelixNewViewMessage) BlockPairData() []byte {
	return x._message.GetBytes(2)
}

func (x *LeanHelixNewViewMessage) RawBlockPairData() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixNewViewMessage) MutateBlockPairData(v []byte) error {
	return x._message.SetBytes(2, v)
}

// builder

type LeanHelixNewViewMessageBuilder struct {
	Body *LeanHelixNewViewProofBuilder
	HeaderSignature primitives.Ed25519Sig
	BlockPairData []byte

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixNewViewMessageBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, w.HeaderSignature)
	w._builder.WriteBytes(buf, w.BlockPairData)
	return nil
}

func (w *LeanHelixNewViewMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixNewViewMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixNewViewMessageBuilder) Build() *LeanHelixNewViewMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixNewViewMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixCommitPreparePPMessageSignedFields

// reader

type LeanHelixCommitPreparePPMessageSignedFields struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixCommitPreparePPMessageSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeBytes,membuffers.TypeBytes,}
var _LeanHelixCommitPreparePPMessageSignedFields_Unions = [][]membuffers.FieldType{}

func LeanHelixCommitPreparePPMessageSignedFieldsReader(buf []byte) *LeanHelixCommitPreparePPMessageSignedFields {
	x := &LeanHelixCommitPreparePPMessageSignedFields{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixCommitPreparePPMessageSignedFields_Scheme, _LeanHelixCommitPreparePPMessageSignedFields_Unions)
	return x
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) BlockHeight() uint64 {
	return x._message.GetUint64(0)
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) RawBlockHeight() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) MutateBlockHeight(v uint64) error {
	return x._message.SetUint64(0, v)
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) View() uint32 {
	return x._message.GetUint32(1)
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) RawView() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) MutateView(v uint32) error {
	return x._message.SetUint32(1, v)
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) BlockHash() primitives.Sha256 {
	return x._message.GetBytes(2)
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) RawBlockHash() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) MutateBlockHash(v primitives.Sha256) error {
	return x._message.SetBytes(2, v)
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) SenderPublicKey() primitives.Ed25519Pkey {
	return x._message.GetBytes(3)
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *LeanHelixCommitPreparePPMessageSignedFields) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(3, v)
}

// builder

type LeanHelixCommitPreparePPMessageSignedFieldsBuilder struct {
	BlockHeight uint64
	View uint32
	BlockHash primitives.Sha256
	SenderPublicKey primitives.Ed25519Pkey

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixCommitPreparePPMessageSignedFieldsBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint64(buf, w.BlockHeight)
	w._builder.WriteUint32(buf, w.View)
	w._builder.WriteBytes(buf, w.BlockHash)
	w._builder.WriteBytes(buf, w.SenderPublicKey)
	return nil
}

func (w *LeanHelixCommitPreparePPMessageSignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixCommitPreparePPMessageSignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixCommitPreparePPMessageSignedFieldsBuilder) Build() *LeanHelixCommitPreparePPMessageSignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixCommitPreparePPMessageSignedFieldsReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixViewChangeSignedFields

// reader

type LeanHelixViewChangeSignedFields struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixViewChangeSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeMessage,membuffers.TypeBytes,}
var _LeanHelixViewChangeSignedFields_Unions = [][]membuffers.FieldType{}

func LeanHelixViewChangeSignedFieldsReader(buf []byte) *LeanHelixViewChangeSignedFields {
	x := &LeanHelixViewChangeSignedFields{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixViewChangeSignedFields_Scheme, _LeanHelixViewChangeSignedFields_Unions)
	return x
}

func (x *LeanHelixViewChangeSignedFields) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixViewChangeSignedFields) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixViewChangeSignedFields) BlockHeight() uint64 {
	return x._message.GetUint64(0)
}

func (x *LeanHelixViewChangeSignedFields) RawBlockHeight() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixViewChangeSignedFields) MutateBlockHeight(v uint64) error {
	return x._message.SetUint64(0, v)
}

func (x *LeanHelixViewChangeSignedFields) View() uint32 {
	return x._message.GetUint32(1)
}

func (x *LeanHelixViewChangeSignedFields) RawView() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixViewChangeSignedFields) MutateView(v uint32) error {
	return x._message.SetUint32(1, v)
}

func (x *LeanHelixViewChangeSignedFields) PreparedProof() *LeanHelixPreparedProof {
	b, s := x._message.GetMessage(2)
	return LeanHelixPreparedProofReader(b[:s])
}

func (x *LeanHelixViewChangeSignedFields) RawPreparedProof() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixViewChangeSignedFields) SenderPublicKey() primitives.Ed25519Pkey {
	return x._message.GetBytes(3)
}

func (x *LeanHelixViewChangeSignedFields) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *LeanHelixViewChangeSignedFields) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(3, v)
}

// builder

type LeanHelixViewChangeSignedFieldsBuilder struct {
	BlockHeight uint64
	View uint32
	PreparedProof *LeanHelixPreparedProofBuilder
	SenderPublicKey primitives.Ed25519Pkey

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixViewChangeSignedFieldsBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint64(buf, w.BlockHeight)
	w._builder.WriteUint32(buf, w.View)
	err = w._builder.WriteMessage(buf, w.PreparedProof)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, w.SenderPublicKey)
	return nil
}

func (w *LeanHelixViewChangeSignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixViewChangeSignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixViewChangeSignedFieldsBuilder) Build() *LeanHelixViewChangeSignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixViewChangeSignedFieldsReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPreparedProof

// reader

type LeanHelixPreparedProof struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixPreparedProof_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeBytesArray,}
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

func (x *LeanHelixPreparedProof) PpBodyIterator() *LeanHelixPreparedProofPpBodyIterator {
	return &LeanHelixPreparedProofPpBodyIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type LeanHelixPreparedProofPpBodyIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixPreparedProofPpBodyIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixPreparedProofPpBodyIterator) NextPpBody() *LeanHelixCommitPreparePPMessageSignedFields {
	b, s := i.iterator.NextMessage()
	return LeanHelixCommitPreparePPMessageSignedFieldsReader(b[:s])
}

func (x *LeanHelixPreparedProof) RawPpBodyArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixPreparedProof) PpPrepareSignatureIterator() *LeanHelixPreparedProofPpPrepareSignatureIterator {
	return &LeanHelixPreparedProofPpPrepareSignatureIterator{iterator: x._message.GetBytesArrayIterator(1)}
}

type LeanHelixPreparedProofPpPrepareSignatureIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixPreparedProofPpPrepareSignatureIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixPreparedProofPpPrepareSignatureIterator) NextPpPrepareSignature() primitives.Ed25519Sig {
	return i.iterator.NextBytes()
}

func (x *LeanHelixPreparedProof) RawPpPrepareSignatureArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

// builder

type LeanHelixPreparedProofBuilder struct {
	PpBody []*LeanHelixCommitPreparePPMessageSignedFieldsBuilder
	PpPrepareSignature []primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixPreparedProofBuilder) arrayOfPpBody() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.PpBody))
	for i, v := range w.PpBody {
		res[i] = v
	}
	return res
}

func (w *LeanHelixPreparedProofBuilder) arrayOfPpPrepareSignature() [][]byte {
	res := make([][]byte, len(w.PpPrepareSignature))
	for i, v := range w.PpPrepareSignature {
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
	err = w._builder.WriteMessageArray(buf, w.arrayOfPpBody())
	if err != nil {
		return
	}
	w._builder.WriteBytesArray(buf, w.arrayOfPpPrepareSignature())
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
// message LeanHelixNewViewSignedFields

// reader

type LeanHelixNewViewSignedFields struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixNewViewSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytes,}
var _LeanHelixNewViewSignedFields_Unions = [][]membuffers.FieldType{}

func LeanHelixNewViewSignedFieldsReader(buf []byte) *LeanHelixNewViewSignedFields {
	x := &LeanHelixNewViewSignedFields{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixNewViewSignedFields_Scheme, _LeanHelixNewViewSignedFields_Unions)
	return x
}

func (x *LeanHelixNewViewSignedFields) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixNewViewSignedFields) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixNewViewSignedFields) BlockHeight() uint64 {
	return x._message.GetUint64(0)
}

func (x *LeanHelixNewViewSignedFields) RawBlockHeight() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixNewViewSignedFields) MutateBlockHeight(v uint64) error {
	return x._message.SetUint64(0, v)
}

func (x *LeanHelixNewViewSignedFields) View() uint32 {
	return x._message.GetUint32(1)
}

func (x *LeanHelixNewViewSignedFields) RawView() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixNewViewSignedFields) MutateView(v uint32) error {
	return x._message.SetUint32(1, v)
}

func (x *LeanHelixNewViewSignedFields) NewViewProof() *LeanHelixNewViewProof {
	b, s := x._message.GetMessage(2)
	return LeanHelixNewViewProofReader(b[:s])
}

func (x *LeanHelixNewViewSignedFields) RawNewViewProof() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixNewViewSignedFields) NvPpBody() *LeanHelixCommitPreparePPMessageSignedFields {
	b, s := x._message.GetMessage(3)
	return LeanHelixCommitPreparePPMessageSignedFieldsReader(b[:s])
}

func (x *LeanHelixNewViewSignedFields) RawNvPpBody() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *LeanHelixNewViewSignedFields) NvPpSignature() primitives.Ed25519Sig {
	return x._message.GetBytes(4)
}

func (x *LeanHelixNewViewSignedFields) RawNvPpSignature() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *LeanHelixNewViewSignedFields) MutateNvPpSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(4, v)
}

func (x *LeanHelixNewViewSignedFields) SenderPublicKey() primitives.Ed25519Pkey {
	return x._message.GetBytes(5)
}

func (x *LeanHelixNewViewSignedFields) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *LeanHelixNewViewSignedFields) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(5, v)
}

// builder

type LeanHelixNewViewSignedFieldsBuilder struct {
	BlockHeight uint64
	View uint32
	NewViewProof *LeanHelixNewViewProofBuilder
	NvPpBody *LeanHelixCommitPreparePPMessageSignedFieldsBuilder
	NvPpSignature primitives.Ed25519Sig
	SenderPublicKey primitives.Ed25519Pkey

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixNewViewSignedFieldsBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint64(buf, w.BlockHeight)
	w._builder.WriteUint32(buf, w.View)
	err = w._builder.WriteMessage(buf, w.NewViewProof)
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.NvPpBody)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, w.NvPpSignature)
	w._builder.WriteBytes(buf, w.SenderPublicKey)
	return nil
}

func (w *LeanHelixNewViewSignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixNewViewSignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixNewViewSignedFieldsBuilder) Build() *LeanHelixNewViewSignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixNewViewSignedFieldsReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewViewProof

// reader

type LeanHelixNewViewProof struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixNewViewProof_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeBytesArray,}
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

func (x *LeanHelixNewViewProof) ViewChangeHeaderIterator() *LeanHelixNewViewProofViewChangeHeaderIterator {
	return &LeanHelixNewViewProofViewChangeHeaderIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type LeanHelixNewViewProofViewChangeHeaderIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixNewViewProofViewChangeHeaderIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixNewViewProofViewChangeHeaderIterator) NextViewChangeHeader() *LeanHelixViewChangeSignedFields {
	b, s := i.iterator.NextMessage()
	return LeanHelixViewChangeSignedFieldsReader(b[:s])
}

func (x *LeanHelixNewViewProof) RawViewChangeHeaderArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixNewViewProof) ViewChangeHeaderSignatureIterator() *LeanHelixNewViewProofViewChangeHeaderSignatureIterator {
	return &LeanHelixNewViewProofViewChangeHeaderSignatureIterator{iterator: x._message.GetBytesArrayIterator(1)}
}

type LeanHelixNewViewProofViewChangeHeaderSignatureIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixNewViewProofViewChangeHeaderSignatureIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixNewViewProofViewChangeHeaderSignatureIterator) NextViewChangeHeaderSignature() primitives.Ed25519Sig {
	return i.iterator.NextBytes()
}

func (x *LeanHelixNewViewProof) RawViewChangeHeaderSignatureArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

// builder

type LeanHelixNewViewProofBuilder struct {
	ViewChangeHeader []*LeanHelixViewChangeSignedFieldsBuilder
	ViewChangeHeaderSignature []primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixNewViewProofBuilder) arrayOfViewChangeHeader() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.ViewChangeHeader))
	for i, v := range w.ViewChangeHeader {
		res[i] = v
	}
	return res
}

func (w *LeanHelixNewViewProofBuilder) arrayOfViewChangeHeaderSignature() [][]byte {
	res := make([][]byte, len(w.ViewChangeHeaderSignature))
	for i, v := range w.ViewChangeHeaderSignature {
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
	err = w._builder.WriteMessageArray(buf, w.arrayOfViewChangeHeader())
	if err != nil {
		return
	}
	w._builder.WriteBytesArray(buf, w.arrayOfViewChangeHeaderSignature())
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
// message LeanHelixBlockProof

// reader

type LeanHelixBlockProof struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _LeanHelixBlockProof_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytesArray,membuffers.TypeBytesArray,membuffers.TypeBytes,}
var _LeanHelixBlockProof_Unions = [][]membuffers.FieldType{}

func LeanHelixBlockProofReader(buf []byte) *LeanHelixBlockProof {
	x := &LeanHelixBlockProof{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixBlockProof_Scheme, _LeanHelixBlockProof_Unions)
	return x
}

func (x *LeanHelixBlockProof) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixBlockProof) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixBlockProof) BlockHeight() uint64 {
	return x._message.GetUint64(0)
}

func (x *LeanHelixBlockProof) RawBlockHeight() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixBlockProof) MutateBlockHeight(v uint64) error {
	return x._message.SetUint64(0, v)
}

func (x *LeanHelixBlockProof) View() uint32 {
	return x._message.GetUint32(1)
}

func (x *LeanHelixBlockProof) RawView() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixBlockProof) MutateView(v uint32) error {
	return x._message.SetUint32(1, v)
}

func (x *LeanHelixBlockProof) BlockHashMask() primitives.Uint256 {
	return x._message.GetBytes(2)
}

func (x *LeanHelixBlockProof) RawBlockHashMask() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixBlockProof) MutateBlockHashMask(v primitives.Uint256) error {
	return x._message.SetBytes(2, v)
}

func (x *LeanHelixBlockProof) BlockHash() primitives.Uint256 {
	return x._message.GetBytes(3)
}

func (x *LeanHelixBlockProof) RawBlockHash() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *LeanHelixBlockProof) MutateBlockHash(v primitives.Uint256) error {
	return x._message.SetBytes(3, v)
}

func (x *LeanHelixBlockProof) PkeyIterator() *LeanHelixBlockProofPkeyIterator {
	return &LeanHelixBlockProofPkeyIterator{iterator: x._message.GetBytesArrayIterator(4)}
}

type LeanHelixBlockProofPkeyIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixBlockProofPkeyIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixBlockProofPkeyIterator) NextPkey() primitives.Ed25519Pkey {
	return i.iterator.NextBytes()
}

func (x *LeanHelixBlockProof) RawPkeyArray() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *LeanHelixBlockProof) SignatureIterator() *LeanHelixBlockProofSignatureIterator {
	return &LeanHelixBlockProofSignatureIterator{iterator: x._message.GetBytesArrayIterator(5)}
}

type LeanHelixBlockProofSignatureIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixBlockProofSignatureIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixBlockProofSignatureIterator) NextSignature() primitives.Ed25519Sig {
	return i.iterator.NextBytes()
}

func (x *LeanHelixBlockProof) RawSignatureArray() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *LeanHelixBlockProof) RandomSeedSignature() primitives.Bls1Sig {
	return x._message.GetBytes(6)
}

func (x *LeanHelixBlockProof) RawRandomSeedSignature() []byte {
	return x._message.RawBufferForField(6, 0)
}

func (x *LeanHelixBlockProof) MutateRandomSeedSignature(v primitives.Bls1Sig) error {
	return x._message.SetBytes(6, v)
}

// builder

type LeanHelixBlockProofBuilder struct {
	BlockHeight uint64
	View uint32
	BlockHashMask primitives.Uint256
	BlockHash primitives.Uint256
	Pkey []primitives.Ed25519Pkey
	Signature []primitives.Ed25519Sig
	RandomSeedSignature primitives.Bls1Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixBlockProofBuilder) arrayOfPkey() [][]byte {
	res := make([][]byte, len(w.Pkey))
	for i, v := range w.Pkey {
		res[i] = v
	}
	return res
}

func (w *LeanHelixBlockProofBuilder) arrayOfSignature() [][]byte {
	res := make([][]byte, len(w.Signature))
	for i, v := range w.Signature {
		res[i] = v
	}
	return res
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
	w._builder.Reset()
	w._builder.WriteUint64(buf, w.BlockHeight)
	w._builder.WriteUint32(buf, w.View)
	w._builder.WriteBytes(buf, w.BlockHashMask)
	w._builder.WriteBytes(buf, w.BlockHash)
	w._builder.WriteBytesArray(buf, w.arrayOfPkey())
	w._builder.WriteBytesArray(buf, w.arrayOfSignature())
	w._builder.WriteBytes(buf, w.RandomSeedSignature)
	return nil
}

func (w *LeanHelixBlockProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixBlockProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
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

