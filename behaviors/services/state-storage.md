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
## `ReadKeys` (method)
> Retrieve the values from a set of keys belonging to one contract.

#### Make sure state is fully synced
* Check block height = last_commited_block height.
  * If block height = last_commited_block height + 1
    hold response for up to X = 2 sec waiting for potential StateDiff commit from the block storage.
  * If block height > last_commited_block height + 1 or upon timeout, return OUT_OF_SYNC response and initiate a `Sync Flow`

#### Read the state
* Return the ordered list of values read from the state store.

&nbsp;
## `GetStateHash` (method)
> Returns the state hash (merkle tree root)

#### Check State Storage is updated to the latest block
Check that block_height is consistent with current block_height, otherwise return status = NOT_UPDATED.

#### Return the state root
Return the state merkle tree root.