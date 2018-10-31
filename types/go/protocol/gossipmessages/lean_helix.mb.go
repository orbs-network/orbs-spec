// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package gossipmessages

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/consensus"
)

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixMessage (non serializable)

type LeanHelixMessage struct {
	MessageType consensus.LeanHelixMessageType
	Content     primitives.LeanHelixMessageContent
	BlockPair   *protocol.BlockPairContainer
}

func (x *LeanHelixMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{MessageType:%s,Content:%s,BlockPair:%s,}", x.StringMessageType(), x.StringContent(), x.StringBlockPair())
}

func (x *LeanHelixMessage) StringMessageType() (res string) {
	res = fmt.Sprintf("%x", x.MessageType)
	return
}

func (x *LeanHelixMessage) StringContent() (res string) {
	res = fmt.Sprintf("%s", x.Content)
	return
}

func (x *LeanHelixMessage) StringBlockPair() (res string) {
	res = x.BlockPair.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixRandomSeedShare

// reader

type LeanHelixRandomSeedShare struct {
	// RandomSeedShare primitives.Bls1Sig

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixRandomSeedShare) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RandomSeedShare:%s,}", x.StringRandomSeedShare())
}

var _LeanHelixRandomSeedShare_Scheme = []membuffers.FieldType{membuffers.TypeBytes}
var _LeanHelixRandomSeedShare_Unions = [][]membuffers.FieldType{}

func LeanHelixRandomSeedShareReader(buf []byte) *LeanHelixRandomSeedShare {
	x := &LeanHelixRandomSeedShare{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixRandomSeedShare_Scheme, _LeanHelixRandomSeedShare_Unions)
	return x
}

func (x *LeanHelixRandomSeedShare) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixRandomSeedShare) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixRandomSeedShare) Equal(y *LeanHelixRandomSeedShare) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixRandomSeedShare) RandomSeedShare() primitives.Bls1Sig {
	return primitives.Bls1Sig(x._message.GetBytes(0))
}

func (x *LeanHelixRandomSeedShare) RawRandomSeedShare() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixRandomSeedShare) MutateRandomSeedShare(v primitives.Bls1Sig) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *LeanHelixRandomSeedShare) StringRandomSeedShare() string {
	return fmt.Sprintf("%s", x.RandomSeedShare())
}

// builder

type LeanHelixRandomSeedShareBuilder struct {
	RandomSeedShare primitives.Bls1Sig

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *LeanHelixRandomSeedShareBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteBytes(buf, []byte(w.RandomSeedShare))
	return nil
}

func (w *LeanHelixRandomSeedShareBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixRandomSeedShareBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixRandomSeedShareBuilder) Build() *LeanHelixRandomSeedShare {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixRandomSeedShareReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums
