# `_Committee` System Service

The system service charged with maintaining validators reputation (based on accepted block proposals) and for obtaining an ordered list of validators - committee.\
Implemented on the `Native` processor.\
Note: Committee members coincide with elected validators (obtained by calling `_Elections.getElectedValidatorsOrbsAddress()`) at the beginning of a new consensus round. Upon election event, the newly elected validators (during block execution) might differ. Thus, Reputation update should run before election result is reflected.

#### Permissions
* This is a system contract that runs under `System` permissions.


#### State
> Maintain a counter of unsuccessful block proposals for each committee member.\
> `ToleranceLevel` aims to smooth effects not caused directly by the block proposer - "missed blocks measure" is dependent on consensus agreement (takes into account other committee members).
> The tolerance level is currently high (4) to account for "empty block mechanism", which relies on switching block proposers during a long empty block duration. \
> `ReputationBottomCap` aims to optimize the trade-off between the damage to the system throughput, inflicted by a malfunctioning block proposer, and the damage to the block proposer - allowing it get back on track in a reasonable amount of time. \
> Currently, these values are hardcoded, could be further optimized and in the future perhaps support update. 
* Constants:
    * ToleranceLevel (4)
    * ReputationBottomCap (10)
* State variables:
    * MembersReputation (map):
        * MemberID (OrbsAddress)
        * MissedBlocksCounter (0)

&nbsp;
## `_init` (method)

#### Permissions
* `Internal` (caller must be the same service).
* `ReadWrite` (does not changes state).

#### Behavior
* Empty.



## `UpdateReputation` (method)
> Dependent on `_Elections` system contract. \
> If no such contract calls to its methods fail no state updates take place. \
> An accepted block proposal initializes the proposer's reputation.\
> Support calls only by "protocol special" empty signature which is reserved for system.

#### Permissions
* `External` (caller can only be system.contract call).
* `ReadWrite` (might change state).

#### Behavior
* OrderedCommittee = Get ordered list of validators by calling `getOrderedCommittee()`.
* BlockProposer = Get block proposer by calling `Environment.GetBlockProposer`.
* Reset the block proposer `MembersReputation.MissedBlocksCounter`.
* Iterate over OrderedCommittee until BlockProposer found:
    * For all unsuccessful proposals - increase the missed blocks `MembersReputation.MissedBlocksCounter`.



## `getOrderedCommittee` (method)
> Note: 
#### Permissions
* `External & Internal` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* ElectedValidators = Get list of elected validators by calling `_Elections.getElectedValidatorsOrbsAddress()`.
* Set the rating of each committee member:
    * Get missedBlocksCounter from state.MembersReputation(map). If no such member, set missedBlocksCounter to zero.
    * Rate = min(Tolerance - missedBlocksCounter, 0)
    * Rate = max(-ReputationBottomCap, Rate)
* Sort the ElectedValidators according to Rate: 
    * Use Sha(ElectedValidatorAddress, BlockHeight)
    * Multiply by 2^(Rate)
    
Return the sorted validators list.
