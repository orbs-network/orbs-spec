# Public Api

Provides external public gateway interface (like JSON over HTTP) to the network (which is normally used by clients).
Since responses are synchronous, maintains a list of all active client sessions.

Currently a single instance per virtual chain per node.

&nbsp;
## `Public Interfaces`

#### JSON over HTTP
* Requests are HTTP POST with JSON body and JSON response.
* TODO: Binary fields such as addresses are encoded with Base58.
* See request and response [encoding](../../interfaces/protocol/encoding/public-api/json-over-http.md).

&nbsp;
## `SendTransaction` (method) <!-- pass 1 -->

> Execute a transaction under consensus that may change state (write)

#### Check request validity
* Correct protocol version.
* Correct virtual chain.

#### Forward transaction
* Calculate the transaction `tx_id` (see transaction format).
* Maintain session context in order to eventually return a response and associate it with the `tx_id`.
* Send transaction to the network by calling `TransactionPool.AddNewTransaction`.
* If session resets or timeouts, clear session context.

&nbsp;
## `ReturnTransactionResults` (method) <!-- pass 1 -->

> Called by TransactionPool on committed blocks to let PublicApi respond to their waiting clients

* Locate the relevant session contexts based on `tx_id` of every transaction.
* Respond to client using data from the transaction receipt.

&nbsp;
## `GetTransactionStatus` (method)

> Check the status of previously sent transaction

#### Check request validity
* Check the transaction timestamp, verifying that it's not a future transaction plus grace (eg. 5 seconds).

#### Query transaction status
* Queries the transactions pool by calling `TransactionPool.GetTransactionReceipt`.
  * If found returns PENDING, if committed return COMMITTED with the receipt.
* Queries the block storage by calling `BlockStorage.GetTransactionReceipt`.
  * If found returns COMMITTED with the receipt, else returns NO_RECORD_FOUND.

&nbsp;
## `CallMethod` (method)

> Run a read only method based on the current block height and returns the output arguments.

#### Retrains session info
  * Retains session info so the response will eventually arrive to this client.

#### Check request validity
* Performs checks on the transaction:
  * Version, Transaction format, virtual chain.
* Batch tarnsactions and queries the Consensus core on the latest block height by calling `ConsensusBuilder.GetLatestBlockHeight`.
  * Note that the method call is executed asynchronous to the block creation and the block height on which the execution is perfrormed may vary up to few blocks from the latest block.

#### Forward call
* Sends the trasnactions to the VM for execution by calling `VirtualMachine.RunLocalMethod` with the latest block height.
