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
* All messages should be signed by nodes and authenticated upon reception.
* Reconnect to topology peers when disconnected with a configurable polling interval.

&nbsp;
## `On Message Received`

* Identify target service according to the message broadcast group.
* Call `Target.GossipMessageReceived` with message content.

&nbsp;
## `BroadcastMessage` RPC

* Send message to each of the nodes.

&nbsp;
## `UnicastMessage` RPC

* Send message to the specific recipient node.
