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

The management service supports automatic upgrading for all binary images of dockerized instances given to Boyar. These include all node services and all instances running virtual chains (ONG). Note that Boyar itself is not a dockerized node service and has its own special and separate upgrade flow.

The service polls a provided docker registry and eagerly upgrades to new versions as they are published to the registry. The default docker registry is [Docker Hub](https://hub.docker.com/). Canary virtual chains have their own specially tagged images that are normally published earlier. Binary version upgrades are always moving forward, meaning if the service sees an older version as the most recent in Docker Hub, it will not rollback and ignore.

Since the management service configures all other services in the system and provides their launch arguments, it has intimate knowledge about each service and which launch configuration arguments it expects.

The tagging convention for the docker registry images is described in [NAMING.md](../version-release/NAMING.md).

&nbsp;

## Private Networks

The Orbs architecture is primarily designed to accommodate the public Orbs network where Ethereum mainnet is the source of truth for network state. Nevertheless, the Orbs codebase supports running an isolated collection of nodes in a private network mode.

When running in a private network mode, Ethereum mainnet is no longer relevant as the source of truth. In this scenario, the management service is replaced by static JSON files managed directly by the private network administrator.

&nbsp;

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

Current state endpoint example: `/vchains/123/management`

In infrequent cases, ONG is required to audit historic data for security purposes. This happens when a virtual chain is first downloading the block history from peers (and verifying it) and when a virtual chain is running in audit mode. When this is the case, the ONG output endpoint must provide historic snapshots.

When an ONG instance accesses the historic data endpoint, it provides the historic ref time as an argument. Historic snapshots are paged by management service in daily intervals - meaning the snapshot for ref time X contains aggregated data from all Ethereum events with ref time between `X - [X mod (day)]` and `X - [X mod (day)] + day`, where day is `60 * 60 * 24`.  If there are no events for a certain topic in this timeframe, the aggregated data is not empty and reflects the latest event prior to the timeframe.

Historic state endpoint example: `/vchains/123/management/1582619954`

The last parameter is the requested historic ref time (as a Unix timestamp, without any rounding or paging constraints, this ref time is rounded internally by the reader service to the relevant snapshot page which is then returned in full).

&nbsp;

## Multiple Protocol Versions

The official version of the Orbs protocol for the public network at any given time appears on Ethereum in a dedicated contract. When the protocol version changes, an Ethereum event is reported. This means the network protocol version at any point in history is tied to the ref time (Ethereum block time) at that point.

The management service naturally behaves in accordance with the Orbs protocol specifications. For example, the address of the Ethereum contract to query is defined as part of the protocol. Protocol versions change over time. The management service is required to be fully backwards compatible and support all historic protocol versions to date.

When the management service handles requests for current ref time, it switches to a new protocol version as soon as it sees the relevant Ethereum block on the Ethereum blockchain (listens to the tip). When historic data is requested at a historic ref time, the response should satisfy the protocol version at that historic ref time. If the protocol version changed in the middle of a historic snapshot, the first half should conform to the first protocol version and the second half to the second protocol version.

&nbsp;

## Uncoordinated Virtual Chain Restarts

One of the requirements of the consensus algorithm is that the network maintains a running quorum of nodes at any given time. This presents a challenge with software upgrades and node restarts since we don’t want all nodes to go offline at the same time.

To accommodate this, **any change** the reader service performs to any virtual chain instance or any node service that may cause it to restart and suffer downtime, should follow the gradual [ROLLOUT.md](../version-release/ROLLOUT.md) flow based on randomized timeslots.

When a new virtual chain is created, its subscription indicates the genesis ref time of when its consensus starts. The reader service launches the container in advance to make sure it is running by that time.

&nbsp;

## Detailed Specification

### Concepts

* Current reference time represents the latest time of management data read (Etehreum finality time).
  * All newer entries are ignored and must not be cached (by default should not be logged).
  * All entries earlier than the current reference time are final.
* Two type of entries:
  * State - eg. Topology, a state valid for the current reference time.
  * Events entries - a state valid from timestamp until the next entry timestamp.
* Entries are not immediately derived from Ethereum events and may be based on multiple events, accumulated state.
* An entry reference time indicates the time the entry takes effect. Note: necessarily derived from the Ethereum event timestamp.
* Each section in each page must have valid entries for the entire period, this may require adding an entry prior to the time window to be valid in its beginning.

### Global fields

* `CurrentRefTime` is the latest Ethereum block time under finality and sync state:
    * `CurrentRefTime = timeOfLatestBlockNumber - 40`
    * Note: Grace calculation was moved to ONG. Events are reported only if they were emitted with time (block) <=  `CurrentRefTime`
    * Note: ONG is not aware of finality at all.
    * Note: at any point in time the management service may reflect a state of ethereum that is in the past, if the scan process didn’t catch up with the current state yet.
* `PageStartRefTime` is the earliest ref time guaranteed to be covered by the response. Note that it’s ok if events outside this range appear in the response.
* `PageEndRefTime` is the latest ref time covered by the response (always `PageEndRefTime <= CurrentRefTime`).

### All entries

* All explicit and implicit entries within `[PageStartRefTime .. PageEndRefTime]` must be included in the page.
* First entry - All pages must include a first valid entry that represents the status before the first entry is generated in the range. The first entry is the latest explicit / implicit entry generated for each category with `referenceTime <= PageStartRefTime`.
* Additional entries with a timestamp before or after the `PageStartRefTime` / `PageEndRefTime` may be included. All entries with `first entry <= timestamp <= PageEndRefTime` must be included in the page (no holes).
* Entries with `timestamp > CurrentRefTime` (newer blocks) are not final and may be changed. If included they are ignored by the consumer, and must not be cached.
* If multiple Ethereum events of the same type are emitted on the same block (same `referenceTime`), a single entry is built from the accumulated data based on the events order. (logIndex). If multiple entries are entered with the same timestamp, only the first one is considered by the consumer.
* Implicit entries -  Implicit entries referenceTime is derived from the data of the latest entry emitted prior to that reference time and not from the block timestamp. There are two types of implicit entries that are generated based on another entry data:
  * `ProtocolVersion` entries - generated based on ProtocolVersionChanged entries.
  * `Subscription.expired` entries - generated based on SubscriptionChanged entries.

### Virtual chain data

#### `GenesisRefTime`

The Genesis reference time derived from: 

`Subscriptions.VcCreated(uint256 vcid, address owner, uint32 genesisReferenceTime)`

Note: currently genesis reference is reported in SubscriptionChange event in block, moving to VcCreated, timestamp.

#### `Committee`

Committee entries are generated upon any change in the Committee members or their weight.

There are 3 type of Ethereum events that trigger committee entries:

* `Commmittee.CommitteeChanged(address[] addrs, address[] orbsAddrs, Uint256[] weights, bool[] certified)`<br>
    Indicating a new committee along with its members’ addresses and weights.
* `Commmittee.CommitteeMemberChanged(address addr, address orbsAddr, uint256 weight, bool certified, string currentStatus, string previousStatus)`<br>
    CurrentStatus / previousStatus - Committee, Standby, Candidate. Indicating a change in a member’s membership.
* `ValidaorsRegistery.ValidatorDataUpdated(address addr, bytes4 ip, address orbsAddr, string name, string website, string contact)`<br>
    Orbs address update. (TBD - gas saving). The event generates a new committee entry with the same member but an updated Orbs address. The events should be ordered together with the Committee and Standby events.

In order to build the committee entries, one has to find the last CommitteeChanged event and apply the delta on each event. 

Note: A block or even a single transaction may include multiple committee changed / CommitteeMemberChanged events. In such case, the entry should hold the accumulated state.

Note: the members in the committee entry are ordered by the {weight, Eth address}

Trigger events:
* `Commmittee.CommitteeChanged`
* `Commmittee.CommitteeMemberChanged`
* `ValidaorsRegistery.ValidatorDataUpdated` - Generating an entry on validator’s Orbs address update (pending final decision on removing Orbs address from the events).

Fields:
* `Committee.RefTime`: the block time of the event.
* `Committee` members: all members that are in Committee status at the time of the event, based on the accumulated state.
* `Committee.EthAddress` - from events.
* `Committee.OrbsAddress`  - from events. See removal of Orbs address from Ethereum events
*  `Committee.Weight` - `AdjustedConsensusWeight(event.weight)`<br>
    `AdjustedConsensusWeight = Max(event.weight, average committee weight) = Max(event.weight, total committee weight / number of committee members)`
    <br>
    Note: currently the event has an Effective Stake field, modifying to Weight.
    <br>
    In the Ethereum events, combined stake of the validator operating this node and the delegators who have chosen them as guardian (weight field in the event) are reported in Orbitons (10^-18 ORBS). The Committee.Weight units should be changed to 1 ORBS, to allow the field to fit inside a JSON number (float64).
* `Committee.IdentityType` - event.certified
    Ignored by the consumer.
    A virtual chain database should only include its relevant members.

#### `CurrentTopology` 

`CurrentTopology` is the superset of all the committees in the last 12 hours. Notice that this includes the committee that was updated in the last event prior to the 12 hour mark.

* Members: union of the committee members in the last 12 hours + latest standby members.
* `member.EthAddress` - member’s EthAddress (note: currently not in the JSON, added for completeness).
* `member.OrbsAddress` - member’s OrbsAddress.
* `member.Ip` - taken from the up to date topology map (see Topology Map).
* `member.Port` - based on the VCID.

Note: the members in the committee entry are ordered by the {Orbs address}.

Topology map:

* Represents the up to date IP address table for each member (Ethereum address).
* Updated on `ValidatorsRegistration.ValidatorDataUpdated(address addr, bytes4 ip, address orbsAddr, string name, string website, string contact, bool newValidator)`.

#### `Subscription`

Subscription entries hold the relevant subscription change entries for the specific virtual chain.

There are two types of events:

* active - represents a subscription start time. An active entry is set for a new VC or an expired one.
* expired - represents a subscription expiry time. An expired entry is set following an active entry. If a subscription event is emitted while a virtual chain is active, it’s expired entry is updated.

Subscription entries are based on `Subscription.SubscriptionChanged(uint256 vcid, uint256 genRef, uint256 expiresAt, string tier, string deploymentSubset)` events.

Trigger events:
* `Subscription.SubscriptionChanged`
* `Subscription.VcConfigRecordChanged`

Fields:
* Active entries:
  * `ActiveSubscription.RefTime = event.timestamp`
  * `ActiveSubscription.Status = active`
  * `ActiveSubscription.Tier = event.tier`
  * `ActiveSubscription.RolloutGroup = event.deploymentSubset`
  * `ActiveSubscription.Params` - the state of the Virtual chains’s parameter records. Presented as a set of {key,value}. Based on `VcConfigRecordChanged(uint256 vcid, string key, string value)` events.
* Expired entries:
  * `ActiveSubscription.RefTime = expiresAt`
  * `ActiveSubscription.Status = expired`
  * `ActiveSubscription.Tier = event.tier`
  * `ActiveSubscription.RolloutGroup = event.deploymentSubset`
  * `ActiveSubscription.Params` - the state of the Virtual chains’s parameter records. Presented as a set of {key,value}. Based on `VcConfigRecordChanged(uint256 vcid, string key, string value)` events.

#### `ProtocolVersion`

Protocol entries: Implicit entries, emitted for each `ProtocolVersionChanged(string deploymentSubset, uint256 currentVersion, uint256 nextVersion, uint256 fromTimestamp)` event.

At any point in time there can only be a single pending protocol change entry with `timestamp > Current referenceTime`. Once a new event is emitted, generating a new entry (may be with timestamp before or after the first one), the new entry overrides the current one.

Notes on the Ethereum protocol contract: 
* The protocol contract holds a single pending protocol change entry.
* In order to prevent scenarios where an event is emitted too recently to its timestamp, a protocol event can’t be triggered an hour prior to the pending one timestamp.
 
Trigger events:
* `Protocol.ProtocolVersionChanged`

Fields:
* `ProtocolVersion.referenceTime = event.fromTimestamp`
* `ProtocolVersion.RolloutGroup = event.deploymentSubset`. Ignored by the consumer.
* A virtual chain database should only include its relevant `RolloutGroup` events.
* `ProtocolVersion.Version = event.nextVersion`

#### `Contract Registry`

* The contract addresses are provided by the registry contract.
* Addresses update are indicated by the `ContractAddressUpdated(string contractName, address addr)` event.
* The relevant contract names are: protocol, committee (currently 2 separate committees), validatorsRegistration, subscriptions.

The management service is expected to maintain a range table of block range to address for each contract. Events for each range are emitted by the contract address.

### Standbys & candidate Guardians

Standby Guardians represent the list of `N=22+5` validators added to the topology. Both the standby Validator and his committee and standby peers need to be aware of the Validator being a standby member.

Ordered Candidate list:

* Guardian RFS = TRUE, as obtained from valid (based on referenceTime) `Election.ValidatorStatusUpdated(address, RFS = TRUE, RFC)` event.
* Not in committee
* Ordered by: {NonStale, EffectiveStake, EthereumAddress}
  * `RFS_timestamp`:
    * Last `ValidatorStatusUpdated(addr, RFS = TRUE).ReferenceTime`
    * A validator in the committee is considered RFS for every block he’s in the committee. We can optimize and update only when leaves the committee but then we need to make sure it’s reflected immediately in topology calculation. Note: a validator that is pushed out of the committee is treated as if he just sent RFS.
  * NonStale:
    * `RFS_timestamp + STALE_TIMEOUT >= Ethereum referenceTime`
    * `STALE_TIMEOUT = 7 days`
  * EffectiveStake:
    * Obtained from `Election.StakeChanged(addr, selfStake, delegated_stake, effective_stake)`

Current Standy:

The top 5 Guardians in the Ordered Candidate list.

### Registered guardians

All registered guardians ordered by {EffectiveStake, EthereumAddress}

Data: all registration fields.
Metadata{`REWARDS_FREQUENCY_SEC`}

Effective stake is obtained from `Election.StakeChanged(addr, selfStake, delegated_stake, effective_stake)`. Note: when storing the `EffectiveStake` map, to reduce memory footprint, the service can clear entries with `EffectiveStake = 0`.

### Coordinated upgrade rules

* Continuously update the latest relevant image and VC parameters.
  * When a new image is identified, set a new `rollout_time`.
* `vc_rollout_window`
  * Determined by the image type: 
    * Regular - 24h 
    * Urgent - 1h
* Determine upgrade time:<br>
  `rollout_time = local_time + [H(address, current_time) mod upgrade_window]`
* If `local_time > rollout_time`, indicate Boyar to upgrade.
* On a service restart, always upgrade to the newest available revision.
* Node services: Service1, Service2 rollout according to the same `rollout_time`.

Notes:
* No strong assumptions on management service sampling time, assume ~ 1 min.
* Assumed reset time ~ 5-10 min.
  * Impact the urgent `rollout_window`.
* Upgrade decision is a local, uncoordinated decision, works also for audit nodes.
* Recoverable - no need for persistent state for the upgrade status.

### Extra definitions

* `ExternalPort` per virtual chain is defined as such:<br>
    `vc_index = VirtualChainId - 1000000`<br>
    `ExternalGossipPort = vc_index + GOSSIP_PORT_BASE`<br>
    `GOSSIP_PORT_BASE = 10000`
    
* `NodeServices` dictionary keys (names of the services) are defined under [Image naming and tagging](../version-release/NAMING.md).
    * These keys also define the docker image names in the docker registry:<br>
        ```
        {DOCKER-NAMESPACE}/{SERVICE-NAME}
        ```
        For example: `orbsnetwork/signer`
    * These keys also define internal network endpoints:
        ```
        http://{SERVICE-NAME}:{INTERNAL-PORT}
        ```

        For example: `http://management-service:8080`

### Private Network

Reference Time - Consensus in private:
* Consensus in private chains where the data may be updated by the operators at different times is based on agreement on `block.timestamp` as a common ground.
* `VC referenceTime = min(JSON.referenceTime, block.timestamp)`
  * No impact on public instances as block.timestamp is at least Ethereum finality time before the referenceTime.
* Update entries are expected to be added prior to taking effect by at least a quorum of validators.
* New configurations will take effect when block.timestamp reaches the configuration.
* As there is consensus on the block timestamp, there’s no need for grace among the validators, therefore `REFERENCE_TIME_CONSENSUS_GRACE = 0`.

CurrentRefTime:
* The reference of the latest entry that was updated. Including implicit entries. Example:
  * A committee member will be added on 1/1/2021 12AM; ReferenceTime = 1609459200

