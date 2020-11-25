# Continuous PoS State Update

Orbs PoS V2.5 architecture continuously updates the PoS state, there are no timely actions (every day, 3 days, etc) and the committee and rewards are continuously updated based don’t last action performed in the system. For example, when a delegator (delegated to a guardian in the committee) stakes additional tokens, the guardian's effective state, the committee weight and their rewards are immediately updated.

### Rewards Assignment
The rewards architecture is designed to constantly maintain a participant up-to-date rewards. At any point in time, a participant, delegator or guardian, may view and claim their current reward balance. The rewards architecture takes inspiration from the rewards per share architecture common in farming pools.
Upon every action that impacts the staking rewards, such as staking of tokens a delegator, 3 hierarchies of rewards are updates:
* Global rewards allocation - indicating allocation of rewards for the entire participants based on the time passed since the last allocation
* Guardian rewards assignment - assigns rewards to the guardian, based on the global rewards allocation and the guardian weight in the committee.
* Delegator rewards assignment - assigns rewards to the delegator, based on the guardian rewards assignment and the delegator stake.

The continuous rewards assignment architecture is designed in a scalable manner, supporting an infinite number of participants without increasing the gas costs of the transactions. A key part of the architecture is calculating the rewards assigned to a participant since his last assignment as needed upon a query. This implies that the current rewards balance may differ from the contact state and the query functions should be used to retrieve the up-to-date values.

The staking rewards updates were designed to minimize the gas cost of frequent actions such as stake and unstake or change in delegation. For each change that impacts specific participants: delegators and guardians, an event describing the change is emitted. Even though events are emitted only for the specific participants, one may use the entire set of the 3 level rewards assignment events to generate the continuous assignment over time graph for any participant.
