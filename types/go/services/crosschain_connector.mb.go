// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
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
	InputArgument []*protocol.MethodArgument
	EthereumBlockHeight uint64
	EthereumBlockHeightMode protocol.BlockHeightMode
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractOutput (non serializable)

type EthereumCallContractOutput struct {
	OutputArgument []*protocol.MethodArgument
	CallStatus protocol.CallMethodStatus
	EthereumBlockHeight uint64
	EthereumBlockTimestamp uint64
}

/////////////////////////////////////////////////////////////////////////////
// enums

