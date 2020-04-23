// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package gossiptopics

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service HeaderSync

type HeaderSync interface {
	BroadcastHeaderAvailabilityRequest(ctx context.Context, input *HeaderAvailabilityRequestInput) (*EmptyOutput, error)
	SendHeaderAvailabilityResponse(ctx context.Context, input *HeaderAvailabilityResponseInput) (*EmptyOutput, error)
	SendHeaderSyncRequest(ctx context.Context, input *HeaderSyncRequestInput) (*EmptyOutput, error)
	SendHeaderSyncResponse(ctx context.Context, input *HeaderSyncResponseInput) (*EmptyOutput, error)
	RegisterHeaderSyncHandler(handler HeaderSyncHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service HeaderSyncHandler

type HeaderSyncHandler interface {
	HandleHeaderAvailabilityRequest(ctx context.Context, input *HeaderAvailabilityRequestInput) (*EmptyOutput, error)
	HandleHeaderAvailabilityResponse(ctx context.Context, input *HeaderAvailabilityResponseInput) (*EmptyOutput, error)
	HandleHeaderSyncRequest(ctx context.Context, input *HeaderSyncRequestInput) (*EmptyOutput, error)
	HandleHeaderSyncResponse(ctx context.Context, input *HeaderSyncResponseInput) (*EmptyOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message HeaderAvailabilityRequestInput (non serializable)

type HeaderAvailabilityRequestInput struct {
	Message *gossipmessages.HeaderAvailabilityRequestMessage
}

func (x *HeaderAvailabilityRequestInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Message:%s,}", x.StringMessage())
}

func (x *HeaderAvailabilityRequestInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HeaderAvailabilityResponseInput (non serializable)

type HeaderAvailabilityResponseInput struct {
	RecipientNodeAddress primitives.NodeAddress
	Message              *gossipmessages.HeaderAvailabilityResponseMessage
}

func (x *HeaderAvailabilityResponseInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientNodeAddress:%s,Message:%s,}", x.StringRecipientNodeAddress(), x.StringMessage())
}

func (x *HeaderAvailabilityResponseInput) StringRecipientNodeAddress() (res string) {
	res = fmt.Sprintf("%s", x.RecipientNodeAddress)
	return
}

func (x *HeaderAvailabilityResponseInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HeaderSyncRequestInput (non serializable)

type HeaderSyncRequestInput struct {
	RecipientNodeAddress primitives.NodeAddress
	Message              *gossipmessages.HeaderSyncRequestMessage
}

func (x *HeaderSyncRequestInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientNodeAddress:%s,Message:%s,}", x.StringRecipientNodeAddress(), x.StringMessage())
}

func (x *HeaderSyncRequestInput) StringRecipientNodeAddress() (res string) {
	res = fmt.Sprintf("%s", x.RecipientNodeAddress)
	return
}

func (x *HeaderSyncRequestInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message HeaderSyncResponseInput (non serializable)

type HeaderSyncResponseInput struct {
	RecipientNodeAddress primitives.NodeAddress
	Message              *gossipmessages.HeaderSyncResponseMessage
}

func (x *HeaderSyncResponseInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientNodeAddress:%s,Message:%s,}", x.StringRecipientNodeAddress(), x.StringMessage())
}

func (x *HeaderSyncResponseInput) StringRecipientNodeAddress() (res string) {
	res = fmt.Sprintf("%s", x.RecipientNodeAddress)
	return
}

func (x *HeaderSyncResponseInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
