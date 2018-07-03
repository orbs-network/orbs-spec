// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossip/leanhelixconsensus"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusAlgo

type ConsensusAlgo interface {
	OnNewConsensusRound(*OnNewConsensusRoundInput) (*OnNewConsensusRoundOutput, error)
	GossipMessageReceived(*GossipMessageReceivedInput) (*GossipMessageReceivedOutput, error)
	OnPrePrepareReceived(*OnPrePrepareReceivedInput) (*OnPrePrepareReceivedOutput, error)
	OnPrepareReceived(*OnPrepareReceivedInput) (*OnPrepareReceivedOutput, error)
	OnCommitReceived(*OnCommitReceivedInput) (*OnCommitReceivedOutput, error)
	OnViewChangeReceived(*OnViewChangeReceivedInput) (*OnViewChangeReceivedOutput, error)
	OnNewViewReceived(*OnNewViewReceivedInput) (*OnNewViewReceivedOutput, error)
	AcknowledgeTransactionsBlockConsensus(*AcknowledgeTransactionsBlockConsensusInput) (*AcknowledgeTransactionsBlockConsensusOutput, error)
	AcknowledgeResultsBlockConsensus(*AcknowledgeResultsBlockConsensusInput) (*AcknowledgeResultsBlockConsensusOutput, error)
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
// message AcknowledgeTransactionsBlockConsensusInput

// reader

type AcknowledgeTransactionsBlockConsensusInput struct {
	message membuffers.Message
}

var m_AcknowledgeTransactionsBlockConsensusInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}
var m_AcknowledgeTransactionsBlockConsensusInput_Unions = [][]membuffers.FieldType{}

func AcknowledgeTransactionsBlockConsensusInputReader(buf []byte) *AcknowledgeTransactionsBlockConsensusInput {
	x := &AcknowledgeTransactionsBlockConsensusInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_AcknowledgeTransactionsBlockConsensusInput_Scheme, m_AcknowledgeTransactionsBlockConsensusInput_Unions)
	return x
}

func (x *AcknowledgeTransactionsBlockConsensusInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *AcknowledgeTransactionsBlockConsensusInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *AcknowledgeTransactionsBlockConsensusInput) TransactionsBlockHeader() *protocol.TransactionsBlockHeader {
	b, s := x.message.GetMessage(0)
	return protocol.TransactionsBlockHeaderReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawTransactionsBlockHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusInput) TransactionsBlockMetaData() *protocol.TransactionsBlockMetaData {
	b, s := x.message.GetMessage(1)
	return protocol.TransactionsBlockMetaDataReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawTransactionsBlockMetaData() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusInput) TransactionsBlockProof() *protocol.TransactionsBlockProof {
	b, s := x.message.GetMessage(2)
	return protocol.TransactionsBlockProofReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawTransactionsBlockProof() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusInput) PreviousCommittedTransactionsBlockHeader() *protocol.TransactionsBlockHeader {
	b, s := x.message.GetMessage(3)
	return protocol.TransactionsBlockHeaderReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawPreviousCommittedTransactionsBlockHeader() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusInput) PreviousCommittedTransactionsBlockMetaData() *protocol.TransactionsBlockMetaData {
	b, s := x.message.GetMessage(4)
	return protocol.TransactionsBlockMetaDataReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawPreviousCommittedTransactionsBlockMetaData() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusInput) PreviousCommittedTransactionsBlockProof() *protocol.TransactionsBlockProof {
	b, s := x.message.GetMessage(5)
	return protocol.TransactionsBlockProofReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawPreviousCommittedTransactionsBlockProof() []byte {
	return x.message.RawBufferForField(5, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusInput) PreviousCommittedResultsBlockProof() *protocol.ResultsBlockProof {
	b, s := x.message.GetMessage(6)
	return protocol.ResultsBlockProofReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawPreviousCommittedResultsBlockProof() []byte {
	return x.message.RawBufferForField(6, 0)
}

// builder

type AcknowledgeTransactionsBlockConsensusInputBuilder struct {
	builder membuffers.Builder
	TransactionsBlockHeader *protocol.TransactionsBlockHeaderBuilder
	TransactionsBlockMetaData *protocol.TransactionsBlockMetaDataBuilder
	TransactionsBlockProof *protocol.TransactionsBlockProofBuilder
	PreviousCommittedTransactionsBlockHeader *protocol.TransactionsBlockHeaderBuilder
	PreviousCommittedTransactionsBlockMetaData *protocol.TransactionsBlockMetaDataBuilder
	PreviousCommittedTransactionsBlockProof *protocol.TransactionsBlockProofBuilder
	PreviousCommittedResultsBlockProof *protocol.ResultsBlockProofBuilder
}

func (w *AcknowledgeTransactionsBlockConsensusInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.TransactionsBlockHeader)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.TransactionsBlockMetaData)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.TransactionsBlockProof)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedTransactionsBlockHeader)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedTransactionsBlockMetaData)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedTransactionsBlockProof)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedResultsBlockProof)
	if err != nil {
		return
	}
	return nil
}

