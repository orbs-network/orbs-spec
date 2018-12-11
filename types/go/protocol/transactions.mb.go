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
	// OutputEventsArray []byte
	// OutputArgumentArray []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *TransactionReceipt) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Txhash:%s,ExecutionResult:%s,OutputEventsArray:%s,OutputArgumentArray:%s,}", x.StringTxhash(), x.StringExecutionResult(), x.StringOutputEventsArray(), x.StringOutputArgumentArray())
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

func (x *TransactionReceipt) OutputEventsArray() []byte {
	return x._message.GetBytes(2)
}

func (x *TransactionReceipt) RawOutputEventsArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *TransactionReceipt) RawOutputEventsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *TransactionReceipt) MutateOutputEventsArray(v []byte) error {
	return x._message.SetBytes(2, v)
}

func (x *TransactionReceipt) StringOutputEventsArray() string {
	return fmt.Sprintf("%x", x.OutputEventsArray())
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
	OutputEventsArray   []byte
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
	w._builder.WriteBytes(buf, w.OutputEventsArray)
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
// message Query

// reader

type Query struct {
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

func (x *Query) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ProtocolVersion:%s,VirtualChainId:%s,Timestamp:%s,Signer:%s,ContractName:%s,MethodName:%s,InputArgumentArray:%s,}", x.StringProtocolVersion(), x.StringVirtualChainId(), x.StringTimestamp(), x.StringSigner(), x.StringContractName(), x.StringMethodName(), x.StringInputArgumentArray())
}

var _Query_Scheme = []membuffers.FieldType{membuffers.TypeUint32, membuffers.TypeUint32, membuffers.TypeUint64, membuffers.TypeMessage, membuffers.TypeString, membuffers.TypeString, membuffers.TypeBytes}
var _Query_Unions = [][]membuffers.FieldType{}

func QueryReader(buf []byte) *Query {
	x := &Query{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _Query_Scheme, _Query_Unions)
	return x
}

func (x *Query) IsValid() bool {
	return x._message.IsValid()
}

func (x *Query) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *Query) Equal(y *Query) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *Query) ProtocolVersion() primitives.ProtocolVersion {
	return primitives.ProtocolVersion(x._message.GetUint32(0))
}

func (x *Query) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Query) MutateProtocolVersion(v primitives.ProtocolVersion) error {
	return x._message.SetUint32(0, uint32(v))
}

func (x *Query) StringProtocolVersion() string {
	return fmt.Sprintf("%s", x.ProtocolVersion())
}

func (x *Query) VirtualChainId() primitives.VirtualChainId {
	return primitives.VirtualChainId(x._message.GetUint32(1))
}

func (x *Query) RawVirtualChainId() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *Query) MutateVirtualChainId(v primitives.VirtualChainId) error {
	return x._message.SetUint32(1, uint32(v))
}

func (x *Query) StringVirtualChainId() string {
	return fmt.Sprintf("%s", x.VirtualChainId())
}

func (x *Query) Timestamp() primitives.TimestampNano {
	return primitives.TimestampNano(x._message.GetUint64(2))
}

func (x *Query) RawTimestamp() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *Query) MutateTimestamp(v primitives.TimestampNano) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *Query) StringTimestamp() string {
	return fmt.Sprintf("%s", x.Timestamp())
}

func (x *Query) Signer() *Signer {
	b, s := x._message.GetMessage(3)
	return SignerReader(b[:s])
}

func (x *Query) RawSigner() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *Query) RawSignerWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *Query) StringSigner() string {
	return x.Signer().String()
}

func (x *Query) ContractName() primitives.ContractName {
	return primitives.ContractName(x._message.GetString(4))
}

func (x *Query) RawContractName() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *Query) MutateContractName(v primitives.ContractName) error {
	return x._message.SetString(4, string(v))
}

func (x *Query) StringContractName() string {
	return fmt.Sprintf("%s", x.ContractName())
}

func (x *Query) MethodName() primitives.MethodName {
	return primitives.MethodName(x._message.GetString(5))
}

func (x *Query) RawMethodName() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *Query) MutateMethodName(v primitives.MethodName) error {
	return x._message.SetString(5, string(v))
}

func (x *Query) StringMethodName() string {
	return fmt.Sprintf("%s", x.MethodName())
}

func (x *Query) InputArgumentArray() []byte {
	return x._message.GetBytes(6)
}

func (x *Query) RawInputArgumentArray() []byte {
	return x._message.RawBufferForField(6, 0)
}

func (x *Query) RawInputArgumentArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(6, 0)
}

func (x *Query) MutateInputArgumentArray(v []byte) error {
	return x._message.SetBytes(6, v)
}

func (x *Query) StringInputArgumentArray() string {
	return fmt.Sprintf("%x", x.InputArgumentArray())
}

