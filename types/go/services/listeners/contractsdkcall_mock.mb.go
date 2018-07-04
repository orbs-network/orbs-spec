// AUTO GENERATED FILE (by membufc proto compiler)
package listeners

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service ContractSdkCallListener

type MockContractSdkCallListener struct {
	mock.Mock
}

func (s *MockContractSdkCallListener) SdkCall(input *SdkCallInput) (*SdkCallOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SdkCallOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

