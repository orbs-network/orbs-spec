# `_GlobalPreOrder` System Service
The system service charged with approving every transaction on a virtual chain (eg. checking subscription).

Implemented on the `Native` processor.

#### Permissions
* This is a system contract that runs under `System` permissions.

&nbsp;
## `_init` (method)

#### Permissions
* `Internal` (caller must be the same service).
* `ReadWrite` (potentially changes state).

#### Behavior
* Empty.

&nbsp;
## `approve` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* Check if the elected validators set is outdated:
    * Get the elected validators set expiration (`ValidUntil`), by calling `_ElectedValidators.getElectedValidators()`.
    * Reject if block.timestamp > `ValidUntil`.
    
* Get the virtual chain by calling `Environment.GetVirtualChain`.
* Read the subscription data from Ethereum by calling `Ethereum.CallContract`.
  * Optimization: Maintain a cached value of the subscription status, refresh it every 1000 blocks.
* Check if the subscription is valid.
