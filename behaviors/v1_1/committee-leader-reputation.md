# Committee with Leader Reputation
Provide an ordered committee based on reputation - 
Initial reputation targets tolerance against malfunctioning participant (validator) by reducing its probability to be elected as leader.

Manage reputation using a smart contract. This state based reputation is consumed upon retrieval of an ordered committee - the query is performed by the ConsesnsuContext service 
to provide the ConsensusAlgo an ordered committee for consensus round setup at a given height.

Reputation state update is based on [triggers-blockPost](../flows/trigger.md) mechanism. 
The [Reputation](../smart-contracts/system/_Reputation.md) system contract is deployed on each virtual chain.
<!--Note: Audit node - verification different than participant?-->
## Flow

#### Participant Services
* `ConsensusContext` - Gets the ordered committee - list of validators - from state.
* `VirtualMachine` - Expose the block proposer in execution context - for state update, and retrieve ordered committee from state. 

#### RequestOrderingCommittee 
* At the beginning of each consensus round (block height), the node queries the `ConsensusContext` for a sorted list of nodes (ordered committee). \
The `ConsensusContext` retrieves an ordered validators list from state by calling the `_Committee.getOrderedCommittee` (using the `VirtualMachine`).

#### Update validators reputation
* Default reputation provides a measure for validator activity as a leader.
* The block proposer (appearing in the block header), if successful will resume the top reputation value.\
All unsuccessful block proposers will gain one miss.
* The state is recorded in the [Committee](../smart-contracts/system/_Committee.md) system contract. \
* The update of reputation is triggered by the `triggers.blockPost` call. 

## SDK update
#### Block.GetBlockProposer
* Returns the block proposer address.

## ConsensusService update
* Note: Block proposer provided according to block (in leader change the block might not be replaced).

#### RequestBlockProposal
* Provide the block proposer address in method call.

#### ValidateBlockProposal
* Provide the block proposer address in method call.