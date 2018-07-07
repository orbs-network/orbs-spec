// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service StateStorage

type StateStorage interface {
	CommitStateDiff(input *CommitStateDiffInput) (*CommitStateDiffOutput, error)
	ReadKeys(input *ReadKeysInput) (*ReadKeysOutput, error)
	GetStateStorageBlockHeight(input *GetStateStorageBlockHeightInput) (*GetStateStorageBlockHeightOutput, error)
	GetStateHash(input *GetStateHashInput) (*GetStateHashOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message ReadKeysInput

// reader

type ReadKeysInput struct {
	message membuffers.Message
}

var m_ReadKeysInput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeMessage,membuffers.TypeBytesArray,}
var m_ReadKeysInput_Unions = [][]membuffers.FieldType{}

func ReadKeysInputReader(buf []byte) *ReadKeysInput {
	x := &ReadKeysInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ReadKeysInput_Scheme, m_ReadKeysInput_Unions)
	return x
}

func (x *ReadKeysInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ReadKeysInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ReadKeysInput) BlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *ReadKeysInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ReadKeysInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *ReadKeysInput) Contract() *protocol.ContractAddress {
	b, s := x.message.GetMessage(1)
	return protocol.ContractAddressReader(b[:s])
}

func (x *ReadKeysInput) RawContract() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ReadKeysInput) KeyIterator() *ReadKeysInputKeyIterator {
	return &ReadKeysInputKeyIterator{iterator: x.message.GetBytesArrayIterator(2)}
}

type ReadKeysInputKeyIterator struct {
	iterator *membuffers.Iterator
}

func (i *ReadKeysInputKeyIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ReadKeysInputKeyIterator) NextKey() []byte {
	return i.iterator.NextBytes()
}

func (x *ReadKeysInput) RawKeyArray() []byte {
	return x.message.RawBufferForField(2, 0)
}

// builder

type ReadKeysInputBuilder struct {
	builder membuffers.Builder
	BlockHeight uint64
	Contract *protocol.ContractAddressBuilder
	Key [][]byte
}

func (w *ReadKeysInputBuilder) Write(buf []byte) (err error) {
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
	err = w.builder.WriteMessage(buf, w.Contract)
	if err != nil {
		return
	}
	w.builder.WriteBytesArray(buf, w.Key)
	return nil
}

func (w *ReadKeysInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ReadKeysInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ReadKeysInputBuilder) Build() *ReadKeysInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ReadKeysInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ReadKeysOutput

// reader

type ReadKeysOutput struct {
	message membuffers.Message
}

var m_ReadKeysOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeBytesArray,}
var m_ReadKeysOutput_Unions = [][]membuffers.FieldType{}

func ReadKeysOutputReader(buf []byte) *ReadKeysOutput {
	x := &ReadKeysOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ReadKeysOutput_Scheme, m_ReadKeysOutput_Unions)
	return x
}

func (x *ReadKeysOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ReadKeysOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ReadKeysOutput) Status() protocol.RequestStatus {
	return protocol.RequestStatus(x.message.GetUint16(0))
}

func (x *ReadKeysOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ReadKeysOutput) MutateStatus(v protocol.RequestStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *ReadKeysOutput) BlobIterator() *ReadKeysOutputBlobIterator {
	return &ReadKeysOutputBlobIterator{iterator: x.message.GetBytesArrayIterator(1)}
}

type ReadKeysOutputBlobIterator struct {
	iterator *membuffers.Iterator
}

func (i *ReadKeysOutputBlobIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ReadKeysOutputBlobIterator) NextBlob() []byte {
	return i.iterator.NextBytes()
}

func (x *ReadKeysOutput) RawBlobArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type ReadKeysOutputBuilder struct {
	builder membuffers.Builder
	Status protocol.RequestStatus
	Blob [][]byte
}

func (w *ReadKeysOutputBuilder) Write(buf []byte) (err error) {
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
	w.builder.WriteBytesArray(buf, w.Blob)
	return nil
}

func (w *ReadKeysOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ReadKeysOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ReadKeysOutputBuilder) Build() *ReadKeysOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ReadKeysOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateStorageBlockHeightInput

// reader

type GetStateStorageBlockHeightInput struct {
	message membuffers.Message
}

var m_GetStateStorageBlockHeightInput_Scheme = []membuffers.FieldType{}
var m_GetStateStorageBlockHeightInput_Unions = [][]membuffers.FieldType{}

func GetStateStorageBlockHeightInputReader(buf []byte) *GetStateStorageBlockHeightInput {
	x := &GetStateStorageBlockHeightInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetStateStorageBlockHeightInput_Scheme, m_GetStateStorageBlockHeightInput_Unions)
	return x
}

func (x *GetStateStorageBlockHeightInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetStateStorageBlockHeightInput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type GetStateStorageBlockHeightInputBuilder struct {
	builder membuffers.Builder
}

func (w *GetStateStorageBlockHeightInputBuilder) Write(buf []byte) (err error) {
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

func (w *GetStateStorageBlockHeightInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetStateStorageBlockHeightInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetStateStorageBlockHeightInputBuilder) Build() *GetStateStorageBlockHeightInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetStateStorageBlockHeightInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateStorageBlockHeightOutput

// reader

type GetStateStorageBlockHeightOutput struct {
	message membuffers.Message
}

var m_GetStateStorageBlockHeightOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint64,}
var m_GetStateStorageBlockHeightOutput_Unions = [][]membuffers.FieldType{}

func GetStateStorageBlockHeightOutputReader(buf []byte) *GetStateStorageBlockHeightOutput {
	x := &GetStateStorageBlockHeightOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetStateStorageBlockHeightOutput_Scheme, m_GetStateStorageBlockHeightOutput_Unions)
	return x
}

func (x *GetStateStorageBlockHeightOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetStateStorageBlockHeightOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetStateStorageBlockHeightOutput) LastCommittedBlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *GetStateStorageBlockHeightOutput) RawLastCommittedBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetStateStorageBlockHeightOutput) MutateLastCommittedBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *GetStateStorageBlockHeightOutput) LastCommittedBlockTimestamp() uint64 {
	return x.message.GetUint64(1)
}

func (x *GetStateStorageBlockHeightOutput) RawLastCommittedBlockTimestamp() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *GetStateStorageBlockHeightOutput) MutateLastCommittedBlockTimestamp(v uint64) error {
	return x.message.SetUint64(1, v)
}

// builder

type GetStateStorageBlockHeightOutputBuilder struct {
	builder membuffers.Builder
	LastCommittedBlockHeight uint64
	LastCommittedBlockTimestamp uint64
}

func (w *GetStateStorageBlockHeightOutputBuilder) Write(buf []byte) (err error) {
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

func (w *GetStateStorageBlockHeightOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetStateStorageBlockHeightOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetStateStorageBlockHeightOutputBuilder) Build() *GetStateStorageBlockHeightOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetStateStorageBlockHeightOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message CommitStateDiffInput

// reader

type CommitStateDiffInput struct {
	message membuffers.Message
}

var m_CommitStateDiffInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessageArray,membuffers.TypeUint64,}
var m_CommitStateDiffInput_Unions = [][]membuffers.FieldType{}

func CommitStateDiffInputReader(buf []byte) *CommitStateDiffInput {
	x := &CommitStateDiffInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_CommitStateDiffInput_Scheme, m_CommitStateDiffInput_Unions)
	return x
}

func (x *CommitStateDiffInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *CommitStateDiffInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *CommitStateDiffInput) ResultsBlockHeader() *protocol.ResultsBlockHeader {
	b, s := x.message.GetMessage(0)
	return protocol.ResultsBlockHeaderReader(b[:s])
}

func (x *CommitStateDiffInput) RawResultsBlockHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *CommitStateDiffInput) ContractStateDiffIterator() *CommitStateDiffInputContractStateDiffIterator {
	return &CommitStateDiffInputContractStateDiffIterator{iterator: x.message.GetMessageArrayIterator(1)}
}

type CommitStateDiffInputContractStateDiffIterator struct {
	iterator *membuffers.Iterator
}

func (i *CommitStateDiffInputContractStateDiffIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *CommitStateDiffInputContractStateDiffIterator) NextContractStateDiff() *protocol.ContractStateDiff {
	b, s := i.iterator.NextMessage()
	return protocol.ContractStateDiffReader(b[:s])
}

func (x *CommitStateDiffInput) RawContractStateDiffArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *CommitStateDiffInput) LastCommittedBlockHeight() uint64 {
	return x.message.GetUint64(2)
}

func (x *CommitStateDiffInput) RawLastCommittedBlockHeight() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *CommitStateDiffInput) MutateLastCommittedBlockHeight(v uint64) error {
	return x.message.SetUint64(2, v)
}

// builder

type CommitStateDiffInputBuilder struct {
	builder membuffers.Builder
	ResultsBlockHeader *protocol.ResultsBlockHeaderBuilder
	ContractStateDiff []*protocol.ContractStateDiffBuilder
	LastCommittedBlockHeight uint64
}

func (w *CommitStateDiffInputBuilder) arrayOfContractStateDiff() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.ContractStateDiff))
	for i, v := range w.ContractStateDiff {
		res[i] = v
	}
	return res
}

