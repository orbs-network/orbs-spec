// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service TransactionPool

type TransactionPool interface {
	gossiptopics.TransactionRelayHandler
	AddNewTransaction(ctx context.Context, input *AddNewTransactionInput) (*AddNewTransactionOutput, error)
	GetCommittedTransactionReceipt(ctx context.Context, input *GetCommittedTransactionReceiptInput) (*GetCommittedTransactionReceiptOutput, error)
	GetTransactionsForOrdering(ctx context.Context, input *GetTransactionsForOrderingInput) (*GetTransactionsForOrderingOutput, error)
	ValidateTransactionsForOrdering(ctx context.Context, input *ValidateTransactionsForOrderingInput) (*ValidateTransactionsForOrderingOutput, error)
	CommitTransactionReceipts(ctx context.Context, input *CommitTransactionReceiptsInput) (*CommitTransactionReceiptsOutput, error)
	RegisterTransactionResultsHandler(handler handlers.TransactionResultsHandler)
}

/////////////////////////////////////////////////////////////////////////////
// message AddNewTransactionInput (non serializable)

type AddNewTransactionInput struct {
	SignedTransaction *protocol.SignedTransaction
}

func (x *AddNewTransactionInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedTransaction:%s,}", x.StringSignedTransaction())
}

func (x *AddNewTransactionInput) StringSignedTransaction() (res string) {
	res = x.SignedTransaction.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message AddNewTransactionOutput (non serializable)

type AddNewTransactionOutput struct {
	TransactionStatus  protocol.TransactionStatus
	TransactionReceipt *protocol.TransactionReceipt
	BlockHeight        primitives.BlockHeight
	BlockTimestamp     primitives.TimestampNano
}

func (x *AddNewTransactionOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionStatus:%s,TransactionReceipt:%s,BlockHeight:%s,BlockTimestamp:%s,}", x.StringTransactionStatus(), x.StringTransactionReceipt(), x.StringBlockHeight(), x.StringBlockTimestamp())
}

func (x *AddNewTransactionOutput) StringTransactionStatus() (res string) {
	res = fmt.Sprintf("%x", x.TransactionStatus)
	return
}

func (x *AddNewTransactionOutput) StringTransactionReceipt() (res string) {
	res = x.TransactionReceipt.String()
	return
}

func (x *AddNewTransactionOutput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *AddNewTransactionOutput) StringBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.BlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetCommittedTransactionReceiptInput (non serializable)

type GetCommittedTransactionReceiptInput struct {
	Txhash               primitives.Sha256
	TransactionTimestamp primitives.TimestampNano
}

func (x *GetCommittedTransactionReceiptInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Txhash:%s,TransactionTimestamp:%s,}", x.StringTxhash(), x.StringTransactionTimestamp())
}

func (x *GetCommittedTransactionReceiptInput) StringTxhash() (res string) {
	res = fmt.Sprintf("%s", x.Txhash)
	return
}

func (x *GetCommittedTransactionReceiptInput) StringTransactionTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.TransactionTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetCommittedTransactionReceiptOutput (non serializable)

type GetCommittedTransactionReceiptOutput struct {
	TransactionStatus  protocol.TransactionStatus
	TransactionReceipt *protocol.TransactionReceipt
	BlockHeight        primitives.BlockHeight
	BlockTimestamp     primitives.TimestampNano
}

func (x *GetCommittedTransactionReceiptOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionStatus:%s,TransactionReceipt:%s,BlockHeight:%s,BlockTimestamp:%s,}", x.StringTransactionStatus(), x.StringTransactionReceipt(), x.StringBlockHeight(), x.StringBlockTimestamp())
}

func (x *GetCommittedTransactionReceiptOutput) StringTransactionStatus() (res string) {
	res = fmt.Sprintf("%x", x.TransactionStatus)
	return
}

func (x *GetCommittedTransactionReceiptOutput) StringTransactionReceipt() (res string) {
	res = x.TransactionReceipt.String()
	return
}

func (x *GetCommittedTransactionReceiptOutput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *GetCommittedTransactionReceiptOutput) StringBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.BlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsForOrderingInput (non serializable)

type GetTransactionsForOrderingInput struct {
	CurrentBlockHeight       primitives.BlockHeight
	PrevBlockTimestamp       primitives.TimestampNano
	PrevBlockReferenceTime   primitives.TimestampSeconds
	MaxTransactionsSetSizeKb uint32
	MaxNumberOfTransactions  uint32
}

func (x *GetTransactionsForOrderingInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,PrevBlockTimestamp:%s,PrevBlockReferenceTime:%s,MaxTransactionsSetSizeKb:%s,MaxNumberOfTransactions:%s,}", x.StringCurrentBlockHeight(), x.StringPrevBlockTimestamp(), x.StringPrevBlockReferenceTime(), x.StringMaxTransactionsSetSizeKb(), x.StringMaxNumberOfTransactions())
}

func (x *GetTransactionsForOrderingInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *GetTransactionsForOrderingInput) StringPrevBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockTimestamp)
	return
}

func (x *GetTransactionsForOrderingInput) StringPrevBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockReferenceTime)
	return
}

func (x *GetTransactionsForOrderingInput) StringMaxTransactionsSetSizeKb() (res string) {
	res = fmt.Sprintf("%v", x.MaxTransactionsSetSizeKb)
	return
}

func (x *GetTransactionsForOrderingInput) StringMaxNumberOfTransactions() (res string) {
	res = fmt.Sprintf("%v", x.MaxNumberOfTransactions)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionsForOrderingOutput (non serializable)

type GetTransactionsForOrderingOutput struct {
	SignedTransactions         []*protocol.SignedTransaction
	ProposedBlockTimestamp     primitives.TimestampNano
	ProposedBlockReferenceTime primitives.TimestampSeconds
}

func (x *GetTransactionsForOrderingOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SignedTransactions:%s,ProposedBlockTimestamp:%s,ProposedBlockReferenceTime:%s,}", x.StringSignedTransactions(), x.StringProposedBlockTimestamp(), x.StringProposedBlockReferenceTime())
}

func (x *GetTransactionsForOrderingOutput) StringSignedTransactions() (res string) {
	res = "["
	for _, v := range x.SignedTransactions {
		res += v.String() + ","
	}
	res += "]"
	return
}

func (x *GetTransactionsForOrderingOutput) StringProposedBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ProposedBlockTimestamp)
	return
}

func (x *GetTransactionsForOrderingOutput) StringProposedBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.ProposedBlockReferenceTime)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsForOrderingInput (non serializable)

type ValidateTransactionsForOrderingInput struct {
	CurrentBlockHeight        primitives.BlockHeight
	CurrentBlockTimestamp     primitives.TimestampNano
	CurrentBlockReferenceTime primitives.TimestampSeconds
	PrevBlockReferenceTime    primitives.TimestampSeconds
	SignedTransactions        []*protocol.SignedTransaction
}

func (x *ValidateTransactionsForOrderingInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,CurrentBlockTimestamp:%s,CurrentBlockReferenceTime:%s,PrevBlockReferenceTime:%s,SignedTransactions:%s,}", x.StringCurrentBlockHeight(), x.StringCurrentBlockTimestamp(), x.StringCurrentBlockReferenceTime(), x.StringPrevBlockReferenceTime(), x.StringSignedTransactions())
}

func (x *ValidateTransactionsForOrderingInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *ValidateTransactionsForOrderingInput) StringCurrentBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockTimestamp)
	return
}

func (x *ValidateTransactionsForOrderingInput) StringCurrentBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockReferenceTime)
	return
}

func (x *ValidateTransactionsForOrderingInput) StringPrevBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockReferenceTime)
	return
}

func (x *ValidateTransactionsForOrderingInput) StringSignedTransactions() (res string) {
	res = "["
	for _, v := range x.SignedTransactions {
		res += v.String() + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsForOrderingOutput (non serializable)

type ValidateTransactionsForOrderingOutput struct {
}

func (x *ValidateTransactionsForOrderingOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message CommitTransactionReceiptsInput (non serializable)

type CommitTransactionReceiptsInput struct {
	ResultsBlockHeader       *protocol.ResultsBlockHeader
	TransactionReceipts      []*protocol.TransactionReceipt
	LastCommittedBlockHeight primitives.BlockHeight
}

func (x *CommitTransactionReceiptsInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ResultsBlockHeader:%s,TransactionReceipts:%s,LastCommittedBlockHeight:%s,}", x.StringResultsBlockHeader(), x.StringTransactionReceipts(), x.StringLastCommittedBlockHeight())
}

func (x *CommitTransactionReceiptsInput) StringResultsBlockHeader() (res string) {
	res = x.ResultsBlockHeader.String()
	return
}

func (x *CommitTransactionReceiptsInput) StringTransactionReceipts() (res string) {
	res = "["
	for _, v := range x.TransactionReceipts {
		res += v.String() + ","
	}
	res += "]"
	return
}

func (x *CommitTransactionReceiptsInput) StringLastCommittedBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.LastCommittedBlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CommitTransactionReceiptsOutput (non serializable)

type CommitTransactionReceiptsOutput struct {
	NextDesiredBlockHeight   primitives.BlockHeight
	LastCommittedBlockHeight primitives.BlockHeight
}

func (x *CommitTransactionReceiptsOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{NextDesiredBlockHeight:%s,LastCommittedBlockHeight:%s,}", x.StringNextDesiredBlockHeight(), x.StringLastCommittedBlockHeight())
}

func (x *CommitTransactionReceiptsOutput) StringNextDesiredBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.NextDesiredBlockHeight)
	return
}

func (x *CommitTransactionReceiptsOutput) StringLastCommittedBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.LastCommittedBlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
