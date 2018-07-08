// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message BlockPair

// reader

type BlockPair struct {
	message membuffers.Message
}

var _BlockPair_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,}
var _BlockPair_Unions = [][]membuffers.FieldType{}

func BlockPairReader(buf []byte) *BlockPair {
	x := &BlockPair{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _BlockPair_Scheme, _BlockPair_Unions)
	return x
}

func (x *BlockPair) IsValid() bool {
	return x.message.IsValid()
}

func (x *BlockPair) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *BlockPair) TransactionsBlock() *TransactionsBlock {
	b, s := x.message.GetMessage(0)
	return TransactionsBlockReader(b[:s])
}

func (x *BlockPair) RawTransactionsBlock() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *BlockPair) ResultsBlock() *ResultsBlock {
	b, s := x.message.GetMessage(1)
	return ResultsBlockReader(b[:s])
}

func (x *BlockPair) RawResultsBlock() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type BlockPairBuilder struct {
	builder membuffers.Builder
	TransactionsBlock *TransactionsBlockBuilder
	ResultsBlock *ResultsBlockBuilder
}

func (w *BlockPairBuilder) Write(buf []byte) (err error) {
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
	err = w.builder.WriteMessage(buf, w.ResultsBlock)
	if err != nil {
		return
	}
	return nil
}

func (w *BlockPairBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *BlockPairBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *BlockPairBuilder) Build() *BlockPair {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return BlockPairReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionsBlock

// reader

type TransactionsBlock struct {
	message membuffers.Message
}

var _TransactionsBlock_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}
var _TransactionsBlock_Unions = [][]membuffers.FieldType{}

func TransactionsBlockReader(buf []byte) *TransactionsBlock {
	x := &TransactionsBlock{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _TransactionsBlock_Scheme, _TransactionsBlock_Unions)
	return x
}

func (x *TransactionsBlock) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionsBlock) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionsBlock) Header() *TransactionsBlockHeader {
	b, s := x.message.GetMessage(0)
	return TransactionsBlockHeaderReader(b[:s])
}

func (x *TransactionsBlock) RawHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionsBlock) MetaData() *TransactionsBlockMetaData {
	b, s := x.message.GetMessage(1)
	return TransactionsBlockMetaDataReader(b[:s])
}

func (x *TransactionsBlock) RawMetaData() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *TransactionsBlock) Body() *TransactionsBlockBody {
	b, s := x.message.GetMessage(2)
	return TransactionsBlockBodyReader(b[:s])
}

func (x *TransactionsBlock) RawBody() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *TransactionsBlock) BlockProof() *TransactionsBlockProof {
	b, s := x.message.GetMessage(3)
	return TransactionsBlockProofReader(b[:s])
}

func (x *TransactionsBlock) RawBlockProof() []byte {
	return x.message.RawBufferForField(3, 0)
}

// builder

type TransactionsBlockBuilder struct {
	builder membuffers.Builder
	Header *TransactionsBlockHeaderBuilder
	MetaData *TransactionsBlockMetaDataBuilder
	Body *TransactionsBlockBodyBuilder
	BlockProof *TransactionsBlockProofBuilder
}

func (w *TransactionsBlockBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Header)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.MetaData)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.Body)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.BlockProof)
	if err != nil {
		return
	}
	return nil
}

func (w *TransactionsBlockBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionsBlockBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *TransactionsBlockBuilder) Build() *TransactionsBlock {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionsBlockReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ResultsBlock

// reader

type ResultsBlock struct {
	message membuffers.Message
}

var _ResultsBlock_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeMessage,}
var _ResultsBlock_Unions = [][]membuffers.FieldType{}

func ResultsBlockReader(buf []byte) *ResultsBlock {
	x := &ResultsBlock{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _ResultsBlock_Scheme, _ResultsBlock_Unions)
	return x
}

func (x *ResultsBlock) IsValid() bool {
	return x.message.IsValid()
}

