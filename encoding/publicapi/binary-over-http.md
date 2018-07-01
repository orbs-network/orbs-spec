# Public Api Binary over HTTP Encoding

* Requests are HTTP POST with raw binary data of the serialized types.

* HTTP content type is `application/octet-stream`.

&nbsp;
## `/v1/call-method`

* Calls `PublicApi.CallMethod` with input and output encoded as the method's input and output.

&nbsp;
## `/v1/send-transaction`

* Calls `PublicApi.SendTransaction` with input and output encoded as the method's input and output.

&nbsp;
### `/v1/get-transaction-status`

* Calls `PublicApi.GetTransactionStatus` with input and output encoded as the method's input and output.