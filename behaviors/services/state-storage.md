# State Storage <!-- tal will finish -->

Holds the latest state under consensus, meaning all of the state variables for all services in a virtual chain.

Currently a single instance per virtual chain per node.

&nbsp;
## `Data Structures`

#### State store
* Stores state data per service by key (key-value store).
* Up to date for a specific block height.
* Maintains a snapshot history of configurable number of blocks (eg. 5).
* Separated into services, for each one:
  * Keep a merkle tree of all state variables (keys and values) under it.
  * Keys are hashes, values are blobs (byte arrays).
  * A non exist key returns a 0 value.
* Default system services:
  * `_Deployments`
    * Contains metadata about every service and library that were deployed on the system.
    * Key: hash(`<namespace>.Processor`), Value: processor for this namespace
      * Meta example: `_Deployments.Processor = Native`
* All service merkle trees are held inside a parent merkle tree
  * Key: hash(service name), Value: merkle tree of the service state variables.
* State root hash is the merkle root of the parent merkle tree.
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

#### Block Height check
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
