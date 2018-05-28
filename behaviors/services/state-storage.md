# State Storage <!-- tal will finish -->

Holds the latest state under consensus, meaning all of the state variables for all services in a virtual chain.

Currently a single instance per virtual chain per node.

&nbsp;
## `Data Structures`

#### State store
* Stores state data per service by key (key-value store).
* Up to date for a specific block height.
* Maintains a snapshot history of configurable number of blocks (eg. 5).
* Data structure:
  * For each service keep a merkle tree of address, value.
  * Keys and values are blobs (byte arrays).
  * A non exist address returns a 0 value.
* Can be persistent to optimize service bootstrap (be careful of partial writes).

#### Sync state
* `last_committed_block`
* `next_desired_block_height`

&nbsp;
## `Init` <!-- oded will finish -->

TODO

&nbsp;
## `CommitStateDiff` <!-- tal will finish -->

> Commits a new block (the state diff from it is what we care about).

#### Consistency check
* If given block_height != `next_desired_block_height` discard and return `next_desired_block_height`

#### Commit state
* Update the state_diff.
  * Update the state merkle tree while updating each state.
* Update the `last_committed_block` height
* Increment the `next_desired_block_height` and return it

&nbsp;
## `ReadKeys` (method) <!-- tal will finish -->

> Retrieve the values (updated to a certain block height) from a set of keys belonging to a contract.

#### Check synchronization status
* If requested block height is in the future but `last_committed_block` is close to it (configurable distance) block and wait
* If requested block height is in the future but `last_committed_block` is far, fail
* If requested block height is in the past and beyond the snapshot history, fail

#### Return the values
* Respond with the values

&nbsp;
## `GetStateHash` (method) <!-- tal will finish -->

> Returns the state hash (merkle tree root) updated to a certain block height.

#### Check synchronization status
* If requested block height is in the future but `last_committed_block` is close to it (configurable distance) block and wait
* If requested block height is in the future but `last_committed_block` is far, fail
* If requested block height is in the past and beyond the snapshot history, fail

#### Return the state root
* Return the state merkle tree root.
