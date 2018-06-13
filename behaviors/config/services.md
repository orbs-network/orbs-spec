# Per-Service Configuration

Configuration is an in-memory object given to every service on initialization. The object can be created based on static configuration files but will ultimately transition to being controlled by the management virtual chain via smart contracts.

#### Common fields (relevant to all services)

```json
{
  "service-id": "tp1",
  "service-type": "TransactionPool",
}
```

#### Public Api

```json
{
  "send-transaction-timeout-sec": "300",
  "get-transaction-status-grace-sec": "5"
}
```


#### Transaction Pool

```json
{
  "pending-pool-max-size-mb": "20480",
  "pending-pool-clear-expired-interval-sec": "10",
  "committed-pool-clear-expired-interval-sec": "30",
  "propagation-batching-transaction-num": "1000",
  "propagation-batching-timeout-sec": "5",
  "query-sync-grace-block-num": "0"
}
```

#### State Storage

```json
{
  "recent-block-snapshot-num": "5",
  "query-sync-grace-block-num": "0"
}
```

#### Block Storage

```json
{
  "block-sync-commit-timeout-sec": "8",
  "transaction-search-grace-sec": "5",
  "query-sync-grace-block-num": "0",
  "max-blocks-per-sync-batch": "10000"
}
```