func (w *AcknowledgeTransactionsBlockConsensusInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *AcknowledgeTransactionsBlockConsensusInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *AcknowledgeTransactionsBlockConsensusInputBuilder) Build() *AcknowledgeTransactionsBlockConsensusInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return AcknowledgeTransactionsBlockConsensusInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message AcknowledgeTransactionsBlockConsensusOutput

// reader

type AcknowledgeTransactionsBlockConsensusOutput struct {
	message membuffers.Message
}

var m_AcknowledgeTransactionsBlockConsensusOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,}
var m_AcknowledgeTransactionsBlockConsensusOutput_Unions = [][]membuffers.FieldType{}

func AcknowledgeTransactionsBlockConsensusOutputReader(buf []byte) *AcknowledgeTransactionsBlockConsensusOutput {
	x := &AcknowledgeTransactionsBlockConsensusOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_AcknowledgeTransactionsBlockConsensusOutput_Scheme, m_AcknowledgeTransactionsBlockConsensusOutput_Unions)
	return x
}

func (x *AcknowledgeTransactionsBlockConsensusOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *AcknowledgeTransactionsBlockConsensusOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *AcknowledgeTransactionsBlockConsensusOutput) Status() ValidateBlockConsensus {
	return ValidateBlockConsensus(x.message.GetUint16(0))
}

func (x *AcknowledgeTransactionsBlockConsensusOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusOutput) MutateStatus(v ValidateBlockConsensus) error {
	return x.message.SetUint16(0, uint16(v))
}

// builder

type AcknowledgeTransactionsBlockConsensusOutputBuilder struct {
	builder membuffers.Builder
	Status ValidateBlockConsensus
}

func (w *AcknowledgeTransactionsBlockConsensusOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.Status))
	return nil
}

func (w *AcknowledgeTransactionsBlockConsensusOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *AcknowledgeTransactionsBlockConsensusOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *AcknowledgeTransactionsBlockConsensusOutputBuilder) Build() *AcknowledgeTransactionsBlockConsensusOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return AcknowledgeTransactionsBlockConsensusOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message AcknowledgeResultsBlockConsensusInput

// reader

type AcknowledgeResultsBlockConsensusInput struct {
	message membuffers.Message
}

var m_AcknowledgeResultsBlockConsensusInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}
var m_AcknowledgeResultsBlockConsensusInput_Unions = [][]membuffers.FieldType{}

func AcknowledgeResultsBlockConsensusInputReader(buf []byte) *AcknowledgeResultsBlockConsensusInput {
	x := &AcknowledgeResultsBlockConsensusInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_AcknowledgeResultsBlockConsensusInput_Scheme, m_AcknowledgeResultsBlockConsensusInput_Unions)
	return x
}

func (x *AcknowledgeResultsBlockConsensusInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *AcknowledgeResultsBlockConsensusInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *AcknowledgeResultsBlockConsensusInput) ResultsBlockHeader() *protocol.ResultsBlockHeader {
	b, s := x.message.GetMessage(0)
	return protocol.ResultsBlockHeaderReader(b[:s])
}

func (x *AcknowledgeResultsBlockConsensusInput) RawResultsBlockHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *AcknowledgeResultsBlockConsensusInput) ResultsBlockProof() *protocol.ResultsBlockProof {
	b, s := x.message.GetMessage(1)
	return protocol.ResultsBlockProofReader(b[:s])
}

func (x *AcknowledgeResultsBlockConsensusInput) RawResultsBlockProof() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *AcknowledgeResultsBlockConsensusInput) PreviousResultsBlockHeader() *protocol.ResultsBlockHeader {
	b, s := x.message.GetMessage(2)
	return protocol.ResultsBlockHeaderReader(b[:s])
}

