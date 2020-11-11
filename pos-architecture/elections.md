# Orbs Elections
Orbs election processes is comprised of two phases: delegation to Guardians and election of Validators. Any token holder can delegate her voting weight (stake) to a Guardian to vote on her behalf. Orbs Validators are elected based on the Guardians votes and their stake. Validators are first approved by the Guardians that can vote out Validators that do not follow the protocol or misbehave. Next, the approved Validators are elected based on their ORBS stake. 

Orbs elections occur periodically every 20,000 Ethereum blocks (approximately 3 days). The election event occurs on a specific Ethereum block (election_block). The election results are calculated based on the participants ORBS stake at the election_block and the delegations and votes that were registered prior to it. 

## Voting and Delegation on Ethereum
Delegation and voting are performed over Ethereum smart contracts, allowing any individual to calculate and verify the election results. The use of Ethereum, a blockchain external to the network, enables to avoid conflict of interest, where the ones protecting the election’s probity are the ones being elected. 

## Token Holders Delegation
Token holders may choose to participate as Guardians or delegate their stake to a Guardian they trust. Delegators are encouraged to delegate to identified and reputable Guardians. The voting weight of each Delegator is equal to the number of ORBS tokens in their balance at the time of each election event. In order to participate as a Delegator, a token holder must have at least 10,000 ORBS tokens at the time of the election event. A Delegator may trust another Delegator with the selection of the Guardian, allowing the Guardian to receive the voting weight of the hierarchy of all Delegators behind it. A delegation may be modified at any time and persists unless modified.

## Guardians Voting
Guardians enforce the protocol by monitoring the network and vote to approve Validators accordingly. A Guardian that identifies Validators that do not follow the protocol can vote them out. In each vote, a Guardian can vote out up to 3 Validators, if all Validators operate correctly, a Guardian may approve all Validators by voting with an empty list. A Guardian may cast a vote at any time which remains valid for up to 45,500 Ethereum blocks (approximately one week). In order to participate in an election, a guardian needs to have a valid vote at the time of the election event.

The voting power of each Guardian is proportional to the number of ORBS tokens that are delegated to that Guardian, including the Guardian’s own tokens, at the time of the election event. Tokens can be delegated to Guardians either directly by their owner, or indirectly through a series of delegations. When voting to vote out multiple Validators, the voting weight for each Validator equals to the Guardian’s total voting weight.

### Guardians Registration
In order to participate in the network, Guardians are required to first register using a smart contract on Ethereum. As part of the registration, a Guardian provides a name and a website url, increasing trust by Delegators and other stakeholders. In addition, a Guardian is required to deposit 16 Ether as part of the registration. A Guardian may choose to unregister and receive the deposit back, a minimal participation period of two weeks is required from registration until a Guardian can unregister. A registered guardian implicitly delegates to herself, taking precedence over an explicit delegation.

### Validators Participation
Validators run the Orbs network, maintaining its security, availability and performance. Validators operate the virtual chains for all active applications and expose public interfaces that allow developers and clients to interact with the network. Validators are elected for election term based on the Guardians votes and their own stake. 

### Validators Registration
In order to participate in the network, Validators are required to first register using a smart contract on Ethereum. As part of the registration, a Validator provides name and website and addresses of her Orbs node.

During the launch period of the network and its initial set-up, Validators undergo a due diligence process in order to evaluate their technical abilities and enabling gradual rollout. This is to ensure the healthy functioning of the network in this critical period and to avoid the network interference by problematic actors during its launch. After a Validator candidate was approved she can connect to the network and then stand to be elected as active Validator.

## Election Results Calculation
On each election up to 22 Validators can be elected. First, the voting weight of each Guardian is calculated based on the balances and delegations state at the time of the election. Next, the vote-out votes for each Validators are counted, based on the Guardians’ vote at the time of the election. A Validator that received 70% or more of the total voting weight of the election is disqualified and can’t participate in the election. The remaining Validators are approved to be elected. Out of the approved Validators, top 22 are elected based on their own stake. A Validator that is consistently voted out by the Guardians in 3 consecutive elections will be removed from the candidates list and required to re-register in order to be elected.

### Paced Validator Additions
In order to prevent rapid changes in the active Validators, Validators change are performed gradually. Up to 3 Validators may be removed and one new Validator can be elected in each election. The top 21 Validators that were approved by the Guardians with the most stake are re-elected. Out of the remaining Validators approved by the Guardians, the Validator with the most stake is elected.

