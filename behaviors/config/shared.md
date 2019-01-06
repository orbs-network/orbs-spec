# Shared Configuration

Configuration is an in-memory object given to every service on initialization. The data can be created based on static configuration files but will ultimately be controlled by the management virtual chain via smart contracts.

Since the system must remain backwards compatible forever (able to audit the old blocks in the chain), changes in configuration can only take place from a certain block height. Until that block height is reached, the original configuration must remain in the system. One way to implement this is place the entire configuration dictionary as a key of a parent map containing block ranges.

#### Virtual chain

```json
{
  "virtual-chain-id": "0x803E2207"
}
```

#### Federation nodes

```json
{
  "nodes": {
    "0x20F082AA2BC5": {
      "reputation": 10
    },
    "0x81CC82AA2BC5": {
      "reputation": 10
    }
  }
}
```

* Keys are node ids (node public keys) encoded in hex.

#### Topology

```json
{
  "node-topology": {
    "0x20F082AA2BC5": {
      "gossip": "cool.node.com:6705"
    },
    "0x81CC82AA2BC5": {
      "gossip": "34.70.102.55:6705"
    },
  },
  "service-topology": {
    "tp1": {
      "type": "TransactionPool",
      "endpoint": "127.0.0.1:80001"
    },
    "ss1": {
      "type": "StateStorage",
      "endpoint": "127.0.0.1:80002"
    }
  }
}
```

* Node topology keys are node ids (node public keys) encoded in hex.

#### Transaction expiration

```json
{
  "transaction-expiration-window-sec": "1800"
}
```

#### Cross services grace wait timeout

```json
{
  "query-grace-timeout-millis": "1000"
}
```

#### Cross services grace wait timeout

```json
{
  "minimum-time-between-empty-blocks-sec": "10"
}
```