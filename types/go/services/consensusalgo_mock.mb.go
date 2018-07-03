// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusAlgo

type MockConsensusAlgo struct {
	mock.Mock
}

func (s *MockConsensusAlgo) OnNewConsensusRound(input *OnNewConsensusRoundInput) (*OnNewConsensusRoundOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*OnNewConsensusRoundOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusAlgo) GossipMessageReceived(input *services.GossipMessageReceivedInput) (*services.GossipMessageReceivedOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*services.GossipMessageReceivedOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusAlgo) OnPrePrepareReceived(input *OnPrePrepareReceivedInput) (*OnPrePrepareReceivedOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*OnPrePrepareReceivedOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusAlgo) OnPrepareReceived(input *OnPrepareReceivedInput) (*OnPrepareReceivedOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*OnPrepareReceivedOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusAlgo) OnCommitReceived(input *OnCommitReceivedInput) (*OnCommitReceivedOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*OnCommitReceivedOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusAlgo) OnViewChangeReceived(input *OnViewChangeReceivedInput) (*OnViewChangeReceivedOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*OnViewChangeReceivedOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusAlgo) OnNewViewReceived(input *OnNewViewReceivedInput) (*OnNewViewReceivedOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*OnNewViewReceivedOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusAlgo) AcknowledgeTransactionsBlockConsensus(input *AcknowledgeTransactionsBlockConsensusInput) (*AcknowledgeTransactionsBlockConsensusOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*AcknowledgeTransactionsBlockConsensusOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusAlgo) AcknowledgeResultsBlockConsensus(input *AcknowledgeResultsBlockConsensusInput) (*AcknowledgeResultsBlockConsensusOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*AcknowledgeResultsBlockConsensusOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

