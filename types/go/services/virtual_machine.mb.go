// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service VirtualMachine

type VirtualMachine interface {
	ProcessTransactionSet(*ProcessTransactionSetInput) (*ProcessTransactionSetOutput, error)
	RunLocalMethod(*RunLocalMethodInput) (*RunLocalMethodOutput, error)
	TransactionSetPreOrder(*TransactionSetPreOrderInput) (*TransactionSetPreOrderOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessTransactionSetInput

// reader

type ProcessTransactionSetInput struct {
	message membuffers.Message
}

var m_ProcessTransactionSetInput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeMessageArray,}
var m_ProcessTransactionSetInput_Unions = [][]membuffers.FieldType{}

func ProcessTransactionSetInputReader(buf []byte) *ProcessTransactionSetInput {
	x := &ProcessTransactionSetInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ProcessTransactionSetInput_Scheme, m_ProcessTransactionSetInput_Unions)
	return x
}

func (x *ProcessTransactionSetInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ProcessTransactionSetInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ProcessTransactionSetInput) BlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *ProcessTransactionSetInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ProcessTransactionSetInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *ProcessTransactionSetInput) SignedTransactionIterator() *ProcessTransactionSetInputSignedTransactionIterator {
	return &ProcessTransactionSetInputSignedTransactionIterator{iterator: x.message.GetMessageArrayIterator(1)}
}

type ProcessTransactionSetInputSignedTransactionIterator struct {
	iterator *membuffers.Iterator
}

func (i *ProcessTransactionSetInputSignedTransactionIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ProcessTransactionSetInputSignedTransactionIterator) NextSignedTransaction() *protocol.SignedTransaction {
	b, s := i.iterator.NextMessage()
	return protocol.SignedTransactionReader(b[:s])
}

func (x *ProcessTransactionSetInput) RawSignedTransactionArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type ProcessTransactionSetInputBuilder struct {
	builder membuffers.Builder
	BlockHeight uint64
	SignedTransaction []*protocol.SignedTransactionBuilder
}

func (w *ProcessTransactionSetInputBuilder) arrayOfSignedTransaction() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.SignedTransaction))
	for i, v := range w.SignedTransaction {
		res[i] = v
	}
	return res
}

func (w *ProcessTransactionSetInputBuilder) Write(buf []byte) (err error) {
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
	err = w.builder.WriteMessageArray(buf, w.arrayOfSignedTransaction())
	if err != nil {
		return
	}
	return nil
}

func (w *ProcessTransactionSetInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ProcessTransactionSetInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ProcessTransactionSetInputBuilder) Build() *ProcessTransactionSetInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ProcessTransactionSetInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessTransactionSetOutput

// reader

type ProcessTransactionSetOutput struct {
	message membuffers.Message
}

var m_ProcessTransactionSetOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeMessageArray,membuffers.TypeMessageArray,}
var m_ProcessTransactionSetOutput_Unions = [][]membuffers.FieldType{}

func ProcessTransactionSetOutputReader(buf []byte) *ProcessTransactionSetOutput {
	x := &ProcessTransactionSetOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ProcessTransactionSetOutput_Scheme, m_ProcessTransactionSetOutput_Unions)
	return x
}

func (x *ProcessTransactionSetOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ProcessTransactionSetOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ProcessTransactionSetOutput) Status() protocol.RequestStatus {
	return protocol.RequestStatus(x.message.GetUint16(0))
}

func (x *ProcessTransactionSetOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ProcessTransactionSetOutput) MutateStatus(v protocol.RequestStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *ProcessTransactionSetOutput) TransactionReceiptIterator() *ProcessTransactionSetOutputTransactionReceiptIterator {
	return &ProcessTransactionSetOutputTransactionReceiptIterator{iterator: x.message.GetMessageArrayIterator(1)}
}

type ProcessTransactionSetOutputTransactionReceiptIterator struct {
	iterator *membuffers.Iterator
}

func (i *ProcessTransactionSetOutputTransactionReceiptIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ProcessTransactionSetOutputTransactionReceiptIterator) NextTransactionReceipt() *protocol.TransactionReceipt {
	b, s := i.iterator.NextMessage()
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *ProcessTransactionSetOutput) RawTransactionReceiptArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ProcessTransactionSetOutput) ContractStateDiffIterator() *ProcessTransactionSetOutputContractStateDiffIterator {
	return &ProcessTransactionSetOutputContractStateDiffIterator{iterator: x.message.GetMessageArrayIterator(2)}
}

