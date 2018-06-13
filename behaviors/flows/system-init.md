# System Init Flow

On system Init, the node sychronizes to its local persistant database and the network.

## Participants in this flow:
<!--
* Persistent services
  * `Consensus Algo`
  * `Block Storage`
  * `State Storage`

* Non-persistent services
  * `Transaction Pool`
  * `Gossip`

* Stateless services
  * `Consensus Context`
  * `Virtual Machine`
  * `Processors`
  * `Public API`
-->

## Assumptions
* 2f+1 of the network are synched to the latest configurations.
  * Multiple configuration changes require to commit blocks between them such that it's gurnteed that 2f+1 nodes have applied the new configuration.

## General
* Configuration changes take effect starting on block N.


## `Consensus Algo`

* Initialize configuration and persistent data.
* Subscribe to relevant gossip topics.
* Start the consensus algorithm.

## `Block Storage`

* Initialize configuration and persistent data.
* Subscribe to relevant gossip topics.
* Attempt block synchronization process from other nodes to see if anyone has new blocks.

## `State Storage`

* Initialize configuration and persistent data.

## `Transaction Pool`

* Initialize configuration and persistent data.
* Subscribe to relevant gossip topics.

## `Gossip`

* Initialize configuration and persistent data.
* Connect to the gossip endpoints of all relevant peer nodes.
