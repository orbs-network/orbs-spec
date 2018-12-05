// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package protocol

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message Transaction

// reader

type Transaction struct {
	// ProtocolVersion primitives.ProtocolVersion
	// VirtualChainId primitives.VirtualChainId
	// Timestamp primitives.TimestampNano
	// Signer Signer
	// ContractName primitives.ContractName
	// MethodName primitives.MethodName
	// InputArgumentArray []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *Transaction) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ProtocolVersion:%s,VirtualChainId:%s,Timestamp:%s,Signer:%s,ContractName:%s,MethodName:%s,InputArgumentArray:%s,}", x.StringProtocolVersion(), x.StringVirtualChainId(), x.StringTimestamp(), x.StringSigner(), x.StringContractName(), x.StringMethodName(), x.StringInputArgumentArray())
}

var _Transaction_Scheme = []membuffers.FieldType{membuffers.TypeUint32, membuffers.TypeUint32, membuffers.TypeUint64, membuffers.TypeMessage, membuffers.TypeString, membuffers.TypeString, membuffers.TypeBytes}
var _Transaction_Unions = [][]membuffers.FieldType{}

func TransactionReader(buf []byte) *Transaction {
	x := &Transaction{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _Transaction_Scheme, _Transaction_Unions)
	return x
}

func (x *Transaction) IsValid() bool {
	return x._message.IsValid()
}

func (x *Transaction) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *Transaction) Equal(y *Transaction) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *Transaction) ProtocolVersion() primitives.ProtocolVersion {
	return primitives.ProtocolVersion(x._message.GetUint32(0))
}

func (x *Transaction) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Transaction) MutateProtocolVersion(v primitives.ProtocolVersion) error {
	return x._message.SetUint32(0, uint32(v))
}

func (x *Transaction) StringProtocolVersion() string {
	return fmt.Sprintf("%s", x.ProtocolVersion())
}

func (x *Transaction) VirtualChainId() primitives.VirtualChainId {
	return primitives.VirtualChainId(x._message.GetUint32(1))
}

func (x *Transaction) RawVirtualChainId() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *Transaction) MutateVirtualChainId(v primitives.VirtualChainId) error {
	return x._message.SetUint32(1, uint32(v))
}

func (x *Transaction) StringVirtualChainId() string {
	return fmt.Sprintf("%s", x.VirtualChainId())
}

func (x *Transaction) Timestamp() primitives.TimestampNano {
	return primitives.TimestampNano(x._message.GetUint64(2))
}

func (x *Transaction) RawTimestamp() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *Transaction) MutateTimestamp(v primitives.TimestampNano) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *Transaction) StringTimestamp() string {
	return fmt.Sprintf("%s", x.Timestamp())
}

func (x *Transaction) Signer() *Signer {
	b, s := x._message.GetMessage(3)
	return SignerReader(b[:s])
}

func (x *Transaction) RawSigner() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *Transaction) RawSignerWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *Transaction) StringSigner() string {
	return x.Signer().String()
}

func (x *Transaction) ContractName() primitives.ContractName {
	return primitives.ContractName(x._message.GetString(4))
}

func (x *Transaction) RawContractName() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *Transaction) MutateContractName(v primitives.ContractName) error {
	return x._message.SetString(4, string(v))
}

func (x *Transaction) StringContractName() string {
	return fmt.Sprintf("%s", x.ContractName())
}

func (x *Transaction) MethodName() primitives.MethodName {
	return primitives.MethodName(x._message.GetString(5))
}

func (x *Transaction) RawMethodName() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *Transaction) MutateMethodName(v primitives.MethodName) error {
	return x._message.SetString(5, string(v))
}

func (x *Transaction) StringMethodName() string {
	return fmt.Sprintf("%s", x.MethodName())
}

func (x *Transaction) InputArgumentArray() []byte {
	return x._message.GetBytes(6)
}

func (x *Transaction) RawInputArgumentArray() []byte {
	return x._message.RawBufferForField(6, 0)
}

func (x *Transaction) RawInputArgumentArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(6, 0)
}

func (x *Transaction) MutateInputArgumentArray(v []byte) error {
	return x._message.SetBytes(6, v)
}

func (x *Transaction) StringInputArgumentArray() string {
	return fmt.Sprintf("%x", x.InputArgumentArray())
}

// builder

