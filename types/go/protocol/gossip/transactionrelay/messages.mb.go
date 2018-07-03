// AUTO GENERATED FILE (by membufc proto compiler)
package transactionrelay

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message Message

// reader

type Message struct {
	message membuffers.Message
}

var m_Message_Scheme = []membuffers.FieldType{membuffers.TypeUnion,}
var m_Message_Unions = [][]membuffers.FieldType{{membuffers.TypeMessage,}}

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

type MessageTransactionsRelayMessage uint16

const (
	MessageTransactionsRelayMessageForwardedTransactions MessageTransactionsRelayMessage = 0
)

func (x *Message) TransactionsRelayMessage() MessageTransactionsRelayMessage {
	return MessageTransactionsRelayMessage(x.message.GetUint16(0))
}

func (x *Message) IsTransactionsRelayMessageForwardedTransactions() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 0)
	return is
}

func (x *Message) TransactionsRelayMessageForwardedTransactions() ForwardedTransactionsMessage {
	_, off := x.message.IsUnionIndex(0, 0, 0)
	return x.message.GetMessageInOffset(off)
}



func (x *Message) RawTransactionsRelayMessage() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type MessageBuilder struct {
	builder membuffers.Builder
	TransactionsRelayMessage MessageTransactionsRelayMessage
	ForwardedTransactions ForwardedTransactionsMessage
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
	w.builder.WriteUnionIndex(buf, uint16(w.TransactionsRelayMessage))
	switch w.TransactionsRelayMessage {
	case MessageTransactionsRelayMessageForwardedTransactions:
		w.builder.WriteMessage(buf, w.ForwardedTransactions)
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
// message ForwardedTransactionsMessage

// reader

type ForwardedTransactionsMessage struct {
	message membuffers.Message
}

var m_ForwardedTransactionsMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var m_ForwardedTransactionsMessage_Unions = [][]membuffers.FieldType{}

func ForwardedTransactionsMessageReader(buf []byte) *ForwardedTransactionsMessage {
	x := &ForwardedTransactionsMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ForwardedTransactionsMessage_Scheme, m_ForwardedTransactionsMessage_Unions)
	return x
}

func (x *ForwardedTransactionsMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *ForwardedTransactionsMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ForwardedTransactionsMessage) ForwardedTransactionsSignedFields() *ForwardedTransactionsSignedFields {
	b, s := x.message.GetMessage(0)
	return ForwardedTransactionsSignedFieldsReader(b[:s])
}

func (x *ForwardedTransactionsMessage) RawForwardedTransactionsSignedFields() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ForwardedTransactionsMessage) Signature() []byte {
	return x.message.GetBytes(1)
}

func (x *ForwardedTransactionsMessage) RawSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ForwardedTransactionsMessage) MutateSignature(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type ForwardedTransactionsMessageBuilder struct {
	builder membuffers.Builder
	ForwardedTransactionsSignedFields *ForwardedTransactionsSignedFieldsBuilder
	Signature []byte
}

func (w *ForwardedTransactionsMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.ForwardedTransactionsSignedFields)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *ForwardedTransactionsMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ForwardedTransactionsMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ForwardedTransactionsMessageBuilder) Build() *ForwardedTransactionsMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ForwardedTransactionsMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ForwardedTransactionsSignedFields

// reader

type ForwardedTransactionsSignedFields struct {
	message membuffers.Message
}

var m_ForwardedTransactionsSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeBytesArray,}
var m_ForwardedTransactionsSignedFields_Unions = [][]membuffers.FieldType{}

func ForwardedTransactionsSignedFieldsReader(buf []byte) *ForwardedTransactionsSignedFields {
	x := &ForwardedTransactionsSignedFields{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ForwardedTransactionsSignedFields_Scheme, m_ForwardedTransactionsSignedFields_Unions)
	return x
}

func (x *ForwardedTransactionsSignedFields) IsValid() bool {
	return x.message.IsValid()
}

func (x *ForwardedTransactionsSignedFields) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ForwardedTransactionsSignedFields) MessageType() *ForwardedTransactionsSignedFields {
	b, s := x.message.GetMessage(0)
	return ForwardedTransactionsSignedFieldsReader(b[:s])
}

func (x *ForwardedTransactionsSignedFields) RawMessageType() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ForwardedTransactionsSignedFields) GwNodePublicKey() []byte {
	return x.message.GetBytes(1)
}

func (x *ForwardedTransactionsSignedFields) RawGwNodePublicKey() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ForwardedTransactionsSignedFields) MutateGwNodePublicKey(v []byte) error {
	return x.message.SetBytes(1, v)
}

func (x *ForwardedTransactionsSignedFields) TransactionDataIterator() *ForwardedTransactionsSignedFieldsTransactionDataIterator {
	return &ForwardedTransactionsSignedFieldsTransactionDataIterator{iterator: x.message.GetBytesArrayIterator(2)}
}

type ForwardedTransactionsSignedFieldsTransactionDataIterator struct {
	iterator *membuffers.Iterator
}

func (i *ForwardedTransactionsSignedFieldsTransactionDataIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ForwardedTransactionsSignedFieldsTransactionDataIterator) NextTransactionData() []byte {
	return i.iterator.NextBytes()
}

func (x *ForwardedTransactionsSignedFields) RawTransactionDataArray() []byte {
	return x.message.RawBufferForField(2, 0)
}

// builder

type ForwardedTransactionsSignedFieldsBuilder struct {
	builder membuffers.Builder
	MessageType *ForwardedTransactionsSignedFieldsBuilder
	GwNodePublicKey []byte
	TransactionData [][]byte
}

func (w *ForwardedTransactionsSignedFieldsBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.MessageType)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.GwNodePublicKey)
	w.builder.WriteBytesArray(buf, w.TransactionData)
	return nil
}

func (w *ForwardedTransactionsSignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ForwardedTransactionsSignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ForwardedTransactionsSignedFieldsBuilder) Build() *ForwardedTransactionsSignedFields {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ForwardedTransactionsSignedFieldsReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type TransactionsRelayMessageType uint16

const (
	TRANSACTION_RELAY_RESERVED TransactionsRelayMessageType = 0
	TRANSACTION_RELAY_FORWARDED_TRANSACTIONS TransactionsRelayMessageType = 1
)

