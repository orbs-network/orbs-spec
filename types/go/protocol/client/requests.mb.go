// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package client

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionRequest

// reader

type SendTransactionRequest struct {
	// SignedTransaction protocol.SignedTransaction

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *SendTransactionRequest) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedTransaction:%s,}", x.StringSignedTransaction())
}

var _SendTransactionRequest_Scheme = []membuffers.FieldType{membuffers.TypeMessage}
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

func (x *SendTransactionRequest) Equal(y *SendTransactionRequest) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *SendTransactionRequest) SignedTransaction() *protocol.SignedTransaction {
	b, s := x._message.GetMessage(0)
	return protocol.SignedTransactionReader(b[:s])
}

func (x *SendTransactionRequest) RawSignedTransaction() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SendTransactionRequest) RawSignedTransactionWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *SendTransactionRequest) StringSignedTransaction() string {
	return x.SignedTransaction().String()
}

// builder

type SendTransactionRequestBuilder struct {
	SignedTransaction *protocol.SignedTransactionBuilder

	// internal
	// implements membuffers.Builder
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
	// RequestResult RequestResult
	// TransactionStatus protocol.TransactionStatus
	// TransactionReceipt protocol.TransactionReceipt

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *SendTransactionResponse) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RequestResult:%s,TransactionStatus:%s,TransactionReceipt:%s,}", x.StringRequestResult(), x.StringTransactionStatus(), x.StringTransactionReceipt())
}

var _SendTransactionResponse_Scheme = []membuffers.FieldType{membuffers.TypeMessage, membuffers.TypeUint16, membuffers.TypeMessage}
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

func (x *SendTransactionResponse) Equal(y *SendTransactionResponse) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *SendTransactionResponse) RequestResult() *RequestResult {
	b, s := x._message.GetMessage(0)
	return RequestResultReader(b[:s])
}

func (x *SendTransactionResponse) RawRequestResult() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SendTransactionResponse) RawRequestResultWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *SendTransactionResponse) StringRequestResult() string {
	return x.RequestResult().String()
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

func (x *SendTransactionResponse) TransactionReceipt() *protocol.TransactionReceipt {
	b, s := x._message.GetMessage(2)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *SendTransactionResponse) RawTransactionReceipt() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *SendTransactionResponse) RawTransactionReceiptWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *SendTransactionResponse) StringTransactionReceipt() string {
	return x.TransactionReceipt().String()
}

// builder

type SendTransactionResponseBuilder struct {
	RequestResult      *RequestResultBuilder
	TransactionStatus  protocol.TransactionStatus
	TransactionReceipt *protocol.TransactionReceiptBuilder

	// internal
	// implements membuffers.Builder
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
	err = w._builder.WriteMessage(buf, w.RequestResult)
	if err != nil {
		return
	}
	w._builder.WriteUint16(buf, uint16(w.TransactionStatus))
	err = w._builder.WriteMessage(buf, w.TransactionReceipt)
	if err != nil {
		return
	}
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
// message RunQueryRequest

// reader

type RunQueryRequest struct {
	// SignedQuery protocol.SignedQuery

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *RunQueryRequest) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedQuery:%s,}", x.StringSignedQuery())
}

var _RunQueryRequest_Scheme = []membuffers.FieldType{membuffers.TypeMessage}
var _RunQueryRequest_Unions = [][]membuffers.FieldType{}

func RunQueryRequestReader(buf []byte) *RunQueryRequest {
	x := &RunQueryRequest{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _RunQueryRequest_Scheme, _RunQueryRequest_Unions)
	return x
}

func (x *RunQueryRequest) IsValid() bool {
	return x._message.IsValid()
}

func (x *RunQueryRequest) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *RunQueryRequest) Equal(y *RunQueryRequest) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *RunQueryRequest) SignedQuery() *protocol.SignedQuery {
	b, s := x._message.GetMessage(0)
	return protocol.SignedQueryReader(b[:s])
}

func (x *RunQueryRequest) RawSignedQuery() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *RunQueryRequest) RawSignedQueryWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *RunQueryRequest) StringSignedQuery() string {
	return x.SignedQuery().String()
}

// builder

