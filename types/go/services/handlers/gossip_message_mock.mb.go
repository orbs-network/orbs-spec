// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package handlers

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionRelayGossipHandler

type MockTransactionRelayGossipHandler struct {
	mock.Mock
}

func (s *MockTransactionRelayGossipHandler) HandleForwardedTransactions(input *HandleForwardedTransactionsInput) (*GossipMessageHandlerOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GossipMessageHandlerOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

/////////////////////////////////////////////////////////////////////////////
// service BlockSyncGossipHandler

type MockBlockSyncGossipHandler struct {
	mock.Mock
}

func (s *MockBlockSyncGossipHandler) HandleBlockAvailabilityRequest(input *HandleBlockAvailabilityRequestInput) (*GossipMessageHandlerOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GossipMessageHandlerOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSyncGossipHandler) HandleBlockAvailabilityResponse(input *HandleBlockAvailabilityResponseInput) (*GossipMessageHandlerOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GossipMessageHandlerOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSyncGossipHandler) HandleBlockSyncRequest(input *HandleBlockSyncRequestInput) (*GossipMessageHandlerOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GossipMessageHandlerOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockSyncGossipHandler) HandleBlockSyncResponse(input *HandleBlockSyncResponseInput) (*GossipMessageHandlerOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GossipMessageHandlerOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

/////////////////////////////////////////////////////////////////////////////
// service LeanHelixConsensusGossipHandler

type MockLeanHelixConsensusGossipHandler struct {
	mock.Mock
}

func (s *MockLeanHelixConsensusGossipHandler) HandleLeanHelixPrePrepare(input *HandleLeanHelixPrePrepareInput) (*GossipMessageHandlerOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GossipMessageHandlerOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensusGossipHandler) HandleLeanHelixPrepare(input *HandleLeanHelixPrepareInput) (*GossipMessageHandlerOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GossipMessageHandlerOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensusGossipHandler) HandleLeanHelixCommit(input *HandleLeanHelixCommitInput) (*GossipMessageHandlerOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GossipMessageHandlerOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensusGossipHandler) HandleLeanHelixViewChange(input *HandleLeanHelixViewChangeInput) (*GossipMessageHandlerOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GossipMessageHandlerOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockLeanHelixConsensusGossipHandler) HandleLeanHelixNewView(input *HandleLeanHelixNewViewInput) (*GossipMessageHandlerOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GossipMessageHandlerOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

