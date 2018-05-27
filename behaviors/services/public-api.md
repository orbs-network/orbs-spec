# Public Api

Provides external public gateway interface (like JSON over HTTP) to the network (which is normally used by clients).
Since responses are synchronous, maintains a list of all active client sessions.

Currently a single instance per virtual chain per node.

&nbsp;
## `Public Interfaces`

#### JSON over HTTP
* Requests are HTTP POST with JSON body and JSON response.
* TODO: Binary fields such as addresses are encoded with Base58.
* See request and response [encoding](../../interfaces/protocol/public-api/json-over-http.md).

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
* Chcek time_stamp is within valid window: < last_block_time + 2sec, else return:
    * tx_id
    * status = INVALID_REQUEST
    * last block time_stamp
    (All other fields are don't care)

#### Check local sessions database
* If tx_id is present in the local sessions database, return:
    * tx_id
    * status = PENDING
    * last block time_stamp
    (All other fields are don't care)

#### Check transaction pool
* Call `TransactionPool.GetTransactionStatus`
    * If status = PENDING, return the response. //TBD hold committed status in transaction pool.

#### Check block storage
* Call `BlockStorage.GetTransactionReceipt`, return the response.

// TODO rate limiter.

&nbsp;
## `CallContract` (method)

> Run a read only contract that returns data

#### Check request validity
* Check protocol version
* Check signature
* todo

#### Forward call
* Forward call to `VirtualMachine.CallContract`.



#### Process
* If the transaction has an open session, return PENDING
* Else call `BlockStorage.GetTransactionReceipt`. // throttling mechanism.
