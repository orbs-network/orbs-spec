# Transaction Pool

Holds pending transactions until added to a block, and keeps committed transactions for a time window to avoid duplication. Proposes transactions for new blocks and validates proposals by other nodes.

Currently a single instance per virtual chain per node.

#### Interacts with services

* `VirtualMachine` - Uses it to perform pre order checks on transactions (like subscription is valid).
* `PublicApi` - Notifies it when transactions it's the gateway for have completed.
* `Gossip` - Uses it to communicate with other nodes.

&nbsp;
## `Data Structures`

#### Pending transaction pool
* Holds transactions until they are added to a block and helps preventing transaction duplication.
* Needs to support efficient query by `txhash`.
* Needs to be sorted by time to allow preparing block proposals according to policy.
* Associates every transaction with the node id (public key) of the gateway that added it to the network.
  * Must only hold a single copy of a transaction (`txhash`) regardless of its associated gateway node (the first that added it).
* No need to be persistent, pending transactions may be lost without impact.
* [Configurable](../config/services.md) max size.
* [Configurable](../config/services.md) interval to clear expired transactions.
  * Transaction is expired if its timestamp is later than current time plus the [configurable](../config/shared.md) expiration window (eg. 30 min).
  * Notify public api about transactions it needs to respond to:
    * If we are marked as the gateway for this transaction in the pending pool, it was originated by the node's public api.
    * If indeed local, update the registered public api service by calling its `HandleTransactionsBlock`.
      * Provide block height and timestamp according to the last committed block.

#### Committed transaction pool
* Holds the receipts of committed transactions to return their result and avoid transaction duplication.
  * Also holds the block height and block timestamp for every transaction.
* Needs to support efficient query by `txhash`.
* No need to be persistent, can re-sync from block storage.
* No limit on max size (depends on expiration window).
* [Configurable](../config/services.md) interval to clear expired transactions.
  * Transaction is expired if its timestamp is later than `last_committed_block` timestamp plus the [configurable](../config/shared.md) expiration window (eg. 30 min).

#### Synchronization state
* `last_committed_block` - The last valid committed block that the transaction pool is synchronized to (persistent if the pools are).

&nbsp;
## `Init` (flow)

* Initialize the [configuration](../config/services.md).
* Load persistent data.
* If no persistent data, init `last_committed_block` to empty (symbolize the empty genesis block) and the pools to empty.
* Subscribe to gossip messages in topic `TRANSACTION_RELAY` by calling `Gossip.TopicSubscribe`.

&nbsp;
## `AddNewTransaction` (method)

> Add a new transaction from a client to the network (propagate to all pending transaction pools). Runs on the gateway node.

#### Check transaction validity
* Correct protocol version.
* Valid fields:
  * Signer - known scheme, valid key length.
