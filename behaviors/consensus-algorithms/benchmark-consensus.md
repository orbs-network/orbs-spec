# Benchmark Consensus Spec
> Algorithm with constant leader used for performance benchmarks, should not be used in production environments since it is not fully decentralized.

## Overview
* One node is designated as leader and will constantly hold this role.
* All other nodes are familiar with the identity of the leader.
* Leader announces which blocks are committed to all other nodes and waits until enough of them are synchronized.
* Leader does not rely on other nodes for validation purposes.

## Leader flow
* Broadcast the last committed block as committed to all nodes with a `COMMIT` message.
* Wait until a quorum of other nodes answer with valid (signed) `COMMITTED` messages (to acknoledge that the block was processed).
* Propose a new block (Transactions and Results).
* Commit the block locally.
* Repeat the process.

## Non-leader flow
* Wait for a valid (signed) `COMMIT` message from the leader.
* If this is the next expected block height, commit the block locally.
* Reply with a `COMMITTED` message to the leader (even if this was not the next expected block height).
* Repeat the process.

## Quorom size
* The larger the quorom, the more risk of liveness issues due to nodes not answering.
* Quorom that is too small will skip synchronizing too many nodes and create fragmentation in the network and reduce performance.
* Recommended quorom size is 2/3 of the nodes.

## Synchronization
* The `COMMITTED` messages include the `last_committed_block_height` of the sender.
* The leader will only move to the next block after a quorom confirms that they are synchronized to the current block.
* The leader will ignore nodes that claim to be synchronized to future block heights or past block heights.
* Nodes that are not synced with the leader will have to rely on regular block sync (from Block Storage) to catch up.

## Liveness
* Leader may become stuck if it does not receive enough `COMMITTED` replies or fails on some internal error.
* In this case, the leader will broadcast its last committed block once again and continue the flow from this point.

## Initialization
* The leader starts the process by initializing its last committed block to a genesis block (fake empty block at height 0 with all proper signatures).
* Nodes that receive a `COMMIT` for the genesis block (height 0) will not commit it locally but will still reply with a `COMMITTED` message.
* Once the Block Storage of the leader initializes, if it contains persistent blocks, it will notify the Consensus Algo upon registration about the last persistent block.
  * This is the common case flow of system starts where the leader has existing persistent blocks in storage.
  * In this case the leader consensus algo initializes with last committed 0 and tries to propose block height 1.
  * While saving this proposal 1 to block storage, block storage will fail since a different block already exists with height 1.
  * This process will continue looping endlessly (more proposals of block height 1 that will keep failing during save).
  * Block sync will eventually trigger (either during startup or because a few seconds without successful commits passed).
  * During block sync pre-synchronization, block storage will notify consensus algo of the latest persistent block.
  * Once consensus algo is up to date with the latest persistent block, it can propose new blocks, save them and move forward.