type TransactionBuilder struct {
	ProtocolVersion    primitives.ProtocolVersion
	VirtualChainId     primitives.VirtualChainId
	Timestamp          primitives.TimestampNano
	Signer             *SignerBuilder
	ContractName       primitives.ContractName
	MethodName         primitives.MethodName
	InputArgumentArray []byte

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *TransactionBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint32(buf, uint32(w.ProtocolVersion))
	w._builder.WriteUint32(buf, uint32(w.VirtualChainId))
	w._builder.WriteUint64(buf, uint64(w.Timestamp))
	err = w._builder.WriteMessage(buf, w.Signer)
	if err != nil {
		return
	}
	w._builder.WriteString(buf, string(w.ContractName))
	w._builder.WriteString(buf, string(w.MethodName))
	w._builder.WriteBytes(buf, w.InputArgumentArray)
	return nil
}

func (w *TransactionBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TransactionBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *TransactionBuilder) Build() *Transaction {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SignedTransaction

// reader

type SignedTransaction struct {
	// Transaction Transaction
	// Signature []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *SignedTransaction) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Transaction:%s,Signature:%s,}", x.StringTransaction(), x.StringSignature())
}

var _SignedTransaction_Scheme = []membuffers.FieldType{membuffers.TypeMessage, membuffers.TypeBytes}
var _SignedTransaction_Unions = [][]membuffers.FieldType{}

func SignedTransactionReader(buf []byte) *SignedTransaction {
	x := &SignedTransaction{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _SignedTransaction_Scheme, _SignedTransaction_Unions)
	return x
}

func (x *SignedTransaction) IsValid() bool {
	return x._message.IsValid()
}

func (x *SignedTransaction) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *SignedTransaction) Equal(y *SignedTransaction) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *SignedTransaction) Transaction() *Transaction {
	b, s := x._message.GetMessage(0)
	return TransactionReader(b[:s])
}

func (x *SignedTransaction) RawTransaction() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SignedTransaction) RawTransactionWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *SignedTransaction) StringTransaction() string {
	return x.Transaction().String()
}

func (x *SignedTransaction) Signature() []byte {
	return x._message.GetBytes(1)
}

func (x *SignedTransaction) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *SignedTransaction) RawSignatureWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *SignedTransaction) MutateSignature(v []byte) error {
	return x._message.SetBytes(1, v)
}

func (x *SignedTransaction) StringSignature() string {
	return fmt.Sprintf("%x", x.Signature())
}

// builder

type SignedTransactionBuilder struct {
	Transaction *TransactionBuilder
	Signature   []byte

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *SignedTransactionBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.Transaction)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *SignedTransactionBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *SignedTransactionBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *SignedTransactionBuilder) Build() *SignedTransaction {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SignedTransactionReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionReceipt

// reader

type TransactionReceipt struct {
	// Txhash primitives.Sha256
	// ExecutionResult ExecutionResult
	// OutputEvent []byte
	// OutputArgumentArray []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *TransactionReceipt) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Txhash:%s,ExecutionResult:%s,OutputEvent:%s,OutputArgumentArray:%s,}", x.StringTxhash(), x.StringExecutionResult(), x.StringOutputEvent(), x.StringOutputArgumentArray())
}

var _TransactionReceipt_Scheme = []membuffers.FieldType{membuffers.TypeBytes, membuffers.TypeUint16, membuffers.TypeBytes, membuffers.TypeBytes}
var _TransactionReceipt_Unions = [][]membuffers.FieldType{}

func TransactionReceiptReader(buf []byte) *TransactionReceipt {
	x := &TransactionReceipt{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _TransactionReceipt_Scheme, _TransactionReceipt_Unions)
	return x
}

func (x *TransactionReceipt) IsValid() bool {
	return x._message.IsValid()
}

func (x *TransactionReceipt) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *TransactionReceipt) Equal(y *TransactionReceipt) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *TransactionReceipt) Txhash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(0))
}

func (x *TransactionReceipt) RawTxhash() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *TransactionReceipt) MutateTxhash(v primitives.Sha256) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *TransactionReceipt) StringTxhash() string {
	return fmt.Sprintf("%s", x.Txhash())
}

func (x *TransactionReceipt) ExecutionResult() ExecutionResult {
	return ExecutionResult(x._message.GetUint16(1))
}

func (x *TransactionReceipt) RawExecutionResult() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *TransactionReceipt) MutateExecutionResult(v ExecutionResult) error {
	return x._message.SetUint16(1, uint16(v))
}

func (x *TransactionReceipt) StringExecutionResult() string {
	return x.ExecutionResult().String()
}

func (x *TransactionReceipt) OutputEvent() []byte {
	return x._message.GetBytes(2)
}

