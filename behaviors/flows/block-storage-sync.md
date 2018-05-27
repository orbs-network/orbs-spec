# Continuous BlockStorage Synchronization Flow

BlockStorage needs to always be synchronized to the latest block in the virtual chain (latest committed block height). It can identify when it is potentially out of sync when another notifies it of a later block.

When BlockStorage is out of sync, it attempts to synchronize with other BlockStorage instances of other nodes.

BlockStorage is also charged with synchronizing StateStorage and TransactionPool to the latest block. Whenever they go out of sync, they attempt to synchronize with BlockStorage inside the node.

## Participants in this flow

## Flow

* Block sync is performed by the `Block Storage`.
* The `Block Storage` identifies out of sync when an out of sync commit it received from the `Consensus Core` or when a subscription request is received from the `State Storage` or `Transaction Pool` to a block that isn't present.
* The `Block Storage` verfies the laetst block height by quering the `Consensus Core`.
* If out of sync, the `Block Storage` sends a request for sync to a random peer node using `Gossip`.
*
* To be completed.
