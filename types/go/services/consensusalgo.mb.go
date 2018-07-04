// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossip/leanhelixconsensus"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusAlgo

type ConsensusAlgo interface {
	OnPrePrepareReceived(*OnPrePrepareReceivedInput) (*OnPrePrepareReceivedOutput, error)
	OnPrepareReceived(*OnPrepareReceivedInput) (*OnPrepareReceivedOutput, error)
	OnCommitReceived(*OnCommitReceivedInput) (*OnCommitReceivedOutput, error)
	OnViewChangeReceived(*OnViewChangeReceivedInput) (*OnViewChangeReceivedOutput, error)
	OnNewViewReceived(*OnNewViewReceivedInput) (*OnNewViewReceivedOutput, error)
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
// message OnPrePrepareReceivedInput

// reader

type OnPrePrepareReceivedInput struct {
	message membuffers.Message
}

var m_OnPrePrepareReceivedInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_OnPrePrepareReceivedInput_Unions = [][]membuffers.FieldType{}

func OnPrePrepareReceivedInputReader(buf []byte) *OnPrePrepareReceivedInput {
	x := &OnPrePrepareReceivedInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnPrePrepareReceivedInput_Scheme, m_OnPrePrepareReceivedInput_Unions)
	return x
}

func (x *OnPrePrepareReceivedInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnPrePrepareReceivedInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *OnPrePrepareReceivedInput) PrePrepare() *leanhelixconsensus.PrePrepareMessage {
	b, s := x.message.GetMessage(0)
	return leanhelixconsensus.PrePrepareMessageReader(b[:s])
}

func (x *OnPrePrepareReceivedInput) RawPrePrepare() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type OnPrePrepareReceivedInputBuilder struct {
	builder membuffers.Builder
	PrePrepare *leanhelixconsensus.PrePrepareMessageBuilder
}

func (w *OnPrePrepareReceivedInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.PrePrepare)
	if err != nil {
		return
	}
	return nil
}

func (w *OnPrePrepareReceivedInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnPrePrepareReceivedInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnPrePrepareReceivedInputBuilder) Build() *OnPrePrepareReceivedInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnPrePrepareReceivedInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnPrePrepareReceivedOutput

// reader

type OnPrePrepareReceivedOutput struct {
	message membuffers.Message
}

var m_OnPrePrepareReceivedOutput_Scheme = []membuffers.FieldType{}
var m_OnPrePrepareReceivedOutput_Unions = [][]membuffers.FieldType{}

func OnPrePrepareReceivedOutputReader(buf []byte) *OnPrePrepareReceivedOutput {
	x := &OnPrePrepareReceivedOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnPrePrepareReceivedOutput_Scheme, m_OnPrePrepareReceivedOutput_Unions)
	return x
}

func (x *OnPrePrepareReceivedOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnPrePrepareReceivedOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type OnPrePrepareReceivedOutputBuilder struct {
	builder membuffers.Builder
}

func (w *OnPrePrepareReceivedOutputBuilder) Write(buf []byte) (err error) {
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

func (w *OnPrePrepareReceivedOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnPrePrepareReceivedOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnPrePrepareReceivedOutputBuilder) Build() *OnPrePrepareReceivedOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnPrePrepareReceivedOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnPrepareReceivedInput

// reader

type OnPrepareReceivedInput struct {
	message membuffers.Message
}

var m_OnPrepareReceivedInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_OnPrepareReceivedInput_Unions = [][]membuffers.FieldType{}

func OnPrepareReceivedInputReader(buf []byte) *OnPrepareReceivedInput {
	x := &OnPrepareReceivedInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnPrepareReceivedInput_Scheme, m_OnPrepareReceivedInput_Unions)
	return x
}

func (x *OnPrepareReceivedInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnPrepareReceivedInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *OnPrepareReceivedInput) Prepare() *leanhelixconsensus.PrepareMessage {
	b, s := x.message.GetMessage(0)
	return leanhelixconsensus.PrepareMessageReader(b[:s])
}

func (x *OnPrepareReceivedInput) RawPrepare() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type OnPrepareReceivedInputBuilder struct {
	builder membuffers.Builder
	Prepare *leanhelixconsensus.PrepareMessageBuilder
}

func (w *OnPrepareReceivedInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Prepare)
	if err != nil {
		return
	}
	return nil
}

