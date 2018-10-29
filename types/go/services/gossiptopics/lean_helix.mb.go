// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package gossiptopics

import (
	"fmt"
	"context"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service LeanHelix

type LeanHelix interface {
	SendLeanHelixPrePrepare(ctx context.Context, input *LeanHelixPrePrepareInput) (*EmptyOutput, error)
	SendLeanHelixPrepare(ctx context.Context, input *LeanHelixPrepareInput) (*EmptyOutput, error)
	SendLeanHelixCommit(ctx context.Context, input *LeanHelixCommitInput) (*EmptyOutput, error)
	SendLeanHelixViewChange(ctx context.Context, input *LeanHelixViewChangeInput) (*EmptyOutput, error)
	SendLeanHelixNewView(ctx context.Context, input *LeanHelixNewViewInput) (*EmptyOutput, error)
	RegisterLeanHelixHandler(handler LeanHelixHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service LeanHelixHandler

type LeanHelixHandler interface {
	HandleLeanHelixPrePrepare(ctx context.Context, input *LeanHelixPrePrepareInput) (*EmptyOutput, error)
	HandleLeanHelixPrepare(ctx context.Context, input *LeanHelixPrepareInput) (*EmptyOutput, error)
	HandleLeanHelixCommit(ctx context.Context, input *LeanHelixCommitInput) (*EmptyOutput, error)
	HandleLeanHelixViewChange(ctx context.Context, input *LeanHelixViewChangeInput) (*EmptyOutput, error)
	HandleLeanHelixNewView(ctx context.Context, input *LeanHelixNewViewInput) (*EmptyOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrePrepareInput (non serializable)

type LeanHelixPrePrepareInput struct {
	RecipientsList *RecipientsList
	Message *gossipmessages.LeanHelixPrePrepareMessage
}

func (x *LeanHelixPrePrepareInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientsList:%s,Message:%s,}", x.StringRecipientsList(), x.StringMessage())
}

func (x *LeanHelixPrePrepareInput) StringRecipientsList() (res string) {
	res = x.RecipientsList.String()
	return
}

func (x *LeanHelixPrePrepareInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrepareInput (non serializable)

type LeanHelixPrepareInput struct {
	RecipientsList *RecipientsList
	Message *gossipmessages.LeanHelixPrepareMessage
}

func (x *LeanHelixPrepareInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientsList:%s,Message:%s,}", x.StringRecipientsList(), x.StringMessage())
}

func (x *LeanHelixPrepareInput) StringRecipientsList() (res string) {
	res = x.RecipientsList.String()
	return
}

func (x *LeanHelixPrepareInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixCommitInput (non serializable)

type LeanHelixCommitInput struct {
	RecipientsList *RecipientsList
	Message *gossipmessages.LeanHelixCommitMessage
}

func (x *LeanHelixCommitInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientsList:%s,Message:%s,}", x.StringRecipientsList(), x.StringMessage())
}

func (x *LeanHelixCommitInput) StringRecipientsList() (res string) {
	res = x.RecipientsList.String()
	return
}

func (x *LeanHelixCommitInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixViewChangeInput (non serializable)

type LeanHelixViewChangeInput struct {
	RecipientsList *RecipientsList
	Message *gossipmessages.LeanHelixViewChangeMessage
}

func (x *LeanHelixViewChangeInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientsList:%s,Message:%s,}", x.StringRecipientsList(), x.StringMessage())
}

func (x *LeanHelixViewChangeInput) StringRecipientsList() (res string) {
	res = x.RecipientsList.String()
	return
}

func (x *LeanHelixViewChangeInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewViewInput (non serializable)

type LeanHelixNewViewInput struct {
	RecipientsList *RecipientsList
	Message *gossipmessages.LeanHelixNewViewMessage
}

func (x *LeanHelixNewViewInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientsList:%s,Message:%s,}", x.StringRecipientsList(), x.StringMessage())
}

func (x *LeanHelixNewViewInput) StringRecipientsList() (res string) {
	res = x.RecipientsList.String()
	return
}

func (x *LeanHelixNewViewInput) StringMessage() (res string) {
	res = x.Message.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums

