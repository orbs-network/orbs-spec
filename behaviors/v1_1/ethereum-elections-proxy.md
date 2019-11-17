# Ethereum Elections proxy
> Mirror the Ethereum PoS voting system onto an Orbs VC. \
This feature facilitates the calculation of the elections results in the Orbs VC. \
Its benefits include: the removal of the dependency in an external mirroring agent and the significant reduction of the Ethereum node requirements - (state archive now supports only ~100 past states - for finality purposes - instead of 5000).


### Design outline:
* The Ethereum PoS proxy state is derived from processing the Ethereum events.
* Continuously syncing the "proxy state" against the Ethereum node, using the `Triggers` mechanism (=> under consensus).
* The current design targets the `_Elections` system contract as the sole client of the `_ElectionsProxy`.
    * Syncing is based on a "pause \ resume" mechanism - governed by the elections events - Syncing until the Elections closing time. 
    * The snapshot of PoS proxy state is maintained in the  [`_ElectionsProxy`](../smart-contracts/system/_ElectionsProxy.md) system contract. 
        * The state snapshot guarantees temporarily consistency across all state queries (same Ethereum block number holds until next syncing step).
        * The proxy simultaneously advances all "Topics" in the PoS state.
    * State revisions are not supported but can easily be added if such a need arise.
    * The `_Elections` system contract deduces the elections results based on the snapshot of PoS proxy state at the target elections closing block. 
        * Elections closing block is the maximal Ethereum block whose timestamp is less than the elections closing time.
    * After the Elections results is calculated the next Elections closing time is set and syncing resumes.

* The proxy syncing runs inside the scope of the Triggers transaction.
    * To avoid failing the triggers transaction entirely
        * errors in calls to `EthereumConnector.EthereumGetLogs` are "wrapped" and handled inside the contract logic (bypassing the `Processor` default panic interrupt behavior). 
    * "Transactional update": apply updates in an atomic fashion after successful processing.
        * Store last_processed (block_number, log_index) to prevent duplication.
        * Assumes processing a single event log is atomic.
        * Future support for triggers isolation will avoid this course and this assumption.      
       
* Bootstrap - some of the Topics require syncing across a long duration of time (ERC20 for example, currently spans over a year an half of Ethereum blocks).
    * Supporting fast sync according to the syncing distance is maintained in the `_ElectionsProxy`.
    * Fallback - if this is not feasible, the external mirroring agent will be used (for at least another election processing), to allow for the long syncing period. 
* Ethereum PoS contracts upgrade requires an upgrade to the Ethereum proxy.
* "Auditability" - the Proxy stores in state a counter per event type, to allow for partial data ingress verification.
* Assumptions: 
    * ["TimeBased Elections"](./time-based-election-periods.md).
    * "Low latency": syncing involves multiple calls to the Ethereum node during the block execution.
    * Support for whitelisted Ethereum nodes only ['Go-Ether' and 'Infura']. The behavior of logs query in various Ethereum node implementations might differ in a way that damages the VC's liveness. 
    * Processing a single event log is atomic.
    
## Elections syncing flow
* During the block execution the `Triggers` mechanism calls the `_ElectionsProxy` system contract `sync()` method.
* The `_ElecionsProxy` retrieves a batch of events' logs across all topics and processes them to update its state. See the [`_ElectionsProxy`](../smart-contracts/system/_ElectionsProxy.md) for a detailed logic.
* Note: syncing to next target is initiated by the `_Elections` system contract.  

#### Participant Services
* `VirtualMachine` - Exposes the Ethereum SDK for events logs retrieval and Ethereum block info queries.
* `EthereumConnector` - Provides the VM with the appropriate Ethereum calls.


## Update to _Elections system contract
* Some of the existing logic should now be extracted \ provided by the ethereum proxy.
#### updateElectionResults()
> Triggered as part of the `Triggers` transaction. Uses the _ElectionsProxy to deduce results.
* If the sync target was not yet initiated \ outdated, set the sync target and Return.
    * Get by calling to `_ElectionsProxy.getSyncTarget()` and compare with the upcoming elections closing time.
    * Set by calling to `_ElectionsProxy.setSyncTarget(timestamp)` with the upcoming elections closing time as the timestamp.
* If the ethereum proxy has synced (`_ElectionsProxy._isSynced()`), start processing the elections results.
    * Use the state queries provided by the `_ElectionsProxy`.
    * Store the results (and related information) in state, with revisions support. 


