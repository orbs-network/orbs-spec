# Subscription Manager

Maintains status of active subscriptions for virtual chains.

&nbsp;
## `GetSubscriptionStatus` (method)
> Check the status of a virtual chain subscription

* Get subscription details from the Ethereum subscription contract with `SidechainConnector.CallEthereumContract`.
* A subscription is active if the amount of deposited tokens is greater than 0.
