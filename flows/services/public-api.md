# Public Api

Provides external public gateway interface (like JSON over HTTP) to the network (normally used by clients).

&nbsp;
## `Public Interfaces`

#### JSON over HTTP
* All requests are HTTP POST with JSON body and JSON response.
* Binary fields such as addresses are encoded with Base58.
* See request and response [encoding](../../interfaces/protocol/public-api.md).

&nbsp;
## `SendTransaction` (rpc)

#### Check request validity
* Correct protocol version.

#### Forward call
* Forward call to `TransactionPool.AddNewPendingTransaction`.

&nbsp;
## `CallContract` (rpc)

#### Check request validity
* todo.

#### Forward call
* Forward call to `VirtualMachine.CallContract`.

&nbsp;
## `GetTransactionStatus` (rpc)

#### Check request validity
* todo.

#### Forward call
* Forward call to `TransactionPool.GetTransactionStatus`.
