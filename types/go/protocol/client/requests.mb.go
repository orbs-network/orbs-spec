// AUTO GENERATED FILE (by membufc proto compiler v0.3.6)
package client

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

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
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *RequestResultBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUint16(buf, uint16(w.RequestStatus))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteUint64(buf, uint64(w.BlockTimestamp))
	return nil
}

func (w *RequestResultBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint16(prefix, offsetFromStart, "RequestResult.RequestStatus", uint16(w.RequestStatus))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "RequestResult.BlockHeight", uint64(w.BlockHeight))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "RequestResult.BlockTimestamp", uint64(w.BlockTimestamp))
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

func RequestResultBuilderFromRaw(raw []byte) *RequestResultBuilder {
	return &RequestResultBuilder{_overrideWithRawBuffer: raw}
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

func (x *TransactionRef) RawTxhashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
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
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *TransactionRefBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUint64(buf, uint64(w.TransactionTimestamp))
	w._builder.WriteBytes(buf, []byte(w.Txhash))
	return nil
}

func (w *TransactionRefBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint32(prefix, offsetFromStart, "TransactionRef.ProtocolVersion", uint32(w.ProtocolVersion))
	w._builder.HexDumpUint32(prefix, offsetFromStart, "TransactionRef.VirtualChainId", uint32(w.VirtualChainId))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "TransactionRef.TransactionTimestamp", uint64(w.TransactionTimestamp))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "TransactionRef.Txhash", []byte(w.Txhash))
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

func TransactionRefBuilderFromRaw(raw []byte) *TransactionRefBuilder {
	return &TransactionRefBuilder{_overrideWithRawBuffer: raw}
}

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
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *SendTransactionRequestBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessage(buf, w.SignedTransaction)
	if err != nil {
		return
	}
	return nil
}

func (w *SendTransactionRequestBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "SendTransactionRequest.SignedTransaction", w.SignedTransaction)
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

func SendTransactionRequestBuilderFromRaw(raw []byte) *SendTransactionRequestBuilder {
	return &SendTransactionRequestBuilder{_overrideWithRawBuffer: raw}
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
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *SendTransactionResponseBuilder) Write(buf []byte) (err error) {
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

func (w *SendTransactionResponseBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "SendTransactionResponse.RequestResult", w.RequestResult)
	if err != nil {
		return
	}
	w._builder.HexDumpUint16(prefix, offsetFromStart, "SendTransactionResponse.TransactionStatus", uint16(w.TransactionStatus))
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "SendTransactionResponse.TransactionReceipt", w.TransactionReceipt)
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

func SendTransactionResponseBuilderFromRaw(raw []byte) *SendTransactionResponseBuilder {
	return &SendTransactionResponseBuilder{_overrideWithRawBuffer: raw}
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
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *RunQueryRequestBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessage(buf, w.SignedQuery)
	if err != nil {
		return
	}
	return nil
}

func (w *RunQueryRequestBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "RunQueryRequest.SignedQuery", w.SignedQuery)
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

func RunQueryRequestBuilderFromRaw(raw []byte) *RunQueryRequestBuilder {
	return &RunQueryRequestBuilder{_overrideWithRawBuffer: raw}
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
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *RunQueryResponseBuilder) Write(buf []byte) (err error) {
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

func (w *RunQueryResponseBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "RunQueryResponse.RequestResult", w.RequestResult)
	if err != nil {
		return
	}
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "RunQueryResponse.QueryResult", w.QueryResult)
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

func RunQueryResponseBuilderFromRaw(raw []byte) *RunQueryResponseBuilder {
	return &RunQueryResponseBuilder{_overrideWithRawBuffer: raw}
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
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *GetTransactionStatusRequestBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessage(buf, w.TransactionRef)
	if err != nil {
		return
	}
	return nil
}

func (w *GetTransactionStatusRequestBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetTransactionStatusRequest.TransactionRef", w.TransactionRef)
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

func GetTransactionStatusRequestBuilderFromRaw(raw []byte) *GetTransactionStatusRequestBuilder {
	return &GetTransactionStatusRequestBuilder{_overrideWithRawBuffer: raw}
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
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *GetTransactionStatusResponseBuilder) Write(buf []byte) (err error) {
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

func (w *GetTransactionStatusResponseBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetTransactionStatusResponse.RequestResult", w.RequestResult)
	if err != nil {
		return
	}
	w._builder.HexDumpUint16(prefix, offsetFromStart, "GetTransactionStatusResponse.TransactionStatus", uint16(w.TransactionStatus))
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetTransactionStatusResponse.TransactionReceipt", w.TransactionReceipt)
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

func GetTransactionStatusResponseBuilderFromRaw(raw []byte) *GetTransactionStatusResponseBuilder {
	return &GetTransactionStatusResponseBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionReceiptProofRequest

// reader

type GetTransactionReceiptProofRequest struct {
	// TransactionRef TransactionRef

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *GetTransactionReceiptProofRequest) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionRef:%s,}", x.StringTransactionRef())
}

var _GetTransactionReceiptProofRequest_Scheme = []membuffers.FieldType{membuffers.TypeMessage}
var _GetTransactionReceiptProofRequest_Unions = [][]membuffers.FieldType{}

func GetTransactionReceiptProofRequestReader(buf []byte) *GetTransactionReceiptProofRequest {
	x := &GetTransactionReceiptProofRequest{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _GetTransactionReceiptProofRequest_Scheme, _GetTransactionReceiptProofRequest_Unions)
	return x
}

func (x *GetTransactionReceiptProofRequest) IsValid() bool {
	return x._message.IsValid()
}

func (x *GetTransactionReceiptProofRequest) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *GetTransactionReceiptProofRequest) Equal(y *GetTransactionReceiptProofRequest) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *GetTransactionReceiptProofRequest) TransactionRef() *TransactionRef {
	b, s := x._message.GetMessage(0)
	return TransactionRefReader(b[:s])
}

func (x *GetTransactionReceiptProofRequest) RawTransactionRef() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetTransactionReceiptProofRequest) RawTransactionRefWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *GetTransactionReceiptProofRequest) StringTransactionRef() string {
	return x.TransactionRef().String()
}

// builder

type GetTransactionReceiptProofRequestBuilder struct {
	TransactionRef *TransactionRefBuilder

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *GetTransactionReceiptProofRequestBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessage(buf, w.TransactionRef)
	if err != nil {
		return
	}
	return nil
}

func (w *GetTransactionReceiptProofRequestBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetTransactionReceiptProofRequest.TransactionRef", w.TransactionRef)
	if err != nil {
		return
	}
	return nil
}

func (w *GetTransactionReceiptProofRequestBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *GetTransactionReceiptProofRequestBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *GetTransactionReceiptProofRequestBuilder) Build() *GetTransactionReceiptProofRequest {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionReceiptProofRequestReader(buf)
}

func GetTransactionReceiptProofRequestBuilderFromRaw(raw []byte) *GetTransactionReceiptProofRequestBuilder {
	return &GetTransactionReceiptProofRequestBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionReceiptProofResponse

// reader

type GetTransactionReceiptProofResponse struct {
	// RequestResult RequestResult
	// TransactionStatus protocol.TransactionStatus
	// TransactionReceipt protocol.TransactionReceipt
	// PackedProof primitives.PackedReceiptProof

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *GetTransactionReceiptProofResponse) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RequestResult:%s,TransactionStatus:%s,TransactionReceipt:%s,PackedProof:%s,}", x.StringRequestResult(), x.StringTransactionStatus(), x.StringTransactionReceipt(), x.StringPackedProof())
}

