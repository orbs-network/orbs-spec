# Public Api

Provides external public gateway interface (like JSON over HTTP) to the network which is normally used by clients. Since responses are synchronous, maintains a list of all active client sessions.

Currently a single instance per virtual chain per node.

#### Interacts with services

* `BlockStorage` - Gets the latest committed block height from it and queries old transactions from it.
* `VirtualMachine` - Uses it to execute methods.
* `TransactionPool` - Provides it with new transactions.

&nbsp;
## `Public Interfaces`

#### Binary HTTP POST
* Requests are HTTP POST with binary data of the serialized types.
* See request and response [encoding](../../encoding/client/binary-over-http.md).

#### JSON over HTTP
* Requests are HTTP POST with textual JSON body and JSON response.
* See request and response [encoding](../../encoding/client/json-over-http.md).

#### HTTP GET
* Requests are HTTP GET with all parameters encoded on the URL.
* See request and response [encoding](../../encoding/client/http-get.md).
* Limited to `CallMethod` requests.

&nbsp;
## `Init` (flow)

* Initialize the [configuration](../config/services.md).
* Register to handle transactions results blocks by calling `TransactionPool.TransactionResultsHandler`.

&nbsp;
## `CallMethod` (method)

> Public interface: Run a read only service method without consensus based on the current block height. The call is synchronous.

#### Check request validity
* Correct protocol version.
* Correct virtual chain.
* Notes: 
  * The request format is validated by the HTTP server.
  * Upon a validity error, retrun an error status with empty block height and timestamp (as they may not be relevant).

#### Forward call
* Execute call on the virtual machine by calling `VirtualMachine.RunLocalMethod` indicating recent [block height](../../terminology.md).
* Return the result along with the reference block height and timestamp.

&nbsp;
## `SendTransaction` (method)

> Public interface: Execute a transaction on a service under consensus that may change state (write). The call is synchronous.

#### Check request validity
* Correct protocol version.
* Correct virtual chain.
* Notes: 
  * The request format is validated by the HTTP server.
  * Upon a validity error, retrun an error status with empty block height and timestamp (as they may not be relevant).
  
#### Forward transaction
* Calculate the transaction `tx_id` (see transaction format for structure).
* Send transaction to the network by calling `TransactionPool.AddNewTransaction`.
  * On failure, send response to client along with the reference block height and timestamp.
* Block until `HandleTransactionResults` or `HandleTransactionError` are called with the relevant `tx_id`.
* If a [configurable](../config/services.md) timeout expires during the block, fail.
  * Note: Beware of having the forwarded transaction fail somewhere else and swallowed without calling `HandleTransactionResults`, `HandleTransactionError`.

&nbsp;
## `GetTransactionStatus` (method)

> Public interface: Query the status of previously sent transaction.

#### Check request validity
* Correct protocol version.
* Correct virtual chain.
* Notes: 
  * The request format is validated by the HTTP server.
  * Upon a validity error, retrun an error status with empty block height and timestamp (as they may not be relevant).

#### Query transaction status
* Query the transactions pool by calling `TransactionPool.GetCommittedTransactionReceipt`.
  * If the return status is `PENDING` or `TIMESTAMP_AHEAD_OF_NODE_TIME`, return with the corresponding block_hieght and timestamp and an empty receipt.
  * If the return status is `COMMITTED`, return with the receipt and corresponding block_hieght and timestamp. 
* If not found in transaction pool (`NO_RECORD_FOUND`), it might be an older transaction, widen our search.
* Query the block storage by calling `BlockStorage.GetTransactionReceipt`.
  * If found return status `COMMITTED` with the receipt, else return status `NO_RECORD_FOUND` along with the reference block height and timestamp.

&nbsp;
## `GetTransactionReceiptProof` (method)
<!-- TODO: consider providing the receipt as an input -->

> Public interface: Returns a receipt along with a proof for its inclusion in a block.

#### Check request validity
* Correct protocol version.
* Correct virtual chain.
* Notes: 
  * The request format is validated by the HTTP server.
  * Upon a validity error, return an error status with empty block height and timestamp (as they may not be relevant).

#### Query the transaction status and receipt
* Query the transaction status by calling `GetTransactionStatus`.
  * If no receipt was found return `NO_RECORD_FOUND` along with the reference block height and timestamp that were returned by `GetTransactionStatus`. 

#### Get a receipt proof
* Get a receipt proof by calling `BlockStorage.GenerateReceiptProof`.
* Return status `COMMITTED` along with the provided proof, block_height and timestamp.

&nbsp;
## TransactionResults Handler

> Handles transaction results enables the public api to respond to the waiting clients, called by `TransactionPool`. 

#### `HandleTransactionResults`
> Returns the results of committed transaction set.
* For every transaction:
  * Locate the relevant blocking `SendTransaction` contexts based on the `tx_id`.
  * Unblock them to respond to the client using the data from the transaction receipt.
    * Set response.transaction_status = COMMITTTED. 

&nbsp;
#### `HandleTransactionError`
> Returns the result of a rejected transaction
* Locate the relevant blocking `SendTransaction` contexts based on the `tx_id`.
  * Unblock them to respond to the client using the data from the response.
    * Set response.transaction_receipt = NULL.
