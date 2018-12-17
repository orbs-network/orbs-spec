// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package protocol

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message MethodArgument

// reader

type MethodArgument struct {
	// Name string
	// Type MethodArgumentType

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *MethodArgument) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Name:%s,Type:%s,}", x.StringName(), x.StringType())
}

var _MethodArgument_Scheme = []membuffers.FieldType{membuffers.TypeString, membuffers.TypeUnion}
var _MethodArgument_Unions = [][]membuffers.FieldType{{membuffers.TypeUint32, membuffers.TypeUint64, membuffers.TypeString, membuffers.TypeBytes}}

func MethodArgumentReader(buf []byte) *MethodArgument {
	x := &MethodArgument{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _MethodArgument_Scheme, _MethodArgument_Unions)
	return x
}

func (x *MethodArgument) IsValid() bool {
	return x._message.IsValid()
}

func (x *MethodArgument) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *MethodArgument) Equal(y *MethodArgument) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *MethodArgument) Name() string {
	return x._message.GetString(0)
}

func (x *MethodArgument) RawName() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *MethodArgument) RawNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *MethodArgument) MutateName(v string) error {
	return x._message.SetString(0, v)
}

func (x *MethodArgument) StringName() string {
	return fmt.Sprintf(x.Name())
}

type MethodArgumentType uint16

const (
	METHOD_ARGUMENT_TYPE_UINT_32_VALUE MethodArgumentType = 0
	METHOD_ARGUMENT_TYPE_UINT_64_VALUE MethodArgumentType = 1
	METHOD_ARGUMENT_TYPE_STRING_VALUE  MethodArgumentType = 2
	METHOD_ARGUMENT_TYPE_BYTES_VALUE   MethodArgumentType = 3
)

func (x *MethodArgument) Type() MethodArgumentType {
	return MethodArgumentType(x._message.GetUnionIndex(1, 0))
}

func (x *MethodArgument) IsTypeUint32Value() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 0)
	return is
}

func (x *MethodArgument) Uint32Value() uint32 {
	is, off := x._message.IsUnionIndex(1, 0, 0)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetUint32InOffset(off)
}

func (x *MethodArgument) StringUint32Value() string {
	return fmt.Sprintf("%x", x.Uint32Value())
}

func (x *MethodArgument) MutateUint32Value(v uint32) error {
	is, off := x._message.IsUnionIndex(1, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint32InOffset(off, v)
	return nil
}

func (x *MethodArgument) IsTypeUint64Value() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 1)
	return is
}

func (x *MethodArgument) Uint64Value() uint64 {
	is, off := x._message.IsUnionIndex(1, 0, 1)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetUint64InOffset(off)
}

func (x *MethodArgument) StringUint64Value() string {
	return fmt.Sprintf("%x", x.Uint64Value())
}

func (x *MethodArgument) MutateUint64Value(v uint64) error {
	is, off := x._message.IsUnionIndex(1, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint64InOffset(off, v)
	return nil
}

func (x *MethodArgument) IsTypeStringValue() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 2)
	return is
}

func (x *MethodArgument) StringValue() string {
	is, off := x._message.IsUnionIndex(1, 0, 2)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetStringInOffset(off)
}

func (x *MethodArgument) StringStringValue() string {
	return fmt.Sprintf(x.StringValue())
}

func (x *MethodArgument) MutateStringValue(v string) error {
	is, off := x._message.IsUnionIndex(1, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetStringInOffset(off, v)
	return nil
}

func (x *MethodArgument) IsTypeBytesValue() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 3)
	return is
}

func (x *MethodArgument) BytesValue() []byte {
	is, off := x._message.IsUnionIndex(1, 0, 3)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetBytesInOffset(off)
}

func (x *MethodArgument) StringBytesValue() string {
	return fmt.Sprintf("%x", x.BytesValue())
}

func (x *MethodArgument) MutateBytesValue(v []byte) error {
	is, off := x._message.IsUnionIndex(1, 0, 3)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetBytesInOffset(off, v)
	return nil
}

func (x *MethodArgument) RawType() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *MethodArgument) RawTypeWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *MethodArgument) StringType() string {
	switch x.Type() {
	case METHOD_ARGUMENT_TYPE_UINT_32_VALUE:
		return "(Uint32Value)" + x.StringUint32Value()
	case METHOD_ARGUMENT_TYPE_UINT_64_VALUE:
		return "(Uint64Value)" + x.StringUint64Value()
	case METHOD_ARGUMENT_TYPE_STRING_VALUE:
		return "(StringValue)" + x.StringStringValue()
	case METHOD_ARGUMENT_TYPE_BYTES_VALUE:
		return "(BytesValue)" + x.StringBytesValue()
	}
	return "(Unknown)"
}

