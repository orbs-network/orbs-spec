// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package gossiptopics

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockSync

type BlockSync interface {
	BroadcastBlockAvailabilityRequest(input *BlockAvailabilityRequestInput) (*BlockSyncOutput, error)
	SendBlockAvailabilityResponse(input *BlockAvailabilityResponseInput) (*BlockSyncOutput, error)
	SendBlockRequest(input *BlockSyncRequestInput) (*BlockSyncOutput, error)
	SendBlockResponse(input *BlockSyncResponseInput) (*BlockSyncOutput, error)
	RegisterBlockSyncHandler(handler BlockSyncHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service BlockSyncHandler

type BlockSyncHandler interface {
	HandleBlockAvailabilityRequest(input *BlockAvailabilityRequestInput) (*BlockSyncOutput, error)
	HandleBlockAvailabilityResponse(input *BlockAvailabilityResponseInput) (*BlockSyncOutput, error)
	HandleBlockSyncRequest(input *BlockSyncRequestInput) (*BlockSyncOutput, error)
	HandleBlockSyncResponse(input *BlockSyncResponseInput) (*BlockSyncOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilityRequestInput (non serializable)

type BlockAvailabilityRequestInput struct {
	RecipientList *gossipmessages.RecipientsList
	BlockAvailabilityRequestMessage *gossipmessages.BlockAvailabilityRequestMessage
}

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilityResponseInput (non serializable)

type BlockAvailabilityResponseInput struct {
	RecipientList *gossipmessages.RecipientsList
	BlockAvailabilityResponseMessage *gossipmessages.BlockAvailabilityResponseMessage
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestInput (non serializable)

type BlockSyncRequestInput struct {
	RecipientList *gossipmessages.RecipientsList
	BlockSyncRequestMessage *gossipmessages.BlockSyncRequestMessage
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseInput (non serializable)

type BlockSyncResponseInput struct {
	RecipientList *gossipmessages.RecipientsList
	BlockSyncResponseMessage *gossipmessages.BlockSyncRequestMessage
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncOutput (non serializable)

type BlockSyncOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// enums