<!--  * Contract name - TODO . -->
* Sender virtual chain matches the transaction pool's virtual chain.
* Check Transaction timestamp:
  * Only accept transactions that haven't expired.
    * Transaction is expired if its timestamp is earlier than current time minus the [configurable](../config/shared.md) expiration window (eg. 30 min).
  * Only accept transactions with timestamp in sync with the node (that aren't in the future).
    * Transaction timestamp is in sync if it is earlier than the last committed block timestamp + [configurable](../config/services.md) sync grace window (eg. 3 min).
    * Note that a transaction may be rejected due to either future timestamp or node's loss of sync.
* Transaction (`txhash`) doesn't already exist in the pending pool or committed pool (duplicate).
* Verify pre order checks (like signature and subscription) by calling `VirtualMachine.TransactionSetPreOrder`.
* On any failure, return the relevant error status and an empty receipt.
  * For an already committed transaction, return the receipt.

#### Add transaction to pending pool
* Add transaction to pending transaction pool if pool is not full.
* Associate the transaction in the pool with the local node id (public key) as gateway.

#### Batch before broadcast
* Batch until a [configurable](../config/services.md) number of transactions is reached or until a [configurable](../config/services.md) time passes.
* Sign the batch with the local node private key.
* Broadcast the batch to all other transaction pools by calling `Gossip.SendMessage`.

&nbsp;
## `OnForwardedTransactions` (method)

> Add a transaction batch forwarded from another node to the pending transaction pool. Part of transaction propagation.

#### Check batch validity
* Check the batch signature, discard on error.

#### Add transaction to pending pool
* Add transaction to pending transaction pool if pool is not full.
* Associate the transaction in the pool with the batch node id (public key) as gateway.

&nbsp;
## `GetTransactionsForOrdering` (method)

> Proposes a set of N transaction for block building based on the block building policy (currently first come first served). Called by the block builder during consensus.

#### Check synchronization status
* If requested block height is in the future but `last_committed_block` is close to it ([configurable](../config/services.md) sync grace distance) block the call until requested height is committed. Or fail on timeout.
* If requested block height is in the future but `last_committed_block` is far, fail.
* If requested block height is in the past, panic.

#### Create proposal
* Prepare a transactions list proposal (current policy is first come first serve).

#### Check transactions validity
* For each transaction:
  * Correct protocol version.
  * Valid fields (sender address, contract address).
  * Sender virtual chain matches contract virtual chain and matches the transaction pool's virtual chain.
  * Check transaction timestamp, accept only transactions that haven't expired.
    * Transaction is expired if its timestamp is earlier than current time minus the [configurable](../config/shared.md) expiration window (eg. 30 min).
* Transaction wasn't already committed (exist in the committed pool).
* Verify pre order checks (like signature and subscription) for all transactions by calling `VirtualMachine.TransactionSetPreOrder`.

&nbsp;
## `ValidateTransactionsForOrdering` (method)

> Verifies that an ordered list of transactions complies with the ordering policy, called by the block builder during consensus when validating a new block proposal.

#### Check synchronization status
* If requested block height is in the future but `last_committed_block` is close to it ([configurable](../config/services.md) sync grace distance) block the call until requested height is committed. Or fail on timeout.
* If requested block height is in the future but `last_committed_block` is far, fail.
* If requested block height is in the past, panic.

#### Check transactions validity
* For each transaction:
  * Correct protocol version.
  * Valid fields (sender address, contract address).
  * Sender virtual chain matches contract virtual chain and matches the transaction pool's virtual chain.
  * Check transaction timestamp, accept only transactions that haven't expired.
    * Transaction is expired if its timestamp is earlier than current time minus the [configurable](../config/shared.md) expiration window (eg. 30 min).
* Transaction wasn't already committed (exist in the committed pool).
* Verify pre order checks (like signature and subscription) for all transactions by calling `VirtualMachine.TransactionSetPreOrder`.

#### Check proposal policy
* Future: Verify that the proposal matches the selection policy (current policy is first come first serve).

&nbsp;
## `CommitTransactionsReceipts` (method)

> Commits a new approved block (the transaction receipts from it is what we care about). This is the main way to update the transaction pool as new blocks are generated and approved in the system.

#### Check block height
* We assume here that the caller of this method inside the node is trusted and has already verified the block.
* We want to commit blocks in sequence, so make sure the given [block height](../../terminology.md) is the next of `last_committed_block`.
* If not, discard the commit and return the next desired block height (which is the next of `last_committed_block`).
* Optimization: If transaction pool missed a day, it only needs to receive the last configurable expiration window (eg. 30 min). The beginning of the day can be skipped. This will probably require a new method from block storage that translates past time stamp to block height.

#### Ignore blocks with an expired timestamp
* If the block timestamp is beyond the expiration time (and so are its transactions), it can be ignored.
  * Transaction is expired if its timestamp is earlier than current time minus the [configurable](../config/shared.md) expiration window (eg. 30 min).
  * If the pools contain any expired transactions, they have periodical timers that clear them.

#### Commit receipts
* For each transaction receipt:
  * Add the receipt, block height and block timestamp to the committed pool.
  * Notify public api about transactions it needs to respond to:
    * If we are marked as the gateway for this transaction in the pending pool, it was originated by the node's public api.
    * If indeed local, update the registered public api service by calling its `HandleTransactionsBlock`.
  * Remove the corresponding transactions (based on their `txhash`) from the pending pool.
* Update `last_committed_block` to match the given block.

&nbsp;
## `GetTransactionReceipt` (method)

> Returns the transaction receipt for a past transaction based on its id. Used when a client asks to query transaction status for an older transaction.
* If `txhash` is present in the pending transaction pool, return status `PENDING` along with the last committed block height and timestamp.
* If `txhash` is present in the committed transaction pool, return status `COMMITTED` and the receipt.
* else return status `NO_RECORD_FOUND` along with the last committed block height and timestamp.

&nbsp;
## Gossip Messages Handlers

> Handles gossip messages from other nodes. Relevant messages include transaction relay messages.

#### `HandleForwardedTransactions`
* Handles `TRANSACTION_RELAY_FORWARDED_TRANSACTIONS` messages, see `send-transaction` flow.



<!-- TODO: oded, add the diagrams again

![alt text][transaction_preorder_validation_flow] <br/><br/>

[transaction_preorder_validation_flow]: transaction_preorder_validation_flow.png "PreOrder Validation Flow"

-->
