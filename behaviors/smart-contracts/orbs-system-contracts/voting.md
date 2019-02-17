## Orbs Vote Processing Contracts
> Orbs voting is performed on the Etehreum platform and processed on Orbs as described in the [voting flow](../../flows/voting.md)).

&nbsp;
### Elections
Ownership: none

#### RecordDelegationData(stakeholder, to, delegation_block_height, delegation_tx_index, updated_by)
> Access: internal
> Process an Ethereum delegation transaction and logs it on Orbs.

* If updated_type[`stakeholder`] = `VOTING_CONTRACT` and `updated_by` = `TRANSFER` fail the transaction.
* Check the current account delegation:
  * if the current stakeholder_last_update[`delegator`].`block_height`, `tx_index`} is larger than the {`delegation_block_height`,`delegation_tx_index`} fail the transaction.
* Update delegation map:
  * stakeholder_last_update[`stakeholder`] = {`delegation_block_height`, `delegation_tx_index`}
  * updated_type[`stakeholder`] = `updated_by`.
  * Set agent[`stakeholder`] = `to`.
  

#### RecordDelegation(bytes txid)
> Access: external
> Records an Ethereum delegation by vote contract transaction by calling `RecordDelegationData`.
> RecordDelegation may only be called during the vote receording period.

* Check the that vote mirroring period did not end
  * `ethereum_block_height` < `election_block_height` + `parameter.VOTE_MIRORING_PERIOD_LENGTH_IN_BLOCKS`.
    * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`.
    * If not return an error indicating: Resubmit in the next elections.
* Read the transction's `Delegate(stakeholder, to, delegate_counter)` event emitted by the `OrbsVoring` contract.
  * If no event was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next elections.
  * `delegation_block_height` = `ethereum_block_height`.
  * `delegation_tx_index` = `ethereum_tx_index`.
* Process the delegation by calling `RecordDelegationData(stakeholder, to, delegation_block_height, delegation_block_index, VOTING_CONTRACT)`


#### RecordDelegationByTransfer(bytes txid)
> Access: external
> Records an Ethereum delegation by transfer transaction by calling `RecordDelegationData`.
> RecordDelegation may only be called during the vote mirroring period.

* Check the that vote mirroring period did not end
  * `ethereum_block_height` < `election_block_height` + `parameter.VOTE_MIRORING_PERIOD_LENGTH_IN_BLOCKS`.
    * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`.
    * If not return an error indicating: Resubmit in the next elections.
* Read the transction's `transfer(from, to, tokens` event emitted by the `ERC20` contract.
  * If no event was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check `tokens` = 7. 
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next elections.
  * `delegation_block_height` = `ethereum_block_height`.
  * `delegation_tx_index` = `ethereum_tx_index`.
* Process the delegation by calling `RecordDelegationData(stakeholder, to, delegation_block_height, delegation_block_index, TRANSFER)`


#### RecordVote(byets txid)
> Access: external
> Records a vote transaction taht was sent on Ethereum
> RecordVote may only be called during the vote mirroring period.

* Check the that vote mirroring period did not end
  * `ethereum_block_height` < `election_block_height` + `parameter.VOTE_MIRORING_PERIOD_LENGTH_IN_BLOCKS`.
    * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`.
    * If not return an error indicating: Resubmit in the next election.
* Read the transction's `Vote(activist, nodes_list, vote_counter)` event emitted by the `OrbsVoring` contract.
  * If no was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next elections.
  * `vote_block_height` = `ethereum_block_height`.
  * `vote_tx_index` = `ethereum_tx_index`.
* Check the current `activist`'s vote:
  * if the current activist_last_update[`stakeholder`].{`block_height`, `tx_index`} is larger than the {`vote_block_height`,`vote_tx_index`} fail the transaction.
* Update `activist`'s list and map:
  * voter_last_update[`activist`] = {`vote_block_height`, `vote_tx_index`}
  * Set vote[`activist`] = `nodes_list`.


#### ProcessVoting() : bool
> Access: external
> Process the election results, updates the nodes registery and the next election data.
* State:
  * `current_activist`
  * `current_stakeholder`

* Parameters:
  * `ELECTION_CYCLE_IN_BLOCKS`
  * `VOTE_MIRRORING_PERIOD_LENGTH_IN_BLOCKS`

* Check that the vote recording period has ended
  * `ethereum_block_height` > `election_block_height` + `parameter.VOTE_RECORDING_PERIOD_LENGTH_IN_BLOCKS`
    * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`
    * If not return an error indicating: recording period has not ended.
