# Orbs Network Protocol Specification

> V2 release (other [versions](VERSIONS.md))

## Table of Contents

### 1. [Node Architecture](node-architecture/)

Outlines all the various subsystems that make an Orbs network node.

### 2. [Proof-of-Stake Architecture](pos-architecture/)

Outlines PoS in the Orbs Universe with incentives, elections and rewards implemented over Ethereum.

### 3. [Virtual Chain Architecture](vchain-architecture/)

The heart of an Orbs node is a set of *virtual chains*, standalone blockchain instances executing smart contracts and implementing Helix consensus. The interfaces between the various services of a virtual chain are defined [here](interfaces/).

### 4. [Version Release](version-release/)

Explanation of the flow for a version relase where each of the node subsystems can be upgraded and deployed to validators.

### 5. [Encoding](encoding/)

Specification for the binary encoding of protocol network messages (both client protocol and gossip protocol).
