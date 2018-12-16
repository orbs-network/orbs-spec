// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package protocol

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message Signer

// reader

type Signer struct {
	// Scheme SignerScheme

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *Signer) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Scheme:%s,}", x.StringScheme())
}

var _Signer_Scheme = []membuffers.FieldType{membuffers.TypeUnion}
var _Signer_Unions = [][]membuffers.FieldType{{membuffers.TypeMessage, membuffers.TypeMessage}}

func SignerReader(buf []byte) *Signer {
	x := &Signer{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _Signer_Scheme, _Signer_Unions)
	return x
}

func (x *Signer) IsValid() bool {
	return x._message.IsValid()
}

func (x *Signer) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *Signer) Equal(y *Signer) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

type SignerScheme uint16

const (
	SIGNER_SCHEME_EDDSA                 SignerScheme = 0
	SIGNER_SCHEME_SMART_CONTRACT_CALLER SignerScheme = 1
)

func (x *Signer) Scheme() SignerScheme {
	return SignerScheme(x._message.GetUnionIndex(0, 0))
}

func (x *Signer) IsSchemeEddsa() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 0)
	return is
}

func (x *Signer) Eddsa() *EdDSA01Signer {
	is, off := x._message.IsUnionIndex(0, 0, 0)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	b, s := x._message.GetMessageInOffset(off)
	return EdDSA01SignerReader(b[:s])
}

func (x *Signer) StringEddsa() string {
	return x.Eddsa().String()
}

func (x *Signer) IsSchemeSmartContractCaller() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 1)
	return is
}

func (x *Signer) SmartContractCaller() *SmartContractCaller {
	is, off := x._message.IsUnionIndex(0, 0, 1)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	b, s := x._message.GetMessageInOffset(off)
	return SmartContractCallerReader(b[:s])
}

func (x *Signer) StringSmartContractCaller() string {
	return x.SmartContractCaller().String()
}

func (x *Signer) RawScheme() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Signer) RawSchemeWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *Signer) StringScheme() string {
	switch x.Scheme() {
	case SIGNER_SCHEME_EDDSA:
		return "(Eddsa)" + x.StringEddsa()
	case SIGNER_SCHEME_SMART_CONTRACT_CALLER:
		return "(SmartContractCaller)" + x.StringSmartContractCaller()
	}
	return "(Unknown)"
}

// builder

type SignerBuilder struct {
	Scheme              SignerScheme
	Eddsa               *EdDSA01SignerBuilder
	SmartContractCaller *SmartContractCallerBuilder

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *SignerBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteUnionIndex(buf, uint16(w.Scheme))
	switch w.Scheme {
	case SIGNER_SCHEME_EDDSA:
		w._builder.WriteMessage(buf, w.Eddsa)
	case SIGNER_SCHEME_SMART_CONTRACT_CALLER:
		w._builder.WriteMessage(buf, w.SmartContractCaller)
	}
	return nil
}

func (w *SignerBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUnionIndex(prefix, offsetFromStart, "Signer.Scheme", uint16(w.Scheme))
	switch w.Scheme {
	case SIGNER_SCHEME_EDDSA:
		w._builder.HexDumpMessage(prefix, offsetFromStart, "Signer.Eddsa", w.Eddsa)
	case SIGNER_SCHEME_SMART_CONTRACT_CALLER:
		w._builder.HexDumpMessage(prefix, offsetFromStart, "Signer.SmartContractCaller", w.SmartContractCaller)
	}
	return nil
}

func (w *SignerBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *SignerBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *SignerBuilder) Build() *Signer {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SignerReader(buf)
}

func SignerBuilderFromRaw(raw []byte) *SignerBuilder {
	return &SignerBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message EdDSA01Signer

// reader

type EdDSA01Signer struct {
	// NetworkType SignerNetworkType
	// SignerPublicKey primitives.Ed25519PublicKey

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *EdDSA01Signer) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{NetworkType:%s,SignerPublicKey:%s,}", x.StringNetworkType(), x.StringSignerPublicKey())
}

var _EdDSA01Signer_Scheme = []membuffers.FieldType{membuffers.TypeUint16, membuffers.TypeBytes}
var _EdDSA01Signer_Unions = [][]membuffers.FieldType{}

func EdDSA01SignerReader(buf []byte) *EdDSA01Signer {
	x := &EdDSA01Signer{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _EdDSA01Signer_Scheme, _EdDSA01Signer_Unions)
	return x
}

func (x *EdDSA01Signer) IsValid() bool {
	return x._message.IsValid()
}

