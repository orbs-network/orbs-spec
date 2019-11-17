# `_ElectionsProxy` System Service

The system contract charged with maintaining a proxy to the Orbs PoS state, reflected from Ethereum. \
This service queries the Ethereum node for relevant events and updates its state accordingly. \
This service is governed by the `_Elections` system contract. \
Granularity of syncing data is by block. All relevant events in a blocks range are fetched, stored in a queue and then processed.
* State machine design pattern: 
    * (init) - set up the proxy and transition to Ready state
    * READY - syncing target has been reached. 
        * Transition to "SYNCING_DATA" once a new syncing target is set.
    * SYNCING_DATA - single step includes fetching a batch of events logs by blocks and processing this data to update the PoS state. 
        * Transition to "Paused" once the batch has been processed.
            * Transition to "Paused" is based on the distance from target - on bootstrap it avoids pausing.
        * Notes: 
            * Supports multiple iterations during a single step
            * 'Infura' Ethereum node poses a limit on the number of returned entries. 
                * This requires to split a long range into multiple calls with small batch size.
                * The EthereumConnector might also return a block range capped accordingly. 
    * PAUSED - pause for a duration of time 
        * Takes into account the distance from target - decreasing the pausing duration when closing in.
        * Transition to "SYNCING_DATA" once enough time has passed. 
* The state machine performs a single step during the trigger tx.
* First syncing target is also set by the `_Elections` system contract, as it finds the PoS state not synced and target timestamp not set. 

#### Permissions
* This is a system contract that runs under `System` permissions.


#### State
> Maintain a list of Topics with the relevant data and each Topic's syncing progress.

* Constants:
    * multiple ETHEREUM_CONTRACT_ADDRESS
    * multiple (EVENT_TYPE, EVENT_NAME, EVENT_ABI)
    * STARTING_BLOCK - starting point for syncing process: the maximal Ethereum block number prior to all contract deployments (lowest contract deploy block number - 1).
    * SYNCING_BATCH_SIZE - limit the number of ethereum blocks, to fetch events logs from, on a single call.
    * MAX_SYNCING_ITERATIONS_PER_STEP - limit multiple fetching and processing during a single step.
    * BOOTSTRAP_DISTANCE (240 ~ 1hrs of blocks) - criteria to perform multiple calls per step, and allow for to allow for pausing of the syncing process.
    * PAUSING_DURATION (10 minutes ~ 40 Ethereum blocks) - duration to pause the syncing when target time is far.
    * PAUSING_DURATION_CLOSE (1 minute ~ 4 Ethereum blocks) - duration to pause when target time is close.
    * PAUSING_THRESHOLD - distance from target which dictates pausing duration (if target is close - pausing duration is decreased).
    * FINALITY_TIME_SHIFT - used to dictate when the Ethereum state is finalized and reflected in the EthereumConnector.
    * Syncing States (as described above): 
        * READY
        * SYNCING_DATA
        * PAUSED
        
