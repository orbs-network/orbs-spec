// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// service Management

type Management interface {
	GetCurrentReference(ctx context.Context, input *GetCurrentReferenceInput) (*GetCurrentReferenceOutput, error)
	GetGenesisReference(ctx context.Context, input *GetGenesisReferenceInput) (*GetGenesisReferenceOutput, error)
	GetProtocolVersion(ctx context.Context, input *GetProtocolVersionInput) (*GetProtocolVersionOutput, error)
	GetSubscriptionStatus(ctx context.Context, input *GetSubscriptionStatusInput) (*GetSubscriptionStatusOutput, error)
	GetCommittee(ctx context.Context, input *GetCommitteeInput) (*GetCommitteeOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message GetCurrentReferenceInput (non serializable)

type GetCurrentReferenceInput struct {
}

func (x *GetCurrentReferenceInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message GetCurrentReferenceOutput (non serializable)

type GetCurrentReferenceOutput struct {
	CurrentReference primitives.TimestampSeconds
}

func (x *GetCurrentReferenceOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentReference:%s,}", x.StringCurrentReference())
}

func (x *GetCurrentReferenceOutput) StringCurrentReference() (res string) {
	res = fmt.Sprintf("%s", x.CurrentReference)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetGenesisReferenceInput (non serializable)

type GetGenesisReferenceInput struct {
}

func (x *GetGenesisReferenceInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message GetGenesisReferenceOutput (non serializable)

type GetGenesisReferenceOutput struct {
	GenesisReference primitives.TimestampSeconds
	CurrentReference primitives.TimestampSeconds
}

func (x *GetGenesisReferenceOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{GenesisReference:%s,CurrentReference:%s,}", x.StringGenesisReference(), x.StringCurrentReference())
}

func (x *GetGenesisReferenceOutput) StringGenesisReference() (res string) {
	res = fmt.Sprintf("%s", x.GenesisReference)
	return
}

func (x *GetGenesisReferenceOutput) StringCurrentReference() (res string) {
	res = fmt.Sprintf("%s", x.CurrentReference)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetProtocolVersionInput (non serializable)

type GetProtocolVersionInput struct {
	Reference primitives.TimestampSeconds
}

func (x *GetProtocolVersionInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Reference:%s,}", x.StringReference())
}

func (x *GetProtocolVersionInput) StringReference() (res string) {
	res = fmt.Sprintf("%s", x.Reference)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetProtocolVersionOutput (non serializable)

type GetProtocolVersionOutput struct {
	ProtocolVersion primitives.ProtocolVersion
}

func (x *GetProtocolVersionOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ProtocolVersion:%s,}", x.StringProtocolVersion())
}

func (x *GetProtocolVersionOutput) StringProtocolVersion() (res string) {
	res = fmt.Sprintf("%s", x.ProtocolVersion)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetSubscriptionStatusInput (non serializable)

type GetSubscriptionStatusInput struct {
	Reference primitives.TimestampSeconds
}

func (x *GetSubscriptionStatusInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Reference:%s,}", x.StringReference())
}

func (x *GetSubscriptionStatusInput) StringReference() (res string) {
	res = fmt.Sprintf("%s", x.Reference)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetSubscriptionStatusOutput (non serializable)

type GetSubscriptionStatusOutput struct {
	SubscriptionStatusIsActive bool
}

func (x *GetSubscriptionStatusOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SubscriptionStatusIsActive:%s,}", x.StringSubscriptionStatusIsActive())
}

func (x *GetSubscriptionStatusOutput) StringSubscriptionStatusIsActive() (res string) {
	res = fmt.Sprintf("%v", x.SubscriptionStatusIsActive)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetCommitteeInput (non serializable)

type GetCommitteeInput struct {
	Reference primitives.TimestampSeconds
}

func (x *GetCommitteeInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Reference:%s,}", x.StringReference())
}

func (x *GetCommitteeInput) StringReference() (res string) {
	res = fmt.Sprintf("%s", x.Reference)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetCommitteeOutput (non serializable)

type GetCommitteeOutput struct {
	Committee []primitives.NodeAddress
}

func (x *GetCommitteeOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Committee:%s,}", x.StringCommittee())
}

func (x *GetCommitteeOutput) StringCommittee() (res string) {
	res = "["
	for _, v := range x.Committee {
		res += fmt.Sprintf("%s", v) + ","
	}
	res += "]"
	return
}
