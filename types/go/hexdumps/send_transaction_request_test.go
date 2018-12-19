package hexdumps

import (
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/protocol/client"
	"testing"
)

func TestSendTransactionRequest(t *testing.T) {
	builder := &client.SendTransactionRequestBuilder{
		SignedTransaction: &protocol.SignedTransactionBuilder{
			Transaction: &protocol.TransactionBuilder{
				ProtocolVersion: 1,
				VirtualChainId:  42,
				Timestamp:       0x1122334455,
				Signer: &protocol.SignerBuilder{
					Scheme: protocol.SIGNER_SCHEME_EDDSA,
					Eddsa: &protocol.EdDSA01SignerBuilder{
						NetworkType:     protocol.NETWORK_TYPE_TEST_NET,
						SignerPublicKey: stubByteArrayWithLength(16, 0x44),
					},
				},
				ContractName: "BenchmarkToken",
				MethodName:   "transfer",
				InputArgumentArray: (&protocol.MethodArgumentArrayBuilder{
					Arguments: []*protocol.MethodArgumentBuilder{
						{
							Name:        "amount",
							Type:        protocol.METHOD_ARGUMENT_TYPE_UINT_64_VALUE,
							Uint64Value: 0x3333,
						},
					},
				}).Build().RawArgumentsArray(),
			},
			Signature: stubByteArrayWithLength(16, 0x22),
		},
	}

	fmt.Println()
	fmt.Printf("*** Hex dump:\n\n")
	builder.HexDump("", 0)
	fmt.Println()

	fmt.Printf("*** Raw:\n\n")
	raw := builder.Build().Raw()
	membuffers.HexDumpRawInLines(raw, 0x20)
	fmt.Println()
}
