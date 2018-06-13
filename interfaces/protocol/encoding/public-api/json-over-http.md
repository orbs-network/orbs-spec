# Public Api JSON over HTTP Encoding

&nbsp;
## `/v1/call-method`

#### Request
```json
{
  "contractAddressBase58": "todo",
  "senderAddressBase58": "todo",
  "payload": "todo"
}
```

#### Response
```json
{

}
```

#### Implementation
* Calls `PublicApi.CallMethod`.

&nbsp;
## `/v1/send-transaction`

#### Request
```json
{
  "header": {
    "version": "todo",
    "senderAddressBase58": "todo",
    "timestamp": "todo",
    "contractAddressBase58": "todo"
  },
  "payload": "todo"
}
```

#### Response
```json
{
  "result": "ok"
}
```

#### Implementation
* Calls `PublicApi.SendTransaction`.

&nbsp;
### `/v1/get-transaction-status`

#### Request
```json
{
  "txid": "todo"
}
```

#### Response
```json
{
  "status": "COMMITTED",
  "receipt": {
    "success": "todo"
  }
}
```

#### Implementation
* Calls `PublicApi.GetTransactionStatus`.
