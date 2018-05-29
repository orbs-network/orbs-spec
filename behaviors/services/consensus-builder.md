# Consensus Builder

Builds and validates the actual content of blocks for the consensus process. Responsible primarily for the creation of new blocks (populating with transactions from the pool) and for content validation of proposed blocks (verifying transaction results).

Currently a single instance per virtual chain per node.

### Interacts with services

* ConsensusAlgo - achieves consensus on a block, uses lite-helix, a PBFT based algorithm.
* TransactionsPool - manages the pending and committed transaction pools, proposes and validates the trasnactions mixture of new blocks.
* VirtualMachine - used to execute the transactions.
* BlockStorage - called to commit new blocks upon consensus.

![alt text][consensus_core_interfaces] <br/><br/>

[consensus_core_interfaces]: consensus_core_interfaces.png "Consensus - Core Interfaces"

&nbsp;
## `Init` <!-- oded will finish -->

&nbsp;
## `Data Structures` <!-- tal can finish -->

#### Sync state
* `last_committed_block`

&nbsp;
## `RequestOrderingCommittee` (method) <!-- tal can finish -->

> Returns a sorted list of nodes that participate in the approval committee for the ordering of a given block height.

* Committee members = all nodes' ids (Public Key).
* Order the nodes' ids based on the sorting algorithm.

&nbsp;
## `RequestValidationCommittee` (method) <!-- tal can finish -->

> Returns a sorted list of nodes that participate in the approval committee for the execution validation of a given block height.

* Committee members = all nodes' ids (Public Key).
* Order the nodes' ids based on the sorting algorithm.


&nbsp;
## `RequestNewTransactionsBlock` (method) <!-- tal can finish -->

> Performed by the leader only, upon request from the algorithm to perform the ordering phase.

* If not synchronized with block height (has the last block header), don't participate and fail.

#### Ordering Block Creation
* Get the pending transactions by calling `TransactionPool.GetTransactionsForOrdering`.
  * if there are no transactions to append:
    * Wait configurable empty_block_wait = 0.5sec and retry once.
    * If still empty continue with an empty block (# of transaction = 0).
* Build Ordering block
  * Current protocol version (0x1)
  * Virtual chain
  * Block height is given, panic if it's not incremented from previous block (the latest).
  * Hash pointer to the previous (latest) transactions block: SHA256(Block header)
  * 64b unix time stamp
  * The merkle root hash of the transactions in the block
  * Placeholder: Metadata - holds reputaion / algorithm data
  * SHA256 of the block metadata.

* Cache the transactions block for the execution validation part.
* You can optimize warm up by running the logic in `RequestNewResultsBlock` right now.

&nbsp;
## `RequestNewResultsBlock` (method) <!-- tal can finish -->

> Performed by the leader only, upon request from the algorithm to perform the execution phase.

* If not synchronized with block height (has the last block header), don't participate and fail.

* The transactions block for this block height should be cached from previous call to `RequestNewTransactionsBlock`.
* Execute the ordered transactions set by calling `VirtualMachine.ProcessTransactionSet`, creating receipts and a state diff.
* Build Results Block
  * Current protocol version (0x1)
  * Virtual chain
  * Block height is given, panic if not it's incremented from previous block (the latest).
  * Hash pointer to the previous (latest) results block: SHA256(Block header)
  * 64b unix time stamp
  * The merkle root hash of the transactions receipts in the block
  * The hash of the state diff in the block
  * Hash pointer to the transactions block of the same height: SHA256(Block header)
  * Merkle root of the state prior to the block execution, retrived by calling `StateStroage.GetStateHash`
  * Bloom filter
    * Set H(1,tx_id) for each transaction's tx_id (concat byte with value 0x01 with tx_id and insert)

&nbsp;
## `ValidateTransactionsBlock` (method) <!-- tal can finish -->

> Performed upon request from the algorithm recieving a block proposal.

* If not synchronized with block height (has the last block header), don't participate and fail.

#### Approve transactions for processing
* Check that the transactions were not expired. Transaction is considered expired if transaction timestamp > current timestamp + expiration window.
* Call `TransactionPool.ValidateTransactionsForOrdering` to validate PreOrder checks and no duplication.

### Ordering Block Checks
* Check protocol verison
* Check that virtual chain matches consensus virtual chain
* Check block_height = previous block_height + 1
* Check prev_block_hash = SHA256(Block header(previously committed block))
* Check time_stamp within +/-2sec of local timestamp, and time_stamp > previous commited block.time_stamp
* Check transaction merkle root hash
* Check Metadata hash

* If one of the Transactions Block checks fails, return INVALID status, else return VALID.

&nbsp;
## `ValidateResultsBlock` (method) <!-- tal can finish -->

* If not synchronized with block height (has the last block header), don't participate and fail.

* Check protocol verison
* Execute the trasnactions in the transactions Block by calling `VirtualMachine.ProcessTransactionSet`, creating receipts and a state diff.
  * Compare the receipts merkle root and state diff hash. (Supports only determistic execution)
* Check that virtual chain matches consensus virtual chain
* Check block_height = previous block_height + 1
* Check prev_block_hash = SHA256(Block header(previously committed block))
* Check time_stamp within +/-2sec of local timestamp, and time_stamp > previous commited block.time_stamp
* Check ordering block hash.
* Check state merkle root hash equal local state merkle root hash (before executing the block). Get the state hash by calling `StateStroage.GetStateHash`

* If one of the Results Block checks fails, return INVALID status, else return VALID.