type ProcessTransactionSetOutputContractStateDiffIterator struct {
	iterator *membuffers.Iterator
}

func (i *ProcessTransactionSetOutputContractStateDiffIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ProcessTransactionSetOutputContractStateDiffIterator) NextContractStateDiff() *protocol.ContractStateDiff {
	b, s := i.iterator.NextMessage()
	return protocol.ContractStateDiffReader(b[:s])
}

func (x *ProcessTransactionSetOutput) RawContractStateDiffArray() []byte {
	return x.message.RawBufferForField(2, 0)
}

// builder

type ProcessTransactionSetOutputBuilder struct {
	builder membuffers.Builder
	Status protocol.RequestStatus
	TransactionReceipt []*protocol.TransactionReceiptBuilder
	ContractStateDiff []*protocol.ContractStateDiffBuilder
}

func (w *ProcessTransactionSetOutputBuilder) arrayOfTransactionReceipt() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.TransactionReceipt))
	for i, v := range w.TransactionReceipt {
		res[i] = v
	}
	return res
}

func (w *ProcessTransactionSetOutputBuilder) arrayOfContractStateDiff() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.ContractStateDiff))
	for i, v := range w.ContractStateDiff {
		res[i] = v
	}
	return res
}

func (w *ProcessTransactionSetOutputBuilder) Write(buf []byte) (err error) {
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
	err = w.builder.WriteMessageArray(buf, w.arrayOfTransactionReceipt())
	if err != nil {
		return
	}
	err = w.builder.WriteMessageArray(buf, w.arrayOfContractStateDiff())
	if err != nil {
		return
	}
	return nil
}

func (w *ProcessTransactionSetOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ProcessTransactionSetOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ProcessTransactionSetOutputBuilder) Build() *ProcessTransactionSetOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ProcessTransactionSetOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message RunLocalMethodInput

// reader

type RunLocalMethodInput struct {
	message membuffers.Message
}

var m_RunLocalMethodInput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeMessage,}
var m_RunLocalMethodInput_Unions = [][]membuffers.FieldType{}

func RunLocalMethodInputReader(buf []byte) *RunLocalMethodInput {
	x := &RunLocalMethodInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_RunLocalMethodInput_Scheme, m_RunLocalMethodInput_Unions)
	return x
}

func (x *RunLocalMethodInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *RunLocalMethodInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *RunLocalMethodInput) BlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *RunLocalMethodInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *RunLocalMethodInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *RunLocalMethodInput) Transaction() *protocol.Transaction {
	b, s := x.message.GetMessage(1)
	return protocol.TransactionReader(b[:s])
}

func (x *RunLocalMethodInput) RawTransaction() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type RunLocalMethodInputBuilder struct {
	builder membuffers.Builder
	BlockHeight uint64
	Transaction *protocol.TransactionBuilder
}

func (w *RunLocalMethodInputBuilder) Write(buf []byte) (err error) {
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
	err = w.builder.WriteMessage(buf, w.Transaction)
	if err != nil {
		return
	}
	return nil
}

func (w *RunLocalMethodInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *RunLocalMethodInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *RunLocalMethodInputBuilder) Build() *RunLocalMethodInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RunLocalMethodInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message RunLocalMethodOutput

// reader

type RunLocalMethodOutput struct {
	message membuffers.Message
}

var m_RunLocalMethodOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeMessageArray,membuffers.TypeUint64,membuffers.TypeUint64,}
var m_RunLocalMethodOutput_Unions = [][]membuffers.FieldType{}

func RunLocalMethodOutputReader(buf []byte) *RunLocalMethodOutput {
	x := &RunLocalMethodOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_RunLocalMethodOutput_Scheme, m_RunLocalMethodOutput_Unions)
	return x
}

func (x *RunLocalMethodOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *RunLocalMethodOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *RunLocalMethodOutput) Status() protocol.CallMethodStatus {
	return protocol.CallMethodStatus(x.message.GetUint16(0))
}

