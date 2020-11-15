# Orbs Network Protocol Specification

> V2 release

## Table of Contents

### 1. [Node Architecture](node-architecture/README.md)

Outlines all the various subsystems that make an Orbs network node.

### 2. [Proof-of-Stake Architecture](pos-architecture/README.md)

Outlines PoS in the Orbs Universe with incentives, elections and rewards implemented over Ethereum.

### 3. [Virtual Chain Architecture](vchain-architecture/README.md)

The heart of an Orbs node is a set of *virtual chains*, standalone blockchain instances executing smart contracts and implementing Helix consensus. The interfaces between the various services of a virtual chain are defined [here](interfaces/README.md).

### 4. [Version Release](version-release/README.md)

Explanation of the flow for a version relase where each of the node subsystems can be upgraded and deployed to validators.

### 5. [Encoding](encoding/README.md)

Specification for the binary encoding of protocol network messages (both client protocol and gossip protocol).
