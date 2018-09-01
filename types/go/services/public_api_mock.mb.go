// AUTO GENERATED FILE (by membufc proto compiler v0.0.19)
package services

import (
	"github.com/orbs-network/go-mock"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service PublicApi

type MockPublicApi struct {
	mock.Mock
	handlers.MockTransactionResultsHandler
}

func (s *MockPublicApi) SendTransaction(input *SendTransactionInput) (*SendTransactionOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SendTransactionOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockPublicApi) CallMethod(input *CallMethodInput) (*CallMethodOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*CallMethodOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockPublicApi) GetTransactionStatus(input *GetTransactionStatusInput) (*GetTransactionStatusOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*GetTransactionStatusOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

