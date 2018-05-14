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
> Handles messages received from another node. 
* If the message is a consensus allgorithm message, pass to algorithm.
* If the message is a transaction batch
  * Validate batch signature, if mismatch log for reputation calculation.
  * Add transactions to pending pool by calling `TransactionPool.AddNewPendingTransaction`.

