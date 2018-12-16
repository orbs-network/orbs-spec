// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package primitives

import (
	"bytes"
	"fmt"
)

type ProtocolVersion uint32

func (x ProtocolVersion) String() string {
	return fmt.Sprintf("%x", uint32(x))
}

func (x ProtocolVersion) Equal(y ProtocolVersion) bool {
	return x == y
}

func (x ProtocolVersion) KeyForMap() uint32 {
	return uint32(x)
}

type VirtualChainId uint32

func (x VirtualChainId) String() string {
	return fmt.Sprintf("%x", uint32(x))
}

func (x VirtualChainId) Equal(y VirtualChainId) bool {
	return x == y
}

func (x VirtualChainId) KeyForMap() uint32 {
	return uint32(x)
}

type BlockHeight uint64

func (x BlockHeight) String() string {
	return fmt.Sprintf("%x", uint64(x))
}

func (x BlockHeight) Equal(y BlockHeight) bool {
	return x == y
}

func (x BlockHeight) KeyForMap() uint64 {
	return uint64(x)
}

type TimestampNano uint64

func (x TimestampNano) String() string {
	return fmt.Sprintf("%x", uint64(x))
}

func (x TimestampNano) Equal(y TimestampNano) bool {
	return x == y
}

func (x TimestampNano) KeyForMap() uint64 {
	return uint64(x)
}

type ContractName string

func (x ContractName) String() string {
	return fmt.Sprintf(string(x))
}

func (x ContractName) Equal(y ContractName) bool {
	return x == y
}

func (x ContractName) KeyForMap() string {
	return string(x)
}

type MethodName string

func (x MethodName) String() string {
	return fmt.Sprintf(string(x))
}

func (x MethodName) Equal(y MethodName) bool {
	return x == y
}

func (x MethodName) KeyForMap() string {
	return string(x)
}

type EventName string

func (x EventName) String() string {
	return fmt.Sprintf(string(x))
}

func (x EventName) Equal(y EventName) bool {
	return x == y
}

func (x EventName) KeyForMap() string {
	return string(x)
}

type ExecutionContextId uint32

func (x ExecutionContextId) String() string {
	return fmt.Sprintf("%x", uint32(x))
}

func (x ExecutionContextId) Equal(y ExecutionContextId) bool {
	return x == y
}

func (x ExecutionContextId) KeyForMap() uint32 {
	return uint32(x)
}

type LeanHelixMessageContent []byte

func (x LeanHelixMessageContent) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x LeanHelixMessageContent) Equal(y LeanHelixMessageContent) bool {
	return bytes.Equal(x, y)
}

func (x LeanHelixMessageContent) KeyForMap() string {
	return string(x)
}

type MerkleTreeProof []byte

func (x MerkleTreeProof) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x MerkleTreeProof) Equal(y MerkleTreeProof) bool {
	return bytes.Equal(x, y)
}

func (x MerkleTreeProof) KeyForMap() string {
	return string(x)
}

type PackedReceiptProof []byte

func (x PackedReceiptProof) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x PackedReceiptProof) Equal(y PackedReceiptProof) bool {
	return bytes.Equal(x, y)
}

func (x PackedReceiptProof) KeyForMap() string {
	return string(x)
}

type PackedEventsArray []byte

func (x PackedEventsArray) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x PackedEventsArray) Equal(y PackedEventsArray) bool {
	return bytes.Equal(x, y)
}

func (x PackedEventsArray) KeyForMap() string {
	return string(x)
}

type PackedArgumentArray []byte

func (x PackedArgumentArray) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x PackedArgumentArray) Equal(y PackedArgumentArray) bool {
	return bytes.Equal(x, y)
}

func (x PackedArgumentArray) KeyForMap() string {
	return string(x)
}
