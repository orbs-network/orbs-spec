# Management Virtual Chain
> Provide a system wide source of truth across virtual chains.
> Initial feature abstracts the governance process results (occur on the Ethereum network) on an Orbs Validator Node level - across its virtual chains' instances.

### Design principles
* Assumptions:
    * "Trust" holds inside the Node (across its VCs instances)
    * Reference Validators set holds during VC genesis (even if an election event altered the validators set).
    * Current Validators set is "honest" (Quorum wise) during the election transition. 

* Hierarchical VCs design: ManagementVC and dependent virtual chains (coined "ConsumerVC").

* ManagementVC
    * Dedicated VC for cross system source of truth
    * Should be "pure": (V1) dedicated for elected validators - system config (future - subscription, and other cross VCs interaction).
    * Holds the last X states close to tip for state queries by block height (facilitates the inner node sync across its VC instances and in turn the underlying agreement).
    * Holds an Election contract for election information.

* ConsumerVC
    * Continuously sync system config.
        * Update logic - depends on the record expiration.
        * Read logic - holds until it is overwritten.
    * Close to real-time information retrieval, with small time phase shift (to support agreement - only for small inner node syncing issues (including its other VCs instances))
    * Maintain cache per block.timestamp (reference time -> offChain data). This implies the same context for the entire block of txs.
    * ManagementVC availability failures :
        * Unavailable  - currently treated as below.
        * Available outdated information - ConsumerVC "partially halts": application state does not progress, only system state can be updated (elected validators set, reputation..).
        
* Cross VCs communication:
    * Intra Validator node (across its VC instances ) - Direct HTTP Synchronous rpc calls.

* Audit node:
    * Audit node implementation relies on historical VC state queries (by block height).
* ConsumerVC node - block sync:
    * Not affected.
* ManagementVC protocol upgrade:
    * Currently, CrossChainAPI should be backward compatible (updating the CrossChainConnector could partially mask this upgrade).
* Testing environments ..


* Notes:
    * VC query by time (used to obtain latest committed block prior to the provided time) is non-deterministic by nature.
    * We address this query as off-chain data, which is not "fully auditable" (intuitively, similar to block time stamp).
    * Query by height is deterministic (reorgs is currently out of scope).
    * Expired elected validators set - allows partial state progress (only "system state" update: elected set, reputation .. ).
    * Transitioning to a new configuration is done agreement by writing to state. This configuration holds during the entire consensus round (updating state at round i determines the config of round i+1).
 
## VC Interop Flows:

