# Get Receipt Proof Flow

The client requests a receipt along witha proof for a previously sent transaction by providing its `tx_id` and timestamp (that was found inside the transaction itself).

The response is synchronous, so if the node takes a short while to figure out the response, the client blocks. This read does not require an active subscription on the virtual chain.

This read is not under consensus. Multiple queries can take place at the same time as this is fully parallel.

## Participants in this flow

* Client
  * `ClientSdk`

* Gateway node
  * `PublicApi`
  * `TransactionPool`
  * `BlockStorage`

## Assumptions for successful flow

* `BlockStorage` is synchronized to latest committed block.

## Flow

* `ClientSdk` sends a request to `PublicApi`.

* `PublicApi` of gateway node:
  * Queries the transaction status. See [Get Transaction Status Flow](../transaction-status.md)
    * By querying the `TransactionPool` and `BlockStorage`
  * If found, requests the `BlockStorage` to prepare a proof.

  * `BlockStorage` of gateway node:
    * Fetches the receipt block.
    * Calculates the receipts merkle tree and generate a merkle proof for the receipt.
    * Returns the block header, proof and the receipt merkle proof.

  * Responds to the client.
