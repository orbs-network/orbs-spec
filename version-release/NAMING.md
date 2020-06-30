# Image Naming and Tagging

> V2 release

## Docker Registry

Binary images for the various node services are published as Docker containers to a docker registry. These images are upgraded automatically by Boyar based on management configuration given by the management reader service.

### Image Naming Conventions

| Service | Image Name | Source Code |
| ------- | ---------- | ----------- |
| Virtual chain core<br>ONG (orbs-network-go) | `node` | https://github.com/orbs-network/orbs-network-go |
| Signer | `signer` | https://github.com/orbs-network/signer-service |
| Management reader | `management-service` | https://github.com/orbs-network/network-state-reader |
| Rewards service | `rewards-service` | https://github.com/orbs-network/rewards-distributor-v2 |
| Ethereum client | `ethereum-client` | *todo* |

Default registry: `registry.hub.docker.com`

Default namespace: `orbsnetwork`

Combined example - ONG image name: `orbsnetwork/node`

### Version Tagging Conventions

The node’s auto-deploy mechanism requires docker repositories to mark images with tag names that **match** a version according to the following versioning scheme:

```
v{PROTOCOL}.{MINOR}.{PATCH}[-canary][+hotfix]
```

* `{PROTOCOL}` indicates the latest supported protocol version. Can be any non-negative integer (0 and above). Note that when a new protocol version is released, not all services are necessarily released so some services might remain with latest versions tagged with a previous protocol version.
 
* `{MINOR}` indicates the changes in functionality according to semver semantics. It must increase monotonically within the same protocol. It can be any non-negative integer (0 and above).
 
* `{PATCH}` indicates changes in implementation according to semver semantics. It must increase monotonically within the same protocol. It can be any non-negative integer (0 and above).
 
* `-canary` is an optional segment that indicates the canary rollout group. If given, this is not a **main** version but rather a **canary** version that should only be rolled out to canary virtual chains.
 
* `+hotfix` is an optional segment that indicates that this version should be applied faster than normal. Normal gradual rollout takes place over **24h**, versions marked as hotflix roll out over **1h**.

The latest available version according to semver semantics will be deployed.
 
Examples of valid versions:
* `v1.2.3`
* `v1.2.3+hotfix`
* `v1.2.3-canary`
* `v1.2.3-canary+hotfix`
 
Notes:

* The `v` prefix is mandatory and has to be lower case.
* An ecmascript regex definition is below, it can also be found [here](https://regex101.com/r/Ly7O1x/310):
    ```
    /^v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(-canary)?(\+hotfix)?$ /gm
    ```

### Version Tagging Workflow

This is the recommended CI flow for a core node service.

* On push to branch/master
    * Publish as `${last git tag}-${commit hash}` to [`orbsnetworkstaging`](https://hub.docker.com/orgs/orbsnetworkstaging)

* On every merge to master 
    * Publish as `experimental` to [`orbsnetwork`](https://hub.docker.com/orgs/orbsnetwork)

* On every manual semver tag in git
    * If given manual approval in CI to publish to production:
        * Publish `${git tag}` in [`orbsnetwork`](https://hub.docker.com/orgs/orbsnetwork)
