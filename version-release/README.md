# Version Release Flow

> V2 release

## Concepts

* **Protocol version** - A new version may incur changes in the network protocol (how the various entities like validator nodes or clients communicate). The protocol version is a network-wide integer that increments with every major protocol change. This integer is marked clearly on every Orbs block header. Protocol changes often have to take place under consensus where all validator nodes switch together to the new protocol in order to avoid forks. Changes to the protocol version require a network governance decision and are indicated on a dedicated Ethereum contract.

* **Stable vs. canary** - The network supports an early release channel called "canary" for versions that are not yet considered stable. Alternative names for this concept in other places include "alpha" / "release candidate" / "test version". The subscription for a virtual chain indicates whether this virtual chain is stable or canary. When a canary version is released, it is deployed only to canary virtual chains. This allows the network to gain confidence for a version in a production setting using real validator nodes in a low risk environment.

* **Binary image upgrade** - A version release normally includes changes to the codebase of one or more of the services comprising an Orbs node. The new codebase is built and published as a ready-to-run [Docker](https://www.docker.com/) container. Validator nodes see the new image in a [docker registry](https://docs.docker.com/registry/) and may choose to upgrade the services to this new version. Note that a binary image may be published regardless of a protocol version change. For example, a fix for a memory leak does not change the protocol but still requires a new binary image to be released. Node operators may override the docker registry endpoint to opt out from auto upgrading from the default channels.

## Release Workflows

1. **Discussion** - Prior to development of a new feature, open community discussion and request for feedback is encouraged. This is particularly important for proposals of protocol changes that require a community governance decision to launch.

2. **Development** - Any feature developed should eventually become a PR (or a set of PRs) to the relevant repos. New code should be tested and existing tests should pass. Code modifications that require a protocol change must be conditional under a check whether the protocol version is active, while old behavior is preserved if not.

    See [DEVELOPMENT.md](DEVELOPMENT.md)

3. **Rollout to staging** - Staging environment is a personal test network that is owned and operated in full by the developer. It is deployed using standard tooling like Nebula, normally as a private network (since all nodes are named explicitly). New versions must be manually tested on staging using traffic simulations.

    See [STAGING.md](STAGING.md)

4. **Gradual rollout to canary** - Canary environment is the production network running on the actual validator nodes. If possible, new versions should be deployed first to canary virtual chains by publishing them to the canary distribution channel. If the version relies on a protocol change, community governance decision must be made on Ethereum to transition to the new protocol version on all canaries. 

    See [CANARY-RELEASE.md](CANARY-RELEASE.md)

5. **Gradual rollout to stable** - Final rollout to production takes place by publishing the version to the stable distribution channel. If the version relies on a protocol change, community governance decision must be made on Ethereum to transition to the new protocol version on all non-canaries. Rollout to the production network is always gradual to prevent network downtime from all validator nodes going down at once.

    See [STABLE-RELEASE.md](STABLE-RELEASE.md)

## Conventions

* **Image naming and tagging** - How new versions of binary images for the various services of the node should be labeled when published to a docker registry.

    See [NAMING.md](NAMING.md)