func (w *CommitStateDiffInputBuilder) Write(buf []byte) (err error) {
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
	err = w.builder.WriteMessageArray(buf, w.arrayOfContractStateDiff())
	if err != nil {
		return
	}
	w.builder.WriteUint64(buf, w.LastCommittedBlockHeight)
	return nil
}

func (w *CommitStateDiffInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *CommitStateDiffInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *CommitStateDiffInputBuilder) Build() *CommitStateDiffInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return CommitStateDiffInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message CommitStateDiffOutput

// reader

type CommitStateDiffOutput struct {
	message membuffers.Message
}

var m_CommitStateDiffOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,membuffers.TypeUint64,}
var m_CommitStateDiffOutput_Unions = [][]membuffers.FieldType{}

func CommitStateDiffOutputReader(buf []byte) *CommitStateDiffOutput {
	x := &CommitStateDiffOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_CommitStateDiffOutput_Scheme, m_CommitStateDiffOutput_Unions)
	return x
}

func (x *CommitStateDiffOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *CommitStateDiffOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *CommitStateDiffOutput) NextDesiredBlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *CommitStateDiffOutput) RawNextDesiredBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *CommitStateDiffOutput) MutateNextDesiredBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

func (x *CommitStateDiffOutput) LastCommittedBlockHeight() uint64 {
	return x.message.GetUint64(1)
}

func (x *CommitStateDiffOutput) RawLastCommittedBlockHeight() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *CommitStateDiffOutput) MutateLastCommittedBlockHeight(v uint64) error {
	return x.message.SetUint64(1, v)
}

// builder

type CommitStateDiffOutputBuilder struct {
	builder membuffers.Builder
	NextDesiredBlockHeight uint64
	LastCommittedBlockHeight uint64
}

func (w *CommitStateDiffOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint64(buf, w.NextDesiredBlockHeight)
	w.builder.WriteUint64(buf, w.LastCommittedBlockHeight)
	return nil
}

func (w *CommitStateDiffOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *CommitStateDiffOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *CommitStateDiffOutputBuilder) Build() *CommitStateDiffOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return CommitStateDiffOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateHashInput

// reader

type GetStateHashInput struct {
	message membuffers.Message
}

var m_GetStateHashInput_Scheme = []membuffers.FieldType{membuffers.TypeUint64,}
var m_GetStateHashInput_Unions = [][]membuffers.FieldType{}

func GetStateHashInputReader(buf []byte) *GetStateHashInput {
	x := &GetStateHashInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetStateHashInput_Scheme, m_GetStateHashInput_Unions)
	return x
}

func (x *GetStateHashInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetStateHashInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetStateHashInput) BlockHeight() uint64 {
	return x.message.GetUint64(0)
}

func (x *GetStateHashInput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetStateHashInput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(0, v)
}

// builder

type GetStateHashInputBuilder struct {
	builder membuffers.Builder
	BlockHeight uint64
}

func (w *GetStateHashInputBuilder) Write(buf []byte) (err error) {
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

func (w *GetStateHashInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetStateHashInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetStateHashInputBuilder) Build() *GetStateHashInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetStateHashInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateHashOutput

// reader

type GetStateHashOutput struct {
	message membuffers.Message
}

var m_GetStateHashOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint16,membuffers.TypeBytes,}
var m_GetStateHashOutput_Unions = [][]membuffers.FieldType{}

func GetStateHashOutputReader(buf []byte) *GetStateHashOutput {
	x := &GetStateHashOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_GetStateHashOutput_Scheme, m_GetStateHashOutput_Unions)
	return x
}

func (x *GetStateHashOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetStateHashOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetStateHashOutput) Status() protocol.RequestStatus {
	return protocol.RequestStatus(x.message.GetUint16(0))
}

func (x *GetStateHashOutput) RawStatus() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetStateHashOutput) MutateStatus(v protocol.RequestStatus) error {
	return x.message.SetUint16(0, uint16(v))
}

func (x *GetStateHashOutput) StateRootHash() []byte {
	return x.message.GetBytes(1)
}

func (x *GetStateHashOutput) RawStateRootHash() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *GetStateHashOutput) MutateStateRootHash(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type GetStateHashOutputBuilder struct {
	builder membuffers.Builder
	Status protocol.RequestStatus
	StateRootHash []byte
}

func (w *GetStateHashOutputBuilder) Write(buf []byte) (err error) {
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
	w.builder.WriteBytes(buf, w.StateRootHash)
	return nil
}

func (w *GetStateHashOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetStateHashOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetStateHashOutputBuilder) Build() *GetStateHashOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetStateHashOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

