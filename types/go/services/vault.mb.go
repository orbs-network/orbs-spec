// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
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
	EthSign(ctx context.Context, input *NodeSignInput) (*NodeSignOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message NodeSignInput

// reader

type NodeSignInput struct {
	// Data []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *NodeSignInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Data:%s,}", x.StringData())
}

var _NodeSignInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes}
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

func (x *NodeSignInput) Data() []byte {
	return x._message.GetBytes(0)
}

func (x *NodeSignInput) RawData() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *NodeSignInput) RawDataWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *NodeSignInput) MutateData(v []byte) error {
	return x._message.SetBytes(0, v)
}

func (x *NodeSignInput) StringData() string {
	return fmt.Sprintf("%x", x.Data())
}

// builder

type NodeSignInputBuilder struct {
	Data []byte

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
	w._builder.WriteBytes(buf, w.Data)
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
	w._builder.HexDumpBytes(prefix, offsetFromStart, "NodeSignInput.Data", w.Data)
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
