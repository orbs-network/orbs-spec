# Orbs Vote Processing Contracts
> Orbs voting is performed on the Ethereum platform and processed on Orbs as described in the [voting flow](./election-flow.md)

> For more defails on Orbs PoS architecture, see [Orbs PoS architecture](https://github.com/orbs-network/orbs-spec/tree/pos2/pos-architecture)

&nbsp;
## OrbsVoting
Ownership: none

### Global state
#### Delegation and voting
* agent[`delegator`] - the account the delegator delegated to.
* voted_validators[`guardian`] - voted in the election

#### Election data
* `election_block_height` - The last Ethereum block height for delegation and voting, the reference block_height for the stake read. 

#### Reward data
* reward[`address`] - Cumulative rewards awarded to an address (as a Guardian, Delegator or Validator)

#### Last election data
* received_votes[`validator`] - the total stake that was voted to the validator.
* `total_voting_stake` - the total stake that participated in the election.
* `total_in_circulation` - the total stake in circulation.

#### Elections Parameters
* `VOTE_MIRRORING_PERIOD_LENGTH_IN_BLOCKS` - The number of Ethereum blocks after the `election_block_height` during which mirroring is still allowed for the election.
  * Default: 550 (~ 2 hours)
* `ELECTION_CYCLE_IN_BLOCKS` - The number of Ethereum blocks between elections, determines the next  `election_block_height`.
  * Default: 20000 (~ 3 days)
* `VOTING_VALIDITY_TIME` - The number of Ethereum blocks during which a Guardian's voting is valid
  * Default: 45500 (~ 7 days)
* `TRANSITION_PERIOD_LENGTH_IN_BLOCKS` - The number of **Orbs** blocks (of the virtual chains) until an election decision is applied.
  * Default: 0 - the next new committee is applied in the next block.
* `MAXIMUM_ELECTED_VALIDATORS` - the maximum Validators that can initially be elected
  * Default: 22
* `DISAPPROVAL_THRESHOLD_PERCENT` - the percentage of the total participating stake required to disapprove a Validator
  * Default: 70%
* `FIRST_ELECTION_BLOCK` - the Ethereum block of the first election. 
  * 7528900 (approximately April 8, 2019 noon UTC)
* `MINIMUM_VALIDATORS` - the minimal number of Validators to be elected
  * 7

#### Rewards Parameters
* `PARTICIPATION_MAX_ANNUAL_REWARD` - The maximum annual reward awarded to stakeholders (Delegators or Guardians) for participation. (Delegators or Guardians)
  * Default: 60M (ORBS)
* `PARTICIPATION_MAX_STAKE_REWARD_PERCENT` - The maximum award in each election as a percent of the participant stake. (taken into consideration when the total participating stake is low: less than `PARTICIPATION_MAX_REWARD`/`PARTICIPATION_MAX_STAKE_REWARD_PERCENT`).
  * Default: 8%
* `VALIDATOR_INTRODUCTION_PROGRAM_ANNUAL_REWARD` - a fixed reward given to Validators at the initial stage of the network.
  * Default: 1M (ORBS)
* `VALIDATORS_STAKE_REWARD_PERCENT` - the Validator reward as a percent of her stake.
  * Default: 4%
* `GUARDIANS_MAX_ANNUAL_REWARD` - maximum annual award for the Guardians Excellence Program
  * Default: 40M (ORBS)
* `GUARDIANS_MAX_DPOS_REWARD_PERCENT` - The maximum award in each election as a percent of the Guardian delegated stake (including its own). Taken into consideration when the total participating stake in the Excellence Program is low: < `GUARDIANS_MAX_ANNUAL_REWARD`/`GUARDIANS_MAX_DPOS_REWARD_PERCENT`).
  * Default: 10%
* `EXCELLENCE_PROGRAM_NUMBER_OF_GUARDIANS` - number of Guardians with the top delegated stake that participate in the Guardians Excellence Program.
  * Default: 10

