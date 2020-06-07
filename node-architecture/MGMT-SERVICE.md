# Management Service

> V2 release

## External Sources of Truth

### Ethereum

The management service relies on an external source of truth to provide it with raw information regarding the current state of the network. In a public Orbs node, this source of truth is the Ethereum blockchain. A collection of smart contracts on Ethereum reports events reflecting:

* Ref time - Ethereum block [time](https://en.wikipedia.org/wiki/Unix_time). This information does not vary per virtual chain. Ref time reflects the clock of the external source of truth (Ethereum).  

* Elected committee - Changes in which nodes are authorized to sign and validate blocks and their respective weights. This information does not vary per virtual chain and only depends on ref time (Ethereum block time).

* Subscriptions - Which virtual chain IDs are active and which have expired, and what resource limitations should be enforced for each (per subscription tier). This information naturally varies per virtual chain since each virtual chain has its own separate subscription.

* Topology - Network addresses of all nodes that are part of the network (committee and candidates). Node addresses do not vary per virtual chain and only depend on ref time. When gossip ports are added to the addresses, the combination does vary per virtual chain.

* Protocol version - Changes in the official protocol version that must be supported on all nodes participating in the network. This information does not vary per virtual chain and only depends on ref time. Canary virtual chains have their own protocol version that usually increments earlier.

* Contract registry - An Ethereum contract holding the addresses of all other Ethereum contracts (to support contract upgrades). This information does not vary per virtual chain and only depends on ref time.

The management service scans Ethereum in chronological block order from the beginning of history to look for relevant events. The service is unable to respond to queries fully until this synchronization is complete. In addition, since recent Ethereum blocks are susceptible to change (re-org), the service adds finality buffers to make sure Ethereum state is final. This is why at any point in time the management service reflects a state of Ethereum that is several minutes in the past.

### Derived Data According to Protocol

The way to derive some fields is defined in the protocol. For example, the external gossip port per virtual chain is a protocol-defined derivative of the virtual chain ID. Internal service endpoints are derived from the service key and its internal port. For all specifications, see “Definitions of Specific Fields” below.

### Docker Registry

The management service supports automatic upgrading for all binary images of dockerized instances given to Boyar. These include all node services and all instances running virtual chains. Note that Boyar itself is not a node service and cannot be upgraded automatically.

The service polls a provided docker registry and eagerly upgrades to new versions as they are published to the registry. The default docker registry is [Docker Hub](https://hub.docker.com/). Canary virtual chains have their own specially tagged images that are normally published earlier. Binary version upgrades are always moving forward, meaning if the service sees an older version as the most recent in Docker Hub, it will not rollback and ignore.

Since the management service configures all other services in the system and provides their launch arguments, it has intimate knowledge about each service and which launch configuration arguments it expects.

The tagging convention for the docker registry images is described in [NAMING.md](../version-release/NAMING.md).

## Private Networks

The Orbs architecture is primarily designed to accommodate the public Orbs network where Ethereum mainnet is the source of truth for network state. Nevertheless, the Orbs codebase supports running an isolated collection of nodes in a private network mode.

When running in a private network mode, Ethereum mainnet is no longer relevant as the source of truth. In this scenario, the management service is replaced by static JSON files managed directly by the private network administrator.

## Output Endpoints

The management service serves multiple clients through multiple different endpoints. These clients include:

* Boyar - Boyar polls this endpoint to provide it with its dynamic configuration. This configuration includes node-level definitions specifying which node services and which virtual chains should be executed in the node cluster.

    Format: [NODE-MGMT-CONFIG.md](NODE-MGMT-CONFIG.md)

    Endpoint: `node/management`

* ONG - Every virtual chain ONG instance polls this endpoint separately to provide it with its management config. This configuration includes the elected committee, state of subscription, topology, protocol version, etc.

    Format: [VC-MGMT-CONFIG.md](VC-MGMT-CONFIG.md)

    Endpoint: `vchains/{VCHAIN-ID}/management`

### Current Data vs. Historical Data

Every event reported on Ethereum is reported on a specific Ethereum block, that has a unique block time. The Ethereum block time is considered the ref time of the source of truth.

In the majority of cases, the output endpoints should reflect the current state of the network (after a small sync gap and finality buffers). For Boyar, current data is a snapshot of the currently needed services. For ONG, current data reflects aggregated Ethereum events from the last **24 hours** sliding time window. If there are no events for a certain topic in this timeframe, the aggregated data is not empty and reflects the latest event prior to the timeframe.

Current state endpoint: `/vchains/123/management`

In infrequent cases, ONG is required to audit historic data for security purposes. This happens when a virtual chain is first downloading the block history from peers (and verifying it) and when a virtual chain is running in audit mode. When this is the case, the ONG output endpoint must provide historic snapshots.

When an ONG instance accesses the historic data endpoint, it provides the historic ref time as an argument. Historic snapshots are paged by management service in daily intervals - meaning the snapshot for ref time X contains aggregated data from all Ethereum events with ref time between `X - [X mod (day)]` and `X - [X mod (day)] + day`, where day is `60 * 60 * 24`.  If there are no events for a certain topic in this timeframe, the aggregated data is not empty and reflects the latest event prior to the timeframe.

Historic state endpoint: `/vchains/123/management/1582619954`

The last parameter is the requested historic ref time (as a Unix timestamp, without any rounding or paging constraints, this ref time is rounded internally by the reader service to the relevant snapshot page which is then returned in full).

## Multiple Protocol Versions

The official version of the Orbs protocol for the public network at any given time appears on Ethereum in a dedicated contract. When the protocol version changes, an Ethereum event is reported. This means the network protocol version at any point in history is tied to the ref time (Ethereum block time) at that point.

The management service naturally behaves in accordance with the Orbs protocol specifications. For example, the address of the Ethereum contract to query is defined as part of the protocol. Protocol versions change over time. The management service is required to be fully backwards compatible and support all historic protocol versions to date.

When the management service handles requests for current ref time, it switches to a new protocol version as soon as it sees the relevant Ethereum block on the Ethereum blockchain (listens to the tip). When historic data is requested at a historic ref time, the response should satisfy the protocol version at that historic ref time. If the protocol version changed in the middle of a historic snapshot, the first half should conform to the first protocol version and the second half to the second protocol version.

## Uncoordinated Virtual Chain Restarts

One of the requirements of the consensus algorithm is that the network maintains a running quorum of nodes at any given time. This presents a challenge with software upgrades and node restarts since we don’t want all nodes to go offline at the same time.

To accommodate this, **any change** the reader service performs to any virtual chain instance or any node service that may cause it to restart and suffer downtime, should follow the gradual [ROLLOUT.md](../version-release/ROLLOUT.md) flow based on randomized timeslots.

When a new virtual chain is created, its subscription indicates the genesis ref time of when its consensus starts. The reader service launches the container in advance to make sure it is running by that time.

## Definitions of Specific Fields

* `CurrentRefTime` is the latest Ethereum block time under finality and sync state:

    * `CurrentRefTime = timeOfLatestBlockNumber - 100`

    * Note: Grace calculation was moved to ONG. Events are reported only if they were emitted with time (block) <=  `CurrentRefTime`

    * Note: ONG is not aware of finality at all.

    * Note: at any point in time the management service may reflect a state of ethereum that is in the past, if the scan process didn’t catch up with the current state yet.

* `CurrentTopology` is the superset of all the committees in the last 12 hours. Notice that this includes the committee that was updated in the last event prior to the 12 hour mark.

* `ExternalPort` per virtual chain is defined as such:

    `vc_index = VirtualChainId - 1000000`

    `ExternalGossipPort = vc_index + GOSSIP_PORT_BASE`

    `GOSSIP_PORT_BASE = 10000`

* `NodeServices` dictionary keys (names of the services) are defined under [Image naming and tagging](../version-release/NAMING.md).

    * These keys also define the docker image names in the docker registry:

        ```
        {DOCKER-NAMESPACE}/{SERVICE-NAME}
        ```

        For example: `orbsnetwork/signer`

    * These keys also define internal network endpoints:

        ```
        http://{SERVICE-NAME}:{INTERNAL-PORT}
        ```

        For example: `http://management-service:8080`