type RunQueryRequestBuilder struct {
	SignedQuery *protocol.SignedQueryBuilder

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *RunQueryRequestBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.SignedQuery)
	if err != nil {
		return
	}
	return nil
}

func (w *RunQueryRequestBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *RunQueryRequestBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *RunQueryRequestBuilder) Build() *RunQueryRequest {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RunQueryRequestReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message RunQueryResponse

// reader

type RunQueryResponse struct {
	// RequestResult RequestResult
	// QueryResult protocol.QueryResult

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *RunQueryResponse) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RequestResult:%s,QueryResult:%s,}", x.StringRequestResult(), x.StringQueryResult())
}

var _RunQueryResponse_Scheme = []membuffers.FieldType{membuffers.TypeMessage, membuffers.TypeMessage}
var _RunQueryResponse_Unions = [][]membuffers.FieldType{}

func RunQueryResponseReader(buf []byte) *RunQueryResponse {
	x := &RunQueryResponse{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _RunQueryResponse_Scheme, _RunQueryResponse_Unions)
	return x
}

func (x *RunQueryResponse) IsValid() bool {
	return x._message.IsValid()
}

func (x *RunQueryResponse) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *RunQueryResponse) Equal(y *RunQueryResponse) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *RunQueryResponse) RequestResult() *RequestResult {
	b, s := x._message.GetMessage(0)
	return RequestResultReader(b[:s])
}

func (x *RunQueryResponse) RawRequestResult() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *RunQueryResponse) RawRequestResultWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *RunQueryResponse) StringRequestResult() string {
	return x.RequestResult().String()
}

func (x *RunQueryResponse) QueryResult() *protocol.QueryResult {
	b, s := x._message.GetMessage(1)
	return protocol.QueryResultReader(b[:s])
}

func (x *RunQueryResponse) RawQueryResult() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *RunQueryResponse) RawQueryResultWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *RunQueryResponse) StringQueryResult() string {
	return x.QueryResult().String()
}

// builder

type RunQueryResponseBuilder struct {
	RequestResult *RequestResultBuilder
	QueryResult   *protocol.QueryResultBuilder

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *RunQueryResponseBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.RequestResult)
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.QueryResult)
	if err != nil {
		return
	}
	return nil
}

func (w *RunQueryResponseBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *RunQueryResponseBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *RunQueryResponseBuilder) Build() *RunQueryResponse {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RunQueryResponseReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusRequest

// reader

type GetTransactionStatusRequest struct {
	// TransactionRef TransactionRef

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *GetTransactionStatusRequest) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionRef:%s,}", x.StringTransactionRef())
}

var _GetTransactionStatusRequest_Scheme = []membuffers.FieldType{membuffers.TypeMessage}
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

func (x *GetTransactionStatusRequest) Equal(y *GetTransactionStatusRequest) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *GetTransactionStatusRequest) TransactionRef() *TransactionRef {
	b, s := x._message.GetMessage(0)
	return TransactionRefReader(b[:s])
}

func (x *GetTransactionStatusRequest) RawTransactionRef() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetTransactionStatusRequest) RawTransactionRefWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *GetTransactionStatusRequest) StringTransactionRef() string {
	return x.TransactionRef().String()
}

// builder

type GetTransactionStatusRequestBuilder struct {
	TransactionRef *TransactionRefBuilder

	// internal
	// implements membuffers.Builder
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
	err = w._builder.WriteMessage(buf, w.TransactionRef)
	if err != nil {
		return
	}
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
	// RequestResult RequestResult
	// TransactionStatus protocol.TransactionStatus
	// TransactionReceipt protocol.TransactionReceipt

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *GetTransactionStatusResponse) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RequestResult:%s,TransactionStatus:%s,TransactionReceipt:%s,}", x.StringRequestResult(), x.StringTransactionStatus(), x.StringTransactionReceipt())
}

var _GetTransactionStatusResponse_Scheme = []membuffers.FieldType{membuffers.TypeMessage, membuffers.TypeUint16, membuffers.TypeMessage}
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

func (x *GetTransactionStatusResponse) Equal(y *GetTransactionStatusResponse) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *GetTransactionStatusResponse) RequestResult() *RequestResult {
	b, s := x._message.GetMessage(0)
	return RequestResultReader(b[:s])
}

