# Virtual Chain (ONG) Architecture

> V2 release


> virtual chain contracts are not deployed or managed on Polygon network.

&nbsp;
## Terminology

Virtual chains are often referred to as "ONG" which is the codename for the [Orbs-network-go](https://github.com/orbs-network/orbs-network-go) reference implementation for a single virtual chain instance.

Definitions of terms like `node`, `service`, `committee`. Available [here](TERMINOLOGY.md).

&nbsp;
## ONG System Overview

![alt text][system_overview]

[system_overview]: _img/system-overview.png "system_overview"

Overview of a single virtual chain instance with brief introduction to all services is available [here](OVERVIEW.md).

&nbsp;
## High level flows

* [System Init](flows/system-init.md)

  Node and all services init (fresh start or after a restart).

* [Run Query](flows/run-query.md)

  Client calls a read-only method of a service (not under consensus).

* [Send Transaction](flows/send-transaction.md)

  Client sends a transaction that may write to state of a service (under consensus).

* [Get Transaction Status](flows/transaction-status.md)

  Client queries regarding the status and receipt of an old transaction.

* [Block Creation](flows/block-creation.md)

  The main continuous flow of consensus where blocks are created from pooled transactions.

* [Inter Node Sync](flows/inter-node-sync.md)

  Block synchronization between nodes.

* [Intra Node Sync](flows/intra-node-sync.md)

  Block synchronization between services inside a node.

&nbsp;
## Services

* [Public API](services/public-api.md)

  Provides external public gateway interface (like JSON over HTTP) for clients.

* [Transaction Pool](services/transaction-pool.md)

  Holds pending and committed transactions that are propagated between nodes. Helps avoid transaction duplication.

* [Gossip](services/gossip.md)

  Connects different nodes over the network with efficient message broadcast and unicast.

* [Consensus Algo](services/consensus-algo.md)

  Pluggable consensus algorithm (multiple algorithms supported side by side).

* [Consensus Context](services/consensus-context.md)

  Provides the system context for the consensus algorithm and deals with the actual content of blocks.

* [Virtual Machine](services/virtual-machine.md)

  Executes service methods (smart contracts) using various processors and produces state difference as a result.

* [Processor](services/processor.md)

  Stateless execution engine for smart contract methods in an isolated environment (multiple languages supported).

* [Block Storage](services/block-storage.md)

  Holds the long term journal of all confirmed blocks (the actual chain of blocks), performs block sync between nodes.

* [State Storage](services/state-storage.md)

  Holds the latest state under consensus meaning all of the state variables for all deployed services in a virtual chain.

* [Cross-chain Connector](services/crosschain-connector.md)

  Runs nodes for other blockchains like Ethereum and provides read access to them.
