# Development Workflow

> V2 release

## Protocol Changes

A new feature or fix may incur changes in the network protocol (how the various entities like validator nodes or clients communicate). The protocol version is a network-wide integer that increments with every major protocol change. This integer is marked clearly on every Orbs block header.

Protocol versions increment explicitly as an external management action (eg. governance decision on Ethereum). They propagate to every virtual chain or service through [VC-MGMT-CONFIG.md](../node-architecture/VC-MGMT-CONFIG.md).

The protocol backwards compatibility strategy for different areas is:

| Service | Protocol Backwards Compatibility |
| ------- | -------------------------------- |
| Virtual chain core - ONG (orbs-network-go) | |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Block format | Full (all previous versions) |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Local persistent data | Full (all previous versions) |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Public API (for clients) | Full (all previous versions) |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Everywhere else | Short (~1 previous version) |
| Management reader | Full (all previous versions) |
| Ethereum writer | Short (~1 previous version) |
| Boyar | N/A (protocol version agnostic) |

Code that behaves differently under a new protocol version should be wrapped with a condition checking which protocol version is currently active while old behavior is maintained according to the strategy above.

## Release Notes