#### General
* `ETHEREUM_AVG_BLOCK_TIME_SEC` = 13.28
* `SEC_IN_A_YEAR` = 31536000
* `ELECTION_PARTICIPATION_MAX_REWARD` = (`PARTICIPATION_MAX_ANNUAL_REWARD` * `ELECTION_CYCLE_IN_BLOCKS` * `ETHEREUM_AVG_BLOCK_TIME_SEC` / `SEC_IN_A_YEAR`) = 505328
* `ELECTION_GUARDIANS_MAX_REWARD` = (`GUARDIANS_MAX_ANNUAL_REWARD` * `ELECTION_CYCLE_IN_BLOCKS` * `ETHEREUM_AVG_BLOCK_TIME_SEC` / `SEC_IN_A_YEAR`) = 336885
* `ELECTION_VALIDATOR_INTRODUCTION_MAX_REWARD` = `VALIDATOR_INTRODUCTION_PROGRAM_ANNUAL_REWARD` * `ELECTION_CYCLE_IN_BLOCKS` * `ETHEREUM_AVG_BLOCK_TIME_SEC` / `SEC_IN_A_YEAR`) = 8423
* `ELECTION_PARTICIPATION_MAX_STAKE_REWARD_PERCENT` =  `PARTICIPATION_MAX_STAKE_REWARD_PERCENT` * `ELECTION_CYCLE_IN_BLOCKS` * `ETHEREUM_AVG_BLOCK_TIME_SEC` / `SEC_IN_A_YEAR`) = 673769660 / 1T (1E12)
* `ELECTION_GUARDIANS_MAX_DPOS_REWARD_PERCENT` =  `GUARDIANS_MAX_DPOS_REWARD_PERCENT` * `ELECTION_CYCLE_IN_BLOCKS` * `ETHEREUM_AVG_BLOCK_TIME_SEC` / `SEC_IN_A_YEAR`) = 842212075 / 1T (1E12)
* `ELECTION_VALIDATORS_STAKE_REWARD_PERCENT` =  `VALIDATORS_STAKE_REWARD_PERCENT` * `ELECTION_CYCLE_IN_BLOCKS` * `ETHEREUM_AVG_BLOCK_TIME_SEC` / `SEC_IN_A_YEAR`) = 336884830 / 1T (1E12)

### Elections Terminology:
* Effective Elections Block Number
  * The election event block number that the currently elected validators were elected at.
    * Note: can be read from the elections history table.
* Current Elections Block Number
  * The elections that is currently undergoing (either voted for or in processing)
* Next Elections Block Number
  * The next elections after the current elections

### Getting the current election block number
> The current election block number is set as the next election event block number in the global elections schedule.
* Get the current Ethereum block number using `Ethereum.GetBlockNumber()`
* If current Ethereum block < `FIRST_ELECTION_BLOCK` 
  * first election = `FIRST_ELECTION_BLOCK`
* Else
  * first election = the first block number in the elections schedule that is above the current Ethereum block.
<!--
* Implementation suggestion: first election = Current Ethereum block + `ELECTION_CYCLE_IN_BLOCKS` + (`FIRST_ELECTION_BLOCK` - Current Ethereum block) % `ELECTION_CYCLE_IN_BLOCKS`) 
-->

### initFirstElection()
> Called once and init the first election based on the current Ethereum block. Implicitly called by internal reads of the `election_event_block`.
* If `election_event_block` == 0, set `election_event_block` to the first election. 

### mirrorDelegationData(delegator, to, delegation_block_height, delegation_tx_index, updated_by)
> Access: internal
> Process an Ethereum delegation transaction and logs it on Orbs.

#### Check the current account delegation:
* if the current delegator_last_update[`delegator`].`block_height`, `tx_index`} is larger than the {`delegation_block_height`,`delegation_tx_index`} fail the transaction.
* If updated_type[`delegator`] = `VOTING_CONTRACT` and `updated_by` = `TRANSFER` fail the transaction.

