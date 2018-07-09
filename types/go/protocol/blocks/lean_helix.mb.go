// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package blocks

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

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

