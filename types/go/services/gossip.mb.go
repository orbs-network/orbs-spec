// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossip"
)

/////////////////////////////////////////////////////////////////////////////
// service Gossip

type Gossip interface {
	TopicSubscribe(*TopicSubscribeInput) (*TopicSubscribeOutput, error)
	TopicUnsubscribe(*TopicUnsubscribeInput) (*TopicUnsubscribeOutput, error)
	SendMessage(*SendMessageInput) (*SendMessageOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message TopicSubscribeInput

// reader

type TopicSubscribeInput struct {
	message membuffers.Message
}

var m_TopicSubscribeInput_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeString,}
var m_TopicSubscribeInput_Unions = [][]membuffers.FieldType{}

func TopicSubscribeInputReader(buf []byte) *TopicSubscribeInput {
	x := &TopicSubscribeInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TopicSubscribeInput_Scheme, m_TopicSubscribeInput_Unions)
	return x
}

func (x *TopicSubscribeInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *TopicSubscribeInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TopicSubscribeInput) Topic() string {
	return x.message.GetString(0)
}

func (x *TopicSubscribeInput) RawTopic() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TopicSubscribeInput) MutateTopic(v string) error {
	return x.message.SetString(0, v)
}

func (x *TopicSubscribeInput) ServiceId() string {
	return x.message.GetString(1)
}

func (x *TopicSubscribeInput) RawServiceId() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *TopicSubscribeInput) MutateServiceId(v string) error {
	return x.message.SetString(1, v)
}

// builder

type TopicSubscribeInputBuilder struct {
	builder membuffers.Builder
	Topic string
	ServiceId string
}

func (w *TopicSubscribeInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteString(buf, w.Topic)
	w.builder.WriteString(buf, w.ServiceId)
	return nil
}

func (w *TopicSubscribeInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TopicSubscribeInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *TopicSubscribeInputBuilder) Build() *TopicSubscribeInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TopicSubscribeInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TopicSubscribeOutput

// reader

type TopicSubscribeOutput struct {
	message membuffers.Message
}

var m_TopicSubscribeOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,}
var m_TopicSubscribeOutput_Unions = [][]membuffers.FieldType{}

func TopicSubscribeOutputReader(buf []byte) *TopicSubscribeOutput {
	x := &TopicSubscribeOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TopicSubscribeOutput_Scheme, m_TopicSubscribeOutput_Unions)
	return x
}

func (x *TopicSubscribeOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *TopicSubscribeOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TopicSubscribeOutput) SubscriptionToken() uint64 {
	return x.message.GetUint64(0)
}

func (x *TopicSubscribeOutput) RawSubscriptionToken() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TopicSubscribeOutput) MutateSubscriptionToken(v uint64) error {
	return x.message.SetUint64(0, v)
}

// builder

type TopicSubscribeOutputBuilder struct {
	builder membuffers.Builder
	SubscriptionToken uint64
}

func (w *TopicSubscribeOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint64(buf, w.SubscriptionToken)
	return nil
}

func (w *TopicSubscribeOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TopicSubscribeOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *TopicSubscribeOutputBuilder) Build() *TopicSubscribeOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TopicSubscribeOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TopicUnsubscribeInput

// reader

type TopicUnsubscribeInput struct {
	message membuffers.Message
}

var m_TopicUnsubscribeInput_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeString,}
var m_TopicUnsubscribeInput_Unions = [][]membuffers.FieldType{}

func TopicUnsubscribeInputReader(buf []byte) *TopicUnsubscribeInput {
	x := &TopicUnsubscribeInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TopicUnsubscribeInput_Scheme, m_TopicUnsubscribeInput_Unions)
	return x
}

func (x *TopicUnsubscribeInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *TopicUnsubscribeInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TopicUnsubscribeInput) Topic() string {
	return x.message.GetString(0)
}

func (x *TopicUnsubscribeInput) RawTopic() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TopicUnsubscribeInput) MutateTopic(v string) error {
	return x.message.SetString(0, v)
}

