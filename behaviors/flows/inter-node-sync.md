# Continuous Inter Node Block Synchronization Flow <!-- tal can finish -->

BlockStorage needs to always be synchronized to the latest block in the virtual chain (latest committed block height). It can identify when it is potentially out of sync when another notifies it of a later block.

When BlockStorage is out of sync, it attempts to synchronize with other BlockStorage instances of other nodes.

BlockStorage is also charged with synchronizing StateStorage and TransactionPool to the latest block. Whenever they go out of sync, they attempt to synchronize with BlockStorage inside the node.

## Participants in this flow

## Assumptions for successful flow

* `ConsensusCore` is synchronized to latest block (and is the source of truth for block height).

## Flow

* `ConsensusCore` commits newly created blocks to `BlockStorage` (its stateless, they come from consensus algorithm).
* `BlockStorage` identifies that it is missing blocks (out of sync).
* `BlockStorage` sends block availablity request message to all nodes with Gossip with desired block range and block type (one configurable batch).
* Any node willing to help synchronize responds.
* Randomly selects a node.
* Requests the batch with message block sync request through Gossip (unicast).
* The node sends a batch of blocks.
* `BlockStorage` does `BlockStorage.AddBlock` (one flow) on these blocks.
* This continues in endless loop until synchronized.