func (x *RunLocalMethodOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *RunLocalMethodOutput) MutateStatus(v protocol.CallMethodStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *RunLocalMethodOutput) OutputArgumentIterator() *RunLocalMethodOutputOutputArgumentIterator {
	return &RunLocalMethodOutputOutputArgumentIterator{iterator: x.message.GetMessageArrayIterator(1)}
}

type RunLocalMethodOutputOutputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *RunLocalMethodOutputOutputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *RunLocalMethodOutputOutputArgumentIterator) NextOutputArgument() *protocol.MethodArgument {
	b, s := i.iterator.NextMessage()
	return protocol.MethodArgumentReader(b[:s])
}

func (x *RunLocalMethodOutput) RawOutputArgumentArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *RunLocalMethodOutput) ReferenceBlockHeight() uint64 {
	return x.message.GetUint64(2)
}

func (x *RunLocalMethodOutput) RawReferenceBlockHeight() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *RunLocalMethodOutput) MutateReferenceBlockHeight(v uint64) error {
	return x.message.SetUint64(2, v)
}

func (x *RunLocalMethodOutput) ReferenceBlockTimestamp() uint64 {
	return x.message.GetUint64(3)
}

func (x *RunLocalMethodOutput) RawReferenceBlockTimestamp() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *RunLocalMethodOutput) MutateReferenceBlockTimestamp(v uint64) error {
	return x.message.SetUint64(3, v)
}

// builder

type RunLocalMethodOutputBuilder struct {
	builder membuffers.Builder
	Status protocol.CallMethodStatus
	OutputArgument []*protocol.MethodArgumentBuilder
	ReferenceBlockHeight uint64
	ReferenceBlockTimestamp uint64
}

func (w *RunLocalMethodOutputBuilder) arrayOfOutputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.OutputArgument))
	for i, v := range w.OutputArgument {
		res[i] = v
	}
	return res
}

func (w *RunLocalMethodOutputBuilder) Write(buf []byte) (err error) {
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
	err = w.builder.WriteMessageArray(buf, w.arrayOfOutputArgument())
	if err != nil {
		return
	}
	w.builder.WriteUint64(buf, w.ReferenceBlockHeight)
	w.builder.WriteUint64(buf, w.ReferenceBlockTimestamp)
	return nil
}

func (w *RunLocalMethodOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *RunLocalMethodOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *RunLocalMethodOutputBuilder) Build() *RunLocalMethodOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return RunLocalMethodOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSetPreOrderInput

// reader

type TransactionSetPreOrderInput struct {
	message membuffers.Message
}

var m_TransactionSetPreOrderInput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeMessageArray,}
var m_TransactionSetPreOrderInput_Unions = [][]membuffers.FieldType{}

func TransactionSetPreOrderInputReader(buf []byte) *TransactionSetPreOrderInput {
	x := &TransactionSetPreOrderInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TransactionSetPreOrderInput_Scheme, m_TransactionSetPreOrderInput_Unions)
	return x
}

func (x *TransactionSetPreOrderInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionSetPreOrderInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionSetPreOrderInput) BlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *TransactionSetPreOrderInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionSetPreOrderInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *TransactionSetPreOrderInput) SignedTransactionIterator() *TransactionSetPreOrderInputSignedTransactionIterator {
	return &TransactionSetPreOrderInputSignedTransactionIterator{iterator: x.message.GetMessageArrayIterator(1)}
}

type TransactionSetPreOrderInputSignedTransactionIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionSetPreOrderInputSignedTransactionIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionSetPreOrderInputSignedTransactionIterator) NextSignedTransaction() *protocol.SignedTransaction {
	b, s := i.iterator.NextMessage()
	return protocol.SignedTransactionReader(b[:s])
}

func (x *TransactionSetPreOrderInput) RawSignedTransactionArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type TransactionSetPreOrderInputBuilder struct {
	builder membuffers.Builder
	BlockHeight uint64
	SignedTransaction []*protocol.SignedTransactionBuilder
}

func (w *TransactionSetPreOrderInputBuilder) arrayOfSignedTransaction() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.SignedTransaction))
	for i, v := range w.SignedTransaction {
		res[i] = v
	}
	return res
}

func (w *TransactionSetPreOrderInputBuilder) Write(buf []byte) (err error) {
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
	err = w.builder.WriteMessageArray(buf, w.arrayOfSignedTransaction())
	if err != nil {
		return
	}
	return nil
}

