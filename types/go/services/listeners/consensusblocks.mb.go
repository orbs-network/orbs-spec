// AUTO GENERATED FILE (by membufc proto compiler)
package listeners

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusBlocksListener

type ConsensusBlocksListener interface {
	AcknowledgeTransactionsBlockConsensus(*AcknowledgeTransactionsBlockConsensusInput) (*AcknowledgeTransactionsBlockConsensusOutput, error)
	AcknowledgeResultsBlockConsensus(*AcknowledgeResultsBlockConsensusInput) (*AcknowledgeResultsBlockConsensusOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message AcknowledgeTransactionsBlockConsensusInput

// reader

type AcknowledgeTransactionsBlockConsensusInput struct {
	message membuffers.Message
}

var m_AcknowledgeTransactionsBlockConsensusInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}
var m_AcknowledgeTransactionsBlockConsensusInput_Unions = [][]membuffers.FieldType{}

func AcknowledgeTransactionsBlockConsensusInputReader(buf []byte) *AcknowledgeTransactionsBlockConsensusInput {
	x := &AcknowledgeTransactionsBlockConsensusInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_AcknowledgeTransactionsBlockConsensusInput_Scheme, m_AcknowledgeTransactionsBlockConsensusInput_Unions)
	return x
}

func (x *AcknowledgeTransactionsBlockConsensusInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *AcknowledgeTransactionsBlockConsensusInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *AcknowledgeTransactionsBlockConsensusInput) TransactionsBlockHeader() *protocol.TransactionsBlockHeader {
	b, s := x.message.GetMessage(0)
	return protocol.TransactionsBlockHeaderReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawTransactionsBlockHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusInput) TransactionsBlockMetaData() *protocol.TransactionsBlockMetaData {
	b, s := x.message.GetMessage(1)
	return protocol.TransactionsBlockMetaDataReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawTransactionsBlockMetaData() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusInput) TransactionsBlockProof() *protocol.TransactionsBlockProof {
	b, s := x.message.GetMessage(2)
	return protocol.TransactionsBlockProofReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawTransactionsBlockProof() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusInput) PreviousCommittedTransactionsBlockHeader() *protocol.TransactionsBlockHeader {
	b, s := x.message.GetMessage(3)
	return protocol.TransactionsBlockHeaderReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawPreviousCommittedTransactionsBlockHeader() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusInput) PreviousCommittedTransactionsBlockMetaData() *protocol.TransactionsBlockMetaData {
	b, s := x.message.GetMessage(4)
	return protocol.TransactionsBlockMetaDataReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawPreviousCommittedTransactionsBlockMetaData() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusInput) PreviousCommittedTransactionsBlockProof() *protocol.TransactionsBlockProof {
	b, s := x.message.GetMessage(5)
	return protocol.TransactionsBlockProofReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawPreviousCommittedTransactionsBlockProof() []byte {
	return x.message.RawBufferForField(5, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusInput) PreviousCommittedResultsBlockProof() *protocol.ResultsBlockProof {
	b, s := x.message.GetMessage(6)
	return protocol.ResultsBlockProofReader(b[:s])
}

func (x *AcknowledgeTransactionsBlockConsensusInput) RawPreviousCommittedResultsBlockProof() []byte {
	return x.message.RawBufferForField(6, 0)
}

// builder

type AcknowledgeTransactionsBlockConsensusInputBuilder struct {
	builder membuffers.Builder
	TransactionsBlockHeader *protocol.TransactionsBlockHeaderBuilder
	TransactionsBlockMetaData *protocol.TransactionsBlockMetaDataBuilder
	TransactionsBlockProof *protocol.TransactionsBlockProofBuilder
	PreviousCommittedTransactionsBlockHeader *protocol.TransactionsBlockHeaderBuilder
	PreviousCommittedTransactionsBlockMetaData *protocol.TransactionsBlockMetaDataBuilder
	PreviousCommittedTransactionsBlockProof *protocol.TransactionsBlockProofBuilder
	PreviousCommittedResultsBlockProof *protocol.ResultsBlockProofBuilder
}

func (w *AcknowledgeTransactionsBlockConsensusInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.TransactionsBlockHeader)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.TransactionsBlockMetaData)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.TransactionsBlockProof)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedTransactionsBlockHeader)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedTransactionsBlockMetaData)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedTransactionsBlockProof)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedResultsBlockProof)
	if err != nil {
		return
	}
	return nil
}

