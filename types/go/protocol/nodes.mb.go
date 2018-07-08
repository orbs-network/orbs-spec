// AUTO GENERATED FILE (by membufc proto compiler v0.0.11)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message NodeData

// reader

type NodeData struct {
	message membuffers.Message
}

var _NodeData_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var _NodeData_Unions = [][]membuffers.FieldType{}

func NodeDataReader(buf []byte) *NodeData {
	x := &NodeData{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _NodeData_Scheme, _NodeData_Unions)
	return x
}

func (x *NodeData) IsValid() bool {
	return x.message.IsValid()
}

func (x *NodeData) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *NodeData) NodePkey() []byte {
	return x.message.GetBytes(0)
}

func (x *NodeData) RawNodePkey() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *NodeData) MutateNodePkey(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *NodeData) NodeRandomSeedPkey() []byte {
	return x.message.GetBytes(1)
}

func (x *NodeData) RawNodeRandomSeedPkey() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *NodeData) MutateNodeRandomSeedPkey(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type NodeDataBuilder struct {
	builder membuffers.Builder
	NodePkey []byte
	NodeRandomSeedPkey []byte
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
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.NodePkey)
	w.builder.WriteBytes(buf, w.NodeRandomSeedPkey)
	return nil
}

func (w *NodeDataBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *NodeDataBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *NodeDataBuilder) Build() *NodeData {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return NodeDataReader(buf)
}

