// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// message GossipMessageHeader

// reader

type GossipMessageHeader struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _GossipMessageHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint32,membuffers.TypeBytesArray,membuffers.TypeUint16,membuffers.TypeUnion,membuffers.TypeUint32,}
var _GossipMessageHeader_Unions = [][]membuffers.FieldType{{membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeUint16,}}

func GossipMessageHeaderReader(buf []byte) *GossipMessageHeader {
	x := &GossipMessageHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _GossipMessageHeader_Scheme, _GossipMessageHeader_Unions)
	return x
}

func (x *GossipMessageHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *GossipMessageHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *GossipMessageHeader) ProtocolVersion() primitives.ProtocolVersion {
	return primitives.ProtocolVersion(x._message.GetUint32(0))
}

func (x *GossipMessageHeader) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *GossipMessageHeader) MutateProtocolVersion(v primitives.ProtocolVersion) error {
	return x._message.SetUint32(0, uint32(v))
}

func (x *GossipMessageHeader) VirtualChainId() primitives.VirtualChainId {
	return primitives.VirtualChainId(x._message.GetUint32(1))
}

func (x *GossipMessageHeader) RawVirtualChainId() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *GossipMessageHeader) MutateVirtualChainId(v primitives.VirtualChainId) error {
	return x._message.SetUint32(1, uint32(v))
}

func (x *GossipMessageHeader) RecipientPublicKeysIterator() *GossipMessageHeaderRecipientPublicKeysIterator {
	return &GossipMessageHeaderRecipientPublicKeysIterator{iterator: x._message.GetBytesArrayIterator(2)}
}

type GossipMessageHeaderRecipientPublicKeysIterator struct {
	iterator *membuffers.Iterator
}

func (i *GossipMessageHeaderRecipientPublicKeysIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *GossipMessageHeaderRecipientPublicKeysIterator) NextRecipientPublicKeys() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(i.iterator.NextBytes())
}

func (x *GossipMessageHeader) RawRecipientPublicKeysArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *GossipMessageHeader) RecipientMode() RecipientsListMode {
	return RecipientsListMode(x._message.GetUint16(3))
}

func (x *GossipMessageHeader) RawRecipientMode() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *GossipMessageHeader) MutateRecipientMode(v RecipientsListMode) error {
	return x._message.SetUint16(3, uint16(v))
}

type GossipMessageHeaderType uint16

const (
	GossipMessageHeaderTypeTransactionRelayType GossipMessageHeaderType = 0
	GossipMessageHeaderTypeBlockSyncType GossipMessageHeaderType = 1
	GossipMessageHeaderTypeLeanHelixConsensusType GossipMessageHeaderType = 2
)

func (x *GossipMessageHeader) Type() GossipMessageHeaderType {
	return GossipMessageHeaderType(x._message.GetUint16(4))
}

func (x *GossipMessageHeader) IsTypeTransactionRelayType() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 0)
	return is
}

func (x *GossipMessageHeader) TransactionRelayType() gossipmessages.TransactionsRelayMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 0)
	return gossipmessages.TransactionsRelayMessageType(x._message.GetUint16InOffset(off))
}

func (x *GossipMessageHeader) MutateTransactionRelayType(v gossipmessages.TransactionsRelayMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *GossipMessageHeader) IsTypeBlockSyncType() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 1)
	return is
}

func (x *GossipMessageHeader) BlockSyncType() gossipmessages.BlockSyncMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 1)
	return gossipmessages.BlockSyncMessageType(x._message.GetUint16InOffset(off))
}

func (x *GossipMessageHeader) MutateBlockSyncType(v gossipmessages.BlockSyncMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *GossipMessageHeader) IsTypeLeanHelixConsensusType() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 2)
	return is
}

func (x *GossipMessageHeader) LeanHelixConsensusType() gossipmessages.LeanHelixMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 2)
	return gossipmessages.LeanHelixMessageType(x._message.GetUint16InOffset(off))
}

func (x *GossipMessageHeader) MutateLeanHelixConsensusType(v gossipmessages.LeanHelixMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *GossipMessageHeader) RawType() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *GossipMessageHeader) NumPayloads() uint32 {
	return x._message.GetUint32(5)
}

func (x *GossipMessageHeader) RawNumPayloads() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *GossipMessageHeader) MutateNumPayloads(v uint32) error {
	return x._message.SetUint32(5, v)
}

// builder

type GossipMessageHeaderBuilder struct {
	ProtocolVersion primitives.ProtocolVersion
	VirtualChainId primitives.VirtualChainId
	RecipientPublicKeys []primitives.Ed25519Pkey
	RecipientMode RecipientsListMode
	Type GossipMessageHeaderType
	NumPayloads uint32
	TransactionRelayType gossipmessages.TransactionsRelayMessageType
	BlockSyncType gossipmessages.BlockSyncMessageType
	LeanHelixConsensusType gossipmessages.LeanHelixMessageType

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *GossipMessageHeaderBuilder) arrayOfRecipientPublicKeys() [][]byte {
	res := make([][]byte, len(w.RecipientPublicKeys))
	for i, v := range w.RecipientPublicKeys {
		res[i] = v
	}
	return res
}

func (w *GossipMessageHeaderBuilder) Write(buf []byte) (err error) {
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
	case GossipMessageHeaderTypeTransactionRelayType:
		w._builder.WriteUint16(buf, uint16(w.TransactionRelayType))
	case GossipMessageHeaderTypeBlockSyncType:
		w._builder.WriteUint16(buf, uint16(w.BlockSyncType))
	case GossipMessageHeaderTypeLeanHelixConsensusType:
		w._builder.WriteUint16(buf, uint16(w.LeanHelixConsensusType))
	}
	w._builder.WriteUint32(buf, w.NumPayloads)
	return nil
}

func (w *GossipMessageHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *GossipMessageHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *GossipMessageHeaderBuilder) Build() *GossipMessageHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GossipMessageHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type RecipientsListMode uint16

const (
	RECIPIENT_LIST_MODE_SEND_TO_LIST RecipientsListMode = 0
	RECIPIENT_LIST_MODE_SEND_TO_ALL_BUT_LIST RecipientsListMode = 1
)

