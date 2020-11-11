// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service StateStorage

type MockStateStorage struct {
	mock.Mock
}

func (s *MockStateStorage) CommitStateDiff(ctx context.Context, input *CommitStateDiffInput) (*CommitStateDiffOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*CommitStateDiffOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockStateStorage) ReadKeys(ctx context.Context, input *ReadKeysInput) (*ReadKeysOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*ReadKeysOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockStateStorage) GetLastCommittedBlockInfo(ctx context.Context, input *GetLastCommittedBlockInfoInput) (*GetLastCommittedBlockInfoOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetLastCommittedBlockInfoOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockStateStorage) GetStateHash(ctx context.Context, input *GetStateHashInput) (*GetStateHashOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetStateHashOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
