// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossip/blocksync"
	"github.com/orbs-network/orbs-spec/types/go/services/listeners"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockStorage

type BlockStorage interface {
	CommitBlock(*CommitBlockInput) (*CommitBlockOutput, error)
	GetTransactionsBlockHeader(*GetTransactionsBlockHeaderInput) (*GetTransactionsBlockHeaderOutput, error)
	GetResultsBlockHeader(*GetResultsBlockHeaderInput) (*GetResultsBlockHeaderOutput, error)
	GetTransactionReceipt(*GetTransactionReceiptInput) (*GetTransactionReceiptOutput, error)
	GetLastCommittedBlockHeight(*GetLastCommittedBlockHeightInput) (*GetLastCommittedBlockHeightOutput, error)
	OnBlockAvailabilityRequest(*OnBlockAvailabilityRequestInput) (*OnBlockAvailabilityRequestOutput, error)
	OnBlockAvailabilityResponse(*OnBlockAvailabilityResponseInput) (*OnBlockAvailabilityResponseOutput, error)
	OnBlockSyncRequest(*OnBlockSyncRequestInput) (*OnBlockSyncRequestOutput, error)
	OnBlockSyncResponse(*OnBlockSyncResponseInput) (*OnBlockSyncResponseOutput, error)
	ValidateBlockForCommit(*ValidateBlockForCommitInput) (*ValidateBlockForCommitOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message CommitBlockInput

// reader

type CommitBlockInput struct {
	message membuffers.Message
}

var m_CommitBlockInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_CommitBlockInput_Unions = [][]membuffers.FieldType{}

func CommitBlockInputReader(buf []byte) *CommitBlockInput {
	x := &CommitBlockInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_CommitBlockInput_Scheme, m_CommitBlockInput_Unions)
	return x
}

func (x *CommitBlockInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *CommitBlockInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *CommitBlockInput) BlockPair() *protocol.BlockPair {
	b, s := x.message.GetMessage(0)
	return protocol.BlockPairReader(b[:s])
}

func (x *CommitBlockInput) RawBlockPair() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type CommitBlockInputBuilder struct {
	builder membuffers.Builder
	BlockPair *protocol.BlockPairBuilder
}

func (w *CommitBlockInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.BlockPair)
	if err != nil {
		return
	}
	return nil
}

func (w *CommitBlockInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *CommitBlockInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *CommitBlockInputBuilder) Build() *CommitBlockInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return CommitBlockInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message CommitBlockOutput

// reader

type CommitBlockOutput struct {
	message membuffers.Message
}

var m_CommitBlockOutput_Scheme = []membuffers.FieldType{}
var m_CommitBlockOutput_Unions = [][]membuffers.FieldType{}

func CommitBlockOutputReader(buf []byte) *CommitBlockOutput {
	x := &CommitBlockOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_CommitBlockOutput_Scheme, m_CommitBlockOutput_Unions)
	return x
}

func (x *CommitBlockOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *CommitBlockOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type CommitBlockOutputBuilder struct {
	builder membuffers.Builder
}

func (w *CommitBlockOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	return nil
}

func (w *CommitBlockOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *CommitBlockOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *CommitBlockOutputBuilder) Build() *CommitBlockOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return CommitBlockOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsBlockHeaderInput

// reader

type GetTransactionsBlockHeaderInput struct {
	message membuffers.Message
}

var m_GetTransactionsBlockHeaderInput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,}
var m_GetTransactionsBlockHeaderInput_Unions = [][]membuffers.FieldType{}

func GetTransactionsBlockHeaderInputReader(buf []byte) *GetTransactionsBlockHeaderInput {
	x := &GetTransactionsBlockHeaderInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetTransactionsBlockHeaderInput_Scheme, m_GetTransactionsBlockHeaderInput_Unions)
	return x
}

func (x *GetTransactionsBlockHeaderInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetTransactionsBlockHeaderInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetTransactionsBlockHeaderInput) BlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *GetTransactionsBlockHeaderInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetTransactionsBlockHeaderInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

// builder

