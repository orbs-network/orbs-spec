// AUTO GENERATED FILE (by membufc proto compiler v0.0.11)
package handlers

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossip"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionRelayGossipHandler

type TransactionRelayGossipHandler interface {
	HandleForwardedTransactions(input *HandleForwardedTransactionsInput) (*GossipMessageHandlerOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// service BlockSyncGossipHandler

type BlockSyncGossipHandler interface {
	HandleBlockAvailabilityRequest(input *HandleBlockAvailabilityRequestInput) (*GossipMessageHandlerOutput, error)
	HandleBlockAvailabilityResponse(input *HandleBlockAvailabilityResponseInput) (*GossipMessageHandlerOutput, error)
	HandleBlockSyncRequest(input *HandleBlockSyncRequestInput) (*GossipMessageHandlerOutput, error)
	HandleBlockSyncResponse(input *HandleBlockSyncResponseInput) (*GossipMessageHandlerOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// service LeanHelixConsensusGossipHandler

type LeanHelixConsensusGossipHandler interface {
	HandleLeanHelixPrePrepare(input *HandleLeanHelixPrePrepareInput) (*GossipMessageHandlerOutput, error)
	HandleLeanHelixPrepare(input *HandleLeanHelixPrepareInput) (*GossipMessageHandlerOutput, error)
	HandleLeanHelixCommit(input *HandleLeanHelixCommitInput) (*GossipMessageHandlerOutput, error)
	HandleLeanHelixViewChange(input *HandleLeanHelixViewChangeInput) (*GossipMessageHandlerOutput, error)
	HandleLeanHelixNewView(input *HandleLeanHelixNewViewInput) (*GossipMessageHandlerOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message GossipMessageHandlerOutput (non serializable)

type GossipMessageHandlerOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message HandleForwardedTransactionsInput (non serializable)

type HandleForwardedTransactionsInput struct {
	Sender []byte
	Message *gossip.ForwardedTransactionsMessage
}

/////////////////////////////////////////////////////////////////////////////
// message HandleBlockAvailabilityRequestInput (non serializable)

type HandleBlockAvailabilityRequestInput struct {
	Sender []byte
	Message *gossip.BlockSyncAvailabilityRequestMessage
}

/////////////////////////////////////////////////////////////////////////////
// message HandleBlockAvailabilityResponseInput (non serializable)

type HandleBlockAvailabilityResponseInput struct {
	Sender []byte
	Message *gossip.BlockSyncAvailabilityResponseMessage
}

/////////////////////////////////////////////////////////////////////////////
// message HandleBlockSyncRequestInput (non serializable)

type HandleBlockSyncRequestInput struct {
	Sender []byte
	Message *gossip.BlockSyncRequestMessage
}

/////////////////////////////////////////////////////////////////////////////
// message HandleBlockSyncResponseInput (non serializable)

type HandleBlockSyncResponseInput struct {
	Sender []byte
	Message *gossip.BlockSyncResponseMessage
}

/////////////////////////////////////////////////////////////////////////////
// message HandleLeanHelixPrePrepareInput (non serializable)

type HandleLeanHelixPrePrepareInput struct {
	Sender []byte
	Message *gossip.LeanHelixPrePrepareMessage
}

/////////////////////////////////////////////////////////////////////////////
// message HandleLeanHelixPrepareInput (non serializable)

type HandleLeanHelixPrepareInput struct {
	Sender []byte
	Message *gossip.LeanHelixPrepareMessage
}

/////////////////////////////////////////////////////////////////////////////
// message HandleLeanHelixCommitInput (non serializable)

type HandleLeanHelixCommitInput struct {
	Sender []byte
	Message *gossip.LeanHelixCommitMessage
}

/////////////////////////////////////////////////////////////////////////////
// message HandleLeanHelixViewChangeInput (non serializable)

type HandleLeanHelixViewChangeInput struct {
	Sender []byte
	Message *gossip.LeanHelixViewChangeMessage
}

/////////////////////////////////////////////////////////////////////////////
// message HandleLeanHelixNewViewInput (non serializable)

type HandleLeanHelixNewViewInput struct {
	Sender []byte
	Message *gossip.LeanHelixNewViewMessage
}

/////////////////////////////////////////////////////////////////////////////
// enums

