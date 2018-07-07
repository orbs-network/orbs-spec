// AUTO GENERATED FILE (by membufc proto compiler)
package handlers

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionResultsHandler

type TransactionResultsHandler interface {
	HandleTransactionResults(input *HandleTransactionResultsInput) (*HandleTransactionResultsOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionResultsInput

// reader

type HandleTransactionResultsInput struct {
	message membuffers.Message
}

var m_HandleTransactionResultsInput_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeUint64,membuffers.TypeUint64,}
var m_HandleTransactionResultsInput_Unions = [][]membuffers.FieldType{}

func HandleTransactionResultsInputReader(buf []byte) *HandleTransactionResultsInput {
	x := &HandleTransactionResultsInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleTransactionResultsInput_Scheme, m_HandleTransactionResultsInput_Unions)
	return x
}

func (x *HandleTransactionResultsInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleTransactionResultsInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleTransactionResultsInput) TransactionReceiptIterator() *HandleTransactionResultsInputTransactionReceiptIterator {
	return &HandleTransactionResultsInputTransactionReceiptIterator{iterator: x.message.GetMessageArrayIterator(0)}
}

type HandleTransactionResultsInputTransactionReceiptIterator struct {
	iterator *membuffers.Iterator
}

func (i *HandleTransactionResultsInputTransactionReceiptIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *HandleTransactionResultsInputTransactionReceiptIterator) NextTransactionReceipt() *protocol.TransactionReceipt {
	b, s := i.iterator.NextMessage()
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *HandleTransactionResultsInput) RawTransactionReceiptArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleTransactionResultsInput) BlockHeight() uint64 {
	return x.message.GetUint64(1)
}

func (x *HandleTransactionResultsInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *HandleTransactionResultsInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(1, v)
}

func (x *HandleTransactionResultsInput) Timestamp() uint64 {
	return x.message.GetUint64(2)
}

func (x *HandleTransactionResultsInput) RawTimestamp() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *HandleTransactionResultsInput) MutateTimestamp(v uint64) error {
	return x.message.SetUint64(2, v)
}

// builder

type HandleTransactionResultsInputBuilder struct {
	builder membuffers.Builder
	TransactionReceipt []*protocol.TransactionReceiptBuilder
	BlockHeight uint64
	Timestamp uint64
}

func (w *HandleTransactionResultsInputBuilder) arrayOfTransactionReceipt() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.TransactionReceipt))
	for i, v := range w.TransactionReceipt {
		res[i] = v
	}
	return res
}

func (w *HandleTransactionResultsInputBuilder) Write(buf []byte) (err error) {
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

func (w *HandleTransactionResultsInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleTransactionResultsInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleTransactionResultsInputBuilder) Build() *HandleTransactionResultsInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleTransactionResultsInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionResultsOutput

// reader

type HandleTransactionResultsOutput struct {
	message membuffers.Message
}

var m_HandleTransactionResultsOutput_Scheme = []membuffers.FieldType{}
var m_HandleTransactionResultsOutput_Unions = [][]membuffers.FieldType{}

func HandleTransactionResultsOutputReader(buf []byte) *HandleTransactionResultsOutput {
	x := &HandleTransactionResultsOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleTransactionResultsOutput_Scheme, m_HandleTransactionResultsOutput_Unions)
	return x
}

func (x *HandleTransactionResultsOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleTransactionResultsOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type HandleTransactionResultsOutputBuilder struct {
	builder membuffers.Builder
}

func (w *HandleTransactionResultsOutputBuilder) Write(buf []byte) (err error) {
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

func (w *HandleTransactionResultsOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleTransactionResultsOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleTransactionResultsOutputBuilder) Build() *HandleTransactionResultsOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleTransactionResultsOutputReader(buf)
}

