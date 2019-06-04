// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package services

import (
	"bytes"
	"context"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// service Vault

type Vault interface {
	NodeSign(ctx context.Context, input *NodeSignInput) (*NodeSignOutput, error)
	NodeGetNodeAddress(ctx context.Context, input *NodeGetNodeAddressInput) (*NodeGetNodeAddressOutput, error)
	NodeGenerateNewKeyPair(ctx context.Context, input *NodeGenerateNewKeyPairInput) (*NodeGenerateNewKeyPairOutput, error)
	RegisterVirtualChain(ctx context.Context, input *RegisterVirtualChainInput) (*RegisterVirtualChainOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message NodeSignInput

// reader

type NodeSignInput struct {
	// DataHash []byte
	// VirtualChainId primitives.VirtualChainId
	// LocalSignature primitives.EcdsaSecp256K1Sig

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *NodeSignInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{DataHash:%s,VirtualChainId:%s,LocalSignature:%s,}", x.StringDataHash(), x.StringVirtualChainId(), x.StringLocalSignature())
}

var _NodeSignInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes, membuffers.TypeUint32, membuffers.TypeBytes}
var _NodeSignInput_Unions = [][]membuffers.FieldType{}

func NodeSignInputReader(buf []byte) *NodeSignInput {
	x := &NodeSignInput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _NodeSignInput_Scheme, _NodeSignInput_Unions)
	return x
}

func (x *NodeSignInput) IsValid() bool {
	return x._message.IsValid()
}

func (x *NodeSignInput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *NodeSignInput) Equal(y *NodeSignInput) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *NodeSignInput) DataHash() []byte {
	return x._message.GetBytes(0)
}

func (x *NodeSignInput) RawDataHash() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *NodeSignInput) RawDataHashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *NodeSignInput) MutateDataHash(v []byte) error {
	return x._message.SetBytes(0, v)
}

func (x *NodeSignInput) StringDataHash() string {
	return fmt.Sprintf("%x", x.DataHash())
}

func (x *NodeSignInput) VirtualChainId() primitives.VirtualChainId {
	return primitives.VirtualChainId(x._message.GetUint32(1))
}

func (x *NodeSignInput) RawVirtualChainId() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *NodeSignInput) MutateVirtualChainId(v primitives.VirtualChainId) error {
	return x._message.SetUint32(1, uint32(v))
}

func (x *NodeSignInput) StringVirtualChainId() string {
	return fmt.Sprintf("%s", x.VirtualChainId())
}

func (x *NodeSignInput) LocalSignature() primitives.EcdsaSecp256K1Sig {
	return primitives.EcdsaSecp256K1Sig(x._message.GetBytes(2))
}

func (x *NodeSignInput) RawLocalSignature() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *NodeSignInput) RawLocalSignatureWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *NodeSignInput) MutateLocalSignature(v primitives.EcdsaSecp256K1Sig) error {
	return x._message.SetBytes(2, []byte(v))
}

func (x *NodeSignInput) StringLocalSignature() string {
	return fmt.Sprintf("%s", x.LocalSignature())
}

// builder

type NodeSignInputBuilder struct {
	DataHash       []byte
	VirtualChainId primitives.VirtualChainId
	LocalSignature primitives.EcdsaSecp256K1Sig

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *NodeSignInputBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, w.DataHash)
	w._builder.WriteUint32(buf, uint32(w.VirtualChainId))
	w._builder.WriteBytes(buf, []byte(w.LocalSignature))
	return nil
}

func (w *NodeSignInputBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "NodeSignInput.DataHash", w.DataHash)
	w._builder.HexDumpUint32(prefix, offsetFromStart, "NodeSignInput.VirtualChainId", uint32(w.VirtualChainId))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "NodeSignInput.LocalSignature", []byte(w.LocalSignature))
	return nil
}

func (w *NodeSignInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *NodeSignInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *NodeSignInputBuilder) Build() *NodeSignInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return NodeSignInputReader(buf)
}

func NodeSignInputBuilderFromRaw(raw []byte) *NodeSignInputBuilder {
	return &NodeSignInputBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message NodeSignOutput

// reader

type NodeSignOutput struct {
	// Signature primitives.EcdsaSecp256K1Sig

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *NodeSignOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Signature:%s,}", x.StringSignature())
}