type GetTransactionsBlockHeaderInputBuilder struct {
	builder membuffers.Builder
	BlockHeight uint64
}

func (w *GetTransactionsBlockHeaderInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint64(buf, w.BlockHeight)
	return nil
}

func (w *GetTransactionsBlockHeaderInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetTransactionsBlockHeaderInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetTransactionsBlockHeaderInputBuilder) Build() *GetTransactionsBlockHeaderInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionsBlockHeaderInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsBlockHeaderOutput

// reader

type GetTransactionsBlockHeaderOutput struct {
	message membuffers.Message
}

var m_GetTransactionsBlockHeaderOutput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}
var m_GetTransactionsBlockHeaderOutput_Unions = [][]membuffers.FieldType{}

func GetTransactionsBlockHeaderOutputReader(buf []byte) *GetTransactionsBlockHeaderOutput {
	x := &GetTransactionsBlockHeaderOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetTransactionsBlockHeaderOutput_Scheme, m_GetTransactionsBlockHeaderOutput_Unions)
	return x
}

func (x *GetTransactionsBlockHeaderOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetTransactionsBlockHeaderOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetTransactionsBlockHeaderOutput) TransactionsBlockHeader() *protocol.TransactionsBlockHeader {
	b, s := x.message.GetMessage(0)
	return protocol.TransactionsBlockHeaderReader(b[:s])
}

func (x *GetTransactionsBlockHeaderOutput) RawTransactionsBlockHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetTransactionsBlockHeaderOutput) TransactionsBlockMetaData() *protocol.TransactionsBlockMetaData {
	b, s := x.message.GetMessage(1)
	return protocol.TransactionsBlockMetaDataReader(b[:s])
}

func (x *GetTransactionsBlockHeaderOutput) RawTransactionsBlockMetaData() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *GetTransactionsBlockHeaderOutput) TransactionsBlockProof() *protocol.TransactionsBlockProof {
	b, s := x.message.GetMessage(2)
	return protocol.TransactionsBlockProofReader(b[:s])
}

func (x *GetTransactionsBlockHeaderOutput) RawTransactionsBlockProof() []byte {
	return x.message.RawBufferForField(2, 0)
}

// builder

type GetTransactionsBlockHeaderOutputBuilder struct {
	builder membuffers.Builder
	TransactionsBlockHeader *protocol.TransactionsBlockHeaderBuilder
	TransactionsBlockMetaData *protocol.TransactionsBlockMetaDataBuilder
	TransactionsBlockProof *protocol.TransactionsBlockProofBuilder
}

func (w *GetTransactionsBlockHeaderOutputBuilder) Write(buf []byte) (err error) {
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
	return nil
}

func (w *GetTransactionsBlockHeaderOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetTransactionsBlockHeaderOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetTransactionsBlockHeaderOutputBuilder) Build() *GetTransactionsBlockHeaderOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionsBlockHeaderOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetResultsBlockHeaderInput

// reader

type GetResultsBlockHeaderInput struct {
	message membuffers.Message
}

var m_GetResultsBlockHeaderInput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,}
var m_GetResultsBlockHeaderInput_Unions = [][]membuffers.FieldType{}

func GetResultsBlockHeaderInputReader(buf []byte) *GetResultsBlockHeaderInput {
	x := &GetResultsBlockHeaderInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetResultsBlockHeaderInput_Scheme, m_GetResultsBlockHeaderInput_Unions)
	return x
}

func (x *GetResultsBlockHeaderInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetResultsBlockHeaderInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetResultsBlockHeaderInput) BlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *GetResultsBlockHeaderInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetResultsBlockHeaderInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

// builder

type GetResultsBlockHeaderInputBuilder struct {
	builder membuffers.Builder
	BlockHeight uint64
}

func (w *GetResultsBlockHeaderInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint64(buf, w.BlockHeight)
	return nil
}

func (w *GetResultsBlockHeaderInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetResultsBlockHeaderInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetResultsBlockHeaderInputBuilder) Build() *GetResultsBlockHeaderInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetResultsBlockHeaderInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetResultsBlockHeaderOutput

// reader

type GetResultsBlockHeaderOutput struct {
	message membuffers.Message
}

var m_GetResultsBlockHeaderOutput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,}
var m_GetResultsBlockHeaderOutput_Unions = [][]membuffers.FieldType{}

func GetResultsBlockHeaderOutputReader(buf []byte) *GetResultsBlockHeaderOutput {
	x := &GetResultsBlockHeaderOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetResultsBlockHeaderOutput_Scheme, m_GetResultsBlockHeaderOutput_Unions)
	return x
}

func (x *GetResultsBlockHeaderOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetResultsBlockHeaderOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetResultsBlockHeaderOutput) ResultsBlockHeader() *protocol.ResultsBlockHeader {
	b, s := x.message.GetMessage(0)
	return protocol.ResultsBlockHeaderReader(b[:s])
}

func (x *GetResultsBlockHeaderOutput) RawResultsBlockHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetResultsBlockHeaderOutput) ResultsBlockProof() *protocol.ResultsBlockProof {
	b, s := x.message.GetMessage(1)
	return protocol.ResultsBlockProofReader(b[:s])
}

func (x *GetResultsBlockHeaderOutput) RawResultsBlockProof() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type GetResultsBlockHeaderOutputBuilder struct {
	builder membuffers.Builder
	ResultsBlockHeader *protocol.ResultsBlockHeaderBuilder
	ResultsBlockProof *protocol.ResultsBlockProofBuilder
}

func (w *GetResultsBlockHeaderOutputBuilder) Write(buf []byte) (err error) {
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
	return nil
}

func (w *GetResultsBlockHeaderOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetResultsBlockHeaderOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetResultsBlockHeaderOutputBuilder) Build() *GetResultsBlockHeaderOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetResultsBlockHeaderOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionReceiptInput

// reader

type GetTransactionReceiptInput struct {
	message membuffers.Message
}

var m_GetTransactionReceiptInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeUint64,}
var m_GetTransactionReceiptInput_Unions = [][]membuffers.FieldType{}

func GetTransactionReceiptInputReader(buf []byte) *GetTransactionReceiptInput {
	x := &GetTransactionReceiptInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetTransactionReceiptInput_Scheme, m_GetTransactionReceiptInput_Unions)
	return x
}

func (x *GetTransactionReceiptInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetTransactionReceiptInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetTransactionReceiptInput) Txid() []byte {
	return x.message.GetBytes(0)
}

func (x *GetTransactionReceiptInput) RawTxid() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetTransactionReceiptInput) MutateTxid(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *GetTransactionReceiptInput) Timestamp() uint64 {
	return x.message.GetUint64(1)
}

func (x *GetTransactionReceiptInput) RawTimestamp() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *GetTransactionReceiptInput) MutateTimestamp(v uint64) error {
	return x.message.SetUint64(1, v)
}

// builder

type GetTransactionReceiptInputBuilder struct {
	builder membuffers.Builder
	Txid []byte
	Timestamp uint64
}

func (w *GetTransactionReceiptInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Txid)
	w.builder.WriteUint64(buf, w.Timestamp)
	return nil
}

func (w *GetTransactionReceiptInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetTransactionReceiptInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetTransactionReceiptInputBuilder) Build() *GetTransactionReceiptInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionReceiptInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionReceiptOutput

// reader

type GetTransactionReceiptOutput struct {
	message membuffers.Message
}

var m_GetTransactionReceiptOutput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeUint64,membuffers.TypeUint64,}
var m_GetTransactionReceiptOutput_Unions = [][]membuffers.FieldType{}

func GetTransactionReceiptOutputReader(buf []byte) *GetTransactionReceiptOutput {
	x := &GetTransactionReceiptOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetTransactionReceiptOutput_Scheme, m_GetTransactionReceiptOutput_Unions)
	return x
}

func (x *GetTransactionReceiptOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetTransactionReceiptOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetTransactionReceiptOutput) Receipt() *protocol.TransactionReceipt {
	b, s := x.message.GetMessage(0)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *GetTransactionReceiptOutput) RawReceipt() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetTransactionReceiptOutput) BlockHeight() uint64 {
	return x.message.GetUint64(1)
}

