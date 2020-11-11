// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service Vault

type MockVault struct {
	mock.Mock
}

func (s *MockVault) NodeSign(ctx context.Context, input *NodeSignInput) (*NodeSignOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*NodeSignOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockVault) EthSign(ctx context.Context, input *NodeSignInput) (*NodeSignOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*NodeSignOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