var _GetTransactionReceiptProofResponse_Scheme = []membuffers.FieldType{membuffers.TypeMessage, membuffers.TypeUint16, membuffers.TypeMessage, membuffers.TypeBytes}
var _GetTransactionReceiptProofResponse_Unions = [][]membuffers.FieldType{}

func GetTransactionReceiptProofResponseReader(buf []byte) *GetTransactionReceiptProofResponse {
	x := &GetTransactionReceiptProofResponse{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _GetTransactionReceiptProofResponse_Scheme, _GetTransactionReceiptProofResponse_Unions)
	return x
}

func (x *GetTransactionReceiptProofResponse) IsValid() bool {
	return x._message.IsValid()
}

func (x *GetTransactionReceiptProofResponse) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *GetTransactionReceiptProofResponse) Equal(y *GetTransactionReceiptProofResponse) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *GetTransactionReceiptProofResponse) RequestResult() *RequestResult {
	b, s := x._message.GetMessage(0)
	return RequestResultReader(b[:s])
}

func (x *GetTransactionReceiptProofResponse) RawRequestResult() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetTransactionReceiptProofResponse) RawRequestResultWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *GetTransactionReceiptProofResponse) StringRequestResult() string {
	return x.RequestResult().String()
}

func (x *GetTransactionReceiptProofResponse) TransactionStatus() protocol.TransactionStatus {
	return protocol.TransactionStatus(x._message.GetUint16(1))
}

func (x *GetTransactionReceiptProofResponse) RawTransactionStatus() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *GetTransactionReceiptProofResponse) MutateTransactionStatus(v protocol.TransactionStatus) error {
	return x._message.SetUint16(1, uint16(v))
}

func (x *GetTransactionReceiptProofResponse) StringTransactionStatus() string {
	return x.TransactionStatus().String()
}

func (x *GetTransactionReceiptProofResponse) TransactionReceipt() *protocol.TransactionReceipt {
	b, s := x._message.GetMessage(2)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *GetTransactionReceiptProofResponse) RawTransactionReceipt() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *GetTransactionReceiptProofResponse) RawTransactionReceiptWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *GetTransactionReceiptProofResponse) StringTransactionReceipt() string {
	return x.TransactionReceipt().String()
}

func (x *GetTransactionReceiptProofResponse) PackedProof() primitives.PackedReceiptProof {
	return primitives.PackedReceiptProof(x._message.GetBytes(3))
}

func (x *GetTransactionReceiptProofResponse) RawPackedProof() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *GetTransactionReceiptProofResponse) RawPackedProofWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *GetTransactionReceiptProofResponse) MutatePackedProof(v primitives.PackedReceiptProof) error {
	return x._message.SetBytes(3, []byte(v))
}

func (x *GetTransactionReceiptProofResponse) StringPackedProof() string {
	return fmt.Sprintf("%s", x.PackedProof())
}

// builder

type GetTransactionReceiptProofResponseBuilder struct {
	RequestResult      *RequestResultBuilder
	TransactionStatus  protocol.TransactionStatus
	TransactionReceipt *protocol.TransactionReceiptBuilder
	PackedProof        primitives.PackedReceiptProof

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *GetTransactionReceiptProofResponseBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessage(buf, w.RequestResult)
	if err != nil {
		return
	}
	w._builder.WriteUint16(buf, uint16(w.TransactionStatus))
	err = w._builder.WriteMessage(buf, w.TransactionReceipt)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, []byte(w.PackedProof))
	return nil
}

func (w *GetTransactionReceiptProofResponseBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetTransactionReceiptProofResponse.RequestResult", w.RequestResult)
	if err != nil {
		return
	}
	w._builder.HexDumpUint16(prefix, offsetFromStart, "GetTransactionReceiptProofResponse.TransactionStatus", uint16(w.TransactionStatus))
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetTransactionReceiptProofResponse.TransactionReceipt", w.TransactionReceipt)
	if err != nil {
		return
	}
	w._builder.HexDumpBytes(prefix, offsetFromStart, "GetTransactionReceiptProofResponse.PackedProof", []byte(w.PackedProof))
	return nil
}

func (w *GetTransactionReceiptProofResponseBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *GetTransactionReceiptProofResponseBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *GetTransactionReceiptProofResponseBuilder) Build() *GetTransactionReceiptProofResponse {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionReceiptProofResponseReader(buf)
}

func GetTransactionReceiptProofResponseBuilderFromRaw(raw []byte) *GetTransactionReceiptProofResponseBuilder {
	return &GetTransactionReceiptProofResponseBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message GetBlockRequest

// reader

type GetBlockRequest struct {
	// ProtocolVersion primitives.ProtocolVersion
	// VirtualChainId primitives.VirtualChainId
	// BlockHeight primitives.BlockHeight

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *GetBlockRequest) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ProtocolVersion:%s,VirtualChainId:%s,BlockHeight:%s,}", x.StringProtocolVersion(), x.StringVirtualChainId(), x.StringBlockHeight())
}

var _GetBlockRequest_Scheme = []membuffers.FieldType{membuffers.TypeUint32, membuffers.TypeUint32, membuffers.TypeUint64}
var _GetBlockRequest_Unions = [][]membuffers.FieldType{}

func GetBlockRequestReader(buf []byte) *GetBlockRequest {
	x := &GetBlockRequest{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _GetBlockRequest_Scheme, _GetBlockRequest_Unions)
	return x
}

func (x *GetBlockRequest) IsValid() bool {
	return x._message.IsValid()
}

func (x *GetBlockRequest) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *GetBlockRequest) Equal(y *GetBlockRequest) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *GetBlockRequest) ProtocolVersion() primitives.ProtocolVersion {
	return primitives.ProtocolVersion(x._message.GetUint32(0))
}

func (x *GetBlockRequest) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetBlockRequest) MutateProtocolVersion(v primitives.ProtocolVersion) error {
	return x._message.SetUint32(0, uint32(v))
}

func (x *GetBlockRequest) StringProtocolVersion() string {
	return fmt.Sprintf("%s", x.ProtocolVersion())
}

func (x *GetBlockRequest) VirtualChainId() primitives.VirtualChainId {
	return primitives.VirtualChainId(x._message.GetUint32(1))
}

func (x *GetBlockRequest) RawVirtualChainId() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *GetBlockRequest) MutateVirtualChainId(v primitives.VirtualChainId) error {
	return x._message.SetUint32(1, uint32(v))
}

func (x *GetBlockRequest) StringVirtualChainId() string {
	return fmt.Sprintf("%s", x.VirtualChainId())
}

func (x *GetBlockRequest) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *GetBlockRequest) RawBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *GetBlockRequest) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *GetBlockRequest) StringBlockHeight() string {
	return fmt.Sprintf("%s", x.BlockHeight())
}

// builder

type GetBlockRequestBuilder struct {
	ProtocolVersion primitives.ProtocolVersion
	VirtualChainId  primitives.VirtualChainId
	BlockHeight     primitives.BlockHeight

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *GetBlockRequestBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	return nil
}

func (w *GetBlockRequestBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint32(prefix, offsetFromStart, "GetBlockRequest.ProtocolVersion", uint32(w.ProtocolVersion))
	w._builder.HexDumpUint32(prefix, offsetFromStart, "GetBlockRequest.VirtualChainId", uint32(w.VirtualChainId))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "GetBlockRequest.BlockHeight", uint64(w.BlockHeight))
	return nil
}

func (w *GetBlockRequestBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *GetBlockRequestBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *GetBlockRequestBuilder) Build() *GetBlockRequest {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetBlockRequestReader(buf)
}

