# Orbs Core Protocol Specification

V1 release

&nbsp;
## Terminology

Definitions of terms like `node`, `service`, `committee`. Available [here](terminology.md).

&nbsp;
## System Overview

![alt text][system_overview]

[system_overview]: behaviors/_img/system-overview.png "system_overview"

Overview of the entire system with brief introduction to all services is available [here](overview.md).

&nbsp;
## High level flows

* [System Init](behaviors/flows/system-init.md)

  Node and all services init (fresh start or after a restart).

* [Call Method](behaviors/flows/call-method.md)

  Client calls a read-only method of a service (not under consensus).

* [Send Transaction](behaviors/flows/send-transaction.md)

  Client sends a transaction that may write to state of a service (under consensus).

* [Get Transaction Status](behaviors/flows/transaction-status.md)

  Client queries regarding the status and receipt of an old transaction.

* [Block Creation](behaviors/flows/block-creation.md)

  The main continuous flow of consensus where blocks are created from pooled transactions.

* [Inter Node Sync](behaviors/flows/inter-node-sync.md)

  Block synchronization between nodes.

* [Intra Node Sync](behaviors/flows/intra-node-sync.md)

  Block synchronization between services inside a node.

&nbsp;
## Services

* [Public Api](behaviors/services/public-api.md)

  Provides external public gateway interface (like JSON over HTTP) for clients.

* [Transaction Pool](behaviors/services/transaction-pool.md)

  Holds pending and committed transactions that are propagated between nodes. Helps avoid transaction duplication.

* [Gossip](behaviors/services/gossip.md)

  Connects different nodes over the network with efficient message broadcast and unicast.

* [Consensus Algo](behaviors/services/consensus-algo.md)

  Pluggable consensus algorithm (multiple algorithms supported side by side).

* [Consensus Context](behaviors/services/consensus-context.md)

  Provides the system context for the consensus algorithm and deals with the actual content of blocks.

* [Virtual Machine](behaviors/services/virtual-machine.md)

  Executes service methods (smart contracts) using various processors and produces state difference as a result.

* [Processor](behaviors/services/processor.md)

  Stateless execution engine for smart contract methods in an isolated environment (multiple languages supported).

* [Block Storage](behaviors/services/block-storage.md)

  Holds the long term journal of all confirmed blocks (the actual chain of blocks), performs block sync between nodes.

* [State Storage](behaviors/services/state-storage.md)

  Holds the latest state under consensus meaning all of the state variables for all deployed services in a virtual chain.

* [Cross-chain Connector](behaviors/services/crosschain-connector.md)

  Runs nodes for other blockchains like Ethereum and provides read access to them.
