// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionPool

type TransactionPool interface {
	handlers.TransactionRelayGossipHandler
	AddNewTransaction(input *AddNewTransactionInput) (*AddNewTransactionOutput, error)
	GetCommittedTransactionReceipt(input *GetCommittedTransactionReceiptInput) (*GetCommittedTransactionReceiptOutput, error)
	GetTransactionsForOrdering(input *GetTransactionsForOrderingInput) (*GetTransactionsForOrderingOutput, error)
	ValidateTransactionsForOrdering(input *ValidateTransactionsForOrderingInput) (*ValidateTransactionsForOrderingOutput, error)
	CommitTransactionReceipts(input *CommitTransactionReceiptsInput) (*CommitTransactionReceiptsOutput, error)
	RegisterTransactionResultsHandler(handler handlers.TransactionResultsHandler)
}

/////////////////////////////////////////////////////////////////////////////
// message AddNewTransactionInput

// reader

type AddNewTransactionInput struct {
	message membuffers.Message
}

var m_AddNewTransactionInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_AddNewTransactionInput_Unions = [][]membuffers.FieldType{}

func AddNewTransactionInputReader(buf []byte) *AddNewTransactionInput {
	x := &AddNewTransactionInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_AddNewTransactionInput_Scheme, m_AddNewTransactionInput_Unions)
	return x
}

func (x *AddNewTransactionInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *AddNewTransactionInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *AddNewTransactionInput) SignedTransaction() *protocol.SignedTransaction {
	b, s := x.message.GetMessage(0)
	return protocol.SignedTransactionReader(b[:s])
}

func (x *AddNewTransactionInput) RawSignedTransaction() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type AddNewTransactionInputBuilder struct {
	builder membuffers.Builder
	SignedTransaction *protocol.SignedTransactionBuilder
}

func (w *AddNewTransactionInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.SignedTransaction)
	if err != nil {
		return
	}
	return nil
}

func (w *AddNewTransactionInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *AddNewTransactionInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *AddNewTransactionInputBuilder) Build() *AddNewTransactionInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return AddNewTransactionInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message AddNewTransactionOutput

// reader

type AddNewTransactionOutput struct {
	message membuffers.Message
}

var m_AddNewTransactionOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeMessage,membuffers.TypeUint64,membuffers.TypeUint64,}
var m_AddNewTransactionOutput_Unions = [][]membuffers.FieldType{}

func AddNewTransactionOutputReader(buf []byte) *AddNewTransactionOutput {
	x := &AddNewTransactionOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_AddNewTransactionOutput_Scheme, m_AddNewTransactionOutput_Unions)
	return x
}

func (x *AddNewTransactionOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *AddNewTransactionOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *AddNewTransactionOutput) TransactionStatus() protocol.TransactionStatus {
	return protocol.TransactionStatus(x.message.GetUint16(0))
}

func (x *AddNewTransactionOutput) RawTransactionStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *AddNewTransactionOutput) MutateTransactionStatus(v protocol.TransactionStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *AddNewTransactionOutput) TransactionReceipt() *protocol.TransactionReceipt {
	b, s := x.message.GetMessage(1)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *AddNewTransactionOutput) RawTransactionReceipt() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *AddNewTransactionOutput) BlockHeight() uint64 {
	return x.message.GetUint64(2)
}

func (x *AddNewTransactionOutput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *AddNewTransactionOutput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(2, v)
}

func (x *AddNewTransactionOutput) BlockTimestamp() uint64 {
	return x.message.GetUint64(3)
}

func (x *AddNewTransactionOutput) RawBlockTimestamp() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *AddNewTransactionOutput) MutateBlockTimestamp(v uint64) error {
	return x.message.SetUint64(3, v)
}

// builder

type AddNewTransactionOutputBuilder struct {
	builder membuffers.Builder
	TransactionStatus protocol.TransactionStatus
	TransactionReceipt *protocol.TransactionReceiptBuilder
	BlockHeight uint64
	BlockTimestamp uint64
}

func (w *AddNewTransactionOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.TransactionStatus))
	err = w.builder.WriteMessage(buf, w.TransactionReceipt)
	if err != nil {
		return
	}
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteUint64(buf, w.BlockTimestamp)
	return nil
}

func (w *AddNewTransactionOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *AddNewTransactionOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *AddNewTransactionOutputBuilder) Build() *AddNewTransactionOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return AddNewTransactionOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetCommittedTransactionReceiptInput

// reader

