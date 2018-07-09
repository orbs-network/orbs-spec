// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package client

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionRequest

// reader

type SendTransactionRequest struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _SendTransactionRequest_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var _SendTransactionRequest_Unions = [][]membuffers.FieldType{}

func SendTransactionRequestReader(buf []byte) *SendTransactionRequest {
	x := &SendTransactionRequest{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _SendTransactionRequest_Scheme, _SendTransactionRequest_Unions)
	return x
}

func (x *SendTransactionRequest) IsValid() bool {
	return x._message.IsValid()
}

func (x *SendTransactionRequest) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *SendTransactionRequest) SignedTransaction() *protocol.SignedTransaction {
	b, s := x._message.GetMessage(0)
	return protocol.SignedTransactionReader(b[:s])
}

func (x *SendTransactionRequest) RawSignedTransaction() []byte {
	return x._message.RawBufferForField(0, 0)
}

// builder

type SendTransactionRequestBuilder struct {
	SignedTransaction *protocol.SignedTransactionBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *SendTransactionRequestBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.SignedTransaction)
	if err != nil {
		return
	}
	return nil
}

func (w *SendTransactionRequestBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *SendTransactionRequestBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *SendTransactionRequestBuilder) Build() *SendTransactionRequest {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendTransactionRequestReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionResponse

// reader

type SendTransactionResponse struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _SendTransactionResponse_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,}
var _SendTransactionResponse_Unions = [][]membuffers.FieldType{}

func SendTransactionResponseReader(buf []byte) *SendTransactionResponse {
	x := &SendTransactionResponse{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _SendTransactionResponse_Scheme, _SendTransactionResponse_Unions)
	return x
}

func (x *SendTransactionResponse) IsValid() bool {
	return x._message.IsValid()
}

func (x *SendTransactionResponse) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *SendTransactionResponse) Receipt() *protocol.TransactionReceipt {
	b, s := x._message.GetMessage(0)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *SendTransactionResponse) RawReceipt() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SendTransactionResponse) TransactionStatus() protocol.TransactionStatus {
	return protocol.TransactionStatus(x._message.GetUint16(1))
}

func (x *SendTransactionResponse) RawTransactionStatus() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *SendTransactionResponse) MutateTransactionStatus(v protocol.TransactionStatus) error {
	return x._message.SetUint16(1, uint16(v))
}

func (x *SendTransactionResponse) BlockHeight() uint64 {
	return x._message.GetUint64(2)
}

func (x *SendTransactionResponse) RawBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *SendTransactionResponse) MutateBlockHeight(v uint64) error {
	return x._message.SetUint64(2, v)
}

func (x *SendTransactionResponse) BlockTimestamp() uint64 {
	return x._message.GetUint64(3)
}

func (x *SendTransactionResponse) RawBlockTimestamp() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *SendTransactionResponse) MutateBlockTimestamp(v uint64) error {
	return x._message.SetUint64(3, v)
}

// builder

type SendTransactionResponseBuilder struct {
	Receipt *protocol.TransactionReceiptBuilder
	TransactionStatus protocol.TransactionStatus
	BlockHeight uint64
	BlockTimestamp uint64

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *SendTransactionResponseBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.Receipt)
	if err != nil {
		return
	}
	w._builder.WriteUint16(buf, uint16(w.TransactionStatus))
	w._builder.WriteUint64(buf, w.BlockHeight)
	w._builder.WriteUint64(buf, w.BlockTimestamp)
	return nil
}

func (w *SendTransactionResponseBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *SendTransactionResponseBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *SendTransactionResponseBuilder) Build() *SendTransactionResponse {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendTransactionResponseReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message CallMethodRequest

// reader

type CallMethodRequest struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _CallMethodRequest_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var _CallMethodRequest_Unions = [][]membuffers.FieldType{}

func CallMethodRequestReader(buf []byte) *CallMethodRequest {
	x := &CallMethodRequest{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _CallMethodRequest_Scheme, _CallMethodRequest_Unions)
	return x
}

func (x *CallMethodRequest) IsValid() bool {
	return x._message.IsValid()
}

func (x *CallMethodRequest) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *CallMethodRequest) Transaction() *protocol.Transaction {
	b, s := x._message.GetMessage(0)
	return protocol.TransactionReader(b[:s])
}

