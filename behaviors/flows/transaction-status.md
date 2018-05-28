# Transaction Status Flow <!-- tal will finish -->

The client queries the status for a transaction. (tx_id and timestamp)

The response is synchronous, so if the node takes a short while to figure out the response, the client blocks. This read does not require an active subscription on the virtual chain.

This query is performed on the receiving node and does not requrie a consensus or communication with other nodes.

Returns the tarnsaction receipt for committed transactions, otherwise returns the transaction status.

## Participants in this flow

* Client
  * `ClientSdk`

* Receiving node
  * `PublicApi`
  * `Transaction Pool`
  * `Block Storage`

## Assumptions for successful flow

## Flow

* `ClientSdk` sends query to `PublicApi`.

* `PublicApi` of receiving node:
  * Retains session info so the response will eventually arrive to this client.
  * Check the transaction timestamp, verifying that it's not a future transaction plus grace (eg. 5 seconds).
  * Queries the transactions pool by calling `TransactionPool.GetTransactionReceipt`.
    * If found returns PENDING, if committed return COMMITTED with the receipt.
  * Queries the block storage by calling `BlockStorage.GetTransactionReceipt`.
    * If found returns COMMITTED with the receipt, else returns NO_RECORD_FOUND.

* `BlockStorage`
  * Go over all the bloom filters in the blocks where the transaction could be, based on its timestamp and the configurable time_window.
    * Lookup the transaction timestamp in the block header timestamp bloom filter and the tx_id in the block header tx_id bloom filter.
    * On match, fetchs the block and search the tx_id in the block receipts.
        * If found, returns the receipt
  * If not found on all relevant blocks, returns NULL.