type GetCommittedTransactionReceiptInput struct {
	message membuffers.Message
}

var m_GetCommittedTransactionReceiptInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeUint64,}
var m_GetCommittedTransactionReceiptInput_Unions = [][]membuffers.FieldType{}

func GetCommittedTransactionReceiptInputReader(buf []byte) *GetCommittedTransactionReceiptInput {
	x := &GetCommittedTransactionReceiptInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetCommittedTransactionReceiptInput_Scheme, m_GetCommittedTransactionReceiptInput_Unions)
	return x
}

func (x *GetCommittedTransactionReceiptInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetCommittedTransactionReceiptInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetCommittedTransactionReceiptInput) Txid() []byte {
	return x.message.GetBytes(0)
}

func (x *GetCommittedTransactionReceiptInput) RawTxid() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetCommittedTransactionReceiptInput) MutateTxid(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *GetCommittedTransactionReceiptInput) Timestamp() uint64 {
	return x.message.GetUint64(1)
}

func (x *GetCommittedTransactionReceiptInput) RawTimestamp() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *GetCommittedTransactionReceiptInput) MutateTimestamp(v uint64) error {
	return x.message.SetUint64(1, v)
}

// builder

type GetCommittedTransactionReceiptInputBuilder struct {
	builder membuffers.Builder
	Txid []byte
	Timestamp uint64
}

func (w *GetCommittedTransactionReceiptInputBuilder) Write(buf []byte) (err error) {
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
	w.builder.WriteUint64(buf, w.Timestamp)
	return nil
}

func (w *GetCommittedTransactionReceiptInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetCommittedTransactionReceiptInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetCommittedTransactionReceiptInputBuilder) Build() *GetCommittedTransactionReceiptInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetCommittedTransactionReceiptInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetCommittedTransactionReceiptOutput

// reader

type GetCommittedTransactionReceiptOutput struct {
	message membuffers.Message
}

var m_GetCommittedTransactionReceiptOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeMessage,membuffers.TypeUint64,membuffers.TypeUint64,}
var m_GetCommittedTransactionReceiptOutput_Unions = [][]membuffers.FieldType{}

func GetCommittedTransactionReceiptOutputReader(buf []byte) *GetCommittedTransactionReceiptOutput {
	x := &GetCommittedTransactionReceiptOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetCommittedTransactionReceiptOutput_Scheme, m_GetCommittedTransactionReceiptOutput_Unions)
	return x
}

func (x *GetCommittedTransactionReceiptOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetCommittedTransactionReceiptOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetCommittedTransactionReceiptOutput) TransactionStatus() protocol.TransactionStatus {
	return protocol.TransactionStatus(x.message.GetUint16(0))
}

func (x *GetCommittedTransactionReceiptOutput) RawTransactionStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetCommittedTransactionReceiptOutput) MutateTransactionStatus(v protocol.TransactionStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *GetCommittedTransactionReceiptOutput) TransactionReceipt() *protocol.TransactionReceipt {
	b, s := x.message.GetMessage(1)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *GetCommittedTransactionReceiptOutput) RawTransactionReceipt() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *GetCommittedTransactionReceiptOutput) BlockHeight() uint64 {
	return x.message.GetUint64(2)
}

func (x *GetCommittedTransactionReceiptOutput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *GetCommittedTransactionReceiptOutput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(2, v)
}

func (x *GetCommittedTransactionReceiptOutput) BlockTimestamp() uint64 {
	return x.message.GetUint64(3)
}

func (x *GetCommittedTransactionReceiptOutput) RawBlockTimestamp() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *GetCommittedTransactionReceiptOutput) MutateBlockTimestamp(v uint64) error {
	return x.message.SetUint64(3, v)
}

// builder

type GetCommittedTransactionReceiptOutputBuilder struct {
	builder membuffers.Builder
	TransactionStatus protocol.TransactionStatus
	TransactionReceipt *protocol.TransactionReceiptBuilder
	BlockHeight uint64
	BlockTimestamp uint64
}

func (w *GetCommittedTransactionReceiptOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.TransactionStatus))
	err = w.builder.WriteMessage(buf, w.TransactionReceipt)
	if err != nil {
		return
	}
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteUint64(buf, w.BlockTimestamp)
	return nil
}

func (w *GetCommittedTransactionReceiptOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetCommittedTransactionReceiptOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetCommittedTransactionReceiptOutputBuilder) Build() *GetCommittedTransactionReceiptOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetCommittedTransactionReceiptOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsForOrderingInput

