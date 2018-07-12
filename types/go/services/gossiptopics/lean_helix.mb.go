// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package gossiptopics

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// service LeanHelix

type LeanHelix interface {
	SendLeanHelixPrePrepare(input *LeanHelixPrePrepareInput) (*LeanHelixOutput, error)
	SendLeanHelixPrepare(input *LeanHelixPrepareInput) (*LeanHelixOutput, error)
	SendLeanHelixCommit(input *LeanHelixCommitInput) (*LeanHelixOutput, error)
	SendLeanHelixViewChange(input *LeanHelixViewChangeInput) (*LeanHelixOutput, error)
	SendLeanHelixNewView(input *LeanHelixNewViewInput) (*LeanHelixOutput, error)
	RegisterLeanHelixHandler(handler LeanHelixHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service LeanHelixHandler

type LeanHelixHandler interface {
	HandleLeanHelixPrePrepare(input *LeanHelixPrePrepareInput) (*LeanHelixOutput, error)
	HandleLeanHelixPrepare(input *LeanHelixPrepareInput) (*LeanHelixOutput, error)
	HandleLeanHelixCommit(input *LeanHelixCommitInput) (*LeanHelixOutput, error)
	HandleLeanHelixViewChange(input *LeanHelixViewChangeInput) (*LeanHelixOutput, error)
	HandleLeanHelixNewView(input *LeanHelixNewViewInput) (*LeanHelixOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrePrepareInput (non serializable)

type LeanHelixPrePrepareInput struct {
	RecipientList *gossipmessages.RecipientsList
	PrePrepareMessage *gossipmessages.LeanHelixPrePrepareMessage
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrepareInput (non serializable)

type LeanHelixPrepareInput struct {
	RecipientList *gossipmessages.RecipientsList
	PrepareMessage *gossipmessages.LeanHelixPrepareMessage
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixCommitInput (non serializable)

type LeanHelixCommitInput struct {
	RecipientList *gossipmessages.RecipientsList
	CommitMessage *gossipmessages.LeanHelixCommitMessage
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixViewChangeInput (non serializable)

type LeanHelixViewChangeInput struct {
	RecipientList *gossipmessages.RecipientsList
	ViewChange *gossipmessages.LeanHelixViewChangeMessage
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewViewInput (non serializable)

type LeanHelixNewViewInput struct {
	RecipientList *gossipmessages.RecipientsList
	NewView *gossipmessages.LeanHelixNewView
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixOutput (non serializable)

type LeanHelixOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// enums

