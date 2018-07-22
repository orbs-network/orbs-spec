// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package gossiptopics

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service LeanHelix

type LeanHelix interface {
	SendLeanHelixPrePrepare(input *LeanHelixPrePrepareInput) (*EmptyOutput, error)
	SendLeanHelixPrepare(input *LeanHelixPrepareInput) (*EmptyOutput, error)
	SendLeanHelixCommit(input *LeanHelixCommitInput) (*EmptyOutput, error)
	SendLeanHelixViewChange(input *LeanHelixViewChangeInput) (*EmptyOutput, error)
	SendLeanHelixNewView(input *LeanHelixNewViewInput) (*EmptyOutput, error)
	RegisterLeanHelixHandler(handler LeanHelixHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service LeanHelixHandler

type LeanHelixHandler interface {
	HandleLeanHelixPrePrepare(input *LeanHelixPrePrepareInput) (*EmptyOutput, error)
	HandleLeanHelixPrepare(input *LeanHelixPrepareInput) (*EmptyOutput, error)
	HandleLeanHelixCommit(input *LeanHelixCommitInput) (*EmptyOutput, error)
	HandleLeanHelixViewChange(input *LeanHelixViewChangeInput) (*EmptyOutput, error)
	HandleLeanHelixNewView(input *LeanHelixNewViewInput) (*EmptyOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrePrepareInput (non serializable)

type LeanHelixPrePrepareInput struct {
	RecipientsList *RecipientsList
	Message *gossipmessages.LeanHelixPrePrepareMessage
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrepareInput (non serializable)

type LeanHelixPrepareInput struct {
	RecipientsList *RecipientsList
	Message *gossipmessages.LeanHelixPrepareMessage
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixCommitInput (non serializable)

type LeanHelixCommitInput struct {
	RecipientsList *RecipientsList
	Message *gossipmessages.LeanHelixCommitMessage
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixViewChangeInput (non serializable)

type LeanHelixViewChangeInput struct {
	RecipientsList *RecipientsList
	Message *gossipmessages.LeanHelixViewChangeMessage
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewViewInput (non serializable)

type LeanHelixNewViewInput struct {
	RecipientsList *RecipientsList
	Message *gossipmessages.LeanHelixNewViewMessage
}

/////////////////////////////////////////////////////////////////////////////
// enums

