// AUTO GENERATED FILE (by membufc proto compiler v0.0.32)
package gossiptopics

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockSync

type MockBlockSync struct {
	mock.Mock
}

func (s *MockBlockSync) BroadcastBlockAvailabilityRequest(ctx context.Context, input *BlockAvailabilityRequestInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSync) SendBlockAvailabilityResponse(ctx context.Context, input *BlockAvailabilityResponseInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSync) SendBlockSyncRequest(ctx context.Context, input *BlockSyncRequestInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSync) SendBlockSyncResponse(ctx context.Context, input *BlockSyncResponseInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSync) RegisterBlockSyncHandler(handler BlockSyncHandler) {
	s.Called(handler)
}

/////////////////////////////////////////////////////////////////////////////
// service BlockSyncHandler

type MockBlockSyncHandler struct {
	mock.Mock
}

func (s *MockBlockSyncHandler) HandleBlockAvailabilityRequest(ctx context.Context, input *BlockAvailabilityRequestInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSyncHandler) HandleBlockAvailabilityResponse(ctx context.Context, input *BlockAvailabilityResponseInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSyncHandler) HandleBlockSyncRequest(ctx context.Context, input *BlockSyncRequestInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSyncHandler) HandleBlockSyncResponse(ctx context.Context, input *BlockSyncResponseInput) (*EmptyOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
