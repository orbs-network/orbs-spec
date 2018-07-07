// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusContext

type ConsensusContext interface {
	RequestNewTransactionsBlock(input *RequestNewTransactionsBlockInput) (*RequestNewTransactionsBlockOutput, error)
	RequestNewResultsBlock(input *RequestNewResultsBlockInput) (*RequestNewResultsBlockOutput, error)
	ValidateTransactionsBlock(input *ValidateTransactionsBlockInput) (*ValidateTransactionsBlockOutput, error)
	ValidateResultsBlock(input *ValidateResultsBlockInput) (*ValidateResultsBlockOutput, error)
	RequestOrderingCommittee(input *RequestCommitteeInput) (*RequestCommitteeOutput, error)
	RequestValidationCommittee(input *RequestCommitteeInput) (*RequestCommitteeOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewTransactionsBlockInput

// reader

type RequestNewTransactionsBlockInput struct {
	message membuffers.Message
}

var m_RequestNewTransactionsBlockInput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint32,membuffers.TypeUint32,membuffers.TypeBytes,}
var m_RequestNewTransactionsBlockInput_Unions = [][]membuffers.FieldType{}

func RequestNewTransactionsBlockInputReader(buf []byte) *RequestNewTransactionsBlockInput {
	x := &RequestNewTransactionsBlockInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_RequestNewTransactionsBlockInput_Scheme, m_RequestNewTransactionsBlockInput_Unions)
	return x
}

func (x *RequestNewTransactionsBlockInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *RequestNewTransactionsBlockInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *RequestNewTransactionsBlockInput) BlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *RequestNewTransactionsBlockInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *RequestNewTransactionsBlockInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *RequestNewTransactionsBlockInput) MaxBlockSizeKb() uint32 {
	return x.message.GetUint32(1)
}

func (x *RequestNewTransactionsBlockInput) RawMaxBlockSizeKb() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *RequestNewTransactionsBlockInput) MutateMaxBlockSizeKb(v uint32) error {
	return x.message.SetUint32(1, v)
}

func (x *RequestNewTransactionsBlockInput) MaxNumberOfTransactions() uint32 {
	return x.message.GetUint32(2)
}

func (x *RequestNewTransactionsBlockInput) RawMaxNumberOfTransactions() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *RequestNewTransactionsBlockInput) MutateMaxNumberOfTransactions(v uint32) error {
	return x.message.SetUint32(2, v)
}

func (x *RequestNewTransactionsBlockInput) PrevBlockHash() []byte {
	return x.message.GetBytes(3)
}

func (x *RequestNewTransactionsBlockInput) RawPrevBlockHash() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *RequestNewTransactionsBlockInput) MutatePrevBlockHash(v []byte) error {
	return x.message.SetBytes(3, v)
}

// builder

type RequestNewTransactionsBlockInputBuilder struct {
	builder membuffers.Builder
	BlockHeight uint64
	MaxBlockSizeKb uint32
	MaxNumberOfTransactions uint32
	PrevBlockHash []byte
}

func (w *RequestNewTransactionsBlockInputBuilder) Write(buf []byte) (err error) {
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
	w.builder.WriteUint32(buf, w.MaxBlockSizeKb)
	w.builder.WriteUint32(buf, w.MaxNumberOfTransactions)
	w.builder.WriteBytes(buf, w.PrevBlockHash)
	return nil
}

func (w *RequestNewTransactionsBlockInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *RequestNewTransactionsBlockInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *RequestNewTransactionsBlockInputBuilder) Build() *RequestNewTransactionsBlockInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RequestNewTransactionsBlockInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewTransactionsBlockOutput

// reader

type RequestNewTransactionsBlockOutput struct {
	message membuffers.Message
}

var m_RequestNewTransactionsBlockOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeMessage,}
var m_RequestNewTransactionsBlockOutput_Unions = [][]membuffers.FieldType{}

