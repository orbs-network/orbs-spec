// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package gossipmessages

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message TempKillMe

// reader

type TempKillMe struct {

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *TempKillMe) String() string {
	return fmt.Sprintf("{}")
}

var _TempKillMe_Scheme = []membuffers.FieldType{}
var _TempKillMe_Unions = [][]membuffers.FieldType{}

func TempKillMeReader(buf []byte) *TempKillMe {
	x := &TempKillMe{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _TempKillMe_Scheme, _TempKillMe_Unions)
	return x
}

func (x *TempKillMe) IsValid() bool {
	return x._message.IsValid()
}

func (x *TempKillMe) Raw() []byte {
	return x._message.RawBuffer()
}

// builder

type TempKillMeBuilder struct {

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *TempKillMeBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	return nil
}

func (w *TempKillMeBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TempKillMeBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *TempKillMeBuilder) Build() *TempKillMe {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TempKillMeReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ForwardedTransactionsMessage (non serializable)

type ForwardedTransactionsMessage struct {
	Sender *SenderSignature
	SignedTransactions []*protocol.SignedTransaction
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

