# CrossChainAPI

> Provides node-internal cross-VCs interface used by other VCs instances of the validator node (not exposed to other nodes).
> Responses are synchronous.
> Currently a single instance per virtual chain per node.


#### Interacts with services

* `VirtualMachine` - Uses it to execute the smart contract calls at a given blockHeight.
* `BlockStorage` - Uses it to respond to queries regarding committed blocks information (height, timestamp).

&nbsp;
## `CrossVC Interfaces`


&nbsp;
## `Init` (flow)

* Initialize the configuration.




&nbsp;
## `GetBlockInfoByTime` (method)

> Returns the blockInfo (height, timestamp) of the latest committed block before the given reference time.
> Query the blockStorage by timestamp. The call is synchronous.

#### Forward call
* Gets the block info by calling `BlockStorage.GetCommittedBlockInfoByTime` with the given timestamp .
* Return the result.
* Notes: 
  * The request format is validated by the HTTP server.

&nbsp;
## `QueryContract` (method)

> Run a read only service method without consensus based on state at the given block height. The call is synchronous.

* Get the block info by querying the blockStorage calling `BlockStorage.GetCommittedBlockInfoByHeight` with the given blockHeight.
#### Forward call
* Execute call on the virtual machine by calling `VirtualMachine.CallSystemContract` using the blockHeight and its corresponding timestamp.
* Return the result.
* Notes: 
  * The request format is validated by the HTTP server.