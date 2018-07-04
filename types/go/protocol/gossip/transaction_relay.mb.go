// AUTO GENERATED FILE (by membufc proto compiler)
package gossip

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

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

func (x *ForwardedTransactionsMessage) Body() *ForwardedTransactionsSignedFields {
	b, s := x.message.GetMessage(0)
	return ForwardedTransactionsSignedFieldsReader(b[:s])
}

func (x *ForwardedTransactionsMessage) RawBody() []byte {
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
	Body *ForwardedTransactionsSignedFieldsBuilder
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
	err = w.builder.WriteMessage(buf, w.Body)
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

var m_ForwardedTransactionsSignedFields_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeMessageArray,}
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

func (x *ForwardedTransactionsSignedFields) GwNodePublicKey() []byte {
	return x.message.GetBytes(0)
}

func (x *ForwardedTransactionsSignedFields) RawGwNodePublicKey() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ForwardedTransactionsSignedFields) MutateGwNodePublicKey(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *ForwardedTransactionsSignedFields) TransactionIterator() *ForwardedTransactionsSignedFieldsTransactionIterator {
	return &ForwardedTransactionsSignedFieldsTransactionIterator{iterator: x.message.GetMessageArrayIterator(1)}
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
	return x.message.RawBufferForField(1, 0)
}

// builder

type ForwardedTransactionsSignedFieldsBuilder struct {
	builder membuffers.Builder
	GwNodePublicKey []byte
	Transaction []*protocol.SignedTransactionBuilder
}

func (w *ForwardedTransactionsSignedFieldsBuilder) arrayOfTransaction() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.Transaction))
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
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.GwNodePublicKey)
	err = w.builder.WriteMessageArray(buf, w.arrayOfTransaction())
	if err != nil {
		return
	}
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

