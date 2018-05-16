# State Storage

Holds the latest state under consensus, meaning the state variables for all contracts in a virtual chain.

&nbsp;
## `Data Structures`

#### State store
* Stores state data for all contracts by key (key-value store).
* Data structure:
  * For each smart contract keep a merkle tree of address, value.
* Value stored per key is a free-form string. // TODO serialized form.
* No need to be persistent.
* Updated with the state diff and top block height by the block storage.
* Upon loosing sync, re-syncs with the block storage.

&nbsp;
## `CommitStateDiff`
> Commits teh state diff of a new block

#### Consistency check
* Check block height = last_commited_block height + 1. 
* If block height > last_commited_block + 1 initate sync flow.
* If block height < last_commited_block + 1 silently discard.

#### Commit state
* Update the state_diff.
  * Update the state merkle tree while updating each state.
* When done update the last_committed_block height

&nbsp;
## `Check in-sync` 

## `ReadKeys` (method)
> Retrieve the values from a set of keys belonging to one contract.
* Update last_requested_block = block height.
* Check block height = last_commited_block height.  
  * If equal reset OUT_OF_SYNC, stop the `Sync flow`.
* If block height = last_commited_block height + 1
  * hold response for up to X = 2 sec waiting for potential StateDiff commit from the block storage.
  * Upon timeout return OUT_OF_SYNC response and initiate a `Sync Flow` starting with last_commited_block height + 1.
* If block height > last_commited_block height + 1:
  * return OUT_OF_SYNC response and initiate a `Sync Flow` starting with last_commited_block height + 1.
* If block height < last_commited_block + 1 return UNEXEPCETED_BLOCK_HEIGHT.

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
## `Sync Flow`
> Syncs the state storage based on the block storage state diff.
* Call `BlockStorage.RequestStateDiffUpdate` with consumer_block_height = last_commited_block + 1, target_block_height = MAX_UINT64.