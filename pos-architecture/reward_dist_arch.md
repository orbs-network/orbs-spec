# Reward Distribution Architecture

The rewards and fee distribution logic (implemented in StakingRewards.sol and FeesAndBootstrapRewards.sol) may appear as the most complex logic implemented in of the Orbs PoS contracts. However, the architecture (which is based on defi farming pools) is structured around one key idea - the rewards-per-share technique for efficient pro-rata distriubution of funds. It owuld be much easier to unerstand these contracts once this technique is understood.

But before getting into that, let's start with a quick summary of the problem - the reward distribution model. 
- The PoS echosystem has two main roles: guardians, and delegators.
- Each delegator delegates to a single guardian.
- Guardians are also delegators (they are self-delegating).
- The protocol assignes rewards to each elected guardian (committee member), which in turn distributes a portion of the rewards to each of her delegators.
- All distbituions are pro-rata, according to the corresponding weight of a guardian compared to the rest (for guardians), and according the the corresponding weight of the delegator compared to other delegators of the same guardian (for delegators).
- The portion of the guardian rewards that is assigned to the delegators is configurable, with the default set to 66.67%.

Example:
Say we have 2 elected guardians: A (with 20% of the committee weight) and B (80%).
A has the following delegators: dA_1 (stake 30), dA_2 (stake 70).
And B: dB_1 (stake 40), dA_2 (stake 60) - delegated to B.

Assuming in a given period, the procotol assigned a total rewards of 1000 Orbs to all guardians and delegators.
Also let's assume a 66% guardian cut, instead of 66.67% for simplicity.

The distirubion would then be as follows:
- A holds 20% of the committee weight, so A and her delegators will all receive a total of 20% * 1000 = 200 Orbs combined.
  - A takes a guardian cut of 34% * 200 = 68 Orbs
  - The remaining 132 Orbs are distributed to A's delegators.
  - dA_1 hold 30% of the total stake delegated to A, so she gets 30% * 132 = 39.6 Orbs.
  - dA_2 hold 70% of the total stake delegated to A, so she gets 70% * 132 = 92.4 Orbs.
- B holds 20% of the committee weight, so B and her delegators will all receive a total of 80% * 1000 = 800 Orbs combined.
  - B takes a guardian cut of 34% * 800 = 272 Orbs
  - The remaining 528 Orbs are distributed to A's delegators.
  - dB_1 hold 40% of the total stake delegated to B, so she gets 40% * 528 = 211.2 Orbs.
  - dB_2 hold 60% of the total stake delegated to B, so she gets 60% * 528 = 316.8 Orbs.
  
(Note - in practice, A and B are also delegators (with a self delegation), and would also be rewarded just as any other delegator, in additoin to their guardian cut)

So, how would we implement this distribution logic? A naive approach would be to:
- Keep track on the PoS state (delegations, weights, etc.)
- On any change, calculate the total rewards to be distributed for the time period since the last change (1000 in our example).
- Distribute the rewards pro-rata to all guardians and delegators, as seen in the example above.
The problem with this method is that it's very inefficient when changes are frequent, and when the amount of participants is large.

The rewards-per-share technique allows to implement this disribution logic efficiently. 
Let's start by explaining how we use this technique to implement the guardian rewards:
- We keep a global variable named `stakingRewardsPerWeight`. It holds the total rewards that would have been distributed to a theoretical elected guardian (and her delegators) with an absolute weight of 1, since the begining of time.
  - This variable is updated on every change (stake increase, delegation change, etc.)
  - It monotonically increases with time.
- For each guardian, we keep a variable names `lastStakingRewardsPerToken`. It holds the value of `stakingRewardsPerToken` that was during the last guardian weight change.













  
  





