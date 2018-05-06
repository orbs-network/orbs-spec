# Consensus

Implements the consensus algorithm, controlling the flow of transactions across nodes for ordering, execution, and validation until they are committed to block storage and their state effects are committed to state storage.

Stub implementation before the official consensus implementation is using Raft.

&nbsp;
## `Block Creation`

* Continuous process performed by the leader only (by ticks in configurable interval).
* Always synchronized with the latest block in block storage by calling `BlockStorage.GetLastBlock`.
* All pending transactions are taken from pending transaction pool by calling `TransactionPool.GetAllPendingTransactions`.
* They are all executed by calling `VirtualMachine.ProcessTransactionSet`, creating a state diff.
* New block is created.
  * Current protocol version.
  * Pointing to the previous block (the latest) hash.
  * Block height is incremented from previous block (the latest).
  * Contains the state diff from the execution.
* Appended to Raft by sending the "append" Raft message.

&nbsp;
## `Raft Messages`

* When Raft needs to broadcast a message call `Gossip.BroadcastMessage` with a "RaftMessage" message.
* When Raft needs to unicast a message call `Gossip.UnicastMessage` with a "RaftMessage" message.

&nbsp;
## `OnLeaderElected` (method)
> Handle Raft "leaderElected" event (called when a new leader is elected)

* New leader node starts block creation process, old leader stops.

&nbsp;
## `OnCommitted` (method)
> Handle Raft "committed" event (called when a new block has been committed by the leader)

* Add block to block storage by calling `BlockStorage.AddBlock`.

&nbsp;
## `GossipMessageReceived` (method)
> Handle a gossip message from another node

#### On "RaftMessage" message
* Forward message to Raft implementation using Raft "received" event.
