# Gossip

Connects different nodes over the network with efficient message broadcast and unicast. This is the main method nodes in the network use to communicate among themselves.

Currently a single instance per virtual chain per node.

&nbsp;
## `Network Requirements`

* Fully connected, every node is connected to every other node.
* Network topology is known in advance and configurable.
  * Assume all services are found on all nodes.
  * Assume single instance of a service on a node.
* All communication is direct.
  * TODO forwarding scheme for large (PrePrepare) messages.
* No special support for retransmission except standard TCP stack.
* All messages should be signed by nodes and authenticated upon reception. (Peer2Peer)
* Reconnect to topology peers when disconnected with a configurable polling interval.
* Call `OnMessageReceived` when a gossip message is received.
* See inter node message [encoding](../../interfaces/protocol/gossip/json-over-websocket.md).

&nbsp;
## `Init Flow`
* Read configuration file:
  * Federation nodes data (map of public keys to network address)
  * Broadcast group list for each virtual chain.

&nbsp;
## `OnMessageReceived`
> Called when a gossip message is received from another node, forwarded to the services that are subscribed to the message.
* Call `Target.GossipMessageReceived` with message content.
* The target service is responsible to identify the topic and message type and process the message accordingly.

&nbsp;
## `SendMessage` (method)
> Sends a message to a group a nodes, the message is forwarded to the services subscribed to the topic. Used fro Boradcast, Multicast and Unicast communication.
* The gossip message holds the list of destination nodes.
  * On broadcast indicatation, forward the message to all the nodes of the virtual chain.
* The list of nodes to forward the message to is indicated as part of the gossip message.
* Sends a message to each of the nodes' gossip.

&nbsp;
## `TopicSubscribe` (method)
> Used by services to subscribe to a specific topic.
Returns a token that identifies the subscription and can be used in order to identify the message's subscription within the receving service.

&nbsp;
## `TopicUnsubscribe` (method)
> Used by services to unsubscribe from a specific topic.
