// AUTO GENERATED FILE (by membufc proto compiler v0.0.14)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/blockproofs"
)

/////////////////////////////////////////////////////////////////////////////
// message BlockPair

// reader

type BlockPair struct {
	// TransactionsBlock TransactionsBlock
	// ResultsBlock ResultsBlock

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _BlockPair_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,}
var _BlockPair_Unions = [][]membuffers.FieldType{}

func BlockPairReader(buf []byte) *BlockPair {
	x := &BlockPair{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _BlockPair_Scheme, _BlockPair_Unions)
	return x
}

func (x *BlockPair) IsValid() bool {
	return x._message.IsValid()
}

func (x *BlockPair) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *BlockPair) TransactionsBlock() *TransactionsBlock {
	b, s := x._message.GetMessage(0)
	return TransactionsBlockReader(b[:s])
}

func (x *BlockPair) RawTransactionsBlock() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *BlockPair) ResultsBlock() *ResultsBlock {
	b, s := x._message.GetMessage(1)
	return ResultsBlockReader(b[:s])
}

func (x *BlockPair) RawResultsBlock() []byte {
	return x._message.RawBufferForField(1, 0)
}

// builder

type BlockPairBuilder struct {
	TransactionsBlock *TransactionsBlockBuilder
	ResultsBlock *ResultsBlockBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
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
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.TransactionsBlock)
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.ResultsBlock)
	if err != nil {
		return
	}
	return nil
}

func (w *BlockPairBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *BlockPairBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
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
	// Header TransactionsBlockHeader
	// Metadata TransactionsBlockMetadata
	// SignedTransactionsOpaque [][]byte
	// BlockProof TransactionsBlockProof

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _TransactionsBlock_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessage,membuffers.TypeBytesArray,membuffers.TypeMessage,}
var _TransactionsBlock_Unions = [][]membuffers.FieldType{}

func TransactionsBlockReader(buf []byte) *TransactionsBlock {
	x := &TransactionsBlock{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _TransactionsBlock_Scheme, _TransactionsBlock_Unions)
	return x
}

func (x *TransactionsBlock) IsValid() bool {
	return x._message.IsValid()
}

func (x *TransactionsBlock) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *TransactionsBlock) Header() *TransactionsBlockHeader {
	b, s := x._message.GetMessage(0)
	return TransactionsBlockHeaderReader(b[:s])
}

func (x *TransactionsBlock) RawHeader() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *TransactionsBlock) Metadata() *TransactionsBlockMetadata {
	b, s := x._message.GetMessage(1)
	return TransactionsBlockMetadataReader(b[:s])
}

func (x *TransactionsBlock) RawMetadata() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *TransactionsBlock) SignedTransactionsOpaqueIterator() *TransactionsBlockSignedTransactionsOpaqueIterator {
	return &TransactionsBlockSignedTransactionsOpaqueIterator{iterator: x._message.GetBytesArrayIterator(2)}
}

type TransactionsBlockSignedTransactionsOpaqueIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionsBlockSignedTransactionsOpaqueIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionsBlockSignedTransactionsOpaqueIterator) NextSignedTransactionsOpaque() []byte {
	return i.iterator.NextBytes()
}

func (x *TransactionsBlock) RawSignedTransactionsOpaqueArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *TransactionsBlock) BlockProof() *TransactionsBlockProof {
	b, s := x._message.GetMessage(3)
	return TransactionsBlockProofReader(b[:s])
}

func (x *TransactionsBlock) RawBlockProof() []byte {
	return x._message.RawBufferForField(3, 0)
}

// builder

type TransactionsBlockBuilder struct {
	Header *TransactionsBlockHeaderBuilder
	Metadata *TransactionsBlockMetadataBuilder
	SignedTransactionsOpaque [][]byte
	BlockProof *TransactionsBlockProofBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
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
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.Header)
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.Metadata)
	if err != nil {
		return
	}
	w._builder.WriteBytesArray(buf, w.SignedTransactionsOpaque)
	err = w._builder.WriteMessage(buf, w.BlockProof)
	if err != nil {
		return
	}
	return nil
}