func (x *ResultsBlock) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ResultsBlock) Header() *ResultsBlockHeader {
	b, s := x.message.GetMessage(0)
	return ResultsBlockHeaderReader(b[:s])
}

func (x *ResultsBlock) RawHeader() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ResultsBlock) Body() *ResultsBlockBody {
	b, s := x.message.GetMessage(1)
	return ResultsBlockBodyReader(b[:s])
}

func (x *ResultsBlock) RawBody() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ResultsBlock) BlockProof() *ResultsBlockProof {
	b, s := x.message.GetMessage(2)
	return ResultsBlockProofReader(b[:s])
}

func (x *ResultsBlock) RawBlockProof() []byte {
	return x.message.RawBufferForField(2, 0)
}

// builder

type ResultsBlockBuilder struct {
	builder membuffers.Builder
	Header *ResultsBlockHeaderBuilder
	Body *ResultsBlockBodyBuilder
	BlockProof *ResultsBlockProofBuilder
}

func (w *ResultsBlockBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Header)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.Body)
	if err != nil {
		return
	}
	err = w.builder.WriteMessage(buf, w.BlockProof)
	if err != nil {
		return
	}
	return nil
}

func (w *ResultsBlockBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ResultsBlockBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ResultsBlockBuilder) Build() *ResultsBlock {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ResultsBlockReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionsBlockHeader

// reader

type TransactionsBlockHeader struct {
	message membuffers.Message
}

var _TransactionsBlockHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeBytes,membuffers.TypeUint64,membuffers.TypeBytes,membuffers.TypeBytes,}
var _TransactionsBlockHeader_Unions = [][]membuffers.FieldType{}

func TransactionsBlockHeaderReader(buf []byte) *TransactionsBlockHeader {
	x := &TransactionsBlockHeader{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _TransactionsBlockHeader_Scheme, _TransactionsBlockHeader_Unions)
	return x
}

func (x *TransactionsBlockHeader) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionsBlockHeader) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionsBlockHeader) Version() uint32 {
	return x.message.GetUint32(0)
}

func (x *TransactionsBlockHeader) RawVersion() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionsBlockHeader) MutateVersion(v uint32) error {
	return x.message.SetUint32(0, v)
}

func (x *TransactionsBlockHeader) VirtualChain() uint32 {
	return x.message.GetUint32(1)
}

func (x *TransactionsBlockHeader) RawVirtualChain() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *TransactionsBlockHeader) MutateVirtualChain(v uint32) error {
	return x.message.SetUint32(1, v)
}

func (x *TransactionsBlockHeader) BlockHeight() uint64 {
	return x.message.GetUint64(2)
}

func (x *TransactionsBlockHeader) RawBlockHeight() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *TransactionsBlockHeader) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(2, v)
}

func (x *TransactionsBlockHeader) PrevBlockHashPtr() primitives.Sha256 {
	return x.message.GetBytes(3)
}

func (x *TransactionsBlockHeader) RawPrevBlockHashPtr() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *TransactionsBlockHeader) MutatePrevBlockHashPtr(v primitives.Sha256) error {
	return x.message.SetBytes(3, v)
}

func (x *TransactionsBlockHeader) Timestamp() uint64 {
	return x.message.GetUint64(4)
}

func (x *TransactionsBlockHeader) RawTimestamp() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *TransactionsBlockHeader) MutateTimestamp(v uint64) error {
	return x.message.SetUint64(4, v)
}

func (x *TransactionsBlockHeader) TransactionsRootHash() primitives.Sha256 {
	return x.message.GetBytes(5)
}

func (x *TransactionsBlockHeader) RawTransactionsRootHash() []byte {
	return x.message.RawBufferForField(5, 0)
}

func (x *TransactionsBlockHeader) MutateTransactionsRootHash(v primitives.Sha256) error {
	return x.message.SetBytes(5, v)
}

func (x *TransactionsBlockHeader) MetaDataHash() primitives.Sha256 {
	return x.message.GetBytes(6)
}

func (x *TransactionsBlockHeader) RawMetaDataHash() []byte {
	return x.message.RawBufferForField(6, 0)
}

