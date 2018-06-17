# Call Method Flow

The client performs a read only method call on a service. This returns the result based on a somewhat latest block (might miss by a couple of blocks). Nevertheless, the block this is based upon is returned in the result.

The response is synchronous, so if the node takes a short while to figure out the response, the client blocks. This read does not require an active subscription on the virtual chain.

This read is not under consensus. Multiple reads can take place at the same time as this is fully parallel.

## Participants in this flow

* Client
  * `ClientSdk`

* Gateway node
  * `PublicApi`
  * `BlockStorage`
  * `VirtualMachine`
  * `Processor`
  * `StateStorage`
  * `CrosschainConnector`

## Assumptions for successful flow

* `BlockStorage` is synchronized to latest committed block.

## Flow

* `ClientSdk` sends request to `PublicApi`.

* `PublicApi` of gateway node:
  * Gets the block height for the call from `BlockStorage`.
  * Sends the call for execution in `VirtualMachine`.

  * `VirtualMachine` of gateway node:
    * Executes the smart contract on the relevant `Processor`.
    * Depending on contract code may reads state from `StateStorage` or `CrosschainConnector`.

  * Responds to the client.
