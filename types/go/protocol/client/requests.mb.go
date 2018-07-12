// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package client

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionRequest

// reader

type SendTransactionRequest struct {
	// SignedTransaction protocol.SignedTransaction

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *SendTransactionRequest) String() string {
	return fmt.Sprintf("{SignedTransaction:%s,}", x.StringSignedTransaction())
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

func (x *SendTransactionRequest) StringSignedTransaction() string {
	return x.SignedTransaction().String()
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
	// TransactionReceipt protocol.TransactionReceipt
	// TransactionStatus protocol.TransactionStatus
	// BlockHeight primitives.BlockHeight
	// BlockTimestamp primitives.Timestamp

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *SendTransactionResponse) String() string {
	return fmt.Sprintf("{TransactionReceipt:%s,TransactionStatus:%s,BlockHeight:%s,BlockTimestamp:%s,}", x.StringTransactionReceipt(), x.StringTransactionStatus(), x.StringBlockHeight(), x.StringBlockTimestamp())
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

func (x *SendTransactionResponse) TransactionReceipt() *protocol.TransactionReceipt {
	b, s := x._message.GetMessage(0)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *SendTransactionResponse) RawTransactionReceipt() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SendTransactionResponse) StringTransactionReceipt() string {
	return x.TransactionReceipt().String()
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

func (x *SendTransactionResponse) StringTransactionStatus() string {
	return x.TransactionStatus().String()
}

func (x *SendTransactionResponse) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *SendTransactionResponse) RawBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *SendTransactionResponse) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *SendTransactionResponse) StringBlockHeight() string {
	return fmt.Sprintf("%x", x.BlockHeight())
}

func (x *SendTransactionResponse) BlockTimestamp() primitives.Timestamp {
	return primitives.Timestamp(x._message.GetUint64(3))
}

func (x *SendTransactionResponse) RawBlockTimestamp() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *SendTransactionResponse) MutateBlockTimestamp(v primitives.Timestamp) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *SendTransactionResponse) StringBlockTimestamp() string {
	return fmt.Sprintf("%x", x.BlockTimestamp())
}

// builder

type SendTransactionResponseBuilder struct {
	TransactionReceipt *protocol.TransactionReceiptBuilder
	TransactionStatus protocol.TransactionStatus
	BlockHeight primitives.BlockHeight
	BlockTimestamp primitives.Timestamp

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
	err = w._builder.WriteMessage(buf, w.TransactionReceipt)
	if err != nil {
		return
	}
	w._builder.WriteUint16(buf, uint16(w.TransactionStatus))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteUint64(buf, uint64(w.BlockTimestamp))
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
	// Transaction protocol.Transaction

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *CallMethodRequest) String() string {
	return fmt.Sprintf("{Transaction:%s,}", x.StringTransaction())
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

func (x *CallMethodRequest) StringTransaction() string {
	return x.Transaction().String()
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
	// OutputArguments []protocol.MethodArgument
	// CallResult protocol.ExecutionResult
	// BlockHeight primitives.BlockHeight
	// BlockTimestamp primitives.Timestamp

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *CallMethodResponse) String() string {
	return fmt.Sprintf("{OutputArguments:%s,CallResult:%s,BlockHeight:%s,BlockTimestamp:%s,}", x.StringOutputArguments(), x.StringCallResult(), x.StringBlockHeight(), x.StringBlockTimestamp())
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

func (x *CallMethodResponse) OutputArgumentsIterator() *CallMethodResponseOutputArgumentsIterator {
	return &CallMethodResponseOutputArgumentsIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type CallMethodResponseOutputArgumentsIterator struct {
	iterator *membuffers.Iterator
}

func (i *CallMethodResponseOutputArgumentsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *CallMethodResponseOutputArgumentsIterator) NextOutputArguments() *protocol.MethodArgument {
	b, s := i.iterator.NextMessage()
	return protocol.MethodArgumentReader(b[:s])
}

func (x *CallMethodResponse) RawOutputArgumentsArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *CallMethodResponse) StringOutputArguments() (res string) {
	res = "["
	for i := x.OutputArgumentsIterator(); i.HasNext(); {
		res += i.NextOutputArguments().String() + ","
	}
	res += "]"
	return
}

func (x *CallMethodResponse) CallResult() protocol.ExecutionResult {
	return protocol.ExecutionResult(x._message.GetUint16(1))
}

func (x *CallMethodResponse) RawCallResult() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *CallMethodResponse) MutateCallResult(v protocol.ExecutionResult) error {
	return x._message.SetUint16(1, uint16(v))
}

func (x *CallMethodResponse) StringCallResult() string {
	return x.CallResult().String()
}

func (x *CallMethodResponse) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *CallMethodResponse) RawBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *CallMethodResponse) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *CallMethodResponse) StringBlockHeight() string {
	return fmt.Sprintf("%x", x.BlockHeight())
}

func (x *CallMethodResponse) BlockTimestamp() primitives.Timestamp {
	return primitives.Timestamp(x._message.GetUint64(3))
}