func (x *EdDSA01Signer) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *EdDSA01Signer) Equal(y *EdDSA01Signer) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *EdDSA01Signer) NetworkType() SignerNetworkType {
	return SignerNetworkType(x._message.GetUint16(0))
}

func (x *EdDSA01Signer) RawNetworkType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *EdDSA01Signer) MutateNetworkType(v SignerNetworkType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *EdDSA01Signer) StringNetworkType() string {
	return x.NetworkType().String()
}

func (x *EdDSA01Signer) SignerPublicKey() primitives.Ed25519PublicKey {
	return primitives.Ed25519PublicKey(x._message.GetBytes(1))
}

func (x *EdDSA01Signer) RawSignerPublicKey() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *EdDSA01Signer) RawSignerPublicKeyWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *EdDSA01Signer) MutateSignerPublicKey(v primitives.Ed25519PublicKey) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *EdDSA01Signer) StringSignerPublicKey() string {
	return fmt.Sprintf("%s", x.SignerPublicKey())
}

// builder

type EdDSA01SignerBuilder struct {
	NetworkType     SignerNetworkType
	SignerPublicKey primitives.Ed25519PublicKey

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *EdDSA01SignerBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteUint16(buf, uint16(w.NetworkType))
	w._builder.WriteBytes(buf, []byte(w.SignerPublicKey))
	return nil
}

func (w *EdDSA01SignerBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint16(prefix, offsetFromStart, "EdDSA01Signer.NetworkType", uint16(w.NetworkType))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "EdDSA01Signer.SignerPublicKey", []byte(w.SignerPublicKey))
	return nil
}

func (w *EdDSA01SignerBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *EdDSA01SignerBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *EdDSA01SignerBuilder) Build() *EdDSA01Signer {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return EdDSA01SignerReader(buf)
}

func EdDSA01SignerBuilderFromRaw(raw []byte) *EdDSA01SignerBuilder {
	return &EdDSA01SignerBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message SmartContractCaller

// reader

type SmartContractCaller struct {
	// ContractName primitives.ContractName

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *SmartContractCaller) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContractName:%s,}", x.StringContractName())
}

var _SmartContractCaller_Scheme = []membuffers.FieldType{membuffers.TypeString}
var _SmartContractCaller_Unions = [][]membuffers.FieldType{}

func SmartContractCallerReader(buf []byte) *SmartContractCaller {
	x := &SmartContractCaller{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _SmartContractCaller_Scheme, _SmartContractCaller_Unions)
	return x
}

func (x *SmartContractCaller) IsValid() bool {
	return x._message.IsValid()
}

func (x *SmartContractCaller) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *SmartContractCaller) Equal(y *SmartContractCaller) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *SmartContractCaller) ContractName() primitives.ContractName {
	return primitives.ContractName(x._message.GetString(0))
}

func (x *SmartContractCaller) RawContractName() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SmartContractCaller) RawContractNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *SmartContractCaller) MutateContractName(v primitives.ContractName) error {
	return x._message.SetString(0, string(v))
}

func (x *SmartContractCaller) StringContractName() string {
	return fmt.Sprintf("%s", x.ContractName())
}

// builder

type SmartContractCallerBuilder struct {
	ContractName primitives.ContractName

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *SmartContractCallerBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteString(buf, string(w.ContractName))
	return nil
}

func (w *SmartContractCallerBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpString(prefix, offsetFromStart, "SmartContractCaller.ContractName", string(w.ContractName))
	return nil
}

func (w *SmartContractCallerBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *SmartContractCallerBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *SmartContractCallerBuilder) Build() *SmartContractCaller {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SmartContractCallerReader(buf)
}

func SmartContractCallerBuilderFromRaw(raw []byte) *SmartContractCallerBuilder {
	return &SmartContractCallerBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// enums

type SignerNetworkType uint16

const (
	NETWORK_TYPE_RESERVED SignerNetworkType = 0
	NETWORK_TYPE_MAIN_NET SignerNetworkType = 77
	NETWORK_TYPE_TEST_NET SignerNetworkType = 84
)

func (n SignerNetworkType) String() string {
	switch n {
	case NETWORK_TYPE_RESERVED:
		return "NETWORK_TYPE_RESERVED"
	case NETWORK_TYPE_MAIN_NET:
		return "NETWORK_TYPE_MAIN_NET"
	case NETWORK_TYPE_TEST_NET:
		return "NETWORK_TYPE_TEST_NET"
	}
	return "UNKNOWN"
}