### ConsumerVC RequestOrderedCommittee flow (Read)
* At the beginning of each consensus round (block height), the node queries the `ConsensusContext` for a sorted list of nodes (committee).
* The `ConsensusContext` retrieves the current Elected Validators Set from state by calling `_ElectedValidators.getElectedValidators()` (using the `VM.CallSystemContract`).
    * Poll until success.
    * If the Elected Validators Set is empty (state update has yet to occur) retrieves the Reference Set.
        * If `MANAGEMENT_MODE` is in `local` mode - retrieves from config file.    
        * If `remote` mode - calls `_ElectedValidators.getReferenceValidatorsSet()`.
            * Inside service: 
                <!-- Note: this is dependent on archive mode * Retrieve the Virtual Chain Reference blockHeight (virtual chain's genesis time according to the ManagementVC blockHeight) from the subscription - SDK call `Mgmt.CallContract('_Subscription', 'getSubscriptionDetails')`.
                    * If blockHeight  == 0 (reference blockHeight in subscription has not been written to state yet), Fail.
                * Retrieve the Reference Set from the ManagementVC by calling `Mgmt.CallContractAtBlock('_ElectedValidators', 'getElectedValidators', 'blockHeight')`. -->
            * Retrieve the Virtual Chain Reference Validators Set (virtual chain's validators at genesis) from the subscription using the SDK call `Mgmt.CallContract('_Subscription', 'getReferenceValidatorsSet')`.
* The `ConsensusContext` retrieves the ordered committee by calling `_Committee.getOrderedCommitteeForAddresses` with the Elected Validators Set as argument.
* Note: ManagementVC contract call is detailed in the update flow.


#### Participant Services
##### ConsumerVC
* `ConsensusAlgo` - ConsensusAlgo requests committee for new consensus round.
* `ConsensusContext` - Retrieves the ordered committee based on state.
* `VirtualMachine` - Provides the `ConsensusContext` with state queries access.
* `CrossChainConnector` - Optional - Get Reference elected validators from ManagementVC.

##### ManagementVC - Optional (reference set remote)
* Detailed in update flow.


### ConsumerVC block creation \ validation flow
* If the elected validators set is outdated the block should be empty - with only a single system transaction (`Trigger tx`).
    * Enforced using the pre order checks - `_GlobalPreOrder` contract checks the elected validators `ValidUntil` (call `_ElectedValidators.getElectedValidators()`) against the block.timestamp.
    * See [Block creation flow](../flows/block-creation-before-commit.md) for further details.


### ConsumerVC update elected validators flow
* `ConsensusAlgo` calls the `ConsensusContext` to create \ validate a new block proposal.
* `ConsensusContext` calls the VM to execute the block transactions using a proposer \ validator consensus role accordingly.
* The `VirtualMachine` executes the `Trigger transaction` which in turn calls the  [_ElectedValidators](../smart-contracts/system/_ElectedValidators.md) sync method.
* If the elected validators set should be updated the sync method retrieves an updated record by using the SDK call `Mgmt.CallContract('_ElectedValidators', 'getElectedValidators')`. 
* `Mgmt.CallContract` acquires the new record using two consecutive queries to the ManagementVC:
    * Get the ManagementVC BlockHeight - by calling `ManagementConnector.GetBlockInfoByTime` using current block.timestamp.
        * `ManagementConnector` deduces a "finality reference time" based on provided time reference and config value `MANAGEMENT_FINALITY_TIME_COMPONENT`.
        * `ManagementConnector` http calls the ManagementVC `CrosschainAPI.GetBlockInfoByTime` using the "finality reference time".
            * `CrosschainAPI` gets the block info by querying the `BlockStorage`.
        * BlockProposer logs the query result in the transaction receipt - as key value in the offchain_data section.
        * Validator verifies this logged information against its own query's result (no extra logging).
            * If the proposed ManagementVC BlockHeight is invalid the Validator fails the entire block.
    * Get the ManagementVC elected validators record - by calling `ManagementConnector.CallContract` using the fetched BlockHeight.
        * `ManagementConnector` http calls the ManagementVC `CrosschainAPI.CallContract`.
                * `CrosschainAPI` relays the call to the VM.
        * Both the BlockProposer and Validators log the query result in the transaction receipt.
* Notes:
    * The elected validators set is recorded in the VC state inside the [_ElectedValidators](../smart-contracts/system/_ElectedValidators.md) system contract.
    * The "finality reference time" is deterministically deduced from the block time; Aims to overcome inter-node sync across VCs.
    * The ManagementVC BlockHeight is non-deterministic (up to 1 block height difference when using the same time reference) - requires a softer verification rule (similar to block time stamp), and leaves some room for manipulation.
    * The SDK call `Mgmt.CallContract` also returns a call_result indicating valid or invalid completion of the call.
        * If the call was not completed successfully the record will not be updated. Failure tolerance coincides with the underlying consensus process. 
    
#### Participant Services
##### ConsumerVC
* `ConsensusContext` - Flags the VM with the node's consensus role: either as leader (proposer) or non-leader (validator).
* `VirtualMachine` - Exposes and validates the ManagementVC queries' results. Fails the entire block (ProcessTxSet - return Error) when validation fails.  Write\read offChain data to tx receipt. Provide Mgmt.SDK calls using `CrossChainConnector`. 
* `CrossChainConnector` - Retrieves the ManagementVC BlockHeight. Retrieves the elected validators from ManagementVC.

##### ManagementVC
* `CrosschainAPI` - Exposes "queries API" and serves the queries' results as obtained from the VM.
* `VirtualMachine` - Obtains Elected validators from state by block-height (supports historical state query).
* `BlockStorage` -  Obtains blockInfo (height, timestamp) by time reference or by height.



##

## Node config update
* MANAGEMENT_FINALITY_TIME_COMPONENT - reduce current block time by this constant (~ 10 seconds).
* MANAGEMENT_VIRTUAL_CHAIN_ID - the ManagementVC virtual chain id.
* MANAGEMENT_PROTOCOL_VERSION. <!-- - //TODO: upgrade -->
* MANAGEMENT_ENDPOINT - IP:Port.
<!-- * MANAGEMENT_MAX_SYNC_REJECT_TIME - allowed time drift of config - continue running with old information (~ 10 minutes) - used for pre order in `_GlobalPreOrder`. -->
* MANAGEMENT_RETRY_INTERVAL - polling param.
* MANAGEMENT_MODE - remote or local. 


## SDK update

#### Mgmt.CallContract(contract_name, method_name, input_arguments) : call_result, output_arguments

* Output: 
    * call_result indicating successful call or not. 
    * output_arguments of a remote contract call on the ManagementVC.

* If `MANAGEMENT_MODE` is local return `CALL_RESULT_FAILURE`, nil.
* Get the ManagementVC-BlockHeight according to reference timestamp (block.timestamp).
    * queried_result := ManagementConnector.GetBlockInfoByTime(block.timestamp). 
    * Note: non-deterministic result.

* BlockProposer logs the non-deterministic data (ManagementVC-BlockHeight) in transaction receipt
    * If ConsensusRole == CONSENSUS_ROLE_PROPOSER, write ManagementVC-BlockHeight into the transaction receipt - inside the offchain_data section.
* Validator verifies non-deterministic data (ManagementVC-BlockHeight)
    * If ConsensusRole == CONSENSUS_ROLE_VALIDATOR, read proposed ManagementVC-BlockHeight from transaction receipt and verify against queried result.
        * Verify proposed == queried_result OR proposed == queried_result - 1.
    * If verification fails, fail the block.
* Use proposed ManagementVC-BlockHeight to query the ManagementVC state.
    * (call_result, output_arguments) = ManagementConnector.CallContract(proposed_blockHeight, contract_name, method_name, input_arguments)
    * Log the call output into the transaction receipt - inside the offchain_data section.
* Usage: query ManagementVC state.


#### Mgmt.CallContractAtBlock(block_height, contract_name, method_name, input_arguments) : call_result, output_arguments
* Output: 
    * call_result indicating successful call or not. 
    * output_arguments of a remote contract call on the ManagementVC.
* If `MANAGEMENT_MODE` is local return `CALL_RESULT_FAILURE`, nil.
* (call_result, output_arguments) = ManagementConnector.CallContract(block_height, contract_name, method_name, input_arguments)
* Log the call output into the transaction receipt - inside the offchain_data section.


 ### Update Elected Validators Flow Diagram
 
 ![alt text][update_elected_validators_flow] <br/><br/>
 
 [update_elected_validators_flow]: ../../behaviors/_img/update_elected_validators_consumer_vc.png "Update Flow"



 ### Get Committee Flow Diagram
 
 ![alt text][get_committee_flow] <br/><br/>
 
 [get_committee_flow]: ../../behaviors/_img/get_committee_consumer_vc.png "Get Flow"
