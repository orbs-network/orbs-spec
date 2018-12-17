// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package protocol

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/consensus"
)

/////////////////////////////////////////////////////////////////////////////
// message BlockPairContainer (non serializable)

type BlockPairContainer struct {
	TransactionsBlock *TransactionsBlockContainer
	ResultsBlock      *ResultsBlockContainer
}

func (x *BlockPairContainer) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionsBlock:%s,ResultsBlock:%s,}", x.StringTransactionsBlock(), x.StringResultsBlock())
}

func (x *BlockPairContainer) StringTransactionsBlock() (res string) {
	res = x.TransactionsBlock.String()
	return
}

func (x *BlockPairContainer) StringResultsBlock() (res string) {
	res = x.ResultsBlock.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionsBlockContainer (non serializable)

type TransactionsBlockContainer struct {
	Header             *TransactionsBlockHeader
	Metadata           *TransactionsBlockMetadata
	SignedTransactions []*SignedTransaction
	BlockProof         *TransactionsBlockProof
}

func (x *TransactionsBlockContainer) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Header:%s,Metadata:%s,SignedTransactions:%s,BlockProof:%s,}", x.StringHeader(), x.StringMetadata(), x.StringSignedTransactions(), x.StringBlockProof())
}

func (x *TransactionsBlockContainer) StringHeader() (res string) {
	res = x.Header.String()
	return
}

func (x *TransactionsBlockContainer) StringMetadata() (res string) {
	res = x.Metadata.String()
	return
}

func (x *TransactionsBlockContainer) StringSignedTransactions() (res string) {
	res = "["
	for _, v := range x.SignedTransactions {
		res += v.String() + ","
	}
	res += "]"
	return
}

func (x *TransactionsBlockContainer) StringBlockProof() (res string) {
	res = x.BlockProof.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ResultsBlockContainer (non serializable)

type ResultsBlockContainer struct {
	Header                  *ResultsBlockHeader
	TransactionReceipts     []*TransactionReceipt
	ContractStateDiffs      []*ContractStateDiff
	TransactionsBloomFilter *TransactionsBloomFilter
	BlockProof              *ResultsBlockProof
}

func (x *ResultsBlockContainer) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Header:%s,TransactionReceipts:%s,ContractStateDiffs:%s,TransactionsBloomFilter:%s,BlockProof:%s,}", x.StringHeader(), x.StringTransactionReceipts(), x.StringContractStateDiffs(), x.StringTransactionsBloomFilter(), x.StringBlockProof())
}

func (x *ResultsBlockContainer) StringHeader() (res string) {
	res = x.Header.String()
	return
}

func (x *ResultsBlockContainer) StringTransactionReceipts() (res string) {
	res = "["
	for _, v := range x.TransactionReceipts {
		res += v.String() + ","
	}
	res += "]"
	return
}

func (x *ResultsBlockContainer) StringContractStateDiffs() (res string) {
	res = "["
	for _, v := range x.ContractStateDiffs {
		res += v.String() + ","
	}
	res += "]"
	return
}

func (x *ResultsBlockContainer) StringTransactionsBloomFilter() (res string) {
	res = x.TransactionsBloomFilter.String()
	return
}

func (x *ResultsBlockContainer) StringBlockProof() (res string) {
	res = x.BlockProof.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionsBlockHeader

// reader

type TransactionsBlockHeader struct {
	// ProtocolVersion primitives.ProtocolVersion
	// VirtualChainId primitives.VirtualChainId
	// BlockHeight primitives.BlockHeight
	// PrevBlockHashPtr primitives.Sha256
	// Timestamp primitives.TimestampNano
	// TransactionsRootHash primitives.MerkleSha256
	// MetadataHash primitives.Sha256
	// NumSignedTransactions uint32

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *TransactionsBlockHeader) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ProtocolVersion:%s,VirtualChainId:%s,BlockHeight:%s,PrevBlockHashPtr:%s,Timestamp:%s,TransactionsRootHash:%s,MetadataHash:%s,NumSignedTransactions:%s,}", x.StringProtocolVersion(), x.StringVirtualChainId(), x.StringBlockHeight(), x.StringPrevBlockHashPtr(), x.StringTimestamp(), x.StringTransactionsRootHash(), x.StringMetadataHash(), x.StringNumSignedTransactions())
}

var _TransactionsBlockHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint32, membuffers.TypeUint32, membuffers.TypeUint64, membuffers.TypeBytes, membuffers.TypeUint64, membuffers.TypeBytes, membuffers.TypeBytes, membuffers.TypeUint32}
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

