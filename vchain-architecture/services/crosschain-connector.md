# Cross-chain Connector

Runs nodes for other blockchains like Ethereum and provides read access to them.

#### Interacts with services

* None - Passive, just provides services to others upon request.

&nbsp;
## `Supported Cross-chains`

#### Ethereum
* An official Ethereum node connected to main net.
* Runs node locally to avoid trust in third party.
* Node should be kept synchronized and live.
* The service connects to the Ethereum node using a configurable endpoint URL.

&nbsp;
## `Init` (flow)

* Test connection to the relevant node.

&nbsp;
## `CallEthereumContract` (method)

> Run a read only contract on Ethereum and return its data.

* Read only call to the Ethereum node using given arguments through IPC.
* The ABI needs to be defined for the ethereum connector to be able to make the call
  * The ABI must contain only one function with the provided function name.
* The contract address need to be supplied
* Needs to support calling with arguments and receiving any output supported by etheruem (solidity)

&nbsp;
## `EthereumGetTransactionLogs` (method)
> Query the log associated with a transaction.

* Query the Ethereum node using the given arguments through IPC.
* The ABI needs to be defined for the ethereum connector to be able to make the call
  * The ABI must contain only one event with the provided event name.
* Filters:
  * Emitting contract address
  * Event name
* Returns a list of events from the desried transaction's receipt that matches the event signature and emitting contract.