#### Update delegation map:
* delegator_last_update[`delegator`] = {`delegation_block_height`, `delegation_tx_index`}
* updated_type[`delegator`] = `updated_by`.
* if (`to` == 0) or (`to` == `delegator`):
  * Set agent[`delegator`] = address(0).
* Else
  * Set agent[`delegator`] = `to`.
  

### mirrorDelegation(bytes txid)
> Access: external
> Mirrors an Ethereum delegation by vote contract transaction by calling `mirrorDelegationData`.
> mirrorDelegation may only be called during the vote mirroring period.

#### Check the that vote processing did not start yet
* If processing has started return an error indicating: Resubmit in the next elections.

#### Read and check the Delegate event
* Read the transaction's `Delegate(delegator, to, delegate_counter)` event emitted by the `OrbsVoting` contract.
  * If no event was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next elections.
  * `delegation_block_height` = `ethereum_block_height`.
  * `delegation_tx_index` = `ethereum_tx_index`.

#### Process the delegation
* Process the delegation by calling `mirrorDelegationData(delegator, to, delegation_block_height, delegation_block_index, VOTING_CONTRACT)`

### mirrorDelegationByTransfer(bytes txid)
> Access: external
> Mirrors an Ethereum delegation by transfer transaction by calling `mirrorDelegationData`.
> mirrorDelegation may only be called during the vote mirroring period.

#### Check the that vote processing did not start yet
* If processing has started return an error indicating: Resubmit in the next elections.

#### Read and check the transfer event
* Read the transaction's `transfer(from, to, tokens` event emitted by the `ERC20` contract.
  * If no event was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check `tokens` = 7. 
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next elections.
  * `delegation_block_height` = `ethereum_block_height`.
  * `delegation_tx_index` = `ethereum_tx_index`.

#### Process the delegation
* Process the delegation by calling `mirrorDelegationData(delegator, to, delegation_block_height, delegation_block_index, TRANSFER)`


### mirrorVote(bytes txid)
> Access: external
> Mirrors a vote transaction that was sent on Ethereum
> mirrorVote may only be called during the vote mirroring period.

#### Check the that vote mirroring period did not end
  * `ethereum_block_height` < `election_block_height` + `parameter.VOTE_MIRRORING_PERIOD_LENGTH_IN_BLOCKS`.
    * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`.
    * If not return an error indicating: Resubmit in the next election.

#### Read and check the vote event
* Read the transaction's `vote(guardian, validators_list, vote_counter)` event emitted by the `OrbsVoting` contract.
  * If no was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next elections.
  * `vote_block_height` = `ethereum_block_height`.
  * `vote_tx_index` = `ethereum_tx_index`.

#### Check the current `guardian`'s vote:
* if the current guardian_last_update[`delegator`].{`block_height`, `tx_index`} is larger than the {`vote_block_height`,`vote_tx_index`} fail the transaction.

#### Check that the guardian is in the due-diligence list 
* Call the Ethereum guardians contract: Ethereum.Guardians.isGuardian(`guardian`).

#### Update vote map:
* Update `guardian`'s votes map:
  * guardian_last_update[`guardian`] = {`vote_block_height`, `vote_tx_index`}
  * Set voted_validators[`guardian`] = `validators_list`.


### ProcessVoting() : bool
> Access: external
> Process the election results, updates OrbsValidatorsConfig and the next election data.
#### State
* `current_guardian`
* `current_delegator`
* `processing_state`

#### Check that the vote mirroring period has ended
* `ethereum_block_height` > `election_block_height` + `parameter.VOTE_MIRRORING_PERIOD_LENGTH_IN_BLOCKS`
  * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`
  * If not return an error indicating: mirroring period has not ended.

#### If processing did not start initiate guardians processing
* Check `election_in_process` != `election_ethereum_block_height`
  * `election_in_process` = `election_ethereum_block_height` 
  * `process_guardians` = TRUE

