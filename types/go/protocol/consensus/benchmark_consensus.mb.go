// AUTO GENERATED FILE (by membufc proto compiler v0.3.6)
package consensus

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message BenchmarkConsensusBlockProof

// reader

type BenchmarkConsensusBlockProof struct {
	// BlockRef BenchmarkConsensusBlockRef
	// Nodes []BenchmarkConsensusSenderSignature
	// Placeholder []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *BenchmarkConsensusBlockProof) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockRef:%s,Nodes:%s,Placeholder:%s,}", x.StringBlockRef(), x.StringNodes(), x.StringPlaceholder())
}

var _BenchmarkConsensusBlockProof_Scheme = []membuffers.FieldType{membuffers.TypeMessage, membuffers.TypeMessageArray, membuffers.TypeBytes}
var _BenchmarkConsensusBlockProof_Unions = [][]membuffers.FieldType{}

func BenchmarkConsensusBlockProofReader(buf []byte) *BenchmarkConsensusBlockProof {
	x := &BenchmarkConsensusBlockProof{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BenchmarkConsensusBlockProof_Scheme, _BenchmarkConsensusBlockProof_Unions)
	return x
}

func (x *BenchmarkConsensusBlockProof) IsValid() bool {
	return x._message.IsValid()
}

func (x *BenchmarkConsensusBlockProof) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BenchmarkConsensusBlockProof) Equal(y *BenchmarkConsensusBlockProof) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *BenchmarkConsensusBlockProof) BlockRef() *BenchmarkConsensusBlockRef {
	b, s := x._message.GetMessage(0)
	return BenchmarkConsensusBlockRefReader(b[:s])
}

func (x *BenchmarkConsensusBlockProof) RawBlockRef() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BenchmarkConsensusBlockProof) RawBlockRefWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *BenchmarkConsensusBlockProof) StringBlockRef() string {
	return x.BlockRef().String()
}

func (x *BenchmarkConsensusBlockProof) NodesIterator() *BenchmarkConsensusBlockProofNodesIterator {
	return &BenchmarkConsensusBlockProofNodesIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type BenchmarkConsensusBlockProofNodesIterator struct {
	iterator *membuffers.Iterator
}

func (i *BenchmarkConsensusBlockProofNodesIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *BenchmarkConsensusBlockProofNodesIterator) NextNodes() *BenchmarkConsensusSenderSignature {
	b, s := i.iterator.NextMessage()
	return BenchmarkConsensusSenderSignatureReader(b[:s])
}

func (x *BenchmarkConsensusBlockProof) RawNodesArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BenchmarkConsensusBlockProof) RawNodesArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *BenchmarkConsensusBlockProof) StringNodes() (res string) {
	res = "["
	for i := x.NodesIterator(); i.HasNext(); {
		res += i.NextNodes().String() + ","
	}
	res += "]"
	return
}

func (x *BenchmarkConsensusBlockProof) Placeholder() []byte {
	return x._message.GetBytes(2)
}

func (x *BenchmarkConsensusBlockProof) RawPlaceholder() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BenchmarkConsensusBlockProof) RawPlaceholderWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *BenchmarkConsensusBlockProof) MutatePlaceholder(v []byte) error {
	return x._message.SetBytes(2, v)
}

func (x *BenchmarkConsensusBlockProof) StringPlaceholder() string {
	return fmt.Sprintf("%x", x.Placeholder())
}

// builder

type BenchmarkConsensusBlockProofBuilder struct {
	BlockRef    *BenchmarkConsensusBlockRefBuilder
	Nodes       []*BenchmarkConsensusSenderSignatureBuilder
	Placeholder []byte

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *BenchmarkConsensusBlockProofBuilder) arrayOfNodes() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Nodes))
	for i, v := range w.Nodes {
		res[i] = v
	}
	return res
}

func (w *BenchmarkConsensusBlockProofBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessage(buf, w.BlockRef)
	if err != nil {
		return
	}
	err = w._builder.WriteMessageArray(buf, w.arrayOfNodes())
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, w.Placeholder)
	return nil
}

func (w *BenchmarkConsensusBlockProofBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "BenchmarkConsensusBlockProof.BlockRef", w.BlockRef)
	if err != nil {
		return
	}
	err = w._builder.HexDumpMessageArray(prefix, offsetFromStart, "BenchmarkConsensusBlockProof.Nodes", w.arrayOfNodes())
	if err != nil {
		return
	}
	w._builder.HexDumpBytes(prefix, offsetFromStart, "BenchmarkConsensusBlockProof.Placeholder", w.Placeholder)
	return nil
}

func (w *BenchmarkConsensusBlockProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BenchmarkConsensusBlockProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BenchmarkConsensusBlockProofBuilder) Build() *BenchmarkConsensusBlockProof {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BenchmarkConsensusBlockProofReader(buf)
}

func BenchmarkConsensusBlockProofBuilderFromRaw(raw []byte) *BenchmarkConsensusBlockProofBuilder {
	return &BenchmarkConsensusBlockProofBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message BenchmarkConsensusSenderSignature

// reader

type BenchmarkConsensusSenderSignature struct {
	// SenderNodeAddress primitives.NodeAddress
	// Signature primitives.EcdsaSecp256K1Sig

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *BenchmarkConsensusSenderSignature) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SenderNodeAddress:%s,Signature:%s,}", x.StringSenderNodeAddress(), x.StringSignature())
}

var _BenchmarkConsensusSenderSignature_Scheme = []membuffers.FieldType{membuffers.TypeBytes, membuffers.TypeBytes}
var _BenchmarkConsensusSenderSignature_Unions = [][]membuffers.FieldType{}

func BenchmarkConsensusSenderSignatureReader(buf []byte) *BenchmarkConsensusSenderSignature {
	x := &BenchmarkConsensusSenderSignature{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BenchmarkConsensusSenderSignature_Scheme, _BenchmarkConsensusSenderSignature_Unions)
	return x
}

func (x *BenchmarkConsensusSenderSignature) IsValid() bool {
	return x._message.IsValid()
}

func (x *BenchmarkConsensusSenderSignature) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BenchmarkConsensusSenderSignature) Equal(y *BenchmarkConsensusSenderSignature) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *BenchmarkConsensusSenderSignature) SenderNodeAddress() primitives.NodeAddress {
	return primitives.NodeAddress(x._message.GetBytes(0))
}

func (x *BenchmarkConsensusSenderSignature) RawSenderNodeAddress() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BenchmarkConsensusSenderSignature) RawSenderNodeAddressWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *BenchmarkConsensusSenderSignature) MutateSenderNodeAddress(v primitives.NodeAddress) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *BenchmarkConsensusSenderSignature) StringSenderNodeAddress() string {
	return fmt.Sprintf("%s", x.SenderNodeAddress())
}

func (x *BenchmarkConsensusSenderSignature) Signature() primitives.EcdsaSecp256K1Sig {
	return primitives.EcdsaSecp256K1Sig(x._message.GetBytes(1))
}

func (x *BenchmarkConsensusSenderSignature) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BenchmarkConsensusSenderSignature) RawSignatureWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *BenchmarkConsensusSenderSignature) MutateSignature(v primitives.EcdsaSecp256K1Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *BenchmarkConsensusSenderSignature) StringSignature() string {
	return fmt.Sprintf("%s", x.Signature())
}

// builder

type BenchmarkConsensusSenderSignatureBuilder struct {
	SenderNodeAddress primitives.NodeAddress
	Signature         primitives.EcdsaSecp256K1Sig

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *BenchmarkConsensusSenderSignatureBuilder) Write(buf []byte) (err error) {
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

func (w *BenchmarkConsensusSenderSignatureBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "BenchmarkConsensusSenderSignature.SenderNodeAddress", []byte(w.SenderNodeAddress))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "BenchmarkConsensusSenderSignature.Signature", []byte(w.Signature))
	return nil
}

func (w *BenchmarkConsensusSenderSignatureBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BenchmarkConsensusSenderSignatureBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BenchmarkConsensusSenderSignatureBuilder) Build() *BenchmarkConsensusSenderSignature {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BenchmarkConsensusSenderSignatureReader(buf)
}

func BenchmarkConsensusSenderSignatureBuilderFromRaw(raw []byte) *BenchmarkConsensusSenderSignatureBuilder {
	return &BenchmarkConsensusSenderSignatureBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message BenchmarkConsensusBlockRef

// reader

type BenchmarkConsensusBlockRef struct {
	// PlaceholderType BenchmarkConsensusPlaceholder
	// BlockHeight primitives.BlockHeight
	// PlaceholderView uint64
	// BlockHash primitives.Sha256

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *BenchmarkConsensusBlockRef) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{PlaceholderType:%s,BlockHeight:%s,PlaceholderView:%s,BlockHash:%s,}", x.StringPlaceholderType(), x.StringBlockHeight(), x.StringPlaceholderView(), x.StringBlockHash())
}

var _BenchmarkConsensusBlockRef_Scheme = []membuffers.FieldType{membuffers.TypeUint16, membuffers.TypeUint64, membuffers.TypeUint64, membuffers.TypeBytes}
var _BenchmarkConsensusBlockRef_Unions = [][]membuffers.FieldType{}

func BenchmarkConsensusBlockRefReader(buf []byte) *BenchmarkConsensusBlockRef {
	x := &BenchmarkConsensusBlockRef{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BenchmarkConsensusBlockRef_Scheme, _BenchmarkConsensusBlockRef_Unions)
	return x
}

func (x *BenchmarkConsensusBlockRef) IsValid() bool {
	return x._message.IsValid()
}

func (x *BenchmarkConsensusBlockRef) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BenchmarkConsensusBlockRef) Equal(y *BenchmarkConsensusBlockRef) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *BenchmarkConsensusBlockRef) PlaceholderType() BenchmarkConsensusPlaceholder {
	return BenchmarkConsensusPlaceholder(x._message.GetUint16(0))
}