// builder

type QueryBuilder struct {
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

func (w *QueryBuilder) Write(buf []byte) (err error) {
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

func (w *QueryBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *QueryBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *QueryBuilder) Build() *Query {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return QueryReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SignedQuery

// reader

type SignedQuery struct {
	// Transaction Transaction
	// Signature []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *SignedQuery) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Transaction:%s,Signature:%s,}", x.StringTransaction(), x.StringSignature())
}

var _SignedQuery_Scheme = []membuffers.FieldType{membuffers.TypeMessage, membuffers.TypeBytes}
var _SignedQuery_Unions = [][]membuffers.FieldType{}

func SignedQueryReader(buf []byte) *SignedQuery {
	x := &SignedQuery{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _SignedQuery_Scheme, _SignedQuery_Unions)
	return x
}

func (x *SignedQuery) IsValid() bool {
	return x._message.IsValid()
}

func (x *SignedQuery) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *SignedQuery) Equal(y *SignedQuery) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *SignedQuery) Transaction() *Transaction {
	b, s := x._message.GetMessage(0)
	return TransactionReader(b[:s])
}

func (x *SignedQuery) RawTransaction() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SignedQuery) RawTransactionWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *SignedQuery) StringTransaction() string {
	return x.Transaction().String()
}

func (x *SignedQuery) Signature() []byte {
	return x._message.GetBytes(1)
}

func (x *SignedQuery) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *SignedQuery) RawSignatureWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *SignedQuery) MutateSignature(v []byte) error {
	return x._message.SetBytes(1, v)
}

func (x *SignedQuery) StringSignature() string {
	return fmt.Sprintf("%x", x.Signature())
}

// builder

type SignedQueryBuilder struct {
	Transaction *TransactionBuilder
	Signature   []byte

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *SignedQueryBuilder) Write(buf []byte) (err error) {
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

func (w *SignedQueryBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *SignedQueryBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *SignedQueryBuilder) Build() *SignedQuery {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SignedQueryReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message QueryResult

// reader

type QueryResult struct {
	// ExecutionResult ExecutionResult
	// OutputEventsArray []byte
	// OutputArgumentArray []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *QueryResult) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ExecutionResult:%s,OutputEventsArray:%s,OutputArgumentArray:%s,}", x.StringExecutionResult(), x.StringOutputEventsArray(), x.StringOutputArgumentArray())
}

var _QueryResult_Scheme = []membuffers.FieldType{membuffers.TypeUint16, membuffers.TypeBytes, membuffers.TypeBytes}
var _QueryResult_Unions = [][]membuffers.FieldType{}

func QueryResultReader(buf []byte) *QueryResult {
	x := &QueryResult{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _QueryResult_Scheme, _QueryResult_Unions)
	return x
}

func (x *QueryResult) IsValid() bool {
	return x._message.IsValid()
}

func (x *QueryResult) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *QueryResult) Equal(y *QueryResult) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *QueryResult) ExecutionResult() ExecutionResult {
	return ExecutionResult(x._message.GetUint16(0))
}

func (x *QueryResult) RawExecutionResult() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *QueryResult) MutateExecutionResult(v ExecutionResult) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *QueryResult) StringExecutionResult() string {
	return x.ExecutionResult().String()
}

func (x *QueryResult) OutputEventsArray() []byte {
	return x._message.GetBytes(1)
}

func (x *QueryResult) RawOutputEventsArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *QueryResult) RawOutputEventsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *QueryResult) MutateOutputEventsArray(v []byte) error {
	return x._message.SetBytes(1, v)
}

func (x *QueryResult) StringOutputEventsArray() string {
	return fmt.Sprintf("%x", x.OutputEventsArray())
}

func (x *QueryResult) OutputArgumentArray() []byte {
	return x._message.GetBytes(2)
}

func (x *QueryResult) RawOutputArgumentArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *QueryResult) RawOutputArgumentArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *QueryResult) MutateOutputArgumentArray(v []byte) error {
	return x._message.SetBytes(2, v)
}

func (x *QueryResult) StringOutputArgumentArray() string {
	return fmt.Sprintf("%x", x.OutputArgumentArray())
}

// builder

type QueryResultBuilder struct {
	ExecutionResult     ExecutionResult
	OutputEventsArray   []byte
	OutputArgumentArray []byte

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *QueryResultBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint16(buf, uint16(w.ExecutionResult))
	w._builder.WriteBytes(buf, w.OutputEventsArray)
	w._builder.WriteBytes(buf, w.OutputArgumentArray)
	return nil
}

func (w *QueryResultBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *QueryResultBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *QueryResultBuilder) Build() *QueryResult {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return QueryResultReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums
