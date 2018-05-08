# Configuration Manager

Responsible for updating the relevant services on federation and subscription configuation changes.
* Maintains status of active subscriptions for virtual chains, updates the Public API and the Consensus on the subscription status. (push)
* Maintains status of the active federation members, updates the Consensus.

The configurations are updated periodically. Subscriptions and Federation status read from Ethereum's smart contracts are accessed using the `SidechainConnector`.

&nbsp;
## Priodic update 
New configueration takes effect every K = 1000 blocks (when processing the block with block_height mod K = 0)
### Subscriptions
* Get subscription details from the Ethereum subscription contract with `SidechainConnector.CallEthereumContract`.
* Update the relevant services using: `Public-API.UpdateSubscriptionStatus` and `Consensus.UpdateSubscriptionStatus`.
* A subscription is active if the amount of deposited tokens is greater than 0.

### Federation Members
* Get subscription details from the Ethereum subscription contract with `SidechainConnector.CallEthereumContract`.
* Update the relevant services using: `Consensus.UpdateFederationStatus`.
