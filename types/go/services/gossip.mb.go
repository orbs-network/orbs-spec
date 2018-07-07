// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossip"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service Gossip

type Gossip interface {
	BroadcastForwardedTransactions(input *BroadcastForwardedTransactionsInput) (*SendGossipMessageOutput, error)
	BroadcastBlockSyncAvailabilityRequest(input *BroadcastBlockSyncAvailabilityRequestInput) (*SendGossipMessageOutput, error)
	SendBlockSyncAvailabilityResponse(input *SendBlockSyncAvailabilityResponseInput) (*SendGossipMessageOutput, error)
	SendBlockSyncRequest(input *SendBlockSyncRequestInput) (*SendGossipMessageOutput, error)
	SendBlockSyncResponse(input *SendBlockSyncResponseInput) (*SendGossipMessageOutput, error)
	SendLeanHelixPrePrepare(input *SendLeanHelixPrePrepareInput) (*SendGossipMessageOutput, error)
	SendLeanHelixPrepare(input *SendLeanHelixPrepareInput) (*SendGossipMessageOutput, error)
	SendLeanHelixCommit(input *SendLeanHelixCommitInput) (*SendGossipMessageOutput, error)
	SendLeanHelixViewChange(input *SendLeanHelixViewChangeInput) (*SendGossipMessageOutput, error)
	SendLeanHelixNewView(input *SendLeanHelixNewViewInput) (*SendGossipMessageOutput, error)
	RegisterTransactionRelayGossipHandler(handler handlers.TransactionRelayGossipHandler)
	RegisterBlockSyncGossipHandler(handler handlers.BlockSyncGossipHandler)
	RegisterLeanHelixConsensusGossipHandler(handler handlers.LeanHelixConsensusGossipHandler)
}

/////////////////////////////////////////////////////////////////////////////
// message SendGossipMessageOutput

// reader

type SendGossipMessageOutput struct {
	message membuffers.Message
}

var m_SendGossipMessageOutput_Scheme = []membuffers.FieldType{}
var m_SendGossipMessageOutput_Unions = [][]membuffers.FieldType{}

func SendGossipMessageOutputReader(buf []byte) *SendGossipMessageOutput {
	x := &SendGossipMessageOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SendGossipMessageOutput_Scheme, m_SendGossipMessageOutput_Unions)
	return x
}

func (x *SendGossipMessageOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendGossipMessageOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type SendGossipMessageOutputBuilder struct {
	builder membuffers.Builder
}

func (w *SendGossipMessageOutputBuilder) Write(buf []byte) (err error) {
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

func (w *SendGossipMessageOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendGossipMessageOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendGossipMessageOutputBuilder) Build() *SendGossipMessageOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendGossipMessageOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BroadcastForwardedTransactionsInput

// reader

type BroadcastForwardedTransactionsInput struct {
	message membuffers.Message
}

var m_BroadcastForwardedTransactionsInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_BroadcastForwardedTransactionsInput_Unions = [][]membuffers.FieldType{}

func BroadcastForwardedTransactionsInputReader(buf []byte) *BroadcastForwardedTransactionsInput {
	x := &BroadcastForwardedTransactionsInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_BroadcastForwardedTransactionsInput_Scheme, m_BroadcastForwardedTransactionsInput_Unions)
	return x
}

func (x *BroadcastForwardedTransactionsInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *BroadcastForwardedTransactionsInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *BroadcastForwardedTransactionsInput) Message() *gossip.ForwardedTransactionsMessage {
	b, s := x.message.GetMessage(0)
	return gossip.ForwardedTransactionsMessageReader(b[:s])
}

func (x *BroadcastForwardedTransactionsInput) RawMessage() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type BroadcastForwardedTransactionsInputBuilder struct {
	builder membuffers.Builder
	Message *gossip.ForwardedTransactionsMessageBuilder
}

func (w *BroadcastForwardedTransactionsInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *BroadcastForwardedTransactionsInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *BroadcastForwardedTransactionsInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *BroadcastForwardedTransactionsInputBuilder) Build() *BroadcastForwardedTransactionsInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BroadcastForwardedTransactionsInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BroadcastBlockSyncAvailabilityRequestInput

// reader

type BroadcastBlockSyncAvailabilityRequestInput struct {
	message membuffers.Message
}

var m_BroadcastBlockSyncAvailabilityRequestInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_BroadcastBlockSyncAvailabilityRequestInput_Unions = [][]membuffers.FieldType{}

func BroadcastBlockSyncAvailabilityRequestInputReader(buf []byte) *BroadcastBlockSyncAvailabilityRequestInput {
	x := &BroadcastBlockSyncAvailabilityRequestInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_BroadcastBlockSyncAvailabilityRequestInput_Scheme, m_BroadcastBlockSyncAvailabilityRequestInput_Unions)
	return x
}

