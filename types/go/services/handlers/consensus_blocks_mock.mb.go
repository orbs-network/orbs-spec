// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package handlers

import (
	"github.com/orbs-network/go-mock"
	"context"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusBlocksHandler

type MockConsensusBlocksHandler struct {
	mock.Mock
}

func (s *MockConsensusBlocksHandler) HandleBlockConsensus(ctx context.Context, input *HandleBlockConsensusInput) (*HandleBlockConsensusOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*HandleBlockConsensusOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