## SDK / EthereumConnector new functions
* See [Ethereum Connector](../services/crosschain-connector.md) for further details.
#### ether.GetBlockInfoByTime(uint64 ethereum_timestamp) : (uint64 block_number, uint64 block_timestamp).
>  Returns the block info (number, timestamp) of the maximal Ethereum block such that its timestamp < ethereum_timestamp. 
* Input: reference timestamp.
* Output: Ethereum block number and its corresponding timestamp.
* Retrieve the result by calling `EthereumConnector.EthereumGetBlockInfoByTime(block.timestamp, ethereum_timestamp)`.
    * If the provided ethereum_timestamp does not comply with the finality rule, the connector will fail the call.
* On error return (`0`,`0`).

&nbsp;
#### ether.GetBlockInfo() : (uint64 block_number, uint64 block_timestamp).
> Returns the block info (number, timestamp) of the finality Ethereum block.
* Output: Ethereum block number and its corresponding timestamp.
* Retrieve the result by calling `EthereumConnector.EthereumGetBlockInfo(block.timestamp)`.
* On error return (`0`,`0`).

#### ether.GetLogs(events_filter, from_block, to_block) : (log_record_list, block_number)
> Returns a list of event-records which comply to the given filter params. 
* Input: 
    * events_filter : list of tuples (event_type, ethereum_contract_address, event_name, event_ABI) 
        * contract_address - the Ethereum contract address from which logs should originate.
        * event_name - the event name is used to retrieve the event id from the contract ABI.
        * event_type - an id which uniquely mapped to the tuple signature (contract_address, event_name).
    * from_block - Ethereum block number.
    * to_block - Ethereum block number.
* Retrieve the result by calling `EthereumConnector.EthereumGetLogs(block.timestamp, array of [contract_address, event_name, event_ABI], from_block, to_block)`.   
    * If the provided block number range does not comply with the finality rule, the connector will fail the call.
    * The EthereumConnector may return an actual to_block which differs from the given argument; in cases where the range is not feasible (for example the range is in the future).
    * Foreach event log - append back the event_type according to the provided events_filter mapping.
    * Sort the resulting logs by the log's blockNumber and logIndex.
* Output:
    * log_record_list - list of events' logs.
        * log_record := (event_type, blockNumber, logIndex, contract_address, event_name, event_data)
        * Note: logIndex is the location of event log in block.
    * to_block - actual range bound - to_block (allow the EthereumConnector some freedom, rather than "success\failure").
* On error return (nil,`0`).


## ElectionsProxy Contract Outline
> The elections proxy contract continuously syncs the PoS state and provides state query access. \
See [_ElectionsProxy](../../behaviors/smart-contracts/system/_ElectionsProxy.md) for further details.


### Interface

#### setSyncTarget (elections_closing_time)
* Called after calculating the Elections results setting the next elections closing time as the sync target.

#### getSyncTarget() : (timestamp, block_number)
* Returns the current sync target. The timestamp indicates the elections closing time. The block_number (if available - not `0`) is the corresponding elections closing block.

#### isSynced() : bool
* If the PoS state has synced the sync target blockNumber.

#### sync()
* Conditionally sync each the PoS state.

#### PoS state queries 
* getDelegatorInfo(address) : delegatorInfo 
    * delegator_info := (address, balance, stake, cool_down, delegate_to) 
* getGuardians() : list of addresses
* getGuardianInfo(address) : guardianInfo
    * guardian_info := (address, vote)
        * vote := (timestamp, list of addresses (to vote out))
* getValidators() : list of addresses
* getValidatorInfo(address) : validatorInfo
    * validator_info := (address, vote)

* Notes:
    * getValidators() and getGuardians() can return both list of tuple (address, info) avoiding multi calls.

#### state update events
* DelegatorInfoUpdate(address, newBalance, newStake, newCooldDown, newDelegateTO)
* GuardiansUpdate(newGuardians : list addresses)
* ValidatorsUpdate(newValidators : list addresses)
* EthereumEventRecorded(event_type, newEventsCounter)
* EthereumLogProcessed(block_number, log_index)

#### Audit support
* getEventCounter(event_type) : returns the counter of the given event_type.
* getLastProcessedLogInfo() : (block_number, log_index) returns the last_processed field.
