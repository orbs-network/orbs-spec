// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package gossipmessages

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message HeaderAvailabilityRequestMessage (non serializable)

type HeaderAvailabilityRequestMessage struct {
	SignedBatchRange *HeaderSyncRange
	Sender           *SenderSignature
}

func (x *HeaderAvailabilityRequestMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedBatchRange:%s,Sender:%s,}", x.StringSignedBatchRange(), x.StringSender())
}

func (x *HeaderAvailabilityRequestMessage) StringSignedBatchRange() (res string) {
	res = x.SignedBatchRange.String()
	return
}

func (x *HeaderAvailabilityRequestMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HeaderAvailabilityResponseMessage (non serializable)

type HeaderAvailabilityResponseMessage struct {
	SignedBatchRange *HeaderSyncRange
	Sender           *SenderSignature
}

func (x *HeaderAvailabilityResponseMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedBatchRange:%s,Sender:%s,}", x.StringSignedBatchRange(), x.StringSender())
}

func (x *HeaderAvailabilityResponseMessage) StringSignedBatchRange() (res string) {
	res = x.SignedBatchRange.String()
	return
}

func (x *HeaderAvailabilityResponseMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HeaderSyncRequestMessage (non serializable)

type HeaderSyncRequestMessage struct {
	SignedChunkRange *HeaderSyncRange
	Sender           *SenderSignature
}

func (x *HeaderSyncRequestMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedChunkRange:%s,Sender:%s,}", x.StringSignedChunkRange(), x.StringSender())
}

func (x *HeaderSyncRequestMessage) StringSignedChunkRange() (res string) {
	res = x.SignedChunkRange.String()
	return
}

func (x *HeaderSyncRequestMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HeaderSyncResponseMessage (non serializable)

type HeaderSyncResponseMessage struct {
	SignedChunkRange *HeaderSyncRange
	Sender           *SenderSignature
	HeaderWithProof  []*ResultsBlockHeaderWithProof
}

func (x *HeaderSyncResponseMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedChunkRange:%s,Sender:%s,HeaderWithProof:%s,}", x.StringSignedChunkRange(), x.StringSender(), x.StringHeaderWithProof())
}

func (x *HeaderSyncResponseMessage) StringSignedChunkRange() (res string) {
	res = x.SignedChunkRange.String()
	return
}

func (x *HeaderSyncResponseMessage) StringSender() (res string) {
	res = x.Sender.String()
	return
}

func (x *HeaderSyncResponseMessage) StringHeaderWithProof() (res string) {
	res = "["
	for _, v := range x.HeaderWithProof {
		res += v.String() + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HeaderSyncRange

// reader

type HeaderSyncRange struct {
	// HeaderType HeaderType
	// FirstBlockHeight primitives.BlockHeight
	// LastBlockHeight primitives.BlockHeight
	// LastCommittedBlockHeight primitives.BlockHeight

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *HeaderSyncRange) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{HeaderType:%s,FirstBlockHeight:%s,LastBlockHeight:%s,LastCommittedBlockHeight:%s,}", x.StringHeaderType(), x.StringFirstBlockHeight(), x.StringLastBlockHeight(), x.StringLastCommittedBlockHeight())
}

var _HeaderSyncRange_Scheme = []membuffers.FieldType{membuffers.TypeUint16, membuffers.TypeUint64, membuffers.TypeUint64, membuffers.TypeUint64}
var _HeaderSyncRange_Unions = [][]membuffers.FieldType{}

func HeaderSyncRangeReader(buf []byte) *HeaderSyncRange {
	x := &HeaderSyncRange{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _HeaderSyncRange_Scheme, _HeaderSyncRange_Unions)
	return x
}

func (x *HeaderSyncRange) IsValid() bool {
	return x._message.IsValid()
}

func (x *HeaderSyncRange) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *HeaderSyncRange) Equal(y *HeaderSyncRange) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *HeaderSyncRange) HeaderType() HeaderType {
	return HeaderType(x._message.GetUint16(0))
}

func (x *HeaderSyncRange) RawHeaderType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *HeaderSyncRange) MutateHeaderType(v HeaderType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *HeaderSyncRange) StringHeaderType() string {
	return x.HeaderType().String()
}

func (x *HeaderSyncRange) FirstBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *HeaderSyncRange) RawFirstBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *HeaderSyncRange) MutateFirstBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *HeaderSyncRange) StringFirstBlockHeight() string {
	return fmt.Sprintf("%s", x.FirstBlockHeight())
}

func (x *HeaderSyncRange) LastBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(2))
}