func GetBlockRequestBuilderFromRaw(raw []byte) *GetBlockRequestBuilder {
	return &GetBlockRequestBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message GetBlockResponse

// reader

type GetBlockResponse struct {
	// RequestResult RequestResult
	// TransactionsBlockHeader protocol.TransactionsBlockHeader
	// TransactionsBlockMetadata protocol.TransactionsBlockMetadata
	// SignedTransactions []protocol.SignedTransaction
	// TransactionsBlockProof protocol.TransactionsBlockProof
	// ResultsBlockHeader protocol.ResultsBlockHeader
	// TransactionReceipts []protocol.TransactionReceipt
	// ContractStateDiffs []protocol.ContractStateDiff
	// ResultsBlockProof protocol.ResultsBlockProof

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *GetBlockResponse) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RequestResult:%s,TransactionsBlockHeader:%s,TransactionsBlockMetadata:%s,SignedTransactions:%s,TransactionsBlockProof:%s,ResultsBlockHeader:%s,TransactionReceipts:%s,ContractStateDiffs:%s,ResultsBlockProof:%s,}", x.StringRequestResult(), x.StringTransactionsBlockHeader(), x.StringTransactionsBlockMetadata(), x.StringSignedTransactions(), x.StringTransactionsBlockProof(), x.StringResultsBlockHeader(), x.StringTransactionReceipts(), x.StringContractStateDiffs(), x.StringResultsBlockProof())
}

var _GetBlockResponse_Scheme = []membuffers.FieldType{membuffers.TypeMessage, membuffers.TypeMessage, membuffers.TypeMessage, membuffers.TypeMessageArray, membuffers.TypeMessage, membuffers.TypeMessage, membuffers.TypeMessageArray, membuffers.TypeMessageArray, membuffers.TypeMessage}
var _GetBlockResponse_Unions = [][]membuffers.FieldType{}

func GetBlockResponseReader(buf []byte) *GetBlockResponse {
	x := &GetBlockResponse{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _GetBlockResponse_Scheme, _GetBlockResponse_Unions)
	return x
}

func (x *GetBlockResponse) IsValid() bool {
	return x._message.IsValid()
}

func (x *GetBlockResponse) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *GetBlockResponse) Equal(y *GetBlockResponse) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *GetBlockResponse) RequestResult() *RequestResult {
	b, s := x._message.GetMessage(0)
	return RequestResultReader(b[:s])
}

func (x *GetBlockResponse) RawRequestResult() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GetBlockResponse) RawRequestResultWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *GetBlockResponse) StringRequestResult() string {
	return x.RequestResult().String()
}

func (x *GetBlockResponse) TransactionsBlockHeader() *protocol.TransactionsBlockHeader {
	b, s := x._message.GetMessage(1)
	return protocol.TransactionsBlockHeaderReader(b[:s])
}

func (x *GetBlockResponse) RawTransactionsBlockHeader() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *GetBlockResponse) RawTransactionsBlockHeaderWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *GetBlockResponse) StringTransactionsBlockHeader() string {
	return x.TransactionsBlockHeader().String()
}

func (x *GetBlockResponse) TransactionsBlockMetadata() *protocol.TransactionsBlockMetadata {
	b, s := x._message.GetMessage(2)
	return protocol.TransactionsBlockMetadataReader(b[:s])
}

func (x *GetBlockResponse) RawTransactionsBlockMetadata() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *GetBlockResponse) RawTransactionsBlockMetadataWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *GetBlockResponse) StringTransactionsBlockMetadata() string {
	return x.TransactionsBlockMetadata().String()
}

func (x *GetBlockResponse) SignedTransactionsIterator() *GetBlockResponseSignedTransactionsIterator {
	return &GetBlockResponseSignedTransactionsIterator{iterator: x._message.GetMessageArrayIterator(3)}
}

type GetBlockResponseSignedTransactionsIterator struct {
	iterator *membuffers.Iterator
}

func (i *GetBlockResponseSignedTransactionsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *GetBlockResponseSignedTransactionsIterator) NextSignedTransactions() *protocol.SignedTransaction {
	b, s := i.iterator.NextMessage()
	return protocol.SignedTransactionReader(b[:s])
}

