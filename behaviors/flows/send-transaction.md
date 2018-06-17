# Send Transaction Flow

The client sends a transaction on a service that requires consensus and may write to state. The transaction is limited in execution to a time window. This is to avoid duplication (will not execute twice) and to kill transactions that are pending execution for too long.

The response is synchronous, so if the node takes a short while to figure out the response, the client blocks. Processing requires an active subscription on the virtual chain.

Transaction is processed under consensus (this is part of the [continuous block creation flow](block-creation.md) as this flow ends when the transaction has been added to the pending pool and propagated to all nodes). Transactions are performed serially since their side effects can influence one another.

## Participants in this flow

* Client
  * `ClientSdk`

* Gateway node
  * `PublicApi`
  * `TransactionPool`
  * `VirtualMachine`
  * `Processor`
  * `StateStorage`
  * `CrosschainConnector`
  * `Gossip`

* All other nodes
  * `Gossip`
  * `TransactionPool`

## Assumptions for successful flow

* No assumptions on synchronization.

## Flow

* `ClientSdk` sends request to `PublicApi`.

* `PublicApi` of gateway node:
  * Adds the transaction as pending to `TransactionPool`.

  * `TransactionPool` of gateway node:
    * Executes pre order checks by calling `VirtualMachine`.

    * `VirtualMachine` of gateway node:
      * Executes the subscription check smart contract on the native `Processor`.
      * Depending on contract code may reads state from `StateStorage` or `CrosschainConnector`.

    * Adds transaction to pending transaction pool.
    * Prepares a batch of transactions for gossip and signs them as their gateway.
    * Broadcasts the batch to all nodes with `Gossip`.

    * `TransactionPool` of all nodes:
      * Checks the batch signature of the gateway and adds transactions to pending pool.
      * Waits until the transaction is committed to the blockchain under consensus (added to a new block).
        * This is part of the [continuous block creation flow](block-creation.md)).
        * This process may take a few seconds.
      * After it is committed, it's moved from the pending pool to the committed pool in all nodes.

    * Returns the result to `PublicApi`.

  * Responds to the client.

&nbsp;
## Send Transaction Flow Diagram

![alt text][send_transaction_flow] <br/><br/>

[send_transaction_flow]: ../_img/send_transaction_flow.png "Send transction"
