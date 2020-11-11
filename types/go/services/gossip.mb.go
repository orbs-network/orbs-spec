// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
)

/////////////////////////////////////////////////////////////////////////////
// service Gossip

type Gossip interface {
	gossiptopics.TransactionRelay
	gossiptopics.BlockSync
	gossiptopics.LeanHelix
	gossiptopics.BenchmarkConsensus
	UpdateTopology(ctx context.Context, input *UpdateTopologyInput) (*UpdateTopologyOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message UpdateTopologyInput (non serializable)

type UpdateTopologyInput struct {
	Peers []*GossipPeer
}

func (x *UpdateTopologyInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Peers:%s,}", x.StringPeers())
}

func (x *UpdateTopologyInput) StringPeers() (res string) {
	res = "["
	for _, v := range x.Peers {
		res += v.String() + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message UpdateTopologyOutput (non serializable)

type UpdateTopologyOutput struct {
}

func (x *UpdateTopologyOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message GossipPeer (non serializable)

type GossipPeer struct {
	Address  primitives.NodeAddress
	Endpoint string
	Port     uint32
}

func (x *GossipPeer) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Address:%s,Endpoint:%s,Port:%s,}", x.StringAddress(), x.StringEndpoint(), x.StringPort())
}

func (x *GossipPeer) StringAddress() (res string) {
	res = fmt.Sprintf("%s", x.Address)
	return
}

func (x *GossipPeer) StringEndpoint() (res string) {
	res = fmt.Sprintf("%s", x.Endpoint)
	return
}

func (x *GossipPeer) StringPort() (res string) {
	res = fmt.Sprintf("%v", x.Port)
	return
}
