# `_Deployments` System Service

The system service charged with deploying new smart contracts (services or libraries) onto a virtual chain.

Implemented on the `Native` processor.

#### Permissions
* This is a system contract that runs under `System` permissions.

&nbsp;
## `_init` (method)

#### Permissions
* `Internal` (caller must be the same service).
* `ReadWrite` (potentially changes state).

#### Behavior
* Must be empty to avoid chicken and egg problem since used to deploy all other contracts.

&nbsp;
## `getInfo` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly` (potentially changes state).

#### Behavior
* Takes service name as argument.
* Read processor type from state key `<name>.Processor` by calling `State.ReadUint32ByKey`.
* If an empty value is read, contract is not deployed.

&nbsp;
## `deployService` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite` (potentially changes state).

#### Behavior
* Takes service name, processor type and code as arguments.
* Make sure the service isn't already deployed by calling `getInfo`.
* Write processor type to state under key `<name>.Processor` by calling `State.WriteUint32ByKey`.
* Write code to state under key `<name>.Code` by calling `State.WriteBytesByKey`.
* Init the service by calling `Service.CallMethod` with method `_init`.
  * This is an `Internal` method, requires `System` permissions.

&nbsp;
## `getCode` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly` (potentially changes state).

#### Behavior
* Takes service name as argument.
* Read code from state key `<name>.Code` by calling `State.ReadBytesByKey`.
