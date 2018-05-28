# Shared Configuration

#### Virtual Chain

```json
{
  "virtual-chain-id": "0x803E2207"
}
```

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

* Node IDs are node public keys encoded in hex
