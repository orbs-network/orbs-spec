// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package gossipmessages

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncAvailabilityRequestHeader

// reader

type BlockSyncAvailabilityRequestHeader struct {
	// SenderPublicKey primitives.Ed25519Pkey
	// Signature primitives.Ed25519Sig

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *BlockSyncAvailabilityRequestHeader) String() string {
	return fmt.Sprintf("{SenderPublicKey:%s,Signature:%s,}", x.StringSenderPublicKey(), x.StringSignature())
}

var _BlockSyncAvailabilityRequestHeader_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var _BlockSyncAvailabilityRequestHeader_Unions = [][]membuffers.FieldType{}

func BlockSyncAvailabilityRequestHeaderReader(buf []byte) *BlockSyncAvailabilityRequestHeader {
	x := &BlockSyncAvailabilityRequestHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncAvailabilityRequestHeader_Scheme, _BlockSyncAvailabilityRequestHeader_Unions)
	return x
}

func (x *BlockSyncAvailabilityRequestHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncAvailabilityRequestHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncAvailabilityRequestHeader) SenderPublicKey() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(x._message.GetBytes(0))
}

func (x *BlockSyncAvailabilityRequestHeader) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncAvailabilityRequestHeader) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *BlockSyncAvailabilityRequestHeader) StringSenderPublicKey() string {
	return fmt.Sprintf("%x", x.SenderPublicKey())
}

func (x *BlockSyncAvailabilityRequestHeader) Signature() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(x._message.GetBytes(1))
}

func (x *BlockSyncAvailabilityRequestHeader) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncAvailabilityRequestHeader) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *BlockSyncAvailabilityRequestHeader) StringSignature() string {
	return fmt.Sprintf("%x", x.Signature())
}

// builder

type BlockSyncAvailabilityRequestHeaderBuilder struct {
	SenderPublicKey primitives.Ed25519Pkey
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncAvailabilityRequestHeaderBuilder) Write(buf []byte) (err error) {
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

func (w *BlockSyncAvailabilityRequestHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncAvailabilityRequestHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncAvailabilityRequestHeaderBuilder) Build() *BlockSyncAvailabilityRequestHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncAvailabilityRequestHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncAvailabilityResponseHeader

// reader

type BlockSyncAvailabilityResponseHeader struct {
	// SenderPublicKey primitives.Ed25519Pkey
	// Signature primitives.Ed25519Sig

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *BlockSyncAvailabilityResponseHeader) String() string {
	return fmt.Sprintf("{SenderPublicKey:%s,Signature:%s,}", x.StringSenderPublicKey(), x.StringSignature())
}

var _BlockSyncAvailabilityResponseHeader_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var _BlockSyncAvailabilityResponseHeader_Unions = [][]membuffers.FieldType{}

func BlockSyncAvailabilityResponseHeaderReader(buf []byte) *BlockSyncAvailabilityResponseHeader {
	x := &BlockSyncAvailabilityResponseHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncAvailabilityResponseHeader_Scheme, _BlockSyncAvailabilityResponseHeader_Unions)
	return x
}

func (x *BlockSyncAvailabilityResponseHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncAvailabilityResponseHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncAvailabilityResponseHeader) SenderPublicKey() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(x._message.GetBytes(0))
}

func (x *BlockSyncAvailabilityResponseHeader) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncAvailabilityResponseHeader) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *BlockSyncAvailabilityResponseHeader) StringSenderPublicKey() string {
	return fmt.Sprintf("%x", x.SenderPublicKey())
}

func (x *BlockSyncAvailabilityResponseHeader) Signature() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(x._message.GetBytes(1))
}

func (x *BlockSyncAvailabilityResponseHeader) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncAvailabilityResponseHeader) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *BlockSyncAvailabilityResponseHeader) StringSignature() string {
	return fmt.Sprintf("%x", x.Signature())
}

// builder

