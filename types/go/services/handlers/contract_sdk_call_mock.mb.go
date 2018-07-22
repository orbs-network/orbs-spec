// AUTO GENERATED FILE (by membufc proto compiler v0.0.17)
package handlers

import (
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service ContractSdkCallHandler

type MockContractSdkCallHandler struct {
	mock.Mock
}

func (s *MockContractSdkCallHandler) HandleSdkCall(input *HandleSdkCallInput) (*HandleSdkCallOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*HandleSdkCallOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

