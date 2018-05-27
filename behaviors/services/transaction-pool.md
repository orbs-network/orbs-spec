# Transaction Pool

Holds pending transactions until added to a block, and keeps committed transactions for a time window to avoid duplication. Proposes transactions for new blocks and validates proposals by other nodes.

Currently a single instance per virtual chain per node.

![alt text][transaction_preorder_validation_flow] <br/><br/>

[transaction_preorder_validation_flow]: transaction_preorder_validation_flow.png "PreOrder Validation Flow"

&nbsp;
## Init Flow
* Subscribe to transactions messages.

&nbsp;
## `Data Structures`

#### Pending transaction pool
* Holds transactions until they are added to a block and helps preventing transaction duplication.
* Associates every transaction with the node id (public key) of the gateway that added it to the network.
* No need to be persistent.
* Configurable max size.
* Configurable interval to clear expired transactions by comparing `time_stamp` to a configurable time window (relative to current time).

#### Committed transaction pool
* Holds the ID of transactionss after they are added to the blockchain to avoid transaction duplication and return their result.
* No need to be persistent, re-sync from the block storage.
* No limit on max size, max_size is determined by the exeperation window x maximum transaction rate. (on full committed pool - fatal error, drop all new committed blocks until not full)
* Configurable interval to clear expired transactions. Transaction is considered expired if transaction timestamp > last block.timestamp + expiration window.

&nbsp;
## `AddNewTransaction` (method) <!-- pass 1 -->

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
## `AddForwardedTransactionBatch` (method) <!-- pass 1 -->

> Add a transaction batch forwarded from another node to the pending transaction pool

#### Check batch validity
* Check the batch signature, discard on error.

#### Add transaction to pending pool
* Add transaction to pending transaction pool if pool is not full.
* Associate the transaction in the pool with the batch node id (public key).






&nbsp;
## `GetPendingTransactions` (method)
> Returns a set of N trasnactions for block building based on the block building policy (first come first served)

#### Make sure the transaction pool is in sync
* Update last_requested_block = block height.
* Check block height = last_commited_block height + 1.  
  * If equal reset OUT_OF_SYNC, stop the `Sync flow`.
* If block height = last_commited_block height + 1
  * hold response for up to X = 2 sec waiting for potential receipts commit from the block storage.
  * Upon timeout set OUT_OF_SYNC, return OUT_OF_SYNC response and initiate a `Sync Flow` starting with last_commited_block height + 1.
* If block height > last_commited_block height + 1:
  * set OUT_OF_SYNC, return OUT_OF_SYNC response and initiate a `Sync Flow` starting with last_commited_block height + 1.
* If block height < last_commited_block + 1 return UNEXEPCETED_BLOCK_HEIGHT.

### Returns a list of trasnactions
* Prepare a transactions list proposal, verify that the trasnactions are valid for ordering by calling `CheckPreOrder`.
* Return an ordered list of transactions from the pending transaction pool.
  (Notice that the transactions are not evicted from the pool until `MarkCommittedTransactions`)

&nbsp;
## `ValidateTransactionsForOrdering` (method)
> Verifies that an ordered list of transactions complies with the ordering rules, called by the consensus-core when validating a new block proposal:

### Check that the pools are up to date
* Check block height = last_commited_block height + 1.
* If block height > last_commited_block + 1 initate sync flow and return OUT_OF_SYNC status.
* If block height < last_commited_block + 1 return UNEXEPCETED_BLOCK_HEIGHT.

### Check transaction list
* Verify that the trasnactions are valid for ordering by calling `CheckPreOrder`.
* TODO handle non-full blocks and pending transactions check.

### Return status
* Return a valid / invalid status.

&nbsp;
## `MarkCommittedTransactions` (method)
> Mark committed transactions as committed, removes them from the pending pool and notify the PublicAPI on local transactions that were committed.
* For each transaction receipt:
    * Add the tx_id to the committed pool
    * Lookup the the corresponding transaction in the pending pool, if local (originated by the node's Public API), update the local PublicAPI using `PublicAPI.UpdateTransactionsResponse`.
    * Remove the corresponding transactions (based on their tx_id) from the pending pool.

&nbsp;
## `GetTransactionStatus` (method)
> Return the status of a transaction

#### Check pending transactions pool
* If tx_id is present in the pending transaction pool, return status = PENDING else retun status = NO_RECORD_FOUND.

&nbsp;
## `GossipMessageReceived` (method)

> Handle a gossip message from another node.

#### `TransactionsBatch` message <!-- pass 1 -->
* Call `AddForwardedTransactionBatch`.

&nbsp;
## `UpdateSubscriptionStatus` (method)
> Updates the transaction pool Subscription status.
* Update the local subscription status for the virtual chain, takes effect starting from the indicated block.

&nbsp;
## `Sync Flow`
> Syncs the state storage based on the block storage state diff.
* Call `BlockStorage.RequestReceiptsUpdate` with consumer_block_height = last_commited_block + 1, target_block_height = MAX_UINT64.