func (w *TransactionsBlockBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TransactionsBlockBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
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
	// Header ResultsBlockHeader
	// TransactionReceipts []TransactionReceipt
	// ContractStateDiffs []ContractStateDiff
	// BlockProof ResultsBlockProof

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _ResultsBlock_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeMessageArray,membuffers.TypeMessageArray,membuffers.TypeMessage,}
var _ResultsBlock_Unions = [][]membuffers.FieldType{}

func ResultsBlockReader(buf []byte) *ResultsBlock {
	x := &ResultsBlock{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ResultsBlock_Scheme, _ResultsBlock_Unions)
	return x
}

func (x *ResultsBlock) IsValid() bool {
	return x._message.IsValid()
}

func (x *ResultsBlock) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ResultsBlock) Header() *ResultsBlockHeader {
	b, s := x._message.GetMessage(0)
	return ResultsBlockHeaderReader(b[:s])
}

func (x *ResultsBlock) RawHeader() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ResultsBlock) TransactionReceiptsIterator() *ResultsBlockTransactionReceiptsIterator {
	return &ResultsBlockTransactionReceiptsIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type ResultsBlockTransactionReceiptsIterator struct {
	iterator *membuffers.Iterator
}

func (i *ResultsBlockTransactionReceiptsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ResultsBlockTransactionReceiptsIterator) NextTransactionReceipts() *TransactionReceipt {
	b, s := i.iterator.NextMessage()
	return TransactionReceiptReader(b[:s])
}

func (x *ResultsBlock) RawTransactionReceiptsArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *ResultsBlock) ContractStateDiffsIterator() *ResultsBlockContractStateDiffsIterator {
	return &ResultsBlockContractStateDiffsIterator{iterator: x._message.GetMessageArrayIterator(2)}
}

type ResultsBlockContractStateDiffsIterator struct {
	iterator *membuffers.Iterator
}

func (i *ResultsBlockContractStateDiffsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ResultsBlockContractStateDiffsIterator) NextContractStateDiffs() *ContractStateDiff {
	b, s := i.iterator.NextMessage()
	return ContractStateDiffReader(b[:s])
}

func (x *ResultsBlock) RawContractStateDiffsArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *ResultsBlock) BlockProof() *ResultsBlockProof {
	b, s := x._message.GetMessage(3)
	return ResultsBlockProofReader(b[:s])
}

func (x *ResultsBlock) RawBlockProof() []byte {
	return x._message.RawBufferForField(3, 0)
}

// builder

type ResultsBlockBuilder struct {
	Header *ResultsBlockHeaderBuilder
	TransactionReceipts []*TransactionReceiptBuilder
	ContractStateDiffs []*ContractStateDiffBuilder
	BlockProof *ResultsBlockProofBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *ResultsBlockBuilder) arrayOfTransactionReceipts() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.TransactionReceipts))
	for i, v := range w.TransactionReceipts {
		res[i] = v
	}
	return res
}

func (w *ResultsBlockBuilder) arrayOfContractStateDiffs() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.ContractStateDiffs))
	for i, v := range w.ContractStateDiffs {
		res[i] = v
	}
	return res
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
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.Header)
	if err != nil {
		return
	}
	err = w._builder.WriteMessageArray(buf, w.arrayOfTransactionReceipts())
	if err != nil {
		return
	}
	err = w._builder.WriteMessageArray(buf, w.arrayOfContractStateDiffs())
	if err != nil {
		return
	}
	err = w._builder.WriteMessage(buf, w.BlockProof)
	if err != nil {
		return
	}
	return nil
}

func (w *ResultsBlockBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ResultsBlockBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
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
	// ProtocolVersion primitives.ProtocolVersion
	// VirtualChainId primitives.VirtualChainId
	// BlockHeight primitives.BlockHeight
	// PrevBlockHashPtr primitives.Sha256
	// Timestamp primitives.Timestamp
	// TransactionsRootHash primitives.MerkleSha256
	// MetadataHash primitives.Sha256

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _TransactionsBlockHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeBytes,membuffers.TypeUint64,membuffers.TypeBytes,membuffers.TypeBytes,}
var _TransactionsBlockHeader_Unions = [][]membuffers.FieldType{}

