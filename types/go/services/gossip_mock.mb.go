// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"github.com/orbs-network/go-mock"
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
)

/////////////////////////////////////////////////////////////////////////////
// service Gossip

type MockGossip struct {
	mock.Mock
	gossiptopics.MockTransactionRelay
	gossiptopics.MockBlockSync
	gossiptopics.MockLeanHelix
	gossiptopics.MockBenchmarkConsensus
}

func (s *MockGossip) UpdateTopology(ctx context.Context, input *UpdateTopologyInput) (*UpdateTopologyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*UpdateTopologyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
