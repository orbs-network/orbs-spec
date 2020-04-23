// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package gossipmessages

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/consensus"
)

/////////////////////////////////////////////////////////////////////////////
// message Header

// reader

type Header struct {
	// ProtocolVersion primitives.ProtocolVersion
	// VirtualChainId primitives.VirtualChainId
	// RecipientNodeAddresses []primitives.NodeAddress
	// RecipientMode RecipientsListMode
	// Topic HeaderTopic

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *Header) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ProtocolVersion:%s,VirtualChainId:%s,RecipientNodeAddresses:%s,RecipientMode:%s,Topic:%s,}", x.StringProtocolVersion(), x.StringVirtualChainId(), x.StringRecipientNodeAddresses(), x.StringRecipientMode(), x.StringTopic())
}

var _Header_Scheme = []membuffers.FieldType{membuffers.TypeUint32, membuffers.TypeUint32, membuffers.TypeBytesArray, membuffers.TypeUint16, membuffers.TypeUnion}
var _Header_Unions = [][]membuffers.FieldType{{membuffers.TypeUint16, membuffers.TypeUint16, membuffers.TypeUint16, membuffers.TypeUint16, membuffers.TypeUint16}}

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
	return fmt.Sprintf("%s", x.ProtocolVersion())
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
	return fmt.Sprintf("%s", x.VirtualChainId())
}

func (x *Header) RecipientNodeAddressesIterator() *HeaderRecipientNodeAddressesIterator {
	return &HeaderRecipientNodeAddressesIterator{iterator: x._message.GetBytesArrayIterator(2)}
}

type HeaderRecipientNodeAddressesIterator struct {
	iterator *membuffers.Iterator
}

func (i *HeaderRecipientNodeAddressesIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *HeaderRecipientNodeAddressesIterator) NextRecipientNodeAddresses() primitives.NodeAddress {
	return primitives.NodeAddress(i.iterator.NextBytes())
}

func (x *Header) RawRecipientNodeAddressesArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *Header) RawRecipientNodeAddressesArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *Header) StringRecipientNodeAddresses() (res string) {
	res = "["
	for i := x.RecipientNodeAddressesIterator(); i.HasNext(); {
		res += fmt.Sprintf("%s", i.NextRecipientNodeAddresses()) + ","
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
	HEADER_TOPIC_TRANSACTION_RELAY   HeaderTopic = 0
	HEADER_TOPIC_BLOCK_SYNC          HeaderTopic = 1
	HEADER_TOPIC_LEAN_HELIX          HeaderTopic = 2
	HEADER_TOPIC_BENCHMARK_CONSENSUS HeaderTopic = 3
	HEADER_TOPIC_HEADER_SYNC         HeaderTopic = 4
)

func (x *Header) Topic() HeaderTopic {
	return HeaderTopic(x._message.GetUnionIndex(4, 0))
}

func (x *Header) IsTopicTransactionRelay() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 0)
	return is
}

func (x *Header) TransactionRelay() TransactionsRelayMessageType {
	is, off := x._message.IsUnionIndex(4, 0, 0)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
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
	is, off := x._message.IsUnionIndex(4, 0, 1)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
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
	is, off := x._message.IsUnionIndex(4, 0, 2)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
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
	is, off := x._message.IsUnionIndex(4, 0, 3)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
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

func (x *Header) IsTopicHeaderSync() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 4)
	return is
}

func (x *Header) HeaderSync() HeaderSyncMessageType {
	is, off := x._message.IsUnionIndex(4, 0, 4)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return HeaderSyncMessageType(x._message.GetUint16InOffset(off))
}

func (x *Header) StringHeaderSync() string {
	return x.HeaderSync().String()
}

func (x *Header) MutateHeaderSync(v HeaderSyncMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 4)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *Header) RawTopic() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *Header) RawTopicWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(4, 0)
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
	case HEADER_TOPIC_HEADER_SYNC:
		return "(HeaderSync)" + x.StringHeaderSync()
	}
	return "(Unknown)"
}

// builder

type HeaderBuilder struct {
	ProtocolVersion        primitives.ProtocolVersion
	VirtualChainId         primitives.VirtualChainId
	RecipientNodeAddresses []primitives.NodeAddress
	RecipientMode          RecipientsListMode
	Topic                  HeaderTopic
	TransactionRelay       TransactionsRelayMessageType
	BlockSync              BlockSyncMessageType
	LeanHelix              consensus.LeanHelixMessageType
	BenchmarkConsensus     consensus.BenchmarkConsensusMessageType
	HeaderSync             HeaderSyncMessageType

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *HeaderBuilder) arrayOfRecipientNodeAddresses() [][]byte {
	res := make([][]byte, len(w.RecipientNodeAddresses))
	for i, v := range w.RecipientNodeAddresses {
		res[i] = v
	}
	return res
}

