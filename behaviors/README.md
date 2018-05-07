# High Level Flows

&nbsp;
## Client sends a transaction (write)

* Assuming state storage and block storage are synchronized to latest block.
* Client request arrives to `PublicApi` through a public protocol.
* `PublicApi` performs (optimization) checks on the transaction:
  * Stateless checks like signature validations.
  * Execute a pre order contract by forwarding to `VirtualMachine`.
    * State needed for execution is read from `StateStorage`.
  * Checks virtual chain subscription.
* `PublicApi` batches several transactions that passed and signs them.
* `PublicApi` forwards the batch to `Gossip` as a broadcast to `Consensus`. // what if we want to switch to a consensus algorithm that doesn't do a full broadcast?
  * `Gossip` checks the batch signature of `PublicApi`.
  * `Gossip` makes sure this `PublicApi` is valid for the virtual chain (through config).
  * `Consensus` receives message and adds to `TransactionPool` (with affiliation to signing `PublicApi`).
* It's now waiting until one of the nodes adds it to a block.
* `PublicApi` waits for transaction receipt and response is returned to client.

&nbsp;
## Client calls a contract (read)

* Assuming state storage and block storage are synchronized to latest block.
* Client request arrives to `PublicApi` through a public protocol.
* Contract call is forwarded to `VirtualMachine` for execution.
  * State needed for execution is read from `StateStorage`.
* Response is returned to client.

&nbsp;
## Continuous block creation

* `Consensus` participates in ordering.
  * For example, consensus process among nodes chooses a leader for next block.
  * The leader fills the block with pending transactions from `TransactionPool`.
  * Consensus is performed over block order (block is given a block number).
* `Consensus` participates in execution in parallel.
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