func (x *AcknowledgeResultsBlockConsensusInput) RawPreviousResultsBlockHeader() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *AcknowledgeResultsBlockConsensusInput) PreviousCommittedTransactionsBlockMetaData() *protocol.TransactionsBlockMetaData {
	b, s := x.message.GetMessage(3)
	return protocol.TransactionsBlockMetaDataReader(b[:s])
}

func (x *AcknowledgeResultsBlockConsensusInput) RawPreviousCommittedTransactionsBlockMetaData() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *AcknowledgeResultsBlockConsensusInput) PreviousCommittedTransactionsBlockProof() *protocol.TransactionsBlockProof {
	b, s := x.message.GetMessage(4)
	return protocol.TransactionsBlockProofReader(b[:s])
}

func (x *AcknowledgeResultsBlockConsensusInput) RawPreviousCommittedTransactionsBlockProof() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *AcknowledgeResultsBlockConsensusInput) PreviousCommittedResultsBlockProof() *protocol.ResultsBlockProof {
	b, s := x.message.GetMessage(5)
	return protocol.ResultsBlockProofReader(b[:s])
}

func (x *AcknowledgeResultsBlockConsensusInput) RawPreviousCommittedResultsBlockProof() []byte {
	return x.message.RawBufferForField(5, 0)
}

// builder

type AcknowledgeResultsBlockConsensusInputBuilder struct {
	builder membuffers.Builder
	ResultsBlockHeader *protocol.ResultsBlockHeaderBuilder
	ResultsBlockProof *protocol.ResultsBlockProofBuilder
	PreviousResultsBlockHeader *protocol.ResultsBlockHeaderBuilder
	PreviousCommittedTransactionsBlockMetaData *protocol.TransactionsBlockMetaDataBuilder
	PreviousCommittedTransactionsBlockProof *protocol.TransactionsBlockProofBuilder
	PreviousCommittedResultsBlockProof *protocol.ResultsBlockProofBuilder
}

func (w *AcknowledgeResultsBlockConsensusInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.ResultsBlockHeader)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.ResultsBlockProof)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousResultsBlockHeader)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedTransactionsBlockMetaData)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedTransactionsBlockProof)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedResultsBlockProof)
	if err != nil {
		return
	}
	return nil
}

func (w *AcknowledgeResultsBlockConsensusInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *AcknowledgeResultsBlockConsensusInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *AcknowledgeResultsBlockConsensusInputBuilder) Build() *AcknowledgeResultsBlockConsensusInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return AcknowledgeResultsBlockConsensusInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message AcknowledgeResultsBlockConsensusOutput

// reader

type AcknowledgeResultsBlockConsensusOutput struct {
	message membuffers.Message
}

var m_AcknowledgeResultsBlockConsensusOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,}
var m_AcknowledgeResultsBlockConsensusOutput_Unions = [][]membuffers.FieldType{}

func AcknowledgeResultsBlockConsensusOutputReader(buf []byte) *AcknowledgeResultsBlockConsensusOutput {
	x := &AcknowledgeResultsBlockConsensusOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_AcknowledgeResultsBlockConsensusOutput_Scheme, m_AcknowledgeResultsBlockConsensusOutput_Unions)
	return x
}

func (x *AcknowledgeResultsBlockConsensusOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *AcknowledgeResultsBlockConsensusOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *AcknowledgeResultsBlockConsensusOutput) Status() ValidateBlockConsensus {
	return ValidateBlockConsensus(x.message.GetUint16(0))
}

func (x *AcknowledgeResultsBlockConsensusOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *AcknowledgeResultsBlockConsensusOutput) MutateStatus(v ValidateBlockConsensus) error {
	return x.message.SetUint16(0, uint16(v))
}

// builder

type AcknowledgeResultsBlockConsensusOutputBuilder struct {
	builder membuffers.Builder
	Status ValidateBlockConsensus
}

func (w *AcknowledgeResultsBlockConsensusOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.Status))
	return nil
}

func (w *AcknowledgeResultsBlockConsensusOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *AcknowledgeResultsBlockConsensusOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *AcknowledgeResultsBlockConsensusOutputBuilder) Build() *AcknowledgeResultsBlockConsensusOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return AcknowledgeResultsBlockConsensusOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type ValidateBlockConsensus uint16

const (
	VALIDATE_BLOCK_CONSENSUS_RESERVED ValidateBlockConsensus = 0
	VALIDATE_BLOCK_CONSENSUS_VALID_FOR_COMMIT ValidateBlockConsensus = 1
	VALIDATE_BLOCK_CONSENSUS_INVALID_FOR_COMMIT ValidateBlockConsensus = 2
)

