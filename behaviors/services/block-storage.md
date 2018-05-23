# Block Storage
Holds the long term journal of all confirmed blocks. Provides the state stroage with the state diff and the trasnaction pool with the transaction receipts. Syncs with other nodes when missing blocks are required.

&nbsp;
## Init Flow
* Subscribe to node sync gossip messages.
* Read persistent data
  * Start updating the state storage and transaction pool once they request to.

&nbsp;
## `Data Structures`

#### Block database
* Holds all blocks with ability to query efficiently by block height or block ID.
* Able to store blocks out of order (upon sync lost) and verify the headers when receiving the missing blocks / block headers.
* Able to return efficiently the last_commited_block - was verified (by validationing the headers hash chain) block height in sync in the database.
* Able to efficiently fetch the headers, receipts and state diff independently.
* Needs to be persistent.
* When empty, initialized with the genesis block.
* Data is not sharded, all blocks are stored locally on all machines.
* Recommended implementation: LevelDB
  * Fast key-value storage library written at Google.
  * Provides an ordered mapping from string keys to string values.

&nbsp;
## `Block Synchronization`

* The synchronization process is initialled upon either:
  * Receving CommitBlock block with height <> last_commited_block block + 1.
  * Receving a request to fetch data from blocks that are unavailable to the block storage. (from the state storage or transactions pool)
  * block_timeout = 8 sec has passed since the last block commit.

### `Block Synchronization Flow`
* Identify ndoes that have the desired blocks by broadcasting a BLOCK_AVAILABILITY_REQUEST messsage to all nodes using `Gossip.UnicastMessage`.
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
## `AddBlock` (rpc)
> Add a new block to the database

#### Verify block before adding to database
* Check the block protocol version.
* If the block already exist (block height < last_commited_block block + 1) - sligntly discrad. (may occur during node sync)
* If block height > last_commited_block block + 1 indicates that the block storage is out of sync, add the to the storage without committing it and initiate `Block Synchronization Flow`.
  * The block should be stored such that it can be committed once the previous blocks were committed.

#### Add block to storage
* Store block in database indexed by block height.

#### Commit block 
* If block height = last_commited_block block + 1, check hash pointer, if doesn't match then discard block. 
  * Update last_commited_block = block height.
  * Update the subscribed state storage
    * If the state storage consumer_block_height + 1 = last_commited_block, send StateDiff update to state stroage by calling `StateStorage.CommitStateDiff`.
  * Update the subscribed transaction pool
    * If the transaction pool consumer_block_height + 1 = last_commited_block, send the tarsanctions receipt update to the transaction pool by calling `TransactionPool.MarkCommittedTransactions`.

&nbsp;
## `GetBlocksByHeight` (rpc)
> Return an array of blocks of a range of heights along with the last_commited_block and the last_added_block. (last_commited_block <> last_added_block indicates out of sync).
* If the block range is not present, check that it within the added range by calling `ConsensusCore.GetTopBlockHeight`, if it is, fetch the missing blocks by initiating Block Synchronization Flow. 
* TODO - Currently not used in any V1 flow (may be needed for future contract calls)

&nbsp;
## `GetTransactionsBlockByHeight / GetResultsBlockByHeight` (rpc)
> Similar to `GetBlocksByHeight`, returns only the transactions block / results block parts.

&nbsp;
## `GetLastBlockHeight` (rpc)
> Returns the latest committed and addeded block heigths.

&nbsp;
## `GossipMessageReceived` (rpc)
> Handles a gossip message from another node. Relevant messages include node sync messages.

&nbsp;
## `GetTransactionReceipt` (rpc)
> Returns the transaction receipt for a transaction based on its tx_id and time_stamp.
* Fetch all relevant block headers based on the time_stamp
  * Block headers / bloom filters should be stored such they can be fetched as a bulk. 
  * For 30min time window, with average block time of 2sec, need to fetch ~900 headers ~8MB //TBD final bloom filter size.
* Search for the tx_id in the bloom filter
  * If hit, fetch the Valdiation Block receipts and search for the tx_id.
* If tx_id found, return receipt //TBD proof (block proof + merkle leaf).
* if tx_id wasn't found return receipt = NULL.


![alt text][block_state_pool_flow] <br/><br/>

[block_state_pool_flow]: block_state_pool_flow.png "Block Storage - State Storage / Transaction Pool"

## `RequestReceiptsUpdate` (rpc)
> Used by a transaction pool to subscribe for receipts update and re-sync.
* TODO Checkpoints.
* Check consumer_id, if not exist add to database, if exist update.
* If target_block_height = MAX_UINT64, continue to commit until notified otherwise (subscribed to commits)
* Check the requested range (consumer_block_height - target_block_height), if not present, check that it within the added range by calling `ConsensusCore.GetTopBlock`, if it is, fetch the missing blocks by initiating `Block Synchronization Flow`. 
* Update the consumer_block_height.
* If consumer_block_height block was committed, send the tarsanctions receipt update to the transaction pool by calling `TransactionPool.MarkCommittedTransactions`.
  * Wait for the response and update the consumer_block_height. 
* Repeat until consumer_block_height isn't committed.
* 

## `RequestStateDiffUpdate` (rpc)
> Used by a state storage to subscribe for receipts update and re-sync.
* TODO Checkpoints.
* Check consumer_id, if not exist add to database, if exist update.
* If target_block_height = MAX_UINT64, continue to commit until notified otherwise (subscribed to commits)
* Check the requested range (consumer_block_height - target_block_height), if not present, check that it within the added range by calling `ConsensusCore.GetTopBlock`, if it is, fetch the missing blocks by initiating `Block Synchronization Flow`. 
Update the consumer_block_height.
* If consumer_block_height block was committed, send StateDiff update to state stroage by calling `StateStorage.CommitStateDiff`.
  * Wait for the response and update the consumer_block_height. 
* Repeat until consumer_block_height isn't committed.