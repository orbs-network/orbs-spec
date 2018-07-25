// AUTO GENERATED FILE (by membufc proto compiler v0.0.18)
package gossiptopics

import (
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// message RecipientsList (non serializable)

type RecipientsList struct {
	RecipientPublicKeys []primitives.Ed25519Pkey
	RecipientMode gossipmessages.RecipientsListMode
}

func (x *RecipientsList) String() string {
	return fmt.Sprintf("{RecipientPublicKeys:%s,RecipientMode:%s,}", x.StringRecipientPublicKeys(), x.StringRecipientMode())
}

func (x *RecipientsList) StringRecipientPublicKeys() (res string) {
	res = "["
		for _, v := range x.RecipientPublicKeys {
		res += fmt.Sprintf("%x", v) + ","
  }
	res += "]"
	return
}

func (x *RecipientsList) StringRecipientMode() (res string) {
	res = fmt.Sprintf("%x", x.RecipientMode)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EmptyOutput (non serializable)

type EmptyOutput struct {
}

func (x *EmptyOutput) String() string {
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// enums

