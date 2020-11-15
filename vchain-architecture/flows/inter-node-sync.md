# Continuous Inter Node Block Synchronization Flow

`BlockStorage` must be synchronized to the latest block in the virtual chain (latest committed block height). As long as it's synchronized, the node can participate in consensus, and the `ConsensusAlgo` keeps committing blocks to it.

`BlockStorage` can identify when it is potentially out of sync if no new blocks are committed successfully for some time. When this happens, it attempts to synchronize from other `BlockStorage` instances of other nodes.
 Note under Orbs PoS v2 model it is assumed the participants are honest for 12hrs after initiating an exit flow. that these instances aren't necessarily trusted, so full block validations need to happen on every block they provide.

## Participants in this flow

* Primary
  * `BlockStorage`

* Helpers
  * `Gossip`
  * `ConsensusAlgo`
  * `ConsensusContext`

## Assumptions for a successful flow

* Other nodes in the network are more advanced and have consensus over later blocks.

## Flow

* `BlockStorage` of the out of sync node:
  * Identifies that it out of sync (no blocks are committed for some time).
  * Broadcasts a sync request to all nodes with `Gossip`.

* `BlockStorage` of all nodes willing to help respond if they have missing blocks.
  * Negotiate **batch** size (total number of blocks the syncing node gives in this session).

* `BlockStorage` of the out of sync node:
  * Chooses randomly one of the nodes to synchronize with.
  * Starts a batched synchronization process with it through `Gossip`.
  * Negotiate **chunk** size (number of blocks sent together as a chunk during the streaming process of the full batch).
  * Validates the consensus of every untrusted block it receives with `ConsensusAlgo`.

  * `ConsensusAlgo` of the out of sync node:
    * Gets the block's committee from `ConsensusContext` and verifies the block proofs.
