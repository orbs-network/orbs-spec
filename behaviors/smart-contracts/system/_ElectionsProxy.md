# `_ElectionsProxy` System Service

The system service charged with maintaining a proxy to the Orbs PoS state, reflected from Ethereum. \
This service queries the Ethereum node for relevant events and updates its state accordingly. \
This service is governed by the `_Elections` system contract. \
Employs a SyncingScheduler to advance the syncing process of the various Topics.

#### Permissions
* This is a system contract that runs under `System` permissions.


#### State
> Maintain a list of Topics with the relevant data and each Topic's syncing progress.

* Constants:
    * list of contracts info:
        * ETHEREUM_CONTRACT_ADDRESS
        * ETHEREUM_CREATION_BLOCK_NUMBER (the block number with the first actual transaction can be used)
        * list of pairs (CONTRACT_EVENT, CONTRACT_ABI)
    * Topic types: 
        * TokenBalance
        * Staking
        * Delegation
        * Guardians
        * Validators    
        
* State variables:
    * SyncTarget := (timestamp, ethereum_block_number)
    * TopicList := list of Topic
        * Topic := (list of event_sync_info, Topic state)
            * events_counter : counter of all Topic events processed.
            * event_sync_info := (event_key, contract_ABI, last_synced)
                * event_key := (contract_address, event_name)
            * last_synced : block_number 
                * Initially, lastSynced.block_number is set to the contract creation block number - 1 (could alter to the first block number which includes events).
            * Topic.getLastSynced() := minimal last_synced across the Topic's events
            * Topic state - for example, TokenBalance.state holds a mapping from address to value.



&nbsp;
## `_init` (method)

#### Permissions
* `Internal` (caller must be the same service).
* `ReadWrite` (potentially changes state).

#### Behavior
* Init the state variables
    * Set the SyncTarget to `0` (indicating the target was not yet set).
    * Populate the TopicList:
        * Init Topic state record.
        * Set the events_counter to `0`.
        * Populate the Topic list of events info (event_sync_info).

&nbsp; 
## `setSyncTarget` (method)
> Sets the given timestamp as the sync target.

#### Permissions
* `External` (caller can only be `_Elections` system.contract).
* `ReadWrite` (might change state).

#### Behavior
* Set the SyncTarget.timestamp to the given timestamp and the target block_number to `0`.
* Implicitly resumes the syncing process as the condition in isSynced() checks as false.
 

&nbsp; 
## `getSyncTarget` (method)
> Returns the current sync target.

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite` (does not change state).


&nbsp; 
## `isSynced` (method)
> Return true if all Topics are synced

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite` (does not change state).

#### Behavior
* Check all Topics are synced to the sync target blockNumber.
* `0` blockNumber in SyncTarget implies the target timestamp is in the future \ not initialized => not synced.

 

&nbsp; 
## `sync` (method)
> Conditionally sync each Topic based on SyncingScheduler.

#### Permissions
* `External` (caller can only be `_Triggers` system.contract).
* `ReadWrite` (might change state).

#### Behavior
* If already synced (`isSynced()`), Return.
* If syncTarget.block_number is `0`, try to update it.
    * Get the Elections closing block info by calling `ether.GetBlockInfoByTime(syncTarget.ethereum_timestamp)`.
    * Update the syncTarget.block_number (if the resulting block info is valid).
* Get list of Topics to sync, together with their respective block_number range, by calling SyncingScheduler.getTopicsToSync().
* Foreach Topic in the returned list (might be empty)
    * Init a current_topic_state (temporary local variable) to current Topic.state.
    * Init an empty local_event_list (temporary local variable) to hold the local emitted events.
    * Get all events logs records for the topic
        * Loop over the list of Topic.contract_event_key (pair of (contract_address, event_name)).
            * Get events using the SDK call `ether.GetPastEvents(contract_address, event_name, from_block, to_block)`
        * event_log_record := (blockNumber, logIndex, contract_address, event_name, event_data)
    * Sort event records by (blockNumber, logIndex)
    * For each event_log_record
        * Update current_topic_state by processing the event data in accordance to the state transition rules.
        * Do not emit an update event. Store it in the local_event_list.
    * On successful processing.
        * Update the Topic state. Set Topic lastSynced (block_number, block_timestamp).
        * Loop over the local_event_list and emit all events by order.
        
* Notes: 
    * block_number and block_timestamp refer to Ethereum blocks.
    * On error the Ethereum SDK provides an empty result (`nil \ 0`) to avoid failing the entire trigger tx. 
        * The error is handled by breaking its code scope flow.
        * An error during a Topic update should not continue to the Topic's next batch (the SyncingScheduler might schedule multiple batches for a Topic during the same call). 
            * On the other hand, it can try to update the next Topic in the Topics list returned from the SyncingScheduler. 
    * Optimization - not all Topics require processing the events in order (thus, allowing for an inconsistent state until the snapshot).
    




### SyncingScheduler - private component 
> Syncing different Topics in an independent and balanced way. \
The implementation supports a Weighted Round Robin scheduling (for example: ERC20 3 batches, Delegate 1).

* BATCH_SIZE - default range unit.

#### getTopicsToSync() : list of (Topic, batch_interval)
> Returns an ordered list of Topics to sync with their corresponding ethereum block_number range to retrieve.


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


<!--#### Topic types in depth

* Topic types: 
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


 

