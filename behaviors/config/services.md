# Per-Service Configuration

Configuration is an in-memory object given to every service on initialization. The object can be created based on static configuration files but will ultimately be controlled by the management virtual chain via smart contracts.

Since the system must remain backwards compatible forever (able to audit the old blocks in the chain), changes in configuration can only take place from a certain block height. Until that block height is reached, the original configuration must remain in the system. One way to implement this is place the entire configuration dictionary as a key of a parent map containing block ranges.

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

#### Gossip

```json
{
  "reconnection-polling-interval-sec": "10"
}
```

#### Consensus Algo (per algorithm)

```json
{
  "committee-size": "21",
  "block-content-max-size-kb": "1024"
}
```

#### Consensus Context

```json
{
  "minimal-block-transaction-num": "0",
  "minimal-block-delay-sec": "0.5",
  "system-timestamp-allowed-jitter-sec": "2"
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