// builder

type MethodArgumentBuilder struct {
	Name        string
	Type        MethodArgumentType
	Uint32Value uint32
	Uint64Value uint64
	StringValue string
	BytesValue  []byte

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *MethodArgumentBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteString(buf, w.Name)
	w._builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case METHOD_ARGUMENT_TYPE_UINT_32_VALUE:
		w._builder.WriteUint32(buf, w.Uint32Value)
	case METHOD_ARGUMENT_TYPE_UINT_64_VALUE:
		w._builder.WriteUint64(buf, w.Uint64Value)
	case METHOD_ARGUMENT_TYPE_STRING_VALUE:
		w._builder.WriteString(buf, w.StringValue)
	case METHOD_ARGUMENT_TYPE_BYTES_VALUE:
		w._builder.WriteBytes(buf, w.BytesValue)
	}
	return nil
}

func (w *MethodArgumentBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpString(prefix, offsetFromStart, "MethodArgument.Name", w.Name)
	w._builder.HexDumpUnionIndex(prefix, offsetFromStart, "MethodArgument.Type", uint16(w.Type))
	switch w.Type {
	case METHOD_ARGUMENT_TYPE_UINT_32_VALUE:
		w._builder.HexDumpUint32(prefix, offsetFromStart, "MethodArgument.Uint32Value", w.Uint32Value)
	case METHOD_ARGUMENT_TYPE_UINT_64_VALUE:
		w._builder.HexDumpUint64(prefix, offsetFromStart, "MethodArgument.Uint64Value", w.Uint64Value)
	case METHOD_ARGUMENT_TYPE_STRING_VALUE:
		w._builder.HexDumpString(prefix, offsetFromStart, "MethodArgument.StringValue", w.StringValue)
	case METHOD_ARGUMENT_TYPE_BYTES_VALUE:
		w._builder.HexDumpBytes(prefix, offsetFromStart, "MethodArgument.BytesValue", w.BytesValue)
	}
	return nil
}

func (w *MethodArgumentBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *MethodArgumentBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *MethodArgumentBuilder) Build() *MethodArgument {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return MethodArgumentReader(buf)
}

func MethodArgumentBuilderFromRaw(raw []byte) *MethodArgumentBuilder {
	return &MethodArgumentBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message MethodArgumentArray

// reader

type MethodArgumentArray struct {
	// Arguments []MethodArgument

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *MethodArgumentArray) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Arguments:%s,}", x.StringArguments())
}

var _MethodArgumentArray_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray}
var _MethodArgumentArray_Unions = [][]membuffers.FieldType{}

func MethodArgumentArrayReader(buf []byte) *MethodArgumentArray {
	x := &MethodArgumentArray{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _MethodArgumentArray_Scheme, _MethodArgumentArray_Unions)
	return x
}

func (x *MethodArgumentArray) IsValid() bool {
	return x._message.IsValid()
}

func (x *MethodArgumentArray) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *MethodArgumentArray) Equal(y *MethodArgumentArray) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *MethodArgumentArray) ArgumentsIterator() *MethodArgumentArrayArgumentsIterator {
	return &MethodArgumentArrayArgumentsIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type MethodArgumentArrayArgumentsIterator struct {
	iterator *membuffers.Iterator
}

func (i *MethodArgumentArrayArgumentsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *MethodArgumentArrayArgumentsIterator) NextArguments() *MethodArgument {
	b, s := i.iterator.NextMessage()
	return MethodArgumentReader(b[:s])
}

func (x *MethodArgumentArray) RawArgumentsArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *MethodArgumentArray) RawArgumentsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *MethodArgumentArray) StringArguments() (res string) {
	res = "["
	for i := x.ArgumentsIterator(); i.HasNext(); {
		res += i.NextArguments().String() + ","
	}
	res += "]"
	return
}

// builder

type MethodArgumentArrayBuilder struct {
	Arguments []*MethodArgumentBuilder

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *MethodArgumentArrayBuilder) arrayOfArguments() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Arguments))
	for i, v := range w.Arguments {
		res[i] = v
	}
	return res
}

func (w *MethodArgumentArrayBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessageArray(buf, w.arrayOfArguments())
	if err != nil {
		return
	}
	return nil
}

