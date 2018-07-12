// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package blockproofs

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message LeanHelix

// reader

type LeanHelix struct {
	// BlockHeight primitives.BlockHeight
	// View uint32
	// BlockHashMask primitives.Uint256
	// BlockHash primitives.Uint256
	// NodePublicKeys []primitives.Ed25519Pkey
	// NodeSignatures []primitives.Ed25519Sig
	// RandomSeedSignature primitives.Bls1Sig

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *LeanHelix) String() string {
	return fmt.Sprintf("{BlockHeight:%s,View:%s,BlockHashMask:%s,BlockHash:%s,NodePublicKeys:%s,NodeSignatures:%s,RandomSeedSignature:%s,}", x.StringBlockHeight(), x.StringView(), x.StringBlockHashMask(), x.StringBlockHash(), x.StringNodePublicKeys(), x.StringNodeSignatures(), x.StringRandomSeedSignature())
}

var _LeanHelix_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytesArray,membuffers.TypeBytesArray,membuffers.TypeBytes,}
var _LeanHelix_Unions = [][]membuffers.FieldType{}

func LeanHelixReader(buf []byte) *LeanHelix {
	x := &LeanHelix{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelix_Scheme, _LeanHelix_Unions)
	return x
}

func (x *LeanHelix) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelix) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelix) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(0))
}

func (x *LeanHelix) RawBlockHeight() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelix) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(0, uint64(v))
}

func (x *LeanHelix) StringBlockHeight() string {
	return fmt.Sprintf("%x", x.BlockHeight())
}

func (x *LeanHelix) View() uint32 {
	return x._message.GetUint32(1)
}

func (x *LeanHelix) RawView() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelix) MutateView(v uint32) error {
	return x._message.SetUint32(1, v)
}

func (x *LeanHelix) StringView() string {
	return fmt.Sprintf("%x", x.View())
}

func (x *LeanHelix) BlockHashMask() primitives.Uint256 {
	return primitives.Uint256(x._message.GetBytes(2))
}

func (x *LeanHelix) RawBlockHashMask() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelix) MutateBlockHashMask(v primitives.Uint256) error {
	return x._message.SetBytes(2, []byte(v))
}

func (x *LeanHelix) StringBlockHashMask() string {
	return fmt.Sprintf("%x", x.BlockHashMask())
}

func (x *LeanHelix) BlockHash() primitives.Uint256 {
	return primitives.Uint256(x._message.GetBytes(3))
}

func (x *LeanHelix) RawBlockHash() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *LeanHelix) MutateBlockHash(v primitives.Uint256) error {
	return x._message.SetBytes(3, []byte(v))
}

func (x *LeanHelix) StringBlockHash() string {
	return fmt.Sprintf("%x", x.BlockHash())
}

func (x *LeanHelix) NodePublicKeysIterator() *LeanHelixNodePublicKeysIterator {
	return &LeanHelixNodePublicKeysIterator{iterator: x._message.GetBytesArrayIterator(4)}
}

type LeanHelixNodePublicKeysIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixNodePublicKeysIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixNodePublicKeysIterator) NextNodePublicKeys() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(i.iterator.NextBytes())
}

func (x *LeanHelix) RawNodePublicKeysArray() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *LeanHelix) StringNodePublicKeys() (res string) {
	res = "["
	for i := x.NodePublicKeysIterator(); i.HasNext(); {
		res += fmt.Sprintf("%x", i.NextNodePublicKeys()) + ","
	}
	res += "]"
	return
}

func (x *LeanHelix) NodeSignaturesIterator() *LeanHelixNodeSignaturesIterator {
	return &LeanHelixNodeSignaturesIterator{iterator: x._message.GetBytesArrayIterator(5)}
}

type LeanHelixNodeSignaturesIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixNodeSignaturesIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixNodeSignaturesIterator) NextNodeSignatures() primitives.Ed25519Sig {
	return primitives.Ed25519Sig(i.iterator.NextBytes())
}

func (x *LeanHelix) RawNodeSignaturesArray() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *LeanHelix) StringNodeSignatures() (res string) {
	res = "["
	for i := x.NodeSignaturesIterator(); i.HasNext(); {
		res += fmt.Sprintf("%x", i.NextNodeSignatures()) + ","
	}
	res += "]"
	return
}

func (x *LeanHelix) RandomSeedSignature() primitives.Bls1Sig {
	return primitives.Bls1Sig(x._message.GetBytes(6))
}

func (x *LeanHelix) RawRandomSeedSignature() []byte {
	return x._message.RawBufferForField(6, 0)
}

func (x *LeanHelix) MutateRandomSeedSignature(v primitives.Bls1Sig) error {
	return x._message.SetBytes(6, []byte(v))
}

func (x *LeanHelix) StringRandomSeedSignature() string {
	return fmt.Sprintf("%x", x.RandomSeedSignature())
}

// builder

type LeanHelixBuilder struct {
	BlockHeight primitives.BlockHeight
	View uint32
	BlockHashMask primitives.Uint256
	BlockHash primitives.Uint256
	NodePublicKeys []primitives.Ed25519Pkey
	NodeSignatures []primitives.Ed25519Sig
	RandomSeedSignature primitives.Bls1Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixBuilder) arrayOfNodePublicKeys() [][]byte {
	res := make([][]byte, len(w.NodePublicKeys))
	for i, v := range w.NodePublicKeys {
		res[i] = v
	}
	return res
}

func (w *LeanHelixBuilder) arrayOfNodeSignatures() [][]byte {
	res := make([][]byte, len(w.NodeSignatures))
	for i, v := range w.NodeSignatures {
		res[i] = v
	}
	return res
}

func (w *LeanHelixBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, []byte(w.BlockHashMask))
	w._builder.WriteBytes(buf, []byte(w.BlockHash))
	w._builder.WriteBytesArray(buf, w.arrayOfNodePublicKeys())
	w._builder.WriteBytesArray(buf, w.arrayOfNodeSignatures())
	w._builder.WriteBytes(buf, []byte(w.RandomSeedSignature))
	return nil
}

func (w *LeanHelixBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixBuilder) Build() *LeanHelix {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixReader(buf)
}