func (x *TransactionReceipt) RawOutputEvent() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *TransactionReceipt) RawOutputEventWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *TransactionReceipt) MutateOutputEvent(v []byte) error {
	return x._message.SetBytes(2, v)
}

func (x *TransactionReceipt) StringOutputEvent() string {
	return fmt.Sprintf("%x", x.OutputEvent())
}

func (x *TransactionReceipt) OutputArgumentArray() []byte {
	return x._message.GetBytes(3)
}

func (x *TransactionReceipt) RawOutputArgumentArray() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *TransactionReceipt) RawOutputArgumentArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *TransactionReceipt) MutateOutputArgumentArray(v []byte) error {
	return x._message.SetBytes(3, v)
}

func (x *TransactionReceipt) StringOutputArgumentArray() string {
	return fmt.Sprintf("%x", x.OutputArgumentArray())
}

// builder

type TransactionReceiptBuilder struct {
	Txhash              primitives.Sha256
	ExecutionResult     ExecutionResult
	OutputEvent         []byte
	OutputArgumentArray []byte

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *TransactionReceiptBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteBytes(buf, []byte(w.Txhash))
	w._builder.WriteUint16(buf, uint16(w.ExecutionResult))
	w._builder.WriteBytes(buf, w.OutputEvent)
	w._builder.WriteBytes(buf, w.OutputArgumentArray)
	return nil
}

func (w *TransactionReceiptBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TransactionReceiptBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *TransactionReceiptBuilder) Build() *TransactionReceipt {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionReceiptReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OutputEvent

// reader

type OutputEvent struct {
	// ContractName primitives.ContractName
	// Arguments []EventArgument

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *OutputEvent) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContractName:%s,Arguments:%s,}", x.StringContractName(), x.StringArguments())
}

var _OutputEvent_Scheme = []membuffers.FieldType{membuffers.TypeString, membuffers.TypeMessageArray}
var _OutputEvent_Unions = [][]membuffers.FieldType{}

func OutputEventReader(buf []byte) *OutputEvent {
	x := &OutputEvent{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _OutputEvent_Scheme, _OutputEvent_Unions)
	return x
}

func (x *OutputEvent) IsValid() bool {
	return x._message.IsValid()
}

func (x *OutputEvent) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *OutputEvent) Equal(y *OutputEvent) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *OutputEvent) ContractName() primitives.ContractName {
	return primitives.ContractName(x._message.GetString(0))
}

func (x *OutputEvent) RawContractName() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *OutputEvent) MutateContractName(v primitives.ContractName) error {
	return x._message.SetString(0, string(v))
}

func (x *OutputEvent) StringContractName() string {
	return fmt.Sprintf("%s", x.ContractName())
}

func (x *OutputEvent) ArgumentsIterator() *OutputEventArgumentsIterator {
	return &OutputEventArgumentsIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type OutputEventArgumentsIterator struct {
	iterator *membuffers.Iterator
}

func (i *OutputEventArgumentsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *OutputEventArgumentsIterator) NextArguments() *EventArgument {
	b, s := i.iterator.NextMessage()
	return EventArgumentReader(b[:s])
}

func (x *OutputEvent) RawArgumentsArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *OutputEvent) RawArgumentsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *OutputEvent) StringArguments() (res string) {
	res = "["
	for i := x.ArgumentsIterator(); i.HasNext(); {
		res += i.NextArguments().String() + ","
	}
	res += "]"
	return
}

// builder

type OutputEventBuilder struct {
	ContractName primitives.ContractName
	Arguments    []*EventArgumentBuilder

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *OutputEventBuilder) arrayOfArguments() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Arguments))
	for i, v := range w.Arguments {
		res[i] = v
	}
	return res
}

func (w *OutputEventBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteString(buf, string(w.ContractName))
	err = w._builder.WriteMessageArray(buf, w.arrayOfArguments())
	if err != nil {
		return
	}
	return nil
}

func (w *OutputEventBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *OutputEventBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *OutputEventBuilder) Build() *OutputEvent {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OutputEventReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message EventArgument

// reader

type EventArgument struct {
	// Type EventArgumentType

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *EventArgument) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Type:%s,}", x.StringType())
}

var _EventArgument_Scheme = []membuffers.FieldType{membuffers.TypeUnion}
var _EventArgument_Unions = [][]membuffers.FieldType{{membuffers.TypeUint32, membuffers.TypeUint64, membuffers.TypeString, membuffers.TypeBytes}}

func EventArgumentReader(buf []byte) *EventArgument {
	x := &EventArgument{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _EventArgument_Scheme, _EventArgument_Unions)
	return x
}

