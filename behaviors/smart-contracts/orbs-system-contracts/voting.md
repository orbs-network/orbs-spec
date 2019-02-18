# Orbs Vote Processing Contracts
> Orbs voting is performed on the Ethereum platform and processed on Orbs as described in the [voting flow](../../flows/voting.md)).

&nbsp;
## OrbsVoting
Ownership: none

### Global state
* reward[`address`] = uint

### mirrorDelegationData(stakeholder, to, delegation_block_height, delegation_tx_index, updated_by)
> Access: internal
> Process an Ethereum delegation transaction and logs it on Orbs.

#### Check the current account delegation:
* if the current stakeholder_last_update[`stakeholder`].`block_height`, `tx_index`} is larger than the {`delegation_block_height`,`delegation_tx_index`} fail the transaction.
* If updated_type[`stakeholder`] = `VOTING_CONTRACT` and `updated_by` = `TRANSFER` fail the transaction.

#### Update delegation map:
* stakeholder_last_update[`stakeholder`] = {`delegation_block_height`, `delegation_tx_index`}
* updated_type[`stakeholder`] = `updated_by`.
* Set agent[`stakeholder`] = `to`.
  

### mirrorDelegation(bytes txid)
> Access: external
> Mirrors an Ethereum delegation by vote contract transaction by calling `mirrorDelegationData`.
> mirrorDelegation may only be called during the vote mirroring period.

#### Check the that vote mirroring period did not end
* `ethereum_block_height` < `election_block_height` + `parameter.VOTE_MIRRORING_PERIOD_LENGTH_IN_BLOCKS`.
  * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`.
  * If not return an error indicating: Resubmit in the next elections.

#### Read and check the Delegate event
* Read the transaction's `Delegate(stakeholder, to, delegate_counter)` event emitted by the `OrbsVoting` contract.
  * If no event was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next elections.
  * `delegation_block_height` = `ethereum_block_height`.
  * `delegation_tx_index` = `ethereum_tx_index`.

#### Process the delegation
* Process the delegation by calling `mirrorDelegationData(stakeholder, to, delegation_block_height, delegation_block_index, VOTING_CONTRACT)`


### mirrorDelegationByTransfer(bytes txid)
> Access: external
> Mirrors an Ethereum delegation by transfer transaction by calling `mirrorDelegationData`.
> mirrorDelegation may only be called during the vote mirroring period.

#### Check the that vote mirroring period did not end
* `ethereum_block_height` < `election_block_height` + `parameter.VOTE_MIRRORING_PERIOD_LENGTH_IN_BLOCKS`.
  * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`.
  * If not return an error indicating: Resubmit in the next elections.

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
* Process the delegation by calling `mirrorDelegationData(stakeholder, to, delegation_block_height, delegation_block_index, TRANSFER)`


### mirrorVote(bytes txid)
> Access: external
> Mirrors a vote transaction that was sent on Ethereum
> mirrorVote may only be called during the vote mirroring period.

#### Check the that vote mirroring period did not end
  * `ethereum_block_height` < `election_block_height` + `parameter.VOTE_MIRRORING_PERIOD_LENGTH_IN_BLOCKS`.
    * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`.
    * If not return an error indicating: Resubmit in the next election.

#### Read and check the vote event
* Read the transaction's `vote(activist, validators_list, vote_counter)` event emitted by the `OrbsVoting` contract.
  * If no was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next elections.
  * `vote_block_height` = `ethereum_block_height`.
  * `vote_tx_index` = `ethereum_tx_index`.

#### Check the current `activist`'s vote:
* if the current activist_last_update[`stakeholder`].{`block_height`, `tx_index`} is larger than the {`vote_block_height`,`vote_tx_index`} fail the transaction.

#### Check that the activist is in the due-diligence list 
* Call the Ethereum activists contract: Ethereum.Activists.isActivist(`activist`).

#### Update vote map:
* Update `activist`'s list and map:
  * voter_last_update[`activist`] = {`vote_block_height`, `vote_tx_index`}
  * Set vote[`activist`] = `validators_list`.


### ProcessVoting() : bool
> Access: external
> Process the election results, updates OrbsValidatorsConfig and the next election data.
#### State
* `current_activist`
* `current_stakeholder`

#### Check that the vote mirroring period has ended
* `ethereum_block_height` > `election_block_height` + `parameter.VOTE_MIRRORING_PERIOD_LENGTH_IN_BLOCKS`
  * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`
  * If not return an error indicating: mirroring period has not ended.

