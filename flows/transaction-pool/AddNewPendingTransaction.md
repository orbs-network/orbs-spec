# AddNewPendingTransaction

* if transaction is expired return error
* if transaction is not valid return error
  * correct protocol version
  * valid fields (sender address, contract address)
  * sender virtual chain matches contract virtual chain

* call `SubscriptionManager.GetSubscriptionStatus` to get the virtual chain subscription status
* if subscription status is not active return error

* if transaction already exists in pending transaction cache return error
* add transaction to pending transaction cache

* call `Gossip.BroadcastMessage` to broadcast transaction to all nodes
