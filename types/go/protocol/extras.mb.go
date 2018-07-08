// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package protocol

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message PreOrderResult

// reader

type PreOrderResult struct {
	message membuffers.Message
}

var _PreOrderResult_Scheme = []membuffers.FieldType{membuffers.TypeUint16,}
var _PreOrderResult_Unions = [][]membuffers.FieldType{}

func PreOrderResultReader(buf []byte) *PreOrderResult {
	x := &PreOrderResult{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _PreOrderResult_Scheme, _PreOrderResult_Unions)
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

type BlockHeightMode uint16

const (
	BLOCK_HEIGHT_MODE_RESERVED BlockHeightMode = 0
	BLOCK_HEIGHT_MODE_USE_RECENT BlockHeightMode = 1
	BLOCK_HEIGHT_MODE_USE_SPECIFIED BlockHeightMode = 2
)

type AccessScope uint16

const (
	ACCESS_SCOPE_RESERVED AccessScope = 0
	ACCESS_SCOPE_READ_ONLY AccessScope = 1
	ACCESS_SCOPE_READ_WRITE AccessScope = 2
)

type PermissionScope uint16

const (
	PERMISSION_SCOPE_RESERVED PermissionScope = 0
	PERMISSION_SCOPE_SYSTEM PermissionScope = 1
	PERMISSION_SCOPE_SERVICE PermissionScope = 2
)

type DeployServiceStatus uint16

const (
	DEPLOY_SERVICE_STATUS_RESERVED DeployServiceStatus = 0
	DEPLOY_SERVICE_STATUS_SUCCESS DeployServiceStatus = 1
	DEPLOY_SERVICE_STATUS_FAILED DeployServiceStatus = 2
)

type PreOrderStatus uint16

const (
	PRE_ORDER_STATUS_RESERVED PreOrderStatus = 0
	PRE_ORDER_STATUS_VALID PreOrderStatus = 1
	PRE_ORDER_STATUS_INVALID_SIGNATURE PreOrderStatus = 2
	PRE_ORDER_STATUS_INVALID_ADDRESS_SCHEME PreOrderStatus = 3
	PRE_ORDER_STATUS_NOT_APPROVED PreOrderStatus = 4
)

type ValidateBlockConsensus uint16

const (
	VALIDATE_BLOCK_CONSENSUS_RESERVED ValidateBlockConsensus = 0
	VALIDATE_BLOCK_CONSENSUS_VALID_FOR_COMMIT ValidateBlockConsensus = 1
	VALIDATE_BLOCK_CONSENSUS_INVALID_FOR_COMMIT ValidateBlockConsensus = 2
)

