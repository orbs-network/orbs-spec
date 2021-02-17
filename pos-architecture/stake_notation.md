# Stake Notation in Orbs PoS Contracts 
Orbs PoS contacts and in particular the delegation contract maintain multiple variations for the a participant holds and the stake that was delegated to the participant.

The following table explains the meaning of the different values, the table targets advanced users that want to understand the code and internals of the contracts.


| Stake Property         | Description |
| ---------------------- | ----------- |
| selfStake              | The amount of ORBS tokens locked in the staking contract for an address | 
| delegatedStake | If self-delegating equals to the uncappedDelegatedStake else 0. |
| totalDelegatedStake | The sum of the delegatedStake of all participants <br> Note: the totalDelegatedStake may differ from the total stake staked in the staking contact, due to delegation to a non self-delegating address. |
| effictiveStake / committee weight | The weight of a guardian in the committee <br> Min(selfStake / minSelfStake, delegatedStake) <br> default: (8%) - indicates that the guardian as at least 8% of his effectiveStake |
| committeeTotalWeight | The sum of the weight of all the committee members |
