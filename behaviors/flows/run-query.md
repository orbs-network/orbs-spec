# Run Query Flow

The client performs a read only method call on a service. The call returns the result based on a somewhat latest block (might miss by a couple of blocks). Nevertheless, the block this is based upon is returned as part of the result.

The response is synchronous, so if the node takes a short while to figure out the response, the client blocks. This read does not require an active subscription on the virtual chain.

This read is not under consensus. Multiple reads can take place at the same time, allowing them to run in parallel.

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

* `BlockStorage` is synchronized to the latest committed block.

## Flow

* `ClientSdk` sends request to `PublicApi`.

* `PublicApi` of gateway node:
  * Sends the call for execution in `VirtualMachine`.

  * `VirtualMachine` of gateway node:
    * Gets the block height for the call from `StateStorage`.
    * Executes the smart contract on the relevant `Processor`.
    * Depending on the contract code, the state may be read from either  `StateStorage` or `CrosschainConnector`.

  * Responds to the client.

&nbsp;
## Run Query Flow Diagram

![alt text][run_local_method_flow] <br/><br/>

[run_local_method_flow]: ../_img/run_local_method_flow.png "Call Method"