func (x *GetTransactionReceiptOutput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *GetTransactionReceiptOutput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(1, v)
}

func (x *GetTransactionReceiptOutput) BlockTimestamp() uint64 {
	return x.message.GetUint64(2)
}

func (x *GetTransactionReceiptOutput) RawBlockTimestamp() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *GetTransactionReceiptOutput) MutateBlockTimestamp(v uint64) error {
	return x.message.SetUint64(2, v)
}

// builder

type GetTransactionReceiptOutputBuilder struct {
	builder membuffers.Builder
	Receipt *protocol.TransactionReceiptBuilder
	BlockHeight uint64
	BlockTimestamp uint64
}

func (w *GetTransactionReceiptOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Receipt)
	if err != nil {
		return
	}
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteUint64(buf, w.BlockTimestamp)
	return nil
}

func (w *GetTransactionReceiptOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetTransactionReceiptOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetTransactionReceiptOutputBuilder) Build() *GetTransactionReceiptOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionReceiptOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetLastCommittedBlockHeightInput

// reader

type GetLastCommittedBlockHeightInput struct {
	message membuffers.Message
}

var m_GetLastCommittedBlockHeightInput_Scheme = []membuffers.FieldType{}
var m_GetLastCommittedBlockHeightInput_Unions = [][]membuffers.FieldType{}

func GetLastCommittedBlockHeightInputReader(buf []byte) *GetLastCommittedBlockHeightInput {
	x := &GetLastCommittedBlockHeightInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetLastCommittedBlockHeightInput_Scheme, m_GetLastCommittedBlockHeightInput_Unions)
	return x
}

func (x *GetLastCommittedBlockHeightInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetLastCommittedBlockHeightInput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type GetLastCommittedBlockHeightInputBuilder struct {
	builder membuffers.Builder
}

func (w *GetLastCommittedBlockHeightInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	return nil
}

func (w *GetLastCommittedBlockHeightInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetLastCommittedBlockHeightInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetLastCommittedBlockHeightInputBuilder) Build() *GetLastCommittedBlockHeightInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetLastCommittedBlockHeightInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetLastCommittedBlockHeightOutput

// reader

type GetLastCommittedBlockHeightOutput struct {
	message membuffers.Message
}

var m_GetLastCommittedBlockHeightOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint64,}
var m_GetLastCommittedBlockHeightOutput_Unions = [][]membuffers.FieldType{}

func GetLastCommittedBlockHeightOutputReader(buf []byte) *GetLastCommittedBlockHeightOutput {
	x := &GetLastCommittedBlockHeightOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetLastCommittedBlockHeightOutput_Scheme, m_GetLastCommittedBlockHeightOutput_Unions)
	return x
}

func (x *GetLastCommittedBlockHeightOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetLastCommittedBlockHeightOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetLastCommittedBlockHeightOutput) LastCommittedBlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *GetLastCommittedBlockHeightOutput) RawLastCommittedBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetLastCommittedBlockHeightOutput) MutateLastCommittedBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *GetLastCommittedBlockHeightOutput) LastCommittedBlockTimestamp() uint64 {
	return x.message.GetUint64(1)
}

func (x *GetLastCommittedBlockHeightOutput) RawLastCommittedBlockTimestamp() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *GetLastCommittedBlockHeightOutput) MutateLastCommittedBlockTimestamp(v uint64) error {
	return x.message.SetUint64(1, v)
}

// builder

type GetLastCommittedBlockHeightOutputBuilder struct {
	builder membuffers.Builder
	LastCommittedBlockHeight uint64
	LastCommittedBlockTimestamp uint64
}

func (w *GetLastCommittedBlockHeightOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint64(buf, w.LastCommittedBlockHeight)
	w.builder.WriteUint64(buf, w.LastCommittedBlockTimestamp)
	return nil
}

func (w *GetLastCommittedBlockHeightOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetLastCommittedBlockHeightOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetLastCommittedBlockHeightOutputBuilder) Build() *GetLastCommittedBlockHeightOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetLastCommittedBlockHeightOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnBlockAvailabilityRequestInput

// reader

type OnBlockAvailabilityRequestInput struct {
	message membuffers.Message
}

var m_OnBlockAvailabilityRequestInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_OnBlockAvailabilityRequestInput_Unions = [][]membuffers.FieldType{}

func OnBlockAvailabilityRequestInputReader(buf []byte) *OnBlockAvailabilityRequestInput {
	x := &OnBlockAvailabilityRequestInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnBlockAvailabilityRequestInput_Scheme, m_OnBlockAvailabilityRequestInput_Unions)
	return x
}

func (x *OnBlockAvailabilityRequestInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnBlockAvailabilityRequestInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *OnBlockAvailabilityRequestInput) BlockAvailabilityRequest() *blocksync.BlockAvailabilityRequestMessage {
	b, s := x.message.GetMessage(0)
	return blocksync.BlockAvailabilityRequestMessageReader(b[:s])
}

func (x *OnBlockAvailabilityRequestInput) RawBlockAvailabilityRequest() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type OnBlockAvailabilityRequestInputBuilder struct {
	builder membuffers.Builder
	BlockAvailabilityRequest *blocksync.BlockAvailabilityRequestMessageBuilder
}

func (w *OnBlockAvailabilityRequestInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.BlockAvailabilityRequest)
	if err != nil {
		return
	}
	return nil
}

func (w *OnBlockAvailabilityRequestInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnBlockAvailabilityRequestInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnBlockAvailabilityRequestInputBuilder) Build() *OnBlockAvailabilityRequestInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnBlockAvailabilityRequestInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnBlockAvailabilityRequestOutput

// reader

type OnBlockAvailabilityRequestOutput struct {
	message membuffers.Message
}

var m_OnBlockAvailabilityRequestOutput_Scheme = []membuffers.FieldType{}
var m_OnBlockAvailabilityRequestOutput_Unions = [][]membuffers.FieldType{}

func OnBlockAvailabilityRequestOutputReader(buf []byte) *OnBlockAvailabilityRequestOutput {
	x := &OnBlockAvailabilityRequestOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnBlockAvailabilityRequestOutput_Scheme, m_OnBlockAvailabilityRequestOutput_Unions)
	return x
}

func (x *OnBlockAvailabilityRequestOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnBlockAvailabilityRequestOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type OnBlockAvailabilityRequestOutputBuilder struct {
	builder membuffers.Builder
}

func (w *OnBlockAvailabilityRequestOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	return nil
}

func (w *OnBlockAvailabilityRequestOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnBlockAvailabilityRequestOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnBlockAvailabilityRequestOutputBuilder) Build() *OnBlockAvailabilityRequestOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnBlockAvailabilityRequestOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnBlockAvailabilityResponseInput

// reader

type OnBlockAvailabilityResponseInput struct {
	message membuffers.Message
}

var m_OnBlockAvailabilityResponseInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_OnBlockAvailabilityResponseInput_Unions = [][]membuffers.FieldType{}

func OnBlockAvailabilityResponseInputReader(buf []byte) *OnBlockAvailabilityResponseInput {
	x := &OnBlockAvailabilityResponseInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnBlockAvailabilityResponseInput_Scheme, m_OnBlockAvailabilityResponseInput_Unions)
	return x
}

func (x *OnBlockAvailabilityResponseInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnBlockAvailabilityResponseInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *OnBlockAvailabilityResponseInput) BlockAvailabilityResponse() *blocksync.BlockAvailabilityRequestMessage {
	b, s := x.message.GetMessage(0)
	return blocksync.BlockAvailabilityRequestMessageReader(b[:s])
}

func (x *OnBlockAvailabilityResponseInput) RawBlockAvailabilityResponse() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type OnBlockAvailabilityResponseInputBuilder struct {
	builder membuffers.Builder
	BlockAvailabilityResponse *blocksync.BlockAvailabilityRequestMessageBuilder
}

func (w *OnBlockAvailabilityResponseInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.BlockAvailabilityResponse)
	if err != nil {
		return
	}
	return nil
}

func (w *OnBlockAvailabilityResponseInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnBlockAvailabilityResponseInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnBlockAvailabilityResponseInputBuilder) Build() *OnBlockAvailabilityResponseInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnBlockAvailabilityResponseInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnBlockAvailabilityResponseOutput

// reader

type OnBlockAvailabilityResponseOutput struct {
	message membuffers.Message
}

var m_OnBlockAvailabilityResponseOutput_Scheme = []membuffers.FieldType{}
var m_OnBlockAvailabilityResponseOutput_Unions = [][]membuffers.FieldType{}

func OnBlockAvailabilityResponseOutputReader(buf []byte) *OnBlockAvailabilityResponseOutput {
	x := &OnBlockAvailabilityResponseOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnBlockAvailabilityResponseOutput_Scheme, m_OnBlockAvailabilityResponseOutput_Unions)
	return x
}

func (x *OnBlockAvailabilityResponseOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnBlockAvailabilityResponseOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type OnBlockAvailabilityResponseOutputBuilder struct {
	builder membuffers.Builder
}

func (w *OnBlockAvailabilityResponseOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	return nil
}

func (w *OnBlockAvailabilityResponseOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnBlockAvailabilityResponseOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnBlockAvailabilityResponseOutputBuilder) Build() *OnBlockAvailabilityResponseOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnBlockAvailabilityResponseOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnBlockSyncRequestInput

// reader

type OnBlockSyncRequestInput struct {
	message membuffers.Message
}

var m_OnBlockSyncRequestInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_OnBlockSyncRequestInput_Unions = [][]membuffers.FieldType{}

func OnBlockSyncRequestInputReader(buf []byte) *OnBlockSyncRequestInput {
	x := &OnBlockSyncRequestInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnBlockSyncRequestInput_Scheme, m_OnBlockSyncRequestInput_Unions)
	return x
}

func (x *OnBlockSyncRequestInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnBlockSyncRequestInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *OnBlockSyncRequestInput) BlockSyncRequest() *blocksync.BlockSyncRequestMessage {
	b, s := x.message.GetMessage(0)
	return blocksync.BlockSyncRequestMessageReader(b[:s])
}

func (x *OnBlockSyncRequestInput) RawBlockSyncRequest() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type OnBlockSyncRequestInputBuilder struct {
	builder membuffers.Builder
	BlockSyncRequest *blocksync.BlockSyncRequestMessageBuilder
}

func (w *OnBlockSyncRequestInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.BlockSyncRequest)
	if err != nil {
		return
	}
	return nil
}

func (w *OnBlockSyncRequestInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnBlockSyncRequestInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnBlockSyncRequestInputBuilder) Build() *OnBlockSyncRequestInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnBlockSyncRequestInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnBlockSyncRequestOutput

// reader

type OnBlockSyncRequestOutput struct {
	message membuffers.Message
}

var m_OnBlockSyncRequestOutput_Scheme = []membuffers.FieldType{}
var m_OnBlockSyncRequestOutput_Unions = [][]membuffers.FieldType{}

func OnBlockSyncRequestOutputReader(buf []byte) *OnBlockSyncRequestOutput {
	x := &OnBlockSyncRequestOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnBlockSyncRequestOutput_Scheme, m_OnBlockSyncRequestOutput_Unions)
	return x
}

func (x *OnBlockSyncRequestOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnBlockSyncRequestOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type OnBlockSyncRequestOutputBuilder struct {
	builder membuffers.Builder
}

func (w *OnBlockSyncRequestOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	return nil
}

func (w *OnBlockSyncRequestOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnBlockSyncRequestOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnBlockSyncRequestOutputBuilder) Build() *OnBlockSyncRequestOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnBlockSyncRequestOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnBlockSyncResponseInput

// reader

type OnBlockSyncResponseInput struct {
	message membuffers.Message
}

var m_OnBlockSyncResponseInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var m_OnBlockSyncResponseInput_Unions = [][]membuffers.FieldType{}

func OnBlockSyncResponseInputReader(buf []byte) *OnBlockSyncResponseInput {
	x := &OnBlockSyncResponseInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnBlockSyncResponseInput_Scheme, m_OnBlockSyncResponseInput_Unions)
	return x
}

