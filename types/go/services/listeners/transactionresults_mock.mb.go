// AUTO GENERATED FILE (by membufc proto compiler)
package listeners

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionResultsListener

type MockTransactionResultsListener struct {
	mock.Mock
}

func (s *MockTransactionResultsListener) ReturnTransactionResults(input *ReturnTransactionResultsInput) (*ReturnTransactionResultsOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ReturnTransactionResultsOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

