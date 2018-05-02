# Transaction Pool

Holds pending transactions and remembers committed transactions.

&nbsp;
## `Data Structures`

#### Pending transaction pool
* Holds transactions until they are added to a block and helps preventing transaction duplication.
* No need to be persistent.
* Configurable max size.
* Configurable interval to clear expired transactions. Transaction is considered expired if transaction timestamp > current timestamp + expiration window.

#### Committed transaction pool
* Holds transactions after they are added to the blockchain to avoid transaction duplication and return their result.
* No need to be persistent.
* No limit on max size.
* Configurable interval to clear expired transactions. Transaction is considered expired if transaction timestamp > current timestamp + expiration window.

&nbsp;
## `AddNewPendingTransaction`

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
* Call `Gossip.BroadcastMessage` to broadcast transaction to all nodes as "newTransaction" message.

&nbsp;
## `GetAllPendingTransactions`

* Return all transaction ids and transactions from pending transaction pool.

&nbsp;
## `MarkCommittedTransactions`

* Add all receipts to committed transaction pool.
* Delete all the receipts' transactions from the pending transaction pool.

&nbsp;
## `GetTransactionStatus`

* If the transaction is in the committed transaction pool return COMMITTED and its receipt.
* If the transaction is in the pending transaction pool return PENDING.

&nbsp;
## `GossipMessageReceived`

### On "newTransaction" message
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
