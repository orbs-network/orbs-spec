# State Storage

Holds the latest state under consensus, meaning all of the state variables for all [deployed services](../../terminology.md) in a virtual chain. Provides the source of truth inside the node for the state snapshot.

Currently a single instance per virtual chain per node.

#### Interacts with services

* None - Passive, just provides services to others upon request.

&nbsp;
## `Data Structures`

#### State store
* Stores state data per [deployed service](../../terminology.md) by key (key-value store).
* Up to date for a specific [block height](../../terminology.md).
* Maintains an efficient snapshot history of [configurable](../config/services.md) number of past blocks (eg. 5).
* Separated into deployed services, for each one:
  * Keep a merkle tree of all state variables (keys and values) under it.
  * Keys are hashes, values are blobs (byte arrays).
  * A non existent key returns an empty byte array. (implies null value / integer `0` value)
* Default system services:
  * `_Deployments`
    * Contains metadata about every [service](../../terminology.md) and [library](../../terminology.md) smart contracts that were deployed on the system.
    * The format is Key: hash(`<name>.Processor`), Value: processor for this smart contract (service or library).
      * Meta example: `_Deployments.Processor = Native`.
* All service merkle trees are held inside a parent merkle tree.
  * Key: hash(service name), Value: merkle tree of the service state variables.
* The state root hash is the merkle root of the parent merkle tree.
* Can be persistent to optimize service bootstrap (be careful of partial writes).

#### Synchronization state
* `last_committed_block` - The last valid committed block that the state storage is synchronized to (can be persistent).

&nbsp;
## `Init` (flow)

* Initialize the [configuration](../config/services.md).
* Load persistent data.
* If no persistent data, init `last_committed_block` to empty (symbolize the empty genesis block) and the state store to empty.

&nbsp;
## `CommitStateDiff` (method)

> Commits a new approved block (the state diff from it is what we care about). This is the main way to update the current state snapshot as new blocks are generated and approved in the system.

#### Check block height
* We assume here that the caller of this method inside the node is trusted and has already verified the block.
* We want to commit blocks in sequence, so make sure the given [block height](../../terminology.md) is the next of `last_committed_block`.
* If not, discard the commit and return the next desired block height (which is the next of `last_committed_block`).

#### Commit state
* Update the state store according to the block's `state_diff`.
  * Also update the state merkle tree while updating each state value.
  * When setting the zero value (an empty byte array), delete the entry from both state and merkle tree.
* Update `last_committed_block` to match the given block.

&nbsp;
## `ReadKeys` (method)

> Retrieve the values (updated to a certain block height) from a set of keys belonging to a contract. This is the main way for other services to query state in the node.

#### Check synchronization status
* If requested block height is in the future but `last_committed_block` is close to it ([configurable](../config/services.md) sync grace distance) block the call until requested height is committed. Or fail on timeout.
* If requested block height is in the future but `last_committed_block` is far, fail.
* If requested block height is in the past and beyond the snapshot history, fail.

#### Return the values
* If requested contract name does not exist in storage, fail.
* Otherwise, respond with the values from the state store.
* For any requested key not stored return the zero value (an empty byte array)

&nbsp;
## `GetStateStorageBlockHeight` (method)

> Returns the last committed height and timestamp, called by the Virtual Machine as part of the CallMethod flow.

* Return the height and timestamp of `last_committed_block`.

&nbsp;
## `GetStateHash` (method)

> Returns the state hash (merkle tree root) updated to a certain block height. The latest hash is always written inside blocks so this is needed by block creators.

#### Check synchronization status
* If requested block height is in the future but `last_committed_block` is close to it ([configurable](../config/services.md) sync grace distance) block the call until requested height is committed. Or fail on timeout.
* If requested block height is in the future but `last_committed_block` is far, fail.
* If requested block height is in the past and beyond the snapshot history, fail.

#### Return the state root
* Respond with the merkle root from the state store.
