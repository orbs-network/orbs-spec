# Virtual Chain Service Interfaces

A virtual chain instance (ONG) is divided into micro-services that have explicit interfaces defined using Google Protobuf language.

The Orbs virtual chain [reference implementation](https://github.com/orbs-network/orbs-network-go) relies on these interfaces compiled to Golang type definitions. The compliation process is available [here](../types/README.md) and must be performed and published whenever the interfaces change.