type BlockSyncAvailabilityResponseHeaderBuilder struct {
	SenderPublicKey primitives.Ed25519Pkey
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncAvailabilityResponseHeaderBuilder) Write(buf []byte) (err error) {
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

func (w *BlockSyncAvailabilityResponseHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncAvailabilityResponseHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncAvailabilityResponseHeaderBuilder) Build() *BlockSyncAvailabilityResponseHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncAvailabilityResponseHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestHeader

// reader

type BlockSyncRequestHeader struct {
	// SenderPublicKey primitives.Ed25519Pkey
	// Signature primitives.Ed25519Sig

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *BlockSyncRequestHeader) String() string {
	return fmt.Sprintf("{SenderPublicKey:%s,Signature:%s,}", x.StringSenderPublicKey(), x.StringSignature())
}

var _BlockSyncRequestHeader_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var _BlockSyncRequestHeader_Unions = [][]membuffers.FieldType{}

func BlockSyncRequestHeaderReader(buf []byte) *BlockSyncRequestHeader {
	x := &BlockSyncRequestHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncRequestHeader_Scheme, _BlockSyncRequestHeader_Unions)
	return x
}

func (x *BlockSyncRequestHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncRequestHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncRequestHeader) SenderPublicKey() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(x._message.GetBytes(0))
}

func (x *BlockSyncRequestHeader) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncRequestHeader) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *BlockSyncRequestHeader) StringSenderPublicKey() string {
	return fmt.Sprintf("%x", x.SenderPublicKey())
}

func (x *BlockSyncRequestHeader) Signature() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(x._message.GetBytes(1))
}

func (x *BlockSyncRequestHeader) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncRequestHeader) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *BlockSyncRequestHeader) StringSignature() string {
	return fmt.Sprintf("%x", x.Signature())
}

// builder

type BlockSyncRequestHeaderBuilder struct {
	SenderPublicKey primitives.Ed25519Pkey
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncRequestHeaderBuilder) Write(buf []byte) (err error) {
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

func (w *BlockSyncRequestHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncRequestHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncRequestHeaderBuilder) Build() *BlockSyncRequestHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncRequestHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncAvailability

// reader

type BlockSyncAvailability struct {
	// BlockType BlockType
	// FirstAvailableBlockHeight primitives.BlockHeight
	// LastAvailableBlockHeight primitives.BlockHeight
	// LastCommittedBlockHeight primitives.BlockHeight

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *BlockSyncAvailability) String() string {
	return fmt.Sprintf("{BlockType:%s,FirstAvailableBlockHeight:%s,LastAvailableBlockHeight:%s,LastCommittedBlockHeight:%s,}", x.StringBlockType(), x.StringFirstAvailableBlockHeight(), x.StringLastAvailableBlockHeight(), x.StringLastCommittedBlockHeight())
}

var _BlockSyncAvailability_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,}
var _BlockSyncAvailability_Unions = [][]membuffers.FieldType{}

func BlockSyncAvailabilityReader(buf []byte) *BlockSyncAvailability {
	x := &BlockSyncAvailability{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncAvailability_Scheme, _BlockSyncAvailability_Unions)
	return x
}

func (x *BlockSyncAvailability) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncAvailability) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncAvailability) BlockType() BlockType {
	return BlockType(x._message.GetUint16(0))
}

func (x *BlockSyncAvailability) RawBlockType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncAvailability) MutateBlockType(v BlockType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *BlockSyncAvailability) StringBlockType() string {
	return x.BlockType().String()
}

func (x *BlockSyncAvailability) FirstAvailableBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *BlockSyncAvailability) RawFirstAvailableBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncAvailability) MutateFirstAvailableBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *BlockSyncAvailability) StringFirstAvailableBlockHeight() string {
	return fmt.Sprintf("%x", x.FirstAvailableBlockHeight())
}

func (x *BlockSyncAvailability) LastAvailableBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *BlockSyncAvailability) RawLastAvailableBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BlockSyncAvailability) MutateLastAvailableBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *BlockSyncAvailability) StringLastAvailableBlockHeight() string {
	return fmt.Sprintf("%x", x.LastAvailableBlockHeight())
}

func (x *BlockSyncAvailability) LastCommittedBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(3))
}

func (x *BlockSyncAvailability) RawLastCommittedBlockHeight() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *BlockSyncAvailability) MutateLastCommittedBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *BlockSyncAvailability) StringLastCommittedBlockHeight() string {
	return fmt.Sprintf("%x", x.LastCommittedBlockHeight())
}

// builder

