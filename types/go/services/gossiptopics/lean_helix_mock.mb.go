// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package gossiptopics

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service LeanHelix

type MockLeanHelix struct {
	mock.Mock
}

func (s *MockLeanHelix) SendLeanHelixMessage(ctx context.Context, input *LeanHelixInput) (*EmptyOutput, error) {
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

func (s *MockLeanHelixHandler) HandleLeanHelixMessage(ctx context.Context, input *LeanHelixInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
