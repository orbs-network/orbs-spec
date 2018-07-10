// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package gossiptopics

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockSync

type BlockSync interface {
	BroadcastBlockSyncAvailabilityRequest(input *BlockSyncAvailabilityRequestInput) (*BlockSyncOutput, error)
	SendBlockSyncAvailabilityResponse(input *BlockSyncAvailabilityResponseInput) (*BlockSyncOutput, error)
	SendBlockSyncRequest(input *BlockSyncRequestInput) (*BlockSyncOutput, error)
	SendBlockSyncResponse(input *BlockSyncResponseInput) (*BlockSyncOutput, error)
	RegisterBlockSyncHandler(handler BlockSyncHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service BlockSyncHandler

type BlockSyncHandler interface {
	HandleBlockAvailabilityRequest(input *BlockSyncAvailabilityRequestInput) (*BlockSyncOutput, error)
	HandleBlockAvailabilityResponse(input *BlockSyncAvailabilityResponseInput) (*BlockSyncOutput, error)
	HandleBlockSyncRequest(input *BlockSyncRequestInput) (*BlockSyncOutput, error)
	HandleBlockSyncResponse(input *BlockSyncResponseInput) (*BlockSyncOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncAvailabilityRequestInput (non serializable)

type BlockSyncAvailabilityRequestInput struct {
	Header *gossipmessages.BlockSyncAvailabilityRequestHeader
	Request *gossipmessages.BlockSyncAvailability
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncAvailabilityResponseInput (non serializable)

type BlockSyncAvailabilityResponseInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Header *gossipmessages.BlockSyncAvailabilityResponseHeader
	Response *gossipmessages.BlockSyncAvailability
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestInput (non serializable)

type BlockSyncRequestInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Header *gossipmessages.BlockSyncRequestHeader
	Request *gossipmessages.BlockSyncRequest
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseInput (non serializable)

type BlockSyncResponseInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Header *gossipmessages.BlockSyncResponseHeader
	Response *gossipmessages.BlockSyncResponse
	BlockPairs []*protocol.BlockPair
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncOutput (non serializable)

type BlockSyncOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// enums

