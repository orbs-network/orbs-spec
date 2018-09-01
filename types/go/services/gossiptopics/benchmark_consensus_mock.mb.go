// AUTO GENERATED FILE (by membufc proto compiler v0.0.19)
package gossiptopics

import (
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service BenchmarkConsensus

type MockBenchmarkConsensus struct {
	mock.Mock
}

func (s *MockBenchmarkConsensus) BroadcastBenchmarkConsensusCommit(input *BenchmarkConsensusCommitInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBenchmarkConsensus) SendBenchmarkConsensusCommitted(input *BenchmarkConsensusCommittedInput) (*EmptyOutput, error) {
	ret := s.Called(input)
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

func (s *MockBenchmarkConsensusHandler) HandleBenchmarkConsensusCommit(input *BenchmarkConsensusCommitInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBenchmarkConsensusHandler) HandleBenchmarkConsensusCommitted(input *BenchmarkConsensusCommittedInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