func (x *GetBlockResponse) RawSignedTransactionsArray() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *GetBlockResponse) RawSignedTransactionsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *GetBlockResponse) StringSignedTransactions() (res string) {
	res = "["
	for i := x.SignedTransactionsIterator(); i.HasNext(); {
		res += i.NextSignedTransactions().String() + ","
	}
	res += "]"
	return
}

func (x *GetBlockResponse) TransactionsBlockProof() *protocol.TransactionsBlockProof {
	b, s := x._message.GetMessage(4)
	return protocol.TransactionsBlockProofReader(b[:s])
}

func (x *GetBlockResponse) RawTransactionsBlockProof() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *GetBlockResponse) RawTransactionsBlockProofWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(4, 0)
}

func (x *GetBlockResponse) StringTransactionsBlockProof() string {
	return x.TransactionsBlockProof().String()
}

func (x *GetBlockResponse) ResultsBlockHeader() *protocol.ResultsBlockHeader {
	b, s := x._message.GetMessage(5)
	return protocol.ResultsBlockHeaderReader(b[:s])
}

func (x *GetBlockResponse) RawResultsBlockHeader() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *GetBlockResponse) RawResultsBlockHeaderWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(5, 0)
}

func (x *GetBlockResponse) StringResultsBlockHeader() string {
	return x.ResultsBlockHeader().String()
}

func (x *GetBlockResponse) TransactionReceiptsIterator() *GetBlockResponseTransactionReceiptsIterator {
	return &GetBlockResponseTransactionReceiptsIterator{iterator: x._message.GetMessageArrayIterator(6)}
}

type GetBlockResponseTransactionReceiptsIterator struct {
	iterator *membuffers.Iterator
}

func (i *GetBlockResponseTransactionReceiptsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *GetBlockResponseTransactionReceiptsIterator) NextTransactionReceipts() *protocol.TransactionReceipt {
	b, s := i.iterator.NextMessage()
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *GetBlockResponse) RawTransactionReceiptsArray() []byte {
	return x._message.RawBufferForField(6, 0)
}

func (x *GetBlockResponse) RawTransactionReceiptsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(6, 0)
}

func (x *GetBlockResponse) StringTransactionReceipts() (res string) {
	res = "["
	for i := x.TransactionReceiptsIterator(); i.HasNext(); {
		res += i.NextTransactionReceipts().String() + ","
	}
	res += "]"
	return
}

func (x *GetBlockResponse) ContractStateDiffsIterator() *GetBlockResponseContractStateDiffsIterator {
	return &GetBlockResponseContractStateDiffsIterator{iterator: x._message.GetMessageArrayIterator(7)}
}

type GetBlockResponseContractStateDiffsIterator struct {
	iterator *membuffers.Iterator
}

func (i *GetBlockResponseContractStateDiffsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *GetBlockResponseContractStateDiffsIterator) NextContractStateDiffs() *protocol.ContractStateDiff {
	b, s := i.iterator.NextMessage()
	return protocol.ContractStateDiffReader(b[:s])
}

func (x *GetBlockResponse) RawContractStateDiffsArray() []byte {
	return x._message.RawBufferForField(7, 0)
}

func (x *GetBlockResponse) RawContractStateDiffsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(7, 0)
}

func (x *GetBlockResponse) StringContractStateDiffs() (res string) {
	res = "["
	for i := x.ContractStateDiffsIterator(); i.HasNext(); {
		res += i.NextContractStateDiffs().String() + ","
	}
	res += "]"
	return
}

func (x *GetBlockResponse) ResultsBlockProof() *protocol.ResultsBlockProof {
	b, s := x._message.GetMessage(8)
	return protocol.ResultsBlockProofReader(b[:s])
}

func (x *GetBlockResponse) RawResultsBlockProof() []byte {
	return x._message.RawBufferForField(8, 0)
}