type BlockSyncAvailabilityBuilder struct {
	BlockType BlockType
	FirstAvailableBlockHeight primitives.BlockHeight
	LastAvailableBlockHeight primitives.BlockHeight
	LastCommittedBlockHeight primitives.BlockHeight

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncAvailabilityBuilder) Write(buf []byte) (err error) {
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

func (w *BlockSyncAvailabilityBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncAvailabilityBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncAvailabilityBuilder) Build() *BlockSyncAvailability {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncAvailabilityReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequest

// reader

type BlockSyncRequest struct {
	// BlockType BlockType
	// FirstRequestedBlockHeight primitives.BlockHeight
	// LastRequestedBlockHeight primitives.BlockHeight
	// LastCommittedBlockHeight primitives.BlockHeight

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *BlockSyncRequest) String() string {
	return fmt.Sprintf("{BlockType:%s,FirstRequestedBlockHeight:%s,LastRequestedBlockHeight:%s,LastCommittedBlockHeight:%s,}", x.StringBlockType(), x.StringFirstRequestedBlockHeight(), x.StringLastRequestedBlockHeight(), x.StringLastCommittedBlockHeight())
}

var _BlockSyncRequest_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,}
var _BlockSyncRequest_Unions = [][]membuffers.FieldType{}

func BlockSyncRequestReader(buf []byte) *BlockSyncRequest {
	x := &BlockSyncRequest{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncRequest_Scheme, _BlockSyncRequest_Unions)
	return x
}

func (x *BlockSyncRequest) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncRequest) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncRequest) BlockType() BlockType {
	return BlockType(x._message.GetUint16(0))
}

func (x *BlockSyncRequest) RawBlockType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncRequest) MutateBlockType(v BlockType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *BlockSyncRequest) StringBlockType() string {
	return x.BlockType().String()
}

func (x *BlockSyncRequest) FirstRequestedBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *BlockSyncRequest) RawFirstRequestedBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncRequest) MutateFirstRequestedBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *BlockSyncRequest) StringFirstRequestedBlockHeight() string {
	return fmt.Sprintf("%x", x.FirstRequestedBlockHeight())
}

func (x *BlockSyncRequest) LastRequestedBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *BlockSyncRequest) RawLastRequestedBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BlockSyncRequest) MutateLastRequestedBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *BlockSyncRequest) StringLastRequestedBlockHeight() string {
	return fmt.Sprintf("%x", x.LastRequestedBlockHeight())
}

func (x *BlockSyncRequest) LastCommittedBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(3))
}

func (x *BlockSyncRequest) RawLastCommittedBlockHeight() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *BlockSyncRequest) MutateLastCommittedBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *BlockSyncRequest) StringLastCommittedBlockHeight() string {
	return fmt.Sprintf("%x", x.LastCommittedBlockHeight())
}

// builder

type BlockSyncRequestBuilder struct {
	BlockType BlockType
	FirstRequestedBlockHeight primitives.BlockHeight
	LastRequestedBlockHeight primitives.BlockHeight
	LastCommittedBlockHeight primitives.BlockHeight

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncRequestBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUint64(buf, uint64(w.FirstRequestedBlockHeight))
	w._builder.WriteUint64(buf, uint64(w.LastRequestedBlockHeight))
	w._builder.WriteUint64(buf, uint64(w.LastCommittedBlockHeight))
	return nil
}

func (w *BlockSyncRequestBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncRequestBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncRequestBuilder) Build() *BlockSyncRequest {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncRequestReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseHeader

// reader

type BlockSyncResponseHeader struct {
	// SenderPublicKey primitives.Ed25519Pkey
	// Signature primitives.Ed25519Sig

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *BlockSyncResponseHeader) String() string {
	return fmt.Sprintf("{SenderPublicKey:%s,Signature:%s,}", x.StringSenderPublicKey(), x.StringSignature())
}

var _BlockSyncResponseHeader_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var _BlockSyncResponseHeader_Unions = [][]membuffers.FieldType{}

func BlockSyncResponseHeaderReader(buf []byte) *BlockSyncResponseHeader {
	x := &BlockSyncResponseHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncResponseHeader_Scheme, _BlockSyncResponseHeader_Unions)
	return x
}

func (x *BlockSyncResponseHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncResponseHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncResponseHeader) SenderPublicKey() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(x._message.GetBytes(0))
}

func (x *BlockSyncResponseHeader) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncResponseHeader) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *BlockSyncResponseHeader) StringSenderPublicKey() string {
	return fmt.Sprintf("%x", x.SenderPublicKey())
}

func (x *BlockSyncResponseHeader) Signature() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(x._message.GetBytes(1))
}

func (x *BlockSyncResponseHeader) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncResponseHeader) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *BlockSyncResponseHeader) StringSignature() string {
	return fmt.Sprintf("%x", x.Signature())
}

