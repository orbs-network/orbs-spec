// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package gossiptopics

import (
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockSync

type MockBlockSync struct {
	mock.Mock
}

func (s *MockBlockSync) BroadcastBlockAvailabilityRequest(input *BlockAvailabilityRequestInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSync) SendBlockAvailabilityResponse(input *BlockAvailabilityResponseInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSync) SendBlockSyncRequest(input *BlockSyncRequestInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSync) SendBlockSyncResponse(input *BlockSyncResponseInput) (*EmptyOutput, error) {
	ret := s.Called(input)
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

func (s *MockBlockSyncHandler) HandleBlockAvailabilityRequest(input *BlockAvailabilityRequestInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSyncHandler) HandleBlockAvailabilityResponse(input *BlockAvailabilityResponseInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSyncHandler) HandleBlockSyncRequest(input *BlockSyncRequestInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSyncHandler) HandleBlockSyncResponse(input *BlockSyncResponseInput) (*EmptyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*EmptyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

