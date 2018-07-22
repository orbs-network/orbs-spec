// AUTO GENERATED FILE (by membufc proto compiler v0.0.17)
package gossiptopics

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockSync

type BlockSync interface {
	BroadcastBlockAvailabilityRequest(input *BlockAvailabilityRequestInput) (*EmptyOutput, error)
	SendBlockAvailabilityResponse(input *BlockAvailabilityResponseInput) (*EmptyOutput, error)
	SendBlockSyncRequest(input *BlockSyncRequestInput) (*EmptyOutput, error)
	SendBlockSyncResponse(input *BlockSyncResponseInput) (*EmptyOutput, error)
	RegisterBlockSyncHandler(handler BlockSyncHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service BlockSyncHandler

type BlockSyncHandler interface {
	HandleBlockAvailabilityRequest(input *BlockAvailabilityRequestInput) (*EmptyOutput, error)
	HandleBlockAvailabilityResponse(input *BlockAvailabilityResponseInput) (*EmptyOutput, error)
	HandleBlockSyncRequest(input *BlockSyncRequestInput) (*EmptyOutput, error)
	HandleBlockSyncResponse(input *BlockSyncResponseInput) (*EmptyOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilityRequestInput (non serializable)

type BlockAvailabilityRequestInput struct {
	Message *gossipmessages.BlockAvailabilityRequestMessage
}

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilityResponseInput (non serializable)

type BlockAvailabilityResponseInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Message *gossipmessages.BlockAvailabilityResponseMessage
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestInput (non serializable)

type BlockSyncRequestInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Message *gossipmessages.BlockSyncRequestMessage
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseInput (non serializable)

type BlockSyncResponseInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Message *gossipmessages.BlockSyncResponseMessage
}

/////////////////////////////////////////////////////////////////////////////
// enums