func (x *BroadcastBlockSyncAvailabilityRequestInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *BroadcastBlockSyncAvailabilityRequestInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *BroadcastBlockSyncAvailabilityRequestInput) Message() *gossip.BlockSyncAvailabilityRequestMessage {
	b, s := x.message.GetMessage(0)
	return gossip.BlockSyncAvailabilityRequestMessageReader(b[:s])
}

func (x *BroadcastBlockSyncAvailabilityRequestInput) RawMessage() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type BroadcastBlockSyncAvailabilityRequestInputBuilder struct {
	builder membuffers.Builder
	Message *gossip.BlockSyncAvailabilityRequestMessageBuilder
}

func (w *BroadcastBlockSyncAvailabilityRequestInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *BroadcastBlockSyncAvailabilityRequestInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *BroadcastBlockSyncAvailabilityRequestInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *BroadcastBlockSyncAvailabilityRequestInputBuilder) Build() *BroadcastBlockSyncAvailabilityRequestInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BroadcastBlockSyncAvailabilityRequestInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SendBlockSyncAvailabilityResponseInput

// reader

type SendBlockSyncAvailabilityResponseInput struct {
	message membuffers.Message
}

var m_SendBlockSyncAvailabilityResponseInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_SendBlockSyncAvailabilityResponseInput_Unions = [][]membuffers.FieldType{}

func SendBlockSyncAvailabilityResponseInputReader(buf []byte) *SendBlockSyncAvailabilityResponseInput {
	x := &SendBlockSyncAvailabilityResponseInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SendBlockSyncAvailabilityResponseInput_Scheme, m_SendBlockSyncAvailabilityResponseInput_Unions)
	return x
}

func (x *SendBlockSyncAvailabilityResponseInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendBlockSyncAvailabilityResponseInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SendBlockSyncAvailabilityResponseInput) Recipient() []byte {
	return x.message.GetBytes(0)
}

func (x *SendBlockSyncAvailabilityResponseInput) RawRecipient() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SendBlockSyncAvailabilityResponseInput) MutateRecipient(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *SendBlockSyncAvailabilityResponseInput) Message() *gossip.BlockSyncAvailabilityResponseMessage {
	b, s := x.message.GetMessage(1)
	return gossip.BlockSyncAvailabilityResponseMessageReader(b[:s])
}

func (x *SendBlockSyncAvailabilityResponseInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type SendBlockSyncAvailabilityResponseInputBuilder struct {
	builder membuffers.Builder
	Recipient []byte
	Message *gossip.BlockSyncAvailabilityResponseMessageBuilder
}

func (w *SendBlockSyncAvailabilityResponseInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Recipient)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *SendBlockSyncAvailabilityResponseInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendBlockSyncAvailabilityResponseInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendBlockSyncAvailabilityResponseInputBuilder) Build() *SendBlockSyncAvailabilityResponseInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendBlockSyncAvailabilityResponseInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SendBlockSyncRequestInput

// reader

type SendBlockSyncRequestInput struct {
	message membuffers.Message
}

var m_SendBlockSyncRequestInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_SendBlockSyncRequestInput_Unions = [][]membuffers.FieldType{}

func SendBlockSyncRequestInputReader(buf []byte) *SendBlockSyncRequestInput {
	x := &SendBlockSyncRequestInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SendBlockSyncRequestInput_Scheme, m_SendBlockSyncRequestInput_Unions)
	return x
}

func (x *SendBlockSyncRequestInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendBlockSyncRequestInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SendBlockSyncRequestInput) Recipient() []byte {
	return x.message.GetBytes(0)
}

func (x *SendBlockSyncRequestInput) RawRecipient() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SendBlockSyncRequestInput) MutateRecipient(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *SendBlockSyncRequestInput) Message() *gossip.BlockSyncRequestMessage {
	b, s := x.message.GetMessage(1)
	return gossip.BlockSyncRequestMessageReader(b[:s])
}

