# Block Storage

Holds the long term journal of all confirmed blocks. Provides the source of truth inside the node for the current [block height](../../terminology.md). When the node is synchronized with the rest of the network, newly committed blocks arrive internally from `ConsensusAlgo`. When not, the service manually synchronizes the missing blocks from other nodes.

Currently a single instance per virtual chain per node.

#### Interacts with services

* `StateStorage` - Provides it with state diffs when new blocks are committed (also synchronizes it).
* `TransactionPool` - Provides it with transaction receipts when new blocks are committed (also synchronizes it).
* `ConsensusAlgo` - Asks it whether untrusted blocks (given by other nodes) are indeed approved and valid.
* `Gossip` - Uses it to communicate with other nodes.

&nbsp;
## `Data Structures`

#### Block database
* Holds all blocks with ability to query efficiently by [block height](../../terminology.md).
* Able to store blocks out of order (when sync is lost but future blocks keep arriving).
  * These blocks are not yet marked as committed and cannot be considered valid.
  * They will be verified and committed after receiving the chain of missing blocks before them.
* When empty, initialized with the genesis block which is empty for virtual chains.
* The database needs to be persistent.
* Data is not sharded, currently all blocks are stored locally on all machines.

#### Synchronization state
* `last_committed_block` - The last valid committed block that the block storage is synchronized to (persistent).

&nbsp;
## `Init` (flow)

* Initialize the [configuration](../config/services.md).
* Load persistent data.
* If there is persistent data, init `last_commited_block` accordingly.
* If no persistent data, init `last_committed_block` to empty (symbolize the empty genesis block) and the block database to empty.
* When consensus algorithms register by calling `BlockStorage.RegisterConsensusBlocksHandler`, update them with the latest persistent block by calling their `HandleBlockConsensus` with `HANDLE_BLOCK_CONSENSUS_MODE_UPDATE_ONLY`.
  * If `last_committed_block` is empty pass an empty block.
* Subscribe to gossip messages by calling `Gossip.BlockSync.RegisterBlockSyncHandler`.
* Trigger the block synchronization process in the `Inter Node Block Sync` flow.
  * Indicating update only mode.

&nbsp;
## `Inter Node Block Sync` (flow)

> Continuous flow where block storage (the source of truth for committed blocks) is synchronizing itself from other nodes when it discovers it is out of sync.

#### Trigger
* Endless loop which is continuously waiting for a trigger marking the node as out of sync.
* The synchronization process is triggered upon:
  * Too much time has passed without a commit with `CommitBlock`. Based on a [configurable](../config/services.md) timeout (eg. 8 sec).
  * During the Init flow.
* Make sure no more than one synchronization process is active at any given time.

#### Pre-synchornization
* When synchronization is triggered, update all registered consensus algos with the latest persistent block by calling their `HandleBlockConsensus` with `HANDLE_BLOCK_CONSENSUS_MODE_UPDATE_ONLY`.
  * If `last_committed_block` is empty pass an empty block.

#### Synchronization process
* Synchronization is made of multiple **batches** each comprised of multiple **chunks**.
* Identify nodes that have the desired blocks by broadcasting `BLOCK_AVAILABILITY_REQUEST` message with `Gossip.BlockSync.BroadcastBlockAvailabilityRequest`.
  * Request blocks starting from `last_committed_block`.
  * Request [configurable](../config/services.md) **batch** range for this session (total number of blocks the syncing node will give in this session).
  * Nodes will respond with a `BLOCK_AVAILABILITY_RESPONSE` message.
    * Indicating the range from the desired blocks that is available to them and their current last committed block.
    * Also indicating the agreed upon [configurable](../config/services.md) **batch** range for this session.
* Randomly select one of the responding nodes and request the **batch** of blocks from them.
  * If some of the responders are significantly behind and can't fulfill one batch, avoid them in the random selection.
  * Request the first [configurable](../config/services.md) **chunk** in the batch by sending a `BLOCK_SYNC_REQUEST` message with `Gossip.BlockSync.SendBlockSyncRequest`.
  * The receiving node will respond with a `BLOCK_SYNC_RESPONSE` message.
    * The response includes one **chunk** of requested blocks and the node's `last_committed_block` height.
  * Upon reception of the response:
    * Since the blocks are untrusted, validate each of them for commit by calling `ValidateBlockForCommit`.
    * If valid, commit each block by calling `CommitBlock`.
  * Repeat the process until all **chunks** in the batch have arrived (from one syncing node).
* Continue having sessions for more **batches** until all desired blocks have arrived.
  * It is recommended to shuffle the batches among different nodes to avoid putting too much burden on a single node.

&nbsp;
## `Intra Node Block Sync` (flow)

> Continuous flow where block storage (the source of truth for committed blocks) is updating the rest of the services in the node on newly committed blocks.

* This flow is instantiated twice:
  * One instance for the service `TransactionPool` and one for `StateStorage`.
* Endless loop which is continuously synchronizing the service with committed blocks based on the service's next desired block height.
* If `last_commited_block` is lower than the service's next desired block height, block and wait (without polling), else:
  * Push the next desired block to the service:
    * To transaction pool by calling `TransactionPool.CommitTransactionsReceipts`.
    * To state storage by calling `StateStorage.CommitStateDiff`.
  * Remember the service's next desired block height according to the call response.
  * Continue iterating.
* Important: When everything is synchronized (common case), `CommitBlock` should immediately trigger the commit to the service without waiting (this is in the critical path).

&nbsp;
## `CommitBlock` (method)

> Commit a new trusted block (pair, Transactions and Results) after it was approved for commit. Used by the critical path in the consensus algo to commit verified blocks.

