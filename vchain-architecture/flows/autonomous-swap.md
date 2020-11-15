# Autonomous Swap Flow

A client sends transactions to transfer token from an Ethereum ERC20 contract to an Orbs proxy contract amd vise versa. 

<!--
The client sends a transaction on a service that requires consensus and may write to state. The transaction is limited in execution to a time window. This is to avoid duplication (will not execute twice) and to kill transactions that are pending execution for too long.

The response is synchronous, so if the node takes a short while to figure out the response, the client blocks. Processing requires an active subscription on the virtual chain.

Transaction is processed under consensus (this is part of the [continuous block creation flow](block-creation.md) as this flow ends when the transaction has been added to the pending pool and propagated to all nodes). Transactions are performed serially since their side effects can influence one another.
-->

## Participants in this flow

* Client
  * `ClientSdk`

* Node
  * `PublicApi`
  * `TransactionPool`
  * `VirtualMachine` / `Processor`
  * `CrosschainConnector`

* Ethereum Smart Contracts
  * `ERC20`
  * `ASB (Autonomous Swap Bridge`
  * `Orbs Proof Validator`
  * `Federation Contract`

* Orbs Smart Contracts
  * `ERC20Proxy`
  * `ASB (Autonomous Swap Bridge`


## Assumptions for successful flow

* No assumptions on synchronization.

## Deploy Flow

## Ethereum -> Orbs Flow

#### Ethereum 
* `Client` sends an `approve` transaction to Ethereum `ERC20` contract

* `Client` sends an `TransferOut` transaction to the Ethereum `ASB` contract

  * `ASB` Contract `TransferOut`:
    * Assigns a unique `eth_tuid`.
      * Increments `eth_tuid`.
    * Transfers the tokens to the `ASB` account by calling `ERC20.transferFrom`.
    * Logs the transfer event: `TransferOutEvent`.

#### Orbs 
* `ClientSDK` waits the required Ethereum block time until it is considered final. <!-- TODO where/how to implement the wait -->

* `ClientSDK` sends a `TransferIn` transaction to Orbs `ASB` contract.
  * With the Ethereum's transaction `txid`.

  * `ASB` Contract `TransferIn`:
    * Retrieves the `TransferOut` event data by calling `Ethereum.GetTransactionLogs`.

    * `VirtualMachine.SDK`
      * Gets the latest block time.
      * Calls `CrosschainConnector.GetTransactionLogs`.

    * `CrosschainConnector`:
      * Calculates the reference Ethereum block height based on the given timestamp and the finality parameter.
      * Queries the Ethereum node, reading the `TransferOutEvent` events that correspond to the transaction and `ASB` contract.
      * Checks that the transaction's block height is prior to the reference block height.

    * Checks that the `eth_tuid` was not already spent.
    * Transfers the tokens to the requested address by calling the `ERC20Proxy.transfer`.
    * Marks the `eth_tuid` as spent.

## Orbs -> Ethereum Flow

#### Orbs
* `ClientSDK` sends an `approve` transaction to Orbs `ERC20Proxy` contract.

* `ClientSDK` sends an `TransferOut` transaction to the Orbs `ASB` contract.

  * `ASB` Contract `TransferOut`:
    * Assigns a unique `orbs_tuid`.
      * Increments `orbs_tuid`.
    * Logs the transfer event by `TransferOutEvent` by calling the `Logs.Emit` SDK.
    * Transfers the tokens to the `ASB` account by calling `ERC20Proxy.transferFrom`.

    * `VirtualMachine.SDK`
      * Adds the log to the transaction receipt.

* `ClientSDK` sends an `GetTransactionReceiptProof` request to Orbs with the `TransferOut` transaction `txid`.

* `PublicApi` finds the transaction receipt and sends a proof back to the `ClientSDK`
  * This is part of the [get receipt proof flow](get-receipt-proof.md).

* `ClientSDK` sends a `ConstructEthereumProof` call method to the Orbs `ASB` contract.

  * `ASB` Contract `ConstructEthereumProof`:
    * Constract an Ethereum `ASB.TransferIn` transaction payload based on the provided proof.

#### Ethereum 

* `Client` sends a `TransferIn` transaction to Ethereum `ASB` contract.

  * `ASB` Contract `TransferIn`:
    * Calls the `Orbs Proof Validator` contract to verify the proof, see [ethereum-receipt-proof](../data-structures/ethereum-receipt-proof.md).
      * Calls the `Federation contract` to retrieve the proof criteria.
    * Checks that the `orbs_tuid` was not already spent.
    * Transfers the tokens to the requested address by calling the `ERC20.transfer`.
    * Marks the `orbs_tuid` as spent.
    <!-- todo timeout flow -->