func (x *CallMethodRequest) RawTransaction() []byte {
	return x._message.RawBufferForField(0, 0)
}

// builder

type CallMethodRequestBuilder struct {
	Transaction *protocol.TransactionBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *CallMethodRequestBuilder) Write(buf []byte) (err error) {
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
	return nil
}

func (w *CallMethodRequestBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *CallMethodRequestBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *CallMethodRequestBuilder) Build() *CallMethodRequest {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return CallMethodRequestReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message CallMethodResponse

// reader

type CallMethodResponse struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _CallMethodResponse_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,}
var _CallMethodResponse_Unions = [][]membuffers.FieldType{}

func CallMethodResponseReader(buf []byte) *CallMethodResponse {
	x := &CallMethodResponse{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _CallMethodResponse_Scheme, _CallMethodResponse_Unions)
	return x
}

func (x *CallMethodResponse) IsValid() bool {
	return x._message.IsValid()
}

func (x *CallMethodResponse) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *CallMethodResponse) OutputArgumentIterator() *CallMethodResponseOutputArgumentIterator {
	return &CallMethodResponseOutputArgumentIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type CallMethodResponseOutputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *CallMethodResponseOutputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *CallMethodResponseOutputArgumentIterator) NextOutputArgument() *protocol.MethodArgument {
	b, s := i.iterator.NextMessage()
	return protocol.MethodArgumentReader(b[:s])
}

func (x *CallMethodResponse) RawOutputArgumentArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *CallMethodResponse) CallStatus() protocol.CallMethodStatus {
	return protocol.CallMethodStatus(x._message.GetUint16(1))
}

func (x *CallMethodResponse) RawCallStatus() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *CallMethodResponse) MutateCallStatus(v protocol.CallMethodStatus) error {
	return x._message.SetUint16(1, uint16(v))
}

func (x *CallMethodResponse) BlockHeight() uint64 {
	return x._message.GetUint64(2)
}

func (x *CallMethodResponse) RawBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *CallMethodResponse) MutateBlockHeight(v uint64) error {
	return x._message.SetUint64(2, v)
}

func (x *CallMethodResponse) BlockTimestamp() uint64 {
	return x._message.GetUint64(3)
}

func (x *CallMethodResponse) RawBlockTimestamp() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *CallMethodResponse) MutateBlockTimestamp(v uint64) error {
	return x._message.SetUint64(3, v)
}

// builder

type CallMethodResponseBuilder struct {
	OutputArgument []*protocol.MethodArgumentBuilder
	CallStatus protocol.CallMethodStatus
	BlockHeight uint64
	BlockTimestamp uint64

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *CallMethodResponseBuilder) arrayOfOutputArgument() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.OutputArgument))
	for i, v := range w.OutputArgument {
		res[i] = v
	}
	return res
}

func (w *CallMethodResponseBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessageArray(buf, w.arrayOfOutputArgument())
	if err != nil {
		return
	}
	w._builder.WriteUint16(buf, uint16(w.CallStatus))
	w._builder.WriteUint64(buf, w.BlockHeight)
	w._builder.WriteUint64(buf, w.BlockTimestamp)
	return nil
}

func (w *CallMethodResponseBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *CallMethodResponseBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *CallMethodResponseBuilder) Build() *CallMethodResponse {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return CallMethodResponseReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusRequest

// reader

type GetTransactionStatusRequest struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _GetTransactionStatusRequest_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeUint64,}
var _GetTransactionStatusRequest_Unions = [][]membuffers.FieldType{}

