// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package gossip

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/messages"
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
	Header *messages.BlockSyncAvailabilityRequestHeader
	Request *messages.BlockSyncAvailability
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncAvailabilityResponseInput (non serializable)

type BlockSyncAvailabilityResponseInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Header *messages.BlockSyncAvailabilityResponseHeader
	Response *messages.BlockSyncAvailability
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestInput (non serializable)

type BlockSyncRequestInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Header *messages.BlockSyncRequestHeader
	Request *messages.BlockSyncRequest
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseInput (non serializable)

type BlockSyncResponseInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Header *messages.BlockSyncResponseHeader
	Response *messages.BlockSyncResponse
	BlockPairs []*protocol.BlockPair
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncOutput (non serializable)

type BlockSyncOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// enums

