# `_GlobalPreOrder` System Service
The system service charged with approving every transaction on a virtual chain (eg. checking subscription).

Implemented on the `Native` processor.

#### Permissions
* This is a system contract that runs under `System` permissions.

#### State
* Constants: 
     * ELECTED_VALIDATORS_SYNC_GRACE_TIME - allowed time drift of config - continue running with old information (~ 10 minutes) used in preOrder checks.

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
    * Get the elected validators set expiration (`nextSync`), by service call `CallMethod("_ElectedValidators", "getElectedValidatorsNextSync"`)`.
    * Reject if block.timestamp > `nextSync` + `ELECTED_VALIDATORS_SYNC_GRACE_TIME`.

* Get the virtual chain by calling `Environment.GetVirtualChain`.
<!-- * Read the subscription data from the `_Provision` system contract, by service call `CallMethod("_Provision", "getSubscriptionInfo()"`)`. -->
* Read the subscription data from Ethereum by calling `Ethereum.CallContract`.
  * Optimization: Maintain a cached value of the subscription status, refresh it every 1000 blocks.
* Check if the subscription is valid.