func (w *MethodArgumentArrayBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessageArray(prefix, offsetFromStart, "MethodArgumentArray.Arguments", w.arrayOfArguments())
	if err != nil {
		return
	}
	return nil
}

func (w *MethodArgumentArrayBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *MethodArgumentArrayBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *MethodArgumentArrayBuilder) Build() *MethodArgumentArray {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return MethodArgumentArrayReader(buf)
}

func MethodArgumentArrayBuilderFromRaw(raw []byte) *MethodArgumentArrayBuilder {
	return &MethodArgumentArrayBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message StateRecord

// reader

type StateRecord struct {
	// Key primitives.Ripmd160Sha256
	// Value []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *StateRecord) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Key:%s,Value:%s,}", x.StringKey(), x.StringValue())
}

var _StateRecord_Scheme = []membuffers.FieldType{membuffers.TypeBytes, membuffers.TypeBytes}
var _StateRecord_Unions = [][]membuffers.FieldType{}

func StateRecordReader(buf []byte) *StateRecord {
	x := &StateRecord{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _StateRecord_Scheme, _StateRecord_Unions)
	return x
}

func (x *StateRecord) IsValid() bool {
	return x._message.IsValid()
}

func (x *StateRecord) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *StateRecord) Equal(y *StateRecord) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *StateRecord) Key() primitives.Ripmd160Sha256 {
	return primitives.Ripmd160Sha256(x._message.GetBytes(0))
}

func (x *StateRecord) RawKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *StateRecord) RawKeyWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *StateRecord) MutateKey(v primitives.Ripmd160Sha256) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *StateRecord) StringKey() string {
	return fmt.Sprintf("%s", x.Key())
}

func (x *StateRecord) Value() []byte {
	return x._message.GetBytes(1)
}

func (x *StateRecord) RawValue() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *StateRecord) RawValueWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *StateRecord) MutateValue(v []byte) error {
	return x._message.SetBytes(1, v)
}

func (x *StateRecord) StringValue() string {
	return fmt.Sprintf("%x", x.Value())
}

// builder

type StateRecordBuilder struct {
	Key   primitives.Ripmd160Sha256
	Value []byte

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *StateRecordBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, []byte(w.Key))
	w._builder.WriteBytes(buf, w.Value)
	return nil
}

func (w *StateRecordBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "StateRecord.Key", []byte(w.Key))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "StateRecord.Value", w.Value)
	return nil
}

func (w *StateRecordBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *StateRecordBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *StateRecordBuilder) Build() *StateRecord {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return StateRecordReader(buf)
}

func StateRecordBuilderFromRaw(raw []byte) *StateRecordBuilder {
	return &StateRecordBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message ContractStateDiff

// reader

type ContractStateDiff struct {
	// ContractName primitives.ContractName
	// StateDiffs []StateRecord

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *ContractStateDiff) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContractName:%s,StateDiffs:%s,}", x.StringContractName(), x.StringStateDiffs())
}

var _ContractStateDiff_Scheme = []membuffers.FieldType{membuffers.TypeString, membuffers.TypeMessageArray}
var _ContractStateDiff_Unions = [][]membuffers.FieldType{}

func ContractStateDiffReader(buf []byte) *ContractStateDiff {
	x := &ContractStateDiff{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ContractStateDiff_Scheme, _ContractStateDiff_Unions)
	return x
}

func (x *ContractStateDiff) IsValid() bool {
	return x._message.IsValid()
}

func (x *ContractStateDiff) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ContractStateDiff) Equal(y *ContractStateDiff) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *ContractStateDiff) ContractName() primitives.ContractName {
	return primitives.ContractName(x._message.GetString(0))
}

func (x *ContractStateDiff) RawContractName() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ContractStateDiff) RawContractNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *ContractStateDiff) MutateContractName(v primitives.ContractName) error {
	return x._message.SetString(0, string(v))
}

func (x *ContractStateDiff) StringContractName() string {
	return fmt.Sprintf("%s", x.ContractName())
}

func (x *ContractStateDiff) StateDiffsIterator() *ContractStateDiffStateDiffsIterator {
	return &ContractStateDiffStateDiffsIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type ContractStateDiffStateDiffsIterator struct {
	iterator *membuffers.Iterator
}

func (i *ContractStateDiffStateDiffsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ContractStateDiffStateDiffsIterator) NextStateDiffs() *StateRecord {
	b, s := i.iterator.NextMessage()
	return StateRecordReader(b[:s])
}

func (x *ContractStateDiff) RawStateDiffsArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *ContractStateDiff) RawStateDiffsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *ContractStateDiff) StringStateDiffs() (res string) {
	res = "["
	for i := x.StateDiffsIterator(); i.HasNext(); {
		res += i.NextStateDiffs().String() + ","
	}
	res += "]"
	return
}

