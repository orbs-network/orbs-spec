// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package services

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service CrosschainConnector

type CrosschainConnector interface {
	EthereumCallContract(input *EthereumCallContractInput) (*EthereumCallContractOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractInput (non serializable)

type EthereumCallContractInput struct {
	EthereumContractAddress string
	EthereumFunctionCanonicalForm string
	InputArguments []*protocol.MethodArgument
	EthereumBlockHeight uint64
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractOutput (non serializable)

type EthereumCallContractOutput struct {
	OutputArguments []*protocol.MethodArgument
	CallStatus protocol.ExecutionStatus
	EthereumBlockHeight uint64
	EthereumBlockTimestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// enums