// builder

type BlockSyncResponseHeaderBuilder struct {
	SenderPublicKey primitives.Ed25519Pkey
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncResponseHeaderBuilder) Write(buf []byte) (err error) {
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

func (w *BlockSyncResponseHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncResponseHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncResponseHeaderBuilder) Build() *BlockSyncResponseHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncResponseHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponse

// reader

type BlockSyncResponse struct {
	// BlockType BlockType
	// FirstBlockHeight primitives.BlockHeight
	// LastBlockHeight primitives.BlockHeight
	// LastCommittedBlockHeight primitives.BlockHeight
	// BlockHashes []primitives.Sha256

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *BlockSyncResponse) String() string {
	return fmt.Sprintf("{BlockType:%s,FirstBlockHeight:%s,LastBlockHeight:%s,LastCommittedBlockHeight:%s,BlockHashes:%s,}", x.StringBlockType(), x.StringFirstBlockHeight(), x.StringLastBlockHeight(), x.StringLastCommittedBlockHeight(), x.StringBlockHashes())
}

var _BlockSyncResponse_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeBytesArray,}
var _BlockSyncResponse_Unions = [][]membuffers.FieldType{}

func BlockSyncResponseReader(buf []byte) *BlockSyncResponse {
	x := &BlockSyncResponse{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockSyncResponse_Scheme, _BlockSyncResponse_Unions)
	return x
}

func (x *BlockSyncResponse) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockSyncResponse) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockSyncResponse) BlockType() BlockType {
	return BlockType(x._message.GetUint16(0))
}

func (x *BlockSyncResponse) RawBlockType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockSyncResponse) MutateBlockType(v BlockType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *BlockSyncResponse) StringBlockType() string {
	return x.BlockType().String()
}

func (x *BlockSyncResponse) FirstBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *BlockSyncResponse) RawFirstBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BlockSyncResponse) MutateFirstBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *BlockSyncResponse) StringFirstBlockHeight() string {
	return fmt.Sprintf("%x", x.FirstBlockHeight())
}

func (x *BlockSyncResponse) LastBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *BlockSyncResponse) RawLastBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BlockSyncResponse) MutateLastBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *BlockSyncResponse) StringLastBlockHeight() string {
	return fmt.Sprintf("%x", x.LastBlockHeight())
}

func (x *BlockSyncResponse) LastCommittedBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(3))
}

func (x *BlockSyncResponse) RawLastCommittedBlockHeight() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *BlockSyncResponse) MutateLastCommittedBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *BlockSyncResponse) StringLastCommittedBlockHeight() string {
	return fmt.Sprintf("%x", x.LastCommittedBlockHeight())
}

func (x *BlockSyncResponse) BlockHashesIterator() *BlockSyncResponseBlockHashesIterator {
	return &BlockSyncResponseBlockHashesIterator{iterator: x._message.GetBytesArrayIterator(4)}
}

type BlockSyncResponseBlockHashesIterator struct {
	iterator *membuffers.Iterator
}

func (i *BlockSyncResponseBlockHashesIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *BlockSyncResponseBlockHashesIterator) NextBlockHashes() primitives.Sha256 {
	return primitives.Sha256(i.iterator.NextBytes())
}

func (x *BlockSyncResponse) RawBlockHashesArray() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *BlockSyncResponse) StringBlockHashes() (res string) {
	res = "["
	for i := x.BlockHashesIterator(); i.HasNext(); {
		res += fmt.Sprintf("%x", i.NextBlockHashes()) + ","
	}
	res += "]"
	return
}

// builder

type BlockSyncResponseBuilder struct {
	BlockType BlockType
	FirstBlockHeight primitives.BlockHeight
	LastBlockHeight primitives.BlockHeight
	LastCommittedBlockHeight primitives.BlockHeight
	BlockHashes []primitives.Sha256

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *BlockSyncResponseBuilder) arrayOfBlockHashes() [][]byte {
	res := make([][]byte, len(w.BlockHashes))
	for i, v := range w.BlockHashes {
		res[i] = v
	}
	return res
}

func (w *BlockSyncResponseBuilder) Write(buf []byte) (err error) {
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

func (w *BlockSyncResponseBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockSyncResponseBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BlockSyncResponseBuilder) Build() *BlockSyncResponse {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockSyncResponseReader(buf)
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

