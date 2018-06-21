# System Overview

![alt text][system_overview] <br/><br/>

[system_overview]: behaviors/_img/system-overview.png "system_overview"

This diagram shows a network node. The node connects to the outside world through interfaces seen on its four sides:
* End users - Users running a client which queries the network or performs transactions on it.
* App developers - Developers which deploy decentralized apps made of smart contracts to the network.
* Other Orbs nodes - Other nodes similar to this one which together create the decentralized Orbs network.
* Persistency - Persistence layer which stores data for long term and accesses similar storage on other blockchains.

The interface to *end users* is done through `PublicApi`. This service processes their queries and transactions.

The interface to *other Orbs nodes* is always done through `Gossip`. A dedicated service is required because in a network of 1000 nodes, we can't connect every node directly to every other node, therefore intelligent routing between nodes must be used. Through `Gossip`, the `ConsensusAlgo` service communicates with other nodes to perform the consensus process and close new blocks in the chain. The `TransactionPool` service communicates with other nodes to propagate pending transactions so they can eventually be added to a block during the ongoing consensus process. The `BlockStorage` service communicates with other nodes to perform block sync whenever the node is out-of-sync and unable to participate in the consensus process for new blocks.

The interface to *app developers* is done through multiple types of `Processor` services which support various smart contract runtimes and programming languages that execute the contracts written by these developers.

The interface to other blockchains like Ethereum is done through the `CrosschainConnector` service.

Persistency in the node and storage of the actual blockchain state is done by `BlockStorage` and `StateStorage`. They are separate because they have slightly different requirements. `BlockStorage` stores all the blocks for a long term and is the source of truth for the historic log of transactions. `StateStorage` only stores the accumulated recent state snapshot (after processing all past blocks), it is used for fast smart contract execution.

The `VirtualMachine` is charged with executing transactions. It wraps the various `Processors`, uses them for the actual code execution and interfaces between the code running in them and the outside world (for example providing state access).

While `ConsensusAlgo` implements the actual consensus algorithm (like Helix) that relies on multiple different nodes to reach agreement over the next approved block, it is oblivious to the block content (it can do consensus over any content). The `ConsensusContext` service is its counterpart and used to give the actual content of blocks the relevant system context (like holding transactions inside blocks).