func (x *TransactionsBlockHeader) MutateMetaDataHash(v primitives.Sha256) error {
	return x.message.SetBytes(6, v)
}

// builder

type TransactionsBlockHeaderBuilder struct {
	builder membuffers.Builder
	Version uint32
	VirtualChain uint32
	BlockHeight uint64
	PrevBlockHashPtr primitives.Sha256
	Timestamp uint64
	TransactionsRootHash primitives.Sha256
	MetaDataHash primitives.Sha256
}

func (w *TransactionsBlockHeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.Version)
	w.builder.WriteUint32(buf, w.VirtualChain)
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteBytes(buf, w.PrevBlockHashPtr)
	w.builder.WriteUint64(buf, w.Timestamp)
	w.builder.WriteBytes(buf, w.TransactionsRootHash)
	w.builder.WriteBytes(buf, w.MetaDataHash)
	return nil
}

func (w *TransactionsBlockHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionsBlockHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *TransactionsBlockHeaderBuilder) Build() *TransactionsBlockHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionsBlockHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionsBlockMetaData

// reader

type TransactionsBlockMetaData struct {
	message membuffers.Message
}

var _TransactionsBlockMetaData_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint32Array,}
var _TransactionsBlockMetaData_Unions = [][]membuffers.FieldType{}

func TransactionsBlockMetaDataReader(buf []byte) *TransactionsBlockMetaData {
	x := &TransactionsBlockMetaData{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _TransactionsBlockMetaData_Scheme, _TransactionsBlockMetaData_Unions)
	return x
}

func (x *TransactionsBlockMetaData) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionsBlockMetaData) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionsBlockMetaData) Version() uint32 {
	return x.message.GetUint32(0)
}

func (x *TransactionsBlockMetaData) RawVersion() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionsBlockMetaData) MutateVersion(v uint32) error {
	return x.message.SetUint32(0, v)
}

func (x *TransactionsBlockMetaData) NodesReputationIterator() *TransactionsBlockMetaDataNodesReputationIterator {
	return &TransactionsBlockMetaDataNodesReputationIterator{iterator: x.message.GetUint32ArrayIterator(1)}
}

type TransactionsBlockMetaDataNodesReputationIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionsBlockMetaDataNodesReputationIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionsBlockMetaDataNodesReputationIterator) NextNodesReputation() uint32 {
	return i.iterator.NextUint32()
}

func (x *TransactionsBlockMetaData) RawNodesReputationArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type TransactionsBlockMetaDataBuilder struct {
	builder membuffers.Builder
	Version uint32
	NodesReputation []uint32
}

func (w *TransactionsBlockMetaDataBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.Version)
	w.builder.WriteUint32Array(buf, w.NodesReputation)
	return nil
}

func (w *TransactionsBlockMetaDataBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionsBlockMetaDataBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *TransactionsBlockMetaDataBuilder) Build() *TransactionsBlockMetaData {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionsBlockMetaDataReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionsBlockBody

// reader

type TransactionsBlockBody struct {
	message membuffers.Message
}

var _TransactionsBlockBody_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,}
var _TransactionsBlockBody_Unions = [][]membuffers.FieldType{}

func TransactionsBlockBodyReader(buf []byte) *TransactionsBlockBody {
	x := &TransactionsBlockBody{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _TransactionsBlockBody_Scheme, _TransactionsBlockBody_Unions)
	return x
}

func (x *TransactionsBlockBody) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionsBlockBody) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionsBlockBody) TransactionIterator() *TransactionsBlockBodyTransactionIterator {
	return &TransactionsBlockBodyTransactionIterator{iterator: x.message.GetMessageArrayIterator(0)}
}

type TransactionsBlockBodyTransactionIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionsBlockBodyTransactionIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionsBlockBodyTransactionIterator) NextTransaction() *SignedTransaction {
	b, s := i.iterator.NextMessage()
	return SignedTransactionReader(b[:s])
}