func (x *HeaderSyncRange) RawLastBlockHeight() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *HeaderSyncRange) MutateLastBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(2, uint64(v))
}

func (x *HeaderSyncRange) StringLastBlockHeight() string {
	return fmt.Sprintf("%s", x.LastBlockHeight())
}

func (x *HeaderSyncRange) LastCommittedBlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(3))
}

func (x *HeaderSyncRange) RawLastCommittedBlockHeight() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *HeaderSyncRange) MutateLastCommittedBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *HeaderSyncRange) StringLastCommittedBlockHeight() string {
	return fmt.Sprintf("%s", x.LastCommittedBlockHeight())
}

// builder

type HeaderSyncRangeBuilder struct {
	HeaderType               HeaderType
	FirstBlockHeight         primitives.BlockHeight
	LastBlockHeight          primitives.BlockHeight
	LastCommittedBlockHeight primitives.BlockHeight

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *HeaderSyncRangeBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUint16(buf, uint16(w.HeaderType))
	w._builder.WriteUint64(buf, uint64(w.FirstBlockHeight))
	w._builder.WriteUint64(buf, uint64(w.LastBlockHeight))
	w._builder.WriteUint64(buf, uint64(w.LastCommittedBlockHeight))
	return nil
}

func (w *HeaderSyncRangeBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint16(prefix, offsetFromStart, "HeaderSyncRange.HeaderType", uint16(w.HeaderType))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "HeaderSyncRange.FirstBlockHeight", uint64(w.FirstBlockHeight))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "HeaderSyncRange.LastBlockHeight", uint64(w.LastBlockHeight))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "HeaderSyncRange.LastCommittedBlockHeight", uint64(w.LastCommittedBlockHeight))
	return nil
}

func (w *HeaderSyncRangeBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *HeaderSyncRangeBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *HeaderSyncRangeBuilder) Build() *HeaderSyncRange {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HeaderSyncRangeReader(buf)
}

func HeaderSyncRangeBuilderFromRaw(raw []byte) *HeaderSyncRangeBuilder {
	return &HeaderSyncRangeBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message ResultsBlockHeaderWithProof (non serializable)

type ResultsBlockHeaderWithProof struct {
	Header     *protocol.ResultsBlockHeader
	BlockProof *protocol.ResultsBlockProof
}

func (x *ResultsBlockHeaderWithProof) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Header:%s,BlockProof:%s,}", x.StringHeader(), x.StringBlockProof())
}

func (x *ResultsBlockHeaderWithProof) StringHeader() (res string) {
	res = x.Header.String()
	return
}

func (x *ResultsBlockHeaderWithProof) StringBlockProof() (res string) {
	res = x.BlockProof.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums

type HeaderSyncMessageType uint16

const (
	HEADER_SYNC_RESERVED              HeaderSyncMessageType = 0
	HEADER_SYNC_AVAILABILITY_REQUEST  HeaderSyncMessageType = 1
	HEADER_SYNC_AVAILABILITY_RESPONSE HeaderSyncMessageType = 2
	HEADER_SYNC_REQUEST               HeaderSyncMessageType = 3
	HEADER_SYNC_RESPONSE              HeaderSyncMessageType = 4
)

func (n HeaderSyncMessageType) String() string {
	switch n {
	case HEADER_SYNC_RESERVED:
		return "HEADER_SYNC_RESERVED"
	case HEADER_SYNC_AVAILABILITY_REQUEST:
		return "HEADER_SYNC_AVAILABILITY_REQUEST"
	case HEADER_SYNC_AVAILABILITY_RESPONSE:
		return "HEADER_SYNC_AVAILABILITY_RESPONSE"
	case HEADER_SYNC_REQUEST:
		return "HEADER_SYNC_REQUEST"
	case HEADER_SYNC_RESPONSE:
		return "HEADER_SYNC_RESPONSE"
	}
	return "UNKNOWN"
}

type HeaderType uint16

const (
	HEADER_TYPE_RESERVED                        HeaderType = 0
	HEADER_TYPE_RESULTS_BLOCK_HEADER_WITH_PROOF HeaderType = 1
)

func (n HeaderType) String() string {
	switch n {
	case HEADER_TYPE_RESERVED:
		return "HEADER_TYPE_RESERVED"
	case HEADER_TYPE_RESULTS_BLOCK_HEADER_WITH_PROOF:
		return "HEADER_TYPE_RESULTS_BLOCK_HEADER_WITH_PROOF"
	}
	return "UNKNOWN"
}
