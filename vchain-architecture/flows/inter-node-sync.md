# Continuous Inter Node Block Synchronization Flow

`BlockStorage` must be synchronized to the latest block in the virtual chain (latest committed block height). As long as it's synchronized, the node can participate in consensus, and the `ConsensusAlgo` keeps committing blocks to it.

`BlockStorage` can identify when it is potentially out of sync if no new blocks are committed successfully for some time. When this happens, it attempts to synchronize from other `BlockStorage` instances of other nodes.
The main VC fork can be deduced according to Orbs DPoS v2 model on Ethereum. By obtaining the current elected committee the node can verify the blockchain tip  against it. Note, it is assumed the committee members remain honest for 12hrs after initiating an exit flow (this assumption is reassured by the staking mechanism on Ethereum includes a cooling period).
Each block holds a reference to Ethereum time, indicating the next block assigned committee which attests for its consensus proof.
To obtain the relevant committee members for each block Ethereum reference, the node maintains a proxy to Ethereum (management service), which continuously aggregates the information obtained from Ethereum and acts as the PoS source of truth.

## Participants in this flow

* Primary
  * `BlockStorage`

* Helpers
  * `Gossip`
  * `ConsensusAlgo`
  * `ConsensusContext`

## Assumptions for a successful flow

* Other nodes in the network are more advanced and have consensus over later blocks.
* Either the Ethereum reference in the newest block is within honesty grace (12hrs of time.now) or the current committee has not changed significantly. Implying one can guarantee honest attestation for the block.
* There can be only a single valid fork which conforming to the consensus rules. This fork is held by at least one honest node and can be obtained from it.


## Flow

The `BlockStorage` performs the sync in a reverse order - from tip to genesis.
* The sync progress state is maintained using three indices: 
 * InOrder - indicates the length of consecutive blocks from genesis with no gaps
 * LastSynced - temporary, indicating the latest block height committed to block storage
 * Top - the current blockchain tip, indicating the max block height of committed blocks

* `BlockStorage` of the out of sync node:
  * Identifies that it out of sync (no blocks are committed for some time).
  * If there exists a gap (lastSynced is greater than inOrder) close it first.
  * Broadcasts a sync request to all nodes with `Gossip`.

* `BlockStorage` of all nodes willing to help respond if they have missing blocks.
  * Negotiate **batch** size (total number of blocks the syncing node gives in this session).

* `BlockStorage` of the out of sync node:
  * Chooses randomly one of the nodes to synchronize with.
  * Starts a batched synchronization process with it through `Gossip`.
  * Negotiate **chunk** size (number of blocks sent together as a chunk during the streaming process of the full batch).
  * Validates the consensus of every untrusted block it receives with `ConsensusAlgo`.

  * `ConsensusAlgo` of the out of sync node:
    * If syncing the tip - verify Ethereum reference is withing 12hrs of time.now or there exists honest stake (in current committee) which attest for this block.
    * Otherwise, gets the block's committee from `ConsensusContext` and verifies the block proofs.
