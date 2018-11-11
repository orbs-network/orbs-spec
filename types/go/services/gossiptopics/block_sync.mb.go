// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package gossiptopics

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockSync

type BlockSync interface {
	BroadcastBlockAvailabilityRequest(ctx context.Context, input *BlockAvailabilityRequestInput) (*EmptyOutput, error)
	SendBlockAvailabilityResponse(ctx context.Context, input *BlockAvailabilityResponseInput) (*EmptyOutput, error)
	SendBlockSyncRequest(ctx context.Context, input *BlockSyncRequestInput) (*EmptyOutput, error)
	SendBlockSyncResponse(ctx context.Context, input *BlockSyncResponseInput) (*EmptyOutput, error)
	RegisterBlockSyncHandler(handler BlockSyncHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service BlockSyncHandler

type BlockSyncHandler interface {
	HandleBlockAvailabilityRequest(ctx context.Context, input *BlockAvailabilityRequestInput) (*EmptyOutput, error)
	HandleBlockAvailabilityResponse(ctx context.Context, input *BlockAvailabilityResponseInput) (*EmptyOutput, error)
	HandleBlockSyncRequest(ctx context.Context, input *BlockSyncRequestInput) (*EmptyOutput, error)
	HandleBlockSyncResponse(ctx context.Context, input *BlockSyncResponseInput) (*EmptyOutput, error)
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
	RecipientPublicKey primitives.Ed25519PublicKey
	Message            *gossipmessages.BlockAvailabilityResponseMessage
}

func (x *BlockAvailabilityResponseInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientPublicKey:%s,Message:%s,}", x.StringRecipientPublicKey(), x.StringMessage())
}

func (x *BlockAvailabilityResponseInput) StringRecipientPublicKey() (res string) {
	res = fmt.Sprintf("%s", x.RecipientPublicKey)
	return
}

func (x *BlockAvailabilityResponseInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestInput (non serializable)

type BlockSyncRequestInput struct {
	RecipientPublicKey primitives.Ed25519PublicKey
	Message            *gossipmessages.BlockSyncRequestMessage
}

func (x *BlockSyncRequestInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientPublicKey:%s,Message:%s,}", x.StringRecipientPublicKey(), x.StringMessage())
}

func (x *BlockSyncRequestInput) StringRecipientPublicKey() (res string) {
	res = fmt.Sprintf("%s", x.RecipientPublicKey)
	return
}

func (x *BlockSyncRequestInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseInput (non serializable)

type BlockSyncResponseInput struct {
	RecipientPublicKey primitives.Ed25519PublicKey
	Message            *gossipmessages.BlockSyncResponseMessage
}

func (x *BlockSyncResponseInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientPublicKey:%s,Message:%s,}", x.StringRecipientPublicKey(), x.StringMessage())
}

func (x *BlockSyncResponseInput) StringRecipientPublicKey() (res string) {
	res = fmt.Sprintf("%s", x.RecipientPublicKey)
	return
}

func (x *BlockSyncResponseInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
