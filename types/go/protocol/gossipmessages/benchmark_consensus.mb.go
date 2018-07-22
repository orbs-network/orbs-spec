// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package gossipmessages

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"bytes"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/consensus"
)

/////////////////////////////////////////////////////////////////////////////
// message TempKillMeBenchmarkConsensus

// reader

type TempKillMeBenchmarkConsensus struct {

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *TempKillMeBenchmarkConsensus) String() string {
	return fmt.Sprintf("{}")
}

var _TempKillMeBenchmarkConsensus_Scheme = []membuffers.FieldType{}
var _TempKillMeBenchmarkConsensus_Unions = [][]membuffers.FieldType{}

func TempKillMeBenchmarkConsensusReader(buf []byte) *TempKillMeBenchmarkConsensus {
	x := &TempKillMeBenchmarkConsensus{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _TempKillMeBenchmarkConsensus_Scheme, _TempKillMeBenchmarkConsensus_Unions)
	return x
}

func (x *TempKillMeBenchmarkConsensus) IsValid() bool {
	return x._message.IsValid()
}

func (x *TempKillMeBenchmarkConsensus) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *TempKillMeBenchmarkConsensus) Equal(y *TempKillMeBenchmarkConsensus) bool {
  if x == nil && y == nil {
    return true
  }
  if x == nil || y == nil {
    return false
  }
  return bytes.Equal(x.Raw(), y.Raw())
}

// builder

type TempKillMeBenchmarkConsensusBuilder struct {

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *TempKillMeBenchmarkConsensusBuilder) Write(buf []byte) (err error) {
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

func (w *TempKillMeBenchmarkConsensusBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TempKillMeBenchmarkConsensusBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *TempKillMeBenchmarkConsensusBuilder) Build() *TempKillMeBenchmarkConsensus {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TempKillMeBenchmarkConsensusReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message BenchmarkConsensusCommitMessage (non serializable)

type BenchmarkConsensusCommitMessage struct {
	BlockPair *protocol.BlockPairContainer
}

/////////////////////////////////////////////////////////////////////////////
// message BenchmarkConsensusCommittedMessage (non serializable)

type BenchmarkConsensusCommittedMessage struct {
	LastCommittedBlockHeight primitives.BlockHeight
	Sender *consensus.LeanHelixSenderSignature
}

/////////////////////////////////////////////////////////////////////////////
// enums