func (x *BenchmarkConsensusBlockRef) RawPlaceholderType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BenchmarkConsensusBlockRef) MutatePlaceholderType(v BenchmarkConsensusPlaceholder) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *BenchmarkConsensusBlockRef) StringPlaceholderType() string {
	return x.PlaceholderType().String()
}

func (x *BenchmarkConsensusBlockRef) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *BenchmarkConsensusBlockRef) RawBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BenchmarkConsensusBlockRef) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *BenchmarkConsensusBlockRef) StringBlockHeight() string {
	return fmt.Sprintf("%s", x.BlockHeight())
}

func (x *BenchmarkConsensusBlockRef) PlaceholderView() uint64 {
	return x._message.GetUint64(2)
}

func (x *BenchmarkConsensusBlockRef) RawPlaceholderView() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *BenchmarkConsensusBlockRef) MutatePlaceholderView(v uint64) error {
	return x._message.SetUint64(2, v)
}

func (x *BenchmarkConsensusBlockRef) StringPlaceholderView() string {
	return fmt.Sprintf("%x", x.PlaceholderView())
}

func (x *BenchmarkConsensusBlockRef) BlockHash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(3))
}

func (x *BenchmarkConsensusBlockRef) RawBlockHash() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *BenchmarkConsensusBlockRef) RawBlockHashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *BenchmarkConsensusBlockRef) MutateBlockHash(v primitives.Sha256) error {
	return x._message.SetBytes(3, []byte(v))
}

func (x *BenchmarkConsensusBlockRef) StringBlockHash() string {
	return fmt.Sprintf("%s", x.BlockHash())
}

// builder

type BenchmarkConsensusBlockRefBuilder struct {
	PlaceholderType BenchmarkConsensusPlaceholder
	BlockHeight     primitives.BlockHeight
	PlaceholderView uint64
	BlockHash       primitives.Sha256

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *BenchmarkConsensusBlockRefBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUint16(buf, uint16(w.PlaceholderType))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteUint64(buf, w.PlaceholderView)
	w._builder.WriteBytes(buf, []byte(w.BlockHash))
	return nil
}

func (w *BenchmarkConsensusBlockRefBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint16(prefix, offsetFromStart, "BenchmarkConsensusBlockRef.PlaceholderType", uint16(w.PlaceholderType))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "BenchmarkConsensusBlockRef.BlockHeight", uint64(w.BlockHeight))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "BenchmarkConsensusBlockRef.PlaceholderView", w.PlaceholderView)
	w._builder.HexDumpBytes(prefix, offsetFromStart, "BenchmarkConsensusBlockRef.BlockHash", []byte(w.BlockHash))
	return nil
}

func (w *BenchmarkConsensusBlockRefBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BenchmarkConsensusBlockRefBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *BenchmarkConsensusBlockRefBuilder) Build() *BenchmarkConsensusBlockRef {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BenchmarkConsensusBlockRefReader(buf)
}

func BenchmarkConsensusBlockRefBuilderFromRaw(raw []byte) *BenchmarkConsensusBlockRefBuilder {
	return &BenchmarkConsensusBlockRefBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// enums

type BenchmarkConsensusMessageType uint16

const (
	BENCHMARK_CONSENSUS_RESERVED  BenchmarkConsensusMessageType = 0
	BENCHMARK_CONSENSUS_COMMIT    BenchmarkConsensusMessageType = 1
	BENCHMARK_CONSENSUS_COMMITTED BenchmarkConsensusMessageType = 2
)

func (n BenchmarkConsensusMessageType) String() string {
	switch n {
	case BENCHMARK_CONSENSUS_RESERVED:
		return "BENCHMARK_CONSENSUS_RESERVED"
	case BENCHMARK_CONSENSUS_COMMIT:
		return "BENCHMARK_CONSENSUS_COMMIT"
	case BENCHMARK_CONSENSUS_COMMITTED:
		return "BENCHMARK_CONSENSUS_COMMITTED"
	}
	return "UNKNOWN"
}

type BenchmarkConsensusPlaceholder uint16

const (
	BENCHMARK_CONSENSUS_INVALID_0 BenchmarkConsensusPlaceholder = 0
	BENCHMARK_CONSENSUS_INVALID_1 BenchmarkConsensusPlaceholder = 1
	BENCHMARK_CONSENSUS_INVALID_2 BenchmarkConsensusPlaceholder = 2
	BENCHMARK_CONSENSUS_VALID     BenchmarkConsensusPlaceholder = 3
)

func (n BenchmarkConsensusPlaceholder) String() string {
	switch n {
	case BENCHMARK_CONSENSUS_INVALID_0:
		return "BENCHMARK_CONSENSUS_INVALID_0"
	case BENCHMARK_CONSENSUS_INVALID_1:
		return "BENCHMARK_CONSENSUS_INVALID_1"
	case BENCHMARK_CONSENSUS_INVALID_2:
		return "BENCHMARK_CONSENSUS_INVALID_2"
	case BENCHMARK_CONSENSUS_VALID:
		return "BENCHMARK_CONSENSUS_VALID"
	}
	return "UNKNOWN"
}