func (w *OnPrepareReceivedInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnPrepareReceivedInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnPrepareReceivedInputBuilder) Build() *OnPrepareReceivedInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnPrepareReceivedInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnPrepareReceivedOutput

// reader

type OnPrepareReceivedOutput struct {
	message membuffers.Message
}

var m_OnPrepareReceivedOutput_Scheme = []membuffers.FieldType{}
var m_OnPrepareReceivedOutput_Unions = [][]membuffers.FieldType{}

func OnPrepareReceivedOutputReader(buf []byte) *OnPrepareReceivedOutput {
	x := &OnPrepareReceivedOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnPrepareReceivedOutput_Scheme, m_OnPrepareReceivedOutput_Unions)
	return x
}

func (x *OnPrepareReceivedOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnPrepareReceivedOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type OnPrepareReceivedOutputBuilder struct {
	builder membuffers.Builder
}

func (w *OnPrepareReceivedOutputBuilder) Write(buf []byte) (err error) {
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

func (w *OnPrepareReceivedOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnPrepareReceivedOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnPrepareReceivedOutputBuilder) Build() *OnPrepareReceivedOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnPrepareReceivedOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnCommitReceivedInput

// reader

type OnCommitReceivedInput struct {
	message membuffers.Message
}

var m_OnCommitReceivedInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_OnCommitReceivedInput_Unions = [][]membuffers.FieldType{}

func OnCommitReceivedInputReader(buf []byte) *OnCommitReceivedInput {
	x := &OnCommitReceivedInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnCommitReceivedInput_Scheme, m_OnCommitReceivedInput_Unions)
	return x
}

func (x *OnCommitReceivedInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnCommitReceivedInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *OnCommitReceivedInput) Commit() *leanhelixconsensus.CommitMessage {
	b, s := x.message.GetMessage(0)
	return leanhelixconsensus.CommitMessageReader(b[:s])
}

func (x *OnCommitReceivedInput) RawCommit() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type OnCommitReceivedInputBuilder struct {
	builder membuffers.Builder
	Commit *leanhelixconsensus.CommitMessageBuilder
}

func (w *OnCommitReceivedInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Commit)
	if err != nil {
		return
	}
	return nil
}

func (w *OnCommitReceivedInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnCommitReceivedInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnCommitReceivedInputBuilder) Build() *OnCommitReceivedInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnCommitReceivedInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnCommitReceivedOutput

// reader

type OnCommitReceivedOutput struct {
	message membuffers.Message
}

var m_OnCommitReceivedOutput_Scheme = []membuffers.FieldType{}
var m_OnCommitReceivedOutput_Unions = [][]membuffers.FieldType{}

func OnCommitReceivedOutputReader(buf []byte) *OnCommitReceivedOutput {
	x := &OnCommitReceivedOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnCommitReceivedOutput_Scheme, m_OnCommitReceivedOutput_Unions)
	return x
}

func (x *OnCommitReceivedOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnCommitReceivedOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type OnCommitReceivedOutputBuilder struct {
	builder membuffers.Builder
}

func (w *OnCommitReceivedOutputBuilder) Write(buf []byte) (err error) {
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

func (w *OnCommitReceivedOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnCommitReceivedOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnCommitReceivedOutputBuilder) Build() *OnCommitReceivedOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnCommitReceivedOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnViewChangeReceivedInput

// reader

type OnViewChangeReceivedInput struct {
	message membuffers.Message
}

var m_OnViewChangeReceivedInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_OnViewChangeReceivedInput_Unions = [][]membuffers.FieldType{}

func OnViewChangeReceivedInputReader(buf []byte) *OnViewChangeReceivedInput {
	x := &OnViewChangeReceivedInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnViewChangeReceivedInput_Scheme, m_OnViewChangeReceivedInput_Unions)
	return x
}

func (x *OnViewChangeReceivedInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnViewChangeReceivedInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *OnViewChangeReceivedInput) ViewChange() *leanhelixconsensus.ViewChangeMessage {
	b, s := x.message.GetMessage(0)
	return leanhelixconsensus.ViewChangeMessageReader(b[:s])
}

func (x *OnViewChangeReceivedInput) RawViewChange() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type OnViewChangeReceivedInputBuilder struct {
	builder membuffers.Builder
	ViewChange *leanhelixconsensus.ViewChangeMessageBuilder
}

func (w *OnViewChangeReceivedInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.ViewChange)
	if err != nil {
		return
	}
	return nil
}