func (x *TransactionsBlockHeader) Equal(y *TransactionsBlockHeader) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
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

func (x *TransactionsBlockHeader) StringProtocolVersion() string {
	return fmt.Sprintf("%s", x.ProtocolVersion())
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

func (x *TransactionsBlockHeader) StringVirtualChainId() string {
	return fmt.Sprintf("%s", x.VirtualChainId())
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

func (x *TransactionsBlockHeader) StringBlockHeight() string {
	return fmt.Sprintf("%s", x.BlockHeight())
}

func (x *TransactionsBlockHeader) PrevBlockHashPtr() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(3))
}

func (x *TransactionsBlockHeader) RawPrevBlockHashPtr() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *TransactionsBlockHeader) RawPrevBlockHashPtrWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *TransactionsBlockHeader) MutatePrevBlockHashPtr(v primitives.Sha256) error {
	return x._message.SetBytes(3, []byte(v))
}

func (x *TransactionsBlockHeader) StringPrevBlockHashPtr() string {
	return fmt.Sprintf("%s", x.PrevBlockHashPtr())
}

func (x *TransactionsBlockHeader) Timestamp() primitives.TimestampNano {
	return primitives.TimestampNano(x._message.GetUint64(4))
}

func (x *TransactionsBlockHeader) RawTimestamp() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *TransactionsBlockHeader) MutateTimestamp(v primitives.TimestampNano) error {
	return x._message.SetUint64(4, uint64(v))
}

func (x *TransactionsBlockHeader) StringTimestamp() string {
	return fmt.Sprintf("%s", x.Timestamp())
}

func (x *TransactionsBlockHeader) TransactionsRootHash() primitives.MerkleSha256 {
	return primitives.MerkleSha256(x._message.GetBytes(5))
}

func (x *TransactionsBlockHeader) RawTransactionsRootHash() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *TransactionsBlockHeader) RawTransactionsRootHashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(5, 0)
}

func (x *TransactionsBlockHeader) MutateTransactionsRootHash(v primitives.MerkleSha256) error {
	return x._message.SetBytes(5, []byte(v))
}

func (x *TransactionsBlockHeader) StringTransactionsRootHash() string {
	return fmt.Sprintf("%s", x.TransactionsRootHash())
}

func (x *TransactionsBlockHeader) MetadataHash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(6))
}

func (x *TransactionsBlockHeader) RawMetadataHash() []byte {
	return x._message.RawBufferForField(6, 0)
}

func (x *TransactionsBlockHeader) RawMetadataHashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(6, 0)
}

func (x *TransactionsBlockHeader) MutateMetadataHash(v primitives.Sha256) error {
	return x._message.SetBytes(6, []byte(v))
}

func (x *TransactionsBlockHeader) StringMetadataHash() string {
	return fmt.Sprintf("%s", x.MetadataHash())
}

func (x *TransactionsBlockHeader) NumSignedTransactions() uint32 {
	return x._message.GetUint32(7)
}

func (x *TransactionsBlockHeader) RawNumSignedTransactions() []byte {
	return x._message.RawBufferForField(7, 0)
}

func (x *TransactionsBlockHeader) MutateNumSignedTransactions(v uint32) error {
	return x._message.SetUint32(7, v)
}

func (x *TransactionsBlockHeader) StringNumSignedTransactions() string {
	return fmt.Sprintf("%x", x.NumSignedTransactions())
}

// builder

type TransactionsBlockHeaderBuilder struct {
	ProtocolVersion       primitives.ProtocolVersion
	VirtualChainId        primitives.VirtualChainId
	BlockHeight           primitives.BlockHeight
	PrevBlockHashPtr      primitives.Sha256
	Timestamp             primitives.TimestampNano
	TransactionsRootHash  primitives.MerkleSha256
	MetadataHash          primitives.Sha256
	NumSignedTransactions uint32

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *TransactionsBlockHeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteUint32(buf, uint32(w.ProtocolVersion))
	w._builder.WriteUint32(buf, uint32(w.VirtualChainId))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteBytes(buf, []byte(w.PrevBlockHashPtr))
	w._builder.WriteUint64(buf, uint64(w.Timestamp))
	w._builder.WriteBytes(buf, []byte(w.TransactionsRootHash))
	w._builder.WriteBytes(buf, []byte(w.MetadataHash))
	w._builder.WriteUint32(buf, w.NumSignedTransactions)
	return nil
}

func (w *TransactionsBlockHeaderBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint32(prefix, offsetFromStart, "TransactionsBlockHeader.ProtocolVersion", uint32(w.ProtocolVersion))
	w._builder.HexDumpUint32(prefix, offsetFromStart, "TransactionsBlockHeader.VirtualChainId", uint32(w.VirtualChainId))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "TransactionsBlockHeader.BlockHeight", uint64(w.BlockHeight))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "TransactionsBlockHeader.PrevBlockHashPtr", []byte(w.PrevBlockHashPtr))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "TransactionsBlockHeader.Timestamp", uint64(w.Timestamp))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "TransactionsBlockHeader.TransactionsRootHash", []byte(w.TransactionsRootHash))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "TransactionsBlockHeader.MetadataHash", []byte(w.MetadataHash))
	w._builder.HexDumpUint32(prefix, offsetFromStart, "TransactionsBlockHeader.NumSignedTransactions", w.NumSignedTransactions)
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

func TransactionsBlockHeaderBuilderFromRaw(raw []byte) *TransactionsBlockHeaderBuilder {
	return &TransactionsBlockHeaderBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message ResultsBlockHeader

// reader

type ResultsBlockHeader struct {
	// ProtocolVersion primitives.ProtocolVersion
	// VirtualChainId primitives.VirtualChainId
	// BlockHeight primitives.BlockHeight
	// PrevBlockHashPtr primitives.Sha256
	// Timestamp primitives.TimestampNano
	// ReceiptsRootHash primitives.MerkleSha256
	// StateDiffHash primitives.Sha256
	// TransactionsBlockHashPtr primitives.Sha256
	// PreExecutionStateRootHash primitives.MerkleSha256
	// TransactionsBloomFilterHash primitives.Sha256
	// NumTransactionReceipts uint32
	// NumContractStateDiffs uint32

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *ResultsBlockHeader) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ProtocolVersion:%s,VirtualChainId:%s,BlockHeight:%s,PrevBlockHashPtr:%s,Timestamp:%s,ReceiptsRootHash:%s,StateDiffHash:%s,TransactionsBlockHashPtr:%s,PreExecutionStateRootHash:%s,TransactionsBloomFilterHash:%s,NumTransactionReceipts:%s,NumContractStateDiffs:%s,}", x.StringProtocolVersion(), x.StringVirtualChainId(), x.StringBlockHeight(), x.StringPrevBlockHashPtr(), x.StringTimestamp(), x.StringReceiptsRootHash(), x.StringStateDiffHash(), x.StringTransactionsBlockHashPtr(), x.StringPreExecutionStateRootHash(), x.StringTransactionsBloomFilterHash(), x.StringNumTransactionReceipts(), x.StringNumContractStateDiffs())
}

var _ResultsBlockHeader_Scheme = []membuffers.FieldType{membuffers.TypeUint32, membuffers.TypeUint32, membuffers.TypeUint64, membuffers.TypeBytes, membuffers.TypeUint64, membuffers.TypeBytes, membuffers.TypeBytes, membuffers.TypeBytes, membuffers.TypeBytes, membuffers.TypeBytes, membuffers.TypeUint32, membuffers.TypeUint32}
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

func (x *ResultsBlockHeader) Equal(y *ResultsBlockHeader) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
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

func (x *ResultsBlockHeader) StringProtocolVersion() string {
	return fmt.Sprintf("%s", x.ProtocolVersion())
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

func (x *ResultsBlockHeader) StringVirtualChainId() string {
	return fmt.Sprintf("%s", x.VirtualChainId())
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

func (x *ResultsBlockHeader) StringBlockHeight() string {
	return fmt.Sprintf("%s", x.BlockHeight())
}

func (x *ResultsBlockHeader) PrevBlockHashPtr() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(3))
}

func (x *ResultsBlockHeader) RawPrevBlockHashPtr() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *ResultsBlockHeader) RawPrevBlockHashPtrWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *ResultsBlockHeader) MutatePrevBlockHashPtr(v primitives.Sha256) error {
	return x._message.SetBytes(3, []byte(v))
}

func (x *ResultsBlockHeader) StringPrevBlockHashPtr() string {
	return fmt.Sprintf("%s", x.PrevBlockHashPtr())
}

func (x *ResultsBlockHeader) Timestamp() primitives.TimestampNano {
	return primitives.TimestampNano(x._message.GetUint64(4))
}

func (x *ResultsBlockHeader) RawTimestamp() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *ResultsBlockHeader) MutateTimestamp(v primitives.TimestampNano) error {
	return x._message.SetUint64(4, uint64(v))
}

