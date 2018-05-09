# Public Api
Provides external public gateway interface (like JSON over HTTP) to the network (normally used by clients).
Maintains a state for clients sessions and send a response to the client.

&nbsp;
## `Public Interfaces`

#### JSON over HTTP
* All requests are HTTPPOST with JSON body and JSON response.
* Binary fields such as addresses are encoded with Base58.
* See request and response [encoding](../../interfaces/protocol/public-api/json-over-http.md).

&nbsp;
## `SendTransaction` (method)
> Execute a transaction that changes state (write) and returns a response to the client

#### Check request validity
* Correct protocol version.
* Correct virtual chain.

#### Subscribe to message response
* Log the client session along with the tx_id = SHA(Transaction serialization)

#### Forward call
* Send transaction to the network by calling `Consensus.SendTransaction`.

## `UpdateTransactionsResponse` (method)
* Called by the Consensus service upon commiting a new block, provides a list of receipts, each identified by a tx_id.
* Upon reception, match to the open sessions and send response back to the client.


&nbsp;
## `CallContract` (method)
> Run a read only contract that returns data

#### Check request validity
* Check protocol version
* Check signature
* todo

#### Forward call
* Forward call to `VirtualMachine.CallContract`.

&nbsp;
## `GetTransactionStatus` (method)
> Check the status of previously sent transaction

#### Check request validity
* todo.

#### Process
* If the transaction has an open session, return PENDING
* Else call `BlockStorage.GetTransactionReceipt`. // throttling mechanism.

&nbsp;
## `UpdateSubscriptionStatus` (method)
Called by the configuaration manager per virtual chain, updates the virtual chain's subscription level.