var _NodeSignOutput_Scheme = []membuffers.FieldType{membuffers.TypeBytes}
var _NodeSignOutput_Unions = [][]membuffers.FieldType{}

func NodeSignOutputReader(buf []byte) *NodeSignOutput {
	x := &NodeSignOutput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _NodeSignOutput_Scheme, _NodeSignOutput_Unions)
	return x
}

func (x *NodeSignOutput) IsValid() bool {
	return x._message.IsValid()
}

func (x *NodeSignOutput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *NodeSignOutput) Equal(y *NodeSignOutput) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *NodeSignOutput) Signature() primitives.EcdsaSecp256K1Sig {
	return primitives.EcdsaSecp256K1Sig(x._message.GetBytes(0))
}

func (x *NodeSignOutput) RawSignature() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *NodeSignOutput) RawSignatureWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *NodeSignOutput) MutateSignature(v primitives.EcdsaSecp256K1Sig) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *NodeSignOutput) StringSignature() string {
	return fmt.Sprintf("%s", x.Signature())
}

// builder

type NodeSignOutputBuilder struct {
	Signature primitives.EcdsaSecp256K1Sig

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *NodeSignOutputBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, []byte(w.Signature))
	return nil
}

func (w *NodeSignOutputBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "NodeSignOutput.Signature", []byte(w.Signature))
	return nil
}

func (w *NodeSignOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *NodeSignOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *NodeSignOutputBuilder) Build() *NodeSignOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return NodeSignOutputReader(buf)
}

func NodeSignOutputBuilderFromRaw(raw []byte) *NodeSignOutputBuilder {
	return &NodeSignOutputBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message NodeGetNodeAddressInput

// reader

type NodeGetNodeAddressInput struct {

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *NodeGetNodeAddressInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

var _NodeGetNodeAddressInput_Scheme = []membuffers.FieldType{}
var _NodeGetNodeAddressInput_Unions = [][]membuffers.FieldType{}

func NodeGetNodeAddressInputReader(buf []byte) *NodeGetNodeAddressInput {
	x := &NodeGetNodeAddressInput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _NodeGetNodeAddressInput_Scheme, _NodeGetNodeAddressInput_Unions)
	return x
}

func (x *NodeGetNodeAddressInput) IsValid() bool {
	return x._message.IsValid()
}

func (x *NodeGetNodeAddressInput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *NodeGetNodeAddressInput) Equal(y *NodeGetNodeAddressInput) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

// builder

type NodeGetNodeAddressInputBuilder struct {

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *NodeGetNodeAddressInputBuilder) Write(buf []byte) (err error) {
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
	return nil
}

func (w *NodeGetNodeAddressInputBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	return nil
}

func (w *NodeGetNodeAddressInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *NodeGetNodeAddressInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *NodeGetNodeAddressInputBuilder) Build() *NodeGetNodeAddressInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return NodeGetNodeAddressInputReader(buf)
}

func NodeGetNodeAddressInputBuilderFromRaw(raw []byte) *NodeGetNodeAddressInputBuilder {
	return &NodeGetNodeAddressInputBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message NodeGetNodeAddressOutput

// reader

type NodeGetNodeAddressOutput struct {
	// NodeAddress primitives.NodeAddress

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *NodeGetNodeAddressOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{NodeAddress:%s,}", x.StringNodeAddress())
}

var _NodeGetNodeAddressOutput_Scheme = []membuffers.FieldType{membuffers.TypeBytes}
var _NodeGetNodeAddressOutput_Unions = [][]membuffers.FieldType{}

func NodeGetNodeAddressOutputReader(buf []byte) *NodeGetNodeAddressOutput {
	x := &NodeGetNodeAddressOutput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _NodeGetNodeAddressOutput_Scheme, _NodeGetNodeAddressOutput_Unions)
	return x
}

func (x *NodeGetNodeAddressOutput) IsValid() bool {
	return x._message.IsValid()
}

func (x *NodeGetNodeAddressOutput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *NodeGetNodeAddressOutput) Equal(y *NodeGetNodeAddressOutput) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *NodeGetNodeAddressOutput) NodeAddress() primitives.NodeAddress {
	return primitives.NodeAddress(x._message.GetBytes(0))
}

