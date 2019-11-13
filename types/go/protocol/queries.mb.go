// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package protocol

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

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
	// InputArgumentArray primitives.PackedArgumentArray

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

func (x *Query) RawContractNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(4, 0)
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

func (x *Query) RawMethodNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(5, 0)
}

func (x *Query) MutateMethodName(v primitives.MethodName) error {
	return x._message.SetString(5, string(v))
}

func (x *Query) StringMethodName() string {
	return fmt.Sprintf("%s", x.MethodName())
}

func (x *Query) InputArgumentArray() primitives.PackedArgumentArray {
	return primitives.PackedArgumentArray(x._message.GetBytes(6))
}

func (x *Query) RawInputArgumentArray() []byte {
	return x._message.RawBufferForField(6, 0)
}

func (x *Query) RawInputArgumentArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(6, 0)
}

func (x *Query) MutateInputArgumentArray(v primitives.PackedArgumentArray) error {
	return x._message.SetBytes(6, []byte(v))
}

func (x *Query) StringInputArgumentArray() string {
	return fmt.Sprintf("%s", x.InputArgumentArray())
}

// builder

type QueryBuilder struct {
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

func (w *QueryBuilder) Write(buf []byte) (err error) {
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

func (w *QueryBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint32(prefix, offsetFromStart, "Query.ProtocolVersion", uint32(w.ProtocolVersion))
	w._builder.HexDumpUint32(prefix, offsetFromStart, "Query.VirtualChainId", uint32(w.VirtualChainId))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "Query.Timestamp", uint64(w.Timestamp))
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "Query.Signer", w.Signer)
	if err != nil {
		return
	}
	w._builder.HexDumpString(prefix, offsetFromStart, "Query.ContractName", string(w.ContractName))
	w._builder.HexDumpString(prefix, offsetFromStart, "Query.MethodName", string(w.MethodName))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "Query.InputArgumentArray", []byte(w.InputArgumentArray))
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

func QueryBuilderFromRaw(raw []byte) *QueryBuilder {
	return &QueryBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message SignedQuery

// reader

type SignedQuery struct {
	// Query Query
	// Signature []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *SignedQuery) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Query:%s,Signature:%s,}", x.StringQuery(), x.StringSignature())
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

func (x *SignedQuery) Query() *Query {
	b, s := x._message.GetMessage(0)
	return QueryReader(b[:s])
}

func (x *SignedQuery) RawQuery() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SignedQuery) RawQueryWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *SignedQuery) StringQuery() string {
	return x.Query().String()
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
	Query     *QueryBuilder
	Signature []byte

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *SignedQueryBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessage(buf, w.Query)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *SignedQueryBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "SignedQuery.Query", w.Query)
	if err != nil {
		return
	}
	w._builder.HexDumpBytes(prefix, offsetFromStart, "SignedQuery.Signature", w.Signature)
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

func SignedQueryBuilderFromRaw(raw []byte) *SignedQueryBuilder {
	return &SignedQueryBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message QueryResult

// reader

type QueryResult struct {
	// ExecutionResult ExecutionResult
	// OutputArgumentArray primitives.PackedArgumentArray
	// OutputEventsArray primitives.PackedEventsArray

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *QueryResult) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ExecutionResult:%s,OutputArgumentArray:%s,OutputEventsArray:%s,}", x.StringExecutionResult(), x.StringOutputArgumentArray(), x.StringOutputEventsArray())
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

func (x *QueryResult) OutputArgumentArray() primitives.PackedArgumentArray {
	return primitives.PackedArgumentArray(x._message.GetBytes(1))
}

func (x *QueryResult) RawOutputArgumentArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *QueryResult) RawOutputArgumentArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *QueryResult) MutateOutputArgumentArray(v primitives.PackedArgumentArray) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *QueryResult) StringOutputArgumentArray() string {
	return fmt.Sprintf("%s", x.OutputArgumentArray())
}

func (x *QueryResult) OutputEventsArray() primitives.PackedEventsArray {
	return primitives.PackedEventsArray(x._message.GetBytes(2))
}

func (x *QueryResult) RawOutputEventsArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *QueryResult) RawOutputEventsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *QueryResult) MutateOutputEventsArray(v primitives.PackedEventsArray) error {
	return x._message.SetBytes(2, []byte(v))
}

func (x *QueryResult) StringOutputEventsArray() string {
	return fmt.Sprintf("%s", x.OutputEventsArray())
}

// builder

type QueryResultBuilder struct {
	ExecutionResult     ExecutionResult
	OutputArgumentArray primitives.PackedArgumentArray
	OutputEventsArray   primitives.PackedEventsArray

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *QueryResultBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUint16(buf, uint16(w.ExecutionResult))
	w._builder.WriteBytes(buf, []byte(w.OutputArgumentArray))
	w._builder.WriteBytes(buf, []byte(w.OutputEventsArray))
	return nil
}

func (w *QueryResultBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint16(prefix, offsetFromStart, "QueryResult.ExecutionResult", uint16(w.ExecutionResult))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "QueryResult.OutputArgumentArray", []byte(w.OutputArgumentArray))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "QueryResult.OutputEventsArray", []byte(w.OutputEventsArray))
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

func QueryResultBuilderFromRaw(raw []byte) *QueryResultBuilder {
	return &QueryResultBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// enums
