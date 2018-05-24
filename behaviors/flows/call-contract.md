# Call Contract Flow

The client performs a read only contract call. This returns the result based on a somewhat latest block (might miss by a couple of blocks). Nevertheless, the block this is based upon is returned in the result.

The response is synchronous, so if the node takes a short while to figure out the response, the client blocks. This read does not require an active subscription on the virtual chain.

This read is not under consensus. Multiple reads can take place at the same time as this is fully parallel.

## Participants in this flow

* Client
  * `ClientSdk`

* Receiving node
  * `PublicApi`
  * `ConsensusCore`
  * `VirtualMachine`
  * `Processor`
  * `StateStorage`
  * `SidechainConnector`

## Assumptions for successful flow

* `ConsensusCore` is synchronized to somewhat latest block.

## Flow

* `ClientSdk` sends request to `PublicApi`.

* `PublicApi` of receiving node:
  * Retains session info so the response will eventually arrive to this client.
  * Performs inexpensive checks on the call.
  * Prepares a batch of some calls for execution.
  * Marks the batch with the latest block height by calling `ConsensusCore.GetTopBlockHeight`.
  * Sends the batch for execution by calling `VirtualMachine.CallContract`.

* `VirtualMachine` of receiving node:
  * Checks the sender signature.
  * Executes the smart contract on the relevant `Processor`.
  * Reads relevant state for execution from `StateStorage` or `SidechainConnector`.
  * Returns the batch results to `PublicApi`.

* `PublicApi` of receiving node responds to the client.

## Flow diagram

TODO