func RequestNewTransactionsBlockOutputReader(buf []byte) *RequestNewTransactionsBlockOutput {
	x := &RequestNewTransactionsBlockOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_RequestNewTransactionsBlockOutput_Scheme, m_RequestNewTransactionsBlockOutput_Unions)
	return x
}

func (x *RequestNewTransactionsBlockOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *RequestNewTransactionsBlockOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *RequestNewTransactionsBlockOutput) Status() protocol.RequestStatus {
	return protocol.RequestStatus(x.message.GetUint16(0))
}

func (x *RequestNewTransactionsBlockOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *RequestNewTransactionsBlockOutput) MutateStatus(v protocol.RequestStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *RequestNewTransactionsBlockOutput) TransactionsBlock() *protocol.TransactionsBlock {
	b, s := x.message.GetMessage(1)
	return protocol.TransactionsBlockReader(b[:s])
}

func (x *RequestNewTransactionsBlockOutput) RawTransactionsBlock() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type RequestNewTransactionsBlockOutputBuilder struct {
	builder membuffers.Builder
	Status protocol.RequestStatus
	TransactionsBlock *protocol.TransactionsBlockBuilder
}

func (w *RequestNewTransactionsBlockOutputBuilder) Write(buf []byte) (err error) {
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
	err = w.builder.WriteMessage(buf, w.TransactionsBlock)
	if err != nil {
		return
	}
	return nil
}

func (w *RequestNewTransactionsBlockOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *RequestNewTransactionsBlockOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *RequestNewTransactionsBlockOutputBuilder) Build() *RequestNewTransactionsBlockOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RequestNewTransactionsBlockOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewResultsBlockInput

// reader

type RequestNewResultsBlockInput struct {
	message membuffers.Message
}

var m_RequestNewResultsBlockInput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeBytes,}
var m_RequestNewResultsBlockInput_Unions = [][]membuffers.FieldType{}

func RequestNewResultsBlockInputReader(buf []byte) *RequestNewResultsBlockInput {
	x := &RequestNewResultsBlockInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_RequestNewResultsBlockInput_Scheme, m_RequestNewResultsBlockInput_Unions)
	return x
}

func (x *RequestNewResultsBlockInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *RequestNewResultsBlockInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *RequestNewResultsBlockInput) BlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *RequestNewResultsBlockInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *RequestNewResultsBlockInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *RequestNewResultsBlockInput) PrevBlockHash() []byte {
	return x.message.GetBytes(1)
}

func (x *RequestNewResultsBlockInput) RawPrevBlockHash() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *RequestNewResultsBlockInput) MutatePrevBlockHash(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type RequestNewResultsBlockInputBuilder struct {
	builder membuffers.Builder
	BlockHeight uint64
	PrevBlockHash []byte
}

func (w *RequestNewResultsBlockInputBuilder) Write(buf []byte) (err error) {
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
	w.builder.WriteBytes(buf, w.PrevBlockHash)
	return nil
}

func (w *RequestNewResultsBlockInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *RequestNewResultsBlockInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *RequestNewResultsBlockInputBuilder) Build() *RequestNewResultsBlockInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RequestNewResultsBlockInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewResultsBlockOutput

// reader

type RequestNewResultsBlockOutput struct {
	message membuffers.Message
}

var m_RequestNewResultsBlockOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeMessage,}
var m_RequestNewResultsBlockOutput_Unions = [][]membuffers.FieldType{}

func RequestNewResultsBlockOutputReader(buf []byte) *RequestNewResultsBlockOutput {
	x := &RequestNewResultsBlockOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_RequestNewResultsBlockOutput_Scheme, m_RequestNewResultsBlockOutput_Unions)
	return x
}

func (x *RequestNewResultsBlockOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *RequestNewResultsBlockOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *RequestNewResultsBlockOutput) Status() protocol.RequestStatus {
	return protocol.RequestStatus(x.message.GetUint16(0))
}

func (x *RequestNewResultsBlockOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *RequestNewResultsBlockOutput) MutateStatus(v protocol.RequestStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *RequestNewResultsBlockOutput) ResultsBlock() *protocol.ResultsBlock {
	b, s := x.message.GetMessage(1)
	return protocol.ResultsBlockReader(b[:s])
}

