// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package gossiptopics

import (
	"fmt"
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

func (x *BlockAvailabilityRequestInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Message:%s,}", x.StringMessage())
}

func (x *BlockAvailabilityRequestInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BlockAvailabilityResponseInput (non serializable)

type BlockAvailabilityResponseInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Message *gossipmessages.BlockAvailabilityResponseMessage
}

func (x *BlockAvailabilityResponseInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientPublicKey:%s,Message:%s,}", x.StringRecipientPublicKey(), x.StringMessage())
}

func (x *BlockAvailabilityResponseInput) StringRecipientPublicKey() (res string) {
	res = fmt.Sprintf("%x", x.RecipientPublicKey)
	return
}

func (x *BlockAvailabilityResponseInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestInput (non serializable)

type BlockSyncRequestInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Message *gossipmessages.BlockSyncRequestMessage
}

func (x *BlockSyncRequestInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientPublicKey:%s,Message:%s,}", x.StringRecipientPublicKey(), x.StringMessage())
}

func (x *BlockSyncRequestInput) StringRecipientPublicKey() (res string) {
	res = fmt.Sprintf("%x", x.RecipientPublicKey)
	return
}

func (x *BlockSyncRequestInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseInput (non serializable)

type BlockSyncResponseInput struct {
	RecipientPublicKey primitives.Ed25519Pkey
	Message *gossipmessages.BlockSyncResponseMessage
}

func (x *BlockSyncResponseInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientPublicKey:%s,Message:%s,}", x.StringRecipientPublicKey(), x.StringMessage())
}

func (x *BlockSyncResponseInput) StringRecipientPublicKey() (res string) {
	res = fmt.Sprintf("%x", x.RecipientPublicKey)
	return
}

func (x *BlockSyncResponseInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums

