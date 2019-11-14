// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package gossiptopics

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service BenchmarkConsensus

type MockBenchmarkConsensus struct {
	mock.Mock
}

func (s *MockBenchmarkConsensus) BroadcastBenchmarkConsensusCommit(ctx context.Context, input *BenchmarkConsensusCommitInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBenchmarkConsensus) SendBenchmarkConsensusCommitted(ctx context.Context, input *BenchmarkConsensusCommittedInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBenchmarkConsensus) RegisterBenchmarkConsensusHandler(handler BenchmarkConsensusHandler) {
	s.Called(handler)
}

/////////////////////////////////////////////////////////////////////////////
// service BenchmarkConsensusHandler

type MockBenchmarkConsensusHandler struct {
	mock.Mock
}

func (s *MockBenchmarkConsensusHandler) HandleBenchmarkConsensusCommit(ctx context.Context, input *BenchmarkConsensusCommitInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBenchmarkConsensusHandler) HandleBenchmarkConsensusCommitted(ctx context.Context, input *BenchmarkConsensusCommittedInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
