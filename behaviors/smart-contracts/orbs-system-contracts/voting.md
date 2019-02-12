## Orbs Vote Processing Contracts
> Orbs voting is performed on the Etehreum platform and processed on Orbs as described in the [voting flow](../../flows/voting.md)).

&nbsp;
### VoteProcessing
Ownership: none

#### RecordDelegationData(delegator, to, delegation_block_height, delegation_tx_index, updated_by)
> Access: internal
> Process an Ethereum delegation transaction and logs it on Orbs.

* If `updated_by` = `VOTING_CONTRACT` and `updated_by` = `TRANSFER` fail the transaction.
* Check the current account delegation:
  * if thre current delegator_last_update[`delegator`].`block_height`, `tx_index`} is larger than the {`delegation_block_height`,`delegation_tx_index`} fail the transaction.
* Update delegation map:
  * delegator_last_update[`delegator`] = {`delegation_block_height`, `delegation_tx_index`}
  * updated_type[`delegator`] = `updated_by`.
  * If delegatee[`to`] already exist (not empty)  
    * Remove the delegator from delegators[delegatee[`delegator`]].
  * Set delegatee[`delegator`] = `to`.
  

#### RecordDelegation(bytes txid)
> Access: external
> Records an Ethereum delegation by vote contract transaction by calling `RecordDelegationData`.
> RecordDelegation may only be called during the vote receording period.

* Check the that vote recording period did not end
  * `ethereum_block_height` < `election_block_height` + `parameter.VOTE_RECORDING_PERIOD_LENGTH_IN_BLOCKS`.
    * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`.
    * If not return an error indicating: Resubmit in the next election vote receording period.
* Read the transction's `Delegate(delegator, to, delegate_counter)` event emitted by the `OrbsVoring` contract.
  * If no event was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next election vote receording period.
  * `delegation_block_height` = `ethereum_block_height`.
  * `delegation_tx_index` = `ethereum_tx_index`.
* Process the delegation by calling `RecordDelegationData(delegator, to, delegation_block_height, delegation_block_index, VOTING_CONTRACT)`


#### RecordDelegationByTransfer(bytes txid)
> Access: external
> Records an Ethereum delegation by transfer transaction by calling `RecordDelegationData`.
> RecordDelegation may only be called during the vote receording period.

* Check the that vote recording period did not end
  * `ethereum_block_height` < `election_block_height` + `parameter.VOTE_RECORDING_PERIOD_LENGTH_IN_BLOCKS`.
    * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`.
    * If not return an error indicating: Resubmit in the next election vote receording period.
* Read the transction's `transfer(from, to, tokens` event emitted by the `ERC20` contract.
  * If no event was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check `tokens` = 7. 
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next election vote receording period.
  * `delegation_block_height` = `ethereum_block_height`.
  * `delegation_tx_index` = `ethereum_tx_index`.
* Process the delegation by calling `RecordDelegationData(delegator, to, delegation_block_height, delegation_block_index, TRANSFER)`


#### RecordVote(byets txid)
> Access: external
> Records a vote transaction taht was sent on Ethereum
> RecordDelegation may only be called during the vote receording period.

* Check the that vote recording period did not end
  * `ethereum_block_height` < `election_block_height` + `parameter.VOTE_RECORDING_PERIOD_LENGTH_IN_BLOCKS`.
    * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`.
    * If not return an error indicating: Resubmit in the next election vote receording period.
* Read the transction's `Vote(voter, nodes_list, vote_counter)` event emitted by the `OrbsVoring` contract.
  * If no was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next election vote receording period.
  * `vote_block_height` = `ethereum_block_height`.
  * `vote_tx_index` = `ethereum_tx_index`.
* Check the current `voter`'s vote:
  * if the current voter_last_update[`delegator`].{`block_height`, `tx_index`} is larger than the {`vote_block_height`,`vote_tx_index`} fail the transaction.
* Update `voter`'s list and map:
  * voter_last_update[`voter`] = {`delegation_block_height`, `delegation_tx_index`}
  * Set vote[`voter`] = `nodes_list`.


#### ProcessVoting() : bool
> Access: external
> Process the election results, updates the nodes registery and the next election data.
* State:
  * `election_ethereum_block_height` (init: TBD)
  * `apply_results_timestamp` (init: 0)
  * `current voter`

* Parameters:
  * `ELECTION_CYCLE_IN_BLOCKS`
  * `VOTE_RECORDING_PERIOD_LENGTH_IN_BLOCKS`
  * `APPLY_RESULTS_LENGTH_IN_SEC`

* Check that the vote recording period has ended
  * `ethereum_block_height` > `election_block_height` + `parameter.VOTE_RECORDING_PERIOD_LENGTH_IN_BLOCKS`
    * Read the reference final `ethereum_block_height` using `Ethereum.GetBlockHeight`.
    * If not return an error indicating: recording period has not ended.
* Check not `prcoessing_done`.
  * If `prcoessing_done` return TRUE
* If not `voters_done`:
  * Check voter_last_update[`current_voter`] has not expired (< `ethereum_block_height` - `parameter.VOTING_VALIDITY_TIME`)
  * stake[`current_voter`] = Ethereum.balanceOf(current_voter).
  * 
  * current_voter = current_voter.next 
  * 

* If ..
  * Call `ProcessVotes()`.
  * Call `NodesRegistery.UpdateElectionResult`
  * Set `election_ethereum_block_height` = `election_ethereum_block_height` + `ELECTION_CYCLE_IN_BLOCKS`
  * 

* If `ethereum_block_height` > `election_ethereum_block_height` + `VOTE_RECORDING_PERIOD_LENGTH_IN_BLOCKS`
  * Set `vote_recording_period` = `FALSE`
  * Call `CalculateVotes()`
  * Call `CalculateElectedNodes`
  * Set `apply_results_block_height` = current.block_height 


#### CalculateVotes()
> Calculates the active nodes based on the votes:
* data structures in state:
  * total_voted_stake[node] = uint
  * voting_weight[address] = uint

&nbsp;
* Remove active voters from delegator's list:
  * For every `voter` in voters list:
    * `current_delegatee` = delegatee[`voter`]
    * Remove `voter` from delegators[`current_delegatee`]

* Recursively set the voting_weight of each `voter` or `delegator`:
  * If the participant has not delegators that delegated to it then voting_weight = stake[participant]
  * Else voting_weight = sum(delegators voting_weight + stake[participant])

* For (every `voter` in voters list):
  * if number of nodes in vote[`voter`] > parameter.VOTES_PER_TOKEN 
    * weight = voting_weight[`voter`] * parameter.VOTES_PER_TOKEN / (number of nodes in vote[`voter`])
  * Else
    * weight = voting_weight[`voter`]
  * For every node in vote[`voter`]:
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

#### UpdateElectionResult(block_height, nodes_list)
> Stores the list of elected nodes per block height
* Nodes[block_height] = nodes_list
* NodesUpdateBlockHeight[index] = block_height

#### GetElectedNodesByHeight(block_height) : nodes_list
> Return the list of elected nodes per requested block_height
* Find the `update_block_height` = largest block in NodesUpdateBlockHeight that is smaller than the provided block_height.
* Return Nodes[`update_block_height`]

#### GetNumberOfNodeUpdates : uint
> Returns the size of the NodesUpdateBlockHeight list, used for migrations.

#### GetElectedNodesByIndex : (block_height, nodes_list)
> Returns the list of elected nodes and reelvant block_height by index, used for migrations.

