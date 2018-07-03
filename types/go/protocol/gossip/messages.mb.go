// AUTO GENERATED FILE (by membufc proto compiler)
package gossip

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossip/leanhelixconsensus"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossip/blocksync"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossip/transactionrelay"
)

/////////////////////////////////////////////////////////////////////////////
// message Message

// reader

type Message struct {
	message membuffers.Message
}

var m_Message_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeBytesArray,membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeUint32,membuffers.TypeUnion,}
var m_Message_Unions = [][]membuffers.FieldType{{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}}

func MessageReader(buf []byte) *Message {
	x := &Message{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_Message_Scheme, m_Message_Unions)
	return x
}

func (x *Message) IsValid() bool {
	return x.message.IsValid()
}

func (x *Message) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *Message) ProtocolVersion() uint32 {
	return x.message.GetUint32(0)
}

func (x *Message) RawProtocolVersion() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *Message) MutateProtocolVersion(v uint32) error {
	return x.message.SetUint32(0, v)
}

func (x *Message) RecipientsListIterator() *MessageRecipientsListIterator {
	return &MessageRecipientsListIterator{iterator: x.message.GetBytesArrayIterator(1)}
}

type MessageRecipientsListIterator struct {
	iterator *membuffers.Iterator
}

func (i *MessageRecipientsListIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *MessageRecipientsListIterator) NextRecipientsList() []byte {
	return i.iterator.NextBytes()
}

func (x *Message) RawRecipientsListArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *Message) InverseRecipientsList() RecipientsListMode {
	return RecipientsListMode(x.message.GetUint16(2))
}

func (x *Message) RawInverseRecipientsList() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *Message) MutateInverseRecipientsList(v RecipientsListMode) error {
	return x.message.SetUint16(2, uint16(v))
}

func (x *Message) Topic() MessageTopic {
	return MessageTopic(x.message.GetUint16(3))
}

func (x *Message) RawTopic() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *Message) MutateTopic(v MessageTopic) error {
	return x.message.SetUint16(3, uint16(v))
}

func (x *Message) VirtualChain() uint32 {
	return x.message.GetUint32(4)
}

func (x *Message) RawVirtualChain() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *Message) MutateVirtualChain(v uint32) error {
	return x.message.SetUint32(4, v)
}

type MessageMessageTopic uint16

const (
	MessageMessageTopicLeanHelixMessage MessageMessageTopic = 0
	MessageMessageTopicBlockSyncMessage MessageMessageTopic = 1
	MessageMessageTopicTransactionsRelayMessage MessageMessageTopic = 2
)

func (x *Message) MessageTopic() MessageMessageTopic {
	return MessageMessageTopic(x.message.GetUint16(5))
}

func (x *Message) IsMessageTopicLeanHelixMessage() bool {
	is, _ := x.message.IsUnionIndex(5, 0, 0)
	return is
}

func (x *Message) MessageTopicLeanHelixMessage() leanhelixconsensus.Message {
	_, off := x.message.IsUnionIndex(5, 0, 0)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) IsMessageTopicBlockSyncMessage() bool {
	is, _ := x.message.IsUnionIndex(5, 0, 1)
	return is
}

func (x *Message) MessageTopicBlockSyncMessage() blocksync.Message {
	_, off := x.message.IsUnionIndex(5, 0, 1)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) IsMessageTopicTransactionsRelayMessage() bool {
	is, _ := x.message.IsUnionIndex(5, 0, 2)
	return is
}

func (x *Message) MessageTopicTransactionsRelayMessage() transactionrelay.Message {
	_, off := x.message.IsUnionIndex(5, 0, 2)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) RawMessageTopic() []byte {
	return x.message.RawBufferForField(5, 0)
}

// builder

type MessageBuilder struct {
	builder membuffers.Builder
	ProtocolVersion uint32
	RecipientsList [][]byte
	InverseRecipientsList RecipientsListMode
	Topic MessageTopic
	VirtualChain uint32
	MessageTopic MessageMessageTopic
	LeanHelixMessage leanhelixconsensus.Message
	BlockSyncMessage blocksync.Message
	TransactionsRelayMessage transactionrelay.Message
}

func (w *MessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.ProtocolVersion)
	w.builder.WriteBytesArray(buf, w.RecipientsList)
	w.builder.WriteUint16(buf, uint16(w.InverseRecipientsList))
	w.builder.WriteUint16(buf, uint16(w.Topic))
	w.builder.WriteUint32(buf, w.VirtualChain)
	w.builder.WriteUnionIndex(buf, uint16(w.MessageTopic))
	switch w.MessageTopic {
	case MessageMessageTopicLeanHelixMessage:
		w.builder.WriteMessage(buf, w.LeanHelixMessage)
	case MessageMessageTopicBlockSyncMessage:
		w.builder.WriteMessage(buf, w.BlockSyncMessage)
	case MessageMessageTopicTransactionsRelayMessage:
		w.builder.WriteMessage(buf, w.TransactionsRelayMessage)
	}
	return nil
}

func (w *MessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *MessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *MessageBuilder) Build() *Message {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return MessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type MessageTopic uint16

const (
	MESSAGE_TOPIC_RESERVED MessageTopic = 0
	MESSAGE_TOPIC_LEAN_HELIX_CONSENSUS MessageTopic = 1
	MESSAGE_TOPIC_TRANSACTION_RELAY MessageTopic = 2
	MESSAGE_TOPIC_BLOCK_SYNC MessageTopic = 3
)

type RecipientsListMode uint16

const (
	RECIPIENT_LIST_MODE_RESERVED RecipientsListMode = 0
	RECIPIENT_LIST_MODE_SEND_TO_LIST RecipientsListMode = 1
	RECIPIENT_LIST_MODE_SEND_TO_ALL_BUT_LIST RecipientsListMode = 2
)