* Check `election_in_process` != `election_ethereum_block_height`
  * `election_in_process` = `election_ethereum_block_height` 
  * `process_activisits` = TRUE
* If `process_stakeholders` 
  * stake[`current_stakeholder`] = Ethereum.balanceOf(`current_stakeholder`).
  * stakeholders[agent[`current_stakeholder`]].add(`current_stakeholder`). (generate an reverse table)
  * If `current_stakeholder` is last:
    * Clear `process_stakeholders`.
    * Set `calculate_votes`
  * Else `current_stakeholder` = next stakeholder in list. 
* If `process_activisits`:
  * Check activiest_last_update[`current_activist`] has not expired (< `election_block_height` - `parameter.VOTING_VALIDITY_TIME`).
  * stake[`current_activist`] = Ethereum.balanceOf(`current_activist`).
  * If `current_activist` is last:
    * clear `process_activisits`.
    * Set `process_stakeholders`.
    * Else `current_activist` = next activist in list. 
* If `calculate_votes`
  * Call `CalculateVotes()`.
  * Call `CalculateElectedNodes`.
  * Call `NodesRegistery.UpdateElectionResult`.
  * Set `election_ethereum_block_height` = `election_ethereum_block_height` + `ELECTION_CYCLE_IN_BLOCKS`.
  * Clear `calculate_votes`.
* If (`process_stakeholders` OR `process_activisits` OR `calculate_votes`)
  * Return FALSE
* Else 
  * Return TRUE

#### CalculateVotes()
> Calculates the active nodes based on the votes:
* data structures in state:
  * total_voted_stake[`node`] = uint
  * voting_weight[`address`] = uint

* Remove activisits from the stakeholders's list:
  * For every `activist` in activists list:
    * `refernce_agent` = agent[`activist`]
    * Remove `activist` from stakeholders[`refernce_agent`]

* Recursively set the voting_weight of each `activist` or `stakeholder`:
  * If the participant has no `stakeholder`s that delegated to it then voting_weight = stake[participant]
  * Else voting_weight = sum(delegating `stakeholder`s voting_weight) + stake[participant]

* For (every `activist` in activists list):
  * if number of nodes in vote[`activist`] > parameter.VOTES_PER_TOKEN 
    * weight = voting_weight[`activist`] * parameter.VOTES_PER_TOKEN / (number of nodes in vote[`activist`])
  * Else
    * weight = voting_weight[`activist`]
  * For every node in vote[`activist`]:
    * total_voted_stake[`node`] += weight

#### CalculateElectedNodes() 
> Determines the elected nodes based on the election results
> Default: top X nodes are valid.
* Elected nodes = X nodes with the top total_voted_stake[`node`]
* Note: the list of nodes is avaialble by Ethereum.OrbsNodesCandidates.getNodeList()
* Update the state with the elected nodes list.


&nbsp;
### NodesRegistery
> Holds the elected nodes per block height

#### UpdateElectionResult(nodes_list)
> Stores the list of elected nodes per block height
* Nodes[block_height] = nodes_list
* NodesCommencementBlockHeight[index] = current block_height + `parameter.TRANSITION_PERIOD_LENGTH_IN_BLOCKS`

#### GetElectedNodesByHeight(block_height) : nodes_list
> Return the list of elected nodes per requested block_height
* Find the `update_block_height` = largest block in NodesUpdateBlockHeight that is smaller than the provided block_height.
* Return Nodes[`update_block_height`]

#### GetNumberOfNodeUpdates : uint
> Returns the size of the NodesUpdateBlockHeight list, used for migrations.

#### GetElectedNodesByIndex(index) : (block_height, nodes_list)
> Returns the list of elected nodes and reelvant block_height by index, used for migrations.

