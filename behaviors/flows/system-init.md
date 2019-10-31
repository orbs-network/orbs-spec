# System Init Flow

General node initialization, all [services](../../terminology.md) are initialized.

## Participants in this flow

* All services.

## Assumptions for successful flow

* Majority of the nodes (Quorum) are updated to the latest configuration.
* Initial configuration "honestly" holds until it is overwritten. Configuration changes occur under consensus - state update. 

## Flow

* All services:
  * Initialize configuration and persistent data.

* System Contracts Deployment:
  * Different system contracts are deployed based on configuration.
    
* Services that gossip (`ConsensusAlgo`, `BlockStorage`, `TransactionPool`):
  * Subscribe to relevant gossip topics.

* `ConsensusAlgo`:
  * Start the consensus algorithm.

* `BlockStorage`:
  * Attempt block synchronization process from other nodes to see if anyone has new blocks.

* `Gossip`:
  * Connect to the gossip endpoints of all relevant peer nodes.

* `CrosschainConnector`:
  * Establish the appropriate connection (Ethereum or ManagementVC).
    * Ethereum - Start the relevant node.