func (x *GetTransactionStatusResponse) RawRequestResult() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetTransactionStatusResponse) RawRequestResultWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *GetTransactionStatusResponse) StringRequestResult() string {
	return x.RequestResult().String()
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

func (x *GetTransactionStatusResponse) TransactionReceipt() *protocol.TransactionReceipt {
	b, s := x._message.GetMessage(2)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *GetTransactionStatusResponse) RawTransactionReceipt() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *GetTransactionStatusResponse) RawTransactionReceiptWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *GetTransactionStatusResponse) StringTransactionReceipt() string {
	return x.TransactionReceipt().String()
}

// builder

type GetTransactionStatusResponseBuilder struct {
	RequestResult      *RequestResultBuilder
	TransactionStatus  protocol.TransactionStatus
	TransactionReceipt *protocol.TransactionReceiptBuilder

	// internal
	// implements membuffers.Builder
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
	err = w._builder.WriteMessage(buf, w.RequestResult)
	if err != nil {
		return
	}
	w._builder.WriteUint16(buf, uint16(w.TransactionStatus))
	err = w._builder.WriteMessage(buf, w.TransactionReceipt)
	if err != nil {
		return
	}
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
// message GetTransactionProofRequest

// reader

type GetTransactionProofRequest struct {
	// TransactionRef TransactionRef

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *GetTransactionProofRequest) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionRef:%s,}", x.StringTransactionRef())
}

var _GetTransactionProofRequest_Scheme = []membuffers.FieldType{membuffers.TypeMessage}
var _GetTransactionProofRequest_Unions = [][]membuffers.FieldType{}

func GetTransactionProofRequestReader(buf []byte) *GetTransactionProofRequest {
	x := &GetTransactionProofRequest{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _GetTransactionProofRequest_Scheme, _GetTransactionProofRequest_Unions)
	return x
}

func (x *GetTransactionProofRequest) IsValid() bool {
	return x._message.IsValid()
}

func (x *GetTransactionProofRequest) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *GetTransactionProofRequest) Equal(y *GetTransactionProofRequest) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *GetTransactionProofRequest) TransactionRef() *TransactionRef {
	b, s := x._message.GetMessage(0)
	return TransactionRefReader(b[:s])
}

func (x *GetTransactionProofRequest) RawTransactionRef() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetTransactionProofRequest) RawTransactionRefWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *GetTransactionProofRequest) StringTransactionRef() string {
	return x.TransactionRef().String()
}

// builder

type GetTransactionProofRequestBuilder struct {
	TransactionRef *TransactionRefBuilder

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *GetTransactionProofRequestBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.TransactionRef)
	if err != nil {
		return
	}
	return nil
}

func (w *GetTransactionProofRequestBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *GetTransactionProofRequestBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *GetTransactionProofRequestBuilder) Build() *GetTransactionProofRequest {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionProofRequestReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionProofResponse

// reader

type GetTransactionProofResponse struct {
	// RequestResult RequestResult
	// TransactionStatus protocol.TransactionStatus
	// TransactionReceipt protocol.TransactionReceipt
	// Proof primitives.PackedReceiptProof

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *GetTransactionProofResponse) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RequestResult:%s,TransactionStatus:%s,TransactionReceipt:%s,Proof:%s,}", x.StringRequestResult(), x.StringTransactionStatus(), x.StringTransactionReceipt(), x.StringProof())
}

var _GetTransactionProofResponse_Scheme = []membuffers.FieldType{membuffers.TypeMessage, membuffers.TypeUint16, membuffers.TypeMessage, membuffers.TypeBytes}
var _GetTransactionProofResponse_Unions = [][]membuffers.FieldType{}

func GetTransactionProofResponseReader(buf []byte) *GetTransactionProofResponse {
	x := &GetTransactionProofResponse{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _GetTransactionProofResponse_Scheme, _GetTransactionProofResponse_Unions)
	return x
}

func (x *GetTransactionProofResponse) IsValid() bool {
	return x._message.IsValid()
}