func (x *RequestNewResultsBlockOutput) RawResultsBlock() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type RequestNewResultsBlockOutputBuilder struct {
	builder membuffers.Builder
	Status protocol.RequestStatus
	ResultsBlock *protocol.ResultsBlockBuilder
}

func (w *RequestNewResultsBlockOutputBuilder) Write(buf []byte) (err error) {
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
	err = w.builder.WriteMessage(buf, w.ResultsBlock)
	if err != nil {
		return
	}
	return nil
}

func (w *RequestNewResultsBlockOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *RequestNewResultsBlockOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *RequestNewResultsBlockOutputBuilder) Build() *RequestNewResultsBlockOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RequestNewResultsBlockOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsBlockInput

// reader

type ValidateTransactionsBlockInput struct {
	message membuffers.Message
}

var m_ValidateTransactionsBlockInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var m_ValidateTransactionsBlockInput_Unions = [][]membuffers.FieldType{}

func ValidateTransactionsBlockInputReader(buf []byte) *ValidateTransactionsBlockInput {
	x := &ValidateTransactionsBlockInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ValidateTransactionsBlockInput_Scheme, m_ValidateTransactionsBlockInput_Unions)
	return x
}

func (x *ValidateTransactionsBlockInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ValidateTransactionsBlockInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ValidateTransactionsBlockInput) TransactionsBlock() *protocol.TransactionsBlock {
	b, s := x.message.GetMessage(0)
	return protocol.TransactionsBlockReader(b[:s])
}

func (x *ValidateTransactionsBlockInput) RawTransactionsBlock() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ValidateTransactionsBlockInput) PrevBlockHash() []byte {
	return x.message.GetBytes(1)
}

func (x *ValidateTransactionsBlockInput) RawPrevBlockHash() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ValidateTransactionsBlockInput) MutatePrevBlockHash(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type ValidateTransactionsBlockInputBuilder struct {
	builder membuffers.Builder
	TransactionsBlock *protocol.TransactionsBlockBuilder
	PrevBlockHash []byte
}

func (w *ValidateTransactionsBlockInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.TransactionsBlock)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.PrevBlockHash)
	return nil
}

func (w *ValidateTransactionsBlockInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ValidateTransactionsBlockInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ValidateTransactionsBlockInputBuilder) Build() *ValidateTransactionsBlockInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ValidateTransactionsBlockInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsBlockOutput

// reader

type ValidateTransactionsBlockOutput struct {
	message membuffers.Message
}

var m_ValidateTransactionsBlockOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,}
var m_ValidateTransactionsBlockOutput_Unions = [][]membuffers.FieldType{}

func ValidateTransactionsBlockOutputReader(buf []byte) *ValidateTransactionsBlockOutput {
	x := &ValidateTransactionsBlockOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ValidateTransactionsBlockOutput_Scheme, m_ValidateTransactionsBlockOutput_Unions)
	return x
}

func (x *ValidateTransactionsBlockOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ValidateTransactionsBlockOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ValidateTransactionsBlockOutput) Status() protocol.RequestStatus {
	return protocol.RequestStatus(x.message.GetUint16(0))
}

func (x *ValidateTransactionsBlockOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ValidateTransactionsBlockOutput) MutateStatus(v protocol.RequestStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

// builder

type ValidateTransactionsBlockOutputBuilder struct {
	builder membuffers.Builder
	Status protocol.RequestStatus
}

func (w *ValidateTransactionsBlockOutputBuilder) Write(buf []byte) (err error) {
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

func (w *ValidateTransactionsBlockOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ValidateTransactionsBlockOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ValidateTransactionsBlockOutputBuilder) Build() *ValidateTransactionsBlockOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ValidateTransactionsBlockOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateResultsBlockInput

// reader

type ValidateResultsBlockInput struct {
	message membuffers.Message
}

var m_ValidateResultsBlockInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var m_ValidateResultsBlockInput_Unions = [][]membuffers.FieldType{}

func ValidateResultsBlockInputReader(buf []byte) *ValidateResultsBlockInput {
	x := &ValidateResultsBlockInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ValidateResultsBlockInput_Scheme, m_ValidateResultsBlockInput_Unions)
	return x
}

func (x *ValidateResultsBlockInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ValidateResultsBlockInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ValidateResultsBlockInput) ResultsBlock() *protocol.ResultsBlock {
	b, s := x.message.GetMessage(0)
	return protocol.ResultsBlockReader(b[:s])
}

func (x *ValidateResultsBlockInput) RawResultsBlock() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ValidateResultsBlockInput) PrevBlockHash() []byte {
	return x.message.GetBytes(1)
}

func (x *ValidateResultsBlockInput) RawPrevBlockHash() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ValidateResultsBlockInput) MutatePrevBlockHash(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type ValidateResultsBlockInputBuilder struct {
	builder membuffers.Builder
	ResultsBlock *protocol.ResultsBlockBuilder
	PrevBlockHash []byte
}

func (w *ValidateResultsBlockInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.ResultsBlock)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.PrevBlockHash)
	return nil
}

func (w *ValidateResultsBlockInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ValidateResultsBlockInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ValidateResultsBlockInputBuilder) Build() *ValidateResultsBlockInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ValidateResultsBlockInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateResultsBlockOutput

// reader

type ValidateResultsBlockOutput struct {
	message membuffers.Message
}

var m_ValidateResultsBlockOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,}
var m_ValidateResultsBlockOutput_Unions = [][]membuffers.FieldType{}

func ValidateResultsBlockOutputReader(buf []byte) *ValidateResultsBlockOutput {
	x := &ValidateResultsBlockOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ValidateResultsBlockOutput_Scheme, m_ValidateResultsBlockOutput_Unions)
	return x
}

func (x *ValidateResultsBlockOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ValidateResultsBlockOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ValidateResultsBlockOutput) Status() protocol.RequestStatus {
	return protocol.RequestStatus(x.message.GetUint16(0))
}

func (x *ValidateResultsBlockOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ValidateResultsBlockOutput) MutateStatus(v protocol.RequestStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

// builder

type ValidateResultsBlockOutputBuilder struct {
	builder membuffers.Builder
	Status protocol.RequestStatus
}

func (w *ValidateResultsBlockOutputBuilder) Write(buf []byte) (err error) {
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

func (w *ValidateResultsBlockOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ValidateResultsBlockOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ValidateResultsBlockOutputBuilder) Build() *ValidateResultsBlockOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ValidateResultsBlockOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message RequestCommitteeInput

// reader

type RequestCommitteeInput struct {
	message membuffers.Message
}

var m_RequestCommitteeInput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint64,membuffers.TypeUint32,}
var m_RequestCommitteeInput_Unions = [][]membuffers.FieldType{}

func RequestCommitteeInputReader(buf []byte) *RequestCommitteeInput {
	x := &RequestCommitteeInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_RequestCommitteeInput_Scheme, m_RequestCommitteeInput_Unions)
	return x
}

func (x *RequestCommitteeInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *RequestCommitteeInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *RequestCommitteeInput) BlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *RequestCommitteeInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *RequestCommitteeInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *RequestCommitteeInput) RandomSeed() uint64 {
	return x.message.GetUint64(1)
}

func (x *RequestCommitteeInput) RawRandomSeed() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *RequestCommitteeInput) MutateRandomSeed(v uint64) error {
	return x.message.SetUint64(1, v)
}

func (x *RequestCommitteeInput) MaxCommitteeSize() uint32 {
	return x.message.GetUint32(2)
}

func (x *RequestCommitteeInput) RawMaxCommitteeSize() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *RequestCommitteeInput) MutateMaxCommitteeSize(v uint32) error {
	return x.message.SetUint32(2, v)
}

// builder

