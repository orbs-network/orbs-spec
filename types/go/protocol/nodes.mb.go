// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message NodeData

// reader

type NodeData struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _NodeData_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var _NodeData_Unions = [][]membuffers.FieldType{}

func NodeDataReader(buf []byte) *NodeData {
	x := &NodeData{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _NodeData_Scheme, _NodeData_Unions)
	return x
}

func (x *NodeData) IsValid() bool {
	return x._message.IsValid()
}

func (x *NodeData) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *NodeData) NodePkey() primitives.Ed25519Pkey {
	return x._message.GetBytes(0)
}

func (x *NodeData) RawNodePkey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *NodeData) MutateNodePkey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, v)
}

func (x *NodeData) NodeRandomSeedPkey() primitives.Bls1Pkey {
	return x._message.GetBytes(1)
}

func (x *NodeData) RawNodeRandomSeedPkey() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *NodeData) MutateNodeRandomSeedPkey(v primitives.Bls1Pkey) error {
	return x._message.SetBytes(1, v)
}

// builder

type NodeDataBuilder struct {
	NodePkey primitives.Ed25519Pkey
	NodeRandomSeedPkey primitives.Bls1Pkey

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *NodeDataBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteBytes(buf, w.NodePkey)
	w._builder.WriteBytes(buf, w.NodeRandomSeedPkey)
	return nil
}

func (w *NodeDataBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *NodeDataBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *NodeDataBuilder) Build() *NodeData {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return NodeDataReader(buf)
}

