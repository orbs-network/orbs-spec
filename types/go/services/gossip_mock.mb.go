// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service Gossip

type MockGossip struct {
	mock.Mock
}

func (s *MockGossip) BroadcastForwardedTransactions(input *BroadcastForwardedTransactionsInput) (*SendGossipMessageOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SendGossipMessageOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockGossip) BroadcastBlockSyncAvailabilityRequest(input *BroadcastBlockSyncAvailabilityRequestInput) (*SendGossipMessageOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SendGossipMessageOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockGossip) SendBlockSyncAvailabilityResponse(input *SendBlockSyncAvailabilityResponseInput) (*SendGossipMessageOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SendGossipMessageOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockGossip) SendBlockSyncRequest(input *SendBlockSyncRequestInput) (*SendGossipMessageOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SendGossipMessageOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockGossip) SendBlockSyncResponse(input *SendBlockSyncResponseInput) (*SendGossipMessageOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SendGossipMessageOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockGossip) SendLeanHelixPrePrepare(input *SendLeanHelixPrePrepareInput) (*SendGossipMessageOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SendGossipMessageOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockGossip) SendLeanHelixPrepare(input *SendLeanHelixPrepareInput) (*SendGossipMessageOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SendGossipMessageOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockGossip) SendLeanHelixCommit(input *SendLeanHelixCommitInput) (*SendGossipMessageOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SendGossipMessageOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockGossip) SendLeanHelixViewChange(input *SendLeanHelixViewChangeInput) (*SendGossipMessageOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SendGossipMessageOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockGossip) SendLeanHelixNewView(input *SendLeanHelixNewViewInput) (*SendGossipMessageOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SendGossipMessageOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