func (w *HeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteUint32(buf, uint32(w.ProtocolVersion))
	w._builder.WriteUint32(buf, uint32(w.VirtualChainId))
	w._builder.WriteBytesArray(buf, w.arrayOfRecipientNodeAddresses())
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
	case HEADER_TOPIC_HEADER_SYNC:
		w._builder.WriteUint16(buf, uint16(w.HeaderSync))
	}
	return nil
}

func (w *HeaderBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint32(prefix, offsetFromStart, "Header.ProtocolVersion", uint32(w.ProtocolVersion))
	w._builder.HexDumpUint32(prefix, offsetFromStart, "Header.VirtualChainId", uint32(w.VirtualChainId))
	w._builder.HexDumpBytesArray(prefix, offsetFromStart, "Header.RecipientNodeAddresses", w.arrayOfRecipientNodeAddresses())
	w._builder.HexDumpUint16(prefix, offsetFromStart, "Header.RecipientMode", uint16(w.RecipientMode))
	w._builder.HexDumpUnionIndex(prefix, offsetFromStart, "Header.Topic", uint16(w.Topic))
	switch w.Topic {
	case HEADER_TOPIC_TRANSACTION_RELAY:
		w._builder.HexDumpUint16(prefix, offsetFromStart, "Header.TransactionRelay", uint16(w.TransactionRelay))
	case HEADER_TOPIC_BLOCK_SYNC:
		w._builder.HexDumpUint16(prefix, offsetFromStart, "Header.BlockSync", uint16(w.BlockSync))
	case HEADER_TOPIC_LEAN_HELIX:
		w._builder.HexDumpUint16(prefix, offsetFromStart, "Header.LeanHelix", uint16(w.LeanHelix))
	case HEADER_TOPIC_BENCHMARK_CONSENSUS:
		w._builder.HexDumpUint16(prefix, offsetFromStart, "Header.BenchmarkConsensus", uint16(w.BenchmarkConsensus))
	case HEADER_TOPIC_HEADER_SYNC:
		w._builder.HexDumpUint16(prefix, offsetFromStart, "Header.HeaderSync", uint16(w.HeaderSync))
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

func HeaderBuilderFromRaw(raw []byte) *HeaderBuilder {
	return &HeaderBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message SenderSignature

// reader

type SenderSignature struct {
	// SenderNodeAddress primitives.NodeAddress
	// Signature primitives.EcdsaSecp256K1Sig

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *SenderSignature) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SenderNodeAddress:%s,Signature:%s,}", x.StringSenderNodeAddress(), x.StringSignature())
}

var _SenderSignature_Scheme = []membuffers.FieldType{membuffers.TypeBytes, membuffers.TypeBytes}
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

func (x *SenderSignature) SenderNodeAddress() primitives.NodeAddress {
	return primitives.NodeAddress(x._message.GetBytes(0))
}

func (x *SenderSignature) RawSenderNodeAddress() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SenderSignature) RawSenderNodeAddressWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *SenderSignature) MutateSenderNodeAddress(v primitives.NodeAddress) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *SenderSignature) StringSenderNodeAddress() string {
	return fmt.Sprintf("%s", x.SenderNodeAddress())
}

func (x *SenderSignature) Signature() primitives.EcdsaSecp256K1Sig {
	return primitives.EcdsaSecp256K1Sig(x._message.GetBytes(1))
}

func (x *SenderSignature) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *SenderSignature) RawSignatureWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *SenderSignature) MutateSignature(v primitives.EcdsaSecp256K1Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *SenderSignature) StringSignature() string {
	return fmt.Sprintf("%s", x.Signature())
}

// builder

type SenderSignatureBuilder struct {
	SenderNodeAddress primitives.NodeAddress
	Signature         primitives.EcdsaSecp256K1Sig

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *SenderSignatureBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteBytes(buf, []byte(w.SenderNodeAddress))
	w._builder.WriteBytes(buf, []byte(w.Signature))
	return nil
}

func (w *SenderSignatureBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "SenderSignature.SenderNodeAddress", []byte(w.SenderNodeAddress))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "SenderSignature.Signature", []byte(w.Signature))
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

func SenderSignatureBuilderFromRaw(raw []byte) *SenderSignatureBuilder {
	return &SenderSignatureBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// enums

type RecipientsListMode uint16

const (
	RECIPIENT_LIST_MODE_BROADCAST    RecipientsListMode = 0
	RECIPIENT_LIST_MODE_LIST         RecipientsListMode = 1
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