type RequestCommitteeInputBuilder struct {
	builder membuffers.Builder
	BlockHeight uint64
	RandomSeed uint64
	MaxCommitteeSize uint32
}

func (w *RequestCommitteeInputBuilder) Write(buf []byte) (err error) {
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
	w.builder.WriteUint64(buf, w.RandomSeed)
	w.builder.WriteUint32(buf, w.MaxCommitteeSize)
	return nil
}

func (w *RequestCommitteeInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *RequestCommitteeInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *RequestCommitteeInputBuilder) Build() *RequestCommitteeInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RequestCommitteeInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message RequestCommitteeOutput

// reader

type RequestCommitteeOutput struct {
	message membuffers.Message
}

var m_RequestCommitteeOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeMessageArray,}
var m_RequestCommitteeOutput_Unions = [][]membuffers.FieldType{}

func RequestCommitteeOutputReader(buf []byte) *RequestCommitteeOutput {
	x := &RequestCommitteeOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_RequestCommitteeOutput_Scheme, m_RequestCommitteeOutput_Unions)
	return x
}

func (x *RequestCommitteeOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *RequestCommitteeOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *RequestCommitteeOutput) Status() protocol.RequestStatus {
	return protocol.RequestStatus(x.message.GetUint16(0))
}

func (x *RequestCommitteeOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *RequestCommitteeOutput) MutateStatus(v protocol.RequestStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *RequestCommitteeOutput) NodeDataIterator() *RequestCommitteeOutputNodeDataIterator {
	return &RequestCommitteeOutputNodeDataIterator{iterator: x.message.GetMessageArrayIterator(1)}
}

type RequestCommitteeOutputNodeDataIterator struct {
	iterator *membuffers.Iterator
}

func (i *RequestCommitteeOutputNodeDataIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *RequestCommitteeOutputNodeDataIterator) NextNodeData() *NodeData {
	b, s := i.iterator.NextMessage()
	return NodeDataReader(b[:s])
}

func (x *RequestCommitteeOutput) RawNodeDataArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type RequestCommitteeOutputBuilder struct {
	builder membuffers.Builder
	Status protocol.RequestStatus
	NodeData []*NodeDataBuilder
}

func (w *RequestCommitteeOutputBuilder) arrayOfNodeData() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.NodeData))
	for i, v := range w.NodeData {
		res[i] = v
	}
	return res
}

func (w *RequestCommitteeOutputBuilder) Write(buf []byte) (err error) {
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
	err = w.builder.WriteMessageArray(buf, w.arrayOfNodeData())
	if err != nil {
		return
	}
	return nil
}

func (w *RequestCommitteeOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *RequestCommitteeOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *RequestCommitteeOutputBuilder) Build() *RequestCommitteeOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RequestCommitteeOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message NodeData

// reader

type NodeData struct {
	message membuffers.Message
}

var m_NodeData_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,}
var m_NodeData_Unions = [][]membuffers.FieldType{}

func NodeDataReader(buf []byte) *NodeData {
	x := &NodeData{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_NodeData_Scheme, m_NodeData_Unions)
	return x
}

func (x *NodeData) IsValid() bool {
	return x.message.IsValid()
}

func (x *NodeData) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *NodeData) NodePkey() []byte {
	return x.message.GetBytes(0)
}

func (x *NodeData) RawNodePkey() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *NodeData) MutateNodePkey(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *NodeData) NodeRandomSeedPkey() []byte {
	return x.message.GetBytes(1)
}

func (x *NodeData) RawNodeRandomSeedPkey() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *NodeData) MutateNodeRandomSeedPkey(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type NodeDataBuilder struct {
	builder membuffers.Builder
	NodePkey []byte
	NodeRandomSeedPkey []byte
}

func (w *NodeDataBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.NodePkey)
	w.builder.WriteBytes(buf, w.NodeRandomSeedPkey)
	return nil
}

func (w *NodeDataBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *NodeDataBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *NodeDataBuilder) Build() *NodeData {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return NodeDataReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

