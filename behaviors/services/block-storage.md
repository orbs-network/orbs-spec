# Block Storage
Holds the long term journal of all confirmed blocks, used to reconstruct the state. 
Provides the state stroage with the state diff and the trasnaction pool with the transaction receipts.
Syncs with other nodes when missing blocks are required.

&nbsp;
## `Data Structures`

#### Block database
* Holds all blocks with ability to query efficiently by block height or block ID.
* Able to store blocks out of order (upon sync lost) and verify the headers when receiving the missing blocks / block headers.
* Able to return efficiently the last_commited_block - was verified (by validationing the headers hash chain) block height in sync in the database.
* Needs to be persistent.
* When empty, initialized with the genesis block (empty block with no transactions, no previous block hash).
* Data is not sharded, all blocks are stored locally on all machines.
* Recommended implementation: LevelDB
  * Fast key-value storage library written at Google.
  * Provides an ordered mapping from string keys to string values.

&nbsp;
## `Block Synchronization`

* The synchronization process is initialled upon either:
  * Receving CommitBlock block with height <> last_commited_block block + 1.
  * Receving a request to fetch data from blocks that are unavailable to the block storage. (from the state storage or transactions pool)

## `Block Synchronization Flow`
# TODO redefine flow, do we want to optmize consensus operation? Seperate headers from data? Set designated node algorithmically ?

* Ask all block storage nodes if they have the required blocks.
* newer blocks than mine by calling `Gossip.BroadcastMessage` with "HasNewBlocksMessage" message.
* Choose one of the nodes that replies with new blocks and start a sync process to get all of its new blocks:
  * "HasNewBlocks" message to signify that there are new blocks.
  * "HasNewBlocksResponse" message to signify which node was chosen for sync.
  * "SendNewBlocks" message to request the batch of new blocks.
  * "SendNewBlocksResponse" message to contain the batch of blocks as a response.

&nbsp;
## `AddBlock` (method)
> Add a new block to the database

#### Verify block before adding to database
* Check the block protocol version.
* If the block already exist (block height < last_commited_block block + 1) - sligntly discrad. (may occur during node sync)
* If block height = last_commited_block block + 1, check hash pointer if matches `Commit block` else discard block.
* If block height > last_commited_block block + 1 - out_of_sync, add the to the storage without committing it, initiate `Block Synchronization`.
  * The should be storaed in order to commit it once regain sync.

#### Add block to storage
* Store block in database indexed by block height.

#### Commit block 
* Update last_commited_block block = block height.
* Send StateDiff update to state stroage by calling `StateStorage.CommitStateDiff`.
* Send transactions receipts update to the trasnaction pool by calling `TransactionPool.MarkCommittedTransactions`.

&nbsp;
## `GetBlocksByHeight` (method)
> Return an array of blocks of a range of heights along with the last_commited_block and the last_added_block. (last_commited_block <> last_added_block indicates out of sync).
* If the block range is not present, check that it within the added range by calling `ConsensusCore.GetTopBlock`, if it is, fetch the missing blocks by initiating Block Synchronization Flow. 
* TODO - Checkpoints

&nbsp;
## `GetOrderingBlockByHeight / GetValidationBlockByHeight` (method)
> Similar to `GetBlocksByHeight`, returns only the ordering block / validation block parts.

&nbsp;
## `GetLastBlockHeight` (method)
> Returns the latest committed and addeded block heigths.

&nbsp;
## `ProcessGossipMessage` (method)
> Handles a gossip message from another node. Relevant messages include node sync messages.
/* TODO Block synchonizatiob flow
#### On "HasNewBlocksMessage" message
* Call `HasNewBlocks` to generate a response.
* Respond by calling `Gossip.UnicastMessage` with "HasNewBlocksResponse" message.

#### On "HasNewBlocksResponse" message
* Start a sync process with one of the nodes that has new blocks (only one concurrent process allowed).
* When a new sync process starts, send the node "SendNewBlocks" message by calling `Gossip.UnicastMessage`.

#### On "SendNewBlocks" message
* Call `GetBlocks` to generate a response.
* Respond by calling `Gossip.UnicastMessage` with "SendNewBlocksResponse" message.

#### On "SendNewBlocksResponse" message
* Make sure sync process with the sending node is active.
* Add blocks to database sorted by block height by calling `AddBlock`.

`Sync flow`
* Randomly select one of the other nodes and send it a BLOCK_SYNC_REQUEST message using `Gossip.UnicastMessage` with block_height = top_block + 1.
  * If a response does not arrive within X=5sec, resend to another node.
* Upon receiption of a BLOCK_SYNC_RESPONSE message, perfom regular `BlockCommitted` flow.
* Repeat sending requests until receiving a valid PrePrepare message (including block_height = top_block + 1)

*/

## `GetTransactionReceipt` (method)
> returns a transaction receipt for a given transaction ID and timestamp.
* May use the block header blloom filter for transaction_id search.
* TODO return also a proof.

## `GetReceipts` (method)
> Return an array of Validation Blocks with the receipts, header and proof of a range of heights along with the last_commited_block and the last_added_block. (last_commited_block <> last_added_block indicates out of sync). Used by the the tarnsaction-pool for sync.
* If the block range is not present, check that it within the added range by calling `ConsensusCore.GetTopBlock`, if it is, fetch the missing blocks by initiating Block Synchronization Flow. 

#### Verify  

## `GetStateDiff` (method)
> Return an array of Validation Blocks with the state_diff, header and proof of a range of heights along with the last_commited_block and the last_added_block. (last_commited_block <> last_added_block indicates out of sync). Used by the the tarnsaction-pool for sync.
* If the block range is not present, check that it within the added range by calling `ConsensusCore.GetTopBlock`, if it is, fetch the missing blocks by initiating Block Synchronization Flow. 
