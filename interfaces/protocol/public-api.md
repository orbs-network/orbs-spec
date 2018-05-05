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
