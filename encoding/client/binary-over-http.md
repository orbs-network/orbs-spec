# Client Binary over HTTP Encoding

* Requests are HTTP POST with raw binary data of the serialized types.
* HTTP content type is `application/membuffers`.

* The following HTTP headers are encouraged to be returned for verbosity:
  * `X-ORBS-REQUEST-RESULT` - human readable description of the result.
  * `X-ORBS-BLOCK-HEIGHT` - the reference block height for the result as a decimal number.
  * `X-ORBS-BLOCK-TIMESTAMP` - human readable timestamp of the reference block of the result (RFC3339). 
  * `X-ORBS-ERROR-DETAILS` - human readable description of any error if relevant.

* Errors that are non-application level return: 
  * The relevant HTTP status code (eg. `400 Bad Request`, `403 Forbidden`, `404 Not Found`, `429 Too Many Requests`, `500 Internal Server Error`, `503 Service Unavailable`).
  * HTTP content type of `text/plain` with a human readable explanation of the error as body. 
  * Example: If node fails on out of sync, returns `503 Service Unavailable` with text "node out of sync".
  
&nbsp;
## `/api/v1/run-query`

* Calls `PublicApi.RunQuery`.
* Request and response encoded as [MemBuffers](../serialization-format.md) serialized [messages](../../interfaces/protocol/client/requests.proto).

&nbsp;
## `/api/v1/send-transaction`

* Calls `PublicApi.SendTransaction`.
* Request and response encoded as [MemBuffers](../serialization-format.md) serialized [messages](../../interfaces/protocol/client/requests.proto).
* This API is synchronous and can take a long time to process (until a block containing the transaction is closed).

&nbsp;
### `/api/v1/get-transaction-status`

* Calls `PublicApi.GetTransactionStatus`.
* Request and response encoded as [MemBuffers](../serialization-format.md) serialized [messages](../../interfaces/protocol/client/requests.proto).

&nbsp;
### `/api/v1/get-transaction-receipt-proof`

* Calls `PublicApi.GetTransactionReceiptProof`.
* Request and response encoded as [MemBuffers](../serialization-format.md) serialized [messages](../../interfaces/protocol/client/requests.proto).

&nbsp;
## `/api/v1/get-block`

* Calls `PublicApi.GetBlock`.
* Request and response encoded as [MemBuffers](../serialization-format.md) serialized [messages](../../interfaces/protocol/client/requests.proto).

&nbsp;
## `/api/v1/send-transaction-async`

* Similar to `/api/v1/send-transaction` but does not wait until the transaction is fully executed.
* Only waits until the transaction is stored successfully in the transaction pool.