#### Delegators processing
* If `process_delegators` 
  * stake[`current_delegator`] = Ethereum.balanceOf(`current_delegator`).
  * delegators[agent[`current_delegator`]].add(`current_delegator`). (generate an reverse table)
  * If `current_delegator` is last:
    * Clear `process_delegators`.
    * Set `calculate_votes`
  * Else `current_delegator` = next delegator in list. 

#### Guardians processing
* If `process_guardians`:
  * Check guardian_last_update[`current_guardian`] has not expired (< `election_block_height` - `parameter.VOTING_VALIDITY_TIME`).
  * stake[`current_guardian`] = Ethereum.ERC20.balanceOf(`current_guardian`).
  * If `current_guardian` is last:
    * clear `process_guardians`.
    * Set `process_delegators`.
    * Else `current_guardian` = next guardian in list. 

#### Calculations
* If `calculate_votes`
  * Call `CalculateVotes()`.
  * Call `CalculateElectedValidators()` and update `OrbsValidatorsConfig.UpdateElectionResult`.
  * Set `election_ethereum_block_height` = `election_ethereum_block_height` + `parameter.ELECTION_CYCLE_IN_BLOCKS`.
  * Clear `calculate_votes`.

#### Return processing status (completed) 
* If (`process_delegators` OR `process_guardians` OR `calculate_votes`)
  * Return FALSE
* Else 
  * Return TRUE

### CalculateVotes()
> Calculates the active validators based on the votes:

#### Remove the Guardians from the delegation list (implicitly delegate to themselves)
* Get the `guardians_list` from Ethereum 
  * Call `Ethereum.OrbsGuardians.getGuardians()`
* For every `guardian` in guardians list:
  * `reference_agent` = agent[`guardian`]
  * Remove `guardian` from delegators[`reference_agent`]

#### Calculate the hierarchical voting_stake
* Recursively set the voting_stake of each `guardian` or `delegator`, start with the `guardians_list`: 
  * Add participant to participants_list
  * `total_voting_stake` += stake[participant]
  * If the participant has no `delegator`s that delegated to it then voting_stake = stake[participant]
  * Else voting_stake = sum(delegating `delegator`s voting_stake) + stake[participant]

#### Calculate the elected validators
* Get the validators_list
  * Call `Ethereum.OrbsValidators.getValidators()`
  * Alternatively use `Ethereum.OrbsValidators.isValidator(Validator)` based on the voted addresses.

* For (every `guardian` in guardians list):
  * For every validator in voted_validators[`guardian`]:
    * received_votes[`validator`] += voting_stake[`guardian`].

* Calculate `disapproval_threshold` = `total_voting_stake` * `DISAPPROVAL_THRESHOLD_PERCENT`

* elected_validators = validators in `validators_list` with: 
  * received_votes[`validator`] < `disapproval_threshold`.

#### Minimal number of Validators protection
* If number of elected_validators < `MINIMUM_VALIDATORS`
  * elected_validators = validators in `validators_list`

#### ------------------------------------------------------------------------------
#### TOP22 Calculation - Not Implemented

* Generate `currently_elected_approved_list`
  * `currently_elected_approved_list` = validators in `validators_list` with:
    * received_votes[`validator`] < `disapproval_threshold`.
    * in `currently_elected_validators`.

* Find the top `top_candidate`
  * `top_candidate` is the validator in `validators_list` with the most stake[`validator`] that meets:
    * received_votes[`validator`] < `disapproval_threshold`.
    * not in `currently_elected_validators`.

* If `currently_elected_approved_list` < `MAXIMUM_ELECTED_VALIDATORS`
  * elected_validators.add(`currently_elected_approved_list`)
  * elected_validators.add(`top_candidate`)
