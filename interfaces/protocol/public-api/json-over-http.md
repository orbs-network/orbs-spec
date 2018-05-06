# Public Api

## JSON over HTTP Encoding

&nbsp;
### `/public/sendTransaction`

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
### `/public/callContract`

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
* Calls `PublicApi.CallContract`.

&nbsp;
### `/public/getTransactionStatus`

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
