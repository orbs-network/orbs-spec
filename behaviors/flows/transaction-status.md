# Transaction Status Flow

The client queries the status for a transaction. (tx_id and timestamp)

The response is synchronous, so if the node takes a short while to figure out the response, the client blocks. Processing requires an active subscription on the virtual chain. The client's query isn't signed.

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
  * Check the transaction timestamp, verifying that it's not a future transation.
  * Lookup its own session database for PENDING transactions.
    * If found returns PENDING.
  * Queries the tarnsaction pool's pendong transactions pool by calling `TransactionPool.GetTransactionStatus`.
    * If found returns PENDING.
  * Queries the block storage by calling `BlockStorage.GetTransactionReceipt`.
    * If found returns the transaction receipt data, else returns NO_RECORD_FOUND.
* `BlockStorage`
  * Fetchs the relevant ResultsBlock headers' bloomfilter based on the transaction timestamp and the configurable time_window.
  * Lookup the transaction's timestamp in the block header's timestamp bloom filter and the tx_id in the block header's tx_id bloom filter.
    * On match, fetchs the block and searches the tx_id in the block's receipts.
        * If found, returns the receipt
    * If not found on all relevant blocks, returns NULL.