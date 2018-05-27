# Continuous Block Creation Flow

The purpose of consensus is to continuously commit new blocks. Blocks are populated with transactions waiting in TransactionPool.

## Participants in this flow

## Flow

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
