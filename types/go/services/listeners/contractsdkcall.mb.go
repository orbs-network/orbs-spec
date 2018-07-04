// AUTO GENERATED FILE (by membufc proto compiler)
package listeners

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ContractSdkCallListener

type ContractSdkCallListener interface {
	SdkCall(*SdkCallInput) (*SdkCallOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message SdkCallInput

// reader

type SdkCallInput struct {
	message membuffers.Message
}

var m_SdkCallInput_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessageArray,}
var m_SdkCallInput_Unions = [][]membuffers.FieldType{}

func SdkCallInputReader(buf []byte) *SdkCallInput {
	x := &SdkCallInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SdkCallInput_Scheme, m_SdkCallInput_Unions)
	return x
}

func (x *SdkCallInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SdkCallInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SdkCallInput) ContextId() uint32 {
	return x.message.GetUint32(0)
}

func (x *SdkCallInput) RawContextId() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SdkCallInput) MutateContextId(v uint32) error {
	return x.message.SetUint32(0, v)
}

func (x *SdkCallInput) Contract() *protocol.ContractAddress {
	b, s := x.message.GetMessage(1)
	return protocol.ContractAddressReader(b[:s])
}

func (x *SdkCallInput) RawContract() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *SdkCallInput) Method() *protocol.MethodAddress {
	b, s := x.message.GetMessage(2)
	return protocol.MethodAddressReader(b[:s])
}

func (x *SdkCallInput) RawMethod() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *SdkCallInput) InputArgumentIterator() *SdkCallInputInputArgumentIterator {
	return &SdkCallInputInputArgumentIterator{iterator: x.message.GetMessageArrayIterator(3)}
}

type SdkCallInputInputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *SdkCallInputInputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *SdkCallInputInputArgumentIterator) NextInputArgument() *protocol.MethodArgument {
	b, s := i.iterator.NextMessage()
	return protocol.MethodArgumentReader(b[:s])
}

func (x *SdkCallInput) RawInputArgumentArray() []byte {
	return x.message.RawBufferForField(3, 0)
}

// builder

type SdkCallInputBuilder struct {
	builder membuffers.Builder
	ContextId uint32
	Contract *protocol.ContractAddressBuilder
	Method *protocol.MethodAddressBuilder
	InputArgument []*protocol.MethodArgumentBuilder
}

func (w *SdkCallInputBuilder) arrayOfInputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.InputArgument))
	for i, v := range w.InputArgument {
		res[i] = v
	}
	return res
}

func (w *SdkCallInputBuilder) Write(buf []byte) (err error) {
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

func (w *SdkCallInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SdkCallInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SdkCallInputBuilder) Build() *SdkCallInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SdkCallInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SdkCallOutput

// reader

type SdkCallOutput struct {
	message membuffers.Message
}

var m_SdkCallOutput_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,}
var m_SdkCallOutput_Unions = [][]membuffers.FieldType{}

func SdkCallOutputReader(buf []byte) *SdkCallOutput {
	x := &SdkCallOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SdkCallOutput_Scheme, m_SdkCallOutput_Unions)
	return x
}

func (x *SdkCallOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SdkCallOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SdkCallOutput) OutputArgumentIterator() *SdkCallOutputOutputArgumentIterator {
	return &SdkCallOutputOutputArgumentIterator{iterator: x.message.GetMessageArrayIterator(0)}
}

type SdkCallOutputOutputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *SdkCallOutputOutputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *SdkCallOutputOutputArgumentIterator) NextOutputArgument() *protocol.MethodArgument {
	b, s := i.iterator.NextMessage()
	return protocol.MethodArgumentReader(b[:s])
}

func (x *SdkCallOutput) RawOutputArgumentArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type SdkCallOutputBuilder struct {
	builder membuffers.Builder
	OutputArgument []*protocol.MethodArgumentBuilder
}

func (w *SdkCallOutputBuilder) arrayOfOutputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.OutputArgument))
	for i, v := range w.OutputArgument {
		res[i] = v
	}
	return res
}

func (w *SdkCallOutputBuilder) Write(buf []byte) (err error) {
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

func (w *SdkCallOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SdkCallOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SdkCallOutputBuilder) Build() *SdkCallOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SdkCallOutputReader(buf)
}