func (x *NodeGetNodeAddressOutput) RawNodeAddress() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *NodeGetNodeAddressOutput) RawNodeAddressWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *NodeGetNodeAddressOutput) MutateNodeAddress(v primitives.NodeAddress) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *NodeGetNodeAddressOutput) StringNodeAddress() string {
	return fmt.Sprintf("%s", x.NodeAddress())
}

// builder

type NodeGetNodeAddressOutputBuilder struct {
	NodeAddress primitives.NodeAddress

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *NodeGetNodeAddressOutputBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, []byte(w.NodeAddress))
	return nil
}

func (w *NodeGetNodeAddressOutputBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "NodeGetNodeAddressOutput.NodeAddress", []byte(w.NodeAddress))
	return nil
}

func (w *NodeGetNodeAddressOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *NodeGetNodeAddressOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *NodeGetNodeAddressOutputBuilder) Build() *NodeGetNodeAddressOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return NodeGetNodeAddressOutputReader(buf)
}

func NodeGetNodeAddressOutputBuilderFromRaw(raw []byte) *NodeGetNodeAddressOutputBuilder {
	return &NodeGetNodeAddressOutputBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message NodeGenerateNewKeyPairInput

// reader

type NodeGenerateNewKeyPairInput struct {

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *NodeGenerateNewKeyPairInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

var _NodeGenerateNewKeyPairInput_Scheme = []membuffers.FieldType{}
var _NodeGenerateNewKeyPairInput_Unions = [][]membuffers.FieldType{}

func NodeGenerateNewKeyPairInputReader(buf []byte) *NodeGenerateNewKeyPairInput {
	x := &NodeGenerateNewKeyPairInput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _NodeGenerateNewKeyPairInput_Scheme, _NodeGenerateNewKeyPairInput_Unions)
	return x
}

func (x *NodeGenerateNewKeyPairInput) IsValid() bool {
	return x._message.IsValid()
}

func (x *NodeGenerateNewKeyPairInput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *NodeGenerateNewKeyPairInput) Equal(y *NodeGenerateNewKeyPairInput) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

// builder

type NodeGenerateNewKeyPairInputBuilder struct {

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *NodeGenerateNewKeyPairInputBuilder) Write(buf []byte) (err error) {
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
	return nil
}

func (w *NodeGenerateNewKeyPairInputBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	return nil
}

func (w *NodeGenerateNewKeyPairInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *NodeGenerateNewKeyPairInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *NodeGenerateNewKeyPairInputBuilder) Build() *NodeGenerateNewKeyPairInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return NodeGenerateNewKeyPairInputReader(buf)
}

func NodeGenerateNewKeyPairInputBuilderFromRaw(raw []byte) *NodeGenerateNewKeyPairInputBuilder {
	return &NodeGenerateNewKeyPairInputBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message NodeGenerateNewKeyPairOutput

// reader

type NodeGenerateNewKeyPairOutput struct {
	// NodeAddress primitives.NodeAddress

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *NodeGenerateNewKeyPairOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{NodeAddress:%s,}", x.StringNodeAddress())
}

var _NodeGenerateNewKeyPairOutput_Scheme = []membuffers.FieldType{membuffers.TypeBytes}
var _NodeGenerateNewKeyPairOutput_Unions = [][]membuffers.FieldType{}

func NodeGenerateNewKeyPairOutputReader(buf []byte) *NodeGenerateNewKeyPairOutput {
	x := &NodeGenerateNewKeyPairOutput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _NodeGenerateNewKeyPairOutput_Scheme, _NodeGenerateNewKeyPairOutput_Unions)
	return x
}

func (x *NodeGenerateNewKeyPairOutput) IsValid() bool {
	return x._message.IsValid()
}

func (x *NodeGenerateNewKeyPairOutput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *NodeGenerateNewKeyPairOutput) Equal(y *NodeGenerateNewKeyPairOutput) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *NodeGenerateNewKeyPairOutput) NodeAddress() primitives.NodeAddress {
	return primitives.NodeAddress(x._message.GetBytes(0))
}

