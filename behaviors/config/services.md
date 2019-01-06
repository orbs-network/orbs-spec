# Per-Service Configuration

Configuration is an in-memory object given to every service on initialization. The data can be created based on static configuration files but will ultimately be controlled by the management virtual chain via smart contracts.

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
}
```

#### Transaction Pool

```json
{
  "pending-pool-max-size-bytes": "20480000",
  "pending-pool-clear-expired-interval-sec": "10",
  "committed-pool-clear-expired-interval-sec": "30",
  "propagation-batch-size": "100",
  "propagation-batching-timeout-msec": "100",
  "query-sync-grace-block-num": "0",
  "future-timestamp-grace-sec": "5"
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
  "block-sync-commit-timeout-empty-blocks-time": "4",
  "transaction-search-grace-sec": "5",
  "query-sync-grace-block-num": "0",
  "max-blocks-per-sync-batch": "1000",
  "max-blocks-per-sync-chunk": "10"
}
```

#### Cross-chain Connector
<!--- TODO: move to per node configuration file --->
```json
{
  "etheruem-node-dns-address": "ethereum.orbs.network",
  "etheruem-node-port": "8545",
  "number-of-blocks-for-finality": "20"
}
```
