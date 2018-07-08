// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossip"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service Gossip

type Gossip interface {
	BroadcastForwardedTransactions(input *BroadcastForwardedTransactionsInput) (*SendGossipMessageOutput, error)
	BroadcastBlockSyncAvailabilityRequest(input *BroadcastBlockSyncAvailabilityRequestInput) (*SendGossipMessageOutput, error)
	SendBlockSyncAvailabilityResponse(input *SendBlockSyncAvailabilityResponseInput) (*SendGossipMessageOutput, error)
	SendBlockSyncRequest(input *SendBlockSyncRequestInput) (*SendGossipMessageOutput, error)
	SendBlockSyncResponse(input *SendBlockSyncResponseInput) (*SendGossipMessageOutput, error)
	SendLeanHelixPrePrepare(input *SendLeanHelixPrePrepareInput) (*SendGossipMessageOutput, error)
	SendLeanHelixPrepare(input *SendLeanHelixPrepareInput) (*SendGossipMessageOutput, error)
	SendLeanHelixCommit(input *SendLeanHelixCommitInput) (*SendGossipMessageOutput, error)
	SendLeanHelixViewChange(input *SendLeanHelixViewChangeInput) (*SendGossipMessageOutput, error)
	SendLeanHelixNewView(input *SendLeanHelixNewViewInput) (*SendGossipMessageOutput, error)
	RegisterTransactionRelayGossipHandler(handler handlers.TransactionRelayGossipHandler)
	RegisterBlockSyncGossipHandler(handler handlers.BlockSyncGossipHandler)
	RegisterLeanHelixConsensusGossipHandler(handler handlers.LeanHelixConsensusGossipHandler)
}

/////////////////////////////////////////////////////////////////////////////
// message SendGossipMessageOutput (non serializable)

type SendGossipMessageOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// message BroadcastForwardedTransactionsInput (non serializable)

type BroadcastForwardedTransactionsInput struct {
	Message *gossip.ForwardedTransactionsMessage
}

/////////////////////////////////////////////////////////////////////////////
// message BroadcastBlockSyncAvailabilityRequestInput (non serializable)

type BroadcastBlockSyncAvailabilityRequestInput struct {
	Message *gossip.BlockSyncAvailabilityRequestMessage
}

/////////////////////////////////////////////////////////////////////////////
// message SendBlockSyncAvailabilityResponseInput (non serializable)

type SendBlockSyncAvailabilityResponseInput struct {
	Recipient []byte
	Message *gossip.BlockSyncAvailabilityResponseMessage
}

/////////////////////////////////////////////////////////////////////////////
// message SendBlockSyncRequestInput (non serializable)

type SendBlockSyncRequestInput struct {
	Recipient []byte
	Message *gossip.BlockSyncRequestMessage
}

/////////////////////////////////////////////////////////////////////////////
// message SendBlockSyncResponseInput (non serializable)

type SendBlockSyncResponseInput struct {
	Recipient []byte
	Message *gossip.BlockSyncResponseMessage
}

/////////////////////////////////////////////////////////////////////////////
// message SendLeanHelixPrePrepareInput (non serializable)

type SendLeanHelixPrePrepareInput struct {
	Recipient [][]byte
	Message *gossip.LeanHelixPrePrepareMessage
}

/////////////////////////////////////////////////////////////////////////////
// message SendLeanHelixPrepareInput (non serializable)

type SendLeanHelixPrepareInput struct {
	Recipient [][]byte
	Message *gossip.LeanHelixPrepareMessage
}

/////////////////////////////////////////////////////////////////////////////
// message SendLeanHelixCommitInput (non serializable)

type SendLeanHelixCommitInput struct {
	Recipient [][]byte
	RecipientMode gossip.RecipientsListMode
	Message *gossip.LeanHelixCommitMessage
}

/////////////////////////////////////////////////////////////////////////////
// message SendLeanHelixViewChangeInput (non serializable)

type SendLeanHelixViewChangeInput struct {
	Recipient [][]byte
	Message *gossip.LeanHelixViewChangeMessage
}

/////////////////////////////////////////////////////////////////////////////
// message SendLeanHelixNewViewInput (non serializable)

type SendLeanHelixNewViewInput struct {
	Recipient [][]byte
	Message *gossip.LeanHelixNewViewMessage
}

/////////////////////////////////////////////////////////////////////////////
// enums

