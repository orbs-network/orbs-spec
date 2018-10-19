// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package services

import (
	"github.com/orbs-network/go-mock"
	"context"
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service BlockStorage

type MockBlockStorage struct {
	mock.Mock
	gossiptopics.MockBlockSyncHandler
}

func (s *MockBlockStorage) CommitBlock(ctx context.Context, input *CommitBlockInput) (*CommitBlockOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*CommitBlockOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) GetTransactionsBlockHeader(ctx context.Context, input *GetTransactionsBlockHeaderInput) (*GetTransactionsBlockHeaderOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetTransactionsBlockHeaderOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) GetResultsBlockHeader(ctx context.Context, input *GetResultsBlockHeaderInput) (*GetResultsBlockHeaderOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetResultsBlockHeaderOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) GetTransactionReceipt(ctx context.Context, input *GetTransactionReceiptInput) (*GetTransactionReceiptOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetTransactionReceiptOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) GetLastCommittedBlockHeight(ctx context.Context, input *GetLastCommittedBlockHeightInput) (*GetLastCommittedBlockHeightOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetLastCommittedBlockHeightOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) ValidateBlockForCommit(ctx context.Context, input *ValidateBlockForCommitInput) (*ValidateBlockForCommitOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*ValidateBlockForCommitOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockBlockStorage) RegisterConsensusBlocksHandler(handler handlers.ConsensusBlocksHandler) {
	s.Called(handler)
}

