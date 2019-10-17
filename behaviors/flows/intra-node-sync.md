# Continuous Intra Node Block Synchronization Flow

The source of truth for block synchronization in the node is `BlockStorage`. Other node services must be synchronized with the latest block as well, to fulfill their functions.

Block synchronization intra node takes place by push from `BlockStorage` to each of the other services. Whenever a new block is validated and committed to `BlockStorage`, it pushes the block to the other services. If the services are not in sync, it's the responsibility of `BlockStorage` to push all of the missing blocks to them until they are synced.

## Participants in this flow

* Synchronizer
  * `BlockStorage`

* Receiving services
  * `StateStorage`
  * `TransactionPool`

## Assumptions for a successful flow

* `BlockStorage` is synchronized to the latest block.

## Flow

* `BlockStorage` gets a new block committed and pushes the block to the other services.

* Receiving service asks to be synchronized if missing blocks (the pushed block isn't the next needed).

* `BlockStorage` synchronizes the service as needed with all missing blocks.
