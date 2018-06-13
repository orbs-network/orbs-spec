# Get Transaction Status Flow

The client queries the status for a previously sent transaction by providing its `tx_id` and timestamp (that was found inside the transaction itself).

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

* `ClientSdk` sends query to `PublicApi`.

* `PublicApi` of gateway node:
  * Looks for the transaction in `TransactionPool`.
  * If not found, looks for transaction in `BlockStorage`.

  * `BlockStorage` of gateway node:
    * Goes over all the blocks where the transaction could be found according to timestamp (and expiration window).
    * For each block looks for the transaction in the timestamp and id bloom filters.
    * If found, searches for the `tx_id` in the block receipts.

  * Responds to the client.
