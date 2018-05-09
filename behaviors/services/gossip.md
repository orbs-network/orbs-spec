# Gossip

Connects different nodes over the network with efficient message broadcast and unicast.

&nbsp;
## `Network Requirements`

* Fully connected, every node is connected to every other node.
* Network topology is known in advance and configurable.
  * Assume all services are found on all nodes.
  * Assume single instance of a service on a node.
* All communication is direct.
* No special support for retransmission except standard TCP stack.
* All messages should be signed by nodes and authenticated upon reception. (Peer2Peer)
* Reconnect to topology peers when disconnected with a configurable polling interval.
* Call `OnMessageReceived` when a gossip message is received.
* See inter node message [encoding](../../interfaces/protocol/gossip/json-over-websocket.md).

&nbsp;
## `OnMessageReceived` (method)
> Called when a gossip message is received from another node

* Identify target service according to the message target broadcast group.
* Call `Target.GossipMessageReceived` with message content.

&nbsp;
## `BroadcastMessage` (method)
> Broadcasts a message to all nodes

* Send message to each of the nodes.

&nbsp;
## `UnicastMessage` (method)
> Sends a message to a single nodes

* Send message to the specific recipient node.