func (x *GetTransactionProofResponse) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *GetTransactionProofResponse) Equal(y *GetTransactionProofResponse) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *GetTransactionProofResponse) RequestResult() *RequestResult {
	b, s := x._message.GetMessage(0)
	return RequestResultReader(b[:s])
}

func (x *GetTransactionProofResponse) RawRequestResult() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetTransactionProofResponse) RawRequestResultWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *GetTransactionProofResponse) StringRequestResult() string {
	return x.RequestResult().String()
}

func (x *GetTransactionProofResponse) TransactionStatus() protocol.TransactionStatus {
	return protocol.TransactionStatus(x._message.GetUint16(1))
}

func (x *GetTransactionProofResponse) RawTransactionStatus() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *GetTransactionProofResponse) MutateTransactionStatus(v protocol.TransactionStatus) error {
	return x._message.SetUint16(1, uint16(v))
}

func (x *GetTransactionProofResponse) StringTransactionStatus() string {
	return x.TransactionStatus().String()
}

func (x *GetTransactionProofResponse) TransactionReceipt() *protocol.TransactionReceipt {
	b, s := x._message.GetMessage(2)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *GetTransactionProofResponse) RawTransactionReceipt() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *GetTransactionProofResponse) RawTransactionReceiptWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *GetTransactionProofResponse) StringTransactionReceipt() string {
	return x.TransactionReceipt().String()
}

func (x *GetTransactionProofResponse) Proof() primitives.PackedReceiptProof {
	return primitives.PackedReceiptProof(x._message.GetBytes(3))
}

func (x *GetTransactionProofResponse) RawProof() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *GetTransactionProofResponse) MutateProof(v primitives.PackedReceiptProof) error {
	return x._message.SetBytes(3, []byte(v))
}

func (x *GetTransactionProofResponse) StringProof() string {
	return fmt.Sprintf("%s", x.Proof())
}

// builder

type GetTransactionProofResponseBuilder struct {
	RequestResult      *RequestResultBuilder
	TransactionStatus  protocol.TransactionStatus
	TransactionReceipt *protocol.TransactionReceiptBuilder
	Proof              primitives.PackedReceiptProof

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *GetTransactionProofResponseBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.RequestResult)
	if err != nil {
		return
	}
	w._builder.WriteUint16(buf, uint16(w.TransactionStatus))
	err = w._builder.WriteMessage(buf, w.TransactionReceipt)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, []byte(w.Proof))
	return nil
}

func (w *GetTransactionProofResponseBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *GetTransactionProofResponseBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *GetTransactionProofResponseBuilder) Build() *GetTransactionProofResponse {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionProofResponseReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionRef

// reader

type TransactionRef struct {
	// ProtocolVersion primitives.ProtocolVersion
	// VirtualChainId primitives.VirtualChainId
	// TransactionTimestamp primitives.TimestampNano
	// Txhash primitives.Sha256

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *TransactionRef) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ProtocolVersion:%s,VirtualChainId:%s,TransactionTimestamp:%s,Txhash:%s,}", x.StringProtocolVersion(), x.StringVirtualChainId(), x.StringTransactionTimestamp(), x.StringTxhash())
}

var _TransactionRef_Scheme = []membuffers.FieldType{membuffers.TypeUint32, membuffers.TypeUint32, membuffers.TypeUint64, membuffers.TypeBytes}
var _TransactionRef_Unions = [][]membuffers.FieldType{}

func TransactionRefReader(buf []byte) *TransactionRef {
	x := &TransactionRef{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _TransactionRef_Scheme, _TransactionRef_Unions)
	return x
}

func (x *TransactionRef) IsValid() bool {
	return x._message.IsValid()
}

func (x *TransactionRef) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *TransactionRef) Equal(y *TransactionRef) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *TransactionRef) ProtocolVersion() primitives.ProtocolVersion {
	return primitives.ProtocolVersion(x._message.GetUint32(0))
}

func (x *TransactionRef) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *TransactionRef) MutateProtocolVersion(v primitives.ProtocolVersion) error {
	return x._message.SetUint32(0, uint32(v))
}

func (x *TransactionRef) StringProtocolVersion() string {
	return fmt.Sprintf("%s", x.ProtocolVersion())
}

