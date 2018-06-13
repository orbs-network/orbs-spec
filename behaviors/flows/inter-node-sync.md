# Continuous Inter Node Block Synchronization Flow

`BlockStorage` always needs be synchronized to the latest block in the virtual chain (latest committed block height). As long as it's synchronized, the node will be able to participate in consensus and the `ConsensusAlgo` will keep committing blocks to it.

`BlockStorage` can identify when it is potentially out of sync if no new blocks are committed successfully for some time. When this happens, it attempts to synchronize from other `BlockStorage` instances of other nodes. Note that these instances aren't necessarily trusted, so full block validations need to happen on every block they provide.

## Participants in this flow

* Primary
  * `BlockStorage`

* Helpers
  * `Gossip`
  * `ConsensusAlgo`

## Assumptions for successful flow

* Other nodes in the network are more advanced and have consensus over later blocks.

## Flow

* `BlockStorage` of the original node:
  * Identifies that it out of sync (no blocks are committed for some time).
  * Broadcasts a sync request to all nodes with `Gossip`.

* `BlockStorage` of all nodes willing to help respond if they have missing blocks.

* `BlockStorage` of the original node:
  * Chooses randomly one of the nodes to synchronize with.
  * Starts a batched synchronization process with it through `Gossip`.
  * Validates the consensus of every untrusted block it receives with `ConsensusAlgo`.
