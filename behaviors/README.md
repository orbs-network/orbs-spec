# High Level Flows

&nbsp;
## Client calls a contract (read)

* Assuming state storage and block storage are synchronized to latest block.
* Client request arrives to `PublicApi` through a public protocol.
* Contract call is forwarded to `VirtualMachine` for execution.
* State needed for execution is read from `StateStorage`.
* Response is returned to client.

&nbsp;
## Client sends a transaction (write)

* Assuming state storage and block storage are synchronized to latest block.
* Client request arrives to `PublicApi` through a public protocol.
* Transaction is added as pending to `TransactionPool`.
* Transaction is broadcasted to all other nodes through `Gossip`.
* It's now waiting until one of the nodes adds it to a block.

&nbsp;
## Continuous block creation

* Consensus process among nodes chooses a leader for next block.
* The leader fills the block with pending transactions from `TransactionPool`.
* Consensus is performed over block order (block is given a block number).
* Relevant nodes execute the block's transactions when it's the next block.
* Execution takes place on the `VirtualMachine`.
* State needed for execution is read from `StateStorage`.
* Execution results in a state diff (that isn't committed yet).
* Consensus is performed over execution results.
* The block is committed to `BlockStorage`.

&nbsp;
## Continuous state synchronization

* `StateStorage` is continuously building the latest state from `BlockStorage`.
* Waiting until a new block is committed to `BlockStorage`.
* The state diff from the block is committed to `StateStorage`.

&nbsp;
## Continuous block synchronization

* todo