func (x *ResultsBlockHeader) StringTimestamp() string {
	return fmt.Sprintf("%s", x.Timestamp())
}

func (x *ResultsBlockHeader) ReceiptsRootHash() primitives.MerkleSha256 {
	return primitives.MerkleSha256(x._message.GetBytes(5))
}

func (x *ResultsBlockHeader) RawReceiptsRootHash() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *ResultsBlockHeader) RawReceiptsRootHashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(5, 0)
}

func (x *ResultsBlockHeader) MutateReceiptsRootHash(v primitives.MerkleSha256) error {
	return x._message.SetBytes(5, []byte(v))
}

func (x *ResultsBlockHeader) StringReceiptsRootHash() string {
	return fmt.Sprintf("%s", x.ReceiptsRootHash())
}

func (x *ResultsBlockHeader) StateDiffHash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(6))
}

func (x *ResultsBlockHeader) RawStateDiffHash() []byte {
	return x._message.RawBufferForField(6, 0)
}

func (x *ResultsBlockHeader) RawStateDiffHashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(6, 0)
}

func (x *ResultsBlockHeader) MutateStateDiffHash(v primitives.Sha256) error {
	return x._message.SetBytes(6, []byte(v))
}

func (x *ResultsBlockHeader) StringStateDiffHash() string {
	return fmt.Sprintf("%s", x.StateDiffHash())
}

func (x *ResultsBlockHeader) TransactionsBlockHashPtr() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(7))
}

func (x *ResultsBlockHeader) RawTransactionsBlockHashPtr() []byte {
	return x._message.RawBufferForField(7, 0)
}

func (x *ResultsBlockHeader) RawTransactionsBlockHashPtrWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(7, 0)
}

func (x *ResultsBlockHeader) MutateTransactionsBlockHashPtr(v primitives.Sha256) error {
	return x._message.SetBytes(7, []byte(v))
}

func (x *ResultsBlockHeader) StringTransactionsBlockHashPtr() string {
	return fmt.Sprintf("%s", x.TransactionsBlockHashPtr())
}

func (x *ResultsBlockHeader) PreExecutionStateRootHash() primitives.MerkleSha256 {
	return primitives.MerkleSha256(x._message.GetBytes(8))
}

func (x *ResultsBlockHeader) RawPreExecutionStateRootHash() []byte {
	return x._message.RawBufferForField(8, 0)
}

func (x *ResultsBlockHeader) RawPreExecutionStateRootHashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(8, 0)
}

func (x *ResultsBlockHeader) MutatePreExecutionStateRootHash(v primitives.MerkleSha256) error {
	return x._message.SetBytes(8, []byte(v))
}

func (x *ResultsBlockHeader) StringPreExecutionStateRootHash() string {
	return fmt.Sprintf("%s", x.PreExecutionStateRootHash())
}

func (x *ResultsBlockHeader) TransactionsBloomFilterHash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(9))
}

func (x *ResultsBlockHeader) RawTransactionsBloomFilterHash() []byte {
	return x._message.RawBufferForField(9, 0)
}

func (x *ResultsBlockHeader) RawTransactionsBloomFilterHashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(9, 0)
}

func (x *ResultsBlockHeader) MutateTransactionsBloomFilterHash(v primitives.Sha256) error {
	return x._message.SetBytes(9, []byte(v))
}

func (x *ResultsBlockHeader) StringTransactionsBloomFilterHash() string {
	return fmt.Sprintf("%s", x.TransactionsBloomFilterHash())
}

func (x *ResultsBlockHeader) NumTransactionReceipts() uint32 {
	return x._message.GetUint32(10)
}

func (x *ResultsBlockHeader) RawNumTransactionReceipts() []byte {
	return x._message.RawBufferForField(10, 0)
}

func (x *ResultsBlockHeader) MutateNumTransactionReceipts(v uint32) error {
	return x._message.SetUint32(10, v)
}

func (x *ResultsBlockHeader) StringNumTransactionReceipts() string {
	return fmt.Sprintf("%x", x.NumTransactionReceipts())
}

func (x *ResultsBlockHeader) NumContractStateDiffs() uint32 {
	return x._message.GetUint32(11)
}

func (x *ResultsBlockHeader) RawNumContractStateDiffs() []byte {
	return x._message.RawBufferForField(11, 0)
}

func (x *ResultsBlockHeader) MutateNumContractStateDiffs(v uint32) error {
	return x._message.SetUint32(11, v)
}

