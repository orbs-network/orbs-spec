// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package protocol

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message IndexedEvent

// reader

type IndexedEvent struct {
	// ContractName primitives.ContractName
	// EventName string
	// BlockHeight primitives.BlockHeight
	// Timestamp primitives.TimestampNano
	// Txhash primitives.Sha256
	// ExecutionResult ExecutionResult
	// Index uint32
	// Arguments primitives.PackedArgumentArray

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *IndexedEvent) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContractName:%s,EventName:%s,BlockHeight:%s,Timestamp:%s,Txhash:%s,ExecutionResult:%s,Index:%s,Arguments:%s,}", x.StringContractName(), x.StringEventName(), x.StringBlockHeight(), x.StringTimestamp(), x.StringTxhash(), x.StringExecutionResult(), x.StringIndex(), x.StringArguments())
}

var _IndexedEvent_Scheme = []membuffers.FieldType{membuffers.TypeString, membuffers.TypeString, membuffers.TypeUint64, membuffers.TypeUint64, membuffers.TypeBytes, membuffers.TypeUint16, membuffers.TypeUint32, membuffers.TypeBytes}
var _IndexedEvent_Unions = [][]membuffers.FieldType{}

func IndexedEventReader(buf []byte) *IndexedEvent {
	x := &IndexedEvent{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _IndexedEvent_Scheme, _IndexedEvent_Unions)
	return x
}

func (x *IndexedEvent) IsValid() bool {
	return x._message.IsValid()
}

func (x *IndexedEvent) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *IndexedEvent) Equal(y *IndexedEvent) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *IndexedEvent) ContractName() primitives.ContractName {
	return primitives.ContractName(x._message.GetString(0))
}

func (x *IndexedEvent) RawContractName() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *IndexedEvent) RawContractNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *IndexedEvent) MutateContractName(v primitives.ContractName) error {
	return x._message.SetString(0, string(v))
}

func (x *IndexedEvent) StringContractName() string {
	return fmt.Sprintf("%s", x.ContractName())
}

func (x *IndexedEvent) EventName() string {
	return x._message.GetString(1)
}

func (x *IndexedEvent) RawEventName() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *IndexedEvent) RawEventNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *IndexedEvent) MutateEventName(v string) error {
	return x._message.SetString(1, v)
}

func (x *IndexedEvent) StringEventName() string {
	return fmt.Sprintf("%s", x.EventName())
}

func (x *IndexedEvent) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *IndexedEvent) RawBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *IndexedEvent) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *IndexedEvent) StringBlockHeight() string {
	return fmt.Sprintf("%s", x.BlockHeight())
}

func (x *IndexedEvent) Timestamp() primitives.TimestampNano {
	return primitives.TimestampNano(x._message.GetUint64(3))
}

func (x *IndexedEvent) RawTimestamp() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *IndexedEvent) MutateTimestamp(v primitives.TimestampNano) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *IndexedEvent) StringTimestamp() string {
	return fmt.Sprintf("%s", x.Timestamp())
}

func (x *IndexedEvent) Txhash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(4))
}

func (x *IndexedEvent) RawTxhash() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *IndexedEvent) RawTxhashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(4, 0)
}

func (x *IndexedEvent) MutateTxhash(v primitives.Sha256) error {
	return x._message.SetBytes(4, []byte(v))
}

func (x *IndexedEvent) StringTxhash() string {
	return fmt.Sprintf("%s", x.Txhash())
}

func (x *IndexedEvent) ExecutionResult() ExecutionResult {
	return ExecutionResult(x._message.GetUint16(5))
}

func (x *IndexedEvent) RawExecutionResult() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *IndexedEvent) MutateExecutionResult(v ExecutionResult) error {
	return x._message.SetUint16(5, uint16(v))
}

func (x *IndexedEvent) StringExecutionResult() string {
	return x.ExecutionResult().String()
}

func (x *IndexedEvent) Index() uint32 {
	return x._message.GetUint32(6)
}

func (x *IndexedEvent) RawIndex() []byte {
	return x._message.RawBufferForField(6, 0)
}

func (x *IndexedEvent) MutateIndex(v uint32) error {
	return x._message.SetUint32(6, v)
}

func (x *IndexedEvent) StringIndex() string {
	return fmt.Sprintf("%v", x.Index())
}

func (x *IndexedEvent) Arguments() primitives.PackedArgumentArray {
	return primitives.PackedArgumentArray(x._message.GetBytes(7))
}

func (x *IndexedEvent) RawArguments() []byte {
	return x._message.RawBufferForField(7, 0)
}

func (x *IndexedEvent) RawArgumentsWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(7, 0)
}

func (x *IndexedEvent) MutateArguments(v primitives.PackedArgumentArray) error {
	return x._message.SetBytes(7, []byte(v))
}

func (x *IndexedEvent) StringArguments() string {
	return fmt.Sprintf("%s", x.Arguments())
}

// builder

type IndexedEventBuilder struct {
	ContractName    primitives.ContractName
	EventName       string
	BlockHeight     primitives.BlockHeight
	Timestamp       primitives.TimestampNano
	Txhash          primitives.Sha256
	ExecutionResult ExecutionResult
	Index           uint32
	Arguments       primitives.PackedArgumentArray

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *IndexedEventBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteString(buf, string(w.ContractName))
	w._builder.WriteString(buf, w.EventName)
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteUint64(buf, uint64(w.Timestamp))
	w._builder.WriteBytes(buf, []byte(w.Txhash))
	w._builder.WriteUint16(buf, uint16(w.ExecutionResult))
	w._builder.WriteUint32(buf, w.Index)
	w._builder.WriteBytes(buf, []byte(w.Arguments))
	return nil
}

func (w *IndexedEventBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpString(prefix, offsetFromStart, "IndexedEvent.ContractName", string(w.ContractName))
	w._builder.HexDumpString(prefix, offsetFromStart, "IndexedEvent.EventName", w.EventName)
	w._builder.HexDumpUint64(prefix, offsetFromStart, "IndexedEvent.BlockHeight", uint64(w.BlockHeight))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "IndexedEvent.Timestamp", uint64(w.Timestamp))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "IndexedEvent.Txhash", []byte(w.Txhash))
	w._builder.HexDumpUint16(prefix, offsetFromStart, "IndexedEvent.ExecutionResult", uint16(w.ExecutionResult))
	w._builder.HexDumpUint32(prefix, offsetFromStart, "IndexedEvent.Index", w.Index)
	w._builder.HexDumpBytes(prefix, offsetFromStart, "IndexedEvent.Arguments", []byte(w.Arguments))
	return nil
}

func (w *IndexedEventBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *IndexedEventBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *IndexedEventBuilder) Build() *IndexedEvent {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return IndexedEventReader(buf)
}

func IndexedEventBuilderFromRaw(raw []byte) *IndexedEventBuilder {
	return &IndexedEventBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// enums
