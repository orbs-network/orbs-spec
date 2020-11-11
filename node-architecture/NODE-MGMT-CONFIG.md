# Node Management Config

> V2 release

## Scope

This configuration provides Boyar (the node orchestrator daemon) with all management information reflecting which services should be executed in the node cluster (ONG instances for all virtual chains and all global node services) and how they should be configured. This information changes over time and therefore must be checked repeatedly for changes.

This JSON config format is used both for the public Orbs network (where Ethereum is the raw source of truth for this data) and for private networks (where static JSON files managed by admin are the raw source of truth for this data).

### Bootstrap Config vs. Dynamic Config

When the node boots, Boyar uses a lean subset of this configuration to set up the minimal services required for bootstrap like the node management service. The bootstrap configuration is given to Boyar via command-line (normally by Nebula).

After the node management service starts, it accesses the various external sources of truth (eg. Ethereum, docker registry) and generates the fuller dynamic version of this configuration for Boyar to poll.

&nbsp;

## Configuration Format

* [Detailed example](../config-examples/node-management.json)

* [Public bootstrap example](../config-examples/node-management-bootstrap-public.json)

* [Private bootstrap example](../config-examples/node-management-bootstrap-private.json)

&nbsp;

Global fields:

| Field Name | Description |
| ---------- | ----------- |
| `orchestrator` | General information regarding the orchestration of the node. | 
| `services` | Dictionary of all node services (see fields per node service below).<br>The well-known dictionary keys include:<br>&bull;&nbsp;`management-service` - Management service<br>&bull;&nbsp;`signer` - Signer |
| `chains` | Array of all virtual chains (see fields per virtual chain below). |
| `network` | Array of network nodes. |

Fields under the orchestration section:

| Field Name | Description |
| ---------- | ----------- |
| `DynamicManagementConfig` | Settings related to the dynamic node management config to poll. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`Url` | URL of where the next dynamic configuration JSON should be polled from. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`ReadInterval` | How often in seconds the config should be polled. As duration string (s/m/h). |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`ResetTimeout` | How long in seconds should we allow the endpoint to fail providing valid config before resetting into bootstrap mode. As duration string (s/m/h). |
| `StorageDriver` | The storage driver to use for all persistent data.<br>Possible values:<br>&bull;`nfs` - Network File System. |
| `StorageOptions` | Storage driver-specific configurations. |

Fields per node service:

| Field Name | Description |
| ---------- | ----------- |
| `ExternalPort` | External port (outside docker) this instance should be listening on for incoming connections. This port is also open to the Internet. This field is optional. |
| `InternalPort` | Internal port (inside docker) this instance should be listening on for incoming connections. This field is optional, but required if ExternalPort is given. |
| `DockerConfig` | Settings specifying which docker image to run. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`Image` | Image name in the docker registry. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`Tag` | Version of the image in the docker registry. |
| `Config` | Contents of a JSON config file that is created and given to the instance on launch. |
| `Disabled` | Optional boolean field describing whether this node service is disabled. When a service is no longer active and should be removed (eg. protocol update), it should remain in the configuration with this value as true to make sure all nodes remove it explicitly. |

Fields per virtual chain:

| Field Name | Description |
| ---------- | ----------- |
| `Id` | The numerical ID of the virtual chain. |
| `ExternalPort` | External port (outside docker) this instance should be listening on for incoming gossip connections (for other nodes / peers) of this virtual chain. Different for different virtual chains. |
| `InternalPort` | Internal port (inside docker) this instance should be listening on for incoming gossip connections (for other nodes / peers) of this virtual chain. Identical for different virtual chains. |
| `InternalHttpPort` | Internal port (inside docker) this instance should be listening on for incoming HTTP connections (Public API for end-users). These connections are served externally by the Nginx node service which acts as API gateway for all virtual chains in the node. |
| `DockerConfig` | Settings specifying which docker image to run. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`Image` | Image name in the docker registry. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`Tag` | Version of the image in the docker registry. |
| `Config` | Contents of a JSON config file that is created and given to the instance on launch. |
| `Disabled` | Optional boolean field describing whether this virtual chain is disabled. When a virtual chain is no longer active and should be removed (eg. subscription expired), it should remain in the configuration with this value as true to make sure all nodes remove it explicitly. |
