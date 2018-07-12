// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package gossiptopics

import (
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// message RecipientsList (non serializable)

type RecipientsList struct {
	RecipientPublicKeys []primitives.Ed25519Pkey
	RecipientMode gossipmessages.RecipientsListMode
}

/////////////////////////////////////////////////////////////////////////////
// message EmptyOutput (non serializable)

type EmptyOutput struct {
}

/////////////////////////////////////////////////////////////////////////////
// enums

