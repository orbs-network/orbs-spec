// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package services

import (
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service StateStorage

type MockStateStorage struct {
	mock.Mock
}

func (s *MockStateStorage) CommitStateDiff(input *CommitStateDiffInput) (*CommitStateDiffOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*CommitStateDiffOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockStateStorage) ReadKeys(input *ReadKeysInput) (*ReadKeysOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ReadKeysOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockStateStorage) GetStateStorageBlockHeight(input *GetStateStorageBlockHeightInput) (*GetStateStorageBlockHeightOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GetStateStorageBlockHeightOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockStateStorage) GetStateHash(input *GetStateHashInput) (*GetStateHashOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GetStateHashOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

