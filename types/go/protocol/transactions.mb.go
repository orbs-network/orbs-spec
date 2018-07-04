// AUTO GENERATED FILE (by membufc proto compiler)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message Transaction

// reader

type Transaction struct {
	message membuffers.Message
}

var m_Transaction_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint32,membuffers.TypeMessage,membuffers.TypeUint64,membuffers.TypeString,membuffers.TypeString,membuffers.TypeMessageArray,}
var m_Transaction_Unions = [][]membuffers.FieldType{}

func TransactionReader(buf []byte) *Transaction {
	x := &Transaction{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_Transaction_Scheme, m_Transaction_Unions)
	return x
}

func (x *Transaction) IsValid() bool {
	return x.message.IsValid()
}

func (x *Transaction) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *Transaction) ProtocolVersion() uint32 {
	return x.message.GetUint32(0)
}

func (x *Transaction) RawProtocolVersion() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *Transaction) MutateProtocolVersion(v uint32) error {
	return x.message.SetUint32(0, v)
}

func (x *Transaction) VirtualChain() uint32 {
	return x.message.GetUint32(1)
}

func (x *Transaction) RawVirtualChain() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *Transaction) MutateVirtualChain(v uint32) error {
	return x.message.SetUint32(1, v)
}

func (x *Transaction) Sender() *Sender {
	b, s := x.message.GetMessage(2)
	return SenderReader(b[:s])
}

func (x *Transaction) RawSender() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *Transaction) Timestamp() uint64 {
	return x.message.GetUint64(3)
}

func (x *Transaction) RawTimestamp() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *Transaction) MutateTimestamp(v uint64) error {
	return x.message.SetUint64(3, v)
}

func (x *Transaction) ContractName() string {
	return x.message.GetString(4)
}

func (x *Transaction) RawContractName() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *Transaction) MutateContractName(v string) error {
	return x.message.SetString(4, v)
}

func (x *Transaction) MethodName() string {
	return x.message.GetString(5)
}

func (x *Transaction) RawMethodName() []byte {
	return x.message.RawBufferForField(5, 0)
}

func (x *Transaction) MutateMethodName(v string) error {
	return x.message.SetString(5, v)
}

func (x *Transaction) InputArgumentIterator() *TransactionInputArgumentIterator {
	return &TransactionInputArgumentIterator{iterator: x.message.GetMessageArrayIterator(6)}
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
	return x.message.RawBufferForField(6, 0)
}

// builder

type TransactionBuilder struct {
	builder membuffers.Builder
	ProtocolVersion uint32
	VirtualChain uint32
	Sender *SenderBuilder
	Timestamp uint64
	ContractName string
	MethodName string
	InputArgument []*MethodArgumentBuilder
}

func (w *TransactionBuilder) arrayOfInputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.InputArgument))
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
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.ProtocolVersion)
	w.builder.WriteUint32(buf, w.VirtualChain)
	err = w.builder.WriteMessage(buf, w.Sender)
	if err != nil {
		return
	}
	w.builder.WriteUint64(buf, w.Timestamp)
	w.builder.WriteString(buf, w.ContractName)
	w.builder.WriteString(buf, w.MethodName)
	err = w.builder.WriteMessageArray(buf, w.arrayOfInputArgument())
	if err != nil {
		return
	}
	return nil
}

func (w *TransactionBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
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
	message membuffers.Message
}

var m_SignedTransaction_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytesArray,}
var m_SignedTransaction_Unions = [][]membuffers.FieldType{}

func SignedTransactionReader(buf []byte) *SignedTransaction {
	x := &SignedTransaction{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SignedTransaction_Scheme, m_SignedTransaction_Unions)
	return x
}

func (x *SignedTransaction) IsValid() bool {
	return x.message.IsValid()
}

func (x *SignedTransaction) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SignedTransaction) TransactionContent() *Transaction {
	b, s := x.message.GetMessage(0)
	return TransactionReader(b[:s])
}

func (x *SignedTransaction) RawTransactionContent() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SignedTransaction) SignatureDataIterator() *SignedTransactionSignatureDataIterator {
	return &SignedTransactionSignatureDataIterator{iterator: x.message.GetBytesArrayIterator(1)}
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
	return x.message.RawBufferForField(1, 0)
}

// builder

type SignedTransactionBuilder struct {
	builder membuffers.Builder
	TransactionContent *TransactionBuilder
	SignatureData [][]byte
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
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.TransactionContent)
	if err != nil {
		return
	}
	w.builder.WriteBytesArray(buf, w.SignatureData)
	return nil
}

func (w *SignedTransactionBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SignedTransactionBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
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
	message membuffers.Message
}

var m_TransactionReceipt_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeUint32,membuffers.TypeUint16,membuffers.TypeMessageArray,}
var m_TransactionReceipt_Unions = [][]membuffers.FieldType{}

func TransactionReceiptReader(buf []byte) *TransactionReceipt {
	x := &TransactionReceipt{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TransactionReceipt_Scheme, m_TransactionReceipt_Unions)
	return x
}

func (x *TransactionReceipt) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionReceipt) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionReceipt) Txid() []byte {
	return x.message.GetBytes(0)
}

func (x *TransactionReceipt) RawTxid() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionReceipt) MutateTxid(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *TransactionReceipt) Version() uint32 {
	return x.message.GetUint32(1)
}

func (x *TransactionReceipt) RawVersion() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *TransactionReceipt) MutateVersion(v uint32) error {
	return x.message.SetUint32(1, v)
}

func (x *TransactionReceipt) TransactionStatus() ExecutionStatus {
	return ExecutionStatus(x.message.GetUint16(2))
}

func (x *TransactionReceipt) RawTransactionStatus() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *TransactionReceipt) MutateTransactionStatus(v ExecutionStatus) error {
	return x.message.SetUint16(2, uint16(v))
}

func (x *TransactionReceipt) OutputArgumentIterator() *TransactionReceiptOutputArgumentIterator {
	return &TransactionReceiptOutputArgumentIterator{iterator: x.message.GetMessageArrayIterator(3)}
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
	return x.message.RawBufferForField(3, 0)
}

// builder

type TransactionReceiptBuilder struct {
	builder membuffers.Builder
	Txid []byte
	Version uint32
	TransactionStatus ExecutionStatus
	OutputArgument []*MethodArgumentBuilder
}

func (w *TransactionReceiptBuilder) arrayOfOutputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.OutputArgument))
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
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Txid)
	w.builder.WriteUint32(buf, w.Version)
	w.builder.WriteUint16(buf, uint16(w.TransactionStatus))
	err = w.builder.WriteMessageArray(buf, w.arrayOfOutputArgument())
	if err != nil {
		return
	}
	return nil
}

func (w *TransactionReceiptBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionReceiptBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
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

