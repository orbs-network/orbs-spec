// AUTO GENERATED FILE (by membufc proto compiler v0.0.17)
package gossipmessages

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"bytes"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/consensus"
)

/////////////////////////////////////////////////////////////////////////////
// message Header

// reader

type Header struct {
	// ProtocolVersion primitives.ProtocolVersion
	// VirtualChainId primitives.VirtualChainId
	// RecipientPublicKeys []primitives.Ed25519Pkey
	// RecipientMode RecipientsListMode
	// Topic HeaderTopic

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *Header) String() string {
	return fmt.Sprintf("{ProtocolVersion:%s,VirtualChainId:%s,RecipientPublicKeys:%s,RecipientMode:%s,Topic:%s,}", x.StringProtocolVersion(), x.StringVirtualChainId(), x.StringRecipientPublicKeys(), x.StringRecipientMode(), x.StringTopic())
}

var _Header_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint32,membuffers.TypeBytesArray,membuffers.TypeUint16,membuffers.TypeUnion,}
var _Header_Unions = [][]membuffers.FieldType{{membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeUint16,}}

func HeaderReader(buf []byte) *Header {
	x := &Header{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _Header_Scheme, _Header_Unions)
	return x
}

func (x *Header) IsValid() bool {
	return x._message.IsValid()
}

func (x *Header) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *Header) Equal(y *Header) bool {
  if x == nil && y == nil {
    return true
  }
  if x == nil || y == nil {
    return false
  }
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *Header) ProtocolVersion() primitives.ProtocolVersion {
	return primitives.ProtocolVersion(x._message.GetUint32(0))
}

func (x *Header) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Header) MutateProtocolVersion(v primitives.ProtocolVersion) error {
	return x._message.SetUint32(0, uint32(v))
}

func (x *Header) StringProtocolVersion() string {
	return fmt.Sprintf("%x", x.ProtocolVersion())
}

func (x *Header) VirtualChainId() primitives.VirtualChainId {
	return primitives.VirtualChainId(x._message.GetUint32(1))
}

func (x *Header) RawVirtualChainId() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *Header) MutateVirtualChainId(v primitives.VirtualChainId) error {
	return x._message.SetUint32(1, uint32(v))
}

func (x *Header) StringVirtualChainId() string {
	return fmt.Sprintf("%x", x.VirtualChainId())
}

func (x *Header) RecipientPublicKeysIterator() *HeaderRecipientPublicKeysIterator {
	return &HeaderRecipientPublicKeysIterator{iterator: x._message.GetBytesArrayIterator(2)}
}

type HeaderRecipientPublicKeysIterator struct {
	iterator *membuffers.Iterator
}

func (i *HeaderRecipientPublicKeysIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *HeaderRecipientPublicKeysIterator) NextRecipientPublicKeys() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(i.iterator.NextBytes())
}

func (x *Header) RawRecipientPublicKeysArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *Header) StringRecipientPublicKeys() (res string) {
	res = "["
	for i := x.RecipientPublicKeysIterator(); i.HasNext(); {
		res += fmt.Sprintf("%x", i.NextRecipientPublicKeys()) + ","
	}
	res += "]"
	return
}

func (x *Header) RecipientMode() RecipientsListMode {
	return RecipientsListMode(x._message.GetUint16(3))
}

func (x *Header) RawRecipientMode() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *Header) MutateRecipientMode(v RecipientsListMode) error {
	return x._message.SetUint16(3, uint16(v))
}

func (x *Header) StringRecipientMode() string {
	return x.RecipientMode().String()
}

type HeaderTopic uint16

const (
	HEADER_TOPIC_TRANSACTION_RELAY HeaderTopic = 0
	HEADER_TOPIC_BLOCK_SYNC HeaderTopic = 1
	HEADER_TOPIC_LEAN_HELIX HeaderTopic = 2
	HEADER_TOPIC_BENCHMARK_CONSENSUS HeaderTopic = 3
)

func (x *Header) Topic() HeaderTopic {
	return HeaderTopic(x._message.GetUint16(4))
}

func (x *Header) IsTopicTransactionRelay() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 0)
	return is
}

func (x *Header) TransactionRelay() TransactionsRelayMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 0)
	return TransactionsRelayMessageType(x._message.GetUint16InOffset(off))
}

func (x *Header) StringTransactionRelay() string {
	return x.TransactionRelay().String()
}

func (x *Header) MutateTransactionRelay(v TransactionsRelayMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *Header) IsTopicBlockSync() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 1)
	return is
}

func (x *Header) BlockSync() BlockSyncMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 1)
	return BlockSyncMessageType(x._message.GetUint16InOffset(off))
}

func (x *Header) StringBlockSync() string {
	return x.BlockSync().String()
}

func (x *Header) MutateBlockSync(v BlockSyncMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *Header) IsTopicLeanHelix() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 2)
	return is
}

func (x *Header) LeanHelix() consensus.LeanHelixMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 2)
	return consensus.LeanHelixMessageType(x._message.GetUint16InOffset(off))
}

func (x *Header) StringLeanHelix() string {
	return x.LeanHelix().String()
}

func (x *Header) MutateLeanHelix(v consensus.LeanHelixMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *Header) IsTopicBenchmarkConsensus() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 3)
	return is
}

func (x *Header) BenchmarkConsensus() consensus.BenchmarkConsensusMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 3)
	return consensus.BenchmarkConsensusMessageType(x._message.GetUint16InOffset(off))
}