* Else
  * Sort the `currently_elected_approved_list` by their stake
  * elected_validators = `MAXIMUM_ELECTED_VALIDATORS` - 1 validators
  * `bottom_validator` = the bottom validator in `currently_elected_approved_list`.
  * If stake[`bottom_validator`] > stake[`top_candidate`]
    * elected_validators.add(`bottom_validator`)
  * Else
    * elected_validators.add(`top_candidate`)

#### ------------------------------------------------------------------------------

#### Calculate the participation reward
* Calculate the participation reward for the election
  * `election_participation_reward` = min(`ELECTION_PARTICIPATION_MAX_REWARD`, `total_voting_stake` * `ELECTION_PARTICIPATION_MAX_REWARD_PERCENT`)
* Calculate the participation reward
  * For every delegator or guardian:
    * participation_reward[participant] += stake[participant] / `total_voting_stake` * `election_participation_reward`

#### Calculate the Guardians Excellence Program reward
* Calculate the participants in the Guardians Excellence Program for the election
  * `excellence_program_participants` = the top `EXCELLENCE_PROGRAM_NUMBER_OF_GUARDIANS` with the most `voting_stake`.
  * `total_excellence_program_stake` = total `voting_stake` of the excellence_program_participants.  
* Calculate the Guardians Excellence Program reward for the election
  * `election_guardians_reward` = min(`ELECTION_GUARDIANS_MAX_REWARD`, `total_excellence_program_stake` * `ELECTION_GUARDIANS_MAX_DPOS_REWARD_PERCENT`).
* Calculate the Guardians reward
  * For every `guardian` in excellence_program_participants:
    * guardian_excellence_reward[`guardian`] += voting_stake[`guardian`] / `total_excellence_program_stake` * `election_guardians_reward`

#### Calculate the Validators reward
* For every `elected_validator` in `elected_validators`
  * validator_reward[`elected_validator`] += stake[`elected_validator`] * `ELECTION_VALIDATORS_STAKE_REWARD_PERCENT` + `VALIDATOR_INTRODUCTION_PROGRAM_REWARD`


&nbsp;
## Getters Interface

#### General
* getEffectiveElectionBlockNumber() : Ethereum_block
* getCurrentElectionBlockNumber() : Ethereum_block
* getNextElectionBlockNumber() : Ethereum_block
* getElectionPeriod() : Number Of Ethereum_blocks
* getNumberOfElections()
* getElectedValidatorsOrbsAddress() : list of Validators
* getElectedValidatorsEthereumAddress() : list of validators 
* getCurrentEthereumBlockNumber()
* getProcessingStartBlockNumber()
* getMirroringEndBlockNumber()

#### Extended Last Election Results
* GetGuardianVotingWeight(Guardian) : uint
* GetGuardianStake(Guardian) : uint
* GetValidatorVotes(ValidatorEthereumAddress) : uint
* GetValidatorStake(ValidatorEthereumAddress) : uint
* GetTotalStake() : uint
* GetExcellenceProgramGuardians() : list of guardians 

#### Historical data
* getElectedValidatorsOrbsAddressByBlockHeight(orbs_block_height) : elected_validators
* getElectedValidatorsEthereumAddressByBlockHeight() : elected_validators
* getElectedValidatorsOrbsAddressByIndex(index) : elected_validators (Ethereum addresses).
* getElectedValidatorsEthereumAddressByIndex(index) : elected_validators (Ethereum addresses).
* getElectedValidatorsOrbsAddressByIndex(index) : elected_validators (Ethereum addresses).
* getElectionEventBlockNumberByIndex(index) : Election event Ethereum block number
* getElectionResultsBlockHeightByIndex(index) : Apply results orbs_block_height

#### Rewards
* GetCumulativeParticipationReward(delegator) : uint (Integer ORBS)
* GetCumulativeValidatorReward(validatorEthereumAddress) : uint (Integer ORBS)
* GetCumulativeGuardiansExcellenceReward(guardian) : uint (Integer ORBS)



