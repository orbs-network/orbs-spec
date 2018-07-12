// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package services

import (
	"github.com/maraino/go-mock"
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockStorage

type MockBlockStorage struct {
	mock.Mock
	gossiptopics.MockBlockSyncHandler
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

func (s *MockBlockStorage) ValidateBlockForCommit(input *ValidateBlockForCommitInput) (*ValidateBlockForCommitOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ValidateBlockForCommitOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) RegisterConsensusBlocksHandler(handler handlers.ConsensusBlocksHandler) {
	s.Called(handler)
}

