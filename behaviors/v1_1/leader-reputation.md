# Leader Reputation

Provide an ordered committee based on reputation - 
reducing the possibility of a malfunctioning participant (validator) to be elected as leader.

Manage reputation using a smart contract. This state based reputation is consumed by the ConsesnsuContext service 
upon providing the ordered committee to the consensus at height r.

Reputation state update is based on blockPost triggers mechanism. 

<!--Note: Audit node - verification different than participant?-->
## Flow

#### Participant Services

* `ConsensusContext` - Gets the validators reputation and creates an ordered committee. 
Note: when using separate services, consider merging the two calls
 "getElectedValidatorsOrbsAddress", "getElectedValidatorsReputation".
<!--??Note: current block-proposer applies both for ordering and execution (results_block and transactions_block might have different committees)-->
* `VirtualMachine` - Expose leader in execution context - for state update. 
* `Conensus-algo` - provides the block proposer identity.



#### RequestOrderingCommittee 
* A node queries the `ConsensusContext` for a sorted list of nodes (approval committee) 
at the beginning of each consensus round (block height). The `ConsensusContext` integrates the elected set of nodes 
with their reputation to provide the sorted list. 

#### Register Block-proposer
* The block proposer (which also appears in the block header), is registered using trigger transaction, in state by using `reputation` contract, if this contract exists. 
This information is later aggregated to provide basic reputation.  
* Note: Block proposer is validated against the appropriate leader provided by the `consensus-algo` (in leader change the block might not be replaced). 

## SDK update
#### Block.GetBlockProposer
* Returns the block proposer address.

## Consensus-algo update
#### RequestBlockProposal
* Provide the block proposer address in method call.

#### ValidateBlockProposal
* Provide the block proposer address in method call.