func (x *EventArgument) IsValid() bool {
	return x._message.IsValid()
}

func (x *EventArgument) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *EventArgument) Equal(y *EventArgument) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

type EventArgumentType uint16

const (
	EVENT_ARGUMENT_TYPE_UINT_32_VALUE EventArgumentType = 0
	EVENT_ARGUMENT_TYPE_UINT_64_VALUE EventArgumentType = 1
	EVENT_ARGUMENT_TYPE_STRING_VALUE  EventArgumentType = 2
	EVENT_ARGUMENT_TYPE_BYTES_VALUE   EventArgumentType = 3
)

func (x *EventArgument) Type() EventArgumentType {
	return EventArgumentType(x._message.GetUnionIndex(0, 0))
}

func (x *EventArgument) IsTypeUint32Value() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 0)
	return is
}

func (x *EventArgument) Uint32Value() uint32 {
	is, off := x._message.IsUnionIndex(0, 0, 0)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetUint32InOffset(off)
}

func (x *EventArgument) StringUint32Value() string {
	return fmt.Sprintf("%x", x.Uint32Value())
}

func (x *EventArgument) MutateUint32Value(v uint32) error {
	is, off := x._message.IsUnionIndex(0, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint32InOffset(off, v)
	return nil
}

func (x *EventArgument) IsTypeUint64Value() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 1)
	return is
}

func (x *EventArgument) Uint64Value() uint64 {
	is, off := x._message.IsUnionIndex(0, 0, 1)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetUint64InOffset(off)
}

func (x *EventArgument) StringUint64Value() string {
	return fmt.Sprintf("%x", x.Uint64Value())
}

func (x *EventArgument) MutateUint64Value(v uint64) error {
	is, off := x._message.IsUnionIndex(0, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint64InOffset(off, v)
	return nil
}

func (x *EventArgument) IsTypeStringValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 2)
	return is
}

func (x *EventArgument) StringValue() string {
	is, off := x._message.IsUnionIndex(0, 0, 2)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetStringInOffset(off)
}

func (x *EventArgument) StringStringValue() string {
	return fmt.Sprintf(x.StringValue())
}

func (x *EventArgument) MutateStringValue(v string) error {
	is, off := x._message.IsUnionIndex(0, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetStringInOffset(off, v)
	return nil
}

func (x *EventArgument) IsTypeBytesValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 3)
	return is
}

func (x *EventArgument) BytesValue() []byte {
	is, off := x._message.IsUnionIndex(0, 0, 3)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetBytesInOffset(off)
}

func (x *EventArgument) StringBytesValue() string {
	return fmt.Sprintf("%x", x.BytesValue())
}

func (x *EventArgument) MutateBytesValue(v []byte) error {
	is, off := x._message.IsUnionIndex(0, 0, 3)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetBytesInOffset(off, v)
	return nil
}

func (x *EventArgument) RawType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *EventArgument) RawTypeWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *EventArgument) StringType() string {
	switch x.Type() {
	case EVENT_ARGUMENT_TYPE_UINT_32_VALUE:
		return "(Uint32Value)" + x.StringUint32Value()
	case EVENT_ARGUMENT_TYPE_UINT_64_VALUE:
		return "(Uint64Value)" + x.StringUint64Value()
	case EVENT_ARGUMENT_TYPE_STRING_VALUE:
		return "(StringValue)" + x.StringStringValue()
	case EVENT_ARGUMENT_TYPE_BYTES_VALUE:
		return "(BytesValue)" + x.StringBytesValue()
	}
	return "(Unknown)"
}

// builder

type EventArgumentBuilder struct {
	Type        EventArgumentType
	Uint32Value uint32
	Uint64Value uint64
	StringValue string
	BytesValue  []byte

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *EventArgumentBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case EVENT_ARGUMENT_TYPE_UINT_32_VALUE:
		w._builder.WriteUint32(buf, w.Uint32Value)
	case EVENT_ARGUMENT_TYPE_UINT_64_VALUE:
		w._builder.WriteUint64(buf, w.Uint64Value)
	case EVENT_ARGUMENT_TYPE_STRING_VALUE:
		w._builder.WriteString(buf, w.StringValue)
	case EVENT_ARGUMENT_TYPE_BYTES_VALUE:
		w._builder.WriteBytes(buf, w.BytesValue)
	}
	return nil
}

func (w *EventArgumentBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *EventArgumentBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *EventArgumentBuilder) Build() *EventArgument {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return EventArgumentReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums
