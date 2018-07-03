// AUTO GENERATED FILE (by membufc proto compiler)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message ContractAddress

// reader

type ContractAddress struct {
	message membuffers.Message
}

var m_ContractAddress_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var m_ContractAddress_Unions = [][]membuffers.FieldType{}

func ContractAddressReader(buf []byte) *ContractAddress {
	x := &ContractAddress{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ContractAddress_Scheme, m_ContractAddress_Unions)
	return x
}

func (x *ContractAddress) IsValid() bool {
	return x.message.IsValid()
}

func (x *ContractAddress) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ContractAddress) Value() string {
	return x.message.GetString(0)
}

func (x *ContractAddress) RawValue() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ContractAddress) MutateValue(v string) error {
	return x.message.SetString(0, v)
}

// builder

type ContractAddressBuilder struct {
	builder membuffers.Builder
	Value string
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
	w.builder.Reset()
	w.builder.WriteString(buf, w.Value)
	return nil
}

func (w *ContractAddressBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ContractAddressBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
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
	message membuffers.Message
}

var m_MethodAddress_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var m_MethodAddress_Unions = [][]membuffers.FieldType{}

func MethodAddressReader(buf []byte) *MethodAddress {
	x := &MethodAddress{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_MethodAddress_Scheme, m_MethodAddress_Unions)
	return x
}

func (x *MethodAddress) IsValid() bool {
	return x.message.IsValid()
}

func (x *MethodAddress) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *MethodAddress) Value() string {
	return x.message.GetString(0)
}

func (x *MethodAddress) RawValue() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *MethodAddress) MutateValue(v string) error {
	return x.message.SetString(0, v)
}

// builder

type MethodAddressBuilder struct {
	builder membuffers.Builder
	Value string
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
	w.builder.Reset()
	w.builder.WriteString(buf, w.Value)
	return nil
}

func (w *MethodAddressBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *MethodAddressBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
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
	message membuffers.Message
}

var m_MethodArgument_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeUnion,}
var m_MethodArgument_Unions = [][]membuffers.FieldType{{membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeString,membuffers.TypeBytes,}}

func MethodArgumentReader(buf []byte) *MethodArgument {
	x := &MethodArgument{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_MethodArgument_Scheme, m_MethodArgument_Unions)
	return x
}

func (x *MethodArgument) IsValid() bool {
	return x.message.IsValid()
}

func (x *MethodArgument) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *MethodArgument) Name() string {
	return x.message.GetString(0)
}

func (x *MethodArgument) RawName() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *MethodArgument) MutateName(v string) error {
	return x.message.SetString(0, v)
}

type MethodArgumentType uint16

const (
	MethodArgumentTypeUint32 MethodArgumentType = 0
	MethodArgumentTypeUint64 MethodArgumentType = 1
	MethodArgumentTypeString MethodArgumentType = 2
	MethodArgumentTypeBytes MethodArgumentType = 3
)

func (x *MethodArgument) Type() MethodArgumentType {
	return MethodArgumentType(x.message.GetUint16(1))
}

func (x *MethodArgument) IsTypeUint32() bool {
	is, _ := x.message.IsUnionIndex(1, 0, 0)
	return is
}

func (x *MethodArgument) TypeUint32() uint32 {
	_, off := x.message.IsUnionIndex(1, 0, 0)
	return x.message.GetUint32InOffset(off)
}

