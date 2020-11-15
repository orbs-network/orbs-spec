# Continuous Block Creation Flow

The purpose of consensus is to continuously commit new blocks. Blocks are populated with transactions waiting in `TransactionPool`.

We currently support leader-based consensus algorithms, meaning a leader [node](../../terminology.md) is elected and creates a block proposal which is then passed to validator nodes (members of the block [committee](../../terminology.md)) for approval.

## Participants in this flow

* Committee nodes
  * `ConsensusAlgo`
  * `ConsensusContext`
  * `TransactionPool`
  * `VirtualMachine`
  * `Processor`
  * `StateStorage`
  * `CrosschainConnector`
  * `BlockStorage`
  * `Gossip`

* All nodes
  * `ConsensusAlgo`
  * `ConsensusContext`
  * `BlockStorage`
  * `StateStorage`
  * `TransactionPool`
  * `Gossip`
  * `PublicApi`

## Assumptions for successful flow

* `BlockStorage` is synchronized to latest committed block.
  * A node cannot participate in a consensus committee unless this is the case.

## Flow

* [Before commit](block-creation-before-commit.md).

  > First half: steps 1-3 out of 5

* [After commit](block-creation-after-commit.md).

  > Second half: steps 4-5 out of 5
