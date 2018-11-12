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
* The service connects to the Ethereum node using [configurable](../config/services.md) Node DNS address and port.

&nbsp;
## `Init` (flow)

* Test connection to the relevant node.

&nbsp;
## `CallEthereumContract` (method)

> Run a read only contract on Ethereum and return its data.

* Read only call to the Ethereum node using given arguments through IPC.
* The ABI needs to be defined for the ethereum connector to be able to make the call
* The contract address need to be supplied
* Needs to support calling with arguments and receiving any output supported by etheruem (solidity)


&nbsp;
## `EthereumGetTransactionLogs` (method)
> Query the logsassociated with a transaction, with optional filter on contract, event signature.

* Query the Ethereum node using the given arguments through IPC.
* Filters:
    * Contract address (a NULL value indicates no fitler)
    * Event signatrue (a NULL value indicates no fitler). EventSignature uses the following format: `TransferOutEvent(uint32, byte32, uint256)`.

