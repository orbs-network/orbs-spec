// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service CrosschainConnector

type CrosschainConnector interface {
	EthereumCallContract(*EthereumCallContractInput) (*EthereumCallContractOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractInput

// reader

type EthereumCallContractInput struct {
	message membuffers.Message
}

var m_EthereumCallContractInput_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeString,membuffers.TypeMessageArray,membuffers.TypeUint64,membuffers.TypeUint16,}
var m_EthereumCallContractInput_Unions = [][]membuffers.FieldType{}

func EthereumCallContractInputReader(buf []byte) *EthereumCallContractInput {
	x := &EthereumCallContractInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_EthereumCallContractInput_Scheme, m_EthereumCallContractInput_Unions)
	return x
}

func (x *EthereumCallContractInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *EthereumCallContractInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *EthereumCallContractInput) EthereumContractAddress() string {
	return x.message.GetString(0)
}

func (x *EthereumCallContractInput) RawEthereumContractAddress() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *EthereumCallContractInput) MutateEthereumContractAddress(v string) error {
	return x.message.SetString(0, v)
}

func (x *EthereumCallContractInput) EthereumFunctionCanonicalForm() string {
	return x.message.GetString(1)
}

func (x *EthereumCallContractInput) RawEthereumFunctionCanonicalForm() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *EthereumCallContractInput) MutateEthereumFunctionCanonicalForm(v string) error {
	return x.message.SetString(1, v)
}

func (x *EthereumCallContractInput) InputArgumentIterator() *EthereumCallContractInputInputArgumentIterator {
	return &EthereumCallContractInputInputArgumentIterator{iterator: x.message.GetMessageArrayIterator(2)}
}

type EthereumCallContractInputInputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *EthereumCallContractInputInputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *EthereumCallContractInputInputArgumentIterator) NextInputArgument() *protocol.MethodArgument {
	b, s := i.iterator.NextMessage()
	return protocol.MethodArgumentReader(b[:s])
}

func (x *EthereumCallContractInput) RawInputArgumentArray() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *EthereumCallContractInput) EthereumBlockHeight() uint64 {
	return x.message.GetUint64(3)
}

func (x *EthereumCallContractInput) RawEthereumBlockHeight() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *EthereumCallContractInput) MutateEthereumBlockHeight(v uint64) error {
	return x.message.SetUint64(3, v)
}

func (x *EthereumCallContractInput) EthereumBlockHeightMode() BlockHeightMode {
	return BlockHeightMode(x.message.GetUint16(4))
}

func (x *EthereumCallContractInput) RawEthereumBlockHeightMode() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *EthereumCallContractInput) MutateEthereumBlockHeightMode(v BlockHeightMode) error {
	return x.message.SetUint16(4, uint16(v))
}

// builder

type EthereumCallContractInputBuilder struct {
	builder membuffers.Builder
	EthereumContractAddress string
	EthereumFunctionCanonicalForm string
	InputArgument []*protocol.MethodArgumentBuilder
	EthereumBlockHeight uint64
	EthereumBlockHeightMode BlockHeightMode
}

func (w *EthereumCallContractInputBuilder) arrayOfInputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.InputArgument))
	for i, v := range w.InputArgument {
		res[i] = v
	}
	return res
}

func (w *EthereumCallContractInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteString(buf, w.EthereumContractAddress)
	w.builder.WriteString(buf, w.EthereumFunctionCanonicalForm)
	err = w.builder.WriteMessageArray(buf, w.arrayOfInputArgument())
	if err != nil {
		return
	}
	w.builder.WriteUint64(buf, w.EthereumBlockHeight)
	w.builder.WriteUint16(buf, uint16(w.EthereumBlockHeightMode))
	return nil
}

func (w *EthereumCallContractInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *EthereumCallContractInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *EthereumCallContractInputBuilder) Build() *EthereumCallContractInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return EthereumCallContractInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractOutput

// reader

type EthereumCallContractOutput struct {
	message membuffers.Message
}

var m_EthereumCallContractOutput_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,}
var m_EthereumCallContractOutput_Unions = [][]membuffers.FieldType{}

func EthereumCallContractOutputReader(buf []byte) *EthereumCallContractOutput {
	x := &EthereumCallContractOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_EthereumCallContractOutput_Scheme, m_EthereumCallContractOutput_Unions)
	return x
}

func (x *EthereumCallContractOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *EthereumCallContractOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *EthereumCallContractOutput) OutputArgumentIterator() *EthereumCallContractOutputOutputArgumentIterator {
	return &EthereumCallContractOutputOutputArgumentIterator{iterator: x.message.GetMessageArrayIterator(0)}
}

type EthereumCallContractOutputOutputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *EthereumCallContractOutputOutputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *EthereumCallContractOutputOutputArgumentIterator) NextOutputArgument() *protocol.MethodArgument {
	b, s := i.iterator.NextMessage()
	return protocol.MethodArgumentReader(b[:s])
}

func (x *EthereumCallContractOutput) RawOutputArgumentArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *EthereumCallContractOutput) CallStatus() protocol.CallMethodStatus {
	return protocol.CallMethodStatus(x.message.GetUint16(1))
}

func (x *EthereumCallContractOutput) RawCallStatus() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *EthereumCallContractOutput) MutateCallStatus(v protocol.CallMethodStatus) error {
	return x.message.SetUint16(1, uint16(v))
}

func (x *EthereumCallContractOutput) EthereumBlockHeight() uint64 {
	return x.message.GetUint64(2)
}

func (x *EthereumCallContractOutput) RawEthereumBlockHeight() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *EthereumCallContractOutput) MutateEthereumBlockHeight(v uint64) error {
	return x.message.SetUint64(2, v)
}

func (x *EthereumCallContractOutput) EthereumBlockTimestamp() uint64 {
	return x.message.GetUint64(3)
}

func (x *EthereumCallContractOutput) RawEthereumBlockTimestamp() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *EthereumCallContractOutput) MutateEthereumBlockTimestamp(v uint64) error {
	return x.message.SetUint64(3, v)
}

// builder

type EthereumCallContractOutputBuilder struct {
	builder membuffers.Builder
	OutputArgument []*protocol.MethodArgumentBuilder
	CallStatus protocol.CallMethodStatus
	EthereumBlockHeight uint64
	EthereumBlockTimestamp uint64
}

func (w *EthereumCallContractOutputBuilder) arrayOfOutputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.OutputArgument))
	for i, v := range w.OutputArgument {
		res[i] = v
	}
	return res
}

func (w *EthereumCallContractOutputBuilder) Write(buf []byte) (err error) {
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
	w.builder.WriteUint16(buf, uint16(w.CallStatus))
	w.builder.WriteUint64(buf, w.EthereumBlockHeight)
	w.builder.WriteUint64(buf, w.EthereumBlockTimestamp)
	return nil
}

func (w *EthereumCallContractOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *EthereumCallContractOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *EthereumCallContractOutputBuilder) Build() *EthereumCallContractOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return EthereumCallContractOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type BlockHeightMode uint16

const (
	BLOCK_HEIGHT_MODE_RESERVED BlockHeightMode = 0
	BLOCK_HEIGHT_MODE_USE_RECENT BlockHeightMode = 1
	BLOCK_HEIGHT_MODE_USE_SPECIFIED BlockHeightMode = 2
)

