# State Storage

Holds the latest state under consensus, meaning the state variables for all contracts in a virtual chain.

&nbsp;
## `Data Structures`

#### State store
* Stores state data for all contracts by key (key-value store).
* Value stored per key is a free-form string.
* Each contract has its own isolated key address space.
* No need to be persistent.
* Constantly synchronizing with block storage to reflect latest state.
  * Synchronization by continuously polling block storage for new blocks.
  * Poll interval is configurable.
  * New blocks fetched by calling `BlockStorage.GetBlocks`.
  * When new block found, its state diff is merged into the state.

&nbsp;
## `ReadKeys` (method)
> Retrieve the latest state values from a set of keys belonging to one contract

#### Make sure state is fully synced
* Current state block height is equal to block storage latest block (by calling `BlockStorage.GetLastBlock`).
* Allow 5 second timeout for sync to complete.

#### Read the state
* Directly from key value store.
