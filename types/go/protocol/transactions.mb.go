// AUTO GENERATED FILE (by membufc proto compiler v0.0.32)
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
	// InputArgumentArray primitives.PackedArgumentArray

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

func (x *Transaction) RawContractNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(4, 0)
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

func (x *Transaction) RawMethodNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(5, 0)
}

func (x *Transaction) MutateMethodName(v primitives.MethodName) error {
	return x._message.SetString(5, string(v))
}

func (x *Transaction) StringMethodName() string {
	return fmt.Sprintf("%s", x.MethodName())
}

func (x *Transaction) InputArgumentArray() primitives.PackedArgumentArray {
	return primitives.PackedArgumentArray(x._message.GetBytes(6))
}

func (x *Transaction) RawInputArgumentArray() []byte {
	return x._message.RawBufferForField(6, 0)
}

func (x *Transaction) RawInputArgumentArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(6, 0)
}

func (x *Transaction) MutateInputArgumentArray(v primitives.PackedArgumentArray) error {
	return x._message.SetBytes(6, []byte(v))
}

func (x *Transaction) StringInputArgumentArray() string {
	return fmt.Sprintf("%s", x.InputArgumentArray())
}

// builder

type TransactionBuilder struct {
	ProtocolVersion    primitives.ProtocolVersion
	VirtualChainId     primitives.VirtualChainId
	Timestamp          primitives.TimestampNano
	Signer             *SignerBuilder
	ContractName       primitives.ContractName
	MethodName         primitives.MethodName
	InputArgumentArray primitives.PackedArgumentArray

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *TransactionBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUint32(buf, uint32(w.ProtocolVersion))
	w._builder.WriteUint32(buf, uint32(w.VirtualChainId))
	w._builder.WriteUint64(buf, uint64(w.Timestamp))
	err = w._builder.WriteMessage(buf, w.Signer)
	if err != nil {
		return
	}
	w._builder.WriteString(buf, string(w.ContractName))
	w._builder.WriteString(buf, string(w.MethodName))
	w._builder.WriteBytes(buf, []byte(w.InputArgumentArray))
	return nil
}

func (w *TransactionBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint32(prefix, offsetFromStart, "Transaction.ProtocolVersion", uint32(w.ProtocolVersion))
	w._builder.HexDumpUint32(prefix, offsetFromStart, "Transaction.VirtualChainId", uint32(w.VirtualChainId))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "Transaction.Timestamp", uint64(w.Timestamp))
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "Transaction.Signer", w.Signer)
	if err != nil {
		return
	}
	w._builder.HexDumpString(prefix, offsetFromStart, "Transaction.ContractName", string(w.ContractName))
	w._builder.HexDumpString(prefix, offsetFromStart, "Transaction.MethodName", string(w.MethodName))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "Transaction.InputArgumentArray", []byte(w.InputArgumentArray))
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

func TransactionBuilderFromRaw(raw []byte) *TransactionBuilder {
	return &TransactionBuilder{_overrideWithRawBuffer: raw}
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
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *SignedTransactionBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessage(buf, w.Transaction)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *SignedTransactionBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "SignedTransaction.Transaction", w.Transaction)
	if err != nil {
		return
	}
	w._builder.HexDumpBytes(prefix, offsetFromStart, "SignedTransaction.Signature", w.Signature)
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

func SignedTransactionBuilderFromRaw(raw []byte) *SignedTransactionBuilder {
	return &SignedTransactionBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionReceipt

// reader

type TransactionReceipt struct {
	// Txhash primitives.Sha256
	// ExecutionResult ExecutionResult
	// OutputArgumentArray primitives.PackedArgumentArray
	// OutputEventsArray primitives.PackedEventsArray

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *TransactionReceipt) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Txhash:%s,ExecutionResult:%s,OutputArgumentArray:%s,OutputEventsArray:%s,}", x.StringTxhash(), x.StringExecutionResult(), x.StringOutputArgumentArray(), x.StringOutputEventsArray())
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

func (x *TransactionReceipt) RawTxhashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
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

func (x *TransactionReceipt) OutputArgumentArray() primitives.PackedArgumentArray {
	return primitives.PackedArgumentArray(x._message.GetBytes(2))
}

func (x *TransactionReceipt) RawOutputArgumentArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *TransactionReceipt) RawOutputArgumentArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *TransactionReceipt) MutateOutputArgumentArray(v primitives.PackedArgumentArray) error {
	return x._message.SetBytes(2, []byte(v))
}

func (x *TransactionReceipt) StringOutputArgumentArray() string {
	return fmt.Sprintf("%s", x.OutputArgumentArray())
}

func (x *TransactionReceipt) OutputEventsArray() primitives.PackedEventsArray {
	return primitives.PackedEventsArray(x._message.GetBytes(3))
}

func (x *TransactionReceipt) RawOutputEventsArray() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *TransactionReceipt) RawOutputEventsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *TransactionReceipt) MutateOutputEventsArray(v primitives.PackedEventsArray) error {
	return x._message.SetBytes(3, []byte(v))
}

func (x *TransactionReceipt) StringOutputEventsArray() string {
	return fmt.Sprintf("%s", x.OutputEventsArray())
}

// builder

type TransactionReceiptBuilder struct {
	Txhash              primitives.Sha256
	ExecutionResult     ExecutionResult
	OutputArgumentArray primitives.PackedArgumentArray
	OutputEventsArray   primitives.PackedEventsArray

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *TransactionReceiptBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, []byte(w.Txhash))
	w._builder.WriteUint16(buf, uint16(w.ExecutionResult))
	w._builder.WriteBytes(buf, []byte(w.OutputArgumentArray))
	w._builder.WriteBytes(buf, []byte(w.OutputEventsArray))
	return nil
}

func (w *TransactionReceiptBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "TransactionReceipt.Txhash", []byte(w.Txhash))
	w._builder.HexDumpUint16(prefix, offsetFromStart, "TransactionReceipt.ExecutionResult", uint16(w.ExecutionResult))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "TransactionReceipt.OutputArgumentArray", []byte(w.OutputArgumentArray))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "TransactionReceipt.OutputEventsArray", []byte(w.OutputEventsArray))
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

func TransactionReceiptBuilderFromRaw(raw []byte) *TransactionReceiptBuilder {
	return &TransactionReceiptBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// enums