func TransactionsBlockHeaderReader(buf []byte) *TransactionsBlockHeader {
	x := &TransactionsBlockHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _TransactionsBlockHeader_Scheme, _TransactionsBlockHeader_Unions)
	return x
}

func (x *TransactionsBlockHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *TransactionsBlockHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *TransactionsBlockHeader) ProtocolVersion() primitives.ProtocolVersion {
	return primitives.ProtocolVersion(x._message.GetUint32(0))
}

func (x *TransactionsBlockHeader) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *TransactionsBlockHeader) MutateProtocolVersion(v primitives.ProtocolVersion) error {
	return x._message.SetUint32(0, uint32(v))
}

func (x *TransactionsBlockHeader) VirtualChainId() primitives.VirtualChainId {
	return primitives.VirtualChainId(x._message.GetUint32(1))
}

func (x *TransactionsBlockHeader) RawVirtualChainId() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *TransactionsBlockHeader) MutateVirtualChainId(v primitives.VirtualChainId) error {
	return x._message.SetUint32(1, uint32(v))
}

func (x *TransactionsBlockHeader) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *TransactionsBlockHeader) RawBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *TransactionsBlockHeader) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *TransactionsBlockHeader) PrevBlockHashPtr() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(3))
}

func (x *TransactionsBlockHeader) RawPrevBlockHashPtr() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *TransactionsBlockHeader) MutatePrevBlockHashPtr(v primitives.Sha256) error {
	return x._message.SetBytes(3, []byte(v))
}

func (x *TransactionsBlockHeader) Timestamp() primitives.Timestamp {
	return primitives.Timestamp(x._message.GetUint64(4))
}

func (x *TransactionsBlockHeader) RawTimestamp() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *TransactionsBlockHeader) MutateTimestamp(v primitives.Timestamp) error {
	return x._message.SetUint64(4, uint64(v))
}

func (x *TransactionsBlockHeader) TransactionsRootHash() primitives.MerkleSha256 {
	return primitives.MerkleSha256(x._message.GetBytes(5))
}

func (x *TransactionsBlockHeader) RawTransactionsRootHash() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *TransactionsBlockHeader) MutateTransactionsRootHash(v primitives.MerkleSha256) error {
	return x._message.SetBytes(5, []byte(v))
}

func (x *TransactionsBlockHeader) MetadataHash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(6))
}

func (x *TransactionsBlockHeader) RawMetadataHash() []byte {
	return x._message.RawBufferForField(6, 0)
}

func (x *TransactionsBlockHeader) MutateMetadataHash(v primitives.Sha256) error {
	return x._message.SetBytes(6, []byte(v))
}

// builder

type TransactionsBlockHeaderBuilder struct {
	ProtocolVersion primitives.ProtocolVersion
	VirtualChainId primitives.VirtualChainId
	BlockHeight primitives.BlockHeight
	PrevBlockHashPtr primitives.Sha256
	Timestamp primitives.Timestamp
	TransactionsRootHash primitives.MerkleSha256
	MetadataHash primitives.Sha256

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
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
	w._builder.Reset()
	w._builder.WriteUint32(buf, uint32(w.ProtocolVersion))
	w._builder.WriteUint32(buf, uint32(w.VirtualChainId))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteBytes(buf, []byte(w.PrevBlockHashPtr))
	w._builder.WriteUint64(buf, uint64(w.Timestamp))
	w._builder.WriteBytes(buf, []byte(w.TransactionsRootHash))
	w._builder.WriteBytes(buf, []byte(w.MetadataHash))
	return nil
}

func (w *TransactionsBlockHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TransactionsBlockHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *TransactionsBlockHeaderBuilder) Build() *TransactionsBlockHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionsBlockHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ResultsBlockHeader

// reader

type ResultsBlockHeader struct {
	// ProtocolVersion primitives.ProtocolVersion
	// VirtualChainId primitives.VirtualChainId
	// BlockHeight primitives.BlockHeight
	// PrevBlockHashPtr primitives.Sha256
	// Timestamp primitives.Timestamp
	// ReceiptsRootHash primitives.MerkleSha256
	// StateDiffHash primitives.Sha256
	// TransactionsBlockHashPtr primitives.Sha256
	// PreExecutionStateRootHash primitives.MerkleSha256
	// TxhashBloomFilter primitives.BloomFilter
	// TimestampBloomFilter primitives.BloomFilter

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _ResultsBlockHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeBytes,membuffers.TypeUint64,membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytes,}
var _ResultsBlockHeader_Unions = [][]membuffers.FieldType{}

func ResultsBlockHeaderReader(buf []byte) *ResultsBlockHeader {
	x := &ResultsBlockHeader{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ResultsBlockHeader_Scheme, _ResultsBlockHeader_Unions)
	return x
}

func (x *ResultsBlockHeader) IsValid() bool {
	return x._message.IsValid()
}

func (x *ResultsBlockHeader) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ResultsBlockHeader) ProtocolVersion() primitives.ProtocolVersion {
	return primitives.ProtocolVersion(x._message.GetUint32(0))
}

func (x *ResultsBlockHeader) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ResultsBlockHeader) MutateProtocolVersion(v primitives.ProtocolVersion) error {
	return x._message.SetUint32(0, uint32(v))
}

func (x *ResultsBlockHeader) VirtualChainId() primitives.VirtualChainId {
	return primitives.VirtualChainId(x._message.GetUint32(1))
}

func (x *ResultsBlockHeader) RawVirtualChainId() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *ResultsBlockHeader) MutateVirtualChainId(v primitives.VirtualChainId) error {
	return x._message.SetUint32(1, uint32(v))
}

func (x *ResultsBlockHeader) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *ResultsBlockHeader) RawBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *ResultsBlockHeader) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *ResultsBlockHeader) PrevBlockHashPtr() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(3))
}

func (x *ResultsBlockHeader) RawPrevBlockHashPtr() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *ResultsBlockHeader) MutatePrevBlockHashPtr(v primitives.Sha256) error {
	return x._message.SetBytes(3, []byte(v))
}

func (x *ResultsBlockHeader) Timestamp() primitives.Timestamp {
	return primitives.Timestamp(x._message.GetUint64(4))
}

func (x *ResultsBlockHeader) RawTimestamp() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *ResultsBlockHeader) MutateTimestamp(v primitives.Timestamp) error {
	return x._message.SetUint64(4, uint64(v))
}

func (x *ResultsBlockHeader) ReceiptsRootHash() primitives.MerkleSha256 {
	return primitives.MerkleSha256(x._message.GetBytes(5))
}

func (x *ResultsBlockHeader) RawReceiptsRootHash() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *ResultsBlockHeader) MutateReceiptsRootHash(v primitives.MerkleSha256) error {
	return x._message.SetBytes(5, []byte(v))
}

func (x *ResultsBlockHeader) StateDiffHash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(6))
}

func (x *ResultsBlockHeader) RawStateDiffHash() []byte {
	return x._message.RawBufferForField(6, 0)
}

func (x *ResultsBlockHeader) MutateStateDiffHash(v primitives.Sha256) error {
	return x._message.SetBytes(6, []byte(v))
}

func (x *ResultsBlockHeader) TransactionsBlockHashPtr() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(7))
}

func (x *ResultsBlockHeader) RawTransactionsBlockHashPtr() []byte {
	return x._message.RawBufferForField(7, 0)
}

func (x *ResultsBlockHeader) MutateTransactionsBlockHashPtr(v primitives.Sha256) error {
	return x._message.SetBytes(7, []byte(v))
}

func (x *ResultsBlockHeader) PreExecutionStateRootHash() primitives.MerkleSha256 {
	return primitives.MerkleSha256(x._message.GetBytes(8))
}

func (x *ResultsBlockHeader) RawPreExecutionStateRootHash() []byte {
	return x._message.RawBufferForField(8, 0)
}

