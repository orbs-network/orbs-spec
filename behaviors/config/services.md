# Per-Service Configuration

Configuration is an in-memory object given to every service on initialization. The object can be created based on static configuration files but will ultimately transition to being controlled by the management virtual chain via smart contracts.

#### All Services

```json
{
  "service-id": "tp1",
  "service-type": "TransactionPool",
}
```

#### Transaction Pool

```json
{
  "pending-pool-max-size-mb": "20480"
}