func (x *TransactionsBlockBody) RawTransactionArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type TransactionsBlockBodyBuilder struct {
	builder membuffers.Builder
	Transaction []*SignedTransactionBuilder
}

func (w *TransactionsBlockBodyBuilder) arrayOfTransaction() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.Transaction))
	for i, v := range w.Transaction {
		res[i] = v
	}
	return res
}

func (w *TransactionsBlockBodyBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessageArray(buf, w.arrayOfTransaction())
	if err != nil {
		return
	}
	return nil
}

func (w *TransactionsBlockBodyBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionsBlockBodyBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *TransactionsBlockBodyBuilder) Build() *TransactionsBlockBody {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionsBlockBodyReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionsBlockProof

// reader

type TransactionsBlockProof struct {
	message membuffers.Message
}

var _TransactionsBlockProof_Scheme = []membuffers.FieldType{membuffers.TypeBytes,}
var _TransactionsBlockProof_Unions = [][]membuffers.FieldType{}

func TransactionsBlockProofReader(buf []byte) *TransactionsBlockProof {
	x := &TransactionsBlockProof{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _TransactionsBlockProof_Scheme, _TransactionsBlockProof_Unions)
	return x
}

func (x *TransactionsBlockProof) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionsBlockProof) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *TransactionsBlockProof) BlockProof() []byte {
	return x.message.GetBytes(0)
}

func (x *TransactionsBlockProof) RawBlockProof() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionsBlockProof) MutateBlockProof(v []byte) error {
	return x.message.SetBytes(0, v)
}

// builder

type TransactionsBlockProofBuilder struct {
	builder membuffers.Builder
	BlockProof []byte
}

func (w *TransactionsBlockProofBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.BlockProof)
	return nil
}

func (w *TransactionsBlockProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionsBlockProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *TransactionsBlockProofBuilder) Build() *TransactionsBlockProof {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionsBlockProofReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ResultsBlockHeader

// reader

type ResultsBlockHeader struct {
	message membuffers.Message
}

var _ResultsBlockHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeBytes,membuffers.TypeUint64,membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytes,}
var _ResultsBlockHeader_Unions = [][]membuffers.FieldType{}

func ResultsBlockHeaderReader(buf []byte) *ResultsBlockHeader {
	x := &ResultsBlockHeader{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _ResultsBlockHeader_Scheme, _ResultsBlockHeader_Unions)
	return x
}

func (x *ResultsBlockHeader) IsValid() bool {
	return x.message.IsValid()
}

func (x *ResultsBlockHeader) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ResultsBlockHeader) Version() uint32 {
	return x.message.GetUint32(0)
}

func (x *ResultsBlockHeader) RawVersion() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ResultsBlockHeader) MutateVersion(v uint32) error {
	return x.message.SetUint32(0, v)
}

func (x *ResultsBlockHeader) VirtualChain() uint32 {
	return x.message.GetUint32(1)
}

func (x *ResultsBlockHeader) RawVirtualChain() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *ResultsBlockHeader) MutateVirtualChain(v uint32) error {
	return x.message.SetUint32(1, v)
}

func (x *ResultsBlockHeader) BlockHeight() uint64 {
	return x.message.GetUint64(2)
}

func (x *ResultsBlockHeader) RawBlockHeight() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *ResultsBlockHeader) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(2, v)
}

func (x *ResultsBlockHeader) PrevBlockHashPtr() primitives.Sha256 {
	return x.message.GetBytes(3)
}

func (x *ResultsBlockHeader) RawPrevBlockHashPtr() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *ResultsBlockHeader) MutatePrevBlockHashPtr(v primitives.Sha256) error {
	return x.message.SetBytes(3, v)
}

func (x *ResultsBlockHeader) Timestamp() uint64 {
	return x.message.GetUint64(4)
}

func (x *ResultsBlockHeader) RawTimestamp() []byte {
	return x.message.RawBufferForField(4, 0)
}

func (x *ResultsBlockHeader) MutateTimestamp(v uint64) error {
	return x.message.SetUint64(4, v)
}

