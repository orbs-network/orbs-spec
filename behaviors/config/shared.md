# Shared Configuration

Configuration is an in-memory object given to every service on initialization. The object can be created based on static configuration files but will ultimately transition to being controlled by the management virtual chain via smart contracts.

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