// builder

type ContractStateDiffBuilder struct {
	ContractName primitives.ContractName
	StateDiffs   []*StateRecordBuilder

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *ContractStateDiffBuilder) arrayOfStateDiffs() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.StateDiffs))
	for i, v := range w.StateDiffs {
		res[i] = v
	}
	return res
}

func (w *ContractStateDiffBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessageArray(buf, w.arrayOfStateDiffs())
	if err != nil {
		return
	}
	return nil
}

func (w *ContractStateDiffBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpString(prefix, offsetFromStart, "ContractStateDiff.ContractName", string(w.ContractName))
	err = w._builder.HexDumpMessageArray(prefix, offsetFromStart, "ContractStateDiff.StateDiffs", w.arrayOfStateDiffs())
	if err != nil {
		return
	}
	return nil
}

func (w *ContractStateDiffBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ContractStateDiffBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ContractStateDiffBuilder) Build() *ContractStateDiff {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ContractStateDiffReader(buf)
}

func ContractStateDiffBuilderFromRaw(raw []byte) *ContractStateDiffBuilder {
	return &ContractStateDiffBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message Event

// reader

type Event struct {
	// ContractName primitives.ContractName
	// EventName primitives.EventName
	// OutputArgumentArray []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *Event) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContractName:%s,EventName:%s,OutputArgumentArray:%s,}", x.StringContractName(), x.StringEventName(), x.StringOutputArgumentArray())
}

var _Event_Scheme = []membuffers.FieldType{membuffers.TypeString, membuffers.TypeString, membuffers.TypeBytes}
var _Event_Unions = [][]membuffers.FieldType{}

func EventReader(buf []byte) *Event {
	x := &Event{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _Event_Scheme, _Event_Unions)
	return x
}

func (x *Event) IsValid() bool {
	return x._message.IsValid()
}

func (x *Event) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *Event) Equal(y *Event) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *Event) ContractName() primitives.ContractName {
	return primitives.ContractName(x._message.GetString(0))
}

func (x *Event) RawContractName() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Event) RawContractNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *Event) MutateContractName(v primitives.ContractName) error {
	return x._message.SetString(0, string(v))
}

func (x *Event) StringContractName() string {
	return fmt.Sprintf("%s", x.ContractName())
}

func (x *Event) EventName() primitives.EventName {
	return primitives.EventName(x._message.GetString(1))
}

func (x *Event) RawEventName() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *Event) RawEventNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *Event) MutateEventName(v primitives.EventName) error {
	return x._message.SetString(1, string(v))
}

func (x *Event) StringEventName() string {
	return fmt.Sprintf("%s", x.EventName())
}

func (x *Event) OutputArgumentArray() []byte {
	return x._message.GetBytes(2)
}

func (x *Event) RawOutputArgumentArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *Event) RawOutputArgumentArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *Event) MutateOutputArgumentArray(v []byte) error {
	return x._message.SetBytes(2, v)
}

func (x *Event) StringOutputArgumentArray() string {
	return fmt.Sprintf("%x", x.OutputArgumentArray())
}

// builder

type EventBuilder struct {
	ContractName        primitives.ContractName
	EventName           primitives.EventName
	OutputArgumentArray []byte

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *EventBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteString(buf, string(w.EventName))
	w._builder.WriteBytes(buf, w.OutputArgumentArray)
	return nil
}

func (w *EventBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpString(prefix, offsetFromStart, "Event.ContractName", string(w.ContractName))
	w._builder.HexDumpString(prefix, offsetFromStart, "Event.EventName", string(w.EventName))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "Event.OutputArgumentArray", w.OutputArgumentArray)
	return nil
}

func (w *EventBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *EventBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *EventBuilder) Build() *Event {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return EventReader(buf)
}

func EventBuilderFromRaw(raw []byte) *EventBuilder {
	return &EventBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message EventsArray

// reader

type EventsArray struct {
	// Events []Event

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *EventsArray) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Events:%s,}", x.StringEvents())
}

var _EventsArray_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray}
var _EventsArray_Unions = [][]membuffers.FieldType{}

func EventsArrayReader(buf []byte) *EventsArray {
	x := &EventsArray{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _EventsArray_Scheme, _EventsArray_Unions)
	return x
}

func (x *EventsArray) IsValid() bool {
	return x._message.IsValid()
}

