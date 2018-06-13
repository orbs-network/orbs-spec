# Continuous Block Creation Flow <!-- tal will finish -->

The purpose of consensus is to continuously commit new blocks. Blocks are populated with transactions waiting in TransactionPool.

We currently support leader-based consensus algorithms, meaning a leader node is elected and creates a block proposal which is then passed to validator nodes (members of the block committee) for approval.

## Participants in this flow

* Committee nodes
  * `ConsensusAlgo`
  * `ConsensusContext`
  * `TransactionPool`
  * `VirtualMachine`
  * `Processor`
  * `StateStorage`
  * `SidechainConnector`
  * `Gossip`
  * `BlockStorage`

## Assumptions for successful flow

* `BlockStorage` is synchronized to latest committed block.

## Flow <!-- tal can finish --> <!-- move VM->processor->state flow to a different flow -->

* `Consensus Algorithm` starts a new consensus round (block height is decided).
  * Gets the previously committed block from `BlockStorage` (or have it from cache).

  * The `Consensus Algorithm` queries the `ConsensusContext` on the participating committee (all nodes)
  * The `Consensus Algorithm` identifies if it's a leader or a valdiator in the current term.
* Leader node.
  * The `Consensus Algorithm` requests the `ConsensusContext` to build ordering and validation blocks.
  * The `ConsensusContext` builds a new ordering block
    * Requests pending transactions for the `TransactionPool`
    * The `TransactionPool`
      * Perfroms PreOrder checks, similar to the ones done by the node that inserted them to the network.
      * Verifies no duplication with already committed blocks.
      * Calls the `VirtualMachine` to validate the signatrues and execute the pre-order code.
      * Returns a list of ordered transactions.
  * The `ConsensusContext` builds a new validation block
    * Requests the `State Storage` for the root hash before the current block execution.
    * Executes the transactions by forwarding them to `VirtualMachine`.
      * The `VirtualMachine` reads the requried state from the `StateStorage` and executes by calling the native `Processor`.
        * The `StateStorage` verifies that it's in sync with the requested block_height reference.
      * The `VirtualMachine` generates transaction receipts and state diff.
  * The `ConsensusContext` returns the built block proposal to the `Consensus Algo` and maintains a copy of the ordering and validation blocks waiting for them to be committed.
  * The `Consensus Algo` broadcast the block using `Gossip` (PrePrepare).

* Validating nodes
  * The `Consensus Algo` receive the proposed block from the `Gossip` and forward it to the `ConsensusContext` for validation.
  * The `ConsensusContext` validates the ordering block by forwarding it to the `TransactionPool`.
    *  The `TransactionPool`
      * Perfroms PreOrder checks, similar to the ones done by the node that inserted them to the network.
      * Executes a pre order contract by forwarding to `VirtualMachine`.
      * Verifies no duplication with already committed blocks.
  * The `ConsensusContext` validates the validation block.
    * Executes the transactions by forwarding them to `VirtualMachine`.
      * The `VirtualMachine` reads the requried state from the `StateStorage`.
        * The `StateStorage` verifies that it's in sync with the requested block_height reference.
      * The `VirtualMachine` generates transaction receipts and state diff and the `Consens core` matches their root hash with the one of the proposed valdiation block.
    * Requests the `State Storage` for the root hash before the current block execution and matches it with the proposed validation block.
  * The `ConsensusContext` maintains a copy of the ordering and validation blocks and returns a valid indicaion to the `Consensus Algo` taht broadcast it to all nodes using `Gossip` (Prepare)

* The `Consensus Algo` complites the consensus round (Commit) and once a block is Committed, indicates it to the `ConsensusContext` and forward the committed block proofs.
* The `ConsensusContext` attach the block proofs to the ordering and validation blocks, commits them to the `Block Storage` and updates its last_commited_block indications.
* The `Block storage` commits the new validation block state diff to the `State Storage`.
  * The `State Storage` updates the state along with the merkle root.
* The `Block storage` commits the new validation block receipts to the `Transaction Pool`.
* The `Transaction Pool`
  * Looks up the transaction IDs in the pending pool, if the transactions are associated with a `Public API` that is subscribed to the `Transaction Pool` instance (local `Public API`), it sends  the receipt to the `Public API`.
    * The `Public API` reuturns a respond to the Client.
  * Removes the transactions from the pending transction pool.
  * Add the transaction IDs to the commmitted pool.