func (x *ResultsBlockHeader) ReceiptsRootHash() primitives.Sha256 {
	return x.message.GetBytes(5)
}

func (x *ResultsBlockHeader) RawReceiptsRootHash() []byte {
	return x.message.RawBufferForField(5, 0)
}

func (x *ResultsBlockHeader) MutateReceiptsRootHash(v primitives.Sha256) error {
	return x.message.SetBytes(5, v)
}

func (x *ResultsBlockHeader) StateDiffHash() primitives.Sha256 {
	return x.message.GetBytes(6)
}

func (x *ResultsBlockHeader) RawStateDiffHash() []byte {
	return x.message.RawBufferForField(6, 0)
}

func (x *ResultsBlockHeader) MutateStateDiffHash(v primitives.Sha256) error {
	return x.message.SetBytes(6, v)
}

func (x *ResultsBlockHeader) TransactionsBlockHashPtr() primitives.Sha256 {
	return x.message.GetBytes(7)
}

func (x *ResultsBlockHeader) RawTransactionsBlockHashPtr() []byte {
	return x.message.RawBufferForField(7, 0)
}

func (x *ResultsBlockHeader) MutateTransactionsBlockHashPtr(v primitives.Sha256) error {
	return x.message.SetBytes(7, v)
}

func (x *ResultsBlockHeader) PreExecutionStateRootHash() primitives.Sha256 {
	return x.message.GetBytes(8)
}

func (x *ResultsBlockHeader) RawPreExecutionStateRootHash() []byte {
	return x.message.RawBufferForField(8, 0)
}

func (x *ResultsBlockHeader) MutatePreExecutionStateRootHash(v primitives.Sha256) error {
	return x.message.SetBytes(8, v)
}

func (x *ResultsBlockHeader) TxidBloomFilter() []byte {
	return x.message.GetBytes(9)
}

func (x *ResultsBlockHeader) RawTxidBloomFilter() []byte {
	return x.message.RawBufferForField(9, 0)
}

func (x *ResultsBlockHeader) MutateTxidBloomFilter(v []byte) error {
	return x.message.SetBytes(9, v)
}

func (x *ResultsBlockHeader) TimestampBloomFilter() []byte {
	return x.message.GetBytes(10)
}

func (x *ResultsBlockHeader) RawTimestampBloomFilter() []byte {
	return x.message.RawBufferForField(10, 0)
}

func (x *ResultsBlockHeader) MutateTimestampBloomFilter(v []byte) error {
	return x.message.SetBytes(10, v)
}

// builder

type ResultsBlockHeaderBuilder struct {
	builder membuffers.Builder
	Version uint32
	VirtualChain uint32
	BlockHeight uint64
	PrevBlockHashPtr primitives.Sha256
	Timestamp uint64
	ReceiptsRootHash primitives.Sha256
	StateDiffHash primitives.Sha256
	TransactionsBlockHashPtr primitives.Sha256
	PreExecutionStateRootHash primitives.Sha256
	TxidBloomFilter []byte
	TimestampBloomFilter []byte
}

func (w *ResultsBlockHeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.Version)
	w.builder.WriteUint32(buf, w.VirtualChain)
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteBytes(buf, w.PrevBlockHashPtr)
	w.builder.WriteUint64(buf, w.Timestamp)
	w.builder.WriteBytes(buf, w.ReceiptsRootHash)
	w.builder.WriteBytes(buf, w.StateDiffHash)
	w.builder.WriteBytes(buf, w.TransactionsBlockHashPtr)
	w.builder.WriteBytes(buf, w.PreExecutionStateRootHash)
	w.builder.WriteBytes(buf, w.TxidBloomFilter)
	w.builder.WriteBytes(buf, w.TimestampBloomFilter)
	return nil
}

func (w *ResultsBlockHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ResultsBlockHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ResultsBlockHeaderBuilder) Build() *ResultsBlockHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ResultsBlockHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ResultsBlockBody

// reader

type ResultsBlockBody struct {
	message membuffers.Message
}