func GetTransactionStatusRequestReader(buf []byte) *GetTransactionStatusRequest {
	x := &GetTransactionStatusRequest{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _GetTransactionStatusRequest_Scheme, _GetTransactionStatusRequest_Unions)
	return x
}

func (x *GetTransactionStatusRequest) IsValid() bool {
	return x._message.IsValid()
}

func (x *GetTransactionStatusRequest) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *GetTransactionStatusRequest) Txid() primitives.Sha256 {
	return x._message.GetBytes(0)
}

func (x *GetTransactionStatusRequest) RawTxid() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetTransactionStatusRequest) MutateTxid(v primitives.Sha256) error {
	return x._message.SetBytes(0, v)
}

func (x *GetTransactionStatusRequest) Timestamp() uint64 {
	return x._message.GetUint64(1)
}

func (x *GetTransactionStatusRequest) RawTimestamp() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *GetTransactionStatusRequest) MutateTimestamp(v uint64) error {
	return x._message.SetUint64(1, v)
}

// builder

type GetTransactionStatusRequestBuilder struct {
	Txid primitives.Sha256
	Timestamp uint64

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *GetTransactionStatusRequestBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUint64(buf, w.Timestamp)
	return nil
}

func (w *GetTransactionStatusRequestBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *GetTransactionStatusRequestBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *GetTransactionStatusRequestBuilder) Build() *GetTransactionStatusRequest {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionStatusRequestReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusResponse

// reader

type GetTransactionStatusResponse struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _GetTransactionStatusResponse_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,}
var _GetTransactionStatusResponse_Unions = [][]membuffers.FieldType{}

func GetTransactionStatusResponseReader(buf []byte) *GetTransactionStatusResponse {
	x := &GetTransactionStatusResponse{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _GetTransactionStatusResponse_Scheme, _GetTransactionStatusResponse_Unions)
	return x
}

func (x *GetTransactionStatusResponse) IsValid() bool {
	return x._message.IsValid()
}

func (x *GetTransactionStatusResponse) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *GetTransactionStatusResponse) Receipt() *protocol.TransactionReceipt {
	b, s := x._message.GetMessage(0)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *GetTransactionStatusResponse) RawReceipt() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetTransactionStatusResponse) Status() protocol.TransactionStatus {
	return protocol.TransactionStatus(x._message.GetUint16(1))
}

func (x *GetTransactionStatusResponse) RawStatus() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *GetTransactionStatusResponse) MutateStatus(v protocol.TransactionStatus) error {
	return x._message.SetUint16(1, uint16(v))
}

func (x *GetTransactionStatusResponse) BlockHeight() uint64 {
	return x._message.GetUint64(2)
}

func (x *GetTransactionStatusResponse) RawBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *GetTransactionStatusResponse) MutateBlockHeight(v uint64) error {
	return x._message.SetUint64(2, v)
}

func (x *GetTransactionStatusResponse) BlockTimestamp() uint64 {
	return x._message.GetUint64(3)
}

func (x *GetTransactionStatusResponse) RawBlockTimestamp() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *GetTransactionStatusResponse) MutateBlockTimestamp(v uint64) error {
	return x._message.SetUint64(3, v)
}

// builder

type GetTransactionStatusResponseBuilder struct {
	Receipt *protocol.TransactionReceiptBuilder
	Status protocol.TransactionStatus
	BlockHeight uint64
	BlockTimestamp uint64

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *GetTransactionStatusResponseBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.Receipt)
	if err != nil {
		return
	}
	w._builder.WriteUint16(buf, uint16(w.Status))
	w._builder.WriteUint64(buf, w.BlockHeight)
	w._builder.WriteUint64(buf, w.BlockTimestamp)
	return nil
}

func (w *GetTransactionStatusResponseBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *GetTransactionStatusResponseBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *GetTransactionStatusResponseBuilder) Build() *GetTransactionStatusResponse {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionStatusResponseReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