func (x *EventsArray) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *EventsArray) Equal(y *EventsArray) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *EventsArray) EventsIterator() *EventsArrayEventsIterator {
	return &EventsArrayEventsIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type EventsArrayEventsIterator struct {
	iterator *membuffers.Iterator
}

func (i *EventsArrayEventsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *EventsArrayEventsIterator) NextEvents() *Event {
	b, s := i.iterator.NextMessage()
	return EventReader(b[:s])
}

func (x *EventsArray) RawEventsArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *EventsArray) RawEventsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *EventsArray) StringEvents() (res string) {
	res = "["
	for i := x.EventsIterator(); i.HasNext(); {
		res += i.NextEvents().String() + ","
	}
	res += "]"
	return
}

// builder

type EventsArrayBuilder struct {
	Events []*EventBuilder

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *EventsArrayBuilder) arrayOfEvents() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Events))
	for i, v := range w.Events {
		res[i] = v
	}
	return res
}

func (w *EventsArrayBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessageArray(buf, w.arrayOfEvents())
	if err != nil {
		return
	}
	return nil
}

func (w *EventsArrayBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessageArray(prefix, offsetFromStart, "EventsArray.Events", w.arrayOfEvents())
	if err != nil {
		return
	}
	return nil
}

func (w *EventsArrayBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *EventsArrayBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *EventsArrayBuilder) Build() *EventsArray {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return EventsArrayReader(buf)
}

func EventsArrayBuilderFromRaw(raw []byte) *EventsArrayBuilder {
	return &EventsArrayBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// enums

type ExecutionAccessScope uint16

const (
	ACCESS_SCOPE_RESERVED   ExecutionAccessScope = 0
	ACCESS_SCOPE_READ_ONLY  ExecutionAccessScope = 1
	ACCESS_SCOPE_READ_WRITE ExecutionAccessScope = 2
)

func (n ExecutionAccessScope) String() string {
	switch n {
	case ACCESS_SCOPE_RESERVED:
		return "ACCESS_SCOPE_RESERVED"
	case ACCESS_SCOPE_READ_ONLY:
		return "ACCESS_SCOPE_READ_ONLY"
	case ACCESS_SCOPE_READ_WRITE:
		return "ACCESS_SCOPE_READ_WRITE"
	}
	return "UNKNOWN"
}

type ExecutionPermissionScope uint16

const (
	PERMISSION_SCOPE_RESERVED ExecutionPermissionScope = 0
	PERMISSION_SCOPE_SYSTEM   ExecutionPermissionScope = 1
	PERMISSION_SCOPE_SERVICE  ExecutionPermissionScope = 2
)

func (n ExecutionPermissionScope) String() string {
	switch n {
	case PERMISSION_SCOPE_RESERVED:
		return "PERMISSION_SCOPE_RESERVED"
	case PERMISSION_SCOPE_SYSTEM:
		return "PERMISSION_SCOPE_SYSTEM"
	case PERMISSION_SCOPE_SERVICE:
		return "PERMISSION_SCOPE_SERVICE"
	}
	return "UNKNOWN"
}

type ProcessorType uint16

const (
	PROCESSOR_TYPE_RESERVED   ProcessorType = 0
	PROCESSOR_TYPE_NATIVE     ProcessorType = 1
	PROCESSOR_TYPE_JAVASCRIPT ProcessorType = 2
)

func (n ProcessorType) String() string {
	switch n {
	case PROCESSOR_TYPE_RESERVED:
		return "PROCESSOR_TYPE_RESERVED"
	case PROCESSOR_TYPE_NATIVE:
		return "PROCESSOR_TYPE_NATIVE"
	case PROCESSOR_TYPE_JAVASCRIPT:
		return "PROCESSOR_TYPE_JAVASCRIPT"
	}
	return "UNKNOWN"
}

type CrosschainConnectorType uint16

const (
	CROSSCHAIN_CONNECTOR_TYPE_RESERVED CrosschainConnectorType = 0
	CROSSCHAIN_CONNECTOR_TYPE_ETHEREUM CrosschainConnectorType = 1
)

func (n CrosschainConnectorType) String() string {
	switch n {
	case CROSSCHAIN_CONNECTOR_TYPE_RESERVED:
		return "CROSSCHAIN_CONNECTOR_TYPE_RESERVED"
	case CROSSCHAIN_CONNECTOR_TYPE_ETHEREUM:
		return "CROSSCHAIN_CONNECTOR_TYPE_ETHEREUM"
	}
	return "UNKNOWN"
}