#### If processing did not start initiate activists processing
* Check `election_in_process` != `election_ethereum_block_height`
  * `election_in_process` = `election_ethereum_block_height` 
  * `process_activists` = TRUE

#### Stakeholders processing
* If `process_stakeholders` 
  * stake[`current_stakeholder`] = Ethereum.balanceOf(`current_stakeholder`).
  * stakeholders[agent[`current_stakeholder`]].add(`current_stakeholder`). (generate an reverse table)
  * If `current_stakeholder` is last:
    * Clear `process_stakeholders`.
    * Set `calculate_votes`
  * Else `current_stakeholder` = next stakeholder in list. 

#### Activists processing
* If `process_activists`:
  * Check activist_last_update[`current_activist`] has not expired (< `election_block_height` - `parameter.VOTING_VALIDITY_TIME`).
  * stake[`current_activist`] = Ethereum.ERC20.balanceOf(`current_activist`).
  * If `current_activist` is last:
    * clear `process_activists`.
    * Set `process_stakeholders`.
    * Else `current_activist` = next activist in list. 

#### Calculations
* If `calculate_votes`
  * Call `CalculateVotes()`.
  * Call `CalculateElectedValidators()` and update `OrbsValidatorsConfig.UpdateElectionResult`.
  * Set `election_ethereum_block_height` = `election_ethereum_block_height` + `parameter.ELECTION_CYCLE_IN_BLOCKS`.
  * Clear `calculate_votes`.

#### Return processing status (completed) 
* If (`process_stakeholders` OR `process_activists` OR `calculate_votes`)
  * Return FALSE
* Else 
  * Return TRUE

### CalculateVotes()
> Calculates the active validators based on the votes:

#### Remove activists from the stakeholders' list:
* For every `activist` in activists list:
  * `reference_agent` = agent[`activist`]
  * Remove `activist` from stakeholders[`reference_agent`]

#### Calculate the hierarchical voting_weight
* Recursively set the voting_weight of each `activist` or `stakeholder`, start with the list of `activists`: 
  * Add participant to participants_list
  * `total_stake` += stake[participant]
  * If the participant has no `stakeholder`s that delegated to it then voting_weight = stake[participant]
  * Else voting_weight = sum(delegating `stakeholder`s voting_weight) + stake[participant]

#### Calculate the activists votes
* For (every `activist` in activists list):
  * if number of validators in vote[`activist`] > parameter.VOTES_PER_TOKEN 
    * weight = voting_weight[`activist`] * parameter.VOTES_PER_TOKEN / (number of validators in vote[`activist`])
  * Else
    * weight = voting_weight[`activist`]
  * For every validator in vote[`activist`]:
    * voted_stake[`validator`] += weight

#### Calculate the rewards weights
* For every participant in participants_list
  * reward_weight[participant] = stake[participant] / `total_stake` * (1 - parameter.ACTIVITY_REWARD_PERCENT) (1 - 0.8 = 0.2)
* For every `activist` in activists list:
  * reward_weight[`activist`] += voting_weight[`activist`] / `total_stake` * parameter.ACTIVITY_REWARD_PERCENT (0.8)

#### Calculate the rewards - TODO move rewards storage to rewards contract (owned)
* For every participant in participants_list
  * reward[participant] += reward_weight * `parameter.ELECTION_REWARD`

### CalculateElectedValidators() 
> Determines the elected validators based on the election results
> Default: top X validators are valid.

#### Filter non due-diligence validators
* Elected validators = X validators with the top voted_stake[`validator`]
* Note: the list of validators is available by `Ethereum.OrbsValidators.getValidators()` / `Ethereum.OrbsValidators.isValidator()`

&nbsp;
### OrbsValidatorsConfig
> Holds the elected validators per block height

#### UpdateElectionResult(validators_list)
> Stores the list of elected validators per block height
* Validators[block_height] = validators_list
* ValidatorsCommencementBlockHeight[index] = current block_height + `parameter.TRANSITION_PERIOD_LENGTH_IN_BLOCKS`

#### GetElectedValidatorsByHeight(block_height) : validators_list
> Return the list of elected validators per requested block_height
* Find the `update_block_height` = largest block in ValidatorsUpdateBlockHeight that is smaller than the provided block_height.
* Return Validators[`update_block_height`]

#### GetNumberOfValidatorUpdates : uint
> Returns the size of the ValidatorsUpdateBlockHeight list, used for migrations.

#### GetElectedValidatorsByIndex(index) : (block_height, validators_list)
> Returns the list of elected validators and relevant block_height by index, used for migrations.

