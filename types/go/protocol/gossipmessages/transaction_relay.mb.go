// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package gossipmessages

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message ForwardedTransactionsMessage (non serializable)

type ForwardedTransactionsMessage struct {
	Signer *MessageSigner
	ForwardedTransactionSet *ForwardedTransactionSet
}

/////////////////////////////////////////////////////////////////////////////
// message ForwardedTransactionSet

// reader

type ForwardedTransactionSet struct {
	// Transactions []protocol.SignedTransaction

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *ForwardedTransactionSet) String() string {
	return fmt.Sprintf("{Transactions:%s,}", x.StringTransactions())
}

var _ForwardedTransactionSet_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,}
var _ForwardedTransactionSet_Unions = [][]membuffers.FieldType{}

func ForwardedTransactionSetReader(buf []byte) *ForwardedTransactionSet {
	x := &ForwardedTransactionSet{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ForwardedTransactionSet_Scheme, _ForwardedTransactionSet_Unions)
	return x
}

func (x *ForwardedTransactionSet) IsValid() bool {
	return x._message.IsValid()
}

func (x *ForwardedTransactionSet) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ForwardedTransactionSet) TransactionsIterator() *ForwardedTransactionSetTransactionsIterator {
	return &ForwardedTransactionSetTransactionsIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type ForwardedTransactionSetTransactionsIterator struct {
	iterator *membuffers.Iterator
}

func (i *ForwardedTransactionSetTransactionsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ForwardedTransactionSetTransactionsIterator) NextTransactions() *protocol.SignedTransaction {
	b, s := i.iterator.NextMessage()
	return protocol.SignedTransactionReader(b[:s])
}

func (x *ForwardedTransactionSet) RawTransactionsArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ForwardedTransactionSet) StringTransactions() (res string) {
	res = "["
	for i := x.TransactionsIterator(); i.HasNext(); {
		res += i.NextTransactions().String() + ","
	}
	res += "]"
	return
}

// builder

type ForwardedTransactionSetBuilder struct {
	Transactions []*protocol.SignedTransactionBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *ForwardedTransactionSetBuilder) arrayOfTransactions() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Transactions))
	for i, v := range w.Transactions {
		res[i] = v
	}
	return res
}

func (w *ForwardedTransactionSetBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessageArray(buf, w.arrayOfTransactions())
	if err != nil {
		return
	}
	return nil
}

func (w *ForwardedTransactionSetBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ForwardedTransactionSetBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ForwardedTransactionSetBuilder) Build() *ForwardedTransactionSet {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ForwardedTransactionSetReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type TransactionsRelayMessageType uint16

const (
	TRANSACTION_RELAY_RESERVED TransactionsRelayMessageType = 0
	TRANSACTION_RELAY_FORWARDED_TRANSACTIONS TransactionsRelayMessageType = 1
)

func (n TransactionsRelayMessageType) String() string {
	switch n {
	case TRANSACTION_RELAY_RESERVED:
		return "TRANSACTION_RELAY_RESERVED"
	case TRANSACTION_RELAY_FORWARDED_TRANSACTIONS:
		return "TRANSACTION_RELAY_FORWARDED_TRANSACTIONS"
	}
	return "UNKNOWN"
}

