// AUTO GENERATED FILE (by membufc proto compiler)
package handlers

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossip"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionRelayGossipHandler

type TransactionRelayGossipHandler interface {
	HandleForwardedTransactions(*HandleForwardedTransactionsInput) (*GossipMessageHandlerOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// service BlockSyncGossipHandler

type BlockSyncGossipHandler interface {
	HandleBlockAvailabilityRequest(*HandleBlockAvailabilityRequestInput) (*GossipMessageHandlerOutput, error)
	HandleBlockAvailabilityResponse(*HandleBlockAvailabilityResponseInput) (*GossipMessageHandlerOutput, error)
	HandleBlockSyncRequest(*HandleBlockSyncRequestInput) (*GossipMessageHandlerOutput, error)
	HandleBlockSyncResponse(*HandleBlockSyncResponseInput) (*GossipMessageHandlerOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// service LeanHelixConsensusGossipHandler

type LeanHelixConsensusGossipHandler interface {
	HandleLeanHelixPrePrepare(*HandleLeanHelixPrePrepareInput) (*GossipMessageHandlerOutput, error)
	HandleLeanHelixPrepare(*HandleLeanHelixPrepareInput) (*GossipMessageHandlerOutput, error)
	HandleLeanHelixCommit(*HandleLeanHelixCommitInput) (*GossipMessageHandlerOutput, error)
	HandleLeanHelixViewChange(*HandleLeanHelixViewChangeInput) (*GossipMessageHandlerOutput, error)
	HandleLeanHelixNewView(*HandleLeanHelixNewViewInput) (*GossipMessageHandlerOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message GossipMessageHandlerOutput

// reader

type GossipMessageHandlerOutput struct {
	message membuffers.Message
}

var m_GossipMessageHandlerOutput_Scheme = []membuffers.FieldType{}
var m_GossipMessageHandlerOutput_Unions = [][]membuffers.FieldType{}

func GossipMessageHandlerOutputReader(buf []byte) *GossipMessageHandlerOutput {
	x := &GossipMessageHandlerOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GossipMessageHandlerOutput_Scheme, m_GossipMessageHandlerOutput_Unions)
	return x
}

func (x *GossipMessageHandlerOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GossipMessageHandlerOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type GossipMessageHandlerOutputBuilder struct {
	builder membuffers.Builder
}

func (w *GossipMessageHandlerOutputBuilder) Write(buf []byte) (err error) {
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

func (w *GossipMessageHandlerOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GossipMessageHandlerOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GossipMessageHandlerOutputBuilder) Build() *GossipMessageHandlerOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GossipMessageHandlerOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleForwardedTransactionsInput

// reader

type HandleForwardedTransactionsInput struct {
	message membuffers.Message
}

var m_HandleForwardedTransactionsInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_HandleForwardedTransactionsInput_Unions = [][]membuffers.FieldType{}

func HandleForwardedTransactionsInputReader(buf []byte) *HandleForwardedTransactionsInput {
	x := &HandleForwardedTransactionsInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleForwardedTransactionsInput_Scheme, m_HandleForwardedTransactionsInput_Unions)
	return x
}

func (x *HandleForwardedTransactionsInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleForwardedTransactionsInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleForwardedTransactionsInput) Sender() []byte {
	return x.message.GetBytes(0)
}

func (x *HandleForwardedTransactionsInput) RawSender() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleForwardedTransactionsInput) MutateSender(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *HandleForwardedTransactionsInput) Message() *gossip.ForwardedTransactionsMessage {
	b, s := x.message.GetMessage(1)
	return gossip.ForwardedTransactionsMessageReader(b[:s])
}

func (x *HandleForwardedTransactionsInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type HandleForwardedTransactionsInputBuilder struct {
	builder membuffers.Builder
	Sender []byte
	Message *gossip.ForwardedTransactionsMessageBuilder
}

func (w *HandleForwardedTransactionsInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Sender)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *HandleForwardedTransactionsInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleForwardedTransactionsInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleForwardedTransactionsInputBuilder) Build() *HandleForwardedTransactionsInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleForwardedTransactionsInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleBlockAvailabilityRequestInput

// reader

type HandleBlockAvailabilityRequestInput struct {
	message membuffers.Message
}

var m_HandleBlockAvailabilityRequestInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_HandleBlockAvailabilityRequestInput_Unions = [][]membuffers.FieldType{}

func HandleBlockAvailabilityRequestInputReader(buf []byte) *HandleBlockAvailabilityRequestInput {
	x := &HandleBlockAvailabilityRequestInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleBlockAvailabilityRequestInput_Scheme, m_HandleBlockAvailabilityRequestInput_Unions)
	return x
}

func (x *HandleBlockAvailabilityRequestInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleBlockAvailabilityRequestInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleBlockAvailabilityRequestInput) Sender() []byte {
	return x.message.GetBytes(0)
}

func (x *HandleBlockAvailabilityRequestInput) RawSender() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleBlockAvailabilityRequestInput) MutateSender(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *HandleBlockAvailabilityRequestInput) Message() *gossip.BlockSyncAvailabilityRequestMessage {
	b, s := x.message.GetMessage(1)
	return gossip.BlockSyncAvailabilityRequestMessageReader(b[:s])
}

func (x *HandleBlockAvailabilityRequestInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type HandleBlockAvailabilityRequestInputBuilder struct {
	builder membuffers.Builder
	Sender []byte
	Message *gossip.BlockSyncAvailabilityRequestMessageBuilder
}

func (w *HandleBlockAvailabilityRequestInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Sender)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *HandleBlockAvailabilityRequestInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleBlockAvailabilityRequestInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleBlockAvailabilityRequestInputBuilder) Build() *HandleBlockAvailabilityRequestInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleBlockAvailabilityRequestInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleBlockAvailabilityResponseInput

// reader

type HandleBlockAvailabilityResponseInput struct {
	message membuffers.Message
}

var m_HandleBlockAvailabilityResponseInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_HandleBlockAvailabilityResponseInput_Unions = [][]membuffers.FieldType{}

func HandleBlockAvailabilityResponseInputReader(buf []byte) *HandleBlockAvailabilityResponseInput {
	x := &HandleBlockAvailabilityResponseInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleBlockAvailabilityResponseInput_Scheme, m_HandleBlockAvailabilityResponseInput_Unions)
	return x
}

func (x *HandleBlockAvailabilityResponseInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleBlockAvailabilityResponseInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleBlockAvailabilityResponseInput) Sender() []byte {
	return x.message.GetBytes(0)
}

func (x *HandleBlockAvailabilityResponseInput) RawSender() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleBlockAvailabilityResponseInput) MutateSender(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *HandleBlockAvailabilityResponseInput) Message() *gossip.BlockSyncAvailabilityResponseMessage {
	b, s := x.message.GetMessage(1)
	return gossip.BlockSyncAvailabilityResponseMessageReader(b[:s])
}

func (x *HandleBlockAvailabilityResponseInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type HandleBlockAvailabilityResponseInputBuilder struct {
	builder membuffers.Builder
	Sender []byte
	Message *gossip.BlockSyncAvailabilityResponseMessageBuilder
}

func (w *HandleBlockAvailabilityResponseInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Sender)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *HandleBlockAvailabilityResponseInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleBlockAvailabilityResponseInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleBlockAvailabilityResponseInputBuilder) Build() *HandleBlockAvailabilityResponseInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleBlockAvailabilityResponseInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleBlockSyncRequestInput

// reader

type HandleBlockSyncRequestInput struct {
	message membuffers.Message
}

var m_HandleBlockSyncRequestInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_HandleBlockSyncRequestInput_Unions = [][]membuffers.FieldType{}

func HandleBlockSyncRequestInputReader(buf []byte) *HandleBlockSyncRequestInput {
	x := &HandleBlockSyncRequestInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleBlockSyncRequestInput_Scheme, m_HandleBlockSyncRequestInput_Unions)
	return x
}

func (x *HandleBlockSyncRequestInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleBlockSyncRequestInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleBlockSyncRequestInput) Sender() []byte {
	return x.message.GetBytes(0)
}

func (x *HandleBlockSyncRequestInput) RawSender() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleBlockSyncRequestInput) MutateSender(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *HandleBlockSyncRequestInput) Message() *gossip.BlockSyncRequestMessage {
	b, s := x.message.GetMessage(1)
	return gossip.BlockSyncRequestMessageReader(b[:s])
}

func (x *HandleBlockSyncRequestInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type HandleBlockSyncRequestInputBuilder struct {
	builder membuffers.Builder
	Sender []byte
	Message *gossip.BlockSyncRequestMessageBuilder
}

func (w *HandleBlockSyncRequestInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Sender)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *HandleBlockSyncRequestInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleBlockSyncRequestInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleBlockSyncRequestInputBuilder) Build() *HandleBlockSyncRequestInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleBlockSyncRequestInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleBlockSyncResponseInput

// reader

type HandleBlockSyncResponseInput struct {
	message membuffers.Message
}

var m_HandleBlockSyncResponseInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_HandleBlockSyncResponseInput_Unions = [][]membuffers.FieldType{}

func HandleBlockSyncResponseInputReader(buf []byte) *HandleBlockSyncResponseInput {
	x := &HandleBlockSyncResponseInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleBlockSyncResponseInput_Scheme, m_HandleBlockSyncResponseInput_Unions)
	return x
}

func (x *HandleBlockSyncResponseInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleBlockSyncResponseInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleBlockSyncResponseInput) Sender() []byte {
	return x.message.GetBytes(0)
}

func (x *HandleBlockSyncResponseInput) RawSender() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleBlockSyncResponseInput) MutateSender(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *HandleBlockSyncResponseInput) Message() *gossip.BlockSyncResponseMessage {
	b, s := x.message.GetMessage(1)
	return gossip.BlockSyncResponseMessageReader(b[:s])
}

func (x *HandleBlockSyncResponseInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type HandleBlockSyncResponseInputBuilder struct {
	builder membuffers.Builder
	Sender []byte
	Message *gossip.BlockSyncResponseMessageBuilder
}

func (w *HandleBlockSyncResponseInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Sender)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *HandleBlockSyncResponseInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleBlockSyncResponseInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleBlockSyncResponseInputBuilder) Build() *HandleBlockSyncResponseInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleBlockSyncResponseInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleLeanHelixPrePrepareInput

// reader

type HandleLeanHelixPrePrepareInput struct {
	message membuffers.Message
}

var m_HandleLeanHelixPrePrepareInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_HandleLeanHelixPrePrepareInput_Unions = [][]membuffers.FieldType{}

func HandleLeanHelixPrePrepareInputReader(buf []byte) *HandleLeanHelixPrePrepareInput {
	x := &HandleLeanHelixPrePrepareInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleLeanHelixPrePrepareInput_Scheme, m_HandleLeanHelixPrePrepareInput_Unions)
	return x
}

func (x *HandleLeanHelixPrePrepareInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleLeanHelixPrePrepareInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleLeanHelixPrePrepareInput) Sender() []byte {
	return x.message.GetBytes(0)
}

func (x *HandleLeanHelixPrePrepareInput) RawSender() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleLeanHelixPrePrepareInput) MutateSender(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *HandleLeanHelixPrePrepareInput) Message() *gossip.LeanHelixPrePrepareMessage {
	b, s := x.message.GetMessage(1)
	return gossip.LeanHelixPrePrepareMessageReader(b[:s])
}

func (x *HandleLeanHelixPrePrepareInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type HandleLeanHelixPrePrepareInputBuilder struct {
	builder membuffers.Builder
	Sender []byte
	Message *gossip.LeanHelixPrePrepareMessageBuilder
}

func (w *HandleLeanHelixPrePrepareInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Sender)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *HandleLeanHelixPrePrepareInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleLeanHelixPrePrepareInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleLeanHelixPrePrepareInputBuilder) Build() *HandleLeanHelixPrePrepareInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleLeanHelixPrePrepareInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleLeanHelixPrepareInput

// reader

type HandleLeanHelixPrepareInput struct {
	message membuffers.Message
}

var m_HandleLeanHelixPrepareInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_HandleLeanHelixPrepareInput_Unions = [][]membuffers.FieldType{}

func HandleLeanHelixPrepareInputReader(buf []byte) *HandleLeanHelixPrepareInput {
	x := &HandleLeanHelixPrepareInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleLeanHelixPrepareInput_Scheme, m_HandleLeanHelixPrepareInput_Unions)
	return x
}

func (x *HandleLeanHelixPrepareInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleLeanHelixPrepareInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleLeanHelixPrepareInput) Sender() []byte {
	return x.message.GetBytes(0)
}

func (x *HandleLeanHelixPrepareInput) RawSender() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleLeanHelixPrepareInput) MutateSender(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *HandleLeanHelixPrepareInput) Message() *gossip.LeanHelixPrepareMessage {
	b, s := x.message.GetMessage(1)
	return gossip.LeanHelixPrepareMessageReader(b[:s])
}

func (x *HandleLeanHelixPrepareInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type HandleLeanHelixPrepareInputBuilder struct {
	builder membuffers.Builder
	Sender []byte
	Message *gossip.LeanHelixPrepareMessageBuilder
}

func (w *HandleLeanHelixPrepareInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Sender)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *HandleLeanHelixPrepareInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleLeanHelixPrepareInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleLeanHelixPrepareInputBuilder) Build() *HandleLeanHelixPrepareInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleLeanHelixPrepareInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleLeanHelixCommitInput

// reader

