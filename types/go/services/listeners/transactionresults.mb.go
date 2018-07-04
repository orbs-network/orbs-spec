// AUTO GENERATED FILE (by membufc proto compiler)
package listeners

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionResultsListener

type TransactionResultsListener interface {
	ReturnTransactionResults(*ReturnTransactionResultsInput) (*ReturnTransactionResultsOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message ReturnTransactionResultsInput

// reader

type ReturnTransactionResultsInput struct {
	message membuffers.Message
}

var m_ReturnTransactionResultsInput_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeUint64,membuffers.TypeUint64,}
var m_ReturnTransactionResultsInput_Unions = [][]membuffers.FieldType{}

func ReturnTransactionResultsInputReader(buf []byte) *ReturnTransactionResultsInput {
	x := &ReturnTransactionResultsInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ReturnTransactionResultsInput_Scheme, m_ReturnTransactionResultsInput_Unions)
	return x
}

func (x *ReturnTransactionResultsInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ReturnTransactionResultsInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ReturnTransactionResultsInput) TransactionReceiptIterator() *ReturnTransactionResultsInputTransactionReceiptIterator {
	return &ReturnTransactionResultsInputTransactionReceiptIterator{iterator: x.message.GetMessageArrayIterator(0)}
}

type ReturnTransactionResultsInputTransactionReceiptIterator struct {
	iterator *membuffers.Iterator
}

func (i *ReturnTransactionResultsInputTransactionReceiptIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ReturnTransactionResultsInputTransactionReceiptIterator) NextTransactionReceipt() *protocol.TransactionReceipt {
	b, s := i.iterator.NextMessage()
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *ReturnTransactionResultsInput) RawTransactionReceiptArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ReturnTransactionResultsInput) BlockHeight() uint64 {
	return x.message.GetUint64(1)
}

func (x *ReturnTransactionResultsInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ReturnTransactionResultsInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(1, v)
}

func (x *ReturnTransactionResultsInput) Timestamp() uint64 {
	return x.message.GetUint64(2)
}

func (x *ReturnTransactionResultsInput) RawTimestamp() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *ReturnTransactionResultsInput) MutateTimestamp(v uint64) error {
	return x.message.SetUint64(2, v)
}

// builder

type ReturnTransactionResultsInputBuilder struct {
	builder membuffers.Builder
	TransactionReceipt []*protocol.TransactionReceiptBuilder
	BlockHeight uint64
	Timestamp uint64
}

func (w *ReturnTransactionResultsInputBuilder) arrayOfTransactionReceipt() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.TransactionReceipt))
	for i, v := range w.TransactionReceipt {
		res[i] = v
	}
	return res
}

func (w *ReturnTransactionResultsInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessageArray(buf, w.arrayOfTransactionReceipt())
	if err != nil {
		return
	}
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteUint64(buf, w.Timestamp)
	return nil
}

func (w *ReturnTransactionResultsInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ReturnTransactionResultsInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ReturnTransactionResultsInputBuilder) Build() *ReturnTransactionResultsInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ReturnTransactionResultsInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ReturnTransactionResultsOutput

// reader

type ReturnTransactionResultsOutput struct {
	message membuffers.Message
}

var m_ReturnTransactionResultsOutput_Scheme = []membuffers.FieldType{}
var m_ReturnTransactionResultsOutput_Unions = [][]membuffers.FieldType{}

func ReturnTransactionResultsOutputReader(buf []byte) *ReturnTransactionResultsOutput {
	x := &ReturnTransactionResultsOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ReturnTransactionResultsOutput_Scheme, m_ReturnTransactionResultsOutput_Unions)
	return x
}

func (x *ReturnTransactionResultsOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ReturnTransactionResultsOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type ReturnTransactionResultsOutputBuilder struct {
	builder membuffers.Builder
}

func (w *ReturnTransactionResultsOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	return nil
}

func (w *ReturnTransactionResultsOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ReturnTransactionResultsOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ReturnTransactionResultsOutputBuilder) Build() *ReturnTransactionResultsOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ReturnTransactionResultsOutputReader(buf)
}