func (x *ResultsBlockHeader) StringNumContractStateDiffs() string {
	return fmt.Sprintf("%x", x.NumContractStateDiffs())
}

// builder

type ResultsBlockHeaderBuilder struct {
	ProtocolVersion             primitives.ProtocolVersion
	VirtualChainId              primitives.VirtualChainId
	BlockHeight                 primitives.BlockHeight
	PrevBlockHashPtr            primitives.Sha256
	Timestamp                   primitives.TimestampNano
	ReceiptsRootHash            primitives.MerkleSha256
	StateDiffHash               primitives.Sha256
	TransactionsBlockHashPtr    primitives.Sha256
	PreExecutionStateRootHash   primitives.MerkleSha256
	TransactionsBloomFilterHash primitives.Sha256
	NumTransactionReceipts      uint32
	NumContractStateDiffs       uint32

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *ResultsBlockHeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
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
	w._builder.WriteBytes(buf, []byte(w.TransactionsBloomFilterHash))
	w._builder.WriteUint32(buf, w.NumTransactionReceipts)
	w._builder.WriteUint32(buf, w.NumContractStateDiffs)
	return nil
}

func (w *ResultsBlockHeaderBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint32(prefix, offsetFromStart, "ResultsBlockHeader.ProtocolVersion", uint32(w.ProtocolVersion))
	w._builder.HexDumpUint32(prefix, offsetFromStart, "ResultsBlockHeader.VirtualChainId", uint32(w.VirtualChainId))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "ResultsBlockHeader.BlockHeight", uint64(w.BlockHeight))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "ResultsBlockHeader.PrevBlockHashPtr", []byte(w.PrevBlockHashPtr))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "ResultsBlockHeader.Timestamp", uint64(w.Timestamp))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "ResultsBlockHeader.ReceiptsRootHash", []byte(w.ReceiptsRootHash))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "ResultsBlockHeader.StateDiffHash", []byte(w.StateDiffHash))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "ResultsBlockHeader.TransactionsBlockHashPtr", []byte(w.TransactionsBlockHashPtr))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "ResultsBlockHeader.PreExecutionStateRootHash", []byte(w.PreExecutionStateRootHash))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "ResultsBlockHeader.TransactionsBloomFilterHash", []byte(w.TransactionsBloomFilterHash))
	w._builder.HexDumpUint32(prefix, offsetFromStart, "ResultsBlockHeader.NumTransactionReceipts", w.NumTransactionReceipts)
	w._builder.HexDumpUint32(prefix, offsetFromStart, "ResultsBlockHeader.NumContractStateDiffs", w.NumContractStateDiffs)
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

func ResultsBlockHeaderBuilderFromRaw(raw []byte) *ResultsBlockHeaderBuilder {
	return &ResultsBlockHeaderBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionsBlockMetadata

// reader

type TransactionsBlockMetadata struct {

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *TransactionsBlockMetadata) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
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

func (x *TransactionsBlockMetadata) Equal(y *TransactionsBlockMetadata) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

// builder

type TransactionsBlockMetadataBuilder struct {

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *TransactionsBlockMetadataBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	return nil
}

func (w *TransactionsBlockMetadataBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
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

func TransactionsBlockMetadataBuilderFromRaw(raw []byte) *TransactionsBlockMetadataBuilder {
	return &TransactionsBlockMetadataBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionsBloomFilter

// reader

type TransactionsBloomFilter struct {
	// TxhashBloomFilter primitives.BloomFilter
	// TimestampBloomFilter primitives.BloomFilter

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *TransactionsBloomFilter) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TxhashBloomFilter:%s,TimestampBloomFilter:%s,}", x.StringTxhashBloomFilter(), x.StringTimestampBloomFilter())
}

var _TransactionsBloomFilter_Scheme = []membuffers.FieldType{membuffers.TypeBytes, membuffers.TypeBytes}
var _TransactionsBloomFilter_Unions = [][]membuffers.FieldType{}

func TransactionsBloomFilterReader(buf []byte) *TransactionsBloomFilter {
	x := &TransactionsBloomFilter{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _TransactionsBloomFilter_Scheme, _TransactionsBloomFilter_Unions)
	return x
}

func (x *TransactionsBloomFilter) IsValid() bool {
	return x._message.IsValid()
}

func (x *TransactionsBloomFilter) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *TransactionsBloomFilter) Equal(y *TransactionsBloomFilter) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *TransactionsBloomFilter) TxhashBloomFilter() primitives.BloomFilter {
	return primitives.BloomFilter(x._message.GetBytes(0))
}

