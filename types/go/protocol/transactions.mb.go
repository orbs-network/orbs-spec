// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message Transaction

// reader

type Transaction struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _Transaction_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeMessage,membuffers.TypeString,membuffers.TypeString,membuffers.TypeMessageArray,}
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

func (x *Transaction) ProtocolVersion() primitives.ProtocolVersion {
	return primitives.ProtocolVersion(x._message.GetUint32(0))
}

func (x *Transaction) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Transaction) MutateProtocolVersion(v primitives.ProtocolVersion) error {
	return x._message.SetUint32(0, uint32(v))
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

func (x *Transaction) Timestamp() primitives.Timestamp {
	return primitives.Timestamp(x._message.GetUint64(2))
}

func (x *Transaction) RawTimestamp() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *Transaction) MutateTimestamp(v primitives.Timestamp) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *Transaction) Signer() *Signer {
	b, s := x._message.GetMessage(3)
	return SignerReader(b[:s])
}

func (x *Transaction) RawSigner() []byte {
	return x._message.RawBufferForField(3, 0)
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

func (x *Transaction) MethodName() primitives.MethodName {
	return primitives.MethodName(x._message.GetString(5))
}

func (x *Transaction) RawMethodName() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *Transaction) MutateMethodName(v primitives.MethodName) error {
	return x._message.SetString(5, string(v))
}

func (x *Transaction) InputArgumentsIterator() *TransactionInputArgumentsIterator {
	return &TransactionInputArgumentsIterator{iterator: x._message.GetMessageArrayIterator(6)}
}

type TransactionInputArgumentsIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionInputArgumentsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionInputArgumentsIterator) NextInputArguments() *MethodArgument {
	b, s := i.iterator.NextMessage()
	return MethodArgumentReader(b[:s])
}

func (x *Transaction) RawInputArgumentsArray() []byte {
	return x._message.RawBufferForField(6, 0)
}

// builder

type TransactionBuilder struct {
	ProtocolVersion primitives.ProtocolVersion
	VirtualChainId primitives.VirtualChainId
	Timestamp primitives.Timestamp
	Signer *SignerBuilder
	ContractName primitives.ContractName
	MethodName primitives.MethodName
	InputArguments []*MethodArgumentBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *TransactionBuilder) arrayOfInputArguments() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.InputArguments))
	for i, v := range w.InputArguments {
		res[i] = v
	}
	return res
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
	err = w._builder.WriteMessageArray(buf, w.arrayOfInputArguments())
	if err != nil {
		return
	}
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
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _SignedTransaction_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
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

func (x *SignedTransaction) Transaction() *Transaction {
	b, s := x._message.GetMessage(0)
	return TransactionReader(b[:s])
}

func (x *SignedTransaction) RawTransaction() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SignedTransaction) Signature() []byte {
	return x._message.GetBytes(1)
}

func (x *SignedTransaction) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *SignedTransaction) MutateSignature(v []byte) error {
	return x._message.SetBytes(1, v)
}

// builder

type SignedTransactionBuilder struct {
	Transaction *TransactionBuilder
	Signature []byte

	// internal
	membuffers.Builder // interface
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
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _TransactionReceipt_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeUint16,membuffers.TypeMessageArray,}
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

func (x *TransactionReceipt) Txhash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(0))
}

func (x *TransactionReceipt) RawTxhash() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *TransactionReceipt) MutateTxhash(v primitives.Sha256) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *TransactionReceipt) ExecutionStatus() ExecutionStatus {
	return ExecutionStatus(x._message.GetUint16(1))
}

func (x *TransactionReceipt) RawExecutionStatus() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *TransactionReceipt) MutateExecutionStatus(v ExecutionStatus) error {
	return x._message.SetUint16(1, uint16(v))
}

func (x *TransactionReceipt) OutputArgumentsIterator() *TransactionReceiptOutputArgumentsIterator {
	return &TransactionReceiptOutputArgumentsIterator{iterator: x._message.GetMessageArrayIterator(2)}
}

type TransactionReceiptOutputArgumentsIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionReceiptOutputArgumentsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionReceiptOutputArgumentsIterator) NextOutputArguments() *MethodArgument {
	b, s := i.iterator.NextMessage()
	return MethodArgumentReader(b[:s])
}

func (x *TransactionReceipt) RawOutputArgumentsArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

// builder

type TransactionReceiptBuilder struct {
	Txhash primitives.Sha256
	ExecutionStatus ExecutionStatus
	OutputArguments []*MethodArgumentBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *TransactionReceiptBuilder) arrayOfOutputArguments() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.OutputArguments))
	for i, v := range w.OutputArguments {
		res[i] = v
	}
	return res
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
	w._builder.WriteUint16(buf, uint16(w.ExecutionStatus))
	err = w._builder.WriteMessageArray(buf, w.arrayOfOutputArguments())
	if err != nil {
		return
	}
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
// enums