func (w *TransactionSetPreOrderInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionSetPreOrderInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *TransactionSetPreOrderInputBuilder) Build() *TransactionSetPreOrderInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionSetPreOrderInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSetPreOrderOutput

// reader

type TransactionSetPreOrderOutput struct {
	message membuffers.Message
}

var m_TransactionSetPreOrderOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeMessageArray,}
var m_TransactionSetPreOrderOutput_Unions = [][]membuffers.FieldType{}

func TransactionSetPreOrderOutputReader(buf []byte) *TransactionSetPreOrderOutput {
	x := &TransactionSetPreOrderOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TransactionSetPreOrderOutput_Scheme, m_TransactionSetPreOrderOutput_Unions)
	return x
}

func (x *TransactionSetPreOrderOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionSetPreOrderOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionSetPreOrderOutput) Status() protocol.RequestStatus {
	return protocol.RequestStatus(x.message.GetUint16(0))
}

func (x *TransactionSetPreOrderOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionSetPreOrderOutput) MutateStatus(v protocol.RequestStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *TransactionSetPreOrderOutput) PreOrderResultIterator() *TransactionSetPreOrderOutputPreOrderResultIterator {
	return &TransactionSetPreOrderOutputPreOrderResultIterator{iterator: x.message.GetMessageArrayIterator(1)}
}

type TransactionSetPreOrderOutputPreOrderResultIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionSetPreOrderOutputPreOrderResultIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionSetPreOrderOutputPreOrderResultIterator) NextPreOrderResult() *PreOrderResult {
	b, s := i.iterator.NextMessage()
	return PreOrderResultReader(b[:s])
}

func (x *TransactionSetPreOrderOutput) RawPreOrderResultArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type TransactionSetPreOrderOutputBuilder struct {
	builder membuffers.Builder
	Status protocol.RequestStatus
	PreOrderResult []*PreOrderResultBuilder
}

func (w *TransactionSetPreOrderOutputBuilder) arrayOfPreOrderResult() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.PreOrderResult))
	for i, v := range w.PreOrderResult {
		res[i] = v
	}
	return res
}

func (w *TransactionSetPreOrderOutputBuilder) Write(buf []byte) (err error) {
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
	err = w.builder.WriteMessageArray(buf, w.arrayOfPreOrderResult())
	if err != nil {
		return
	}
	return nil
}

func (w *TransactionSetPreOrderOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionSetPreOrderOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *TransactionSetPreOrderOutputBuilder) Build() *TransactionSetPreOrderOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionSetPreOrderOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message PreOrderResult

// reader

type PreOrderResult struct {
	message membuffers.Message
}

var m_PreOrderResult_Scheme = []membuffers.FieldType{membuffers.TypeUint16,}
var m_PreOrderResult_Unions = [][]membuffers.FieldType{}

func PreOrderResultReader(buf []byte) *PreOrderResult {
	x := &PreOrderResult{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_PreOrderResult_Scheme, m_PreOrderResult_Unions)
	return x
}

func (x *PreOrderResult) IsValid() bool {
	return x.message.IsValid()
}

func (x *PreOrderResult) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *PreOrderResult) Status() PreOrderStatus {
	return PreOrderStatus(x.message.GetUint16(0))
}

func (x *PreOrderResult) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *PreOrderResult) MutateStatus(v PreOrderStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

// builder

type PreOrderResultBuilder struct {
	builder membuffers.Builder
	Status PreOrderStatus
}

func (w *PreOrderResultBuilder) Write(buf []byte) (err error) {
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

func (w *PreOrderResultBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *PreOrderResultBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *PreOrderResultBuilder) Build() *PreOrderResult {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return PreOrderResultReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type PreOrderStatus uint16

const (
	PRE_ORDER_STATUS_RESERVED PreOrderStatus = 0
	PRE_ORDER_STATUS_VALID PreOrderStatus = 1
	PRE_ORDER_STATUS_INVALID_SIGNATURE PreOrderStatus = 2
	PRE_ORDER_STATUS_INVALID_ADDRESS_SCHEME PreOrderStatus = 3
	PRE_ORDER_STATUS_NOT_APPROVED PreOrderStatus = 4
)

