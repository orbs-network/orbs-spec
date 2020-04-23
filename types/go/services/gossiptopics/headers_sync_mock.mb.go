// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package gossiptopics

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service HeaderSync

type MockHeaderSync struct {
	mock.Mock
}

func (s *MockHeaderSync) BroadcastHeaderAvailabilityRequest(ctx context.Context, input *HeaderAvailabilityRequestInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockHeaderSync) SendHeaderAvailabilityResponse(ctx context.Context, input *HeaderAvailabilityResponseInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockHeaderSync) SendHeaderSyncRequest(ctx context.Context, input *HeaderSyncRequestInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockHeaderSync) SendHeaderSyncResponse(ctx context.Context, input *HeaderSyncResponseInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockHeaderSync) RegisterHeaderSyncHandler(handler HeaderSyncHandler) {
	s.Called(handler)
}

/////////////////////////////////////////////////////////////////////////////
// service HeaderSyncHandler

type MockHeaderSyncHandler struct {
	mock.Mock
}

func (s *MockHeaderSyncHandler) HandleHeaderAvailabilityRequest(ctx context.Context, input *HeaderAvailabilityRequestInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockHeaderSyncHandler) HandleHeaderAvailabilityResponse(ctx context.Context, input *HeaderAvailabilityResponseInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockHeaderSyncHandler) HandleHeaderSyncRequest(ctx context.Context, input *HeaderSyncRequestInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockHeaderSyncHandler) HandleHeaderSyncResponse(ctx context.Context, input *HeaderSyncResponseInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
