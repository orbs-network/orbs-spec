# Send Transaction Flow

The client sends a transaction that requires consensus and may write to state. The transaction is limited in execution to a time window. This is to avoid duplication (will not execute twice) and to kill transactions that are pending execution for too long.

The response is synchronous, so if the node takes a short while to figure out the response, the client blocks. Processing requires an active subscription on the virtual chain.

Transaction is processed under consensus (this is part of the [continuous block creation flow](block-creation.md) as this flow ends when the transaction has been added to the pending pool and propagated to all nodes). Transactions are performed serially since their side effects can influence one another.

## Participants in this flow

* Client
  * `ClientSdk`

* Receiving node
  * `PublicApi`

* All nodes
  * `TransactionPool`
  * `VirtualMachine`
  * `Processor`
  * `StateStorage`
  * `SidechainConnector`
  * `Gossip`

## Assumptions for successful flow

* No assumptions on synchronization.

## Flow

* `ClientSdk` sends request to `PublicApi`.

* `PublicApi` of receiving node:
  * Retains session info so the response will eventually arrive to this client.
  * Performs inexpensive checks on the transaction.
  * Sends the transaction to `TransactionPool`.

* `TransactionPool` of receiving node:
  * Performs inexpensive checks on the transaction.
  * Executes pre order checks by calling `VirtualMachine.TransactionSetPreOrder`.

* `VirtualMachine` of receiving node:
  * Checks the sender signature.
  * Executes the subscription check smart contract on the native `Processor`.
  * Reads relevant state for execution from `StateStorage` or `SidechainConnector`.

* `TransactionPool` of receiving node:
  * Adds transactions to pending transaction pool.
  * Prepares a batch of transactions for gossip and signs them as their gateway.
  * Broadcasts the batch to all nodes with `Gossip`.

* `Gossip` of all nodes:
  * Checks the batch signature of the gateway.
  * Adds transactions to pending `TransactionPool` and affiliates them with their gateway.

* The transaction is now waiting to be added to a block ([continuous block creation flow](block-creation.md)).
  * After it was added, it's deleted from pending `TransactionPool` of all nodes.
  * The `TransactionPool` of receiving node returns the result to `PublicApi`.

* `PublicApi` of receiving node responds to the client.

## Flow diagram

TODO
