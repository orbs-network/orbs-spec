// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service Indexer

type MockIndexer struct {
	mock.Mock
}

func (s *MockIndexer) GetEvents(ctx context.Context, input *GetEventsInput) (*GetEventsOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*GetEventsOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
