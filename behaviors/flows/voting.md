# Voting Flow

The flow describes the voting for Guardians and Validators in Orbs.

## High level Election Flow
* Once registered, a Validator undergoes a one month qualification period. - `TODO`
* Once qualification is completed, a Validator is added to the candidates list.
* Guardians vote to disapprove up to 3 validators. 
  * Guardians vote can be cast at any time and is valid for a week.
* Elections occur periodically, initially every 3 days.
* A validator is disapproved if it received 70% or more disapproval votes in the election
* If there are N=22 approved Validators that were elected for the previous term:
  * The top 21 approved Validators that were elected for the previous term with the most stake are elected
  * The 22th Validator is compared with the approved new candidate with the most stake, the one with the higher stake is elected.
* Else
  * All the approved Validators that were elected for the previous term are elected.
  * The approved new candidate with the most stake is elected
* A validator that is disapproved by the Guardians in all elections during for over 2 weeks is removed from the candidates list. A validator may re-register, requiring it to redo the qualification period.
  * Initially done by manually (as a proxy for the Guardians)

## Participants in this flow
* Validators (candidates)
  * A set of nodes that underwent due-diligence.
* Elected Validators
  * A Validator candidate that received enough votes.
* Guardian
  * An account that actively participates in Orbs voting.
  * An account is considered an active voter for a period of ACTIVE_VOTING_PERIOD days after voting.
  * Option to require active voters to undergo due-diligence.
* Stakeholder
  * An Orbs account / stakeholder that selected an active voter or another account (delegatee) to act on its behalf.
* Agent
  * An Orbs account that another account delegated its voting to.

## Voting and Delegation

![alt text][hierarchical_voting] <br/><br/>

[hierarchical_voting]: ../_img/hierarchical_voting.png "hierarchical voting"

#### Delegating votes
* Every stakeholder (address) may delegate its stake to an agent (address).
  * The agent may be an Guardian or another stakeholder.
* Delegation may be performed by 2 means:
  * Sending 7 Orbs-satoshi (7x10<sup>-18</sup> Orbs) to an address.
    * Generates a `Transfer(from, to, tokens = 7)` event.
  * Delegating by sending a transaction to Ethereum.OrbsVoting.delegate(address).
    * Generates a `Delegate(from, to)` event.
* Delegation does not expire and requires delegation to another agent in order to change it.
  * Delegation to address 0 cancels the delegation.
* Once a delegation was performed by OrbsVoting contract it takes precedence over sending Orbs-stoshi (regardless of which action was performed later).
* Note: the stake used for the voting is the stakeholder's stake at the election time. (and not the stake at the time of the delegation).

#### Voting to Validators (Disapproval by Guardians)
* A Guardian can vote by sending a transaction to Ethereum.OrbsVoting.vote(address[]).
  * A Guardian must be registered in order to vote, and must be still registered at the time of the election in order for its vote to be counted.
* A vote remains valid for 7 days ~ 40320 blocks (parameter.VOTING_VALIDITY_TIME)
* The voting weight of an Guardian is proportional to the total stake that was deleted to it. (hierarchical delegation)
* A Guardian can vote for the disapproval of up to 3 Validators. (may vote for none)

## Election Flow
![alt text][election_flow] <br/><br/>

[election_flow]: ../_img/election_flow.png "election flow"

* Election event
  * Election is performed every X Ethereum blocks, based on the state in block X.
    * Default: ~3 days (17280), to be reduced at a later stage.
  * Votes and delegations that were done up to the block_height (inclusive) take effect.
* Votes Mirroring Period
  * Starts after the election event for ~2 hours (480 blocks).
  * During the mirroring period, the delegations and votes are recorded on Orbs.
* Votes Processing
  * Performed once the mirroring period is complete.
  * Initiated by a set of transactions.
  * Record and stake, process and updates the elected Validators list.
* Transition period
  * The period between the end of the processing and election results and the commencement.
* Apply new elected Validators.

## Recording the data on Orbs

#### Recording delegation data on Orbs
* Done by an off-chain application that monitors Ethereum's blocks.
  * Sends corresponding transactions (generates the relevant events) to `Orbs.VoteProcessing.recordDelegate(txid)` and `Orbs.VoteProcessing.recordDelegateByTransfer(txid)`
    * Sent before the end of the `Votes Mirroring Period`.
  * Anyone may send the transactions.
  * A transaction is considered final based on the Ethereum connector finality parameter (default 100).
* The VoteProcessing contract reads the log and stores:
  * Stakeholders list
  * Stakeholder -> Agent map.
  * Agent -> Stakeholders map.

#### Recording voting data on Orbs
* Done by an off-chain application that monitors Ethereum's blocks.
  * Sends corresponding transactions (generates the relevant events) to `Orbs.VoteProcessing.recordVote(eth_txid)`
    * Sent during the `Votes Recording Period`.
  * Anyone may send the transactions.
  * A transaction is considered final based on the Ethereum connector finality parameter (default 100).
* The VoteProcessing contract reads the log and corresponding balance and stores:
  * Voters list 
  * Voter -> Validators map.

## Processing()
* Can be initiated by anyone upon completion of the votes mirroring period.
  * Span over a set of transactions (where the last one indicates completion)
* Reads the stake for each stakeholder and Guardian at the time of the election block.
* Eliminates expired votes.
* Calculates the delegated stake for each Guardian - the total stake of the delegation tree that points to it.
  * Note: once an Guardian has registered, it's delegation is ignored.
* Record the election results in the ValidatorsConfig along with the transition Orbs block (A period of time after the last processing block).

#### Rewards calculation
> TBD

&nbsp;
## Contracts specification

[Ethereum voting contracts](../smart-contracts/ethereum-contracts/voting.md)

[Orbs voting contracts](../smart-contracts/orbs-system-contracts/voting.md)

&nbsp;
## Specification

### SDK

#### `GetEthereumBlockHeight`
* Returns the block_height of the first block that meets the finality requirements.
* If the height is not within the finality parameter, return error.

#### `EthereumCallContract`
* Add ethereum_block_height to the input
  * The reference block_height to use. If the reference block height is above the finality block height - fail the transaction. 
  * If equal to 0, use the finality block height based on the timestamp.

#### `EthereumGetTransactionData`
> Returns data associated with a committed transaction

* Query the Ethereum node using the given arguments through IPC.
* Check that the transaction is within the finality parameter (based on the give time_stamp)
* Returns:
  * The committed transaction block_height and transaction_index.