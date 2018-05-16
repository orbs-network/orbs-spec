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

#### Forward call
* Send transaction to the network by calling `TransactionPool.AddNewTransaction`.

#### Maintain sessions context
* Maintain session context in order to return a response and associate a tx_id = SHA256(Transaction).
* TODO: seession timeout (should be ~ 1min)

## `UpdateTransactionsResponse` (method)
* Called by the Consensus service upon commiting a new block, provides a list of receipts, each identified by a tx_id.
* Upon reception, match to the open sessions and send response back to the client.


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