func (x *TransactionsBloomFilter) RawTxhashBloomFilter() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *TransactionsBloomFilter) RawTxhashBloomFilterWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *TransactionsBloomFilter) MutateTxhashBloomFilter(v primitives.BloomFilter) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *TransactionsBloomFilter) StringTxhashBloomFilter() string {
	return fmt.Sprintf("%s", x.TxhashBloomFilter())
}

func (x *TransactionsBloomFilter) TimestampBloomFilter() primitives.BloomFilter {
	return primitives.BloomFilter(x._message.GetBytes(1))
}

func (x *TransactionsBloomFilter) RawTimestampBloomFilter() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *TransactionsBloomFilter) RawTimestampBloomFilterWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *TransactionsBloomFilter) MutateTimestampBloomFilter(v primitives.BloomFilter) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *TransactionsBloomFilter) StringTimestampBloomFilter() string {
	return fmt.Sprintf("%s", x.TimestampBloomFilter())
}

// builder

type TransactionsBloomFilterBuilder struct {
	TxhashBloomFilter    primitives.BloomFilter
	TimestampBloomFilter primitives.BloomFilter

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *TransactionsBloomFilterBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteBytes(buf, []byte(w.TxhashBloomFilter))
	w._builder.WriteBytes(buf, []byte(w.TimestampBloomFilter))
	return nil
}

func (w *TransactionsBloomFilterBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "TransactionsBloomFilter.TxhashBloomFilter", []byte(w.TxhashBloomFilter))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "TransactionsBloomFilter.TimestampBloomFilter", []byte(w.TimestampBloomFilter))
	return nil
}

func (w *TransactionsBloomFilterBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TransactionsBloomFilterBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *TransactionsBloomFilterBuilder) Build() *TransactionsBloomFilter {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionsBloomFilterReader(buf)
}

func TransactionsBloomFilterBuilderFromRaw(raw []byte) *TransactionsBloomFilterBuilder {
	return &TransactionsBloomFilterBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionsBlockProof

// reader

type TransactionsBlockProof struct {
	// ResultsBlockHash primitives.Sha256
	// Type TransactionsBlockProofType

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *TransactionsBlockProof) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ResultsBlockHash:%s,Type:%s,}", x.StringResultsBlockHash(), x.StringType())
}

var _TransactionsBlockProof_Scheme = []membuffers.FieldType{membuffers.TypeBytes, membuffers.TypeUnion}
var _TransactionsBlockProof_Unions = [][]membuffers.FieldType{{membuffers.TypeMessage, membuffers.TypeBytes}}

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

func (x *TransactionsBlockProof) Equal(y *TransactionsBlockProof) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *TransactionsBlockProof) ResultsBlockHash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(0))
}

func (x *TransactionsBlockProof) RawResultsBlockHash() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *TransactionsBlockProof) RawResultsBlockHashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *TransactionsBlockProof) MutateResultsBlockHash(v primitives.Sha256) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *TransactionsBlockProof) StringResultsBlockHash() string {
	return fmt.Sprintf("%s", x.ResultsBlockHash())
}

type TransactionsBlockProofType uint16

const (
	TRANSACTIONS_BLOCK_PROOF_TYPE_BENCHMARK_CONSENSUS TransactionsBlockProofType = 0
	TRANSACTIONS_BLOCK_PROOF_TYPE_LEAN_HELIX          TransactionsBlockProofType = 1
)

func (x *TransactionsBlockProof) Type() TransactionsBlockProofType {
	return TransactionsBlockProofType(x._message.GetUnionIndex(1, 0))
}

func (x *TransactionsBlockProof) IsTypeBenchmarkConsensus() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 0)
	return is
}

func (x *TransactionsBlockProof) BenchmarkConsensus() *consensus.BenchmarkConsensusBlockProof {
	is, off := x._message.IsUnionIndex(1, 0, 0)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	b, s := x._message.GetMessageInOffset(off)
	return consensus.BenchmarkConsensusBlockProofReader(b[:s])
}

func (x *TransactionsBlockProof) StringBenchmarkConsensus() string {
	return x.BenchmarkConsensus().String()
}

func (x *TransactionsBlockProof) IsTypeLeanHelix() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 1)
	return is
}

func (x *TransactionsBlockProof) LeanHelix() primitives.LeanHelixBlockProof {
	is, off := x._message.IsUnionIndex(1, 0, 1)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return primitives.LeanHelixBlockProof(x._message.GetBytesInOffset(off))
}

