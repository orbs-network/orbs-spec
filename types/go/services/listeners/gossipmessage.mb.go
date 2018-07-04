// AUTO GENERATED FILE (by membufc proto compiler)
package listeners

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossip"
)

/////////////////////////////////////////////////////////////////////////////
// service GossipMessageListener

type GossipMessageListener interface {
	GossipMessageReceived(*GossipMessageReceivedInput) (*GossipMessageReceivedOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message GossipMessageReceivedInput

// reader

type GossipMessageReceivedInput struct {
	message membuffers.Message
}

var m_GossipMessageReceivedInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_GossipMessageReceivedInput_Unions = [][]membuffers.FieldType{}

func GossipMessageReceivedInputReader(buf []byte) *GossipMessageReceivedInput {
	x := &GossipMessageReceivedInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GossipMessageReceivedInput_Scheme, m_GossipMessageReceivedInput_Unions)
	return x
}

func (x *GossipMessageReceivedInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GossipMessageReceivedInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GossipMessageReceivedInput) GossipMessage() *gossip.Message {
	b, s := x.message.GetMessage(0)
	return gossip.MessageReader(b[:s])
}

func (x *GossipMessageReceivedInput) RawGossipMessage() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type GossipMessageReceivedInputBuilder struct {
	builder membuffers.Builder
	GossipMessage *gossip.MessageBuilder
}

func (w *GossipMessageReceivedInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.GossipMessage)
	if err != nil {
		return
	}
	return nil
}

func (w *GossipMessageReceivedInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GossipMessageReceivedInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GossipMessageReceivedInputBuilder) Build() *GossipMessageReceivedInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GossipMessageReceivedInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GossipMessageReceivedOutput

// reader

type GossipMessageReceivedOutput struct {
	message membuffers.Message
}

var m_GossipMessageReceivedOutput_Scheme = []membuffers.FieldType{}
var m_GossipMessageReceivedOutput_Unions = [][]membuffers.FieldType{}

func GossipMessageReceivedOutputReader(buf []byte) *GossipMessageReceivedOutput {
	x := &GossipMessageReceivedOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GossipMessageReceivedOutput_Scheme, m_GossipMessageReceivedOutput_Unions)
	return x
}

func (x *GossipMessageReceivedOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GossipMessageReceivedOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type GossipMessageReceivedOutputBuilder struct {
	builder membuffers.Builder
}

func (w *GossipMessageReceivedOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	return nil
}

func (w *GossipMessageReceivedOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GossipMessageReceivedOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GossipMessageReceivedOutputBuilder) Build() *GossipMessageReceivedOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GossipMessageReceivedOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

