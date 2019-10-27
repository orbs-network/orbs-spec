// AUTO GENERATED FILE (by membufc proto compiler v0.0.32)
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
	RecipientNodeAddress primitives.NodeAddress
	Message              *gossipmessages.BlockAvailabilityResponseMessage
}

func (x *BlockAvailabilityResponseInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientNodeAddress:%s,Message:%s,}", x.StringRecipientNodeAddress(), x.StringMessage())
}

func (x *BlockAvailabilityResponseInput) StringRecipientNodeAddress() (res string) {
	res = fmt.Sprintf("%s", x.RecipientNodeAddress)
	return
}

func (x *BlockAvailabilityResponseInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncRequestInput (non serializable)

type BlockSyncRequestInput struct {
	RecipientNodeAddress primitives.NodeAddress
	Message              *gossipmessages.BlockSyncRequestMessage
}

func (x *BlockSyncRequestInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientNodeAddress:%s,Message:%s,}", x.StringRecipientNodeAddress(), x.StringMessage())
}

func (x *BlockSyncRequestInput) StringRecipientNodeAddress() (res string) {
	res = fmt.Sprintf("%s", x.RecipientNodeAddress)
	return
}

func (x *BlockSyncRequestInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message BlockSyncResponseInput (non serializable)

type BlockSyncResponseInput struct {
	RecipientNodeAddress primitives.NodeAddress
	Message              *gossipmessages.BlockSyncResponseMessage
}

func (x *BlockSyncResponseInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientNodeAddress:%s,Message:%s,}", x.StringRecipientNodeAddress(), x.StringMessage())
}

func (x *BlockSyncResponseInput) StringRecipientNodeAddress() (res string) {
	res = fmt.Sprintf("%s", x.RecipientNodeAddress)
	return
}

func (x *BlockSyncResponseInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
