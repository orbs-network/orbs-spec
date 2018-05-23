# Consensus-Core / Block Generator
Responsible for the producing and valdiation of transactions and results blocks.
The consensus core maintains the consistency state (latest block_height) and is responsible for the consistency of the rest of the services.
The consensus core interacts with the follwoing services as part of the consensus process:
* Consensus algorithm - achieves consensus on a block, uses lite-helix, a PBFT based algorithm.
* Transactions pool - manages the pending and committed transaction pools, proposes and validates the trasnactions mixture of new blocks.
* VM - used to execute the transactions.
* Block storage - called to commit new blocks upon consensus.

![alt text][consensus_core_interfaces] <br/><br/>

[consensus_core_interfaces]: consensus_core_interfaces.png "Consensus - Core Interfaces"

&nbsp;
## Init Flow
* Read configuration file:
  * Federation nodes (public keys)
  * Empty_block_wait
  * Out of sync timeout

&nbsp;
## `RequestNewOrderingBlock` (method)
> Performed by the leader only, upon request from the algorithm.

### Ordering Block Creation
* Get the pending transactions by calling `TransactionPool.GetPendingTransactions`.
  * If returns OUT_OF_SYNC retry, until timeout = 10 sec, on timeout return NULL.
  * if there are no transactions to append:
    * Wait empty_block_wait = 0.5sec and retry.
    * If still empty continue with an empty block (# of transaction = 0).
* Build Ordering block
  * Current protocol version (0x1)
  * Virtual chain 
  * Block height is incremented from previous block (the latest).
  * Hash pointer to the previous (latest) block - SHA256(Block header) 
  * 64b unix time stamp 
  * The merkle root hash of the block's transactions
  * Metadata - holds reputaion / algorithm data
  * SHA256 of the block metadata.
* Cache the transactions block for the execution validation part. 

&nbsp;
## `RequestNewResultsBlock` (method)
> Performed by the leader only, upon request from the algorithm.
> The Consensus core receives the transactions block directly from the Ordering Consensus (Same nodes/cores in V1)
* Execute the ordered transactions set by calling `VirtualMachine.ProcessTransactionSet`, creating receipts and a state diff.
  * If returns OUT_OF_SYNC retry, until timeout = 10 sec, on timeout return NULL.
* Build Results Block
  * Current protocol version (0x1)
  * Virtual chain 
  * Block height is incremented from previous block (the latest).
  * Hash pointer to the previous (latest) block - SHA256(Block header)
  * 64b unix time stamp 
  * The merkle root hash of the block's transactions receipts
  * The merkle root of the state diff, retrived by calling `StateStroage.GetStateHash`
    * If returns OUT_OF_SYNC retry, until timeout = 10 sec.
  * Hash pointer to the transactions block of the same height - SHA256(Block header)
  * Merkle root of the state prior to the block execution
  * Bloom filter
    * Set H(1,tx_id) for each transaction's tx_id // TBD sender_address, smart_contract_address.

## `ValidateTransactionsBlock` (method)
> Performed upon request from the algorithm recieving a block proposal. 

### PreOrder checks
> Performs on each transaction in the proposed TransactionsBlock similar checks as the ones done in the transaction pool to verify them under consensus.

#### Check transaction validity
// Moved to transaction pool.

#### Approve transaction for processing
* Not expired. Transaction is considered expired if transaction timestamp > current timestamp + expiration window.
* Check that the Subscription status is active.
* Transaction doesn't already exist in the committed pool (duplicated). Validating by calling `TransactionPool.ValidateTransactionsForOrdering`.

If one of the PreOrder checks fails, return INVALID status.

### Ordering Block Checks
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

If one of the Oredring Block checks fails, return INVALID status.

## `Validate Results Block` (method)
* Check block proof
* Check protocol verison
* Execute the trasnactions in the transactions Block by calling `VirtualMachine.ProcessTransactionSet`, creating receipts and a state diff.
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

If one of the Oredring Block checks fails, return INVALID status.

&nbsp;
## `CommitOrderingBlock` (method)
> Commits the transactions block to the block storage.
* Update the last commited transactions block height. 
* Commit the block to the block storage by calling `BlockStorage.CommitTransactionsBlock`.

&nbsp;
## `CommitResultsBlock` (method)
> Commits the results block to the block storage.
* Update the last commited results block height. 
* Commit the block to the block storage by calling `BlockStorage.CommitResultsBlock`.

&nbsp;
## `RequestOrderingCommittee` (method)
> Reurns a list of nodes to participate in the ordering consensus round.
* Return a list of all the nodes' ids (Public Key)
  * To be updated to random committies.

&nbsp;
## `RequestResultsCommittee` (method)
> Reurns a list of nodes to participate in the execution validation consensus round.
* Return a list of all the nodes' ids (Public Key)
  * To be updated to random committies.
