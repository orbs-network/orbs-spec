## Orbs Vote Processing Contracts
> Orbs voting is performed on the Etehreum platform and processed on Orbs as described in the [voting flow](../../flows/voting.md)).

&nbsp;
### VoteProcessing
Ownership: none

#### RecordDelegationData(delegator, to, delegation_block_height, delegation_tx_index, updated_by)
> Access: internal
> Process an Ethereum delegation transaction and logs it on Orbs along with the delegator stake.

* If `updated_by` = `VOTING_CONTRACT` and `updated_by` = `TRANSFER` fail the transaction.
* Read the `delegator`'s token balance by calling the ERC20 contract `balanceOf` using `ethereum_block_height` = `election_block_height`.
  * `delegator_stake` = uint64(balance / 1E18). (1 = 1 Orbs).
* Check the current account delegation:
  * if thre current update_index[`delegator`].`block_height`, `tx_index`} is larger than the {`delegation_block_height`,`delegation_tx_index`} fail the transaction.
* Update delegation map:
  * update_index[`delegator`] = {`delegation_block_height`, `delegation_tx_index`}
  * updated_type[`delegator`] = `updated_by`.
  * If delegatee[`to`] already exist (not empty)  
    * Remove the delegator from delegators[delegatee[`delegator`]].
  * Set delegatee[`delegator`] = `to`.
  * Add `delegator` to delegators[`to`].
  * Set stake[`delegator`] = `delegator_stake`.


#### RecordDelegation(bytes txid)
> Access: external
> Records an Ethereum delegation by vote contract transaction by calling `RecordDelegationData`.
> RecordDelegation may only be called during the vote receording period.

* Check the `vote_recording_period` = TRUE.
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

* Check the `vote_recording_period` = TRUE.
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
> Records an Ethereum delegation by transfer transaction by calling `RecordDelegationData`.
> RecordDelegation may only be called during the vote receording period.

* Check the `vote_recording_period` = TRUE.
    * If not return an error indicating: Resubmit in the next election vote receording period.
* Read the transction's `Vote(voter, nodes_list, vote_counter)` event emitted by the `OrbsVoring` contract.
  * If no was found fail the transaction.
    * If below the finality threshold, the event should not be returned by `GetTransactionLog`.
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
    * If above return an error indicating: Resubmit in the next election vote receording period.
  * Check that the vote has not expired 
    * if `ethereum_block_height` is below `election_block_height` - `parameter.VOTE_EXPERATION_BLOCKS`, fail the transaction with: expired vote.
  * Check that the `ethereum_block_height` is less or equal to the `election_block_height`.
  * `vote_block_height` = `ethereum_block_height`.
  * `vote_tx_index` = `ethereum_tx_index`.
* Read the `voter`'s token balance by calling the ERC20 contract `balanceOf` using `ethereum_block_height` = `election_block_height`.
  * `voter_stake` = uint64(balance / 1E18). (1 = 1 Orbs).
* Check the current voter vote:
  * if thre current update_index[`delegator`].{`block_height`, `tx_index`} is larger than the {`vote_block_height`,`vote_tx_index`} fail the transaction.
* Update `voter`'s list and map:
  * update_index[`delegator`] = {`delegation_block_height`, `delegation_tx_index`}
  * Set vote[`voter`] = `nodes_list`.
  * Set stake[`voter`] = `voter_stake`.


#### UpdateElectionPhase()
> Updates the current election phase
* State:
  * `election_ethereum_block_height` (init: TBD)
  * `apply_results_block_height` (init: 0)
  * `vote_recording_period` (init: 0)

* Parameters:
  * `ELECTION_CYCLE_IN_BLOCKS`
  * `VOTE_RECORDING_PERIOD_LENGTH_IN_BLOCKS`
  * `APPLY_RESULTS_LENGTH_IN_BLOCKS`

* Get the current Ethereum block height by calling `Ethereum.BlockHeight()`
* If `ethereum_block_height` > `election_ethereum_block_height`
  * Set `vote_recording_period` = `TRUE`.
  * Set `election_ethereum_block_height` = `election_ethereum_block_height` + `ELECTION_CYCLE_IN_BLOCKS`.
* If `ethereum_block_height` > `election_ethereum_block_height` + `VOTE_RECORDING_PERIOD_LENGTH_IN_BLOCKS`
  * Set `vote_recording_period` = `FALSE`.
  * Call `process_votes()`.
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
### NodesRecord
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

