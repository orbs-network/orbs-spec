# Consensus-Core
Respomsible for the producing and valdiation of ordering and validation blocks.
The consensus core maintains the consistency state (latest block_height) and is responsible for the consistency of the rest of the services.
The consensus core interacts with the follwoing services as part of the consensus process:
* Consensus algorithm - achieves consensus on a block, uses lite-helix, a PBFT based algorithm.
* Transactions pool - manages the pending and committed transaction pools, proposes and validates the trasnactions mixture of new blocks.
* VM - used to execute the transactions.
* Block storage - called to commit new blocks upon consensus.

&nbsp;
## `Create Block` (method)
* Process performed by the leader only, initiated by the complition of the previous block round (block committed)
* Get the pending transactions calling `TransactionPool.GetAllPendingTransactions`.
  * if there are no transactions to append, continue with #of transaction = 0.
* Build ordered transactions set based on the ordering policy (in order).
* Perform `Validate PreOrder` check on each transaction.
* Execute the ordered transactions set by calling `VirtualMachine.ProcessTransactionSet`, creating receipts and a state diff.
  * If returns OUT_OF_SYNC retry, until timeout = 10 sec.
* Build Ordering block
  * Current protocol version (0x1)
  * Virtual chain 
  * Block height is incremented from previous block (the latest).
  * Hash pointer to the previous (latest) block - SHA256(Block header) 
  * 64b unix time stamp 
  * The merkle root hash of the block's transactions
  * Metadata - holds reputaion / algorithm data
  * SHA256 of the block metadata.
* Build Execution Validation Block
  * Current protocol version (0x1)
  * Virtual chain 
  * Block height is incremented from previous block (the latest).
  * Hash pointer to the previous (latest) block - SHA256(Block header)
  * 64b unix time stamp 
  * The merkle root hash of the block's transactions receipts
  * The merkle root of the state diff, retrived by calling `StateStroage.GetStateHash`
    * If returns OUT_OF_SYNC retry, until timeout = 10 sec.
  * Hash pointer to the Ordering block of the same height - SHA256(Block header)
  * Merkle root of the state prior to the block execution
  * Bloom filter of the block's receipts

* Block = {Ordering Block, Execution Validation Block}
* Initiate a lite-helix consensus round, appending(block)

## `Validate PreOrder` (method)
> Performs on each transaction in the proposed OrderingBlock similar checks as the ones done in the transaction pool to verify them under consensus.

#### Check transaction validity
* Correct protocol version.
* Valid fields (sender address, contract address).
* Sender virtual chain matches contract virtual chain.
* Check transaction time_stamp, accept only transactions with last block.timestamp -2sec  < time_stamp < block.timestamp expiration window. 
* Valid transaction signature.

#### Approve transaction for processing
* Not expired. Transaction is considered expired if transaction timestamp > current timestamp + expiration window.
* Check that the Subscription status is active.
* Transaction doesn't already exist in the committed pool (duplicated). Validating by calling `TransactionPool.ValidateTransactionsForOrdering`.

## `Validate Ordering Block` (method)
* Check block proof
* Check protocol verison
* Check that virtual chain matches consensus virtual chain 
* Check block_height = previous block_height + 1
* Check prev_block_hash = SHA256(Block header(previously committed block))
* Check time_stamp within +/-2sec of local timestamp, and time_stamp > previous commited block.time_stamp
* Check transaction root hash 
* Check Metadata hash 
* Check ordering policy (not verified)
* Check no previously committed transactions
* Check no duplicated transactions

## `Validate Execution Validation Block` (method)
* Check block proof
* Check protocol verison
* Execute the Ordering Block trasnactions by calling `VirtualMachine.ProcessTransactionSet`, creating receipts and a state diff.
  * If returns OUT_OF_SYNC treat as invalid block.
* Check that virtual chain matches consensus virtual chain 
* Check block_height = previous block_height + 1
* Check prev_block_hash = SHA256(Block header(previously committed block))
* Check time_stamp within +/-2sec of local timestamp, and time_stamp > previous commited block.time_stamp
* Check receipts root hash
* Check state_diff hash 
* Check ordering block hash equal SHA256(Block header(Ordering Block))
* Check state root hash equal local state root hash (before executing the block). Get the state hash by calling `StateStroage.GetStateHash`
  * If returns OUT_OF_SYNC retry.
* Check state_diff hash equal local state_diff hash 
  * Supports only determistic execution 
* Check receipts root hash equal local receipts root hash 
  * Supports only determistic execution 

&nbsp;
## `RequestNewBlock` (method)
Calls `Create Block` return {Ordering Block, Validation Block}

&nbsp;
## `ValidateBlock` (method)
Calls `Validate PreOrder`
Calls `Validate Ordering Block`
Calls `Validate Execution Validation Block`
returns Valid if all are valid.

&nbsp;
## `CommitBlock` (method)
> Initiated on lite-helix block commit. Updates the top_block value and adds the commited block to the block storage by calling:
* `BlockStorage.CommitOrderingBlock`
* `BlockStorage.CommitValidationBlock`

## `UpdateSubscriptionStatus` (method)
> Updates the transaction pool Subscription status.
* Update the local subscription status for the virtual chain, takes effect starting from the indicated block. Used for the subscription validation as part of the 
