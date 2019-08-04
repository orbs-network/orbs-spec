# `_Committee` System Service

The system service charged with maintaining validators reputation (based on accepted block proposals) and for obtaining an ordered list of validators - committee.\
Implemented on the `Native` processor.\
Note: Committee members coincide with elected validators (obtained by calling `_Elections.getElectedValidatorsOrbsAddress()`).\

#### Permissions
* This is a system contract that runs under `System` permissions.


#### State
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
> An accepted block proposal initializes the proposer's reputation.

#### Permissions
* `External` (caller can only be system.contract call).
* `ReadWrite` (might change state).

#### Behavior
* OrderedCommittee = Get ordered list of validators by calling `GetCommittee()`.
* BlockProposer = Get block proposer by calling `Environment.GetBlockProposer`.
* Reset the block proposer `MembersReputation.MissedBlocksCounter`.
* Iterate over OrderedCommittee until BlockProposer found:
    * For all unsuccessful proposals - increase the missed blocks `MembersReputation.MissedBlocksCounter`.



## `GetCommittee` (method)
> Note: 
#### Permissions
* `External` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* ElectedValidators = Get list of elected validators by calling `_Elections.getElectedValidatorsOrbsAddress()`.
* Order the nodes based on the weighted random sorting algorithm (reputation is taken into account here).
* Set each validator reputation :
    * Rate = min(Tolerance - missedBlocksCounter, 0)
    * Rate = max(-ReputationBottomCap, Rate)
    * Weight = 2^(Rate)
* Sort the ElectedValidators according to reputation: 
    * Use Sha(ElectedValidatorAddress, BlockHeight)
    * multiply by Weight
    
Return the sorted validators list.
