// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package gossip

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service LeanHelixConsensus

type MockLeanHelixConsensus struct {
	mock.Mock
}

func (s *MockLeanHelixConsensus) SendLeanHelixPrePrepare(input *LeanHelixPrePrepareInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensus) SendLeanHelixPrepare(input *LeanHelixPrepareInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensus) SendLeanHelixCommit(input *LeanHelixCommitInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensus) SendLeanHelixViewChange(input *LeanHelixViewChangeInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensus) SendLeanHelixNewView(input *LeanHelixNewViewInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensus) RegisterLeanHelixConsensusHandler(handler LeanHelixConsensusHandler) {
	s.Called(handler)
}

/////////////////////////////////////////////////////////////////////////////
// service LeanHelixConsensusHandler

type MockLeanHelixConsensusHandler struct {
	mock.Mock
}

func (s *MockLeanHelixConsensusHandler) HandleLeanHelixPrePrepare(input *LeanHelixPrePrepareInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensusHandler) HandleLeanHelixPrepare(input *LeanHelixPrepareInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensusHandler) HandleLeanHelixCommit(input *LeanHelixCommitInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensusHandler) HandleLeanHelixViewChange(input *LeanHelixViewChangeInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensusHandler) HandleLeanHelixNewView(input *LeanHelixNewViewInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

