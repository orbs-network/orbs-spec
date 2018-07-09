// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package messages

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message ForwardedTransactionsHeader

// reader

type ForwardedTransactionsHeader struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _ForwardedTransactionsHeader_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var _ForwardedTransactionsHeader_Unions = [][]membuffers.FieldType{}

func ForwardedTransactionsHeaderReader(buf []byte) *ForwardedTransactionsHeader {
	x := &ForwardedTransactionsHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ForwardedTransactionsHeader_Scheme, _ForwardedTransactionsHeader_Unions)
	return x
}

func (x *ForwardedTransactionsHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *ForwardedTransactionsHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ForwardedTransactionsHeader) GwNodePublicKey() primitives.Ed25519Pkey {
	return x._message.GetBytes(0)
}

func (x *ForwardedTransactionsHeader) RawGwNodePublicKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ForwardedTransactionsHeader) MutateGwNodePublicKey(v primitives.Ed25519Pkey) error {
	return x._message.SetBytes(0, v)
}

func (x *ForwardedTransactionsHeader) Signature() primitives.Ed25519Sig {
	return x._message.GetBytes(1)
}

func (x *ForwardedTransactionsHeader) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *ForwardedTransactionsHeader) MutateSignature(v primitives.Ed25519Sig) error {
	return x._message.SetBytes(1, v)
}

// builder

type ForwardedTransactionsHeaderBuilder struct {
	GwNodePublicKey primitives.Ed25519Pkey
	Signature primitives.Ed25519Sig

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *ForwardedTransactionsHeaderBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *ForwardedTransactionsHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ForwardedTransactionsHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ForwardedTransactionsHeaderBuilder) Build() *ForwardedTransactionsHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ForwardedTransactionsHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type TransactionsRelayMessageType uint16

const (
	TRANSACTION_RELAY_RESERVED TransactionsRelayMessageType = 0
	TRANSACTION_RELAY_FORWARDED_TRANSACTIONS TransactionsRelayMessageType = 1
)

