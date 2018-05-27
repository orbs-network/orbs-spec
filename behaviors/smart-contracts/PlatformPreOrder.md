# Platform Pre-Order Smart Contract
> The susbcription check is implemented as a hard coded smart contract (meta-programing).

## Interface
Inputs: virtual chain
Outputs: APPROVED / NOT_APPROVED, status_message
Permissions: Read ONLY, Exeternal
> Returns the subscription status of a specific virtual chain. 

* Deployed on the local virtual chain.
* Reads the subscription data from the Ethereum subscription smart contract (getSubscriptionData) using `SideChainconnector.CallEthereumContract`.
    * Maintains a cached value of the subscription status, refresh it every 1000 blocks.
* A susbcription is considered valid if - TODO.
* If the subscription is valid return APPROVED, status_message = ""
* If the subscription is valid return NOT_APPROVED, status_message = "Subscription Error".
