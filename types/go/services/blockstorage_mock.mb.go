// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockStorage

type MockBlockStorage struct {
	mock.Mock
}

func (s *MockBlockStorage) CommitBlock(input *CommitBlockInput) (*CommitBlockOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*CommitBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) GetTransactionsBlockHeader(input *GetTransactionsBlockHeaderInput) (*GetTransactionsBlockHeaderOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GetTransactionsBlockHeaderOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) GetResultsBlockHeader(input *GetResultsBlockHeaderInput) (*GetResultsBlockHeaderOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GetResultsBlockHeaderOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) GetTransactionReceipt(input *GetTransactionReceiptInput) (*GetTransactionReceiptOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GetTransactionReceiptOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) GetLastCommittedBlockHeight(input *GetLastCommittedBlockHeightInput) (*GetLastCommittedBlockHeightOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GetLastCommittedBlockHeightOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) GossipMessageReceived(input *GossipMessageReceivedInput) (*GossipMessageReceivedOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GossipMessageReceivedOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) OnBlockAvailabilityRequest(input *OnBlockAvailabilityRequestInput) (*OnBlockAvailabilityRequestOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*OnBlockAvailabilityRequestOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) OnBlockAvailabilityResponse(input *OnBlockAvailabilityResponseInput) (*OnBlockAvailabilityResponseOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*OnBlockAvailabilityResponseOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) OnBlockSyncRequest(input *OnBlockSyncRequestInput) (*OnBlockSyncRequestOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*OnBlockSyncRequestOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) OnBlockSyncResponse(input *OnBlockSyncResponseInput) (*OnBlockSyncResponseOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*OnBlockSyncResponseOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) ValidateBlockForCommit(input *ValidateBlockForCommitInput) (*ValidateBlockForCommitOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ValidateBlockForCommitOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

