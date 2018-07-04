// AUTO GENERATED FILE (by membufc proto compiler)
package listeners

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service GossipMessageListener

type MockGossipMessageListener struct {
	mock.Mock
}

func (s *MockGossipMessageListener) GossipMessageReceived(input *GossipMessageReceivedInput) (*GossipMessageReceivedOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GossipMessageReceivedOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

