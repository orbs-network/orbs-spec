// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/messages"
)

/////////////////////////////////////////////////////////////////////////////
// message MessageHeader

// reader

type MessageHeader struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _MessageHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint32,membuffers.TypeBytesArray,membuffers.TypeUint16,membuffers.TypeUnion,}
var _MessageHeader_Unions = [][]membuffers.FieldType{{membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeUint16,}}

func MessageHeaderReader(buf []byte) *MessageHeader {
	x := &MessageHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _MessageHeader_Scheme, _MessageHeader_Unions)
	return x
}

func (x *MessageHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *MessageHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *MessageHeader) ProtocolVersion() primitives.ProtocolVersion {
	return primitives.ProtocolVersion(x._message.GetUint32(0))
}

func (x *MessageHeader) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *MessageHeader) MutateProtocolVersion(v primitives.ProtocolVersion) error {
	return x._message.SetUint32(0, uint32(v))
}

func (x *MessageHeader) VirtualChainId() primitives.VirtualChainId {
	return primitives.VirtualChainId(x._message.GetUint32(1))
}

func (x *MessageHeader) RawVirtualChainId() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *MessageHeader) MutateVirtualChainId(v primitives.VirtualChainId) error {
	return x._message.SetUint32(1, uint32(v))
}

func (x *MessageHeader) RecipientPublicKeysIterator() *MessageHeaderRecipientPublicKeysIterator {
	return &MessageHeaderRecipientPublicKeysIterator{iterator: x._message.GetBytesArrayIterator(2)}
}

type MessageHeaderRecipientPublicKeysIterator struct {
	iterator *membuffers.Iterator
}

func (i *MessageHeaderRecipientPublicKeysIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *MessageHeaderRecipientPublicKeysIterator) NextRecipientPublicKeys() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(i.iterator.NextBytes())
}

func (x *MessageHeader) RawRecipientPublicKeysArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *MessageHeader) RecipientMode() RecipientsListMode {
	return RecipientsListMode(x._message.GetUint16(3))
}

func (x *MessageHeader) RawRecipientMode() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *MessageHeader) MutateRecipientMode(v RecipientsListMode) error {
	return x._message.SetUint16(3, uint16(v))
}

type MessageHeaderType uint16

const (
	MessageHeaderTypeTransactionRelay MessageHeaderType = 0
	MessageHeaderTypeBlockSync MessageHeaderType = 1
	MessageHeaderTypeLeanHelixConsensus MessageHeaderType = 2
)

func (x *MessageHeader) Type() MessageHeaderType {
	return MessageHeaderType(x._message.GetUint16(4))
}

func (x *MessageHeader) IsTypeTransactionRelay() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 0)
	return is
}

func (x *MessageHeader) TransactionRelay() messages.TransactionsRelayMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 0)
	return messages.TransactionsRelayMessageType(x._message.GetUint16InOffset(off))
}

func (x *MessageHeader) MutateTransactionRelay(v messages.TransactionsRelayMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *MessageHeader) IsTypeBlockSync() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 1)
	return is
}

func (x *MessageHeader) BlockSync() messages.BlockSyncMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 1)
	return messages.BlockSyncMessageType(x._message.GetUint16InOffset(off))
}

func (x *MessageHeader) MutateBlockSync(v messages.BlockSyncMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *MessageHeader) IsTypeLeanHelixConsensus() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 2)
	return is
}

func (x *MessageHeader) LeanHelixConsensus() messages.LeanHelixMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 2)
	return messages.LeanHelixMessageType(x._message.GetUint16InOffset(off))
}

func (x *MessageHeader) MutateLeanHelixConsensus(v messages.LeanHelixMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *MessageHeader) RawType() []byte {
	return x._message.RawBufferForField(4, 0)
}

// builder

type MessageHeaderBuilder struct {
	ProtocolVersion primitives.ProtocolVersion
	VirtualChainId primitives.VirtualChainId
	RecipientPublicKeys []primitives.Ed25519Pkey
	RecipientMode RecipientsListMode
	Type MessageHeaderType
	TransactionRelay messages.TransactionsRelayMessageType
	BlockSync messages.BlockSyncMessageType
	LeanHelixConsensus messages.LeanHelixMessageType

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *MessageHeaderBuilder) arrayOfRecipientPublicKeys() [][]byte {
	res := make([][]byte, len(w.RecipientPublicKeys))
	for i, v := range w.RecipientPublicKeys {
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
	w._builder.Reset()
	w._builder.WriteUint32(buf, uint32(w.ProtocolVersion))
	w._builder.WriteUint32(buf, uint32(w.VirtualChainId))
	w._builder.WriteBytesArray(buf, w.arrayOfRecipientPublicKeys())
	w._builder.WriteUint16(buf, uint16(w.RecipientMode))
	w._builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case MessageHeaderTypeTransactionRelay:
		w._builder.WriteUint16(buf, uint16(w.TransactionRelay))
	case MessageHeaderTypeBlockSync:
		w._builder.WriteUint16(buf, uint16(w.BlockSync))
	case MessageHeaderTypeLeanHelixConsensus:
		w._builder.WriteUint16(buf, uint16(w.LeanHelixConsensus))
	}
	return nil
}

func (w *MessageHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *MessageHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
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

type RecipientsListMode uint16

const (
	RECIPIENT_LIST_MODE_SEND_TO_LIST RecipientsListMode = 0
	RECIPIENT_LIST_MODE_SEND_TO_ALL_BUT_LIST RecipientsListMode = 1
)