func (x *TransactionsBlockProof) StringLeanHelix() string {
	return fmt.Sprintf("%s", x.LeanHelix())
}

func (x *TransactionsBlockProof) MutateLeanHelix(v primitives.LeanHelixBlockProof) error {
	is, off := x._message.IsUnionIndex(1, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetBytesInOffset(off, []byte(v))
	return nil
}

func (x *TransactionsBlockProof) RawType() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *TransactionsBlockProof) RawTypeWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *TransactionsBlockProof) StringType() string {
	switch x.Type() {
	case TRANSACTIONS_BLOCK_PROOF_TYPE_BENCHMARK_CONSENSUS:
		return "(BenchmarkConsensus)" + x.StringBenchmarkConsensus()
	case TRANSACTIONS_BLOCK_PROOF_TYPE_LEAN_HELIX:
		return "(LeanHelix)" + x.StringLeanHelix()
	}
	return "(Unknown)"
}

// builder

type TransactionsBlockProofBuilder struct {
	ResultsBlockHash   primitives.Sha256
	Type               TransactionsBlockProofType
	BenchmarkConsensus *consensus.BenchmarkConsensusBlockProofBuilder
	LeanHelix          primitives.LeanHelixBlockProof

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *TransactionsBlockProofBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteBytes(buf, []byte(w.ResultsBlockHash))
	w._builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case TRANSACTIONS_BLOCK_PROOF_TYPE_BENCHMARK_CONSENSUS:
		w._builder.WriteMessage(buf, w.BenchmarkConsensus)
	case TRANSACTIONS_BLOCK_PROOF_TYPE_LEAN_HELIX:
		w._builder.WriteBytes(buf, []byte(w.LeanHelix))
	}
	return nil
}

func (w *TransactionsBlockProofBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "TransactionsBlockProof.ResultsBlockHash", []byte(w.ResultsBlockHash))
	w._builder.HexDumpUnionIndex(prefix, offsetFromStart, "TransactionsBlockProof.Type", uint16(w.Type))
	switch w.Type {
	case TRANSACTIONS_BLOCK_PROOF_TYPE_BENCHMARK_CONSENSUS:
		w._builder.HexDumpMessage(prefix, offsetFromStart, "TransactionsBlockProof.BenchmarkConsensus", w.BenchmarkConsensus)
	case TRANSACTIONS_BLOCK_PROOF_TYPE_LEAN_HELIX:
		w._builder.HexDumpBytes(prefix, offsetFromStart, "TransactionsBlockProof.LeanHelix", []byte(w.LeanHelix))
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

func TransactionsBlockProofBuilderFromRaw(raw []byte) *TransactionsBlockProofBuilder {
	return &TransactionsBlockProofBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message ResultsBlockProof

// reader

type ResultsBlockProof struct {
	// TransactionsBlockHash primitives.Sha256
	// Type ResultsBlockProofType

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *ResultsBlockProof) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionsBlockHash:%s,Type:%s,}", x.StringTransactionsBlockHash(), x.StringType())
}

var _ResultsBlockProof_Scheme = []membuffers.FieldType{membuffers.TypeBytes, membuffers.TypeUnion}
var _ResultsBlockProof_Unions = [][]membuffers.FieldType{{membuffers.TypeMessage, membuffers.TypeBytes}}

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

func (x *ResultsBlockProof) Equal(y *ResultsBlockProof) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *ResultsBlockProof) TransactionsBlockHash() primitives.Sha256 {
	return primitives.Sha256(x._message.GetBytes(0))
}

func (x *ResultsBlockProof) RawTransactionsBlockHash() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ResultsBlockProof) RawTransactionsBlockHashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *ResultsBlockProof) MutateTransactionsBlockHash(v primitives.Sha256) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *ResultsBlockProof) StringTransactionsBlockHash() string {
	return fmt.Sprintf("%s", x.TransactionsBlockHash())
}

type ResultsBlockProofType uint16

const (
	RESULTS_BLOCK_PROOF_TYPE_BENCHMARK_CONSENSUS ResultsBlockProofType = 0
	RESULTS_BLOCK_PROOF_TYPE_LEAN_HELIX          ResultsBlockProofType = 1
)

func (x *ResultsBlockProof) Type() ResultsBlockProofType {
	return ResultsBlockProofType(x._message.GetUnionIndex(1, 0))
}

func (x *ResultsBlockProof) IsTypeBenchmarkConsensus() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 0)
	return is
}

