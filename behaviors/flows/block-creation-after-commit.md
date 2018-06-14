# Continuous Block Creation Flow

> Second half: after commit (steps 4-5 out of 5)

*Since the flow is long, this is the second half of the flow: from when the block is committed.*

## Flow

#### Step 4: committee members (with leader) commit the block

`ConsensusAlgo` of committee member:
  * Commits the block in `BlockStorage`.

  * `BlockStorage` of committee member:
    * Stores the block on the persistent chain.
    * Commits the state diff from the Results block to `StateStorage`.
    * Commits the transaction receipts from the Results block to `TransactionPool`.

    * `TransactionPool` of committee member:
      * Moves the transactions from pending to committed pool.
      * If this was the gateway node that added a transaction to the network, notify `PublicApi` to respond to the client.

#### Step 5: non committee members commit the block

`ConsensusAlgo` of non committee member:
  * Validates the consensus of this untrusted block.
  * Gets the block's committee from `ConsensusContext` and verifies the block proofs.
  * Commits the block in `BlockStorage`.

  * `BlockStorage` of non committee member:
    * Stores the block on the persistent chain.
    * Commits the state diff from the Results block to `StateStorage`.
    * Commits the transaction receipts from the Results block to `TransactionPool`.

    * `TransactionPool` of non committee member:
      * Moves the transactions from pending to committed pool.
      * If this was the gateway node that added a transaction to the network, notify `PublicApi` to respond to the client.