// reader

type GetTransactionsForOrderingInput struct {
	message membuffers.Message
}

var m_GetTransactionsForOrderingInput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeUint32,}
var m_GetTransactionsForOrderingInput_Unions = [][]membuffers.FieldType{}

func GetTransactionsForOrderingInputReader(buf []byte) *GetTransactionsForOrderingInput {
	x := &GetTransactionsForOrderingInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetTransactionsForOrderingInput_Scheme, m_GetTransactionsForOrderingInput_Unions)
	return x
}

func (x *GetTransactionsForOrderingInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetTransactionsForOrderingInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetTransactionsForOrderingInput) BlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *GetTransactionsForOrderingInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetTransactionsForOrderingInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *GetTransactionsForOrderingInput) MaxTransactionsSetSizeKb() uint32 {
	return x.message.GetUint32(1)
}

func (x *GetTransactionsForOrderingInput) RawMaxTransactionsSetSizeKb() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *GetTransactionsForOrderingInput) MutateMaxTransactionsSetSizeKb(v uint32) error {
	return x.message.SetUint32(1, v)
}

func (x *GetTransactionsForOrderingInput) MaxNumberOfTransactions() uint32 {
	return x.message.GetUint32(2)
}

func (x *GetTransactionsForOrderingInput) RawMaxNumberOfTransactions() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *GetTransactionsForOrderingInput) MutateMaxNumberOfTransactions(v uint32) error {
	return x.message.SetUint32(2, v)
}

// builder

type GetTransactionsForOrderingInputBuilder struct {
	builder membuffers.Builder
	BlockHeight uint64
	MaxTransactionsSetSizeKb uint32
	MaxNumberOfTransactions uint32
}

func (w *GetTransactionsForOrderingInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteUint32(buf, w.MaxTransactionsSetSizeKb)
	w.builder.WriteUint32(buf, w.MaxNumberOfTransactions)
	return nil
}

func (w *GetTransactionsForOrderingInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetTransactionsForOrderingInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetTransactionsForOrderingInputBuilder) Build() *GetTransactionsForOrderingInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionsForOrderingInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsForOrderingOutput

// reader

type GetTransactionsForOrderingOutput struct {
	message membuffers.Message
}

var m_GetTransactionsForOrderingOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeMessageArray,}
var m_GetTransactionsForOrderingOutput_Unions = [][]membuffers.FieldType{}

func GetTransactionsForOrderingOutputReader(buf []byte) *GetTransactionsForOrderingOutput {
	x := &GetTransactionsForOrderingOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetTransactionsForOrderingOutput_Scheme, m_GetTransactionsForOrderingOutput_Unions)
	return x
}

func (x *GetTransactionsForOrderingOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetTransactionsForOrderingOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetTransactionsForOrderingOutput) Status() protocol.RequestStatus {
	return protocol.RequestStatus(x.message.GetUint16(0))
}

func (x *GetTransactionsForOrderingOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetTransactionsForOrderingOutput) MutateStatus(v protocol.RequestStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *GetTransactionsForOrderingOutput) SignedTransactionIterator() *GetTransactionsForOrderingOutputSignedTransactionIterator {
	return &GetTransactionsForOrderingOutputSignedTransactionIterator{iterator: x.message.GetMessageArrayIterator(1)}
}

type GetTransactionsForOrderingOutputSignedTransactionIterator struct {
	iterator *membuffers.Iterator
}

func (i *GetTransactionsForOrderingOutputSignedTransactionIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *GetTransactionsForOrderingOutputSignedTransactionIterator) NextSignedTransaction() *protocol.SignedTransaction {
	b, s := i.iterator.NextMessage()
	return protocol.SignedTransactionReader(b[:s])
}

func (x *GetTransactionsForOrderingOutput) RawSignedTransactionArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type GetTransactionsForOrderingOutputBuilder struct {
	builder membuffers.Builder
	Status protocol.RequestStatus
	SignedTransaction []*protocol.SignedTransactionBuilder
}

func (w *GetTransactionsForOrderingOutputBuilder) arrayOfSignedTransaction() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.SignedTransaction))
	for i, v := range w.SignedTransaction {
		res[i] = v
	}
	return res
}

func (w *GetTransactionsForOrderingOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.Status))
	err = w.builder.WriteMessageArray(buf, w.arrayOfSignedTransaction())
	if err != nil {
		return
	}
	return nil
}

func (w *GetTransactionsForOrderingOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetTransactionsForOrderingOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetTransactionsForOrderingOutputBuilder) Build() *GetTransactionsForOrderingOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionsForOrderingOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsForOrderingInput

// reader

type ValidateTransactionsForOrderingInput struct {
	message membuffers.Message
}

var m_ValidateTransactionsForOrderingInput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeMessageArray,}
var m_ValidateTransactionsForOrderingInput_Unions = [][]membuffers.FieldType{}

func ValidateTransactionsForOrderingInputReader(buf []byte) *ValidateTransactionsForOrderingInput {
	x := &ValidateTransactionsForOrderingInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ValidateTransactionsForOrderingInput_Scheme, m_ValidateTransactionsForOrderingInput_Unions)
	return x
}

func (x *ValidateTransactionsForOrderingInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ValidateTransactionsForOrderingInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ValidateTransactionsForOrderingInput) BlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *ValidateTransactionsForOrderingInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ValidateTransactionsForOrderingInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *ValidateTransactionsForOrderingInput) TransactionIterator() *ValidateTransactionsForOrderingInputTransactionIterator {
	return &ValidateTransactionsForOrderingInputTransactionIterator{iterator: x.message.GetMessageArrayIterator(1)}
}

type ValidateTransactionsForOrderingInputTransactionIterator struct {
	iterator *membuffers.Iterator
}

func (i *ValidateTransactionsForOrderingInputTransactionIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ValidateTransactionsForOrderingInputTransactionIterator) NextTransaction() *protocol.SignedTransaction {
	b, s := i.iterator.NextMessage()
	return protocol.SignedTransactionReader(b[:s])
}

func (x *ValidateTransactionsForOrderingInput) RawTransactionArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type ValidateTransactionsForOrderingInputBuilder struct {
	builder membuffers.Builder
	BlockHeight uint64
	Transaction []*protocol.SignedTransactionBuilder
}

func (w *ValidateTransactionsForOrderingInputBuilder) arrayOfTransaction() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.Transaction))
	for i, v := range w.Transaction {
		res[i] = v
	}
	return res
}

func (w *ValidateTransactionsForOrderingInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint64(buf, w.BlockHeight)
	err = w.builder.WriteMessageArray(buf, w.arrayOfTransaction())
	if err != nil {
		return
	}
	return nil
}

func (w *ValidateTransactionsForOrderingInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ValidateTransactionsForOrderingInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ValidateTransactionsForOrderingInputBuilder) Build() *ValidateTransactionsForOrderingInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ValidateTransactionsForOrderingInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsForOrderingOutput

// reader

type ValidateTransactionsForOrderingOutput struct {
	message membuffers.Message
}

var m_ValidateTransactionsForOrderingOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,}
var m_ValidateTransactionsForOrderingOutput_Unions = [][]membuffers.FieldType{}

func ValidateTransactionsForOrderingOutputReader(buf []byte) *ValidateTransactionsForOrderingOutput {
	x := &ValidateTransactionsForOrderingOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ValidateTransactionsForOrderingOutput_Scheme, m_ValidateTransactionsForOrderingOutput_Unions)
	return x
}

func (x *ValidateTransactionsForOrderingOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ValidateTransactionsForOrderingOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ValidateTransactionsForOrderingOutput) Status() protocol.RequestStatus {
	return protocol.RequestStatus(x.message.GetUint16(0))
}

func (x *ValidateTransactionsForOrderingOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ValidateTransactionsForOrderingOutput) MutateStatus(v protocol.RequestStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

// builder

type ValidateTransactionsForOrderingOutputBuilder struct {
	builder membuffers.Builder
	Status protocol.RequestStatus
}

func (w *ValidateTransactionsForOrderingOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.Status))
	return nil
}

func (w *ValidateTransactionsForOrderingOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ValidateTransactionsForOrderingOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ValidateTransactionsForOrderingOutputBuilder) Build() *ValidateTransactionsForOrderingOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ValidateTransactionsForOrderingOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message CommitTransactionReceiptsInput

// reader

type CommitTransactionReceiptsInput struct {
	message membuffers.Message
}

var m_CommitTransactionReceiptsInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessageArray,membuffers.TypeUint64,}
var m_CommitTransactionReceiptsInput_Unions = [][]membuffers.FieldType{}

func CommitTransactionReceiptsInputReader(buf []byte) *CommitTransactionReceiptsInput {
	x := &CommitTransactionReceiptsInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_CommitTransactionReceiptsInput_Scheme, m_CommitTransactionReceiptsInput_Unions)
	return x
}

func (x *CommitTransactionReceiptsInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *CommitTransactionReceiptsInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *CommitTransactionReceiptsInput) ResultsBlockHeader() *protocol.ResultsBlockHeader {
	b, s := x.message.GetMessage(0)
	return protocol.ResultsBlockHeaderReader(b[:s])
}

func (x *CommitTransactionReceiptsInput) RawResultsBlockHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *CommitTransactionReceiptsInput) TransactionReceiptIterator() *CommitTransactionReceiptsInputTransactionReceiptIterator {
	return &CommitTransactionReceiptsInputTransactionReceiptIterator{iterator: x.message.GetMessageArrayIterator(1)}
}

type CommitTransactionReceiptsInputTransactionReceiptIterator struct {
	iterator *membuffers.Iterator
}

func (i *CommitTransactionReceiptsInputTransactionReceiptIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *CommitTransactionReceiptsInputTransactionReceiptIterator) NextTransactionReceipt() *protocol.TransactionReceipt {
	b, s := i.iterator.NextMessage()
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *CommitTransactionReceiptsInput) RawTransactionReceiptArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *CommitTransactionReceiptsInput) LastCommittedBlockHeight() uint64 {
	return x.message.GetUint64(2)
}

func (x *CommitTransactionReceiptsInput) RawLastCommittedBlockHeight() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *CommitTransactionReceiptsInput) MutateLastCommittedBlockHeight(v uint64) error {
	return x.message.SetUint64(2, v)
}

// builder

type CommitTransactionReceiptsInputBuilder struct {
	builder membuffers.Builder
	ResultsBlockHeader *protocol.ResultsBlockHeaderBuilder
	TransactionReceipt []*protocol.TransactionReceiptBuilder
	LastCommittedBlockHeight uint64
}

func (w *CommitTransactionReceiptsInputBuilder) arrayOfTransactionReceipt() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.TransactionReceipt))
	for i, v := range w.TransactionReceipt {
		res[i] = v
	}
	return res
}

func (w *CommitTransactionReceiptsInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.ResultsBlockHeader)
	if err != nil {
		return
	}
	err = w.builder.WriteMessageArray(buf, w.arrayOfTransactionReceipt())
	if err != nil {
		return
	}
	w.builder.WriteUint64(buf, w.LastCommittedBlockHeight)
	return nil
}

func (w *CommitTransactionReceiptsInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *CommitTransactionReceiptsInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *CommitTransactionReceiptsInputBuilder) Build() *CommitTransactionReceiptsInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return CommitTransactionReceiptsInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message CommitTransactionReceiptsOutput

// reader

type CommitTransactionReceiptsOutput struct {
	message membuffers.Message
}

var m_CommitTransactionReceiptsOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint64,}
var m_CommitTransactionReceiptsOutput_Unions = [][]membuffers.FieldType{}

func CommitTransactionReceiptsOutputReader(buf []byte) *CommitTransactionReceiptsOutput {
	x := &CommitTransactionReceiptsOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_CommitTransactionReceiptsOutput_Scheme, m_CommitTransactionReceiptsOutput_Unions)
	return x
}

func (x *CommitTransactionReceiptsOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *CommitTransactionReceiptsOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *CommitTransactionReceiptsOutput) NextDesiredBlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *CommitTransactionReceiptsOutput) RawNextDesiredBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *CommitTransactionReceiptsOutput) MutateNextDesiredBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *CommitTransactionReceiptsOutput) LastCommittedBlockHeight() uint64 {
	return x.message.GetUint64(1)
}

func (x *CommitTransactionReceiptsOutput) RawLastCommittedBlockHeight() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *CommitTransactionReceiptsOutput) MutateLastCommittedBlockHeight(v uint64) error {
	return x.message.SetUint64(1, v)
}

// builder

type CommitTransactionReceiptsOutputBuilder struct {
	builder membuffers.Builder
	NextDesiredBlockHeight uint64
	LastCommittedBlockHeight uint64
}

func (w *CommitTransactionReceiptsOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint64(buf, w.NextDesiredBlockHeight)
	w.builder.WriteUint64(buf, w.LastCommittedBlockHeight)
	return nil
}

func (w *CommitTransactionReceiptsOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *CommitTransactionReceiptsOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *CommitTransactionReceiptsOutputBuilder) Build() *CommitTransactionReceiptsOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return CommitTransactionReceiptsOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