func (x *ResultsBlockProof) BenchmarkConsensus() *consensus.BenchmarkConsensusBlockProof {
	is, off := x._message.IsUnionIndex(1, 0, 0)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	b, s := x._message.GetMessageInOffset(off)
	return consensus.BenchmarkConsensusBlockProofReader(b[:s])
}

func (x *ResultsBlockProof) StringBenchmarkConsensus() string {
	return x.BenchmarkConsensus().String()
}

func (x *ResultsBlockProof) IsTypeLeanHelix() bool {
	is, _ := x._message.IsUnionIndex(1, 0, 1)
	return is
}

func (x *ResultsBlockProof) LeanHelix() primitives.LeanHelixBlockProof {
	is, off := x._message.IsUnionIndex(1, 0, 1)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return primitives.LeanHelixBlockProof(x._message.GetBytesInOffset(off))
}

func (x *ResultsBlockProof) StringLeanHelix() string {
	return fmt.Sprintf("%s", x.LeanHelix())
}

func (x *ResultsBlockProof) MutateLeanHelix(v primitives.LeanHelixBlockProof) error {
	is, off := x._message.IsUnionIndex(1, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetBytesInOffset(off, []byte(v))
	return nil
}

func (x *ResultsBlockProof) RawType() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *ResultsBlockProof) RawTypeWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *ResultsBlockProof) StringType() string {
	switch x.Type() {
	case RESULTS_BLOCK_PROOF_TYPE_BENCHMARK_CONSENSUS:
		return "(BenchmarkConsensus)" + x.StringBenchmarkConsensus()
	case RESULTS_BLOCK_PROOF_TYPE_LEAN_HELIX:
		return "(LeanHelix)" + x.StringLeanHelix()
	}
	return "(Unknown)"
}

// builder

type ResultsBlockProofBuilder struct {
	TransactionsBlockHash primitives.Sha256
	Type                  ResultsBlockProofType
	BenchmarkConsensus    *consensus.BenchmarkConsensusBlockProofBuilder
	LeanHelix             primitives.LeanHelixBlockProof

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *ResultsBlockProofBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteBytes(buf, []byte(w.TransactionsBlockHash))
	w._builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case RESULTS_BLOCK_PROOF_TYPE_BENCHMARK_CONSENSUS:
		w._builder.WriteMessage(buf, w.BenchmarkConsensus)
	case RESULTS_BLOCK_PROOF_TYPE_LEAN_HELIX:
		w._builder.WriteBytes(buf, []byte(w.LeanHelix))
	}
	return nil
}

func (w *ResultsBlockProofBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "ResultsBlockProof.TransactionsBlockHash", []byte(w.TransactionsBlockHash))
	w._builder.HexDumpUnionIndex(prefix, offsetFromStart, "ResultsBlockProof.Type", uint16(w.Type))
	switch w.Type {
	case RESULTS_BLOCK_PROOF_TYPE_BENCHMARK_CONSENSUS:
		w._builder.HexDumpMessage(prefix, offsetFromStart, "ResultsBlockProof.BenchmarkConsensus", w.BenchmarkConsensus)
	case RESULTS_BLOCK_PROOF_TYPE_LEAN_HELIX:
		w._builder.HexDumpBytes(prefix, offsetFromStart, "ResultsBlockProof.LeanHelix", []byte(w.LeanHelix))
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

func ResultsBlockProofBuilderFromRaw(raw []byte) *ResultsBlockProofBuilder {
	return &ResultsBlockProofBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// enums

type BlockType uint16

const (
	BLOCK_TYPE_RESERVED           BlockType = 0
	BLOCK_TYPE_BLOCK_PAIR         BlockType = 1
	BLOCK_TYPE_TRANSACTIONS_BLOCK BlockType = 2
	BLOCK_TYPE_RESULTS_BLOCK      BlockType = 3
)

func (n BlockType) String() string {
	switch n {
	case BLOCK_TYPE_RESERVED:
		return "BLOCK_TYPE_RESERVED"
	case BLOCK_TYPE_BLOCK_PAIR:
		return "BLOCK_TYPE_BLOCK_PAIR"
	case BLOCK_TYPE_TRANSACTIONS_BLOCK:
		return "BLOCK_TYPE_TRANSACTIONS_BLOCK"
	case BLOCK_TYPE_RESULTS_BLOCK:
		return "BLOCK_TYPE_RESULTS_BLOCK"
	}
	return "UNKNOWN"
}
