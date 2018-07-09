// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package gossip

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message ForwardedTransactionsMessage

// reader

type ForwardedTransactionsMessage struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _ForwardedTransactionsMessage_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var _ForwardedTransactionsMessage_Unions = [][]membuffers.FieldType{}

func ForwardedTransactionsMessageReader(buf []byte) *ForwardedTransactionsMessage {
	x := &ForwardedTransactionsMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ForwardedTransactionsMessage_Scheme, _ForwardedTransactionsMessage_Unions)
	return x
}

func (x *ForwardedTransactionsMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *ForwardedTransactionsMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ForwardedTransactionsMessage) Body() *ForwardedTransactionsSignedFields {
	b, s := x._message.GetMessage(0)
	return ForwardedTransactionsSignedFieldsReader(b[:s])
}

func (x *ForwardedTransactionsMessage) RawBody() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ForwardedTransactionsMessage) Signature() primitives.Ed25519Sig {
	return x._message.GetBytes(1)
}

func (x *ForwardedTransactionsMessage) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *ForwardedTransactionsMessage) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, v)
}

// builder

type ForwardedTransactionsMessageBuilder struct {
	Body *ForwardedTransactionsSignedFieldsBuilder
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
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
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.Body)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *ForwardedTransactionsMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ForwardedTransactionsMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
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
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _ForwardedTransactionsSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessageArray,}
var _ForwardedTransactionsSignedFields_Unions = [][]membuffers.FieldType{}

func ForwardedTransactionsSignedFieldsReader(buf []byte) *ForwardedTransactionsSignedFields {
	x := &ForwardedTransactionsSignedFields{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ForwardedTransactionsSignedFields_Scheme, _ForwardedTransactionsSignedFields_Unions)
	return x
}

func (x *ForwardedTransactionsSignedFields) IsValid() bool {
	return x._message.IsValid()
}

func (x *ForwardedTransactionsSignedFields) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ForwardedTransactionsSignedFields) GwNodePublicKey() primitives.Ed25519Pkey {
	return x._message.GetBytes(0)
}

func (x *ForwardedTransactionsSignedFields) RawGwNodePublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ForwardedTransactionsSignedFields) MutateGwNodePublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, v)
}

func (x *ForwardedTransactionsSignedFields) TransactionIterator() *ForwardedTransactionsSignedFieldsTransactionIterator {
	return &ForwardedTransactionsSignedFieldsTransactionIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type ForwardedTransactionsSignedFieldsTransactionIterator struct {
	iterator *membuffers.Iterator
}

func (i *ForwardedTransactionsSignedFieldsTransactionIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ForwardedTransactionsSignedFieldsTransactionIterator) NextTransaction() *protocol.SignedTransaction {
	b, s := i.iterator.NextMessage()
	return protocol.SignedTransactionReader(b[:s])
}

func (x *ForwardedTransactionsSignedFields) RawTransactionArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

// builder

type ForwardedTransactionsSignedFieldsBuilder struct {
	GwNodePublicKey primitives.Ed25519Pkey
	Transaction []*protocol.SignedTransactionBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *ForwardedTransactionsSignedFieldsBuilder) arrayOfTransaction() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Transaction))
	for i, v := range w.Transaction {
		res[i] = v
	}
	return res
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
	w._builder.Reset()
	w._builder.WriteBytes(buf, w.GwNodePublicKey)
	err = w._builder.WriteMessageArray(buf, w.arrayOfTransaction())
	if err != nil {
		return
	}
	return nil
}

func (w *ForwardedTransactionsSignedFieldsBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ForwardedTransactionsSignedFieldsBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
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

