// AUTO GENERATED FILE (by membufc proto compiler v0.0.17)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
)

/////////////////////////////////////////////////////////////////////////////
// service Gossip

type Gossip interface {
	gossiptopics.TransactionRelay
	gossiptopics.BlockSync
	gossiptopics.LeanHelix
	gossiptopics.BenchmarkConsensus
}