func (x *CallMethodResponse) RawBlockTimestamp() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *CallMethodResponse) MutateBlockTimestamp(v primitives.Timestamp) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *CallMethodResponse) StringBlockTimestamp() string {
	return fmt.Sprintf("%x", x.BlockTimestamp())
}

// builder

type CallMethodResponseBuilder struct {
	OutputArguments []*protocol.MethodArgumentBuilder
	CallResult protocol.ExecutionResult
	BlockHeight primitives.BlockHeight
	BlockTimestamp primitives.Timestamp

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *CallMethodResponseBuilder) arrayOfOutputArguments() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.OutputArguments))
	for i, v := range w.OutputArguments {
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
	err = w._builder.WriteMessageArray(buf, w.arrayOfOutputArguments())
	if err != nil {
		return
	}
	w._builder.WriteUint16(buf, uint16(w.CallResult))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteUint64(buf, uint64(w.BlockTimestamp))
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
	// TransactionTimestamp primitives.Timestamp
	// Txhash primitives.Sha256

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *GetTransactionStatusRequest) String() string {
	return fmt.Sprintf("{TransactionTimestamp:%s,Txhash:%s,}", x.StringTransactionTimestamp(), x.StringTxhash())
}

var _GetTransactionStatusRequest_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeBytes,}
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

func (x *GetTransactionStatusRequest) TransactionTimestamp() primitives.Timestamp {
	return primitives.Timestamp(x._message.GetUint64(0))
}

func (x *GetTransactionStatusRequest) RawTransactionTimestamp() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetTransactionStatusRequest) MutateTransactionTimestamp(v primitives.Timestamp) error {
	return x._message.SetUint64(0, uint64(v))
}

func (x *GetTransactionStatusRequest) StringTransactionTimestamp() string {
	return fmt.Sprintf("%x", x.TransactionTimestamp())
}

func (x *GetTransactionStatusRequest) Txhash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(1))
}

func (x *GetTransactionStatusRequest) RawTxhash() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *GetTransactionStatusRequest) MutateTxhash(v primitives.Sha256) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *GetTransactionStatusRequest) StringTxhash() string {
	return fmt.Sprintf("%x", x.Txhash())
}

// builder

type GetTransactionStatusRequestBuilder struct {
	TransactionTimestamp primitives.Timestamp
	Txhash primitives.Sha256

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
	w._builder.WriteUint64(buf, uint64(w.TransactionTimestamp))
	w._builder.WriteBytes(buf, []byte(w.Txhash))
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
	// TransactionReceipt protocol.TransactionReceipt
	// TransactionStatus protocol.TransactionStatus
	// BlockHeight primitives.BlockHeight
	// BlockTimestamp primitives.Timestamp

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *GetTransactionStatusResponse) String() string {
	return fmt.Sprintf("{TransactionReceipt:%s,TransactionStatus:%s,BlockHeight:%s,BlockTimestamp:%s,}", x.StringTransactionReceipt(), x.StringTransactionStatus(), x.StringBlockHeight(), x.StringBlockTimestamp())
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

func (x *GetTransactionStatusResponse) TransactionReceipt() *protocol.TransactionReceipt {
	b, s := x._message.GetMessage(0)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *GetTransactionStatusResponse) RawTransactionReceipt() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetTransactionStatusResponse) StringTransactionReceipt() string {
	return x.TransactionReceipt().String()
}

func (x *GetTransactionStatusResponse) TransactionStatus() protocol.TransactionStatus {
	return protocol.TransactionStatus(x._message.GetUint16(1))
}

func (x *GetTransactionStatusResponse) RawTransactionStatus() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *GetTransactionStatusResponse) MutateTransactionStatus(v protocol.TransactionStatus) error {
	return x._message.SetUint16(1, uint16(v))
}

func (x *GetTransactionStatusResponse) StringTransactionStatus() string {
	return x.TransactionStatus().String()
}

func (x *GetTransactionStatusResponse) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *GetTransactionStatusResponse) RawBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *GetTransactionStatusResponse) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *GetTransactionStatusResponse) StringBlockHeight() string {
	return fmt.Sprintf("%x", x.BlockHeight())
}

func (x *GetTransactionStatusResponse) BlockTimestamp() primitives.Timestamp {
	return primitives.Timestamp(x._message.GetUint64(3))
}

func (x *GetTransactionStatusResponse) RawBlockTimestamp() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *GetTransactionStatusResponse) MutateBlockTimestamp(v primitives.Timestamp) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *GetTransactionStatusResponse) StringBlockTimestamp() string {
	return fmt.Sprintf("%x", x.BlockTimestamp())
}

// builder

type GetTransactionStatusResponseBuilder struct {
	TransactionReceipt *protocol.TransactionReceiptBuilder
	TransactionStatus protocol.TransactionStatus
	BlockHeight primitives.BlockHeight
	BlockTimestamp primitives.Timestamp

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
	err = w._builder.WriteMessage(buf, w.TransactionReceipt)
	if err != nil {
		return
	}
	w._builder.WriteUint16(buf, uint16(w.TransactionStatus))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteUint64(buf, uint64(w.BlockTimestamp))
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