type HandleLeanHelixCommitInput struct {
	message membuffers.Message
}

var m_HandleLeanHelixCommitInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_HandleLeanHelixCommitInput_Unions = [][]membuffers.FieldType{}

func HandleLeanHelixCommitInputReader(buf []byte) *HandleLeanHelixCommitInput {
	x := &HandleLeanHelixCommitInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleLeanHelixCommitInput_Scheme, m_HandleLeanHelixCommitInput_Unions)
	return x
}

func (x *HandleLeanHelixCommitInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleLeanHelixCommitInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleLeanHelixCommitInput) Sender() []byte {
	return x.message.GetBytes(0)
}

func (x *HandleLeanHelixCommitInput) RawSender() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleLeanHelixCommitInput) MutateSender(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *HandleLeanHelixCommitInput) Message() *gossip.LeanHelixCommitMessage {
	b, s := x.message.GetMessage(1)
	return gossip.LeanHelixCommitMessageReader(b[:s])
}

func (x *HandleLeanHelixCommitInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type HandleLeanHelixCommitInputBuilder struct {
	builder membuffers.Builder
	Sender []byte
	Message *gossip.LeanHelixCommitMessageBuilder
}

func (w *HandleLeanHelixCommitInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Sender)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *HandleLeanHelixCommitInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleLeanHelixCommitInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleLeanHelixCommitInputBuilder) Build() *HandleLeanHelixCommitInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleLeanHelixCommitInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleLeanHelixViewChangeInput

// reader

type HandleLeanHelixViewChangeInput struct {
	message membuffers.Message
}

var m_HandleLeanHelixViewChangeInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_HandleLeanHelixViewChangeInput_Unions = [][]membuffers.FieldType{}

func HandleLeanHelixViewChangeInputReader(buf []byte) *HandleLeanHelixViewChangeInput {
	x := &HandleLeanHelixViewChangeInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleLeanHelixViewChangeInput_Scheme, m_HandleLeanHelixViewChangeInput_Unions)
	return x
}

func (x *HandleLeanHelixViewChangeInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleLeanHelixViewChangeInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleLeanHelixViewChangeInput) Sender() []byte {
	return x.message.GetBytes(0)
}

func (x *HandleLeanHelixViewChangeInput) RawSender() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleLeanHelixViewChangeInput) MutateSender(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *HandleLeanHelixViewChangeInput) Message() *gossip.LeanHelixViewChangeMessage {
	b, s := x.message.GetMessage(1)
	return gossip.LeanHelixViewChangeMessageReader(b[:s])
}

func (x *HandleLeanHelixViewChangeInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type HandleLeanHelixViewChangeInputBuilder struct {
	builder membuffers.Builder
	Sender []byte
	Message *gossip.LeanHelixViewChangeMessageBuilder
}

func (w *HandleLeanHelixViewChangeInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Sender)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *HandleLeanHelixViewChangeInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleLeanHelixViewChangeInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleLeanHelixViewChangeInputBuilder) Build() *HandleLeanHelixViewChangeInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleLeanHelixViewChangeInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleLeanHelixNewViewInput

// reader

type HandleLeanHelixNewViewInput struct {
	message membuffers.Message
}

var m_HandleLeanHelixNewViewInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_HandleLeanHelixNewViewInput_Unions = [][]membuffers.FieldType{}

func HandleLeanHelixNewViewInputReader(buf []byte) *HandleLeanHelixNewViewInput {
	x := &HandleLeanHelixNewViewInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleLeanHelixNewViewInput_Scheme, m_HandleLeanHelixNewViewInput_Unions)
	return x
}

func (x *HandleLeanHelixNewViewInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleLeanHelixNewViewInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleLeanHelixNewViewInput) Sender() []byte {
	return x.message.GetBytes(0)
}

func (x *HandleLeanHelixNewViewInput) RawSender() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleLeanHelixNewViewInput) MutateSender(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *HandleLeanHelixNewViewInput) Message() *gossip.LeanHelixNewViewMessage {
	b, s := x.message.GetMessage(1)
	return gossip.LeanHelixNewViewMessageReader(b[:s])
}

func (x *HandleLeanHelixNewViewInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type HandleLeanHelixNewViewInputBuilder struct {
	builder membuffers.Builder
	Sender []byte
	Message *gossip.LeanHelixNewViewMessageBuilder
}

func (w *HandleLeanHelixNewViewInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Sender)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *HandleLeanHelixNewViewInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleLeanHelixNewViewInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleLeanHelixNewViewInputBuilder) Build() *HandleLeanHelixNewViewInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleLeanHelixNewViewInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

