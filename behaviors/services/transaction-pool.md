# Transaction Pool

Performs pre-order validation, holds pending transactions, and committed transactions it, propose transactions for new block and validates the transactions order in a propsoed block.
&nbsp;

![alt text][ledger] <br/><br/>

[ledger]: transaction_preorder_validation_flow.png "PreOrder Validation Flow"

## `Data Structures`

#### Pending transaction pool
* Holds transactions until they are added to a block and helps preventing transaction duplication.
* No need to be persistent. Syncs from the block stroage upon loss of sync.
* Configurable max size.
* Configurable interval to clear expired transactions. Transaction is considered expired if transaction timestamp > last block.timestamp + expiration window.
* Sends the transaction receipt to the local relay_gw upon committment of a new block.

#### Committed transaction pool
* Holds the ID of transactionss after they are added to the blockchain to avoid transaction duplication and return their result.
* No need to be persistent, consistency with the commited block is achived as part of the node sync.
* No limit on max size, max_size is determined by the exeperation window x maximum transaction rate. (on full committed pool - fatal error, drop all new committed blocks until not full)
* Configurable interval to clear expired transactions. Transaction is considered expired if transaction timestamp > last block.timestamp + expiration window.

## `CheckPreOrder` (method) // TODO move to VM
> Checks a transaction's validity and approves it for ordering.
> Performed on transactions received both from a local Public API before signing and broadcasting them and as part of the ordering block construction and validation.

#### Check transaction validity
* Correct protocol version.
* Valid fields (sender address, contract address).
* Sender virtual chain matches contract virtual chain matches instance virtual chain.
* Check transaction time_stamp, accept only transactions with last block.timestamp -5 sec < time_stamp < block.timestamp expiration window. 
* Valid transaction signature.

#### Approve transaction for processing
* Not expired. Transaction is considered expired if transaction timestamp > current timestamp + expiration window.
* Check that the Subscription status is active.
    * The virtual chain subscription status is updated periodically by the ConfigurationManager calling the `UpdateSubscriptionStatus` method.
* Transaction doesn't already exist in the pending pool or committed pool (duplicated).


&nbsp;
## `AddNewTransaction` (method)
> Add a new transaction from a client to the network (pending transaction pool)
* If out of sync, hold execution until regain sync.
    * Note: AddNewTransaction is an a-sync event as such it's not consistent across nodes. However gossiping duplicate trunsactions can damage the node's reputation.
* Check that the trasnaction is valid and approved for ordering by calling `CheckPreOrder`.
* If a transaction fails the approval, update the Public API by calling `PublicAPI.UpdateTransactionsResponse` with:
    * tx_id 
    * status
    * current time_stamp
    (All other fields are don't care)

#### Add transaction to pending pool
* Add transaction to pending transaction pool if pool is not full. //TODO full considerations
* Mark the transaction as originated by the local PublicAPI (the node's public key) in order to respond upon committ indication.
* Add transcation to the SendTransactionCache.

### SendTransactionCache
Temporarly delays transactions transmission in order to batch them. (optimization)
If not empty and X = 100 ms have passed since the last batch was sent or cache has more than Y = 100 messages, create batch with up to Y transactions, sign and boradcast the batch to all the other trasnaction pools by calling `Gossip.BroadcastMessage`.

&nbsp;
## `GetPendingTransactions` (method)
> Returns a set of N trasnactions for block building based on the block building policy (first come first served)

### Check that the pools are up to date
* Check block height = last_commited_block height + 1. 
* If block height > last_commited_block + 1 initate sync flow and return OUT_OF_SYNC status.
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

* If the transaction is in the committed transaction pool return COMMITTED and its receipt.
* If the transaction is in the pending transaction pool return PENDING.

&nbsp;
## `GossipMessageReceived` (method)
> Handle a gossip message (TransactionsBatch) from another node.

&nbsp;
#### Check transactions
* Validate the batch signature, if error, silegntly discard.

#### Add transaction to pending pool
* Add transaction to pending transaction pool if pool is not full. //TODO full considerations
* Maintain for each transaction the associated node ID (public key) and the validation result to be used by the reputation mechanism.

&nbsp;
## `UpdateSubscriptionStatus` (method)
> Updates the transaction pool Subscription status.
* Update the local subscription status for the virtual chain, takes effect starting from the indicated block.
