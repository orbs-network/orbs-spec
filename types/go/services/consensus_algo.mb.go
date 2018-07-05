// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusAlgoLeanHelix

type ConsensusAlgoLeanHelix interface {
	handlers.ConsensusBlocksHandler
	handlers.LeanHelixConsensusGossipHandler
	OnNewConsensusRound(*OnNewConsensusRoundInput) (*OnNewConsensusRoundOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message OnNewConsensusRoundInput

// reader

type OnNewConsensusRoundInput struct {
	message membuffers.Message
}

var m_OnNewConsensusRoundInput_Scheme = []membuffers.FieldType{}
var m_OnNewConsensusRoundInput_Unions = [][]membuffers.FieldType{}

func OnNewConsensusRoundInputReader(buf []byte) *OnNewConsensusRoundInput {
	x := &OnNewConsensusRoundInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnNewConsensusRoundInput_Scheme, m_OnNewConsensusRoundInput_Unions)
	return x
}

func (x *OnNewConsensusRoundInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnNewConsensusRoundInput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type OnNewConsensusRoundInputBuilder struct {
	builder membuffers.Builder
}

func (w *OnNewConsensusRoundInputBuilder) Write(buf []byte) (err error) {
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

func (w *OnNewConsensusRoundInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnNewConsensusRoundInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnNewConsensusRoundInputBuilder) Build() *OnNewConsensusRoundInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnNewConsensusRoundInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnNewConsensusRoundOutput

// reader

type OnNewConsensusRoundOutput struct {
	message membuffers.Message
}

var m_OnNewConsensusRoundOutput_Scheme = []membuffers.FieldType{}
var m_OnNewConsensusRoundOutput_Unions = [][]membuffers.FieldType{}

func OnNewConsensusRoundOutputReader(buf []byte) *OnNewConsensusRoundOutput {
	x := &OnNewConsensusRoundOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnNewConsensusRoundOutput_Scheme, m_OnNewConsensusRoundOutput_Unions)
	return x
}

func (x *OnNewConsensusRoundOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnNewConsensusRoundOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type OnNewConsensusRoundOutputBuilder struct {
	builder membuffers.Builder
}

func (w *OnNewConsensusRoundOutputBuilder) Write(buf []byte) (err error) {
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

func (w *OnNewConsensusRoundOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnNewConsensusRoundOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnNewConsensusRoundOutputBuilder) Build() *OnNewConsensusRoundOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnNewConsensusRoundOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

