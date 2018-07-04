// AUTO GENERATED FILE (by membufc proto compiler)
package listeners

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusBlocksListener

type MockConsensusBlocksListener struct {
	mock.Mock
}

func (s *MockConsensusBlocksListener) AcknowledgeTransactionsBlockConsensus(input *AcknowledgeTransactionsBlockConsensusInput) (*AcknowledgeTransactionsBlockConsensusOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*AcknowledgeTransactionsBlockConsensusOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockConsensusBlocksListener) AcknowledgeResultsBlockConsensus(input *AcknowledgeResultsBlockConsensusInput) (*AcknowledgeResultsBlockConsensusOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*AcknowledgeResultsBlockConsensusOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

