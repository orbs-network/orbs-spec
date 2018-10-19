// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package gossiptopics

import (
	"github.com/orbs-network/go-mock"
	"context"
)

/////////////////////////////////////////////////////////////////////////////
// service LeanHelix

type MockLeanHelix struct {
	mock.Mock
}

func (s *MockLeanHelix) SendLeanHelixPrePrepare(ctx context.Context, input *LeanHelixPrePrepareInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) SendLeanHelixPrepare(ctx context.Context, input *LeanHelixPrepareInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) SendLeanHelixCommit(ctx context.Context, input *LeanHelixCommitInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) SendLeanHelixViewChange(ctx context.Context, input *LeanHelixViewChangeInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelix) SendLeanHelixNewView(ctx context.Context, input *LeanHelixNewViewInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
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

func (s *MockLeanHelixHandler) HandleLeanHelixPrePrepare(ctx context.Context, input *LeanHelixPrePrepareInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixHandler) HandleLeanHelixPrepare(ctx context.Context, input *LeanHelixPrepareInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixHandler) HandleLeanHelixCommit(ctx context.Context, input *LeanHelixCommitInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixHandler) HandleLeanHelixViewChange(ctx context.Context, input *LeanHelixViewChangeInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixHandler) HandleLeanHelixNewView(ctx context.Context, input *LeanHelixNewViewInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

