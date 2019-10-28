// AUTO GENERATED FILE (by membufc proto compiler v0.0.32)
package gossiptopics

import (
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol/gossipmessages"
)

/////////////////////////////////////////////////////////////////////////////
// message RecipientsList (non serializable)

type RecipientsList struct {
	RecipientNodeAddresses []primitives.NodeAddress
	RecipientMode          gossipmessages.RecipientsListMode
}

func (x *RecipientsList) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{RecipientNodeAddresses:%s,RecipientMode:%s,}", x.StringRecipientNodeAddresses(), x.StringRecipientMode())
}

func (x *RecipientsList) StringRecipientNodeAddresses() (res string) {
	res = "["
	for _, v := range x.RecipientNodeAddresses {
		res += fmt.Sprintf("%s", v) + ","
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
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// enums
