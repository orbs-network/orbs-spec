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




&nbsp;
## `EthereumGetBlockInfoByTime` (method)
> Retrieve info about the latest Ethereum block such that its timestamp < provided ethereum_timestamp.

* Query the Ethereum node to find the appropriate block using the given arguments.
* Comply to the finality time shift.
    * If the provided ethereum_timestamp > current time - ETHEREUM_FINALITY_TIME_COMPONENT, fail the call.
        * current time is provided as reference_time.
* Returns the block info (uint64 block_number, uint64 block_timestamp) of the resulting block.



&nbsp;
## `EthereumGetLatestBlockInfo` (method)
> Retrieve info about the latest Ethereum block such that its timestamp < reference_time under finality time shift.

* Query the Ethereum node to find the appropriate block using the given arguments.
* Comply to the finality time shift:
    * Find the maximal Ethereum block whose timestamp < reference_time - ETHEREUM_FINALITY_TIME_COMPONENT.
* Returns the block info (uint64 block_number, uint64 block_timestamp) of the resulting block.



&nbsp;
## `EthereumGetLogs` (method)
> Returns all Ethereum logs entries which comply to the given filter params. 
* Query the Ethereum node to find the appropriate logs using the given arguments.
* Comply to the finality time shift:
    * Get Latest finalized block by calling EthereumGetLatestBlockInfo
    * Verify the block range provided is in the finality range.
        * finalized block number >  until_block_number (provided in arguments).
* Derive the event signature key from the provided contract ABI using the event_name.
* The ABI needs to be defined for the ethereum connector to be able to make the call
  * The ABI must contain only one event with the provided event name.
* Returns:
    * call_result - indicates the call success or failure.
    * block_number - indicates the last filtered block_number.
    * list of events logs entries of the form (blockNumber, txIndex, logIndex, ethereum_abi_packed = event_data).