func (x *NodeGenerateNewKeyPairOutput) RawNodeAddress() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *NodeGenerateNewKeyPairOutput) RawNodeAddressWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *NodeGenerateNewKeyPairOutput) MutateNodeAddress(v primitives.NodeAddress) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *NodeGenerateNewKeyPairOutput) StringNodeAddress() string {
	return fmt.Sprintf("%s", x.NodeAddress())
}

// builder

type NodeGenerateNewKeyPairOutputBuilder struct {
	NodeAddress primitives.NodeAddress

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *NodeGenerateNewKeyPairOutputBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, []byte(w.NodeAddress))
	return nil
}

func (w *NodeGenerateNewKeyPairOutputBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "NodeGenerateNewKeyPairOutput.NodeAddress", []byte(w.NodeAddress))
	return nil
}

func (w *NodeGenerateNewKeyPairOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *NodeGenerateNewKeyPairOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *NodeGenerateNewKeyPairOutputBuilder) Build() *NodeGenerateNewKeyPairOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return NodeGenerateNewKeyPairOutputReader(buf)
}

func NodeGenerateNewKeyPairOutputBuilderFromRaw(raw []byte) *NodeGenerateNewKeyPairOutputBuilder {
	return &NodeGenerateNewKeyPairOutputBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message RegisterVirtualChainInput

// reader

type RegisterVirtualChainInput struct {
	// Data VCRegistrationData
	// OwnerSignature primitives.EcdsaSecp256K1Sig

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *RegisterVirtualChainInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Data:%s,OwnerSignature:%s,}", x.StringData(), x.StringOwnerSignature())
}

var _RegisterVirtualChainInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage, membuffers.TypeBytes}
var _RegisterVirtualChainInput_Unions = [][]membuffers.FieldType{}

func RegisterVirtualChainInputReader(buf []byte) *RegisterVirtualChainInput {
	x := &RegisterVirtualChainInput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _RegisterVirtualChainInput_Scheme, _RegisterVirtualChainInput_Unions)
	return x
}

func (x *RegisterVirtualChainInput) IsValid() bool {
	return x._message.IsValid()
}

func (x *RegisterVirtualChainInput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *RegisterVirtualChainInput) Equal(y *RegisterVirtualChainInput) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *RegisterVirtualChainInput) Data() *VCRegistrationData {
	b, s := x._message.GetMessage(0)
	return VCRegistrationDataReader(b[:s])
}

func (x *RegisterVirtualChainInput) RawData() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *RegisterVirtualChainInput) RawDataWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *RegisterVirtualChainInput) StringData() string {
	return x.Data().String()
}

func (x *RegisterVirtualChainInput) OwnerSignature() primitives.EcdsaSecp256K1Sig {
	return primitives.EcdsaSecp256K1Sig(x._message.GetBytes(1))
}

func (x *RegisterVirtualChainInput) RawOwnerSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *RegisterVirtualChainInput) RawOwnerSignatureWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *RegisterVirtualChainInput) MutateOwnerSignature(v primitives.EcdsaSecp256K1Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *RegisterVirtualChainInput) StringOwnerSignature() string {
	return fmt.Sprintf("%s", x.OwnerSignature())
}

// builder

type RegisterVirtualChainInputBuilder struct {
	Data           *VCRegistrationDataBuilder
	OwnerSignature primitives.EcdsaSecp256K1Sig

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *RegisterVirtualChainInputBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessage(buf, w.Data)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, []byte(w.OwnerSignature))
	return nil
}

func (w *RegisterVirtualChainInputBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "RegisterVirtualChainInput.Data", w.Data)
	if err != nil {
		return
	}
	w._builder.HexDumpBytes(prefix, offsetFromStart, "RegisterVirtualChainInput.OwnerSignature", []byte(w.OwnerSignature))
	return nil
}

func (w *RegisterVirtualChainInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *RegisterVirtualChainInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *RegisterVirtualChainInputBuilder) Build() *RegisterVirtualChainInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RegisterVirtualChainInputReader(buf)
}

