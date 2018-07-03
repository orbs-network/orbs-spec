// AUTO GENERATED FILE (by membufc proto compiler)
package services

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service Gossip

type MockGossip struct {
	mock.Mock
}

func (s *MockGossip) TopicSubscribe(input *TopicSubscribeInput) (*TopicSubscribeOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*TopicSubscribeOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockGossip) TopicUnsubscribe(input *TopicUnsubscribeInput) (*TopicUnsubscribeOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*TopicUnsubscribeOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockGossip) SendMessage(input *SendMessageInput) (*SendMessageOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SendMessageOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

