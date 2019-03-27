# Voting Flow

The flow describes the voting for Guardians and Validators in Orbs.

## Participants in this flow
* Validator Candidate
  * A node that registered
* Validator Nominee
  * A Validator Candidate that was approved and participates in the election.
* Elected Validators
  * Validators elected for next election term.
* Guardian
  * An account registered to a Guardian 
  * Actively participates in Orbs voting.
* Delegator
  * An Orbs account that selected an active voter or another account (agent) to act on its behalf.
* Agent
  * An Orbs account that another account delegated its voting to.

![alt text][hierarchical_voting] <br/><br/>

[hierarchical_voting]: ../_img/hierarchical_voting.png "hierarchical voting"

## High level Election Flow

### Registration
> Performed on Ethereum using the registration smart contracts
 
#### Validators Registration
* A Validator candidate registers to the Validators Registry contract and added to the candidates list.
* The Validator is approved and added to the Validators Contract.

#### Guardians Registration
* A Guardian registers tp the Guardians Registry contract.
* Once registered a Guardian vote is considered and the Guardian delegation is ignored (implicitly delegates to herself)

### Delegation and Voting 
> Performed on Ethereum using the voting smart contract

#### Delegation
* A Delegator (token holder) delegates her voting weight (stake) to a Guardian or another Delegator.
  * A delegation is valid until changed.
  * A delegation order is determined by the block_height and block_index the delegation was performed at.
* A Delegator can not delegate to the its own address (loop)
  
#### Voting
* A Guardian votes out up to 3 validators.
  * An empty list indicates no validator to vote out.
* A vote can be cast at any time and is valid for 40320 Ethereum blocks (~ 7 days)

### Delegation and Voting Mirroring
> Performed on Orbs during until the end of the Vote Mirroring Period: 480 blocks (~ 2 hours) after the election event.

### Elections
* All election data is calculated at the election event block.
* The Validator nominees for the election are the registered and approved validators at the time of the election.
* The amount of votes (vote out) for each Validator nominee are calculated
* A Validator that receives more than 70% votes is removed from the election nominees.
* The Validators the remaining validators (approved) are elected for the next voting term.

#### Top 22 Validators with Gradual Changes (Not implemented)
* If the number of approved validators is less than 22 (maximum number of Validators)
  * All the approved Validators that were elected for the previous term are elected.
  * The approved new nominee with the most stake is elected
* Else 
  * The top 21 approved Validators that were elected for the previous term with the most stake are elected
  * The 22th Validator is compared with the approved new nominee with the most stake, the one with the higher stake is elected.

#### Removal of voted out Validators (Not implemented, performed manually)
* A validator that is voted out by the Guardians in all elections for over a week is removed from the candidates list.
  * Initially done by manually (as a proxy for the Guardians)

### Elections Schedule
* Elections occur periodically, initially every 17280 blocks (~ 3 days).
  * The first election occurs at the `FIRST_ELECTION_BLOCK`.
* Once the Vote Mirroring Period is complete, the election results can be processed.
* The results processing consist of:
  * Reading Ethereum data at the time of the elections (stake, delegation, vote, registered Guardians and Validators)
  * Calculating the Guardians voting weight
  * Calculating the elected Validators
  * Storing the elected Validators

## Ethereum and Orbs interfaces

#### Delegating votes
* Every Delegator (token holder) may delegate its stake to an agent (address).
  * The agent may be an Guardian or another Delegator.
* Delegation may be performed by 2 means:
  * Sending 7 Orbs-satoshi (7x10<sup>-18</sup> Orbs) to an address.
    * Generates a `Transfer(from, to, tokens = 7)` event.
  * Delegating by sending a transaction to Ethereum.OrbsVoting.delegate(address).
    * Generates a `Delegate(from, to, delegateCounter)` event.
* Delegation does not expire and requires delegation to another agent in order to change it.
  * Delegation to address 0 cancels the delegation.
* Once a delegation was performed by OrbsVoting contract it takes precedence over sending Orbs-stoshi (regardless of which action was performed later).
* Note: the stake used for the voting is the Delegator's stake at the election time. (and not the stake at the time of the delegation).

#### Voting out Validators (by Guardians)
* A Guardian can vote by sending a transaction to Ethereum.OrbsVoting.voteOut(address[]).
  * Generates a `VoteOut(voter, validators, voteCounter)` event.
* A Guardian must be registered at the time of the election in order for its vote to be counted.
* A vote remains valid for 45,500 blocks (~ 7 days)
* The voting weight of an Guardian is proportional to the total stake that was deleted to it. (hierarchical delegation)
* A Guardian can vote for the disapproval of up to 3 Validators. (may vote for none)

## Election Flow
![alt text][election_flow] <br/><br/>

[election_flow]: ../_img/election_flow.png "election flow"

* Election event
  * Election is performed every 20,000 Ethereum blocks
  * Votes and delegations that were done up to the block_height (inclusive) take effect.
* Votes Mirroring Period
  * Starts after the election event for 545 blocks (~2 hours).
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

#### Recording voting data on Orbs
* Done by an off-chain application that monitors Ethereum's blocks.
  * Sends corresponding transactions (generates the relevant events) to `Orbs.VoteProcessing.recordVote(eth_txid)`
    * Sent during the `Votes Recording Period`.
  * Anyone may send the transactions.
  * A transaction is considered final based on the Ethereum connector finality parameter (default 100).

#### Processing()
* Can be initiated by anyone upon completion of the votes mirroring period.
  * Span over a set of transactions (where the last one indicates completion)
* Reads the stake for each Delegator, Guardian and Validator at the time of the election block.
* Eliminates expired votes.
* Calculates the delegated stake for each Guardian - the total stake of the delegation tree that points to it.
  * Note: once an Guardian has registered, it's delegation is ignored.
* Record the election results in the ValidatorsConfig along with the transition Orbs block (A period of time after the last processing block).

## Rewards calculation

#### Participation reward (Delegators and Guardians)
* Maximum reward per election:
  * 60M x 20000 x 13.28 / 31536000 = 505328 ORBS
* Participation reward for the election
  * Min(493150, 8% of total voting stake)
* Divide the participation reward for the election in proportion to the Delegator/Guardian stake.

#### Guardians Excellence Program reward
* Maximum reward per election:
  * 40M x 20000 x 13.28 / 31536000 = 336885 ORBS
* Guardians reward for the election
  * Min(493150, 10% of voting stake of the top 10 Guardians)
    * The voting stake is the total delegated stake including the Guardian's 
* Divide the reward for the election to the top 10 Guardians in proportion to their stake.

#### Validators reward
* Every elected Validator is awarded with 4% of the Validator stake
* 1M ORBS as part of the Validator Introduction Program.
  * Validator Introduction Program reward per election:
    * 1M x 20000 x 13.28 / 31536000 = 8423 ORBS

&nbsp;
## Contracts specification

[Ethereum voting contracts](../smart-contracts/ethereum-contracts/voting.md)

[Orbs voting contracts](../smart-contracts/orbs-system-contracts/voting.md)
