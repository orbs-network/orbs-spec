// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package gossiptopics

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockSync

type MockBlockSync struct {
	mock.Mock
}

func (s *MockBlockSync) BroadcastBlockAvailabilityRequest(input *BlockAvailabilityRequestInput) (*BlockSyncOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*BlockSyncOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSync) SendBlockAvailabilityResponse(input *BlockAvailabilityResponseInput) (*BlockSyncOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*BlockSyncOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSync) SendBlockRequest(input *BlockSyncRequestInput) (*BlockSyncOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*BlockSyncOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSync) SendBlockResponse(input *BlockSyncResponseInput) (*BlockSyncOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*BlockSyncOutput), ret.Error(1)
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

func (s *MockBlockSyncHandler) HandleBlockAvailabilityRequest(input *BlockAvailabilityRequestInput) (*BlockSyncOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*BlockSyncOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSyncHandler) HandleBlockAvailabilityResponse(input *BlockAvailabilityResponseInput) (*BlockSyncOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*BlockSyncOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSyncHandler) HandleBlockSyncRequest(input *BlockSyncRequestInput) (*BlockSyncOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*BlockSyncOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSyncHandler) HandleBlockSyncResponse(input *BlockSyncResponseInput) (*BlockSyncOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*BlockSyncOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

