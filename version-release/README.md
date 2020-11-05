# Version Release Flow

This document describes the release workflow for docker images participating in the Orbs node deployment.     

Orbs software modules are delivered in several modes:

- npm packages (e.g. [Polygon](https://www.npmjs.com/package/@orbs-network/polygon))
- Compiled Binaries published on S3 (e.g. [Boyar](http://github.com/orbs-network/boyarin))
- Docker images on a public docker registry (e.g orbs-network core node, management service, rewards service, etc)

While Polygon and Boyar are released rarely and require manual intervention to release and to deploy, docker images are automatically deployed by the Management Service running on each Orbs Node. 

The process involves 3 steps:
1. CI automatically builds docker images for each code commit or tag in Github and publishes them to a staging repository in a public docker registry.
1. An existing image is marked for deployment to the production network by tagging adding it the the production deployment repository  
1. Management service notices the requested deployment and schedules an upgrade to be carried out by the node manager agent (Boyar)

Marking an image for deployment by the nodes management services can be done manually by re-tagging an image under the appropriate repository.
For details on how to manually deploy an image see [naming conventions](NAMING.md)

However using `deploy.sh` script simplifies the process of tagging an image for deployment. 

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

4. **Gradual rollout to canary virtual chains** - Canary environment is the production network running on the actual validator nodes. If possible, new versions should be deployed first to canary virtual chains by publishing them to the canary distribution channel. If the version relies on a protocol change, community governance decision must be made on Ethereum to transition to the new protocol version on all canaries. 
    
   Example for marking an `orbs-network-go` image tagged as `v2.0.3` in github for deployment on canary Virtual Chains
   ```shell script
   # in the standard rollout timeframe:
   ./deploy.sh --tag v2.0.3 --canary --standard

   # in the "hotfix"* rollout timeframe:
   ./deploy.sh --tag v2.0.3 --canary --hotfix
   ```
    `*` `--canary` deployments are available only for virtual chain modules, not for node services as there is only one instance of each node.

5. **Gradual rollout to stable virtual chains** - Final rollout to production takes place by publishing the version to the stable distribution channel. If the version relies on a protocol change, community governance decision must be made on Ethereum to transition to the new protocol version on all non-canaries. Rollout to the production network is always gradual to prevent network downtime from all validator nodes going down at once.

   Example for marking an `orbs-network-go` image tagged as `v2.0.3` in github for deployment on stable Virtual Chains
   ```shell script
   # in the standard rollout timeframe:
   ./deploy.sh --tag v2.0.3 --standard

   # in the "hotfix"* rollout timeframe:
   ./deploy.sh --tag v2.0.3 --hotfix
   ```

6. **Immediate rollout to a node service** - Node services whith no requirement for gradual rollout can be deployed immediately. Deploying immediately can be dangerous, create network wide downtime, or compromise consensus. Use this deployment mode with caution

   Example for marking an `ethereum-writer` image tagged as `v2.0.3` in github for immediate deployment on all nodes
   ```shell script
   ./deploy.sh --tag v2.0.3 --repo "ethereum-writer" --immediate
    
## Conventions

* **Image naming and tagging** - How new versions of binary images for the various services of the node should be labeled when published to a docker registry.

    See [NAMING.md](NAMING.md)

## Deploy tool

Deployment tool ([deploy.sh](deploy.sh)) for Orbs node modules. Supported module types are node services, such as signer, ethereum writer, rewards, and management node services, as well as to VChain modules.

Prerequisites: 

- docker CLI installed and logged in with read access to the source image, and write access to the deployment target organization and repository on the default registry
- A reference to a pre-built source docker image to deploy

The reference to the source image must be of this form: 
   `organization/repository:tag`

If the source docker image is not present on the locally it must be available for download at the default registry

Deployment modifiers: --standard, --immediate, --hotfix and --canary.

`-s, --standard` Deploy randomly within a 24 hours window. This is the safest deployment mode with the highest guarantee against downtime or outages. Using --standard indicates to the Orbs node Management Service that this upgrade should be deployed in the Normal rollout window.

`-i, --immediate` Immediate deployment across the network. This deployment mode is unsafe for "management-service" and "node" repos. Use with caution. For VChain modules ("node" repo) or management-service this option is especially risky. Using --immediate indicates to the Orbs node Management Service that this upgrade should be deployed without any rollout window. Use with caution!

`-h, --hotfix` To ensure constant availability of a quorum of nodes, the Orbs Management Service deploys updates to VChain core modules ("node" repo) with a gradual rollout window. Orbs supports two mode of rollout: Standard for a safer and longer rollout window of 24 hours (by default), and Hotfix rollout window for urgent, expedited, rollouts within 1 hour (by default). for more information see https://github.com/orbs-network/orbs-spec
Using --hotfix indicates to the Orbs node Management Service that this upgrade should be deployed in the Hotfix rollout window.

`-c, --canary` This option indicates deployment only to "Canary" VChains. For more information on Canary VChains see https://github.com/orbs-network/orbs-spec. This option is applicable only to Virtual Chain modules.

          Usage: ./$me [OPTIONS]

          -s, --standard   rolling deployment (24 hour deployment window)
          -h, --hotfix     deploy as hotfix (1 hour rolling deployment window)
          -i, --immediate  deploy immediately (no rolling deployment)
          -c, --canary     deploy only to canary vchains, relevant only for "node" repository images
          -t, --tag        the source tag to deploy from (default: "experimental")
          --target-tag     the target tag to deploy to (default: [source tag])
          -r, --repo       the source repository to deploy from (default: "node")
          --target-repo    the target repository to deploy to (default: [source repository])
          -o, --org        the source organization to deploy from (default: orbsnetworkstaging)
          --target-org     the target organization to deplot to (default: orbsnetwork)

          -y               suppress confirmations

### deploy.sh examples - Staging env

```shell script

# deploy current experimental versions to staging
./deploy.sh --target-tag v100.0.0 --target-org orbsnetworkstaging -r node -s
./deploy.sh --target-tag v100.0.0 --target-org orbsnetworkstaging -r signer -s
./deploy.sh --target-tag v100.0.0 --target-org orbsnetworkstaging -r management-service -s
./deploy.sh --target-tag v100.0.0 --target-org orbsnetworkstaging -r ethereum-writer -s
./deploy.sh --target-tag v100.0.0 --target-org orbsnetworkstaging -r rewards-service -s

# reset the bootstrap image to the current experimental version management-service
./deploy.sh --target-tag bootstrap --target-org orbsnetworkstaging --repo management-service -s

```
### deploy.sh examples - Production env

```shell script

# deploy node (orbs core) tagged version v2.0.4 as a canary hotfix with no confirmation prompt
./deploy.sh -t v2.0.4 --canary --hotfix -y

# deploy ethereum writer version v1.2.0
./deploy.sh -t v1.2.0 -r ethereum-writer --immediate

```

### deploy.sh examples - Advanced Production use cases

```shell script

# upgrade the bootstrap management-service image to v1.1.2 
./deploy.sh --tag v1.1.2 --target-tag bootstrap -r management-service

# override rewards-service version based on a later patch by commit hash v1.1.1-2cff1b0e
# Caution - such override may not apply on previously created nodes, and may have some unexpected results
 
./deploy.sh -t v1.1.1-2cff1b0e --target-tag v1.1.1 -r rewards-service
```
