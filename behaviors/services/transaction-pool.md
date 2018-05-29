# Transaction Pool

Holds pending transactions until added to a block, and keeps committed transactions for a time window to avoid duplication. Proposes transactions for new blocks and validates proposals by other nodes.

Currently a single instance per virtual chain per node.

![alt text][transaction_preorder_validation_flow] <br/><br/>

[transaction_preorder_validation_flow]: transaction_preorder_validation_flow.png "PreOrder Validation Flow"

&nbsp;
## `Init` <!-- oded will finish -->

TODO

&nbsp;
## `Data Structures` <!-- tal can finish -->

#### Pending transaction pool
* Holds transactions until they are added to a block and helps preventing transaction duplication.
* Needs to support efficient query by tx_id.
* Needs to be sorted by time to allow preparing block proposals according to policy.
* Associates every transaction with the node id (public key) of the gateway that added it to the network.
* No need to be persistent, re-sync from the block storage.
* Configurable max size.
* Configurable interval to clear expired transactions by comparing `time_stamp` to a configurable time window (relative to current time).

#### Committed transaction pool
* Holds the receipts of committed transactions to return their result and avoid transaction duplication.
* Needs to support efficient query by tx_id.
* No need to be persistent, re-sync from the block storage.
* No limit on max size (depends on expiration window).
* Configurable interval to clear expired transactions. Transaction is considered expired if transaction timestamp > last block.timestamp + configurable expiration window.

#### Sync state
* `last_committed_block`
* `next_desired_block_height`

&nbsp;
## `AddNewTransaction` (method) <!-- tal can finish -->

> Add a new transaction from a client to the network (propagate to all pending transaction pools)

#### Check transaction validity
* Correct protocol version.
* Valid fields (sender address, contract address).
* Sender virtual chain matches contract virtual chain and matches the pool's virtual chain.
* Check transaction `time_stamp`, accept only transactions with configurable time window (relative to current time).
* Transaction doesn't already exist in the pending pool or committed pool (duplicate).
* Verify pre order checks (like signature and subscription) by calling `VirtualMachine.TransactionSetPreOrder`.

#### Add transaction to pending pool
* Add transaction to pending transaction pool if pool is not full.
* Associate the transaction in the pool with the local node id (public key).

#### Batch before broadcast
* Batch until a configurable number of transactions is reached or until a configurable time passes.
* Sign the batch with the local node private key.
* Broadcast the batch to all other transaction pools by calling `Gossip.SendMessage`.

&nbsp;
## `OnForwardedTransactions` (method) <!-- tal can finish -->

> Add a transaction batch forwarded from another node to the pending transaction pool.

#### Check batch validity
* Check the batch signature, discard on error.

#### Add transaction to pending pool
* Add transaction to pending transaction pool if pool is not full.
* Associate the transaction in the pool with the batch node id (public key).

&nbsp;
## `GetTransactionsForOrdering` (method) <!-- tal can finish -->

> Returns a set of N trasnactions for block building based on the block building policy (first come first served)

#### Check synchronization status
* If requested block height is in the future but `last_committed_block` is close to it (configurable distance) block and wait
* If requested block height is in the future but `last_committed_block` is far, fail
* If requested block height is in the past, panic

#### Create proposal
* Prepare a transactions list proposal (policy is first come first serve).

#### Check transactions validity
* For each transaction:
  * Correct protocol version.
  * Valid fields (sender address, contract address).
  * Sender virtual chain matches contract virtual chain and matches the pool's virtual chain.
  * Check transaction `time_stamp`, accept only transactions with configurable time window (relative to current time).
  * Transaction doesn't already exist in the pending pool or committed pool (duplicate).
* Verify pre order checks (like signature and subscription) for all transactions by calling `VirtualMachine.TransactionSetPreOrder`.

&nbsp;
## `ValidateTransactionsForOrdering` (method) <!-- tal can finish -->

> Verifies that an ordered list of transactions complies with the ordering rules, called by the consensus-core when validating a new block proposal:

#### Check synchronization status
* If requested block height is in the future but `last_committed_block` is close to it (configurable distance) block and wait
* If requested block height is in the future but `last_committed_block` is far, fail
* If requested block height is in the past, panic

#### Check transactions validity
* For each transaction:
  * Correct protocol version.
  * Valid fields (sender address, contract address).
  * Sender virtual chain matches contract virtual chain and matches the pool's virtual chain.
  * Check transaction `time_stamp`, accept only transactions with configurable time window (relative to current time).
  * Transaction doesn't already exist in the pending pool or committed pool (duplicate).
* Verify pre order checks (like signature and subscription) for all transactions by calling `VirtualMachine.TransactionSetPreOrder`.
* TODO handle non-full blocks and pending transactions check (check "fairness").

&nbsp;
## `CommitTransactionsReceipts` (method) <!-- tal can finish -->

> Mark committed transactions as committed, removes them from the pending pool and notify the PublicAPI on local transactions that were committed.

#### Block Height
* If given block_height != `next_desired_block_height`:
  * Discard the commit
  * If `next_desired_block_height` is smaller than the first reqruied block height, set it to the first reqruied block height.
    * The first reqruied block height is derived from and the Block Storage's last_committed_block_height and de-duplication configurable time window.
  * Return the `next_desired_block_height`.

#### Commit Receipts
* For each transaction receipt:
  * Add the receipt, block height and block timestamp to the committed pool.
  * Lookup the the corresponding transaction in the pending pool, if local (originated by the node's Public API), update the local PublicAPI using `PublicAPI.ReturnTransactionResults`.
  * Remove the corresponding transactions (based on their tx_id) from the pending pool.
* Update the `last_committed_block_height`.
* Increment `next_desired_block_height`.
* If `next_desired_block_height` is smaller than the first reqruied block height, set it to the first reqruied block height.
  * The first reqruied block height is derived from and the Block Storage's last_committed_block_height and de-duplication configurable time window.
* Return the `next_desired_block_height`.


&nbsp;
## `GetTransactionReceipt` (method) <!-- tal can finish -->

> Return the status of a transaction

* If tx_id is present in the pending transaction pool, return status = PENDING
* If tx_id is present in the committed transaction pool, return status = COMMITTED and the receipt
* else retun status = NO_RECORD_FOUND.

&nbsp;
## `GossipMessageReceived` (method)

> Handle a gossip message from another node.

#### `ForwardedTransactions` message <!-- pass 1 -->
* Call `OnForwardedTransactions`.
