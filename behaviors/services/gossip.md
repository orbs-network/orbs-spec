# Gossip

Connects different nodes over the network with efficient message broadcast and unicast. This is the main method nodes in the network use to communicate among themselves.

Intra node: A pub/sub model where any service subscribes with Gossip to a specific topic (group of messages, all topics are described [here](../../interfaces/protocol/messages.proto)).

Currently a single instance per virtual chain per node.

&nbsp;
## `Network Requirements`

* Fully connected, every node is connected to every other node.
* Network topology is known in advance and configurable.
  * Assume all services are found on all nodes.
  * Assume single instance of a service on a node.
* All communication is direct.
  * TODO Possible exception: forwarding scheme for large (PrePrepare) messages.
* No special support for retransmission except standard TCP stack.
* Recommended: All messages should be signed by nodes and authenticated upon reception (via TLS).
* Reconnect to topology peers when disconnected with a configurable polling interval.
* Call `OnMessageReceived` when a gossip message is received.
* Wire format encoding is [Protobuf over TCP](../../interfaces/protocol/encoding/gossip/protobuf-over-tcp.md).

&nbsp;
## `Data Structures`

#### Network Address Map
* Maps between NodeID (Public Key) to an active socket to the Gossip service of this node.
* Used for inter node communication.
* Assumes every node has a single gossip instance.
* Initialized based on the public gossip endpoints from node topology configuration.

#### Topic Subscriptions Table
* Map between topic to list of service ids that are subscribed to this topic.
* Used for intra node communication.
* Assumes a single instance of every service on every node.
* Generated dynamically by calls to `TopicSubscribe`.

&nbsp;
## `Init` <!-- oded will finish -->

TODO

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

> Called when a gossip message is received from another node (inter-node)

* Look up the the list of subscribed services for this topic in the topic subscription table.
* For each service id:
  * Rely on service topology configuration to locate the service endpoint by service id.
  * Call `Target.GossipMessageReceived` with message content.
  * The target service is responsible to identify the message type and process the message accordingly.

&nbsp;
## `SendMessage` (method)

> Sends an inter node message, the message will be forwarded to the services subscribed to the topic.

* The gossip message holds the list of destination nodes ids (node ids = public keys).
  * NULL value for the list means broadcast to all nodes in the virtual chain.
* For each node id:
  * Rely on Network Address Map to locate the relevant socket.
  * Sends the message on the socket.
