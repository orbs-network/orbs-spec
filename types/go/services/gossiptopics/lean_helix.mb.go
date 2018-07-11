// AUTO GENERATED FILE (by membufc proto compiler v0.0.14)
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
	RecipientPublicKeys []primitives.Ed25519Pkey
	Header *gossipmessages.LeanHelixPrePrepareHeader
	BlockRef *gossipmessages.LeanHelixBlockRef
	Block []byte
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrepareInput (non serializable)

type LeanHelixPrepareInput struct {
	RecipientPublicKeys []primitives.Ed25519Pkey
	Header *gossipmessages.LeanHelixPrepareHeader
	BlockRef *gossipmessages.LeanHelixBlockRef
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixCommitInput (non serializable)

type LeanHelixCommitInput struct {
	RecipientPublicKeys []primitives.Ed25519Pkey
	RecipientMode gossipmessages.RecipientsListMode
	BlockRef *gossipmessages.LeanHelixBlockRef
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixViewChangeInput (non serializable)

type LeanHelixViewChangeInput struct {
	RecipientPublicKeys []primitives.Ed25519Pkey
	Header *gossipmessages.LeanHelixViewChangeHeader
	ViewChange *gossipmessages.LeanHelixViewChange
	Block []byte
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewViewInput (non serializable)

type LeanHelixNewViewInput struct {
	RecipientPublicKeys []primitives.Ed25519Pkey
	Header *gossipmessages.LeanHelixNewViewHeader
	NewView *gossipmessages.LeanHelixNewView
	Block []byte
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixOutput (non serializable)

type LeanHelixOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// enums

