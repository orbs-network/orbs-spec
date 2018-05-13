# Transaction Pool

Holds pending transactions and remembers committed transactions.

&nbsp;
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

&nbsp;
## `AddNewTransaction` (method)
> Add a new transaction from a client to the network (pending transaction pool)

#### Check transaction validity
* Correct protocol version.
* Valid fields (sender address, contract address).
* Sender virtual chain matches contract virtual chain.
* Check transaction time_stamp, accept only transactions with last block.timestamp -2sec  < time_stamp < block.timestamp expiration window. 
* Valid transaction signature.

If a transaction fails the validation, update the Public API by calling `PublicAPI.UpdateTransactionsResponse`.

#### Approve transaction for processing
* Not expired. Transaction is considered expired if transaction timestamp > current timestamp + expiration window.
* Check that the Subscription status is active. The virtual chain subscription status is updated periodically by the through the Consensus service, calling the `UpdateSubscriptionStatus` method.
* Transaction doesn't already exist in the pending pool or committed pool (duplicated).

If a transaction fails the approval, update the Public API by calling `PublicAPI.UpdateTransactionsResponse`.

#### Add transaction to pending pool
* Add transaction to pending transaction pool if pool is not full. //TODO full considerations
* Add transcation to the SendTransactionCache.

### SendTransactionCache
Temporarly delays transactions transmission in order to batch them. (optimization)
If not empty and X = 100 ms have passed since the last batch was sent or cache has more than Y = 100 messages, create batch with up to Y transactions, sign and boradcast the batch to all the other trasnaction pools by calling `Gossip.BroadcastMessage`.

&nbsp;
## `GetAllPendingTransactions` (method)

### Check that the pools are up to date
Check that the requested block_height equal to the block_height the pools are updated to, otherwise initiate Node Sync and return NOT UPDATED.
> Return all currently pending transactions
* Return all transaction ids and transactions from pending transaction pool.

&nbsp;
## `MarkCommittedTransactions` (method)
> Take a group of transactions that were committed and mark them as so

* Add all receipts to committed transaction pool.
* Delete all the receipts' transactions from the pending transaction pool.

&nbsp;
## `GetTransactionStatus` (method)
> Return the status of a transaction

* If the transaction is in the committed transaction pool return COMMITTED and its receipt.
* If the transaction is in the pending transaction pool return PENDING.

&nbsp;
## `AddForwardedTransaction` (method)
> Handle a transaction forwarded by gossip by another node

* Make sure transaction isn't in the committed or pending transaction pools.

#### Check transaction validity
* Correct protocol version.
* Valid fields (sender address, contract address).
* Sender virtual chain matches contract virtual chain.
* Valid transaction signature.

#### Approve transaction for processing
* Not expired. Transaction is considered expired if transaction timestamp > current timestamp + expiration window.
* Subscription status is active. Get the virtual chain subscription status by calling `SubscriptionManager.GetSubscriptionStatus`.
* Transaction doesn't already exist in the pending pool (duplicate).

#### Add transaction to pending pool
* Add transaction to pending transaction pool if pool is not full.

&nbsp;
## `GossipMessageReceived` (method)
> Handle a gossip message from another node

#### On "newTransaction" message
* Call `AddForwardedTransaction`.
