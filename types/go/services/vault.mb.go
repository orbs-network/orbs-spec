// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// service Vault

type Vault interface {
	NodeSign(ctx context.Context, input *NodeSignInput) (*NodeSignOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message NodeSignInput (non serializable)

type NodeSignInput struct {
	Data []byte
}

func (x *NodeSignInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Data:%s,}", x.StringData())
}

func (x *NodeSignInput) StringData() (res string) {
	res = fmt.Sprintf("%x", x.Data)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message NodeSignOutput (non serializable)

type NodeSignOutput struct {
	Signature primitives.EcdsaSecp256K1Sig
}

func (x *NodeSignOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Signature:%s,}", x.StringSignature())
}

func (x *NodeSignOutput) StringSignature() (res string) {
	res = fmt.Sprintf("%s", x.Signature)
	return
}
