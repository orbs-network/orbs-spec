// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
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

var _Transaction_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeMessage,membuffers.TypeUint64,membuffers.TypeString,membuffers.TypeString,membuffers.TypeMessageArray,}
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

func (x *Transaction) ProtocolVersion() uint32 {
	return x._message.GetUint32(0)
}

func (x *Transaction) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Transaction) MutateProtocolVersion(v uint32) error {
	return x._message.SetUint32(0, v)
}

func (x *Transaction) VirtualChain() uint64 {
	return x._message.GetUint64(1)
}

func (x *Transaction) RawVirtualChain() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *Transaction) MutateVirtualChain(v uint64) error {
	return x._message.SetUint64(1, v)
}

func (x *Transaction) Sender() *Sender {
	b, s := x._message.GetMessage(2)
	return SenderReader(b[:s])
}

func (x *Transaction) RawSender() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *Transaction) Timestamp() uint64 {
	return x._message.GetUint64(3)
}

func (x *Transaction) RawTimestamp() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *Transaction) MutateTimestamp(v uint64) error {
	return x._message.SetUint64(3, v)
}

func (x *Transaction) ContractName() string {
	return x._message.GetString(4)
}

func (x *Transaction) RawContractName() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *Transaction) MutateContractName(v string) error {
	return x._message.SetString(4, v)
}

func (x *Transaction) MethodName() string {
	return x._message.GetString(5)
}

func (x *Transaction) RawMethodName() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *Transaction) MutateMethodName(v string) error {
	return x._message.SetString(5, v)
}

func (x *Transaction) InputArgumentIterator() *TransactionInputArgumentIterator {
	return &TransactionInputArgumentIterator{iterator: x._message.GetMessageArrayIterator(6)}
}

type TransactionInputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionInputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionInputArgumentIterator) NextInputArgument() *MethodArgument {
	b, s := i.iterator.NextMessage()
	return MethodArgumentReader(b[:s])
}

func (x *Transaction) RawInputArgumentArray() []byte {
	return x._message.RawBufferForField(6, 0)
}

// builder

type TransactionBuilder struct {
	ProtocolVersion uint32
	VirtualChain uint64
	Sender *SenderBuilder
	Timestamp uint64
	ContractName string
	MethodName string
	InputArgument []*MethodArgumentBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *TransactionBuilder) arrayOfInputArgument() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.InputArgument))
	for i, v := range w.InputArgument {
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
	w._builder.WriteUint32(buf, w.ProtocolVersion)
	w._builder.WriteUint64(buf, w.VirtualChain)
	err = w._builder.WriteMessage(buf, w.Sender)
	if err != nil {
		return
	}
	w._builder.WriteUint64(buf, w.Timestamp)
	w._builder.WriteString(buf, w.ContractName)
	w._builder.WriteString(buf, w.MethodName)
	err = w._builder.WriteMessageArray(buf, w.arrayOfInputArgument())
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

var _SignedTransaction_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytesArray,}
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

func (x *SignedTransaction) TransactionContent() *Transaction {
	b, s := x._message.GetMessage(0)
	return TransactionReader(b[:s])
}

func (x *SignedTransaction) RawTransactionContent() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SignedTransaction) SignatureDataIterator() *SignedTransactionSignatureDataIterator {
	return &SignedTransactionSignatureDataIterator{iterator: x._message.GetBytesArrayIterator(1)}
}

type SignedTransactionSignatureDataIterator struct {
	iterator *membuffers.Iterator
}

func (i *SignedTransactionSignatureDataIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *SignedTransactionSignatureDataIterator) NextSignatureData() []byte {
	return i.iterator.NextBytes()
}

func (x *SignedTransaction) RawSignatureDataArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

// builder

type SignedTransactionBuilder struct {
	TransactionContent *TransactionBuilder
	SignatureData [][]byte

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
	err = w._builder.WriteMessage(buf, w.TransactionContent)
	if err != nil {
		return
	}
	w._builder.WriteBytesArray(buf, w.SignatureData)
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

var _TransactionReceipt_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeUint32,membuffers.TypeUint16,membuffers.TypeMessageArray,}
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

func (x *TransactionReceipt) Txid() primitives.Sha256 {
	return x._message.GetBytes(0)
}

func (x *TransactionReceipt) RawTxid() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *TransactionReceipt) MutateTxid(v primitives.Sha256) error {
	return x._message.SetBytes(0, v)
}

func (x *TransactionReceipt) Version() uint32 {
	return x._message.GetUint32(1)
}

func (x *TransactionReceipt) RawVersion() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *TransactionReceipt) MutateVersion(v uint32) error {
	return x._message.SetUint32(1, v)
}

func (x *TransactionReceipt) TransactionStatus() ExecutionStatus {
	return ExecutionStatus(x._message.GetUint16(2))
}

func (x *TransactionReceipt) RawTransactionStatus() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *TransactionReceipt) MutateTransactionStatus(v ExecutionStatus) error {
	return x._message.SetUint16(2, uint16(v))
}

func (x *TransactionReceipt) OutputArgumentIterator() *TransactionReceiptOutputArgumentIterator {
	return &TransactionReceiptOutputArgumentIterator{iterator: x._message.GetMessageArrayIterator(3)}
}

type TransactionReceiptOutputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionReceiptOutputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionReceiptOutputArgumentIterator) NextOutputArgument() *MethodArgument {
	b, s := i.iterator.NextMessage()
	return MethodArgumentReader(b[:s])
}

func (x *TransactionReceipt) RawOutputArgumentArray() []byte {
	return x._message.RawBufferForField(3, 0)
}

// builder

type TransactionReceiptBuilder struct {
	Txid primitives.Sha256
	Version uint32
	TransactionStatus ExecutionStatus
	OutputArgument []*MethodArgumentBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *TransactionReceiptBuilder) arrayOfOutputArgument() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.OutputArgument))
	for i, v := range w.OutputArgument {
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
	w._builder.WriteBytes(buf, w.Txid)
	w._builder.WriteUint32(buf, w.Version)
	w._builder.WriteUint16(buf, uint16(w.TransactionStatus))
	err = w._builder.WriteMessageArray(buf, w.arrayOfOutputArgument())
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

