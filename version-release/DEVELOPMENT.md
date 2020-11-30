# Development Workflow

> V2 release

## Automated Tests in CI

It is expected that most features added to the system will include some sort of automated testing suite (normally a mix of unit tests, integration tests and end-to-end tests). The entire test suite for a subsystem runs on CI on every push to source control (git). This practice protects against regression bugs and makes sure existing behavior has not been compromised by new code.

## Tagging in Git and Release Notes

Once the code in `master` branch is considered ready for release, the commit marking the release should be tagged in git. The convention is to tag as [semver](https://semver.org), ie. `v1.2.3`.

Some people are confused with deployment scheme, ie. `-hotfix` or `-immediate`. Note that this information is not related to git and should **not** appear in the git tag. The deployment scheme is only part of how images are tagged in the docker repositories.

When creating a new release, please include a short MD file detailing the release notes to aid the community in reviewing the changes.

## Protocol Changes

A new feature or fix may incur changes in the network protocol (how the various entities like validator nodes or clients communicate). The protocol version is a network-wide integer that increments with every major protocol change. This integer is marked clearly on every Orbs block header.

Protocol versions increment explicitly as an external management action (eg. governance decision on Ethereum). They propagate to every virtual chain or service through [VC-MGMT-CONFIG.md](../node-architecture/VC-MGMT-CONFIG.md).

#### Does a change require a protocol change?

To answer whether a code change requires a protocol upgrade ask yourself the following question: if a single node makes the upgrade, could they still participate normally in consensus with all other nodes that have not made the upgrade?

If the answer is positive, a protocol change is probably not required. If the answer is negative, protocol change probably is required.

#### Backwards compatibility

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
| Logs service | N/A (protocol version agnostic) |
| Boyar | N/A (protocol version agnostic) |

Code that behaves differently under a new protocol version should be wrapped with a condition checking which protocol version is currently active while old behavior is maintained according to the strategy above.
