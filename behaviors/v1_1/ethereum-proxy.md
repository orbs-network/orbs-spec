# Ethereum proxy
> Mirror the Ethereum PoS voting system onto an Orbs VC. \
This feature facilitates the calculation of the elections results in the Orbs VC. \
Its benefits include: the removal of the dependency in an external mirroring agent and the significant reduction of the Ethereum node requirements - (state archive now supports only ~100 past states - for finality purposes - instead of 5000).


### Design outline:
* The Ethereum PoS proxy state is derived from processing the Ethereum events.
* Continuously syncing the "proxy state" against the Ethereum node, using the `Triggers` mechanism (=> under consensus).
* The current design targets the `_Elections` system contract as the sole client of the `_ElectionsProxy`.
    * Syncing is based on a "pause \ resume" mechanism - governed by the elections events - Syncing until the Elections closing time. 
    * The snapshot of PoS proxy state is maintained in the `_ElectionsProxy` system contract. 
        * The state snapshot guarantees consistency across all state queries (same block number) only at the target block number.
    * State revisions are not supported but can easily be added if such a need arises.
    * The `_Elections` system contract deduces the elections results based on the target snapshot of PoS proxy state.
    * After the Elections results is calculated the next Elections closing time is set and syncing resumes.

* The proxy syncing runs inside the scope of the Triggers transaction.
   * To avoid failing the triggers transaction entirely
        * errors in calls to `EthereumConnector.EthereumGetLogs` are "wrapped" and handled inside the contract logic (bypassing the `Processor` default panic interrupt behavior). 
   * "Transactional update": apply updates in an atomic way:
        * Store current state in a temporary place holder
        * Retrieve all necessary data
        * Process the data in an ascending order + update temporary state + temporally store emitted events (do not emit yet).
        * Upon successful termination
            * Update the state using the temporary state + emit all temporary events.
            * Update last sync block_number
        * On error, do not update the state (nor the events section in the transaction receipt).
     
* The proxy uses a list of Topics to sync and a SyncingScheduler which dictates how to advance each Topic's syncing independently. 
* Bootstrap - some of the Topics require syncing across a long duration of time (ERC20 for example, currently spans over a year an half of Ethereum blocks).
    * SyncingScheduler should support a fast bootstrap - in the order of a few hours.
        * We currently address this issue by implementing a Weighted Round Robin scheduling.
    * Fallback - if this is not feasible, the external mirroring agent will be used (for at least another election processing), to allow for the long syncing period. 
* Ethereum PoS contracts upgrade requires an upgrade to the Ethereum proxy.
* "Auditability" - the Proxy stores in state a Topic events counter (aggregation across all event types) to allow for partial data ingress verification.
* Assumptions: 
    * ["TimeBased Elections"](./time-based-election-periods.md).
    * The events order inside a single transaction do not affect the state transition (we don't have a simple means to sort logs inside the same tx).
    * "Low latency": syncing involves multiple calls to the Ethereum node during the block execution.
    * Support for whitelisted Ethereum nodes only ['Go-Ether' and 'Infura']. The behavior of logs query in various Ethereum node implementations might differ in a way that damages the VC's liveness. 
    
## Elections syncing flow
* During the block execution the `Triggers` mechanism calls the `_ElectionsProxy` system contract `sync()` method.
* The `_ElecionsProxy` retrieves a batch of events' logs by topic and processes them to update its state. See the contract outline below for a detailed logic.
* Note: syncing is initiated by the `_Elections` system contract.  

#### Participant Services
* `VirtualMachine` - Exposes the Ethereum SDK for events logs retrieval and Ethereum block info queries.  <!-- Logs the results data into the tx receipt.  -->
* `EthereumConnector` - Provides the VM with the appropriate Ethereum calls.


## Update to _Elections system contract
* Some of the existing logic is now extracted \  provided by the ethereum proxy.
#### updateElectionResults()
> Triggered as part of the `Triggers` transaction. Uses the _ElectionsProxy to deduce results.
* If the sync target was not yet initiated \ outdated, set the sync target and Return.
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
* On error return `0`.

&nbsp;
#### ether.GetBlockInfo() : (uint64 block_number, uint64 block_timestamp).
> Returns the block info (number, timestamp) of the finality Ethereum block.
* Output: Ethereum block number and its corresponding timestamp.
* Retrieve the result by calling `EthereumConnector.EthereumGetBlockInfo(block.timestamp)`.
* On error return `0`.

#### ether.GetPastEvents(events_filter, from_block, to_block) : (log_record_list, block_number)
> Returns a list of event-records which comply to the given filter params. 
* Input: 
    * events_filter : list of tuples (contract_address, event_name, event_ABI) 
        * contract_address - the Ethereum contract address from which logs should originate.
        * event_name - the event name is used to retrieve the event id from the contract ABI.
    * from_block - Ethereum block number.
    * to_block - Ethereum block number.
* Retrieve the result by calling `EthereumConnector.EthereumGetLogs(block.timestamp, contract_address, event_name, contract_ABI, from_block, to_block)`.   
    * If the provided block number range does not comply with the finality rule, the connector will fail the call.
    * The EthereumConnector may return an actual to_block which differs from the given argument; in cases where the range is not feasible (for example the range is in the future).
* Output:
    * log_record_list - list of events' logs.
        * log_record := (blockNumber, logIndex, contract_address, event_name, event_data)
        * Note: location of event log in block.
    * to_block - actual range bound - to_block (allow the EthereumConnector some freedom, rather than "success\failure").
* On error return (nil,`0`).


## ElectionsProxy Contract Outline
> The proxy contract continuously syncs the Topics defined below and provides state query access. \
Controlled by the _Elections system contract.\
A SyncingScheduler component is used to advance the various Topics syncing process wisely.\
See [_ElectionsProxy](../../behaviors/smart-contracts/system/_ElectionsProxy.md) for further details.


### Interface

#### setSyncTarget (elections_closing_time)
* Called after calculating the Elections results setting the next elections time as the sync target.

#### getSyncTarget() : (timestamp, block_number)
* Returns the current sync target.

#### isSynced() : bool
* Check all Topics are synced to the sync target blockNumber.

#### sync()
* Conditionally sync each Topic based on SyncingScheduler.

#### SyncingScheduler.getTopicsToSync() : list of (Topic, batch_interval)
* Returns the next Topic to sync with ethereum block_number range to retrieve.

#### PoS state queries 
* getTokenBalanceOf(address) : uint64
* getStakingOf(address) : uint64
* getDelegateOf(address) : address
* getGuardians() : list of addresses
* getValidators() : list of addresses
* getEventsCounterByTopic(Topic type) : uint64

#### state update events
* TokenBalanceUpdate(address, newBalance : uint64)
* StakingUpdate(address, newStake : uint64)
* DelegateUpdate(address, newDelegate : address)
* GuardiansUpdate(newGuardians : list addresses)
* ValidatorsUpdate(newValidators : list addresses)
* EthereumEventRecorded(Topic type, newEventsCounter)
