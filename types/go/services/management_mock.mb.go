// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service Management

type MockManagement struct {
	mock.Mock
}

func (s *MockManagement) GetCurrentReference(ctx context.Context, input *GetCurrentReferenceInput) (*GetCurrentReferenceOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetCurrentReferenceOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockManagement) GetGenesisReference(ctx context.Context, input *GetGenesisReferenceInput) (*GetGenesisReferenceOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetGenesisReferenceOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockManagement) GetProtocolVersion(ctx context.Context, input *GetProtocolVersionInput) (*GetProtocolVersionOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetProtocolVersionOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockManagement) GetSubscriptionStatus(ctx context.Context, input *GetSubscriptionStatusInput) (*GetSubscriptionStatusOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetSubscriptionStatusOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockManagement) GetCommittee(ctx context.Context, input *GetCommitteeInput) (*GetCommitteeOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetCommitteeOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
