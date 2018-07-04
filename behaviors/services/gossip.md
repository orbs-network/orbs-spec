# Gossip

Connects different nodes over the network with efficient message broadcast and unicast. This is the main method nodes in the network use to communicate among themselves.

Inside a node: A pub/sub model where any service subscribes with Gossip to a specific topic (group of messages, all topics are described [here](../../interfaces/protocol/gossip/message.proto)).

Currently a single instance per virtual chain per node.

#### Interacts with services

* None - Passive, just provides services to others upon request.

&nbsp;
## `Network Requirements`

* If the number of nodes is small:
  * Fully connected, every node is connected to every other node.
  * All communication is direct.
    * Optimization: Forwarding scheme for large messages like `PRE_PREPARE`.
* Network topology is known in advance and [configurable](../config/shared.md).
  * Currently assume all services are found on all nodes.
  * Currently assume single instance of a service on a node.
* No special support for retransmission except standard TCP stack.
* Recommended: Socket connections should be signed by nodes and authenticated when opened (via TLS).
  * Not a hard requirement because all communication is signed in the application level.
* Reconnect to topology peers when disconnected with a [configurable](../config/services.md) polling interval.
* Calls `OnMessageReceived` when a gossip message is received.
* Wire format encoding is [Protobuf over TCP](../../encoding/gossip/protobuf-over-tcp.md).

&nbsp;
## `Data Structures`

#### Network address map
* Maps between node id (public key) to an active socket to the Gossip service of this node.
* Used for inter node communication.
* Assumes every node has a single gossip instance.
* Initialized based on the public gossip endpoints from node topology [configuration](../config/shared.md).

#### Topic subscription table
* Map between topic to list of service ids that are subscribed to this topic.
* Used for intra node communication.
* Assumes a single instance of every service on every node.
* Generated dynamically by calls to `TopicSubscribe`.

&nbsp;
## `Init` (flow)

* Initialize the [configuration](../config/services.md).
* Connect to the gossip endpoints of all relevant peer nodes.

&nbsp;
## `TopicSubscribe` (method)

> Used by services to subscribe to a specific topic.

* Add this service id to the the topic subscription table for the requested topic.

&nbsp;
## `TopicUnsubscribe` (method)

> Used by services to unsubscribe from a specific topic.

* Remove this service id from the the topic subscription table for the requested topic.

&nbsp;
## `OnMessageReceived` (method)

> Called internally when a gossip message is received from another node (inter node).

* Lookup the the list of subscribed services for this topic in the topic subscription table.
* For each service id:
  * Rely on service topology [configuration](../config/shared.md) to locate the service endpoint by service id.
  * Call `Target.GossipMessageReceived` with message content.
  * The target service is responsible to identify the message type and process the message accordingly.

&nbsp;
## `SendMessage` (method)

> Sends an inter node message, the message will be forwarded to the services subscribed to the topic on the receiving node.

* The gossip message holds the list of destination nodes ids (public keys).
  * Setting `inverse_recipients` sends to all nodes in the virtual chain except the destination node ids.
  * `NULL` value for the list means broadcast to all nodes in the virtual chain.
* For each node id:
  * Rely on the network address map to locate the relevant socket.
  * Send the message on the socket.
