# State Storage

Holds the latest state under consensus, meaning the state variables for all contracts in a virtual chain.

&nbsp;
## `Data Structures`

#### State store
* Stores state data for all contracts by key (key-value store).
* Data structure:
  * For each smart contract keep a hash table of address, value. 
* Value stored per key is a free-form string.
* No need to be persistent.
* Updated with the state diff and top block height by the Consensus. TBD - block storage.

## `CommitStateDiff`

#### Consistency check
* Check block height = top block height + 1. If not - fatal error. (Need to revert to a checkpoint)

#### Commit state
* Update all the state_diff.
  * Update the state merkle tree.
* When done update the top block height


&nbsp;
## `ReadKeys` (method)
> Retrieve the latest state values from a set of keys belonging to one contract

#### Make sure state is fully synced
* Current state block height is equal to block storage latest block (by calling `BlockStorage.GetLastBlock`).
* Allow 5 second timeout for sync to complete.

#### Read the state
* Directly from key value store.


&nbsp;
## `GetStateHash` (method) //TBD - return to Consensus in CommitStateDiff (guarantees consistency)

#### Check State Storage is updated to the latest block
Check that block_height is consistent with current block_height, otherwise return status = NOT_UPDATED.

#### Return the state root
Return the state merkle tree root hash.