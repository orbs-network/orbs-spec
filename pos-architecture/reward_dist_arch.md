# Reward Distribution Architecture

The rewards and fee distribution logic (implemented in StakingRewards.sol and FeesAndBootstrapRewards.sol) may appear to be the most complex part of the Orbs PoS contracts.
However, the architecture (which is based on defi farming pools) is structured around one key idea - the rewards-per-share technique for efficient pro-rata distriubution of funds.

But before getting into that - let's start with a quick summary of the problem - the reward distribution model. 
- The PoS echosystem has two main roles: guardians, and delegators.
- Each delegator delegates to a single guardian.
- Guardians are also delegators (they are self-delegating).
- Both hold stake in the system, which defines their voting power and their share in the rewards given by the protocol.
- The protocol assignes rewards to each elected guardian (committee member), which in turn distributes a portion of the rewards to each of her delegators.
- All distbituions are pro-rata, according to the corresponding weight of a guardian compared to the rest (for guardians), and according the the corresponding weight of the delegator compared to other delegators of the same guardian (for delegators).
- The portion of the guardian rewards that is assigned to the delegators is configurable, with the default set to 66%.

Example:
Say we have 2 elected guardians: A (with 20% of the committee weight) and B (80%).
A has the following delegators: dA_1 (stake 30), dA_2 (stake 70).
And B: dB_1 (stake 40), dA_2 (stake 60) - delegated to B.


Assuming in a given period, the procotol assigned a total rewards of 1000 Orbs to all guardians and delegators.
The distirubion would than be as follows:
- A holds 20% of the committee weight, so A and her delegators will all receive 20% * 1000 = 200 Orbs in total.
  - A takes a guardian cut of 34% * 200 = 68 Orbs
  - The remaining 132 Orbs are distributed to A's delegators.
  - dA_1 hold 30% of the total stake delegated to A, so she gets 30% * 132 = 39.6 Orbs.
  - dA_2 hold 70% of the total stake delegated to A, so she gets 70% * 132 = 92.4 Orbs.
  
  





