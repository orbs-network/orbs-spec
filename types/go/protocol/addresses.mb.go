// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message Sender

// reader

type Sender struct {
	message membuffers.Message
}

var _Sender_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeUnion,}
var _Sender_Unions = [][]membuffers.FieldType{{membuffers.TypeMessage,membuffers.TypeMessage,}}

func SenderReader(buf []byte) *Sender {
	x := &Sender{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _Sender_Scheme, _Sender_Unions)
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

type SenderScheme uint16

const (
	SenderSchemeEddsa SenderScheme = 0
	SenderSchemeSmartContract SenderScheme = 1
)

func (x *Sender) Scheme() SenderScheme {
	return SenderScheme(x.message.GetUint16(1))
}

func (x *Sender) IsSchemeEddsa() bool {
	is, _ := x.message.IsUnionIndex(1, 0, 0)
	return is
}

func (x *Sender) SchemeEddsa() *EdDSA01Sender {
	_, off := x.message.IsUnionIndex(1, 0, 0)
	b, s := x.message.GetMessageInOffset(off)
	return EdDSA01SenderReader(b[:s])
}

func (x *Sender) IsSchemeSmartContract() bool {
	is, _ := x.message.IsUnionIndex(1, 0, 1)
	return is
}

func (x *Sender) SchemeSmartContract() *SmartContractSender {
	_, off := x.message.IsUnionIndex(1, 0, 1)
	b, s := x.message.GetMessageInOffset(off)
	return SmartContractSenderReader(b[:s])
}

func (x *Sender) RawScheme() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type SenderBuilder struct {
	builder membuffers.Builder
	AddressScheme AddressScheme
	Scheme SenderScheme
	Eddsa *EdDSA01SenderBuilder
	SmartContract *SmartContractSenderBuilder
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
	w.builder.WriteUnionIndex(buf, uint16(w.Scheme))
	switch w.Scheme {
	case SenderSchemeEddsa:
		w.builder.WriteMessage(buf, w.Eddsa)
	case SenderSchemeSmartContract:
		w.builder.WriteMessage(buf, w.SmartContract)
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

var _EdDSA01Sender_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeBytes,}
var _EdDSA01Sender_Unions = [][]membuffers.FieldType{}

func EdDSA01SenderReader(buf []byte) *EdDSA01Sender {
	x := &EdDSA01Sender{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _EdDSA01Sender_Scheme, _EdDSA01Sender_Unions)
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

func (x *EdDSA01Sender) SenderPublicKey() primitives.Ed25519Pkey {
	return x.message.GetBytes(1)
}

func (x *EdDSA01Sender) RawSenderPublicKey() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *EdDSA01Sender) MutateSenderPublicKey(v primitives.Ed25519Pkey) error {
	return x.message.SetBytes(1, v)
}

// builder

type EdDSA01SenderBuilder struct {
	builder membuffers.Builder
	NetworkType NetworkType
	SenderPublicKey primitives.Ed25519Pkey
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

var _SmartContractSender_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeString,membuffers.TypeMessageArray,}
var _SmartContractSender_Unions = [][]membuffers.FieldType{}

func SmartContractSenderReader(buf []byte) *SmartContractSender {
	x := &SmartContractSender{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _SmartContractSender_Scheme, _SmartContractSender_Unions)
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

