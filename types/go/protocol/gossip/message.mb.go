// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package gossip

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message MessageHeader

// reader

type MessageHeader struct {
	message membuffers.Message
}

var _MessageHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeBytesArray,membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeUint32,membuffers.TypeUnion,}
var _MessageHeader_Unions = [][]membuffers.FieldType{{membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeUint16,}}

func MessageHeaderReader(buf []byte) *MessageHeader {
	x := &MessageHeader{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _MessageHeader_Scheme, _MessageHeader_Unions)
	return x
}

func (x *MessageHeader) IsValid() bool {
	return x.message.IsValid()
}

func (x *MessageHeader) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *MessageHeader) ProtocolVersion() uint32 {
	return x.message.GetUint32(0)
}

func (x *MessageHeader) RawProtocolVersion() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *MessageHeader) MutateProtocolVersion(v uint32) error {
	return x.message.SetUint32(0, v)
}

func (x *MessageHeader) RecipientIterator() *MessageHeaderRecipientIterator {
	return &MessageHeaderRecipientIterator{iterator: x.message.GetBytesArrayIterator(1)}
}

type MessageHeaderRecipientIterator struct {
	iterator *membuffers.Iterator
}

func (i *MessageHeaderRecipientIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *MessageHeaderRecipientIterator) NextRecipient() primitives.Ed25519Pkey {
	return i.iterator.NextBytes()
}

func (x *MessageHeader) RawRecipientArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *MessageHeader) RecipientMode() RecipientsListMode {
	return RecipientsListMode(x.message.GetUint16(2))
}

func (x *MessageHeader) RawRecipientMode() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *MessageHeader) MutateRecipientMode(v RecipientsListMode) error {
	return x.message.SetUint16(2, uint16(v))
}

func (x *MessageHeader) Topic() MessageTopic {
	return MessageTopic(x.message.GetUint16(3))
}

func (x *MessageHeader) RawTopic() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *MessageHeader) MutateTopic(v MessageTopic) error {
	return x.message.SetUint16(3, uint16(v))
}

func (x *MessageHeader) VirtualChain() uint32 {
	return x.message.GetUint32(4)
}

func (x *MessageHeader) RawVirtualChain() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *MessageHeader) MutateVirtualChain(v uint32) error {
	return x.message.SetUint32(4, v)
}

type MessageHeaderType uint16

const (
	MessageHeaderTypeTransactionRelay MessageHeaderType = 0
	MessageHeaderTypeBlockSync MessageHeaderType = 1
	MessageHeaderTypeLeanHelixConsensus MessageHeaderType = 2
)

func (x *MessageHeader) Type() MessageHeaderType {
	return MessageHeaderType(x.message.GetUint16(5))
}

func (x *MessageHeader) IsTypeTransactionRelay() bool {
	is, _ := x.message.IsUnionIndex(5, 0, 0)
	return is
}

func (x *MessageHeader) TypeTransactionRelay() TransactionsRelayMessageType {
	_, off := x.message.IsUnionIndex(5, 0, 0)
	return TransactionsRelayMessageType(x.message.GetUint16InOffset(off))
}

func (x *MessageHeader) MutateTypeTransactionRelay(v TransactionsRelayMessageType) error {
	is, off := x.message.IsUnionIndex(5, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *MessageHeader) IsTypeBlockSync() bool {
	is, _ := x.message.IsUnionIndex(5, 0, 1)
	return is
}

func (x *MessageHeader) TypeBlockSync() BlockSyncMessageType {
	_, off := x.message.IsUnionIndex(5, 0, 1)
	return BlockSyncMessageType(x.message.GetUint16InOffset(off))
}

func (x *MessageHeader) MutateTypeBlockSync(v BlockSyncMessageType) error {
	is, off := x.message.IsUnionIndex(5, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *MessageHeader) IsTypeLeanHelixConsensus() bool {
	is, _ := x.message.IsUnionIndex(5, 0, 2)
	return is
}

func (x *MessageHeader) TypeLeanHelixConsensus() LeanHelixMessageType {
	_, off := x.message.IsUnionIndex(5, 0, 2)
	return LeanHelixMessageType(x.message.GetUint16InOffset(off))
}

func (x *MessageHeader) MutateTypeLeanHelixConsensus(v LeanHelixMessageType) error {
	is, off := x.message.IsUnionIndex(5, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *MessageHeader) RawType() []byte {
	return x.message.RawBufferForField(5, 0)
}

// builder

type MessageHeaderBuilder struct {
	builder membuffers.Builder
	ProtocolVersion uint32
	Recipient []primitives.Ed25519Pkey
	RecipientMode RecipientsListMode
	Topic MessageTopic
	VirtualChain uint32
	Type MessageHeaderType
	TransactionRelay TransactionsRelayMessageType
	BlockSync BlockSyncMessageType
	LeanHelixConsensus LeanHelixMessageType
}

func (w *MessageHeaderBuilder) arrayOfRecipient() [][]byte {
	res := make([][]byte, len(w.Recipient))
	for i, v := range w.Recipient {
		res[i] = v
	}
	return res
}

func (w *MessageHeaderBuilder) Write(buf []byte) (err error) {
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
	w.builder.WriteBytesArray(buf, w.arrayOfRecipient())
	w.builder.WriteUint16(buf, uint16(w.RecipientMode))
	w.builder.WriteUint16(buf, uint16(w.Topic))
	w.builder.WriteUint32(buf, w.VirtualChain)
	w.builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case MessageHeaderTypeTransactionRelay:
		w.builder.WriteUint16(buf, uint16(w.TransactionRelay))
	case MessageHeaderTypeBlockSync:
		w.builder.WriteUint16(buf, uint16(w.BlockSync))
	case MessageHeaderTypeLeanHelixConsensus:
		w.builder.WriteUint16(buf, uint16(w.LeanHelixConsensus))
	}
	return nil
}

func (w *MessageHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *MessageHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *MessageHeaderBuilder) Build() *MessageHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return MessageHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type MessageTopic uint16

const (
	MESSAGE_TOPIC_RESERVED MessageTopic = 0
	MESSAGE_TOPIC_TRANSACTION_RELAY MessageTopic = 1
	MESSAGE_TOPIC_BLOCK_SYNC MessageTopic = 2
	MESSAGE_TOPIC_LEAN_HELIX_CONSENSUS MessageTopic = 3
)

type RecipientsListMode uint16

const (
	RECIPIENT_LIST_MODE_RESERVED RecipientsListMode = 0
	RECIPIENT_LIST_MODE_SEND_TO_LIST RecipientsListMode = 1
	RECIPIENT_LIST_MODE_SEND_TO_ALL_BUT_LIST RecipientsListMode = 2
)