func (x *SendBlockSyncRequestInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type SendBlockSyncRequestInputBuilder struct {
	builder membuffers.Builder
	Recipient []byte
	Message *gossip.BlockSyncRequestMessageBuilder
}

func (w *SendBlockSyncRequestInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Recipient)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *SendBlockSyncRequestInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendBlockSyncRequestInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendBlockSyncRequestInputBuilder) Build() *SendBlockSyncRequestInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendBlockSyncRequestInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SendBlockSyncResponseInput

// reader

type SendBlockSyncResponseInput struct {
	message membuffers.Message
}

var m_SendBlockSyncResponseInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessage,}
var m_SendBlockSyncResponseInput_Unions = [][]membuffers.FieldType{}

func SendBlockSyncResponseInputReader(buf []byte) *SendBlockSyncResponseInput {
	x := &SendBlockSyncResponseInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SendBlockSyncResponseInput_Scheme, m_SendBlockSyncResponseInput_Unions)
	return x
}

func (x *SendBlockSyncResponseInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendBlockSyncResponseInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SendBlockSyncResponseInput) Recipient() []byte {
	return x.message.GetBytes(0)
}

func (x *SendBlockSyncResponseInput) RawRecipient() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SendBlockSyncResponseInput) MutateRecipient(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *SendBlockSyncResponseInput) Message() *gossip.BlockSyncResponseMessage {
	b, s := x.message.GetMessage(1)
	return gossip.BlockSyncResponseMessageReader(b[:s])
}

func (x *SendBlockSyncResponseInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type SendBlockSyncResponseInputBuilder struct {
	builder membuffers.Builder
	Recipient []byte
	Message *gossip.BlockSyncResponseMessageBuilder
}

func (w *SendBlockSyncResponseInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Recipient)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *SendBlockSyncResponseInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendBlockSyncResponseInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendBlockSyncResponseInputBuilder) Build() *SendBlockSyncResponseInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendBlockSyncResponseInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SendLeanHelixPrePrepareInput

// reader

type SendLeanHelixPrePrepareInput struct {
	message membuffers.Message
}

var m_SendLeanHelixPrePrepareInput_Scheme = []membuffers.FieldType{membuffers.TypeBytesArray,membuffers.TypeMessage,}
var m_SendLeanHelixPrePrepareInput_Unions = [][]membuffers.FieldType{}

func SendLeanHelixPrePrepareInputReader(buf []byte) *SendLeanHelixPrePrepareInput {
	x := &SendLeanHelixPrePrepareInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SendLeanHelixPrePrepareInput_Scheme, m_SendLeanHelixPrePrepareInput_Unions)
	return x
}

func (x *SendLeanHelixPrePrepareInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendLeanHelixPrePrepareInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SendLeanHelixPrePrepareInput) RecipientIterator() *SendLeanHelixPrePrepareInputRecipientIterator {
	return &SendLeanHelixPrePrepareInputRecipientIterator{iterator: x.message.GetBytesArrayIterator(0)}
}

type SendLeanHelixPrePrepareInputRecipientIterator struct {
	iterator *membuffers.Iterator
}

func (i *SendLeanHelixPrePrepareInputRecipientIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *SendLeanHelixPrePrepareInputRecipientIterator) NextRecipient() []byte {
	return i.iterator.NextBytes()
}

func (x *SendLeanHelixPrePrepareInput) RawRecipientArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SendLeanHelixPrePrepareInput) Message() *gossip.LeanHelixPrePrepareMessage {
	b, s := x.message.GetMessage(1)
	return gossip.LeanHelixPrePrepareMessageReader(b[:s])
}

func (x *SendLeanHelixPrePrepareInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type SendLeanHelixPrePrepareInputBuilder struct {
	builder membuffers.Builder
	Recipient [][]byte
	Message *gossip.LeanHelixPrePrepareMessageBuilder
}

func (w *SendLeanHelixPrePrepareInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytesArray(buf, w.Recipient)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *SendLeanHelixPrePrepareInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendLeanHelixPrePrepareInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendLeanHelixPrePrepareInputBuilder) Build() *SendLeanHelixPrePrepareInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendLeanHelixPrePrepareInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SendLeanHelixPrepareInput

// reader

type SendLeanHelixPrepareInput struct {
	message membuffers.Message
}

var m_SendLeanHelixPrepareInput_Scheme = []membuffers.FieldType{membuffers.TypeBytesArray,membuffers.TypeMessage,}
var m_SendLeanHelixPrepareInput_Unions = [][]membuffers.FieldType{}

func SendLeanHelixPrepareInputReader(buf []byte) *SendLeanHelixPrepareInput {
	x := &SendLeanHelixPrepareInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SendLeanHelixPrepareInput_Scheme, m_SendLeanHelixPrepareInput_Unions)
	return x
}

func (x *SendLeanHelixPrepareInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendLeanHelixPrepareInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SendLeanHelixPrepareInput) RecipientIterator() *SendLeanHelixPrepareInputRecipientIterator {
	return &SendLeanHelixPrepareInputRecipientIterator{iterator: x.message.GetBytesArrayIterator(0)}
}

type SendLeanHelixPrepareInputRecipientIterator struct {
	iterator *membuffers.Iterator
}

func (i *SendLeanHelixPrepareInputRecipientIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *SendLeanHelixPrepareInputRecipientIterator) NextRecipient() []byte {
	return i.iterator.NextBytes()
}

func (x *SendLeanHelixPrepareInput) RawRecipientArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SendLeanHelixPrepareInput) Message() *gossip.LeanHelixPrepareMessage {
	b, s := x.message.GetMessage(1)
	return gossip.LeanHelixPrepareMessageReader(b[:s])
}

func (x *SendLeanHelixPrepareInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type SendLeanHelixPrepareInputBuilder struct {
	builder membuffers.Builder
	Recipient [][]byte
	Message *gossip.LeanHelixPrepareMessageBuilder
}

func (w *SendLeanHelixPrepareInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytesArray(buf, w.Recipient)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *SendLeanHelixPrepareInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendLeanHelixPrepareInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendLeanHelixPrepareInputBuilder) Build() *SendLeanHelixPrepareInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendLeanHelixPrepareInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SendLeanHelixCommitInput

// reader

type SendLeanHelixCommitInput struct {
	message membuffers.Message
}

var m_SendLeanHelixCommitInput_Scheme = []membuffers.FieldType{membuffers.TypeBytesArray,membuffers.TypeUint16,membuffers.TypeMessage,}
var m_SendLeanHelixCommitInput_Unions = [][]membuffers.FieldType{}

func SendLeanHelixCommitInputReader(buf []byte) *SendLeanHelixCommitInput {
	x := &SendLeanHelixCommitInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SendLeanHelixCommitInput_Scheme, m_SendLeanHelixCommitInput_Unions)
	return x
}

func (x *SendLeanHelixCommitInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendLeanHelixCommitInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SendLeanHelixCommitInput) RecipientIterator() *SendLeanHelixCommitInputRecipientIterator {
	return &SendLeanHelixCommitInputRecipientIterator{iterator: x.message.GetBytesArrayIterator(0)}
}

type SendLeanHelixCommitInputRecipientIterator struct {
	iterator *membuffers.Iterator
}

func (i *SendLeanHelixCommitInputRecipientIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *SendLeanHelixCommitInputRecipientIterator) NextRecipient() []byte {
	return i.iterator.NextBytes()
}

func (x *SendLeanHelixCommitInput) RawRecipientArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SendLeanHelixCommitInput) RecipientMode() gossip.RecipientsListMode {
	return gossip.RecipientsListMode(x.message.GetUint16(1))
}

func (x *SendLeanHelixCommitInput) RawRecipientMode() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *SendLeanHelixCommitInput) MutateRecipientMode(v gossip.RecipientsListMode) error {
	return x.message.SetUint16(1, uint16(v))
}

func (x *SendLeanHelixCommitInput) Message() *gossip.LeanHelixCommitMessage {
	b, s := x.message.GetMessage(2)
	return gossip.LeanHelixCommitMessageReader(b[:s])
}

func (x *SendLeanHelixCommitInput) RawMessage() []byte {
	return x.message.RawBufferForField(2, 0)
}

// builder

type SendLeanHelixCommitInputBuilder struct {
	builder membuffers.Builder
	Recipient [][]byte
	RecipientMode gossip.RecipientsListMode
	Message *gossip.LeanHelixCommitMessageBuilder
}

func (w *SendLeanHelixCommitInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytesArray(buf, w.Recipient)
	w.builder.WriteUint16(buf, uint16(w.RecipientMode))
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *SendLeanHelixCommitInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendLeanHelixCommitInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendLeanHelixCommitInputBuilder) Build() *SendLeanHelixCommitInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendLeanHelixCommitInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SendLeanHelixViewChangeInput

