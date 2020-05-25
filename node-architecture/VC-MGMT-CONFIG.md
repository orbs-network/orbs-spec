# Virtual Chain Management Config

> V2 release

## Scope

This configuration provides a virtual chain (ONG instance) with all management information reflecting the state of the network (elected committee, topology, subscription state, ref time, network protocol version). This information changes over time and therefore must be checked repeatedly for changes.

This JSON config format is used both for the public Orbs network (where Ethereum is the raw source of truth for this data) and for private networks (where static JSON files managed by admin are the raw source of truth for this data).

### Ref Time and Events

Some sections of the configuration, such as changes in the elected committee (which nodes are authorized to sign and validate blocks), appear as events on a timeline. Every event indicates the starting ref time of the change.

Ref time is measured in seconds and reflects the external clock of the source of truth of the data. On the public Orbs network the ref time is Ethereum mainnet’s block time. On private networks the ref time is the admin-chosen time when events are written into the static JSON governing the private network. In any case, the current ref time appears as part of the configuration itself (since it is external to the node).

### Current Data vs. Historical Data

The configuration may reflect current data or a snapshot of historic data. This normally depends on which option was requested by its consumer. Historic data is normally used by nodes when they are required to audit past events.

## Configuration Format

* [Example](../config-examples/vc-management.json)

Global fields:

| Field Name | Description |
| ---------- | ----------- |
| `CurrentRefTime` | The current external ref time. For example, when the source of truth is Ethereum blockchain, this reflects the latest Ethereum block time. This field is not provided in snapshots of historic data. |
| `PageStartRefTime` | Since this config response contains multiple events, this is the earliest ref time covered by the response. If the consumer requires earlier events, they should be requested separately (as a new page of historic data). |
`PageEndRefTime` | Since this config response contains multiple events, this is the latest ref time covered by the response. If the consumer requires later events, they should be requested separately (as a new page of historic data). |
| `VirtualChains` | Map of virtual chains by ID containing their respective management config (see fields per virtual chain below). |

Fields per virtual chain:

| Field Name | Description |
| ---------- | ----------- |
| `VirtualChainId` | The numerical ID of the virtual chain. |
| `GenesisRefTime` | The ref time of when this virtual chain was first created. |
| `CurrentTopology` | Details of the nodes that are part of the topology. Based on the topology change events that are relevant to the virtual chain. This field is not provided in snapshots of historic data. A new committee node must be part of topology a few hours before it enters the committee so it has time to sync. A retired committee node must remain part of topology a few hours after it leaves the committee so others have time to sync from it. |
| `CommitteeEvents` | A list of changes in the elected committee of nodes authorized to validate blocks. Events are not incremental. Every event contains the entire state and the ref time where it changed. If the virtual chain subscription is limited to a specific `IdentityType`, only nodes of this type will appear in the committee. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`EffectiveStake` | Combined stake of the validator operating this node and the delegators who have chosen them as guardian. Units are 1K ORBS, notice on the public network that Ethereum events are reported in Orbitons (10^-18 ORBS). |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`IdentityType` | See similar field under subscription events. |
| `SubscriptionEvents` | A list of changes in the virtual chain subscription. Events are not incremental. Every event contains the entire state and the ref time where it changed. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`Status` | Subscription status of the virtual chain.<br>Possible values:<br>&bull;&nbsp;`active` - Subscription is active (paid).<br>&bull;&nbsp;`expired` - Subscription is expired (unpaid). |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`RolloutGroup` | When software versions and protocol versions roll out, they roll out first to certain groups for testing purposes.<br>Possible values:<br>&bull;&nbsp;`canary` - Alpha virtual chains (roll out first).<br>&bull;&nbsp;`ga` - Stable virtual chains (roll out second). |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`IdentityType` | Subscription can be limited to be validated (signed) by nodes that have undergone identity validation.<br>Possible values:<br>&bull;&nbsp;`0` - Not limited (anyone can validate).<br>&bull;&nbsp;`1` - Limited to validators whose identity is verified. |
| `ProtocolVersionEvents` | A list of changes in the official network protocol version. Events are not incremental. Every event contains the entire state and the ref time where it changed. Since the virtual chain subscription indicates a specific RolloutGroup, only protocol version changes for this group will appear. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`RolloutGroup` | See similar field under subscription events. |
| &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;`Version` | Integer describing the protocol version. |
