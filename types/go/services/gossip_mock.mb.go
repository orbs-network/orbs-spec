// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package services

import (
	"github.com/maraino/go-mock"
	"github.com/orbs-network/orbs-spec/types/go/services/gossip"
)

/////////////////////////////////////////////////////////////////////////////
// service Gossip

type MockGossip struct {
	mock.Mock
	gossip.MockTransactionRelay
	gossip.MockBlockSync
	gossip.MockLeanHelixConsensus
}

