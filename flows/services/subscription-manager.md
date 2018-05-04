# Subscription Manager

Maintains status of active subscriptions for virtual chains.

&nbsp;
## `GetSubscriptionStatus` RPC

* Get subscription details from the Ethereum subscription contract with `SidechainConnector.CallEthereumContract`.
* A subscription is active if the amount of deposited tokens is greater than 0.