func (x *TopicUnsubscribeInput) ServiceId() string {
	return x.message.GetString(1)
}

func (x *TopicUnsubscribeInput) RawServiceId() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *TopicUnsubscribeInput) MutateServiceId(v string) error {
	return x.message.SetString(1, v)
}

// builder

type TopicUnsubscribeInputBuilder struct {
	builder membuffers.Builder
	Topic string
	ServiceId string
}

func (w *TopicUnsubscribeInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteString(buf, w.Topic)
	w.builder.WriteString(buf, w.ServiceId)
	return nil
}

func (w *TopicUnsubscribeInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TopicUnsubscribeInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *TopicUnsubscribeInputBuilder) Build() *TopicUnsubscribeInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TopicUnsubscribeInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TopicUnsubscribeOutput

// reader

type TopicUnsubscribeOutput struct {
	message membuffers.Message
}

var m_TopicUnsubscribeOutput_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeString,}
var m_TopicUnsubscribeOutput_Unions = [][]membuffers.FieldType{}

func TopicUnsubscribeOutputReader(buf []byte) *TopicUnsubscribeOutput {
	x := &TopicUnsubscribeOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TopicUnsubscribeOutput_Scheme, m_TopicUnsubscribeOutput_Unions)
	return x
}

func (x *TopicUnsubscribeOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *TopicUnsubscribeOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TopicUnsubscribeOutput) Topic() string {
	return x.message.GetString(0)
}

func (x *TopicUnsubscribeOutput) RawTopic() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TopicUnsubscribeOutput) MutateTopic(v string) error {
	return x.message.SetString(0, v)
}

func (x *TopicUnsubscribeOutput) ServiceId() string {
	return x.message.GetString(1)
}

func (x *TopicUnsubscribeOutput) RawServiceId() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *TopicUnsubscribeOutput) MutateServiceId(v string) error {
	return x.message.SetString(1, v)
}

// builder

type TopicUnsubscribeOutputBuilder struct {
	builder membuffers.Builder
	Topic string
	ServiceId string
}

func (w *TopicUnsubscribeOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteString(buf, w.Topic)
	w.builder.WriteString(buf, w.ServiceId)
	return nil
}

func (w *TopicUnsubscribeOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TopicUnsubscribeOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *TopicUnsubscribeOutputBuilder) Build() *TopicUnsubscribeOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TopicUnsubscribeOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SendMessageInput

// reader

type SendMessageInput struct {
	message membuffers.Message
}

var m_SendMessageInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_SendMessageInput_Unions = [][]membuffers.FieldType{}

func SendMessageInputReader(buf []byte) *SendMessageInput {
	x := &SendMessageInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SendMessageInput_Scheme, m_SendMessageInput_Unions)
	return x
}

func (x *SendMessageInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendMessageInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SendMessageInput) GossipMessage() *gossip.Message {
	b, s := x.message.GetMessage(0)
	return gossip.MessageReader(b[:s])
}

func (x *SendMessageInput) RawGossipMessage() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type SendMessageInputBuilder struct {
	builder membuffers.Builder
	GossipMessage *gossip.MessageBuilder
}

func (w *SendMessageInputBuilder) Write(buf []byte) (err error) {
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

func (w *SendMessageInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendMessageInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendMessageInputBuilder) Build() *SendMessageInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendMessageInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SendMessageOutput

// reader

type SendMessageOutput struct {
	message membuffers.Message
}

var m_SendMessageOutput_Scheme = []membuffers.FieldType{}
var m_SendMessageOutput_Unions = [][]membuffers.FieldType{}

func SendMessageOutputReader(buf []byte) *SendMessageOutput {
	x := &SendMessageOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_SendMessageOutput_Scheme, m_SendMessageOutput_Unions)
	return x
}

func (x *SendMessageOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendMessageOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type SendMessageOutputBuilder struct {
	builder membuffers.Builder
}

func (w *SendMessageOutputBuilder) Write(buf []byte) (err error) {
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

func (w *SendMessageOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendMessageOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendMessageOutputBuilder) Build() *SendMessageOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendMessageOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