#### Final checks before adding
* We assume here that the caller of this method inside the node is trusted and has already done the tests specified by `ValidateBlockForCommit`.
* Correct block protocol version.
* Silently discard the given block if it already exists in the database (panic if it's different from ours under this block height).
* If it doesn't exist, panic if the given block height isn't the next of `last_committed_block`.

#### Commit the block
* Store block in database indexed by block height.
* Update `last_committed_block` to match the given block.
* If any of the intra block syncs (`StateStorage`, `TransactionPool`) is blocking and waiting, wake it up.

&nbsp;
## `ValidateBlockForCommit` (method)

> Validates an untrusted block (pair, Transactions and Results) received during inter node sync before it can be committed.

#### Check the Transactions block (stateless)
* Correct block protocol version.
* Correct virtual chain.
* Check block height:
  * If the block isn't the next of `last_commited_block` according to height, discard.
* Check the block's `transactions_root_hash`:
  * Calculate the merkle root hash of the block's transactions and verify the hash in the header.
* Check the block's metadata hash:
  * Calculate the hash of the block's metadata and verify the hash in the header.

#### Check the Results block (stateless)
* Correct block protocol version.
* Correct virtual chain.
* Check block height:
  * If the block isn't the next of `last_commited_block` according to height, discard.
* Check the block's `receipts_root_hash`:
  * Calculate the merkle root hash of the block's receipts and verify the hash in the header.
* Check the block's `state_diff_hash`:
  * Calculate the hash of the block's state diff and verify the hash in the header.

*Note: The logic up to here also appears in `ConsensusAlgo` and should probably be extracted to avoid duplication.*

#### Check the block consensus
* Identify the relevant service (consensus algorithm) that is registered to handle this block's validation.
  * Check consensus of the block by calling its `HandleBlockConsensus` with `HANDLE_BLOCK_CONSENSUS_MODE_VERIFY_AND_UPDATE`.

&nbsp;
## `GetLastCommittedBlockHeight` (method)

> As the source of truth in the node, returns the height of last committed block.

* Return the height of `last_committed_block`.

&nbsp;
## `GetTransactionReceipt` (method)

> Returns the transaction receipt for a past transaction based on its id and time stamp. Used when a client asks to query transaction status for an older transaction.

* Go over all the blocks where the transaction could be found:
  * Starting where block timestamp is transaction timestamp minus [configurable](../config/services.md) small grace (eg. 5 sec).
  * Finishing where block timestamp is transaction timestamp plus [configurable](../config/shared.md) transaction expiration window (eg. 30 min) plus small grace (eg. 5 sec).
* For each relevant block, look for the transaction in the block header's bloom filters:
  * The transaction timestamp in the timestamp bloom filter (see block format for structure).
  * The transaction `txhash` in the txhash bloom filter (see block format for structure).
  * On match, fetch the block and search for the `txhash` in the block receipts.
    * If found, returns the receipt along with its block height and timestamp.
* If not found on all relevant blocks, returns an empty receipt along with the last committed block height and timestamp.

&nbsp;
## `GetTransactionsBlockHeader` (method)

> Returns a committed Transactions block header and proof given a block height. Used primarily by the consensus algo when it's missing a block.

* If requested block height is in the future but `last_committed_block` is close to it ([configurable](../config/services.md) sync grace distance) block the call until requested height is committed. Or fail on [configurable](../config/shared.md) timeout.
* If requested block height is in the future but `last_committed_block` is far, fail.
* Return the transactions block header, metadata and the transactions block proof.

&nbsp;
## `GetResultsBlockHeader` (method)

> Returns a committed Results block header and proof given a block height. Used primarily by the consensus algo when it's missing a block.

* If requested block height is in the future but `last_committed_block` is close to it ([configurable](../config/services.md) sync grace distance) block the call until requested height is committed. Or fail on [configurable](../config/shared.md) timeout.
* If requested block height is in the future but `last_committed_block` is far, fail.
* Return the results block header, metadata and the results block proof.

&nbsp;
## `GenerateReceiptProof` (method)

> Generates a proof for a receipt inclusion in a block. Returns the transaction receipt for a past transaction based on its id and time stamp, along with the signed block header and receipt merkle proof.

* Get the relevant block and look for the receipt that matches the `txhash`.
  * If no matching receipt was found, return an empty proof. 
* Calculate the receipt merkle proof based on the receipt index
  * Calculate the merkle tree of the block's receipts, using the receipts [Merkle tree format](../data-structures/merkle-tree.md).
    * Consider to maintain a cache of recently calculated merkle trees.
  * Generate a proof for the receipt inclusion based on its index in the block.
* Return the `ResultBlock` header, `ResultBlock` proof, merkle proof and receipt.

&nbsp;
## Gossip Messages Handlers

> Handles gossip messages from other nodes. Relevant messages include block sync messages.

#### `HandleBlockAvailabilityRequest`
* Handles `BLOCK_AVAILABILITY_REQUEST` messages, see `Inter Node Block Sync` flow.

#### `HandleBlockAvailabilityResponse`
* Handles `BLOCK_AVAILABILITY_RESPONSE` messages, see `Inter Node Block Sync` flow.

#### `HandleBlockSyncRequest`
* Handles `BLOCK_SYNC_REQUEST` messages, see `Inter Node Block Sync` flow.

#### `HandleBlockSyncResponse`
* Handles `BLOCK_SYNC_RESPONSE` messages, see `Inter Node Block Sync` flow.



<!--

TODO: oded, add the diagrams again

![alt text][block_state_pool_flow] <br/><br/>

[block_state_pool_flow]: block_state_pool_flow.png "Block Storage - State Storage / Transaction Pool"
-->