* State variables:
    * target_timestamp
    * target_block
    * last_synced : PoS state was updated until this Ethereum block number (all relevant events were retrieved and processed).
    * is_paused : boolean indicating whether the pausing fields below were set already (added for simplicity ).
    * pause_until : timestamp indicating the end of pausing.
    * events_filter : list of event tuple (event_type, ethereum_contract_address, event_name, event_ABI)
    * events_cache_queue : maintain a queue of events cache ( cleared during the processing )
    * syncing_state : indicator of the state machine current state. 
    * events_counter per event type : counter of Topic event processed - for Audit.
    * syncing_bootstrap - If the target is too far - more than 3 days (bootstrap) will perform syncing without pausing.
    
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
    * Set target_block and target_timestamp to `0` (target_timestamp == `0` Or target_block == `0` indicates the target was never set yet - after init).
    * Set last_synced to STARTING_BLOCK.
    * Init the events_cache_queue, is_paused, pause_until, events_counter.
    * Init the PoS State.
    * Set syncing_state to READY.

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
> Return true if the proxy has synced to its target.

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite` (does not change state).

#### Behavior
* Check target_block > 0 and last_synced == target_block.
* `0` blockNumber in sync_target implies the target timestamp is in the future \ not initialized => not synced.

 

&nbsp; 
## `sync` (method)
> Perform a single step of the state machine.

#### Permissions
* `External` (caller can only be `_Triggers` system.contract).
* `ReadWrite` (might change state).

#### Behavior
* Syncing states logic according to current syncing_state :    
    * READY
        * On a single step. 
            * Check whether to start syncing (if target is in the future).
                * If last_synced < target_block, or target_timestamp is set but target_block is unknown (future block)
                    * Transition to SYNCING_DATA. 
                    
    * SYNCING_DATA
        * On a single step.
            * Conditionally perform multiple syncing iterations, then transition to pause if necessary.
            * Calculate the syncing distance and target_block.
                * Get the current Ethereum block info by calling `ether.GetBlockInfo()`.
                    * If current Ethereum block timestamp >= target_timestamp, set the target_block
                        * Get the relevant Ethereum block info by calling `ether.GetBlockInfoByTime(target_timestamp)`.
                        * Set the target_block according to the result.
                    * If target_block is set 
                        * syncing_distance := target_block - last_synced.
                    * Else, 
                        * syncing_distance :=  current Ethereum block number - last_synced.
            * Calculate the number of iterations to perform, based on syncing_distance.
                * If syncing_distance > BOOTSTRAP_DISTANCE
                    * iterations_count := MAX_SYNCING_ITERATIONS_PER_STEP
                * Else, iterations_count := 1
            * Loop for iterations_count 
                * SYNCING_DATA.(FETCHING)     
                    * Fetch the next batch of data
                        * from_block := last_synced + 1
                        * to_block := min(last_synced + SYNCING_BATCH_SIZE, current Ethereum block number, target_block (if already set)).
                        * Get all events logs for the above block range by calling `ether.GetPastEvents(events_filter, from_block, to_block)`   
                            * The SDK returns an ordered list of events logs (by block_number and log_index)
                            * The SDK also returns the actual to_block logs that were able to be fetched without an error - this could be less than to_block.
                            * Each event log record is of the form : (block_number, log_index, contract_address, event_name, event_data)
                        * Store all events (maintain order) in the events_cache_queue.
                    * Notes: 
                        * When fetching data we don't want to mix data from the next elections period - we take into account the elections closing time.
                        * If we reached the fetch part it implies last_synced < current Ethereum block <= target_block.
                    
                * SYNCING_DATA.(PROCESSING_DATA)
                    * Process data in queue 
                        * While events_cache_queue is not empty ( another limit may apply )  
                            * Pop the next event log 
                            * switch event_type
                                * "Transfer" :
                                * "Delegate" :
                                * "Undelegate" :
                                * "Staked" :
                                * "Unstaked"
                                * "Withdrew"
                                * "Restaked"    
                            * increase the events_counter per type
                            * update last_synced according to event.block_number
                * Set the last_synced to the returned to_block (from the fetch phase above).
            * Conditionally pause syncing.
                * syncing_distance := current Ethereum block number - last_synced (has changed after update).
                * If last_synced == target_block (target reached), 
                    * Transition to READY.
                * If syncing_distance > BOOTSTRAP_DISTANCE, Return (do not pause).
                * Transition to PAUSED.
                    
            * Notes:
                * The number of events processed in a single call could be limited if necessary.
    
    * PAUSED
        * On a single step.
            * Check whether to resume syncing.
                * If now() >= target_timestamp OR now() >= pause_until and is_paused == True
                    * Reset is_paused to False 
                    * Transition to SYNCING_DATA
            * Set the pausing duration if necessary.
                * If is_paused == True, Return
                * Calculate the pausing duration.
                    * If target_timestamp - now() <= PAUSING_DURATION (close to target timestamp),
                        * duration := PAUSING_DURATION_CLOSE.  
                        * note: target_timestamp > now() by the first check above.
                    * Else (target_timestamp - now() > PAUSING_DURATION)
                         * duration := PAUSING_DURATION.  
                * Set the pausing fields
                    * pause_until = now() + duration (it is ok to pause until a time pass the target)
                    * Set is_paused to True
                       
* Notes: 
    * block_number refer to Ethereum blocks.
    * On error the Ethereum SDK provides an empty result (`nil \ 0`) to avoid failing the entire trigger tx. 
        * The error is handled by breaking its code scope flow.
    * now() - internal function takes the finality into account := current_block_timestamp - FINALITY_TIME_SHIFT.


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
        
* getEventsCounterByEventType(event_type) : uint64 - for audit.

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
    * Staking  [(Orbs Staking contract address, "Staked", "Unstaked", "Withdrew", "Restaked")] 
    * Delegation  [(OrbsVoting contract address,  "Delegate", "Undelegate"), (Orbs ERC20 contract address, "Transfer")] 
    * Guardians  [(OrbsGuardians contract address, "GuardianRegistered", "GuardianLeft", "GuardianUpdated")]  
    * Validators  [(OrbsValidatorsRegistry contract address, "ValidatorLeft", "ValidatorRegistered", "ValidatorUpdated"), (OrbsValidators, "ValidatorApproved", "ValidatorRemoved")] 

-->


 

