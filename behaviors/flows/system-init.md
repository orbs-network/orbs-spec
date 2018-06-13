# System Init Flow

General node initialization, all [services](../../terminology.md) are initialized.

## Participants in this flow

* All services.

## Assumptions for successful flow

* Majority of the nodes (2f+1) are updated to the latest configuration.
* Configuration changes take effect starting from a future block N, majority needs to update before it's reached.

## Flow

* All services:
  * Initialize configuration and persistent data.

* Services that gossip (`ConsensusAlgo`, `BlockStorage`, `TransactionPool`):
  * Subscribe to relevant gossip topics.

* `ConsensusAlgo`:
  * Start the consensus algorithm.

* `BlockStorage`:
  * Attempt block synchronization process from other nodes to see if anyone has new blocks.

* `Gossip`:
  * Connect to the gossip endpoints of all relevant peer nodes.

* `CrosschainConnector`:
  * Start the relevant node.