func RegisterVirtualChainInputBuilderFromRaw(raw []byte) *RegisterVirtualChainInputBuilder {
	return &RegisterVirtualChainInputBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message RegisterVirtualChainOutput

// reader

type RegisterVirtualChainOutput struct {

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *RegisterVirtualChainOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

var _RegisterVirtualChainOutput_Scheme = []membuffers.FieldType{}
var _RegisterVirtualChainOutput_Unions = [][]membuffers.FieldType{}

func RegisterVirtualChainOutputReader(buf []byte) *RegisterVirtualChainOutput {
	x := &RegisterVirtualChainOutput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _RegisterVirtualChainOutput_Scheme, _RegisterVirtualChainOutput_Unions)
	return x
}

func (x *RegisterVirtualChainOutput) IsValid() bool {
	return x._message.IsValid()
}

func (x *RegisterVirtualChainOutput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *RegisterVirtualChainOutput) Equal(y *RegisterVirtualChainOutput) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

// builder

type RegisterVirtualChainOutputBuilder struct {

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *RegisterVirtualChainOutputBuilder) Write(buf []byte) (err error) {
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
	return nil
}

func (w *RegisterVirtualChainOutputBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	return nil
}

func (w *RegisterVirtualChainOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *RegisterVirtualChainOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *RegisterVirtualChainOutputBuilder) Build() *RegisterVirtualChainOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RegisterVirtualChainOutputReader(buf)
}

func RegisterVirtualChainOutputBuilderFromRaw(raw []byte) *RegisterVirtualChainOutputBuilder {
	return &RegisterVirtualChainOutputBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message VCRegistrationData

// reader

type VCRegistrationData struct {
	// VirtualChainId primitives.VirtualChainId
	// LocalAddress primitives.NodeAddress

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *VCRegistrationData) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{VirtualChainId:%s,LocalAddress:%s,}", x.StringVirtualChainId(), x.StringLocalAddress())
}

var _VCRegistrationData_Scheme = []membuffers.FieldType{membuffers.TypeUint32, membuffers.TypeBytes}
var _VCRegistrationData_Unions = [][]membuffers.FieldType{}

func VCRegistrationDataReader(buf []byte) *VCRegistrationData {
	x := &VCRegistrationData{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _VCRegistrationData_Scheme, _VCRegistrationData_Unions)
	return x
}

func (x *VCRegistrationData) IsValid() bool {
	return x._message.IsValid()
}

func (x *VCRegistrationData) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *VCRegistrationData) Equal(y *VCRegistrationData) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *VCRegistrationData) VirtualChainId() primitives.VirtualChainId {
	return primitives.VirtualChainId(x._message.GetUint32(0))
}

func (x *VCRegistrationData) RawVirtualChainId() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *VCRegistrationData) MutateVirtualChainId(v primitives.VirtualChainId) error {
	return x._message.SetUint32(0, uint32(v))
}

func (x *VCRegistrationData) StringVirtualChainId() string {
	return fmt.Sprintf("%s", x.VirtualChainId())
}

func (x *VCRegistrationData) LocalAddress() primitives.NodeAddress {
	return primitives.NodeAddress(x._message.GetBytes(1))
}

func (x *VCRegistrationData) RawLocalAddress() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *VCRegistrationData) RawLocalAddressWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *VCRegistrationData) MutateLocalAddress(v primitives.NodeAddress) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *VCRegistrationData) StringLocalAddress() string {
	return fmt.Sprintf("%s", x.LocalAddress())
}

// builder

type VCRegistrationDataBuilder struct {
	VirtualChainId primitives.VirtualChainId
	LocalAddress   primitives.NodeAddress

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *VCRegistrationDataBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUint32(buf, uint32(w.VirtualChainId))
	w._builder.WriteBytes(buf, []byte(w.LocalAddress))
	return nil
}

func (w *VCRegistrationDataBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint32(prefix, offsetFromStart, "VCRegistrationData.VirtualChainId", uint32(w.VirtualChainId))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "VCRegistrationData.LocalAddress", []byte(w.LocalAddress))
	return nil
}

func (w *VCRegistrationDataBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *VCRegistrationDataBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *VCRegistrationDataBuilder) Build() *VCRegistrationData {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return VCRegistrationDataReader(buf)
}

func VCRegistrationDataBuilderFromRaw(raw []byte) *VCRegistrationDataBuilder {
	return &VCRegistrationDataBuilder{_overrideWithRawBuffer: raw}
}