func (x *ResultsBlockHeader) MutatePreExecutionStateRootHash(v primitives.MerkleSha256) error {
	return x._message.SetBytes(8, []byte(v))
}

func (x *ResultsBlockHeader) TxhashBloomFilter() primitives.BloomFilter {
	return primitives.BloomFilter(x._message.GetBytes(9))
}

func (x *ResultsBlockHeader) RawTxhashBloomFilter() []byte {
	return x._message.RawBufferForField(9, 0)
}

func (x *ResultsBlockHeader) MutateTxhashBloomFilter(v primitives.BloomFilter) error {
	return x._message.SetBytes(9, []byte(v))
}

func (x *ResultsBlockHeader) TimestampBloomFilter() primitives.BloomFilter {
	return primitives.BloomFilter(x._message.GetBytes(10))
}

func (x *ResultsBlockHeader) RawTimestampBloomFilter() []byte {
	return x._message.RawBufferForField(10, 0)
}

func (x *ResultsBlockHeader) MutateTimestampBloomFilter(v primitives.BloomFilter) error {
	return x._message.SetBytes(10, []byte(v))
}

// builder

type ResultsBlockHeaderBuilder struct {
	ProtocolVersion primitives.ProtocolVersion
	VirtualChainId primitives.VirtualChainId
	BlockHeight primitives.BlockHeight
	PrevBlockHashPtr primitives.Sha256
	Timestamp primitives.Timestamp
	ReceiptsRootHash primitives.MerkleSha256
	StateDiffHash primitives.Sha256
	TransactionsBlockHashPtr primitives.Sha256
	PreExecutionStateRootHash primitives.MerkleSha256
	TxhashBloomFilter primitives.BloomFilter
	TimestampBloomFilter primitives.BloomFilter

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
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
	w._builder.Reset()
	w._builder.WriteUint32(buf, uint32(w.ProtocolVersion))
	w._builder.WriteUint32(buf, uint32(w.VirtualChainId))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteBytes(buf, []byte(w.PrevBlockHashPtr))
	w._builder.WriteUint64(buf, uint64(w.Timestamp))
	w._builder.WriteBytes(buf, []byte(w.ReceiptsRootHash))
	w._builder.WriteBytes(buf, []byte(w.StateDiffHash))
	w._builder.WriteBytes(buf, []byte(w.TransactionsBlockHashPtr))
	w._builder.WriteBytes(buf, []byte(w.PreExecutionStateRootHash))
	w._builder.WriteBytes(buf, []byte(w.TxhashBloomFilter))
	w._builder.WriteBytes(buf, []byte(w.TimestampBloomFilter))
	return nil
}

