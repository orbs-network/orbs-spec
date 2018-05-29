# Block Storage

Holds the long term journal of all confirmed blocks. Continuously synchronizes with other nodes when missing blocks are required.

Currently a single instance per virtual chain per node.

#### Interacts with services

* StateStorage - state diffs
* TransactionPool - transaction receipts.

&nbsp;
## `Init` <!-- oded will finish -->

TODO

&nbsp;
## `Data Structures` <!-- tal will finish -->

#### Block database
* Holds all blocks with ability to query efficiently by block height.
* Able to store blocks out of order (upon sync lost).
  * These blocks are not yet marked as committed.
  * They will be verified and committed after receiving the missing blocks before them.
* Needs to be persistent.
* When empty, initialized with the genesis block.
* Data is not sharded, all blocks are stored locally on all machines.

&nbsp;
## `Inter Node Block Synchronization`

* The synchronization flow is initialled upon:
  * block_timeout = configurable block_commit_timeout (default 8 sec) has passed without a block commit.
  * Init flow.

* Identify nodes that have the desired blocks by broadcasting a `BLOCK_AVAILABILITY_REQUEST` messsage to all nodes using `Gossip.SendMessage`.
  * Request the blocks from the last_committed_block to the highest known block height. (if unknown set to MAX_UINT64)
  * The receiving nodes respond with a `BLOCK_AVAILABILITY_RESPONSE` message
    * Indicating the range out of the desired blocks that is available to them and their current top block.
* Randomly select one of the responding nodes and request a batch of blocks avilable to this node by sending a unicast `BLOCK_SYNC_REQUEST`.
  * The receiving node responds with a `BLOCK_SYNC_RESPONSE` message
    * The message includes the requested blocks and the node's last_committed block height.
  * Upon reception of a response
    * Validate the blocks for commit by performing `Validate Block for Commit`
      * If valid commit the block by calling `Commit Block`.
    * If the responder last_committed block is larger than the maximum known block height, update the following requests accordingly.
* Repeat requesting until receiving all the desired blocks.
  * It is recommanded to shuffle the requests among the nodes that have the block avialable in order not to put too much burden on a single node.

#### `Validate Block for Commit`
> Validates a block received during inter-node sync  

#### Check the Transactions Block Header (stateless)
* Check the block protocol version.
* Check the virtual chain.
* Check block height
  * If the block already exist (block height != last_commited_block block + 1) discard.
* Check transactions_root_hash 
  * Calculate the merkle root hash of the block's transactions verify the hash in the header.
* Check metadata hash
  * Calculate the hash of the block's metadata and verify the hash in the header.

#### Check the Results Block Header (stateless)
* Check the block protocol version.
* Check the virtual chain.
* Check block height
  * If the block already exist (block height != last_commited_block block + 1) discard.
* Check receipts_root_hash
  * Calculate the merkle root hash of the block's transactions verify the hash in the header.
* Check state_diff_hash
  * Calculate the hash of the block's metadata and verify the hash in the header.

#### Validate that the block is under consensus and can be committed.
* Validate that the block can be commited under consensus by calling:
  *  `ConsensusAlgo.AcknowledgeTransactionsBlockConsensus`. 
  *  `ConsensusAlgo.AcknowledgeResultsBlockConsensus`.

&nbsp;
## `Intra Node Block Synchronization` <!-- oded will finish -->

* This flow should be implemented twice - one instance for service TransactionPool and one instance for service StateStorage
* Endless loop which is continuously synchronizing the service with committed blocks based on the service's `next_desired_block_height`
* If `last_commited_block` < `next_desired_block_height` block and wait (without polling), else:
  * Commit `next_desired_block_height` to the service by calling `StateStorage.CommitStateDiff` or `TransactionPool.CommitTransactionsReceipts`.
  * Update `next_desired_block_height` according to the call response.
  * Continue iterating.
* Important: When everything is synchronized (common case), `CommitBlock` should immediately trigger the commit to the service without waiting (this is the critical path).

&nbsp;
## `CommitBlock` (method)
> Commit a new block pair after it was approved for commit.

#### Verify block before adding to database
* Check the block protocol version.
* If the block already exist (block height < last_commited_block block + 1) - silently discard.
* If block height > last_commited_block block + 1, panic.

#### Commit the block
* Store block in database indexed by block height.
* Update last_commited_block = block height.
* If intra block sync of StateStorage is blocking and waiting, wake it up.
* If intra block sync of TransactionPool is blocking and waiting, wake it up.

&nbsp;
## `GetTransactionReceipt` (method) <!-- tal will finish -->

> Returns the transaction receipt for a transaction based on its tx_id and time_stamp.

* Go over all the bloom filters in the blocks where the transaction could be:
  * Starting where block timestamp is transaction timestamp minus small grace (eg. 5 seconds)
  * Finishing where block timestamp is transaction timestamp plus configurable time_window plus small grace (eg. 5 seconds)
* For each block, lookup the transaction timestamp in the block header timestamp bloom filter and the tx_id in the block header tx_id bloom filter.
  * On match, fetchs the block and search the tx_id in the block receipts.
      * If found, returns the receipt
* If not found on all relevant blocks, returns NULL.

&nbsp;
## `GossipMessageReceived` (method)
> Handles a gossip message from another node. Relevant messages include node sync messages.

TODO - add sync flow messages

## `GetTransactionsBlockHeader` (method)
> Returns a committed transactions block header and proof.
* If requested block height is in the future but `last_committed_block` is close to it (configurable distance) block and wait
* If requested block height is in the future but `last_committed_block` is far, fail
* Return the transactions block header, metadata, the transactions block proof and the results block proof.

## `GetResultsBlockHeader` (method)
> Returns a committed results block header and proof.
* If requested block height is in the future but `last_committed_block` is close to it (configurable distance) block and wait
* If requested block height is in the future but `last_committed_block` is far, fail
* Return the results block header, metadata, the transactions block proof and the results block proof.

![alt text][block_state_pool_flow] <br/><br/>

[block_state_pool_flow]: block_state_pool_flow.png "Block Storage - State Storage / Transaction Pool"
