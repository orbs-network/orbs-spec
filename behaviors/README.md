# High Level Flows

&nbsp;
## Client sends a transaction (write)

* Client request arrives to `PublicApi` through a public protocol.
* `PublicApi` 
  * Performs checks on the transaction:
    * Version, Transaction format, virtual chain.
  * Maintain session info.
  * Sends transaction to `Transaction Pool`
* `Transaction Pool` 
  * Performs transction checks. 
  * Execute a pre order contract by forwarding to `VirtualMachine`.
    * State needed for execution is read from `StateStorage`.
    * Checks virtual chain subscription.
  * Checks that its state is in sync.
  * Transactions valid for ordering are batched together, signed and forwarded to `Gossip` to be broadcast to all `Transaction Pools`.
* `Gossip` forwards the transaction batch.
  * TBD move signature, virtual chain check to `Gossip`
* Receiving `Transaction Pool`
  * Checks the batch signature of the originating `Transaction pool` and affiliates the transactions to it.
  * Checks the virtual chain.
* The transaction is now waiting until one of the nodes adds it to a block.
* `PublicApi` waits for transaction response from the `Transaction pool` in order to send response to client.


&nbsp;
## Client calls a contract (read)
> Enables to perform Read Only contract calls. Perfromed on the latest state.
* Client request arrives to `PublicApi` through a public protocol.
* `PublicApi` 
  * Performs checks on the transaction:
    * Version, Transaction format, virtual chain.
  * Maintain session info.
  * Batch up to 100ms or 10 transactions. 
  * Get the latest block height by calling `ConsensusCore.GetTopBlockHeight`.
  * Send the transactions to VM for execution by calling `VirtualMachine.CallContract` with the latest block height. 
  * `VirtualMachine`
    * Checks the Sender and Signatrue (if present: <> 0). If signature isn't present assume sender = 0.
    * Executes the smart contract
    * Returns the result
  * `PublicApi` sends a response to the client.


&nbsp;
## Continuous block creation

* `Consensus Algorithm` intiates a new consensus round.
  * The `Consensus Algorithm` queries the `Consensus Core` on the participating committee (all nodes)
  * The instance of the leader requests the `Consensus Core` to build ordering and validation blocks.
* The `Consensus Core` builds a new ordering block
  * Requests pending transactions for the `TransactionPool`
  * The `TransactionPool` 
    * Perfroms PreOrder checks, similart to the ones done by the node that inserted them to the network.
    * Executes a pre order contract by forwarding to `VirtualMachine`.
    * Verifies no duplication with already committed blocks.
    * Returns a list of ordered transactions.
* The `Consensus Core` builds a new validation block
  * Executes the transactions by forwarding them to `VirtualMachine`.
    * The `VirtualMachine` reads the requried state from the `StateStorage`.
      * The `StateStorage` verifies that it's in sync with the requested block_height reference.
    * The `VirtualMachine` generates transaction receipts and state diff (that isn't committed yet).
  * Requests the `State Storage` for the root hash before the current block execution.
* The `Consensus Core` maintains a copy of the ordering and validation blocks waiting for them to be committed by the `Consensus Algo`
* The `Consensus Algo` of the leader instance receives the ordering block and the validation block from the `Consensus Core` and broadcast them using `Gossip` (PrePrepare).
* The `Consensus Algo` of the non-leader nodes receive the proposed block. and forward them to the `Consensus Core` for validation.
* The `Consensus Core` validates the ordering block by forwarding it to the `TransactionPool`.
  *  The `TransactionPool`
    * Perfroms PreOrder checks, similar to the ones done by the node that inserted them to the network.
    * Executes a pre order contract by forwarding to `VirtualMachine`.
    * Verifies no duplication with already committed blocks.
* The `Consensus Core` validates the validation block.
  * Executes the transactions by forwarding them to `VirtualMachine`.
    * The `VirtualMachine` reads the requried state from the `StateStorage`.
      * The `StateStorage` verifies that it's in sync with the requested block_height reference.
    * The `VirtualMachine` generates transaction receipts and state diff and the `Consens core` matches their root hash with the one of the proposed valdiation block.
  * Requests the `State Storage` for the root hash before the current block execution and matches it with the proposed validation block.
* The `Consensus Core` maintains a copy of the ordering and validation blocks and returns a valid indicaion to the `Consensus Algo` taht broadcast it to all nodes using `Gossip` (Prepare)
* The `Consensus Algo` complites the consensus round (Commit) and once a block is Committed, indicates it to the `Consensus Core` and forward the committed block proofs. 
* The `Consensus Core` attach the block proofs to the ordering and validation blocks, commits them to the `Block Storage` and updates its last_commited_block indications.
* The `Block storage` commits the new validation block state diff to the `State Storage`.
  * The `State Storage` updates the state along with the merkle root.
* The `Block storage` commits the new validation block receipts to the `Transaction Pool`.
* The `Transaction Pool`
  * Looks up the transaction IDs in the pending pool, if the transactions are associated with a `Public API` that is subscribed to the `Transaction Pool` instance (local `Public API`), it sends  the receipt to the `Public API`.
    * The `Public API` reuturns a respond to the Client.
  * Removes the transactions from the pending transction pool.
  * Add the transaction IDs to the commmitted pool.

&nbsp;
## Continuous state synchronization
* The `State Storage` subscribes for state diff commits from the `Block Storage`.
* Any query that is initiated by the `Consensus Core` refers to a block height.
* The `State Storage` verifies that it is synced with the block height in the query.
* If the `State Storage` identifies that it is out of sync, it updates its request subscription and requests blocks starting after its last synched height.
* The `Block Storage` commits the required state diff up to the latest block and continues to commit the state diff on each new block commit.

&nbsp;
## Continuous transaction pool synchronization
* The `Transaction Pool` subscribes for receipts commits from the `Block Storage`.
* Any query that is initiated by the `Consensus Core` refers to a block height.
* The `Transaction Pool` verifies that it is synced with the block height in the query.
* If the `Transaction Pool` identifies that it is out of sync, it updates its request subscription and requests blocks starting after its last synched height.
* The `Block Storage` commits the required receipts up to the latest block and continues to commit receipts on each new block commit.

&nbsp;
## Continuous block synchronization
* Block sync is performed by the `Block Storage`.
* The `Block Storage` identifies out of sync when an out of sync commit it received from the `Consensus Core` or when a subscription request is received from the `State Storage` or `Transaction Pool` to a block that isn't present.
* The `Block Storage` verfies the laetst block height by quering the `Consensus Core`.
* If out of sync, the `Block Storage` sends a request for sync to a random peer node using `Gossip`.
* 
* To be completed.

&nbsp;
## Client Queries Transaction Status

&nbsp;
## System Init