var _ResultsBlockBody_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeMessageArray,}
var _ResultsBlockBody_Unions = [][]membuffers.FieldType{}

func ResultsBlockBodyReader(buf []byte) *ResultsBlockBody {
	x := &ResultsBlockBody{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _ResultsBlockBody_Scheme, _ResultsBlockBody_Unions)
	return x
}

func (x *ResultsBlockBody) IsValid() bool {
	return x.message.IsValid()
}

func (x *ResultsBlockBody) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ResultsBlockBody) TransactionReceiptIterator() *ResultsBlockBodyTransactionReceiptIterator {
	return &ResultsBlockBodyTransactionReceiptIterator{iterator: x.message.GetMessageArrayIterator(0)}
}

type ResultsBlockBodyTransactionReceiptIterator struct {
	iterator *membuffers.Iterator
}

func (i *ResultsBlockBodyTransactionReceiptIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ResultsBlockBodyTransactionReceiptIterator) NextTransactionReceipt() *TransactionReceipt {
	b, s := i.iterator.NextMessage()
	return TransactionReceiptReader(b[:s])
}

func (x *ResultsBlockBody) RawTransactionReceiptArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ResultsBlockBody) ContractStateDiffIterator() *ResultsBlockBodyContractStateDiffIterator {
	return &ResultsBlockBodyContractStateDiffIterator{iterator: x.message.GetMessageArrayIterator(1)}
}

type ResultsBlockBodyContractStateDiffIterator struct {
	iterator *membuffers.Iterator
}

func (i *ResultsBlockBodyContractStateDiffIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ResultsBlockBodyContractStateDiffIterator) NextContractStateDiff() *ContractStateDiff {
	b, s := i.iterator.NextMessage()
	return ContractStateDiffReader(b[:s])
}

func (x *ResultsBlockBody) RawContractStateDiffArray() []byte {
	return x.message.RawBufferForField(1, 0)
}

// builder

type ResultsBlockBodyBuilder struct {
	builder membuffers.Builder
	TransactionReceipt []*TransactionReceiptBuilder
	ContractStateDiff []*ContractStateDiffBuilder
}

func (w *ResultsBlockBodyBuilder) arrayOfTransactionReceipt() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.TransactionReceipt))
	for i, v := range w.TransactionReceipt {
		res[i] = v
	}
	return res
}

func (w *ResultsBlockBodyBuilder) arrayOfContractStateDiff() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.ContractStateDiff))
	for i, v := range w.ContractStateDiff {
		res[i] = v
	}
	return res
}

func (w *ResultsBlockBodyBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
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

func (w *ResultsBlockBodyBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ResultsBlockBodyBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ResultsBlockBodyBuilder) Build() *ResultsBlockBody {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ResultsBlockBodyReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ResultsBlockProof

// reader

type ResultsBlockProof struct {
	message membuffers.Message
}

var _ResultsBlockProof_Scheme = []membuffers.FieldType{membuffers.TypeBytes,}
var _ResultsBlockProof_Unions = [][]membuffers.FieldType{}

func ResultsBlockProofReader(buf []byte) *ResultsBlockProof {
	x := &ResultsBlockProof{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _ResultsBlockProof_Scheme, _ResultsBlockProof_Unions)
	return x
}

func (x *ResultsBlockProof) IsValid() bool {
	return x.message.IsValid()
}

func (x *ResultsBlockProof) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ResultsBlockProof) BlockProof() []byte {
	return x.message.GetBytes(0)
}

func (x *ResultsBlockProof) RawBlockProof() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ResultsBlockProof) MutateBlockProof(v []byte) error {
	return x.message.SetBytes(0, v)
}

// builder

type ResultsBlockProofBuilder struct {
	builder membuffers.Builder
	BlockProof []byte
}

func (w *ResultsBlockProofBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.BlockProof)
	return nil
}

func (w *ResultsBlockProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ResultsBlockProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ResultsBlockProofBuilder) Build() *ResultsBlockProof {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ResultsBlockProofReader(buf)
}

