// AUTO GENERATED FILE (by membufc proto compiler v0.0.14)
package gossiptopics

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service LeanHelix

type MockLeanHelix struct {
	mock.Mock
}

func (s *MockLeanHelix) SendLeanHelixPrePrepare(input *LeanHelixPrePrepareInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) SendLeanHelixPrepare(input *LeanHelixPrepareInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) SendLeanHelixCommit(input *LeanHelixCommitInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) SendLeanHelixViewChange(input *LeanHelixViewChangeInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) SendLeanHelixNewView(input *LeanHelixNewViewInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
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

func (s *MockLeanHelixHandler) HandleLeanHelixPrePrepare(input *LeanHelixPrePrepareInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixHandler) HandleLeanHelixPrepare(input *LeanHelixPrepareInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixHandler) HandleLeanHelixCommit(input *LeanHelixCommitInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixHandler) HandleLeanHelixViewChange(input *LeanHelixViewChangeInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixHandler) HandleLeanHelixNewView(input *LeanHelixNewViewInput) (*LeanHelixOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*LeanHelixOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

