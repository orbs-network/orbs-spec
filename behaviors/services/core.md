# Core
Responsible for controlling the flow of transactions across nodes for ordering, execution, and validation until they are committed to block storage and their state effects are committed to state storage.
The core maintains the consistency state (block_height) and is responsible for the consistency of the rest of the services.
The core utilizes 2 child services:
* Consensus algorithm - achieves consensus on a block, uses lite-helix, a PBFT based algorithm.
* Transactions pool - manages the pending and committed transaction pools

&nbsp;
## `AddPendingTransaction` (method)
Add transaction to pending pool by calling `TransactionPool.AddNewPendingTransaction`.

&nbsp;
## `Block Creation`
* Process performed by the leader only, initiated by the complition of the previous block round (block committed)
* Get the pending transactions calling `TransactionPool.GetAllPendingTransactions`.
  * if there are no transactions to append, continue with #of transaction = 0.
* Build ordered transactions set based on the ordering policy (in order).
* Execute the ordered transactions set by calling `VirtualMachine.ProcessTransactionSet`, creating receipts and a state diff.
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
    * If returns NOT_UPDATED, retry in X=100ms, doubling the wait time for each NOT_UPDATED response.
  * Hash pointer to the Ordering block of the same height - SHA256(Block header)
  * Merkle root of the state prior to the block execution
  * Bloom filter of the block's receipts

* Block = {Ordering Block, Execution Validation Block}
* Initiate a lite-helix consensus round, appending(block)

## `Block Validation - Ordering Block` (method)
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

## `Block Validation - Execution Validation Block` (method)
* Check block proof
* Check protocol verison
* Execute the Ordering Block trasnactions by calling `VirtualMachine.ProcessTransactionSet`, creating receipts and a state diff.
* Check that virtual chain matches consensus virtual chain 
* Check block_height = previous block_height + 1
* Check prev_block_hash = SHA256(Block header(previously committed block))
* Check time_stamp within +/-2sec of local timestamp, and time_stamp > previous commited block.time_stamp
* Check receipts root hash
* Check state_diff hash 
* Check ordering block hash equal SHA256(Block header(Ordering Block))
* Check state root hash equal local state root hash (before executing the block)
* Check state_diff hash equal local state_diff hash 
  * Supports only determistic execution 
* Check receipts root hash equal local receipts root hash 
  * Supports only determistic execution 

&nbsp;
## `Lite-Helix Messages`
* When Lite-Helix sends a broadcast message call `Gossip.BroadcastMessage` with the Lite-Helix message, send message to the Ordering_Nodes group.
* When Lite-Helix sends a unicast message call `Gossip.UnicastMessage` with the Lite-Helix message.

&nbsp;
## `BlockCommitted` (method)
> Initiated on lite-helix block commit. Updates the top_block value and adds the commited block to the block storage, commit state diff to the storage and updates the transaction pool by calling:
* `StateStroage.CommitStateDiff`
* `BlockStorage.CommitOrderingBlock`
* `BlockStorage.CommitValidationBlock`
* `TransactionPool.MarkCommittedTransactions`

&nbsp;
## `GossipMessageReceived` (method)
> Handles messages received from another node. 
* If the message is a consensus allgorithm message, pass to algorithm.
* If the message is a transaction batch
  * Validate batch signature, if mismatch log for reputation calculation.
  * Add transactions to pending pool by calling `TransactionPool.AddNewPendingTransaction`.


&nbsp;
## `Node Sync Flow`
Node sync is initiated upon reception of f+1 consensus mesasge (PrePrepare+Prepare / Commit) with term larger than the expected term. (likely to be called on node init)

`Sync flow`
* Randomly select one of the other nodes and send it a BLOCK_SYNC_REQUEST message using `Gossip.UnicastMessage` with block_height = top_block + 1.
  * If a response does not arrive within X=5sec, resend to another node.
* Upon receiption of a BLOCK_SYNC_RESPONSE message, perfom regular `BlockCommitted` flow.
* Repeat sending requests until receiving a valid PrePrepare message (including block_height = top_block + 1)

