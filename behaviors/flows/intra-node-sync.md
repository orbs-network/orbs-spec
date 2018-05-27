# Continuous Intra Node Block Synchronization Flow <!-- tal can finish -->

TransactionPool

TransactionPool needs to always be synchronized to the latest block in the virtual chain (latest committed block height). It can identify when it is potentially out of sync when another notifies it of a later block.

When TransactionPool is out of sync, it attempts to synchronize with BlockStorage inside the node.

Why synchronized? because needs to remember configurable time (like 30 minutes) to prevent duplication and follow ordering policy.

## Participants in this flow

## Assumptions for successful flow

* `BlockStorage` is synchronized to latest block.

## Flow

* `BlockStorage` commits new blocks to services and specifies its `last_committed_block_height` (not necessarily the block being committed).
* Service handles the commit and responds with its `last_committed_block_height` and its `next_desired_block_height`.
* If `BlockStorage` has the `next_desired_block_height`, it commits it to the service.
* This continues in endless loop until service is synchronized.
