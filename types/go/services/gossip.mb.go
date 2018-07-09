// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/services/gossip"
)

/////////////////////////////////////////////////////////////////////////////
// service Gossip

type Gossip interface {
	gossip.TransactionRelay
	gossip.BlockSync
	gossip.LeanHelixConsensus
}

