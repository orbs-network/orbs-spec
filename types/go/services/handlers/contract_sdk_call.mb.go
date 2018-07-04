// AUTO GENERATED FILE (by membufc proto compiler)
package handlers

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ContractSdkCallHandler

type ContractSdkCallHandler interface {
	HandleSdkCall(*HandleSdkCallInput) (*HandleSdkCallOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleSdkCallInput

// reader

type HandleSdkCallInput struct {
	message membuffers.Message
}

var m_HandleSdkCallInput_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessageArray,}
var m_HandleSdkCallInput_Unions = [][]membuffers.FieldType{}

func HandleSdkCallInputReader(buf []byte) *HandleSdkCallInput {
	x := &HandleSdkCallInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleSdkCallInput_Scheme, m_HandleSdkCallInput_Unions)
	return x
}

func (x *HandleSdkCallInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleSdkCallInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleSdkCallInput) ContextId() uint32 {
	return x.message.GetUint32(0)
}

func (x *HandleSdkCallInput) RawContextId() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleSdkCallInput) MutateContextId(v uint32) error {
	return x.message.SetUint32(0, v)
}

func (x *HandleSdkCallInput) Contract() *protocol.ContractAddress {
	b, s := x.message.GetMessage(1)
	return protocol.ContractAddressReader(b[:s])
}

func (x *HandleSdkCallInput) RawContract() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *HandleSdkCallInput) Method() *protocol.MethodAddress {
	b, s := x.message.GetMessage(2)
	return protocol.MethodAddressReader(b[:s])
}

func (x *HandleSdkCallInput) RawMethod() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *HandleSdkCallInput) InputArgumentIterator() *HandleSdkCallInputInputArgumentIterator {
	return &HandleSdkCallInputInputArgumentIterator{iterator: x.message.GetMessageArrayIterator(3)}
}

type HandleSdkCallInputInputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *HandleSdkCallInputInputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *HandleSdkCallInputInputArgumentIterator) NextInputArgument() *protocol.MethodArgument {
	b, s := i.iterator.NextMessage()
	return protocol.MethodArgumentReader(b[:s])
}

func (x *HandleSdkCallInput) RawInputArgumentArray() []byte {
	return x.message.RawBufferForField(3, 0)
}

// builder

type HandleSdkCallInputBuilder struct {
	builder membuffers.Builder
	ContextId uint32
	Contract *protocol.ContractAddressBuilder
	Method *protocol.MethodAddressBuilder
	InputArgument []*protocol.MethodArgumentBuilder
}

func (w *HandleSdkCallInputBuilder) arrayOfInputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.InputArgument))
	for i, v := range w.InputArgument {
		res[i] = v
	}
	return res
}

func (w *HandleSdkCallInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.ContextId)
	err = w.builder.WriteMessage(buf, w.Contract)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.Method)
	if err != nil {
		return
	}
	err = w.builder.WriteMessageArray(buf, w.arrayOfInputArgument())
	if err != nil {
		return
	}
	return nil
}

func (w *HandleSdkCallInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleSdkCallInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleSdkCallInputBuilder) Build() *HandleSdkCallInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleSdkCallInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleSdkCallOutput

// reader

type HandleSdkCallOutput struct {
	message membuffers.Message
}

var m_HandleSdkCallOutput_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,}
var m_HandleSdkCallOutput_Unions = [][]membuffers.FieldType{}

func HandleSdkCallOutputReader(buf []byte) *HandleSdkCallOutput {
	x := &HandleSdkCallOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleSdkCallOutput_Scheme, m_HandleSdkCallOutput_Unions)
	return x
}

func (x *HandleSdkCallOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleSdkCallOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleSdkCallOutput) OutputArgumentIterator() *HandleSdkCallOutputOutputArgumentIterator {
	return &HandleSdkCallOutputOutputArgumentIterator{iterator: x.message.GetMessageArrayIterator(0)}
}

type HandleSdkCallOutputOutputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *HandleSdkCallOutputOutputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *HandleSdkCallOutputOutputArgumentIterator) NextOutputArgument() *protocol.MethodArgument {
	b, s := i.iterator.NextMessage()
	return protocol.MethodArgumentReader(b[:s])
}

func (x *HandleSdkCallOutput) RawOutputArgumentArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type HandleSdkCallOutputBuilder struct {
	builder membuffers.Builder
	OutputArgument []*protocol.MethodArgumentBuilder
}

func (w *HandleSdkCallOutputBuilder) arrayOfOutputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.OutputArgument))
	for i, v := range w.OutputArgument {
		res[i] = v
	}
	return res
}

func (w *HandleSdkCallOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessageArray(buf, w.arrayOfOutputArgument())
	if err != nil {
		return
	}
	return nil
}

func (w *HandleSdkCallOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleSdkCallOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleSdkCallOutputBuilder) Build() *HandleSdkCallOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleSdkCallOutputReader(buf)
}

