# Client Binary over HTTP Encoding

* Requests are HTTP POST with raw binary data of the serialized types.
* HTTP content type is `application/membuffers`.
* Errors that are non-application level return: 
  * The relevant HTTP status code (eg. `400 Bad Request`, `403 Forbidden`, `404 Not Found`, `429 Too Many Requests`, `500 Internal Server Error`, `503 Service Unavailable`).
  * HTTP content type of `text/plain` with a human readable explanation of the error as body. 
  * Example: If node fails on out of sync, returns `503 Service Unavailable` with text "node out of sync".
  
&nbsp;
## `/v1/api/call-method`

* Calls `PublicApi.CallMethod`.
* Request and response encoded as [MemBuffers](../serialization-format.md) serialized [messages](../../interfaces/protocol/client/requests.proto).

&nbsp;
## `/v1/api/send-transaction`

* Calls `PublicApi.SendTransaction`.
* Request and response encoded as [MemBuffers](../serialization-format.md) serialized [messages](../../interfaces/protocol/client/requests.proto).

&nbsp;
### `/v1/api/get-transaction-status`

* Calls `PublicApi.GetTransactionStatus`.
* Request and response encoded as [MemBuffers](../serialization-format.md) serialized [messages](../../interfaces/protocol/client/requests.proto).