# Ethereum proxy
> Mirror the Ethereum PoS voting system onto an Orbs VC. \
This feature facilitates the calculation of the elections results in the Orbs VC. \
Its benefits include: the removal of the dependency in an external mirroring agent and the significant reduction of the Ethereum node requirements - (state archive now supports only ~100 past states - for finality purposes - instead of 5000).


### Design outline:
* The Ethereum PoS proxy state is derived from processing the Ethereum events.
* Continuously syncing the "proxy state" against the Ethereum node, using the `Triggers` mechanism (=> under consensus).
* Syncing is based on a "pause \ resume" mechanism - governed by the elections events - Syncing until the Elections closing time. After the Elections results is calculated the next Elections closing time is set and syncing resumes.
* The PoS proxy state is maintained in the `_ElectionsProxy` system contract, which also controls the syncing process. The `_Elections` system contract deduces the elections results based on the PoS proxy state.
* The SyncingScheduler component dictates a more flexible and optimized approach to advance each Topic independently and not to query Ethereum node on every Orbs block execution. 
* "Transactional pattern" - before sync get current state into local variable -> get all data and start updating the local variable - on error break and do nothing -> set state based on the new variable  + update last sync block_number.
* To avoid failing the triggers transaction entirely, errors are wrapped in the EthereumConnector call_result and handled inside the contract logic (this only applies to EthereumGetLogs call). Thus, skipping the `Processor` default panic interrupt behavior. 
* Ethereum PoS contracts upgrade \\ TODO.
* "Auditability" - Logging the mirrored data (in the Orbs VC blocks) to facilitate verification of the state correctness. \\ TODO.
* A future optimization would involve caching the data retrieved from calls to avoid multiple calls across topics. 
* Assumptions: 
    * "low latency" The Syncing process involves multiple calls (per each topic) to the Ethereum node, requiring low latency, as this flow occur inside the consensus scope, during the block execution.
    * This feature assumes ["TimeBased Elections"](./time-based-election-periods.md).

## Elections syncing flow
* During the block execution the `Triggers` mechanism calls the `_ElectionsProxy` system contract `sync()` method.
* The `_ElecionsProxy` retrieves a batch of events' logs by topic and processes them to update its state. See the contract outline below for a detailed logic.


#### Participant Services
* `VirtualMachine` - Exposes the Ethereum SDK for events logs retrieval and Ethereum block info queries.  <!-- Logs the results data into the tx receipt.  -->
* `EthereumConnector` - Provides the VM with the appropriate Ethereum calls.



## SDK / EthereumConnector new functions
See [Ethereum Connector](../services/crosschain-connector.md) for further details.
#### ether.GetBlockInfoByTime(uint64 ethereum_timestamp) : (uint64 block_number, uint64 block_timestamp).
>  Returns the block info (number, timestamp) of the latest Ethereum block such that its timestamp < ethereum_timestamp. \
 If ethereum_timestamp > current time - ETHEREUM_FINALITY_TIME_COMPONENT, fail the call.

* Input: reference timestamp.
* Output: Ethereum block number and its corresponding timestamp.
* Retrieve the result by calling `EthereumConnector.EthereumGetBlockInfoByTime(block.timestamp, ethereum_timestamp)`.
* On error return nil.

#### ether.GetLatestBlockInfo() : (uint64 block_number, uint64 block_timestamp).
> Returns the block info (number, timestamp) of the latest Ethereum block which complies with the finality time shift (reference_time := block.timestamp - ETHEREUM_FINALITY_TIME_COMPONENT).
* Output: Ethereum block number and its corresponding timestamp.
* Retrieve the result by calling `EthereumConnector.EthereumGetLatestBlockInfo(block.timestamp)`.
* On error return nil.

#### ether.GetPastEvents(contract_address, event_name, contract_ABI, from_block, until_block) : (log_record_list, block_number, call_result)
> Returns a list of event-records which comply to the given filter params. 
* Input: 
    * contract_address - the Ethereum contract address from which logs should originate.
    * event_name - the event name is used to retrieve the event id from the contract ABI.
    * from_block - Ethereum block number.
    * until_block - Ethereum block number.
<!--* Log the call output as (key, value) into the transaction receipt - inside the off-chain data section.
    * Only on successful call.-->
* Retrieve the result by calling `EthereumConnector.EthereumGetLogs(block.timestamp, contract_address, event_name, contract_ABI, from_block, until_block)`.   
* Output:
    * log_record_list - list of events' logs.
        * log_record := (blockNumber, txIndex, logIndex, contract_address, event_name, event_data)
    * block_number - actual resulting until_block.
    * call_results - success \ failure.


## ElectionsProxy Contract Outline
> The proxy contract continuously syncs the Topics defined below and provides state query access. \
Controlled by the _Elections system contract.\
The SyncingScheduler component is used to advance the various Topics syncing process wisely.\
See [_ElectionsProxy](../../behaviors/smart-contracts/system/_ElectionsProxy.md) for further details.



### SyncingScheduler component
> Syncing different Topics in an independent and balanced way (for example, TokenBalance could sync every 1 minute, while Delegation every 10 minutes - provided the Elections closing time is still out of reach.)


#### getTopicToSync() : Topic, batch_interval
* If isSynced(), Return nil.
* Returns the next Topic to sync with ethereum block_number range to retrieve.
    * Derive the effective sync target based on the Ethereum current block, stored sync target and last synced info.
        * Get the current Ethereum block info (block_number, block_timestamp) by calling `ether.GetLatestBlockInfo()`.
        * Derive the syncTarget.block_number.
            * If syncTarget.block_number is nil &&  currentEthereumBlock.timestamp > syncTarget.ethereum_timestamp (Elections closing time is "finalized" on Ethereum)
                * Get the Elections closing block info by calling `ether.GetBlockInfoByTime(syncTarget.ethereum_timestamp)`.
                * Update the syncTarget.block_number.
        * until_block := min(lastSynced.block_number + batchSize, currentEthereumBlock.block_number, syncTarget.block_number (if not nil)).
        * from_block := min(lastSynced.block_number + 1, .. as above)


#### State outline
> Maintain a list of Topics with the relevant data and each Topic's syncing progress.
* SyncTarget := (timestamp, ethereum_block_number)
* Topic := (contract_event_key list, Topic state, lastSynced)
    * lastSynced : block_number
    * contract_event_key := (contract_address, event_name), contract_creation_block_number
    * Topic types: 
        * TokenBalance
        * Staking
        * Delegation
        * Guardians
        * Validators
* ContractsABIs : mapping (contract_address, contract_ABI)

#### state queries 
* getBalanceOf(address)
* getStakingOf(address)
* getDelegateOf(address)
* getGuardians()
* getGuardianInfo(address)
* getValidators()
* getValidatorInfo(address)

#### state update events
* BalanceUpdated(address, newBalance)


<!--* Topic types: 
    * TokenBalance   [(OrbsToken contract address, events: "Transfer")] 
        * Example for contract_event_key list
            * (OrbsToken_address, "Transfer")
            
        * State : balanceOf : mapping(address => uint256)
        * State query := getBalanceOf(address)
    
    * Staking  [(Orbs Staking contract address, "Staked", "Unstaked", "Withdrew", "Restaked")] 
            event Staked(address indexed stakeOwner, uint256 amount, uint256 totalStakedAmount);
            event Unstaked(address indexed stakeOwner, uint256 amount, uint256 totalStakedAmount);
            event Withdrew(address indexed stakeOwner, uint256 amount, uint256 totalStakedAmount);
            event Restaked(address indexed stakeOwner, uint256 amount, uint256 totalStakedAmount);
        * State : stakingOf : mapping(address => uint256)
        * State query := getStakingOf(address)
        
    * Delegation  [(OrbsVoting contract address,  "Delegate", "Undelegate"), (Orbs ERC20 contract address, "Transfer")] 
        * State: delegateTo : mapping(address => address)
        * State query := getDelegateOf(address)
    * Guardians  [(OrbsGuardians contract address, "GuardianRegistered", "GuardianLeft", "GuardianUpdated")]  
        * State: 
            * guardiansInfo : mapping(address => GuardianInfo)
            * guardiansList : list of addresses
        * State query:
            * getGuardians()
            * getGuardianInfo(address)
        
    * Validators  [(OrbsValidatorsRegistry contract address, "ValidatorLeft", "ValidatorRegistered", "ValidatorUpdated"), (OrbsValidators, "ValidatorApproved", "ValidatorRemoved")] 
        * State: 
            * validatorsInfo : mapping(address => GuardianInfo)
            * validatorsList : list of addresses
        * State query:
            * getValidators()
            * getValidatorInfo(address)
-->


#### setSyncTarget (elections_closing_time)
* Called after calculating the Elections results setting the next elections time as the sync target (`SyncTarget.ethereum_block_number` is set to nil).
    
#### isSynced() : bool
* Check all Topics are synced to the sync target blockNumber.
* `0` blockNumber in SyncTarget implies the target timestamp is in the future => not synced.
 
#### Sync()
* Conditionally sync each Topic based on getTopicToSync().
* If SyncingScheduler.getTopicToSync() is nil, Return.
* Get all events logs records for the topic using SDK call `ether.GetPastEvents(contract_address, event_name, from_block, until_block)`
    * event_log_record := (blockNumber, txIndex, logIndex, contract_address, event_name, event_data)
    * SDK Logs the data into off_chain data section (key, value) of the tx receipt.
* Sort event records by (blockNumber, txIndex, logIndex)
* For each event_log_record
    * Update the state - by processing the event data in accordance to the state transition rules (see Topic specific data processing below). 
    * Emit event on successful state change.
* Update the Topic lastSynced (block_number, block_timestamp)
* Notes: 
    * block_number and block_timestamp refer to Ethereum blocks.
    * On error the Ethereum SDK provides an empty result (nil) to avoid failing the entire trigger tx. The error is handled by breaking its code scope flow.