func (x *Header) StringBenchmarkConsensus() string {
	return x.BenchmarkConsensus().String()
}

func (x *Header) MutateBenchmarkConsensus(v consensus.BenchmarkConsensusMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 3)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *Header) RawTopic() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *Header) StringTopic() string {
	switch x.Topic() {
	case HEADER_TOPIC_TRANSACTION_RELAY:
		return "(TransactionRelay)" + x.StringTransactionRelay()
	case HEADER_TOPIC_BLOCK_SYNC:
		return "(BlockSync)" + x.StringBlockSync()
	case HEADER_TOPIC_LEAN_HELIX:
		return "(LeanHelix)" + x.StringLeanHelix()
	case HEADER_TOPIC_BENCHMARK_CONSENSUS:
		return "(BenchmarkConsensus)" + x.StringBenchmarkConsensus()
	}
	return "(Unknown)"
}

// builder

type HeaderBuilder struct {
	ProtocolVersion primitives.ProtocolVersion
	VirtualChainId primitives.VirtualChainId
	RecipientPublicKeys []primitives.Ed25519Pkey
	RecipientMode RecipientsListMode
	Topic HeaderTopic
	TransactionRelay TransactionsRelayMessageType
	BlockSync BlockSyncMessageType
	LeanHelix consensus.LeanHelixMessageType
	BenchmarkConsensus consensus.BenchmarkConsensusMessageType

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *HeaderBuilder) arrayOfRecipientPublicKeys() [][]byte {
	res := make([][]byte, len(w.RecipientPublicKeys))
	for i, v := range w.RecipientPublicKeys {
		res[i] = v
	}
	return res
}

func (w *HeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint32(buf, uint32(w.ProtocolVersion))
	w._builder.WriteUint32(buf, uint32(w.VirtualChainId))
	w._builder.WriteBytesArray(buf, w.arrayOfRecipientPublicKeys())
	w._builder.WriteUint16(buf, uint16(w.RecipientMode))
	w._builder.WriteUnionIndex(buf, uint16(w.Topic))
	switch w.Topic {
	case HEADER_TOPIC_TRANSACTION_RELAY:
		w._builder.WriteUint16(buf, uint16(w.TransactionRelay))
	case HEADER_TOPIC_BLOCK_SYNC:
		w._builder.WriteUint16(buf, uint16(w.BlockSync))
	case HEADER_TOPIC_LEAN_HELIX:
		w._builder.WriteUint16(buf, uint16(w.LeanHelix))
	case HEADER_TOPIC_BENCHMARK_CONSENSUS:
		w._builder.WriteUint16(buf, uint16(w.BenchmarkConsensus))
	}
	return nil
}

func (w *HeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *HeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *HeaderBuilder) Build() *Header {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SenderSignature

// reader

type SenderSignature struct {
	// SenderPublicKey primitives.Ed25519Pkey
	// Signature primitives.Ed25519Sig

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *SenderSignature) String() string {
	return fmt.Sprintf("{SenderPublicKey:%s,Signature:%s,}", x.StringSenderPublicKey(), x.StringSignature())
}

var _SenderSignature_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var _SenderSignature_Unions = [][]membuffers.FieldType{}

func SenderSignatureReader(buf []byte) *SenderSignature {
	x := &SenderSignature{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _SenderSignature_Scheme, _SenderSignature_Unions)
	return x
}

func (x *SenderSignature) IsValid() bool {
	return x._message.IsValid()
}

func (x *SenderSignature) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *SenderSignature) Equal(y *SenderSignature) bool {
  if x == nil && y == nil {
    return true
  }
  if x == nil || y == nil {
    return false
  }
  return bytes.Equal(x.Raw(), y.Raw())
}

func (x *SenderSignature) SenderPublicKey() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(x._message.GetBytes(0))
}

func (x *SenderSignature) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SenderSignature) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *SenderSignature) StringSenderPublicKey() string {
	return fmt.Sprintf("%x", x.SenderPublicKey())
}

func (x *SenderSignature) Signature() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(x._message.GetBytes(1))
}

func (x *SenderSignature) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *SenderSignature) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *SenderSignature) StringSignature() string {
	return fmt.Sprintf("%x", x.Signature())
}

// builder

type SenderSignatureBuilder struct {
	SenderPublicKey primitives.Ed25519Pkey
	Signature primitives.Ed25519Sig

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *SenderSignatureBuilder) Write(buf []byte) (err error) {
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

func (w *SenderSignatureBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *SenderSignatureBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *SenderSignatureBuilder) Build() *SenderSignature {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SenderSignatureReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type RecipientsListMode uint16

const (
	RECIPIENT_LIST_MODE_BROADCAST RecipientsListMode = 0
	RECIPIENT_LIST_MODE_LIST RecipientsListMode = 1
	RECIPIENT_LIST_MODE_ALL_BUT_LIST RecipientsListMode = 2
)

func (n RecipientsListMode) String() string {
	switch n {
	case RECIPIENT_LIST_MODE_BROADCAST:
		return "RECIPIENT_LIST_MODE_BROADCAST"
	case RECIPIENT_LIST_MODE_LIST:
		return "RECIPIENT_LIST_MODE_LIST"
	case RECIPIENT_LIST_MODE_ALL_BUT_LIST:
		return "RECIPIENT_LIST_MODE_ALL_BUT_LIST"
	}
	return "UNKNOWN"
}

