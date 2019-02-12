# Voting Flow

The flow describes the voting for active delegates and nodes in Orbs.

## Participants in this flow
* Node candidates
  * A set of nodes that underwent due-diligence.
* Elected node
  * A node candidate that received enough votes.
* Active voter
  * An account that actively participates in Orbs voting.
  * An account is considered an active voter for a period of ACTIVE_VOTING_PERIOD days after voting.
  * Option to require active voters to undergo due-diligance.
* Delegator
  * An Orbs account / token holder that selected an active voter or anoher account (delegatee) to act on its behalf.
* Delegatee
  * An Orbs account that another account delegated its voting to.

## Voting and Delegation

![alt text][hierarchical_voting] <br/><br/>

[hierarchical_voting]: ../_img/hierarchical_voting.png "hierarchical voting"

#### Delegating votes
* Every account (delegator) may delegate its stake to an another account
  * Directly to an active voter or to another account.
* Delegation may be perfromed by 2 means:
  * Sending 7 Orbs-satoshi (7x10<sup>-18</sup> Orbs) to an address.
    * Generates a `Transfer(from, to, tokens = 7)` event.
  * Delegating by sending a transaction to Ethereum.OrbsVoring.delegate(address).
    * Generates a `Delegate(from, to)` event.
* Delegation does not expire and requires delegation to another address in order to change it.
  * Delegation to address 0 disables the delegation.
* Once a delegation was performed by OrbsVoring contract it takes precedence over sending Orbs-stoshi (regardless of which action was performed later).
* Note: the stake used for the voting is the delegator's stake at the election time. (and not the stake at the time of the delegation).

#### Voting to nodes (by active voters)
* An active delegate can vote by sending a transaction to Ethereum.OrbsVoring.vote([]address).
* Optional: require the voter to be part of the verified_voters list.
* A vote remains valid for 8 days (parameter.VOTING_VALIDITY_TIME)
* The voting weight of a delegate is proportional to the total stake that was deleted to it. (hierarchical delegation)
* Each delegate can vote to any numebr of nodes.
* Note: the voting weight is proportional to the number of voted nodes:
  * When voting to N<=5(parameter.VOTES_PER_TOKEN) nodes, each node receives the delegates total voting weight.
  * When voting to N>5 nodes, each node receives (5/N) of the weight.

## Election Flow
![alt text][election_flow] <br/><br/>

[election_flow]: ../_img/election_flow.png "election flow"

* Election event
  * Election is performed every X Ethereum blocks, based on the state in block X.
    * Default: ~24 hours (in blocks)
  * Votes and delegations that were done up to the block_height (inclusive) take effect.
* Votes Recording Period
  * Starts after the election trigger for ~2 hours (600 blocks).
  * During the recording period, the delegations, votes and stake are recorded on Orbs.
* Votes Proccesing
  * Performed once the recording is complete.
  * Initiated by a set of transactions.
  * Record and stake, process and updates the elected nodes list.
* Transition period
* Apply new elected nodes (inauguration).

## Recording the data on Orbs

#### Recording delegation data on Orbs
* Done by an off-chain application that monitors Ethereum's blocks.
  * Sends corresponding transactions (generates the relevant events) to Orbs.VoteProcessing.recordDelegation(eth_txid)
    * Sent during the `Votes Recording Period`.
  * Anyone may send the transactions.
  * A transaction is considered final based on the Ethereum connector finality parameter (default 100).
* The VoteProcessing contract reads the log and corresponding balance and stores:
  * Delegators indexed list
  * Delegator -> Delegatee map.
  * Delegatee -> Delegator map.

#### Recording voting data on Orbs
* Done by an off-chain application that monitors Ethereum's blocks.
  * Sends corresponding transactions (generates the relevant events) to Orbs.VoteProcessing.recordVote(eth_txid)
    * Sent during the `Votes Recording Period`.
  * Anyone may send the transactions.
  * A transaction is considered final based on the Ethereum connector finality parameter (default 100).
* The VoteProcessing contract reads the log and corresponding balance and stores:
  * Voters list 
  * Voter -> nodes map.

## Processing
* Initiated be any sender upon complition of the votes recording period.
  * Span over a set of transactions (where the last one indicates complition)
* Reads the stake for each delegator and voter at the election block.
* Eleminates expired votes/
* Calculates the delegated stake for each voter - the total stake of the delegation tree that points to it.
  * Note: once a voter has voted it's delegation is ignored.
* Record the election results in the nodes NodesRegistery along with the transition time in Orbs (A period of time after the last processing block).

#### Rewards calculation
> TBD

&nbsp;
## Contracts specification

[Ethereum voting contracts](../smart-contracts/ethereum-contracts/voting.md)

[Orbs voting contracts](../smart-contracts/orbs-system-contracts/voting.md)

&nbsp;
## Specification

### Consensus Context

#### `RequestOrderingCommittee` / `RequestValidationCommittee` (method)
> Returns a sorted list of nodes (public keys) that participate in the approval committee for the ordering of a given block height. Called by consensus algo.

* The current list of nodes is queried from the NodesRecord contract.
  * Call `VirtualMachine.ProcessQuery(query)`
    * block_height = current_block_height
    * signed_query
      * timestamp = block_timestamp
      * signer = ""
      * contract_name = `NodesRecord`
      * method_name = `GetElectedNodesByHeight`
      * input_argument_array = {`block_height`}
      * signature = ""

* If the size of requested committee is larger than total nodes, select all nodes as the committee.
* Order the nodes based on the weighted random sorting algorithm (reputation is taken into account here).

**Note**: once we support hash in the SDK, we can move also the random committee logic to a smart contract.

### VirtualMachine

####`ProcessQuery`
Add support for query on a specific (current) block height:
* Check state storage block height < block_height < state storage block height + `config.BLOCK_TRACKER_GRACE_DISTANCE`
  * Else panic.

### SDK

#### `GetEthereumBlockHeight`
* Returns the blockheight of the first block that meets the finality requirements.
* If the height is not within the finality parameter, return error.

#### `EthereumCallContract`
* Add ethereum_block_height to the input
  * The reference block_height to use. If the reference block height is above the finality block height - fail the transaction. 
  * If equal to 0, use the finality block height based on the timestamp.