func (x *OnBlockSyncResponseInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnBlockSyncResponseInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *OnBlockSyncResponseInput) BlockSyncResponse() *blocksync.BlockSyncResponseMessage {
	b, s := x.message.GetMessage(0)
	return blocksync.BlockSyncResponseMessageReader(b[:s])
}

func (x *OnBlockSyncResponseInput) RawBlockSyncResponse() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type OnBlockSyncResponseInputBuilder struct {
	builder membuffers.Builder
	BlockSyncResponse *blocksync.BlockSyncResponseMessageBuilder
}

func (w *OnBlockSyncResponseInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.BlockSyncResponse)
	if err != nil {
		return
	}
	return nil
}

func (w *OnBlockSyncResponseInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnBlockSyncResponseInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnBlockSyncResponseInputBuilder) Build() *OnBlockSyncResponseInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnBlockSyncResponseInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message OnBlockSyncResponseOutput

// reader

type OnBlockSyncResponseOutput struct {
	message membuffers.Message
}

var m_OnBlockSyncResponseOutput_Scheme = []membuffers.FieldType{}
var m_OnBlockSyncResponseOutput_Unions = [][]membuffers.FieldType{}

func OnBlockSyncResponseOutputReader(buf []byte) *OnBlockSyncResponseOutput {
	x := &OnBlockSyncResponseOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_OnBlockSyncResponseOutput_Scheme, m_OnBlockSyncResponseOutput_Unions)
	return x
}

func (x *OnBlockSyncResponseOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *OnBlockSyncResponseOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type OnBlockSyncResponseOutputBuilder struct {
	builder membuffers.Builder
}

func (w *OnBlockSyncResponseOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	return nil
}

func (w *OnBlockSyncResponseOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *OnBlockSyncResponseOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *OnBlockSyncResponseOutputBuilder) Build() *OnBlockSyncResponseOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return OnBlockSyncResponseOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateBlockForCommitInput

// reader

type ValidateBlockForCommitInput struct {
	message membuffers.Message
}

var m_ValidateBlockForCommitInput_Scheme = []membuffers.FieldType{}
var m_ValidateBlockForCommitInput_Unions = [][]membuffers.FieldType{}

func ValidateBlockForCommitInputReader(buf []byte) *ValidateBlockForCommitInput {
	x := &ValidateBlockForCommitInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ValidateBlockForCommitInput_Scheme, m_ValidateBlockForCommitInput_Unions)
	return x
}

func (x *ValidateBlockForCommitInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ValidateBlockForCommitInput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type ValidateBlockForCommitInputBuilder struct {
	builder membuffers.Builder
}

func (w *ValidateBlockForCommitInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	return nil
}

func (w *ValidateBlockForCommitInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ValidateBlockForCommitInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ValidateBlockForCommitInputBuilder) Build() *ValidateBlockForCommitInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ValidateBlockForCommitInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateBlockForCommitOutput

// reader

type ValidateBlockForCommitOutput struct {
	message membuffers.Message
}

var m_ValidateBlockForCommitOutput_Scheme = []membuffers.FieldType{}
var m_ValidateBlockForCommitOutput_Unions = [][]membuffers.FieldType{}

func ValidateBlockForCommitOutputReader(buf []byte) *ValidateBlockForCommitOutput {
	x := &ValidateBlockForCommitOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ValidateBlockForCommitOutput_Scheme, m_ValidateBlockForCommitOutput_Unions)
	return x
}

func (x *ValidateBlockForCommitOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ValidateBlockForCommitOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type ValidateBlockForCommitOutputBuilder struct {
	builder membuffers.Builder
}

func (w *ValidateBlockForCommitOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	return nil
}

func (w *ValidateBlockForCommitOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ValidateBlockForCommitOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ValidateBlockForCommitOutputBuilder) Build() *ValidateBlockForCommitOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ValidateBlockForCommitOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