func (w *OnViewChangeReceivedInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnViewChangeReceivedInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnViewChangeReceivedInputBuilder) Build() *OnViewChangeReceivedInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnViewChangeReceivedInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnViewChangeReceivedOutput

// reader

type OnViewChangeReceivedOutput struct {
	message membuffers.Message
}

var m_OnViewChangeReceivedOutput_Scheme = []membuffers.FieldType{}
var m_OnViewChangeReceivedOutput_Unions = [][]membuffers.FieldType{}

func OnViewChangeReceivedOutputReader(buf []byte) *OnViewChangeReceivedOutput {
	x := &OnViewChangeReceivedOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnViewChangeReceivedOutput_Scheme, m_OnViewChangeReceivedOutput_Unions)
	return x
}

func (x *OnViewChangeReceivedOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnViewChangeReceivedOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type OnViewChangeReceivedOutputBuilder struct {
	builder membuffers.Builder
}

func (w *OnViewChangeReceivedOutputBuilder) Write(buf []byte) (err error) {
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

func (w *OnViewChangeReceivedOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnViewChangeReceivedOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnViewChangeReceivedOutputBuilder) Build() *OnViewChangeReceivedOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnViewChangeReceivedOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnNewViewReceivedInput

// reader

type OnNewViewReceivedInput struct {
	message membuffers.Message
}

var m_OnNewViewReceivedInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_OnNewViewReceivedInput_Unions = [][]membuffers.FieldType{}

func OnNewViewReceivedInputReader(buf []byte) *OnNewViewReceivedInput {
	x := &OnNewViewReceivedInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnNewViewReceivedInput_Scheme, m_OnNewViewReceivedInput_Unions)
	return x
}

func (x *OnNewViewReceivedInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnNewViewReceivedInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *OnNewViewReceivedInput) NewView() *leanhelixconsensus.NewViewMessage {
	b, s := x.message.GetMessage(0)
	return leanhelixconsensus.NewViewMessageReader(b[:s])
}

func (x *OnNewViewReceivedInput) RawNewView() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type OnNewViewReceivedInputBuilder struct {
	builder membuffers.Builder
	NewView *leanhelixconsensus.NewViewMessageBuilder
}

func (w *OnNewViewReceivedInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.NewView)
	if err != nil {
		return
	}
	return nil
}

func (w *OnNewViewReceivedInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnNewViewReceivedInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnNewViewReceivedInputBuilder) Build() *OnNewViewReceivedInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnNewViewReceivedInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnNewViewReceivedOutput

// reader

type OnNewViewReceivedOutput struct {
	message membuffers.Message
}

var m_OnNewViewReceivedOutput_Scheme = []membuffers.FieldType{}
var m_OnNewViewReceivedOutput_Unions = [][]membuffers.FieldType{}

func OnNewViewReceivedOutputReader(buf []byte) *OnNewViewReceivedOutput {
	x := &OnNewViewReceivedOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnNewViewReceivedOutput_Scheme, m_OnNewViewReceivedOutput_Unions)
	return x
}

func (x *OnNewViewReceivedOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnNewViewReceivedOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type OnNewViewReceivedOutputBuilder struct {
	builder membuffers.Builder
}

func (w *OnNewViewReceivedOutputBuilder) Write(buf []byte) (err error) {
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

func (w *OnNewViewReceivedOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnNewViewReceivedOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnNewViewReceivedOutputBuilder) Build() *OnNewViewReceivedOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnNewViewReceivedOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

