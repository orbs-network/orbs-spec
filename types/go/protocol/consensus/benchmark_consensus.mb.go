// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
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
	// Sender BenchmarkConsensusSenderSignature

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *BenchmarkConsensusBlockProof) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Sender:%s,}", x.StringSender())
}

var _BenchmarkConsensusBlockProof_Scheme = []membuffers.FieldType{membuffers.TypeMessage}
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

func (x *BenchmarkConsensusBlockProof) Sender() *BenchmarkConsensusSenderSignature {
	b, s := x._message.GetMessage(0)
	return BenchmarkConsensusSenderSignatureReader(b[:s])
}

func (x *BenchmarkConsensusBlockProof) RawSender() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BenchmarkConsensusBlockProof) RawSenderWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *BenchmarkConsensusBlockProof) StringSender() string {
	return x.Sender().String()
}

// builder

type BenchmarkConsensusBlockProofBuilder struct {
	Sender *BenchmarkConsensusSenderSignatureBuilder

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
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
	err = w._builder.WriteMessage(buf, w.Sender)
	if err != nil {
		return
	}
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
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "BenchmarkConsensusBlockProof.Sender", w.Sender)
	if err != nil {
		return
	}
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
	// SenderPublicKey primitives.Ed25519PublicKey
	// Signature primitives.Ed25519Sig

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *BenchmarkConsensusSenderSignature) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SenderPublicKey:%s,Signature:%s,}", x.StringSenderPublicKey(), x.StringSignature())
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

func (x *BenchmarkConsensusSenderSignature) SenderPublicKey() primitives.Ed25519PublicKey {
	return primitives.Ed25519PublicKey(x._message.GetBytes(0))
}

func (x *BenchmarkConsensusSenderSignature) RawSenderPublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BenchmarkConsensusSenderSignature) RawSenderPublicKeyWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *BenchmarkConsensusSenderSignature) MutateSenderPublicKey(v primitives.Ed25519PublicKey) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *BenchmarkConsensusSenderSignature) StringSenderPublicKey() string {
	return fmt.Sprintf("%s", x.SenderPublicKey())
}

func (x *BenchmarkConsensusSenderSignature) Signature() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(x._message.GetBytes(1))
}

func (x *BenchmarkConsensusSenderSignature) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *BenchmarkConsensusSenderSignature) RawSignatureWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *BenchmarkConsensusSenderSignature) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *BenchmarkConsensusSenderSignature) StringSignature() string {
	return fmt.Sprintf("%s", x.Signature())
}

// builder

type BenchmarkConsensusSenderSignatureBuilder struct {
	SenderPublicKey primitives.Ed25519PublicKey
	Signature       primitives.Ed25519Sig

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
	w._builder.WriteBytes(buf, []byte(w.SenderPublicKey))
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
	w._builder.HexDumpBytes(prefix, offsetFromStart, "BenchmarkConsensusSenderSignature.SenderPublicKey", []byte(w.SenderPublicKey))
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