func (w *AcknowledgeTransactionsBlockConsensusInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *AcknowledgeTransactionsBlockConsensusInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *AcknowledgeTransactionsBlockConsensusInputBuilder) Build() *AcknowledgeTransactionsBlockConsensusInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return AcknowledgeTransactionsBlockConsensusInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message AcknowledgeTransactionsBlockConsensusOutput

// reader

type AcknowledgeTransactionsBlockConsensusOutput struct {
	message membuffers.Message
}

var m_AcknowledgeTransactionsBlockConsensusOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,}
var m_AcknowledgeTransactionsBlockConsensusOutput_Unions = [][]membuffers.FieldType{}

func AcknowledgeTransactionsBlockConsensusOutputReader(buf []byte) *AcknowledgeTransactionsBlockConsensusOutput {
	x := &AcknowledgeTransactionsBlockConsensusOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_AcknowledgeTransactionsBlockConsensusOutput_Scheme, m_AcknowledgeTransactionsBlockConsensusOutput_Unions)
	return x
}

func (x *AcknowledgeTransactionsBlockConsensusOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *AcknowledgeTransactionsBlockConsensusOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *AcknowledgeTransactionsBlockConsensusOutput) Status() ValidateBlockConsensus {
	return ValidateBlockConsensus(x.message.GetUint16(0))
}

func (x *AcknowledgeTransactionsBlockConsensusOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *AcknowledgeTransactionsBlockConsensusOutput) MutateStatus(v ValidateBlockConsensus) error {
	return x.message.SetUint16(0, uint16(v))
}

// builder

type AcknowledgeTransactionsBlockConsensusOutputBuilder struct {
	builder membuffers.Builder
	Status ValidateBlockConsensus
}

func (w *AcknowledgeTransactionsBlockConsensusOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.Status))
	return nil
}

func (w *AcknowledgeTransactionsBlockConsensusOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *AcknowledgeTransactionsBlockConsensusOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *AcknowledgeTransactionsBlockConsensusOutputBuilder) Build() *AcknowledgeTransactionsBlockConsensusOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return AcknowledgeTransactionsBlockConsensusOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message AcknowledgeResultsBlockConsensusInput

// reader

type AcknowledgeResultsBlockConsensusInput struct {
	message membuffers.Message
}

var m_AcknowledgeResultsBlockConsensusInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}
var m_AcknowledgeResultsBlockConsensusInput_Unions = [][]membuffers.FieldType{}

func AcknowledgeResultsBlockConsensusInputReader(buf []byte) *AcknowledgeResultsBlockConsensusInput {
	x := &AcknowledgeResultsBlockConsensusInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_AcknowledgeResultsBlockConsensusInput_Scheme, m_AcknowledgeResultsBlockConsensusInput_Unions)
	return x
}

func (x *AcknowledgeResultsBlockConsensusInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *AcknowledgeResultsBlockConsensusInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *AcknowledgeResultsBlockConsensusInput) ResultsBlockHeader() *protocol.ResultsBlockHeader {
	b, s := x.message.GetMessage(0)
	return protocol.ResultsBlockHeaderReader(b[:s])
}

func (x *AcknowledgeResultsBlockConsensusInput) RawResultsBlockHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *AcknowledgeResultsBlockConsensusInput) ResultsBlockProof() *protocol.ResultsBlockProof {
	b, s := x.message.GetMessage(1)
	return protocol.ResultsBlockProofReader(b[:s])
}

func (x *AcknowledgeResultsBlockConsensusInput) RawResultsBlockProof() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *AcknowledgeResultsBlockConsensusInput) PreviousResultsBlockHeader() *protocol.ResultsBlockHeader {
	b, s := x.message.GetMessage(2)
	return protocol.ResultsBlockHeaderReader(b[:s])
}

func (x *AcknowledgeResultsBlockConsensusInput) RawPreviousResultsBlockHeader() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *AcknowledgeResultsBlockConsensusInput) PreviousCommittedTransactionsBlockMetaData() *protocol.TransactionsBlockMetaData {
	b, s := x.message.GetMessage(3)
	return protocol.TransactionsBlockMetaDataReader(b[:s])
}

func (x *AcknowledgeResultsBlockConsensusInput) RawPreviousCommittedTransactionsBlockMetaData() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *AcknowledgeResultsBlockConsensusInput) PreviousCommittedTransactionsBlockProof() *protocol.TransactionsBlockProof {
	b, s := x.message.GetMessage(4)
	return protocol.TransactionsBlockProofReader(b[:s])
}

func (x *AcknowledgeResultsBlockConsensusInput) RawPreviousCommittedTransactionsBlockProof() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *AcknowledgeResultsBlockConsensusInput) PreviousCommittedResultsBlockProof() *protocol.ResultsBlockProof {
	b, s := x.message.GetMessage(5)
	return protocol.ResultsBlockProofReader(b[:s])
}

func (x *AcknowledgeResultsBlockConsensusInput) RawPreviousCommittedResultsBlockProof() []byte {
	return x.message.RawBufferForField(5, 0)
}

// builder

type AcknowledgeResultsBlockConsensusInputBuilder struct {
	builder membuffers.Builder
	ResultsBlockHeader *protocol.ResultsBlockHeaderBuilder
	ResultsBlockProof *protocol.ResultsBlockProofBuilder
	PreviousResultsBlockHeader *protocol.ResultsBlockHeaderBuilder
	PreviousCommittedTransactionsBlockMetaData *protocol.TransactionsBlockMetaDataBuilder
	PreviousCommittedTransactionsBlockProof *protocol.TransactionsBlockProofBuilder
	PreviousCommittedResultsBlockProof *protocol.ResultsBlockProofBuilder
}

func (w *AcknowledgeResultsBlockConsensusInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.ResultsBlockHeader)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.ResultsBlockProof)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousResultsBlockHeader)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedTransactionsBlockMetaData)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedTransactionsBlockProof)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.PreviousCommittedResultsBlockProof)
	if err != nil {
		return
	}
	return nil
}

func (w *AcknowledgeResultsBlockConsensusInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *AcknowledgeResultsBlockConsensusInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *AcknowledgeResultsBlockConsensusInputBuilder) Build() *AcknowledgeResultsBlockConsensusInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return AcknowledgeResultsBlockConsensusInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message AcknowledgeResultsBlockConsensusOutput

// reader

type AcknowledgeResultsBlockConsensusOutput struct {
	message membuffers.Message
}

var m_AcknowledgeResultsBlockConsensusOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,}
var m_AcknowledgeResultsBlockConsensusOutput_Unions = [][]membuffers.FieldType{}

func AcknowledgeResultsBlockConsensusOutputReader(buf []byte) *AcknowledgeResultsBlockConsensusOutput {
	x := &AcknowledgeResultsBlockConsensusOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_AcknowledgeResultsBlockConsensusOutput_Scheme, m_AcknowledgeResultsBlockConsensusOutput_Unions)
	return x
}

func (x *AcknowledgeResultsBlockConsensusOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *AcknowledgeResultsBlockConsensusOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *AcknowledgeResultsBlockConsensusOutput) Status() ValidateBlockConsensus {
	return ValidateBlockConsensus(x.message.GetUint16(0))
}

func (x *AcknowledgeResultsBlockConsensusOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *AcknowledgeResultsBlockConsensusOutput) MutateStatus(v ValidateBlockConsensus) error {
	return x.message.SetUint16(0, uint16(v))
}

// builder

type AcknowledgeResultsBlockConsensusOutputBuilder struct {
	builder membuffers.Builder
	Status ValidateBlockConsensus
}

func (w *AcknowledgeResultsBlockConsensusOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint16(buf, uint16(w.Status))
	return nil
}

func (w *AcknowledgeResultsBlockConsensusOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *AcknowledgeResultsBlockConsensusOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *AcknowledgeResultsBlockConsensusOutputBuilder) Build() *AcknowledgeResultsBlockConsensusOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return AcknowledgeResultsBlockConsensusOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type ValidateBlockConsensus uint16

const (
	VALIDATE_BLOCK_CONSENSUS_RESERVED ValidateBlockConsensus = 0
	VALIDATE_BLOCK_CONSENSUS_VALID_FOR_COMMIT ValidateBlockConsensus = 1
	VALIDATE_BLOCK_CONSENSUS_INVALID_FOR_COMMIT ValidateBlockConsensus = 2
)

