package hexdumps

import (
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"testing"
)

func TestReceiptProof(t *testing.T) {
	builder := &protocol.ReceiptProofBuilder{
		Header: &protocol.ResultsBlockHeaderBuilder{
			ProtocolVersion:             1,
			VirtualChainId:              42,
			BlockHeight:                 0x2222,
			PrevBlockHashPtr:            stubByteArrayWithLength(16, 0x33),
			Timestamp:                   0x1122334455,
			ReceiptsRootHash:            stubByteArrayWithLength(16, 0x44),
			StateDiffHash:               stubByteArrayWithLength(16, 0x55),
			TransactionsBlockHashPtr:    stubByteArrayWithLength(16, 0x66),
			PreExecutionStateRootHash:   stubByteArrayWithLength(16, 0x77),
			TransactionsBloomFilterHash: stubByteArrayWithLength(16, 0x88),
			NumTransactionReceipts:      0x55,
			NumContractStateDiffs:       0x56,
		},
		BlockProof: &protocol.ResultsBlockProofBuilder{
			TransactionsBlockHash: stubByteArrayWithLength(16, 0x99),
			Type:      protocol.RESULTS_BLOCK_PROOF_TYPE_LEAN_HELIX,
			LeanHelix: []byte{0x01, 0x02, 0x03, 0x04},
		},
		ReceiptProof: stubByteArrayWithLength(16, 0x11),
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
