# Contracts Functionality Overview

### contractsRegistry
The contractRegistry manages the PoS contracts list and active managers. The contract is the basis of the system and allows any contract or user to retrieve the up-to-date contract lists. The contractRegistry pushes contract address updates to the PoS managed contracts. The contract registry is the core component of the PoS contracts upgrade and allows to transition to a new contract for each contract in the system, as well as the registry itself.

### stakingContract
The staking contact holds the participants staked tokens. The staking contract notifies the stakingContractHandler on any change in a participant stake, for example upon a stake or unstake actions. 
See: [Orbs Staking Contract High-Level Specification](https://github.com/orbs-network/orbs-staking-contract/blob/master/docs/CONTRACT.md)

### stakingContractHandler
The staking contract handler connects the staking contract to the Orbs PoS contracts and to Orbs contracts governance, managed by the contractRegistry. The stakingContractHandler provides an infrastructure for connecting an additional staking contract or isolating the staking contract actions from the rest of the PoS contracts if needed.

### delegations
The delegations contract manages the delegation database for each participant in Orbs PoS. The delegation contract holds for participant his selected guardian and the participant’s self-stake as well as the stake delegated to the participant. 
On each change in a participant’s delegation, self stake or delegated stake, the contract notifies the stakingRewards contract that calculates the rewards accordingly. In addition, the delegations contract notifies the elections contrac on a participant delegated stake change, masking the delegators notation from the election contract. 

### elections
The election contract connects the delegation contract with the committee, guardians registration and certification contracts. In addition, the election contract implements the voting logic both for voteUnready and voteOut manages the node state machine notification readyForCommittee and readyToSynct. Upon a notification from the delegation contract the elections contract calculates the participant's new effective stake and updates the committee contract. Upon a guardian’s readyForCommittee, the committee, the election contract validates that the guardian is registered and not voted-out and requests the committee contract to add it. If a guardian was voted-unready or voted-out, the elections contract requests the committee contract to remove it from the committee.

### committee
The committee contract manages the current committee state. The committee contract gold the current committee members and their weights. Upon an effective stake change notice from the election contract, the committee contract updates the committee member weight and emits an update event. On a request to join the committee, the committee member checks that candidate is qualified to join, if the candidate weight is higher than the committee member with the lowest weight, the candidate will join the committee instead of the minimal weight guardian. Upon a change in the committee members list or their certification, the contact notifies the stakingRewards and FeesAndBootstrap contracts on the leaving and joining members to update their rewards state.

### stakingRewards
The stakingRewards contract manages the staking rewards state of the participants in the PoS ecosystem. The staking rewards architecture is based on the rewards per share architecture common in framing pools. The contract architecture was designed for Orbs PoS architecture which is based on 3 levels: global allocation, committee guardians receiving their share based on their weight, and delegators receiving a portion of their guardian rewards based on their stake. The staking rewards contract is updated on two events: an update to a delegator stake, and a change in guardian’s committee membership. A participant may claim his staking rewards that are staked in the staking contract and the system state is updated accordingly. The staking withdraws funds from the stakingRewardsWallet holds tokens up to the total unclaimed amount for all participants.

### feesAndBootstrapRewards
The feesAndBootstrapRewards contract manages the fees and bootstrap rewards state of the guardians. The fees and bootstrap rewards architecture is similar to the architecture of the stakingRewards contract. The contract architecture is based on 2 levels: global allocation, committee guardians. The feesAndBootstrapRewards collects fees from two instances of fees wallets the certificated and general virtual chains fees. In addition the feesAndBootstrapRewards withdraws bootstrap rewards from the bootstrap wallet. The may hold tokens up to the total unclaimed amount for all guardians.

### protocol wallets: stakingRewardsWallet and bootstrapRewardsWallet
The stakingRewardsWallet and bootstrapRewardsWallet hold the protocol fund allowing the relevant rewards contract to withdraw funds up to the maximal rate that was set. The protocol wallets provide a simple thus secure mechanism to store the protocol funds. The protocol wallets are disconnected from the PoS contracts governance and management infrastructure and are managed by two types of administrators: a functional administrator that sets the wallet’s client and a migration administrator that controls the withdrawal of funds.

### Fees wallets: generalFeesWallet and certifiedFeesWallet
The fees wallets maintain the subscription fees paid by virtual chains. Two separate instances of fees wallets are deployed one for general virtual chains fees and one for certified virtual chains fees. The fees wallets store the subscription fees in 30 days buckets, to provide an accurate division of the fees for each time period. The two fees wallets client, the feesAndBootstrapRewards, collect the paid fees for the time period since the last collect. 

### guardiansRegistration
The guardiansRegistration contract stores the guardians registration data. And allow guardians to register, modify their data and unregister. In addition the contact stores a metadata map allowing each guardian to store general purpose keys to be queried by the Orbs platform. For example, a link to the guardian identification data may be stored at the "ID_FORM_URL" metadata key.
The guardiansRegistration provides mapping functions, such as from Orbs addresses to Ethereum addresses, used by other contracts to allow function calls using the Orbs address. The guardiansRegistration contract notifies the election contract on changes in guardians registration status.

### certification
The certification contract stores the certification status for each Guardian. The certification data is managed by the certificationManager that may set and clear a guardian’s certification. The certification contract notifies the election contract on any change in a guardian certification.

### subscriptions
The subsceription contract manages the virtual chains’ subscription status. The contract allows to create a virtual chain, modify its properties including metadata properties used by the Orbs platform and extend to a virtual chain subscription. Virtual chain creation and subscription extension are not done by directly interacting with the subscriptions contract but rather by calling the subscription plan contract that updates the subscriptions contract. The subscriptions contact holds a list of valid subscription plan contracts.

### Subscription plan
The subscription plan contract is a stateless contract, responsible for the plan and fees structure of a virtual chain subscription. Virtual chains owners call the subscription plan contract to create a virtual chain and extend its subscription. The subscription plan contracts assigns a tier (plan) and the monthly fees rate based on the contract plan. Multiple subscription contracts may be deployed in the system.

### protocol
The protocol contact is responsible to set and synchronize Orbs platform protocol version upgrades. The protocol contact supports multiple deployment subscets, such as: production and canary, each may operate on the different protocol version. Upon protocol version upgrade a future time is set as the cutoff date for the protocol upgrade.