func (w *ResultsBlockHeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ResultsBlockHeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ResultsBlockHeaderBuilder) Build() *ResultsBlockHeader {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ResultsBlockHeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionsBlockMetadata

// reader

type TransactionsBlockMetadata struct {

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _TransactionsBlockMetadata_Scheme = []membuffers.FieldType{}
var _TransactionsBlockMetadata_Unions = [][]membuffers.FieldType{}

func TransactionsBlockMetadataReader(buf []byte) *TransactionsBlockMetadata {
	x := &TransactionsBlockMetadata{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _TransactionsBlockMetadata_Scheme, _TransactionsBlockMetadata_Unions)
	return x
}

func (x *TransactionsBlockMetadata) IsValid() bool {
	return x._message.IsValid()
}

func (x *TransactionsBlockMetadata) Raw() []byte {
	return x._message.RawBuffer()
}

// builder

type TransactionsBlockMetadataBuilder struct {

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *TransactionsBlockMetadataBuilder) Write(buf []byte) (err error) {
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

func (w *TransactionsBlockMetadataBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TransactionsBlockMetadataBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *TransactionsBlockMetadataBuilder) Build() *TransactionsBlockMetadata {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionsBlockMetadataReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionsBlockProof

// reader

type TransactionsBlockProof struct {
	// Type TransactionsBlockProofType

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _TransactionsBlockProof_Scheme = []membuffers.FieldType{membuffers.TypeUnion,}
var _TransactionsBlockProof_Unions = [][]membuffers.FieldType{{membuffers.TypeMessage,}}

func TransactionsBlockProofReader(buf []byte) *TransactionsBlockProof {
	x := &TransactionsBlockProof{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _TransactionsBlockProof_Scheme, _TransactionsBlockProof_Unions)
	return x
}

func (x *TransactionsBlockProof) IsValid() bool {
	return x._message.IsValid()
}

func (x *TransactionsBlockProof) Raw() []byte {
	return x._message.RawBuffer()
}

type TransactionsBlockProofType uint16

const (
	TRANSACTIONS_BLOCK_PROOF_TYPE_LEAN_HELIX TransactionsBlockProofType = 0
)

func (x *TransactionsBlockProof) Type() TransactionsBlockProofType {
	return TransactionsBlockProofType(x._message.GetUint16(0))
}

func (x *TransactionsBlockProof) IsTypeLeanHelix() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 0)
	return is
}

func (x *TransactionsBlockProof) LeanHelix() *blockproofs.LeanHelix {
	_, off := x._message.IsUnionIndex(0, 0, 0)
	b, s := x._message.GetMessageInOffset(off)
	return blockproofs.LeanHelixReader(b[:s])
}

func (x *TransactionsBlockProof) RawType() []byte {
	return x._message.RawBufferForField(0, 0)
}

// builder

type TransactionsBlockProofBuilder struct {
	Type TransactionsBlockProofType
	LeanHelix *blockproofs.LeanHelixBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
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
	w._builder.Reset()
	w._builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case TRANSACTIONS_BLOCK_PROOF_TYPE_LEAN_HELIX:
		w._builder.WriteMessage(buf, w.LeanHelix)
	}
	return nil
}

func (w *TransactionsBlockProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TransactionsBlockProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *TransactionsBlockProofBuilder) Build() *TransactionsBlockProof {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionsBlockProofReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ResultsBlockProof

// reader

type ResultsBlockProof struct {
	// Type ResultsBlockProofType

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _ResultsBlockProof_Scheme = []membuffers.FieldType{membuffers.TypeUnion,}
var _ResultsBlockProof_Unions = [][]membuffers.FieldType{{membuffers.TypeMessage,}}

func ResultsBlockProofReader(buf []byte) *ResultsBlockProof {
	x := &ResultsBlockProof{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ResultsBlockProof_Scheme, _ResultsBlockProof_Unions)
	return x
}

func (x *ResultsBlockProof) IsValid() bool {
	return x._message.IsValid()
}

func (x *ResultsBlockProof) Raw() []byte {
	return x._message.RawBuffer()
}

type ResultsBlockProofType uint16

const (
	RESULTS_BLOCK_PROOF_TYPE_LEAN_HELIX ResultsBlockProofType = 0
)

func (x *ResultsBlockProof) Type() ResultsBlockProofType {
	return ResultsBlockProofType(x._message.GetUint16(0))
}

func (x *ResultsBlockProof) IsTypeLeanHelix() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 0)
	return is
}

func (x *ResultsBlockProof) LeanHelix() *blockproofs.LeanHelix {
	_, off := x._message.IsUnionIndex(0, 0, 0)
	b, s := x._message.GetMessageInOffset(off)
	return blockproofs.LeanHelixReader(b[:s])
}

func (x *ResultsBlockProof) RawType() []byte {
	return x._message.RawBufferForField(0, 0)
}

// builder

type ResultsBlockProofBuilder struct {
	Type ResultsBlockProofType
	LeanHelix *blockproofs.LeanHelixBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
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
	w._builder.Reset()
	w._builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case RESULTS_BLOCK_PROOF_TYPE_LEAN_HELIX:
		w._builder.WriteMessage(buf, w.LeanHelix)
	}
	return nil
}

func (w *ResultsBlockProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ResultsBlockProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ResultsBlockProofBuilder) Build() *ResultsBlockProof {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ResultsBlockProofReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

