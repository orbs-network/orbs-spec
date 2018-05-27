# Continuous TransactionPool Synchronization Flow

TransactionPool needs to always be synchronized to the latest block in the virtual chain (latest committed block height). It can identify when it is potentially out of sync when another notifies it of a later block.

When TransactionPool is out of sync, it attempts to synchronize with BlockStorage inside the node.

## Participants in this flow

## Flow

* The `Transaction Pool` subscribes for receipts commits from the `Block Storage`.
* Any query that is initiated by the `Consensus Core` refers to a block height.
* The `Transaction Pool` verifies that it is synced with the block height in the query.
* If the `Transaction Pool` identifies that it is out of sync, it updates its request subscription and requests blocks starting after its last synched height.
* The `Block Storage` commits the required receipts up to the latest block and continues to commit receipts on each new block commit.