func (x *MethodArgument) MutateTypeUint32(v uint32) error {
	is, off := x.message.IsUnionIndex(1, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.SetUint32InOffset(off, v)
	return nil
}

func (x *MethodArgument) IsTypeUint64() bool {
	is, _ := x.message.IsUnionIndex(1, 0, 1)
	return is
}

func (x *MethodArgument) TypeUint64() uint64 {
	_, off := x.message.IsUnionIndex(1, 0, 1)
	return x.message.GetUint64InOffset(off)
}

func (x *MethodArgument) MutateTypeUint64(v uint64) error {
	is, off := x.message.IsUnionIndex(1, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.SetUint64InOffset(off, v)
	return nil
}

func (x *MethodArgument) IsTypeString() bool {
	is, _ := x.message.IsUnionIndex(1, 0, 2)
	return is
}

func (x *MethodArgument) TypeString() string {
	_, off := x.message.IsUnionIndex(1, 0, 2)
	return x.message.GetStringInOffset(off)
}

func (x *MethodArgument) MutateTypeString(v string) error {
	is, off := x.message.IsUnionIndex(1, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.SetStringInOffset(off, v)
	return nil
}

func (x *MethodArgument) IsTypeBytes() bool {
	is, _ := x.message.IsUnionIndex(1, 0, 3)
	return is
}

func (x *MethodArgument) TypeBytes() []byte {
	_, off := x.message.IsUnionIndex(1, 0, 3)
	return x.message.GetBytesInOffset(off)
}

func (x *MethodArgument) MutateTypeBytes(v []byte) error {
	is, off := x.message.IsUnionIndex(1, 0, 3)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.SetBytesInOffset(off, v)
	return nil
}

func (x *MethodArgument) RawType() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type MethodArgumentBuilder struct {
	builder membuffers.Builder
	Name string
	Type MethodArgumentType
	Uint32 uint32
	Uint64 uint64
	String string
	Bytes []byte
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
	w.builder.Reset()
	w.builder.WriteString(buf, w.Name)
	w.builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case MethodArgumentTypeUint32:
		w.builder.WriteUint32(buf, w.Uint32)
	case MethodArgumentTypeUint64:
		w.builder.WriteUint64(buf, w.Uint64)
	case MethodArgumentTypeString:
		w.builder.WriteString(buf, w.String)
	case MethodArgumentTypeBytes:
		w.builder.WriteBytes(buf, w.Bytes)
	}
	return nil
}

func (w *MethodArgumentBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *MethodArgumentBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
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
	message membuffers.Message
}

var m_StateDiff_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var m_StateDiff_Unions = [][]membuffers.FieldType{}

func StateDiffReader(buf []byte) *StateDiff {
	x := &StateDiff{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_StateDiff_Scheme, m_StateDiff_Unions)
	return x
}

func (x *StateDiff) IsValid() bool {
	return x.message.IsValid()
}

func (x *StateDiff) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *StateDiff) StateKey() []byte {
	return x.message.GetBytes(0)
}

func (x *StateDiff) RawStateKey() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *StateDiff) MutateStateKey(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *StateDiff) StateValue() []byte {
	return x.message.GetBytes(1)
}

func (x *StateDiff) RawStateValue() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *StateDiff) MutateStateValue(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type StateDiffBuilder struct {
	builder membuffers.Builder
	StateKey []byte
	StateValue []byte
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
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.StateKey)
	w.builder.WriteBytes(buf, w.StateValue)
	return nil
}

func (w *StateDiffBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *StateDiffBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
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
	message membuffers.Message
}

var m_ContractStateDiff_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessageArray,}
var m_ContractStateDiff_Unions = [][]membuffers.FieldType{}

func ContractStateDiffReader(buf []byte) *ContractStateDiff {
	x := &ContractStateDiff{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ContractStateDiff_Scheme, m_ContractStateDiff_Unions)
	return x
}

func (x *ContractStateDiff) IsValid() bool {
	return x.message.IsValid()
}

func (x *ContractStateDiff) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ContractStateDiff) Contract() *ContractAddress {
	b, s := x.message.GetMessage(0)
	return ContractAddressReader(b[:s])
}

func (x *ContractStateDiff) RawContract() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ContractStateDiff) StateDiffIterator() *ContractStateDiffStateDiffIterator {
	return &ContractStateDiffStateDiffIterator{iterator: x.message.GetMessageArrayIterator(1)}
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
	return x.message.RawBufferForField(1, 0)
}

// builder

type ContractStateDiffBuilder struct {
	builder membuffers.Builder
	Contract *ContractAddressBuilder
	StateDiff []*StateDiffBuilder
}

func (w *ContractStateDiffBuilder) arrayOfStateDiff() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.StateDiff))
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
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Contract)
	if err != nil {
		return
	}
	err = w.builder.WriteMessageArray(buf, w.arrayOfStateDiff())
	if err != nil {
		return
	}
	return nil
}

func (w *ContractStateDiffBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ContractStateDiffBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ContractStateDiffBuilder) Build() *ContractStateDiff {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ContractStateDiffReader(buf)
}

