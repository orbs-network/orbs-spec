# Consensus Algo

Implements a raw consensus algorithm. The system supports pluggable consensus, meaning multiple consensus algorithms running side by side on different virtual chains. The ConsensusAlgo abstracts away the details of how the consensus algorithm works and provides the system with a single unified interface.

Communicates with the system through the ConsensusCore service. Controls all the timing of consensus, for example when new blocks are committed or when new blocks need to be proposed and populated with transaction from TransactionPool.

Currently a single instance per virtual chain per node.

&nbsp;
## `Init Flow`
* Subscribe to consensus messages.
* Init view, term to 0.

&nbsp;
## `Lite-Helix Messages`
* When Lite-Helix sends a broadcast message call `Gossip.BroadcastMessage` with the Lite-Helix message, send message to the Ordering_Nodes group.
* When Lite-Helix sends a unicast message call `Gossip.UnicastMessage` with the Lite-Helix message.

## `Lite-Helix Flow`
Init() - View = 0, Term = 0, Leader = first node.

`Append(Block)`
* If leader, Broadcast[PrePrepare, view, term, Block, H(Block), NodeID, Sig]
  * NodeID = Node Public Key
  * H(Block) = SHA256(Execution Validation Block Header)
* Upon reception of PrePrepare message:
  * Check PrePrepare message: view, term, hash, signature
  * Validate Ordering Block
  * Validate Execution Validation Block
  * If all valid, Broadcast[Prepare, view, term, H(Block), NodeID, Sig]
* Upon reception of Prepare message:
  * Check Prepare message: view, term, hash, signature
  * If valid, mark prepare received from nodeID
  * Upon reception of valid Prepare messages from 2f-1 (+1 for PrePrepare, +1 for the node) Broadcast[Commit, view, term, H(Block), NodeID, Sig]
* Upon reception of Commit message:
  * Check Commit message: view, term, hash, signature
  * If valid, mark prepare received from nodeID
  * Upon reception of valid Commit messages from 2f+1 nodes:
    * Call `BlockCommitted(Block)`
    * If leader, create a new block
* Upon timer experation
  * ...


&nbsp;
## `GossipMessageReceived` (method)
> Handles messages received from another node, expect to receive only Consensus messages that cosnesus-algo have subscribed to.
