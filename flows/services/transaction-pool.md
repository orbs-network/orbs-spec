# AddNewPendingTransaction

## Check transaction validity
* Correct protocol version.
* Valid fields (sender address, contract address).
* Sender virtual chain matches contract virtual chain.
* Valid transaction signature.

## Approve transaction for processing
* Not expired. Transaction is considered expired if `Transaction.timestamp` > current timestamp + expiration window.
* Subscription status is active. Get the virtual chain subscription status by calling `SubscriptionManager.GetSubscriptionStatus`.
* Transaction doesn't already exist in the pending pool (duplicate).

## Add transaction to pending pool
* Add transaction to pending transaction cache.
* Call `Gossip.BroadcastMessage` to broadcast transaction to all nodes.

## On any problem
* Return error and log error reason.