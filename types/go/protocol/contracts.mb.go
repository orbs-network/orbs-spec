// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message ContractAddress

// reader

type ContractAddress struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _ContractAddress_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var _ContractAddress_Unions = [][]membuffers.FieldType{}

func ContractAddressReader(buf []byte) *ContractAddress {
	x := &ContractAddress{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ContractAddress_Scheme, _ContractAddress_Unions)
	return x
}

func (x *ContractAddress) IsValid() bool {
	return x._message.IsValid()
}

func (x *ContractAddress) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ContractAddress) Value() string {
	return x._message.GetString(0)
}

func (x *ContractAddress) RawValue() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ContractAddress) MutateValue(v string) error {
	return x._message.SetString(0, v)
}

// builder

type ContractAddressBuilder struct {
	Value string

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *ContractAddressBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteString(buf, w.Value)
	return nil
}

func (w *ContractAddressBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ContractAddressBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ContractAddressBuilder) Build() *ContractAddress {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ContractAddressReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message MethodAddress

// reader

type MethodAddress struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _MethodAddress_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var _MethodAddress_Unions = [][]membuffers.FieldType{}

func MethodAddressReader(buf []byte) *MethodAddress {
	x := &MethodAddress{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _MethodAddress_Scheme, _MethodAddress_Unions)
	return x
}

func (x *MethodAddress) IsValid() bool {
	return x._message.IsValid()
}

func (x *MethodAddress) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *MethodAddress) Value() string {
	return x._message.GetString(0)
}

func (x *MethodAddress) RawValue() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *MethodAddress) MutateValue(v string) error {
	return x._message.SetString(0, v)
}

// builder

type MethodAddressBuilder struct {
	Value string

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *MethodAddressBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteString(buf, w.Value)
	return nil
}

func (w *MethodAddressBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *MethodAddressBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *MethodAddressBuilder) Build() *MethodAddress {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return MethodAddressReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message MethodArgument

// reader

type MethodArgument struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
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

type MethodArgumentType uint16

const (
	MethodArgumentTypeUint32 MethodArgumentType = 0
	MethodArgumentTypeUint64 MethodArgumentType = 1
	MethodArgumentTypeString MethodArgumentType = 2
	MethodArgumentTypeBytes MethodArgumentType = 3
)

func (x *MethodArgument) Type() MethodArgumentType {
	return MethodArgumentType(x._message.GetUint16(1))
}

func (x *MethodArgument) IsTypeUint32() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 0)
	return is
}

func (x *MethodArgument) TypeUint32() uint32 {
	_, off := x._message.IsUnionIndex(1, 0, 0)
	return x._message.GetUint32InOffset(off)
}

func (x *MethodArgument) MutateTypeUint32(v uint32) error {
	is, off := x._message.IsUnionIndex(1, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint32InOffset(off, v)
	return nil
}

func (x *MethodArgument) IsTypeUint64() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 1)
	return is
}

func (x *MethodArgument) TypeUint64() uint64 {
	_, off := x._message.IsUnionIndex(1, 0, 1)
	return x._message.GetUint64InOffset(off)
}

func (x *MethodArgument) MutateTypeUint64(v uint64) error {
	is, off := x._message.IsUnionIndex(1, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint64InOffset(off, v)
	return nil
}

func (x *MethodArgument) IsTypeString() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 2)
	return is
}

func (x *MethodArgument) TypeString() string {
	_, off := x._message.IsUnionIndex(1, 0, 2)
	return x._message.GetStringInOffset(off)
}

func (x *MethodArgument) MutateTypeString(v string) error {
	is, off := x._message.IsUnionIndex(1, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetStringInOffset(off, v)
	return nil
}

func (x *MethodArgument) IsTypeBytes() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 3)
	return is
}

func (x *MethodArgument) TypeBytes() []byte {
	_, off := x._message.IsUnionIndex(1, 0, 3)
	return x._message.GetBytesInOffset(off)
}

func (x *MethodArgument) MutateTypeBytes(v []byte) error {
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

// builder

type MethodArgumentBuilder struct {
	Name string
	Type MethodArgumentType
	Uint32 uint32
	Uint64 uint64
	String string
	Bytes []byte

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
	case MethodArgumentTypeUint32:
		w._builder.WriteUint32(buf, w.Uint32)
	case MethodArgumentTypeUint64:
		w._builder.WriteUint64(buf, w.Uint64)
	case MethodArgumentTypeString:
		w._builder.WriteString(buf, w.String)
	case MethodArgumentTypeBytes:
		w._builder.WriteBytes(buf, w.Bytes)
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
// message StateDiff

// reader

type StateDiff struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _StateDiff_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var _StateDiff_Unions = [][]membuffers.FieldType{}

func StateDiffReader(buf []byte) *StateDiff {
	x := &StateDiff{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _StateDiff_Scheme, _StateDiff_Unions)
	return x
}

func (x *StateDiff) IsValid() bool {
	return x._message.IsValid()
}

func (x *StateDiff) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *StateDiff) StateKey() primitives.Ripmd160Sha256 {
	return x._message.GetBytes(0)
}

func (x *StateDiff) RawStateKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *StateDiff) MutateStateKey(v primitives.Ripmd160Sha256) error {
	return x._message.SetBytes(0, v)
}

func (x *StateDiff) StateValue() []byte {
	return x._message.GetBytes(1)
}

func (x *StateDiff) RawStateValue() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *StateDiff) MutateStateValue(v []byte) error {
	return x._message.SetBytes(1, v)
}

// builder

type StateDiffBuilder struct {
	StateKey primitives.Ripmd160Sha256
	StateValue []byte

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *StateDiffBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteBytes(buf, w.StateKey)
	w._builder.WriteBytes(buf, w.StateValue)
	return nil
}

func (w *StateDiffBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *StateDiffBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *StateDiffBuilder) Build() *StateDiff {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return StateDiffReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ContractStateDiff

// reader

type ContractStateDiff struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _ContractStateDiff_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessageArray,}
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

func (x *ContractStateDiff) Contract() *ContractAddress {
	b, s := x._message.GetMessage(0)
	return ContractAddressReader(b[:s])
}

func (x *ContractStateDiff) RawContract() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ContractStateDiff) StateDiffIterator() *ContractStateDiffStateDiffIterator {
	return &ContractStateDiffStateDiffIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type ContractStateDiffStateDiffIterator struct {
	iterator *membuffers.Iterator
}

func (i *ContractStateDiffStateDiffIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ContractStateDiffStateDiffIterator) NextStateDiff() *StateDiff {
	b, s := i.iterator.NextMessage()
	return StateDiffReader(b[:s])
}

func (x *ContractStateDiff) RawStateDiffArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

// builder

type ContractStateDiffBuilder struct {
	Contract *ContractAddressBuilder
	StateDiff []*StateDiffBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *ContractStateDiffBuilder) arrayOfStateDiff() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.StateDiff))
	for i, v := range w.StateDiff {
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
	err = w._builder.WriteMessage(buf, w.Contract)
	if err != nil {
		return
	}
	err = w._builder.WriteMessageArray(buf, w.arrayOfStateDiff())
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

