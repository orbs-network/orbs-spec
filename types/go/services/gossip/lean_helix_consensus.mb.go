// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package gossip

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/messages"
)

/////////////////////////////////////////////////////////////////////////////
// service LeanHelixConsensus

type LeanHelixConsensus interface {
	SendLeanHelixPrePrepare(input *LeanHelixPrePrepareInput) (*LeanHelixOutput, error)
	SendLeanHelixPrepare(input *LeanHelixPrepareInput) (*LeanHelixOutput, error)
	SendLeanHelixCommit(input *LeanHelixCommitInput) (*LeanHelixOutput, error)
	SendLeanHelixViewChange(input *LeanHelixViewChangeInput) (*LeanHelixOutput, error)
	SendLeanHelixNewView(input *LeanHelixNewViewInput) (*LeanHelixOutput, error)
	RegisterLeanHelixConsensusHandler(handler LeanHelixConsensusHandler)
}

/////////////////////////////////////////////////////////////////////////////
// service LeanHelixConsensusHandler

type LeanHelixConsensusHandler interface {
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
	Header *messages.LeanHelixPrePrepareHeader
	BlockRef *messages.LeanHelixBlockRef
	Block []byte
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixPrepareInput (non serializable)

type LeanHelixPrepareInput struct {
	RecipientPublicKeys []primitives.Ed25519Pkey
	Header *messages.LeanHelixPrepareHeader
	BlockRef *messages.LeanHelixBlockRef
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixCommitInput (non serializable)

type LeanHelixCommitInput struct {
	RecipientPublicKeys []primitives.Ed25519Pkey
	RecipientMode protocol.RecipientsListMode
	BlockRef *messages.LeanHelixBlockRef
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixViewChangeInput (non serializable)

type LeanHelixViewChangeInput struct {
	RecipientPublicKeys []primitives.Ed25519Pkey
	Header *messages.LeanHelixViewChangeHeader
	ViewChange *messages.LeanHelixViewChange
	Block []byte
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixNewViewInput (non serializable)

type LeanHelixNewViewInput struct {
	RecipientPublicKeys []primitives.Ed25519Pkey
	Header *messages.LeanHelixNewViewHeader
	NewView *messages.LeanHelixNewView
	Block []byte
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixOutput (non serializable)

type LeanHelixOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// enums

