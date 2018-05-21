# State Storage

Holds the latest state under consensus, meaning the state variables for all contracts in a virtual chain.

&nbsp;
## `Data Structures`

#### State store
* Stores state data for all contracts by key (key-value store).
* Maintains a snapshot of up to K=5 blocks
  * Potential datastructure: Maintaine by a State Diff Cache in size of X blocks, which is looked up on first on ReadKeys calls.
    * last_committed_block indicates the last block committed to the cache.
    * When block N is committed (to the cache), the state is updated woth block N-K state diff. The atomicity of the state diff cache enables the update of the state not to be atomic.
* Data structure:
  * For each smart contract keep a merkle tree of address, value.
  * Value stored per key is a free-form string. // TODO serialized form.
  * A non exist address returns a 0 value. //TODO
* No need to be persistent.
* Updated with the state diff and top block height by the block storage.
* Upon loosing sync, re-syncs with the block storage.

&nbsp;
## `CommitStateDiff`
> Commits the state diff of a new block

#### Consistency check
* Update local top_block_height = CommitStateDiff.top_block_height
* Check block height:
  * If block height > last_commited_block height + 1, Set OUT_OF_SYNC, initate sync flow.
  * If block height < last_commited_block height + 1 silently discard.
  * If (block height = last_commited_block height + 1) AND (block height = top_block_height), reset OUT_OF_SYNC.

#### Commit state
* Update the state_diff.
  * Update the state merkle tree while updating each state.
* Update the last_committed_block height
* Update the top_block_height

&nbsp;
## `Check in-sync` 
> Verify that the state storage has the available block_height.
  * If block height > last_commited_block height
    * hold response for up to X = 2 sec waiting for potential StateDiff commit from the block storage.
    * Upon timeout return BLOCK_HEIGHT_UNAVIABLE.
  * If block_height < min_available_block_height
    * return BLOCK_HEIGHT_UNAVIABLE.
Note: OUT_OF_SYNC state is triggered only by the block storage updates and isn't triggered by a block height > last_commited_block.

## `ReadKeys` (method)
> Retrieve the values (updated to a certain block height) from a set of keys belonging to a contract. 
> If the requested block height is unavailable returns BLOCK_HEIGHT_UNAVIABLE.

#### Make sure state is in sync
* Performs `Check in-sync` flow.

#### Read the state
* Return the ordered list of values read from the state store.

&nbsp;
## `GetStateHash` (method)
> Returns the state hash (merkle tree root)

#### Make sure state is in sync
* Performs `Check in-sync` flow.

#### Return the state root
* Return the state merkle tree root.

&nbsp;
## `GetStateHash` (method)

&nbsp;
## `Sync Flow`
> Syncs the state storage based on the block storage state diff.
* Call `BlockStorage.RequestStateDiffUpdate` with consumer_block_height = last_commited_block + 1, target_block_height = MAX_UINT64.