// reader

type SendLeanHelixViewChangeInput struct {
	message membuffers.Message
}

var m_SendLeanHelixViewChangeInput_Scheme = []membuffers.FieldType{membuffers.TypeBytesArray,membuffers.TypeMessage,}
var m_SendLeanHelixViewChangeInput_Unions = [][]membuffers.FieldType{}

func SendLeanHelixViewChangeInputReader(buf []byte) *SendLeanHelixViewChangeInput {
	x := &SendLeanHelixViewChangeInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SendLeanHelixViewChangeInput_Scheme, m_SendLeanHelixViewChangeInput_Unions)
	return x
}

func (x *SendLeanHelixViewChangeInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendLeanHelixViewChangeInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SendLeanHelixViewChangeInput) RecipientIterator() *SendLeanHelixViewChangeInputRecipientIterator {
	return &SendLeanHelixViewChangeInputRecipientIterator{iterator: x.message.GetBytesArrayIterator(0)}
}

type SendLeanHelixViewChangeInputRecipientIterator struct {
	iterator *membuffers.Iterator
}

func (i *SendLeanHelixViewChangeInputRecipientIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *SendLeanHelixViewChangeInputRecipientIterator) NextRecipient() []byte {
	return i.iterator.NextBytes()
}

func (x *SendLeanHelixViewChangeInput) RawRecipientArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SendLeanHelixViewChangeInput) Message() *gossip.LeanHelixViewChangeMessage {
	b, s := x.message.GetMessage(1)
	return gossip.LeanHelixViewChangeMessageReader(b[:s])
}

func (x *SendLeanHelixViewChangeInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type SendLeanHelixViewChangeInputBuilder struct {
	builder membuffers.Builder
	Recipient [][]byte
	Message *gossip.LeanHelixViewChangeMessageBuilder
}

func (w *SendLeanHelixViewChangeInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytesArray(buf, w.Recipient)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *SendLeanHelixViewChangeInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendLeanHelixViewChangeInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendLeanHelixViewChangeInputBuilder) Build() *SendLeanHelixViewChangeInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendLeanHelixViewChangeInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SendLeanHelixNewViewInput

// reader

type SendLeanHelixNewViewInput struct {
	message membuffers.Message
}

var m_SendLeanHelixNewViewInput_Scheme = []membuffers.FieldType{membuffers.TypeBytesArray,membuffers.TypeMessage,}
var m_SendLeanHelixNewViewInput_Unions = [][]membuffers.FieldType{}

func SendLeanHelixNewViewInputReader(buf []byte) *SendLeanHelixNewViewInput {
	x := &SendLeanHelixNewViewInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SendLeanHelixNewViewInput_Scheme, m_SendLeanHelixNewViewInput_Unions)
	return x
}

func (x *SendLeanHelixNewViewInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendLeanHelixNewViewInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SendLeanHelixNewViewInput) RecipientIterator() *SendLeanHelixNewViewInputRecipientIterator {
	return &SendLeanHelixNewViewInputRecipientIterator{iterator: x.message.GetBytesArrayIterator(0)}
}

type SendLeanHelixNewViewInputRecipientIterator struct {
	iterator *membuffers.Iterator
}

func (i *SendLeanHelixNewViewInputRecipientIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *SendLeanHelixNewViewInputRecipientIterator) NextRecipient() []byte {
	return i.iterator.NextBytes()
}

func (x *SendLeanHelixNewViewInput) RawRecipientArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SendLeanHelixNewViewInput) Message() *gossip.LeanHelixNewViewMessage {
	b, s := x.message.GetMessage(1)
	return gossip.LeanHelixNewViewMessageReader(b[:s])
}

func (x *SendLeanHelixNewViewInput) RawMessage() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type SendLeanHelixNewViewInputBuilder struct {
	builder membuffers.Builder
	Recipient [][]byte
	Message *gossip.LeanHelixNewViewMessageBuilder
}

func (w *SendLeanHelixNewViewInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytesArray(buf, w.Recipient)
	err = w.builder.WriteMessage(buf, w.Message)
	if err != nil {
		return
	}
	return nil
}

func (w *SendLeanHelixNewViewInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendLeanHelixNewViewInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendLeanHelixNewViewInputBuilder) Build() *SendLeanHelixNewViewInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendLeanHelixNewViewInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

