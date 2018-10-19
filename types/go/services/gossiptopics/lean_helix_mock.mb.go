// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package gossiptopics

import (
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service LeanHelix

type MockLeanHelix struct {
	mock.Mock
}

func (s *MockLeanHelix) SendLeanHelixPrePrepare(input *LeanHelixPrePrepareInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) SendLeanHelixPrepare(input *LeanHelixPrepareInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) SendLeanHelixCommit(input *LeanHelixCommitInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) SendLeanHelixViewChange(input *LeanHelixViewChangeInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) SendLeanHelixNewView(input *LeanHelixNewViewInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) RegisterLeanHelixHandler(handler LeanHelixHandler) {
	s.Called(handler)
}

/////////////////////////////////////////////////////////////////////////////
// service LeanHelixHandler

type MockLeanHelixHandler struct {
	mock.Mock
}

func (s *MockLeanHelixHandler) HandleLeanHelixPrePrepare(input *LeanHelixPrePrepareInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixHandler) HandleLeanHelixPrepare(input *LeanHelixPrepareInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixHandler) HandleLeanHelixCommit(input *LeanHelixCommitInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixHandler) HandleLeanHelixViewChange(input *LeanHelixViewChangeInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixHandler) HandleLeanHelixNewView(input *LeanHelixNewViewInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

