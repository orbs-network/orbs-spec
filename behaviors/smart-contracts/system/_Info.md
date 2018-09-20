# `_Info` System Service

System service providing information to the users of a virtual chain.

Implemented on the `Native` processor.

#### Permissions
* This is a system contract that runs under `System` permissions.

&nbsp;
## `_init` (method)

#### Permissions
* `Internal` (caller must be the same service).
* `ReadWrite` (does not change state).

#### Behavior
* Empty.

&nbsp;
## `getSignerAddress` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* Useful for end users wanting to verify their own address.
* Calls `Address.GetSignerAddress` and returns the result.