func (x *TransactionRef) VirtualChainId() primitives.VirtualChainId {
	return primitives.VirtualChainId(x._message.GetUint32(1))
}

func (x *TransactionRef) RawVirtualChainId() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *TransactionRef) MutateVirtualChainId(v primitives.VirtualChainId) error {
	return x._message.SetUint32(1, uint32(v))
}

func (x *TransactionRef) StringVirtualChainId() string {
	return fmt.Sprintf("%s", x.VirtualChainId())
}

func (x *TransactionRef) TransactionTimestamp() primitives.TimestampNano {
	return primitives.TimestampNano(x._message.GetUint64(2))
}

func (x *TransactionRef) RawTransactionTimestamp() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *TransactionRef) MutateTransactionTimestamp(v primitives.TimestampNano) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *TransactionRef) StringTransactionTimestamp() string {
	return fmt.Sprintf("%s", x.TransactionTimestamp())
}

func (x *TransactionRef) Txhash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(3))
}

func (x *TransactionRef) RawTxhash() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *TransactionRef) MutateTxhash(v primitives.Sha256) error {
	return x._message.SetBytes(3, []byte(v))
}

func (x *TransactionRef) StringTxhash() string {
	return fmt.Sprintf("%s", x.Txhash())
}

// builder

type TransactionRefBuilder struct {
	ProtocolVersion      primitives.ProtocolVersion
	VirtualChainId       primitives.VirtualChainId
	TransactionTimestamp primitives.TimestampNano
	Txhash               primitives.Sha256

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *TransactionRefBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUint64(buf, uint64(w.TransactionTimestamp))
	w._builder.WriteBytes(buf, []byte(w.Txhash))
	return nil
}

func (w *TransactionRefBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TransactionRefBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *TransactionRefBuilder) Build() *TransactionRef {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionRefReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message RequestResult

// reader

type RequestResult struct {
	// RequestStatus protocol.RequestStatus
	// BlockHeight primitives.BlockHeight
	// BlockTimestamp primitives.TimestampNano

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *RequestResult) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RequestStatus:%s,BlockHeight:%s,BlockTimestamp:%s,}", x.StringRequestStatus(), x.StringBlockHeight(), x.StringBlockTimestamp())
}

var _RequestResult_Scheme = []membuffers.FieldType{membuffers.TypeUint16, membuffers.TypeUint64, membuffers.TypeUint64}
var _RequestResult_Unions = [][]membuffers.FieldType{}

func RequestResultReader(buf []byte) *RequestResult {
	x := &RequestResult{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _RequestResult_Scheme, _RequestResult_Unions)
	return x
}

func (x *RequestResult) IsValid() bool {
	return x._message.IsValid()
}

func (x *RequestResult) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *RequestResult) Equal(y *RequestResult) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *RequestResult) RequestStatus() protocol.RequestStatus {
	return protocol.RequestStatus(x._message.GetUint16(0))
}

func (x *RequestResult) RawRequestStatus() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *RequestResult) MutateRequestStatus(v protocol.RequestStatus) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *RequestResult) StringRequestStatus() string {
	return x.RequestStatus().String()
}

func (x *RequestResult) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *RequestResult) RawBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *RequestResult) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *RequestResult) StringBlockHeight() string {
	return fmt.Sprintf("%s", x.BlockHeight())
}

func (x *RequestResult) BlockTimestamp() primitives.TimestampNano {
	return primitives.TimestampNano(x._message.GetUint64(2))
}

func (x *RequestResult) RawBlockTimestamp() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *RequestResult) MutateBlockTimestamp(v primitives.TimestampNano) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *RequestResult) StringBlockTimestamp() string {
	return fmt.Sprintf("%s", x.BlockTimestamp())
}

// builder

type RequestResultBuilder struct {
	RequestStatus  protocol.RequestStatus
	BlockHeight    primitives.BlockHeight
	BlockTimestamp primitives.TimestampNano

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *RequestResultBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint16(buf, uint16(w.RequestStatus))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteUint64(buf, uint64(w.BlockTimestamp))
	return nil
}

func (w *RequestResultBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *RequestResultBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *RequestResultBuilder) Build() *RequestResult {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RequestResultReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums
