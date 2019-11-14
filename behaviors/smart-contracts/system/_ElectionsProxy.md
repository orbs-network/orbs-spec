# `_ElectionsProxy` System Service

The system service charged with maintaining a proxy to the Orbs PoS state, reflected from Ethereum. \
This service queries the Ethereum node for relevant events and updates its state accordingly. \
This service is governed by the `_Elections` system contract. \
Granularity of fetching data is by block. All relevant events are fetched from - to block and stored in a queue.


#### Permissions
* This is a system contract that runs under `System` permissions.


#### State
> Maintain a list of Topics with the relevant data and each Topic's syncing progress.

* Constants:
    * list of contracts info:
        * ETHEREUM_CONTRACT_ADDRESS
        * ETHEREUM_CREATION_BLOCK_NUMBER (the block number with the first actual transaction can be used)
        * list of pairs (EVENT_NAME, EVENT_ABI)
    
    * COLLECTING_BLOCKS_BATCH_SIZE - limit the range of blocks to fetch logs from on a single call.
    * CACHE_BLOCKS_LIMIT - an approximation limit to the size of the cached events - before starting to process their data.
    * Events Syncing States: 
        * IDLE
        * COLLECTING_DATA
        * PROCESSING_DATA
        
* State variables:
    * target_timestamp
    * target_block
    * last_fetched : ethereum_block_number indicating all relevant events were retrieved until this point.
    * last_processed 
    * events_filter : list of event tuple (ethereum_contract_address, event_name, event_ABI)
    * events_cache_queue : maintain a queue of events cache ( cleared during the processing )
    * syncing_phase : maintain current phase of the state machine. 
    * events_counter : counter of all Topic events processed.
    * [ max_calls_per_block := limit the number of Ethereum queries for Testing.]
    
* PoS State:
    * delegates := mapping(address => delegator_info)
        * delegator_info := (address, balance, stake, coolDown, delegate_to)
    * guardians := mapping(address => guardian_info)
        * guardian_info := (address, vote)
            * vote := (timestamp, list of addresses (to vote out))
    * validators := mapping(address => validator_info)
        * validator_info := (address, orbs_address)    


&nbsp;
## `_init` (method)

#### Permissions
* `Internal` (caller must be the same service).
* `ReadWrite` (potentially changes state).

#### Behavior
* Init the state variables
    * Populate the events_filter according to the relevant PoS events.
    * Set the events_counter, target_timestamp, target_block to `0` (target_timestamp == target_block == `0` indicates the target was never set yet).
    * Init the PoS State.

&nbsp; 
## `setSyncTarget` (method)
> Sets the given timestamp as the sync target.

#### Permissions
* `External` (caller can only be `_Elections` system.contract).
* `ReadWrite` (might change state).

#### Behavior
* Set the target_timestamp to the given timestamp and the target_block to `0` (implies future block).
* Implicitly this will resume the syncing process.
 

&nbsp; 
## `getSyncTarget` (method)
> Returns the current sync target.

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite` (does not change state).

&nbsp; 
## `isSynced` (method)
> Return true if all the proxy is synced with its target.

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite` (does not change state).

#### Behavior
* Check target_block > 0 and last_processed == target_block ( last_processed <= last_fetched <= target_block always holds).
* `0` blockNumber in sync_target implies the target timestamp is in the future \ not initialized => not synced.

 

&nbsp; 
## `sync` (method)
> Conditionally sync each Topic based on SyncingScheduler.

#### Permissions
* `External` (caller can only be `_Triggers` system.contract).
* `ReadWrite` (might change state).

#### Behavior
* Syncing states logic according to current syncing_phase :    
    * IDLE
        * Check whether to continue syncing.
            * If last_fetched < target_block, or target_timestamp is set but target_block is unknown (future block) 
                * Transition to COLLECTING_EVENTS 
                    
    * COLLECTING_DATA
        * Check whether to continue collecting more data.
            * Get the current Ethereum block info by calling `ether.GetBlockInfo()`.
                * If current Ethereum block number == last_fetched, Return.
                * If current Ethereum block timestamp >= target_timestamp set the target_block
                    * Get the relevant Ethereum block info by calling `ether.GetBlockInfoByTime(target_timestamp)`.
                    * Set the target_block according to the result.
            * If target_block is set and last_fetched >= target_block, or last_fetched - last_processed >= CACHE_BLOCKS_LIMIT    
                * Transition to PROCESSING_DATA  
             
        * Fetch the next batch of data
            * from_block := last_fetched + 1
            * to_block := min(last_fetched + COLLECTING_BLOCKS_BATCH_SIZE, current Ethereum block number, target_block (if already set)).
            * Get all events logs for the above block range by calling `ether.GetPastEvents(events_filter, from_block, to_block)`   
                * The SDK returns an ordered list of events logs (by block_number and log_index)
                * The SDK also returns the actual to_block logs that were able to be fetched without an error - this could be less than to_block.
                * Each event log record is of the form : (block_number, log_index, contract_address, event_name, event_data)
            * Store all events (maintain order) in the events_cache_queue.
            * Set the last_fetched to the returned to_block.
        * Notes: 
            * When fetching data we don't want to mix data from the next elections period - we take into account the elections closing time.
            * If we reached the fetch part it implies last_fetched < current Ethereum block <= target_block.
            
    * PROCESSING_DATA
        * Process data in queue 
            * While events_cache_queue is not empty ( another limit may apply )  
                * Pop the next event log 
                * switch event type
                    * "Transfer" :
                    * "Delegate" :
                    * "Undelegate" :
                    * "Staked" :
                    * "Unstaked"
                    * "Withdrew"
                    * "Restaked"    
                * increase the events_counter
        * Check whether to continue processing data.
            * If events_cache_queue is empty
                * Set last_processed to last_fetched.
                * If last_processed == target_block and target_block > 0 (i.e. synced)
                    * Transition to IDLE  
                * Transition to COLLECTING_DATA 
                    * implies last_processed == last_fetched < target_block
        * Notes:
            * The number of events processed in a single call could be limited if necessary.
                  
* Notes: 
    * Transitions to a different syncing state sets the syncing_phase; could directly start the next state processing. 
    * block_number refer to Ethereum blocks.
    * On error the Ethereum SDK provides an empty result (`nil \ 0`) to avoid failing the entire trigger tx. 
        * The error is handled by breaking its code scope flow.    


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
        
* getEventsCounter() : uint64 - for audit.

* Notes:
    * getValidators() and getGuardians() can return both list of tuple (address, info) avoiding multi calls.

#### state update events
* DelegatorInfoUpdate(address, newBalance, newStake, newCooldDown, newDelegateTO)
* GuardiansUpdate(newGuardians : list addresses)
* ValidatorsUpdate(newValidators : list addresses)
* EthereumEventRecorded(newEventsCounter)


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


 

