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

#### Contract
* Write state value `Native` in key `_Deployments.Processor` by calling `State.WriteStringKey`.

&nbsp;
## `deployService` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite` (potentially changes state).

#### Contract
* Make sure the service isn't already deployed by calling `Service.IsDeployed`.
* Init the service by calling `Service.CallMethod` with method `_init`.
  * This is an `Internal` method, requires `System` permissions.
* Write state value `Native` in key `<name>.Processor` by calling `State.WriteStringKey`.
