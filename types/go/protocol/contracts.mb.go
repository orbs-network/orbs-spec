// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message MethodArgument

// reader

type MethodArgument struct {
	// Name string
	// Type MethodArgumentType

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *MethodArgument) String() string {
	return fmt.Sprintf("{Name:%s,Type:%s,}", x.StringName(), x.StringType())
}

var _MethodArgument_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeUnion,}
var _MethodArgument_Unions = [][]membuffers.FieldType{{membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeString,membuffers.TypeBytes,}}

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

func (x *MethodArgument) Name() string {
	return x._message.GetString(0)
}

func (x *MethodArgument) RawName() []byte {
	return x._message.RawBufferForField(0, 0)
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
	METHOD_ARGUMENT_TYPE_STRING_VALUE MethodArgumentType = 2
	METHOD_ARGUMENT_TYPE_BYTES_VALUE MethodArgumentType = 3
)

func (x *MethodArgument) Type() MethodArgumentType {
	return MethodArgumentType(x._message.GetUint16(1))
}

func (x *MethodArgument) IsTypeUint32Value() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 0)
	return is
}

func (x *MethodArgument) Uint32Value() uint32 {
	_, off := x._message.IsUnionIndex(1, 0, 0)
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
	_, off := x._message.IsUnionIndex(1, 0, 1)
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
	_, off := x._message.IsUnionIndex(1, 0, 2)
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
	_, off := x._message.IsUnionIndex(1, 0, 3)
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
	Name string
	Type MethodArgumentType
	Uint32Value uint32
	Uint64Value uint64
	StringValue string
	BytesValue []byte

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *MethodArgumentBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
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

/////////////////////////////////////////////////////////////////////////////
// message StateRecord

// reader

type StateRecord struct {
	// Key primitives.Ripmd160Sha256
	// Value []byte

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *StateRecord) String() string {
	return fmt.Sprintf("{Key:%s,Value:%s,}", x.StringKey(), x.StringValue())
}

var _StateRecord_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
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

func (x *StateRecord) Key() primitives.Ripmd160Sha256 {
	return primitives.Ripmd160Sha256(x._message.GetBytes(0))
}

func (x *StateRecord) RawKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *StateRecord) MutateKey(v primitives.Ripmd160Sha256) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *StateRecord) StringKey() string {
	return fmt.Sprintf("%x", x.Key())
}

func (x *StateRecord) Value() []byte {
	return x._message.GetBytes(1)
}

func (x *StateRecord) RawValue() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *StateRecord) MutateValue(v []byte) error {
	return x._message.SetBytes(1, v)
}

func (x *StateRecord) StringValue() string {
	return fmt.Sprintf("%x", x.Value())
}

// builder

type StateRecordBuilder struct {
	Key primitives.Ripmd160Sha256
	Value []byte

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *StateRecordBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteBytes(buf, []byte(w.Key))
	w._builder.WriteBytes(buf, w.Value)
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

/////////////////////////////////////////////////////////////////////////////
// message ContractStateDiff

// reader

type ContractStateDiff struct {
	// ContractName primitives.ContractName
	// StateDiffs []StateRecord

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *ContractStateDiff) String() string {
	return fmt.Sprintf("{ContractName:%s,StateDiffs:%s,}", x.StringContractName(), x.StringStateDiffs())
}

var _ContractStateDiff_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeMessageArray,}
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

func (x *ContractStateDiff) ContractName() primitives.ContractName {
	return primitives.ContractName(x._message.GetString(0))
}

func (x *ContractStateDiff) RawContractName() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ContractStateDiff) MutateContractName(v primitives.ContractName) error {
	return x._message.SetString(0, string(v))
}

func (x *ContractStateDiff) StringContractName() string {
	return fmt.Sprintf("%x", x.ContractName())
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
	StateDiffs []*StateRecordBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
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
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteString(buf, string(w.ContractName))
	err = w._builder.WriteMessageArray(buf, w.arrayOfStateDiffs())
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

/////////////////////////////////////////////////////////////////////////////
// enums

type ExecutionAccessScope uint16

const (
	ACCESS_SCOPE_RESERVED ExecutionAccessScope = 0
	ACCESS_SCOPE_READ_ONLY ExecutionAccessScope = 1
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
	PERMISSION_SCOPE_SYSTEM ExecutionPermissionScope = 1
	PERMISSION_SCOPE_SERVICE ExecutionPermissionScope = 2
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

