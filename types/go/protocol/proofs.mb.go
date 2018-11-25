// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package protocol

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message ReceiptProof

// reader

type ReceiptProof struct {
	// Header ResultsBlockHeader
	// BlockProof ResultsBlockProof
	// ReceiptProof primitives.MerkleTreeProof
	// ReceiptIndex TransactionReceipt
	// Receipt TransactionReceipt

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *ReceiptProof) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Header:%s,BlockProof:%s,ReceiptProof:%s,ReceiptIndex:%s,Receipt:%s,}", x.StringHeader(), x.StringBlockProof(), x.StringReceiptProof(), x.StringReceiptIndex(), x.StringReceipt())
}

var _ReceiptProof_Scheme = []membuffers.FieldType{membuffers.TypeMessage, membuffers.TypeMessage, membuffers.TypeBytes, membuffers.TypeMessage, membuffers.TypeMessage}
var _ReceiptProof_Unions = [][]membuffers.FieldType{}

func ReceiptProofReader(buf []byte) *ReceiptProof {
	x := &ReceiptProof{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ReceiptProof_Scheme, _ReceiptProof_Unions)
	return x
}

func (x *ReceiptProof) IsValid() bool {
	return x._message.IsValid()
}

func (x *ReceiptProof) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ReceiptProof) Equal(y *ReceiptProof) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *ReceiptProof) Header() *ResultsBlockHeader {
	b, s := x._message.GetMessage(0)
	return ResultsBlockHeaderReader(b[:s])
}

func (x *ReceiptProof) RawHeader() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ReceiptProof) RawHeaderWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *ReceiptProof) StringHeader() string {
	return x.Header().String()
}

func (x *ReceiptProof) BlockProof() *ResultsBlockProof {
	b, s := x._message.GetMessage(1)
	return ResultsBlockProofReader(b[:s])
}

func (x *ReceiptProof) RawBlockProof() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *ReceiptProof) RawBlockProofWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *ReceiptProof) StringBlockProof() string {
	return x.BlockProof().String()
}

func (x *ReceiptProof) ReceiptProof() primitives.MerkleTreeProof {
	return primitives.MerkleTreeProof(x._message.GetBytes(2))
}

func (x *ReceiptProof) RawReceiptProof() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *ReceiptProof) MutateReceiptProof(v primitives.MerkleTreeProof) error {
	return x._message.SetBytes(2, []byte(v))
}

func (x *ReceiptProof) StringReceiptProof() string {
	return fmt.Sprintf("%s", x.ReceiptProof())
}

func (x *ReceiptProof) ReceiptIndex() *TransactionReceipt {
	b, s := x._message.GetMessage(3)
	return TransactionReceiptReader(b[:s])
}

func (x *ReceiptProof) RawReceiptIndex() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *ReceiptProof) RawReceiptIndexWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *ReceiptProof) StringReceiptIndex() string {
	return x.ReceiptIndex().String()
}

func (x *ReceiptProof) Receipt() *TransactionReceipt {
	b, s := x._message.GetMessage(4)
	return TransactionReceiptReader(b[:s])
}

func (x *ReceiptProof) RawReceipt() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *ReceiptProof) RawReceiptWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(4, 0)
}

func (x *ReceiptProof) StringReceipt() string {
	return x.Receipt().String()
}

// builder

type ReceiptProofBuilder struct {
	Header       *ResultsBlockHeaderBuilder
	BlockProof   *ResultsBlockProofBuilder
	ReceiptProof primitives.MerkleTreeProof
	ReceiptIndex *TransactionReceiptBuilder
	Receipt      *TransactionReceiptBuilder

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *ReceiptProofBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.Header)
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.BlockProof)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, []byte(w.ReceiptProof))
	err = w._builder.WriteMessage(buf, w.ReceiptIndex)
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.Receipt)
	if err != nil {
		return
	}
	return nil
}

func (w *ReceiptProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ReceiptProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ReceiptProofBuilder) Build() *ReceiptProof {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ReceiptProofReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums
