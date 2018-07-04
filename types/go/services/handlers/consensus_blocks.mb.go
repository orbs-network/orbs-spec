// AUTO GENERATED FILE (by membufc proto compiler)
package handlers

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusBlocksHandler

type ConsensusBlocksHandler interface {
	HandleTransactionsBlock(*HandleTransactionsBlockInput) (*HandleTransactionsBlockOutput, error)
	HandleResultsBlock(*HandleResultsBlockInput) (*HandleResultsBlockOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionsBlockInput

// reader

type HandleTransactionsBlockInput struct {
	message membuffers.Message
}

var m_HandleTransactionsBlockInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}
var m_HandleTransactionsBlockInput_Unions = [][]membuffers.FieldType{}

func HandleTransactionsBlockInputReader(buf []byte) *HandleTransactionsBlockInput {
	x := &HandleTransactionsBlockInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleTransactionsBlockInput_Scheme, m_HandleTransactionsBlockInput_Unions)
	return x
}

func (x *HandleTransactionsBlockInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleTransactionsBlockInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleTransactionsBlockInput) TransactionsBlockHeader() *protocol.TransactionsBlockHeader {
	b, s := x.message.GetMessage(0)
	return protocol.TransactionsBlockHeaderReader(b[:s])
}

func (x *HandleTransactionsBlockInput) RawTransactionsBlockHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleTransactionsBlockInput) TransactionsBlockMetaData() *protocol.TransactionsBlockMetaData {
	b, s := x.message.GetMessage(1)
	return protocol.TransactionsBlockMetaDataReader(b[:s])
}

func (x *HandleTransactionsBlockInput) RawTransactionsBlockMetaData() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *HandleTransactionsBlockInput) TransactionsBlockProof() *protocol.TransactionsBlockProof {
	b, s := x.message.GetMessage(2)
	return protocol.TransactionsBlockProofReader(b[:s])
}

func (x *HandleTransactionsBlockInput) RawTransactionsBlockProof() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *HandleTransactionsBlockInput) PreviousCommittedTransactionsBlockHeader() *protocol.TransactionsBlockHeader {
	b, s := x.message.GetMessage(3)
	return protocol.TransactionsBlockHeaderReader(b[:s])
}

func (x *HandleTransactionsBlockInput) RawPreviousCommittedTransactionsBlockHeader() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *HandleTransactionsBlockInput) PreviousCommittedTransactionsBlockMetaData() *protocol.TransactionsBlockMetaData {
	b, s := x.message.GetMessage(4)
	return protocol.TransactionsBlockMetaDataReader(b[:s])
}

func (x *HandleTransactionsBlockInput) RawPreviousCommittedTransactionsBlockMetaData() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *HandleTransactionsBlockInput) PreviousCommittedTransactionsBlockProof() *protocol.TransactionsBlockProof {
	b, s := x.message.GetMessage(5)
	return protocol.TransactionsBlockProofReader(b[:s])
}

func (x *HandleTransactionsBlockInput) RawPreviousCommittedTransactionsBlockProof() []byte {
	return x.message.RawBufferForField(5, 0)
}

func (x *HandleTransactionsBlockInput) PreviousCommittedResultsBlockProof() *protocol.ResultsBlockProof {
	b, s := x.message.GetMessage(6)
	return protocol.ResultsBlockProofReader(b[:s])
}

func (x *HandleTransactionsBlockInput) RawPreviousCommittedResultsBlockProof() []byte {
	return x.message.RawBufferForField(6, 0)
}

// builder

type HandleTransactionsBlockInputBuilder struct {
	builder membuffers.Builder
	TransactionsBlockHeader *protocol.TransactionsBlockHeaderBuilder
	TransactionsBlockMetaData *protocol.TransactionsBlockMetaDataBuilder
	TransactionsBlockProof *protocol.TransactionsBlockProofBuilder
	PreviousCommittedTransactionsBlockHeader *protocol.TransactionsBlockHeaderBuilder
	PreviousCommittedTransactionsBlockMetaData *protocol.TransactionsBlockMetaDataBuilder
	PreviousCommittedTransactionsBlockProof *protocol.TransactionsBlockProofBuilder
	PreviousCommittedResultsBlockProof *protocol.ResultsBlockProofBuilder
}

func (w *HandleTransactionsBlockInputBuilder) Write(buf []byte) (err error) {
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

func (w *HandleTransactionsBlockInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleTransactionsBlockInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleTransactionsBlockInputBuilder) Build() *HandleTransactionsBlockInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleTransactionsBlockInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleTransactionsBlockOutput

// reader

type HandleTransactionsBlockOutput struct {
	message membuffers.Message
}

var m_HandleTransactionsBlockOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,}
var m_HandleTransactionsBlockOutput_Unions = [][]membuffers.FieldType{}

func HandleTransactionsBlockOutputReader(buf []byte) *HandleTransactionsBlockOutput {
	x := &HandleTransactionsBlockOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleTransactionsBlockOutput_Scheme, m_HandleTransactionsBlockOutput_Unions)
	return x
}

func (x *HandleTransactionsBlockOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleTransactionsBlockOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleTransactionsBlockOutput) Status() ValidateBlockConsensus {
	return ValidateBlockConsensus(x.message.GetUint16(0))
}

func (x *HandleTransactionsBlockOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleTransactionsBlockOutput) MutateStatus(v ValidateBlockConsensus) error {
	return x.message.SetUint16(0, uint16(v))
}

// builder

type HandleTransactionsBlockOutputBuilder struct {
	builder membuffers.Builder
	Status ValidateBlockConsensus
}

func (w *HandleTransactionsBlockOutputBuilder) Write(buf []byte) (err error) {
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

func (w *HandleTransactionsBlockOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleTransactionsBlockOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleTransactionsBlockOutputBuilder) Build() *HandleTransactionsBlockOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleTransactionsBlockOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleResultsBlockInput

// reader

type HandleResultsBlockInput struct {
	message membuffers.Message
}

var m_HandleResultsBlockInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}
var m_HandleResultsBlockInput_Unions = [][]membuffers.FieldType{}

func HandleResultsBlockInputReader(buf []byte) *HandleResultsBlockInput {
	x := &HandleResultsBlockInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleResultsBlockInput_Scheme, m_HandleResultsBlockInput_Unions)
	return x
}

func (x *HandleResultsBlockInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleResultsBlockInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleResultsBlockInput) ResultsBlockHeader() *protocol.ResultsBlockHeader {
	b, s := x.message.GetMessage(0)
	return protocol.ResultsBlockHeaderReader(b[:s])
}

func (x *HandleResultsBlockInput) RawResultsBlockHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleResultsBlockInput) ResultsBlockProof() *protocol.ResultsBlockProof {
	b, s := x.message.GetMessage(1)
	return protocol.ResultsBlockProofReader(b[:s])
}

func (x *HandleResultsBlockInput) RawResultsBlockProof() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *HandleResultsBlockInput) PreviousResultsBlockHeader() *protocol.ResultsBlockHeader {
	b, s := x.message.GetMessage(2)
	return protocol.ResultsBlockHeaderReader(b[:s])
}

func (x *HandleResultsBlockInput) RawPreviousResultsBlockHeader() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *HandleResultsBlockInput) PreviousCommittedTransactionsBlockMetaData() *protocol.TransactionsBlockMetaData {
	b, s := x.message.GetMessage(3)
	return protocol.TransactionsBlockMetaDataReader(b[:s])
}

func (x *HandleResultsBlockInput) RawPreviousCommittedTransactionsBlockMetaData() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *HandleResultsBlockInput) PreviousCommittedTransactionsBlockProof() *protocol.TransactionsBlockProof {
	b, s := x.message.GetMessage(4)
	return protocol.TransactionsBlockProofReader(b[:s])
}

func (x *HandleResultsBlockInput) RawPreviousCommittedTransactionsBlockProof() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *HandleResultsBlockInput) PreviousCommittedResultsBlockProof() *protocol.ResultsBlockProof {
	b, s := x.message.GetMessage(5)
	return protocol.ResultsBlockProofReader(b[:s])
}

func (x *HandleResultsBlockInput) RawPreviousCommittedResultsBlockProof() []byte {
	return x.message.RawBufferForField(5, 0)
}

// builder

type HandleResultsBlockInputBuilder struct {
	builder membuffers.Builder
	ResultsBlockHeader *protocol.ResultsBlockHeaderBuilder
	ResultsBlockProof *protocol.ResultsBlockProofBuilder
	PreviousResultsBlockHeader *protocol.ResultsBlockHeaderBuilder
	PreviousCommittedTransactionsBlockMetaData *protocol.TransactionsBlockMetaDataBuilder
	PreviousCommittedTransactionsBlockProof *protocol.TransactionsBlockProofBuilder
	PreviousCommittedResultsBlockProof *protocol.ResultsBlockProofBuilder
}

func (w *HandleResultsBlockInputBuilder) Write(buf []byte) (err error) {
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

func (w *HandleResultsBlockInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleResultsBlockInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleResultsBlockInputBuilder) Build() *HandleResultsBlockInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleResultsBlockInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message HandleResultsBlockOutput

// reader

type HandleResultsBlockOutput struct {
	message membuffers.Message
}

var m_HandleResultsBlockOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,}
var m_HandleResultsBlockOutput_Unions = [][]membuffers.FieldType{}

func HandleResultsBlockOutputReader(buf []byte) *HandleResultsBlockOutput {
	x := &HandleResultsBlockOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_HandleResultsBlockOutput_Scheme, m_HandleResultsBlockOutput_Unions)
	return x
}

func (x *HandleResultsBlockOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *HandleResultsBlockOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *HandleResultsBlockOutput) Status() ValidateBlockConsensus {
	return ValidateBlockConsensus(x.message.GetUint16(0))
}

func (x *HandleResultsBlockOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *HandleResultsBlockOutput) MutateStatus(v ValidateBlockConsensus) error {
	return x.message.SetUint16(0, uint16(v))
}

// builder

type HandleResultsBlockOutputBuilder struct {
	builder membuffers.Builder
	Status ValidateBlockConsensus
}

func (w *HandleResultsBlockOutputBuilder) Write(buf []byte) (err error) {
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

func (w *HandleResultsBlockOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *HandleResultsBlockOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *HandleResultsBlockOutputBuilder) Build() *HandleResultsBlockOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HandleResultsBlockOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type ValidateBlockConsensus uint16

const (
	VALIDATE_BLOCK_CONSENSUS_RESERVED ValidateBlockConsensus = 0
	VALIDATE_BLOCK_CONSENSUS_VALID_FOR_COMMIT ValidateBlockConsensus = 1
	VALIDATE_BLOCK_CONSENSUS_INVALID_FOR_COMMIT ValidateBlockConsensus = 2
)