func (x *GetBlockResponse) RawResultsBlockProofWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(8, 0)
}

func (x *GetBlockResponse) StringResultsBlockProof() string {
	return x.ResultsBlockProof().String()
}

// builder

type GetBlockResponseBuilder struct {
	RequestResult             *RequestResultBuilder
	TransactionsBlockHeader   *protocol.TransactionsBlockHeaderBuilder
	TransactionsBlockMetadata *protocol.TransactionsBlockMetadataBuilder
	SignedTransactions        []*protocol.SignedTransactionBuilder
	TransactionsBlockProof    *protocol.TransactionsBlockProofBuilder
	ResultsBlockHeader        *protocol.ResultsBlockHeaderBuilder
	TransactionReceipts       []*protocol.TransactionReceiptBuilder
	ContractStateDiffs        []*protocol.ContractStateDiffBuilder
	ResultsBlockProof         *protocol.ResultsBlockProofBuilder

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *GetBlockResponseBuilder) arrayOfSignedTransactions() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.SignedTransactions))
	for i, v := range w.SignedTransactions {
		res[i] = v
	}
	return res
}

func (w *GetBlockResponseBuilder) arrayOfTransactionReceipts() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.TransactionReceipts))
	for i, v := range w.TransactionReceipts {
		res[i] = v
	}
	return res
}

func (w *GetBlockResponseBuilder) arrayOfContractStateDiffs() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.ContractStateDiffs))
	for i, v := range w.ContractStateDiffs {
		res[i] = v
	}
	return res
}

func (w *GetBlockResponseBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessage(buf, w.RequestResult)
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.TransactionsBlockHeader)
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.TransactionsBlockMetadata)
	if err != nil {
		return
	}
	err = w._builder.WriteMessageArray(buf, w.arrayOfSignedTransactions())
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.TransactionsBlockProof)
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.ResultsBlockHeader)
	if err != nil {
		return
	}
	err = w._builder.WriteMessageArray(buf, w.arrayOfTransactionReceipts())
	if err != nil {
		return
	}
	err = w._builder.WriteMessageArray(buf, w.arrayOfContractStateDiffs())
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.ResultsBlockProof)
	if err != nil {
		return
	}
	return nil
}

func (w *GetBlockResponseBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetBlockResponse.RequestResult", w.RequestResult)
	if err != nil {
		return
	}
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetBlockResponse.TransactionsBlockHeader", w.TransactionsBlockHeader)
	if err != nil {
		return
	}
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetBlockResponse.TransactionsBlockMetadata", w.TransactionsBlockMetadata)
	if err != nil {
		return
	}
	err = w._builder.HexDumpMessageArray(prefix, offsetFromStart, "GetBlockResponse.SignedTransactions", w.arrayOfSignedTransactions())
	if err != nil {
		return
	}
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetBlockResponse.TransactionsBlockProof", w.TransactionsBlockProof)
	if err != nil {
		return
	}
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetBlockResponse.ResultsBlockHeader", w.ResultsBlockHeader)
	if err != nil {
		return
	}
	err = w._builder.HexDumpMessageArray(prefix, offsetFromStart, "GetBlockResponse.TransactionReceipts", w.arrayOfTransactionReceipts())
	if err != nil {
		return
	}
	err = w._builder.HexDumpMessageArray(prefix, offsetFromStart, "GetBlockResponse.ContractStateDiffs", w.arrayOfContractStateDiffs())
	if err != nil {
		return
	}
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "GetBlockResponse.ResultsBlockProof", w.ResultsBlockProof)
	if err != nil {
		return
	}
	return nil
}

func (w *GetBlockResponseBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *GetBlockResponseBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *GetBlockResponseBuilder) Build() *GetBlockResponse {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetBlockResponseReader(buf)
}

func GetBlockResponseBuilderFromRaw(raw []byte) *GetBlockResponseBuilder {
	return &GetBlockResponseBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// enums
