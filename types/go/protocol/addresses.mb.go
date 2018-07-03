// AUTO GENERATED FILE (by membufc proto compiler)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message Sender

// reader

type Sender struct {
	message membuffers.Message
}

var m_Sender_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUnion,}
var m_Sender_Unions = [][]membuffers.FieldType{{membuffers.TypeMessage,membuffers.TypeMessage,}}

func SenderReader(buf []byte) *Sender {
	x := &Sender{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_Sender_Scheme, m_Sender_Unions)
	return x
}

func (x *Sender) IsValid() bool {
	return x.message.IsValid()
}

func (x *Sender) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *Sender) AddressScheme() AddressScheme {
	return AddressScheme(x.message.GetUint16(0))
}

func (x *Sender) RawAddressScheme() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *Sender) MutateAddressScheme(v AddressScheme) error {
	return x.message.SetUint16(0, uint16(v))
}

type SenderSenderData uint16

const (
	SenderSenderDataEddsaSender SenderSenderData = 0
	SenderSenderDataSmartContractSender SenderSenderData = 1
)

func (x *Sender) SenderData() SenderSenderData {
	return SenderSenderData(x.message.GetUint16(1))
}

func (x *Sender) IsSenderDataEddsaSender() bool {
	is, _ := x.message.IsUnionIndex(1, 0, 0)
	return is
}

func (x *Sender) SenderDataEddsaSender() EdDSA01Sender {
	_, off := x.message.IsUnionIndex(1, 0, 0)
	return x.message.GetMessageInOffset(off)
}



func (x *Sender) IsSenderDataSmartContractSender() bool {
	is, _ := x.message.IsUnionIndex(1, 0, 1)
	return is
}

func (x *Sender) SenderDataSmartContractSender() SmartContractSender {
	_, off := x.message.IsUnionIndex(1, 0, 1)
	return x.message.GetMessageInOffset(off)
}



func (x *Sender) RawSenderData() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type SenderBuilder struct {
	builder membuffers.Builder
	AddressScheme AddressScheme
	SenderData SenderSenderData
	EddsaSender EdDSA01Sender
	SmartContractSender SmartContractSender
}

func (w *SenderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.AddressScheme))
	w.builder.WriteUnionIndex(buf, uint16(w.SenderData))
	switch w.SenderData {
	case SenderSenderDataEddsaSender:
		w.builder.WriteMessage(buf, w.EddsaSender)
	case SenderSenderDataSmartContractSender:
		w.builder.WriteMessage(buf, w.SmartContractSender)
	}
	return nil
}

func (w *SenderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SenderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SenderBuilder) Build() *Sender {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SenderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message EdDSA01Sender

// reader

type EdDSA01Sender struct {
	message membuffers.Message
}

var m_EdDSA01Sender_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeBytes,}
var m_EdDSA01Sender_Unions = [][]membuffers.FieldType{}

func EdDSA01SenderReader(buf []byte) *EdDSA01Sender {
	x := &EdDSA01Sender{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_EdDSA01Sender_Scheme, m_EdDSA01Sender_Unions)
	return x
}

func (x *EdDSA01Sender) IsValid() bool {
	return x.message.IsValid()
}

func (x *EdDSA01Sender) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *EdDSA01Sender) NetworkType() NetworkType {
	return NetworkType(x.message.GetUint16(0))
}

func (x *EdDSA01Sender) RawNetworkType() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *EdDSA01Sender) MutateNetworkType(v NetworkType) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *EdDSA01Sender) SenderPublicKey() []byte {
	return x.message.GetBytes(1)
}

func (x *EdDSA01Sender) RawSenderPublicKey() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *EdDSA01Sender) MutateSenderPublicKey(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type EdDSA01SenderBuilder struct {
	builder membuffers.Builder
	NetworkType NetworkType
	SenderPublicKey []byte
}

func (w *EdDSA01SenderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.NetworkType))
	w.builder.WriteBytes(buf, w.SenderPublicKey)
	return nil
}

func (w *EdDSA01SenderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *EdDSA01SenderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *EdDSA01SenderBuilder) Build() *EdDSA01Sender {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return EdDSA01SenderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SmartContractSender

// reader

type SmartContractSender struct {
	message membuffers.Message
}

var m_SmartContractSender_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeString,membuffers.TypeMessageArray,}
var m_SmartContractSender_Unions = [][]membuffers.FieldType{}

func SmartContractSenderReader(buf []byte) *SmartContractSender {
	x := &SmartContractSender{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SmartContractSender_Scheme, m_SmartContractSender_Unions)
	return x
}

func (x *SmartContractSender) IsValid() bool {
	return x.message.IsValid()
}

func (x *SmartContractSender) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SmartContractSender) NetworkType() NetworkType {
	return NetworkType(x.message.GetUint16(0))
}

func (x *SmartContractSender) RawNetworkType() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SmartContractSender) MutateNetworkType(v NetworkType) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *SmartContractSender) ContractName() string {
	return x.message.GetString(1)
}

func (x *SmartContractSender) RawContractName() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *SmartContractSender) MutateContractName(v string) error {
	return x.message.SetString(1, v)
}

func (x *SmartContractSender) SenderInputArgumentIterator() *SmartContractSenderSenderInputArgumentIterator {
	return &SmartContractSenderSenderInputArgumentIterator{iterator: x.message.GetMessageArrayIterator(2)}
}

type SmartContractSenderSenderInputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *SmartContractSenderSenderInputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *SmartContractSenderSenderInputArgumentIterator) NextSenderInputArgument() *MethodArgument {
	b, s := i.iterator.NextMessage()
	return MethodArgumentReader(b[:s])
}

func (x *SmartContractSender) RawSenderInputArgumentArray() []byte {
	return x.message.RawBufferForField(2, 0)
}

// builder

type SmartContractSenderBuilder struct {
	builder membuffers.Builder
	NetworkType NetworkType
	ContractName string
	SenderInputArgument []*MethodArgumentBuilder
}

func (w *SmartContractSenderBuilder) arrayOfSenderInputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.SenderInputArgument))
	for i, v := range w.SenderInputArgument {
		res[i] = v
	}
	return res
}

func (w *SmartContractSenderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.NetworkType))
	w.builder.WriteString(buf, w.ContractName)
	err = w.builder.WriteMessageArray(buf, w.arrayOfSenderInputArgument())
	if err != nil {
		return
	}
	return nil
}

func (w *SmartContractSenderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SmartContractSenderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SmartContractSenderBuilder) Build() *SmartContractSender {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SmartContractSenderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type NetworkType uint16

const (
	NETWORK_TYPE_RESERVED NetworkType = 0
	NETWORK_TYPE_MAIN_NET NetworkType = 77
	NETWORK_TYPE_TEST_NET NetworkType = 84
)

type AddressScheme uint16

const (
	ADDRESS_SCHEME_RESERVED AddressScheme = 0
	ADDRESS_SCHEME_EDDSA_20B AddressScheme = 1
	ADDRESS_SCHEME_SMART_CONTRACT AddressScheme = 2
)

