# Public Api

Provides external public gateway interface (like JSON over HTTP) to the network (normally used by clients).

&nbsp;
## `Public Interfaces`

#### JSON over HTTP
* All requests are HTTP POST with JSON body and JSON response.
* Binary fields such as addresses are encoded with Base58.
* See request and response [encoding](../../interfaces/protocol/public-api/json-over-http.md).

&nbsp;
## `SendTransaction` (method)
> Execute a transaction that changes state (write)

#### Check request validity
* Correct protocol version.

#### Forward call
* Forward call to `TransactionPool.AddNewPendingTransaction`.

&nbsp;
## `CallContract` (method)
> Run a read only contract that returns data

#### Check request validity
* todo.

#### Forward call
* Forward call to `VirtualMachine.CallContract`.

&nbsp;
## `GetTransactionStatus` (method)
> Check the status of previously sent transaction

#### Check request validity
* todo.

#### Forward call
* Forward call to `TransactionPool.GetTransactionStatus`.
