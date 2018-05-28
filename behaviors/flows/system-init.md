# System Init Flow

On system Init, the node sychronizes to its local persistant database and the network.

## Participants in this flow

* Persistent services
  * `Block Storage`
  * `Consensus Algo`
  * `State Storage`

* Stateful services
  * `State Storage`
  * `Transaction Pool`
  * `SideChain Connector` 
  * `Consensus Core`

* Configuration files
  * `Consensus Core`
  * `Gossip`

## Assumptions
* 2f+1 of the network are synched to the latest configurations.
  * Multiple configuration changes require to commit blocks between them such that it's gurnteed that 2f+1 nodes have applied the new configuration.

## General
* Configuration changes take effect starting on block N.
 

## `Consensus Core`
* Priodically query the `Block Storage` until receiving 


## `Consensus Algo`

#### Persistent database
* Maintains a persistent database of algorithm required persistent data: Prepare database of latest term.

#### Init Flow
* Subscribe to Consensus `Gossip` messages.
* Read Consnesus persistent data.



* Attempts to receive Consensus Commit messages.
  * Upon reception commits block to block storage
* Wait configurable max 


## `Block Storage`

#### Persistent database
* Maintains a persistent database of:
  * Blocks (in-order and out-of-order committs)
  * Last committed height 

#### Init Flow
* Subscribe to Node Sync `Gossip` messages.
* Initiste inter-node sync flow 
  * Until no CommitBlock was received, set `max_block_height` = MAX_UINT64, indicating all valid blocks.
* 
