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
* Returns a list of events from the desired transaction's receipt that matches the event signature and emitting contract.




&nbsp;
## `EthereumGetBlockInfoByTime` (method)
> Returns the block number and timestamp of the maximal Ethereum block before the reference time.
* Query the Ethereum node to find the appropriate block using the given arguments.
* If the provided reference timestamp > current time - ETHEREUM_FINALITY_TIME_COMPONENT, fail the call.
* Find the maximal Ethereum block number whose matching timestamp < reference time.
* If resulting block number > finality block number, fail the call.

&nbsp;
## `EthereumGetBlockInfo` (method)
> Returns the block number and timestamp of the current Ethereum block.
* Query the Ethereum node to find the appropriate block using the given arguments.
* The current Ethereum block returned, complies to the finality rule:
    * Find the maximal Ethereum block number whose matching timestamp < current time - ETHEREUM_FINALITY_TIME_COMPONENT.
    * Reduce this block number by ETHEREUM_FINALITY_BLOCKS_COMPONENT
    * Return the block info of the resulting block number. 


&nbsp;
## `EthereumGetLogs` (method)
> Returns all Ethereum logs entries which comply to the given filter params.\
This method behaves differently on various Ethereum node implementations [Currently supports only 'Go-Ether' and 'Infura'].
* Query the Ethereum node to find the appropriate logs using the given arguments.
* Comply to the finality rule:
    * Verify the block range provided is in the finality range.
        * Get finality block number
        * If to_block > finality block number, set to_block to the finality block number.
        * If from_block > finality block number, fail the call.
* Derive the event signature key from the provided contract ABI using the event_name.
* The ABI needs to be defined for the ethereum connector to be able to make the call
    * The ABI must contain only one event with the provided event name.
* Derive the resulting filter by providing an array of contract_addresses and an array of Topics which contains a single array of event_signatures.
* In case of an "-32005" error - "query returned more than 10k results" (currently exists only in 'Infura' implementation) :
    * split the range into smaller batches. Fetch the smaller chunks and aggregate them.
* On any other error, fail the call. 
