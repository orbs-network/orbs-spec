# Consensus Core

Provides the interface to the system for ConsensusAlgo. Responsible primarily for the creation of new blocks (populating the proposed block with transactions from TransactionPool) and for the validation of proposed blocks (like transaction results).

Holds the source of truth for the latest block height in the virtual chain.

Currently a single instance per virtual chain per node.

### Interacts with services

* ConsensusAlgo - achieves consensus on a block, uses lite-helix, a PBFT based algorithm.
* TransactionsPool - manages the pending and committed transaction pools, proposes and validates the trasnactions mixture of new blocks.
* VirtualMachine - used to execute the transactions.
* BlockStorage - called to commit new blocks upon consensus.

![alt text][consensus_core_interfaces] <br/><br/>

[consensus_core_interfaces]: consensus_core_interfaces.png "Consensus - Core Interfaces"

&nbsp;
## `Init Flow`
* Read configuration file:
  * Federation nodes (public keys)
  * Empty_block_wait
  * Out of sync timeout

&nbsp;
## `RequestCommittee` (method)

> Returns a sorted list of nodes that participate in the approval committee for a given block height.
* Committee members = all nodes' ids (Public Key).
* Order the nodes' ids based on the sorting algorithm.

&nbsp;
## `RequestNewTransactionsBlock` (method)

> Performed by the leader only, upon request from the algorithm to perform the ordering phase.

#### Ordering Block Creation
* Get the pending transactions by calling `TransactionPool.GetPendingTransactions`.
  * if there are no transactions to append:
    * Wait configurable empty_block_wait = 0.5sec and retry once.
    * If still empty continue with an empty block (# of transaction = 0).
* Build Ordering block
  * Current protocol version (0x1)
  * Virtual chain
  * Block height is given, confirm it's incremented from previous block (the latest).
  * Hash pointer to the previous (latest) block: SHA256(Block header)
  * 64b unix time stamp
  * The merkle root hash of the transactions in the block
  * Placeholder: Metadata - holds reputaion / algorithm data
  * SHA256 of the block metadata.

* Cache the transactions block for the execution validation part.
* You can optimize warm up by running the logic in `RequestNewResultsBlock` right now.

&nbsp;
## `RequestNewResultsBlock` (method)

> Performed by the leader only, upon request from the algorithm.

> The Consensus core receives the transactions block directly from the Ordering Consensus (Same nodes/cores in V1)
* Execute the ordered transactions set by calling `VirtualMachine.ProcessTransactionSet`, creating receipts and a state diff.
* Build Results Block
  * Current protocol version (0x1)
  * Virtual chain
  * Block height is incremented from previous block (the latest).
  * Hash pointer to the previous (latest) block - SHA256(Block header)
  * 64b unix time stamp
  * The merkle root hash of the block's transactions receipts
  * The merkle root of the state diff
  * Hash pointer to the transactions block of the same height - SHA256(Block header)
  * Merkle root of the state prior to the block execution, retrived by calling `StateStroage.GetStateHash`
  * Bloom filter
    * Set H(1,tx_id) for each transaction's tx_id // TBD sender_address, smart_contract_address.

## `ValidateTransactionsBlock` (method)
> Performed upon request from the algorithm recieving a block proposal.

#### Approve transactions for processing
* Check that the transactions were not expired. Transaction is considered expired if transaction timestamp > current timestamp + expiration window.
* Call `TransactionPool.ValidateTransactionsForOrdering` to validate PreOrder checks and no duplication.

### Ordering Block Checks
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

If one of the Transactions Block checks fails, return INVALID status.

## `Validate Results Block` (method)
* Check protocol verison
* Execute the trasnactions in the transactions Block by calling `VirtualMachine.ProcessTransactionSet`, creating receipts and a state diff.
  * Compare the receipts root and state diff hash. (Supports only determistic execution)
* Check that virtual chain matches consensus virtual chain
* Check block_height = previous block_height + 1
* Check prev_block_hash = SHA256(Block header(previously committed block))
* Check time_stamp within +/-2sec of local timestamp, and time_stamp > previous commited block.time_stamp
* Check ordering block hash.
* Check state root hash equal local state root hash (before executing the block). Get the state hash by calling `StateStroage.GetStateHash`

If one of the Results Block checks fails, return INVALID status.

&nbsp;
## `CommitBlock` (method)
> Commits the transactions and resutls block to the block storage.
* Match the commit to the block pair by verifying the blocks hash.
* Update the last commited transactions block height.
* Commit the block to the block storage by calling `BlockStorage.CommitTransactionsBlock`.
