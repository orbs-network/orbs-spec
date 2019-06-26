# Query Value

Enables the client to read a state key of a deployed smart contract based on a somewhat latest block (might miss by a couple of blocks). Nevertheless, the block this is based upon is returned in the result.

The response is synchronous, so if the node takes a short while to figure out the response, the client blocks. This read does not require an active subscription on the virtual chain.

Multiple reads can take place at the same time as this is fully parallel. Query Value does not return a merkle proof for the queried state, may be added in a future interface.

## Participants in this flow

* Client
  * `ClientSdk`

* Gateway node
  * `PublicApi`
  * `VirtualMachine`
  * `StateStorage`

## Assumptions for successful flow

* `BlockStorage` is synchronized to latest committed block.

## Flow

* `ClientSdk` sends request to `PublicApi`.

* `PublicApi` of gateway node:
  * Sends the call for execution in `VirtualMachine`.

  * `VirtualMachine` of gateway node:
    * Gets the block height for the call from `StateStorage`.
    * Reads the queried key value from the `StateStorage`

  * Responds to the client.

&nbsp;
## Run Query Flow Diagram

![alt text][run_local_method_flow] <br/><br/>

[run_local_method_flow]: ../_img/run_local_method_flow.png "Call Method"
