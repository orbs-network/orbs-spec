# Continuous StateStorage Synchronization Flow

StateStorage needs to always be synchronized to the latest block in the virtual chain (latest committed block height). It can identify when it is potentially out of sync when another notifies it of a later block.

When StateStorage is out of sync, it attempts to synchronize with BlockStorage inside the node.

## Participants in this flow

## Flow

* The `State Storage` subscribes for state diff commits from the `Block Storage`.
* Any query that is initiated by the `Consensus Core` refers to a block height.
* The `State Storage` verifies that it is synced with the block height in the query.
* If the `State Storage` identifies that it is out of sync, it updates its request subscription and requests blocks starting after its last synched height.
* The `Block Storage` commits the required state diff up to the latest block and continues to commit the state diff on each new block commit.
