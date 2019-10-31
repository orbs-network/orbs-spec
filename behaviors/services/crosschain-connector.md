# CrossChain Connector

Provides read access to other blockchains like Ethereum and ManagementVC.

#### Interacts with services

* None - Passive, just provides services to others upon request.

&nbsp;
## `Supported Cross-chains`

## Ethereum
* An official Ethereum node connected to main net.
* Runs node locally to avoid trust in third party.
* Node should be kept synchronized and live.
* The service connects to the Ethereum node using a configurable endpoint URL.


&nbsp;
### `Init` (flow)

* Test connection to the relevant Ethereum node.

&nbsp;
### `CallEthereumContract` (method)

> Run a read only contract on Ethereum and return its data.

* Read only call to the Ethereum node using given arguments through IPC.
* The ABI needs to be defined for the ethereum connector to be able to make the call
  * The ABI must contain only one function with the provided function name.
* The contract address need to be supplied
* Needs to support calling with arguments and receiving any output supported by etheruem (solidity)

&nbsp;
### `EthereumGetTransactionLogs` (method)
> Query the log associated with a transaction.

* Query the Ethereum node using the given arguments through IPC.
* The ABI needs to be defined for the ethereum connector to be able to make the call
  * The ABI must contain only one event with the provided event name.
* Filters:
  * Emitting contract address
  * Event name
* Returns a list of events from the desried transaction's receipt that matches the event signature and emitting contract.





## ManagementVC
* Node runs the ManagementVC instance locally - assumes trusted data.
* The ManagementVC instance should be synchronized and live.
* The service connects to the ManagementVC instance using a configurable endpoint URL (intra-node communication).
* NodeConfig - see relevant configuration in [management-chain-readme](../v1_1/management-chain.md):


&nbsp;
### `Init` (flow)
* Test connection to the ManagementVC instance.


&nbsp;
### `GetBlockInfoByTime` (method)
> Query the ManagementVC about its latest committed block prior to a timestamp.
* Shift the provided time reference to overcome inter-node sync across VCs (increase agreement chance).
* This "finality reference time" is deterministically deduced from the provided block time (injects the ManagementVC VirtualChainID and ProtocolVersion).
* NodeConfig: Add time shift constant for ManagementVC.
* Returns the block-info (BlockHeight, BlockTimestamp) of the latest committed block just before the "finality reference time".
* Cache the mapping between the provided time reference and block-info, to save on ManagementVC calls.
* Usage: this method is used as a setup phase, to retrieve ManagementVC state records at a given BlockHeight, under agreement.


&nbsp;
### `CallContract` (method)
> Run a read only contract call on the ManagementVC and return its data.
* Run Call on the ManagementVC using a state at a given blockHeight.
* Input: contract_name, method_name, arguments and blockHeight (injects the ManagementVC VirtualChainID and ProtocolVersion).
* Returns the execution result and output arguments.
* Usage: retrieve ManagementVC state data at a given BlockHeight under agreement.
