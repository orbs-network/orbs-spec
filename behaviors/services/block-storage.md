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
## `Inter Node Block Synchronization` <!-- oded will finish -->

* The synchronization process is initialled upon either:
  * Receving CommitBlock block with height <> last_commited_block block + 1.
  * Receving a request to fetch data from blocks that are unavailable to the block storage. (from the state storage or transactions pool)
  * block_timeout = 8 sec has passed since the last block commit.

### `Inter Node Block Synchronization` <!-- oded will finish, combine -->
* Identify ndoes that have the desired blocks by broadcasting a BLOCK_AVAILABILITY_REQUEST messsage to all nodes using `Gossip.SendMessage`.
  * If the last block is unknown, set the lsat_block to MAX_UINT64.
  * The receiving nodes respond with a BLOCK_AVAILABILITY_RESPONSE message indicating the range out of the desired blocks that is available to them and their current top block.
* Randomly select one of the nodes that responsded and request a batch of blocks avilable at this node.
  * Request the blocks in order
  * Request up to max_block_sync_batch = 100 in each request.
  * The receiving node responds with a BLOCK_SYNC_RESPONSE message that include the requested blocks. In addition, the message includes the node's top block height.
* Upon reception of a response, add the blocks in order by performing `AddBlock`.
* If the node attempt to sync to the latest block and the received top block is not already available to the node, update the requested last block to the received top block.
* Repeat requesting until receiving all the desired blocks.
  * It is recommanded to shuffle the requests among the nodes that have the block avialable in order not to put too much burden on a single node.

&nbsp;
## `Intra Node Block Synchronization` <!-- oded will finish -->

* This flow should be implemented twice - one instance for service TransactionPool and one instance for service StateStorage
* Endless loop which is continuously synchronizing the service with committed blocks based on the service's `next_desired_block_height`
* If `last_commited_block` < `next_desired_block_height` block and wait (without polling), else:
  * Commit `next_desired_block_height` to the service by calling `StateStorage.CommitStateDiff` or `TransactionPool.CommitTransactionsReceipts`.
  * Update `next_desired_block_height` according to the call response.
  * Continue iterating.
* Important: When everything is synchronized (common case), `AddBlock` should immediately trigger the commit to the service without waiting (this is the critical path).

&nbsp;
## `AddBlock` (method) <!-- tal can finish -->

> Add a new block pair (transactions block and results block) to the database.

#### Verify block before adding to database
* Check the block protocol version.
* If the block already exist (block height < last_commited_block block + 1) - sligntly discrad. (may occur during node sync)
* If block height > last_commited_block block + 1 indicates that the block storage is possibly out of sync, add the to the storage without committing it and initiate `Block Synchronization Flow`.
  * The block should be stored such that it can be committed once the previous blocks were committed.

#### Add block to storage
* Store block in database indexed by block height.

#### Commit block
* If block height = last_commited_block block + 1, check hash pointer, if doesn't match then discard block.
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

![alt text][block_state_pool_flow] <br/><br/>

[block_state_pool_flow]: block_state_pool_flow.png "Block Storage - State Storage / Transaction Pool"
