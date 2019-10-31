# Management Virtual Chain
> Provide a system wide source of truth across virtual chains.
> Initial feature abstracts the governance process results (occurs on the Ethereum network) on an Orbs Validator Node level - across its virtual chains' instances.

### Design principles
* Assumptions:
    * "Trust" holds inside the Node (across its VCs instances)
    * Reference Validators set holds during VC genesis and until the state is updated (even if an election event altered the validators set - to avoid network split).
    * Current Validators set is "honest" (Quorum wise) during the election transition (PoS property). 

* Hierarchical VCs design: ManagementVC and dependent virtual chains (coined "ConsumerVC").
* Custom system contracts deployment for the different VCs flavors (ManagementVC, ConsumerVC, PrivateVC). 
    * Currently, this only holds for `_ElectedValidators` system contract (interchangeably coined "Elected Validators Provider"). Where there are two types of providers: remote and local.

* ManagementVC
    * Dedicated VC for cross system source of truth
    * Should be "pure": (V1) dedicated for elected validators - system config (future - subscription, and other cross VCs interaction).
    * Holds the last X states close to tip for state queries by block height (facilitates the inner node sync across its VC instances and in turn the underlying agreement).
    * Holds the `_Election` contract for Ethereum election information.
        * The `_Election` contract takes the role of the "Elected Validators Provider" on the ManagementVC.
    * Holds the `_Provision` system contract which provides information about the different ConsumerVCs subscription and initial config - includes the Reference Validators Set (currently not supported).

* ConsumerVC
    * Continuously sync system config.
        * Update logic - depends on the record expiration.
        * Read logic - holds until it is overwritten.
    * Close to real-time information retrieval, with small time phase shift (to support agreement - only for small inner node syncing issues (including its other VCs instances))
    * Maintain cache per block.timestamp (reference time -> offChain data). This implies the same context for the entire block of txs.
    * ManagementVC availability failures :
        * Available outdated information - ConsumerVC "partially halts": application state does not progress, only system state can be updated (elected validators set, reputation..).
        * Unavailable - currently treated the same as above.
    * Elected validators provider := [_ElectedValidators](../smart-contracts/system/_ElectedValidators.md) system contract. 
        * The Reference Elected Validators is fetched from the ManagementVC by querying the `_Provision` contract.
    

* PrivateVC
    * Does not depend on another VC to work.
    * Hierarchical VCs design can still hold where other Private-ConsumerVCs use it as their ManagementVC.
    * Elected validators provider := [_ElectedValidators](../smart-contracts/system/_ElectedValidators.md) system contract in "dummy" mode - with rudimentary logic. 
    
* Cross VCs communication:
    * Intra Validator node (across its VC instances ) - Direct HTTP Synchronous rpc calls.

* Audit node:
    * Audit node implementation relies on historical VC state queries (by block height).
* Node - block sync:
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

## VC Flows:

### RequestOrderedCommittee flow (Read)
* At the beginning of each consensus round (block height), the node queries the `ConsensusContext` for a sorted list of nodes (committee).
* The `ConsensusContext` retrieves the current Elected Validators Set from state.
    * Poll until success (retry on error).
        * Get the Elected Validators state record  by calling `_ElectedValidators.getElectedValidators()` (using the `VM.CallSystemContract`).
        * If the Elected Validators state record is empty (state update has yet to occur) the `ConsensusContext`  the `_ElectedValidators` will return the Reference Validator Set by calling `_Provision.getReferenceElectedValidators()` (currently not supported).
            * Note: the ConsumerVC will query the ManagementVC `_Provision` system contract. 
    * If the resulting Elected Validator record is empty (with no error) get the Reference Validators Set from local config file.
    * Note: empty set with no error implies "Local Reference Validators" mode.
* The `ConsensusContext` retrieves the ordered committee by calling `_Committee.getOrderedCommitteeForAddresses` with the Elected Validators Set as argument.
* Note: ManagementVC contract call is detailed in the update flow.
* Note: The ManagementVC, ConsumerVC and PrivateVC behaviors coincide here (abstracted away by the Elected Provider system contract).

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
    * Enforced by using the pre order checks - `_GlobalPreOrder` contract which determines whether the current elected validators record is outdated.
    * See [Block creation flow](../flows/block-creation-before-commit.md) for further details.


### ConsumerVC update elected validators flow
* `ConsensusAlgo` calls the `ConsensusContext` to create \ validate a new block proposal.
* `ConsensusContext` calls the VM to execute the block transactions using a proposer \ validator consensus role accordingly.
* The `VirtualMachine` executes the `Trigger transaction` which in turn calls the `_ElectedValidators.sync()` method.
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
    * ManagementVC elected validators update occurs in `_Elections.sync()` which in turn processes the Elections result as relayed from Ethereum (depends on the elections time). 
    * PrivateVC elected validators update does nothing - `_ElectedValidators.sync()` is a dummy method. 
 
#### Participant Services
##### ConsumerVC
* `ConsensusContext` - Flags the VM with the node's consensus role: either as leader (proposer) or non-leader (validator).
* `VirtualMachine` - Exposes and validates the ManagementVC queries' results. Fails the entire block (ProcessTxSet - return Error) when validation fails.  Write\read offChain data to tx receipt. Provide Mgmt.SDK calls using `CrossChainConnector`. 
* `CrossChainConnector` - Retrieves the ManagementVC BlockHeight. Retrieves the elected validators from ManagementVC.

##### ManagementVC
* `CrosschainAPI` - Exposes "queries API" and serves the queries' results as obtained from the VM.
* `VirtualMachine` - Obtains Elected validators from state by block-height (supports historical state query).
* `BlockStorage` -  Obtains blockInfo (height, timestamp) by time reference or by height.



### Node Init update
* When the system initiates it deploys a subset of the system contracts (in the repository) based on the configuration. 


## Node config update
* ELECTED_VALIDATORS_PROVIDER - used in custom system contracts deployment. The system contract name which implements the "Elected Validators Provider Interface" - differentiates between ManagementVC, ConsumerVC and PrivateVC.
<!-- * PROVISION_PROVIDER - used in custom system contracts deployment. The system contract name which implements the "Provision Interface" - differentiates between ManagementVC, ConsumerVC and PrivateVC. -->
* STATE_STORAGE_HISTORY_SNAPSHOT_NUM - number of past states the state storage maintains. On ManagementVC this should be increased (~10 seconds => 20 blocks).
* MANAGEMENT_FINALITY_TIME_COMPONENT - ManagementConnector reduces current block time by this constant (~ 10 seconds) on calls to the ManagementVC.
* MANAGEMENT_VIRTUAL_CHAIN_ID - the ManagementVC virtual chain id.
* MANAGEMENT_PROTOCOL_VERSION. <!-- - //TODO: upgrade -->
* MANAGEMENT_ENDPOINT - IP:Port.
* ELECTED_VALIDATORS_RETRY_INTERVAL - polling param.

## Elected Validators Provider Interface
* getElectedValidators() : (validators_list, election_number, next_sync)
    * Return the state record or Reference Set.
        * Reference Validators set (elected validators state record was not set) flavors:
            * local (ManagementVC and PrivateVC) - return nil. The `ConsensusContext` will subsequently retrieve the Reference Set from local config file.
            * remote (ConsumerVC) - get reference set from ManagementVC (currently not supported).
    * On the ManagementVC, next_sync coincides with the election time.


* sync()
    * Conditionally update the elected validators (if the state record is outdated := next_sync is in the past):
        * ManagementVC - processes the Elections result as relayed from Ethereum.
        * ConsumerVC - fetches the new elected validators set from ManagementVC.
        * PrivateVC - does nothing.
        
* getElectedValidatorsNextSync() : next_sync
    * Used for PreOrder checks. 


<!--## Provision Provider Interface
* getSubscriptionInfo(VirtualChainID): (subscription_record)
    * Used for PreOrder checks in ConsumerVC.
    
* sync()
    * Conditionally update the subscription info.
     

* getReferenceValidatorsSet(VirtualChainID): validators_list
   * Used in ConsumerVC genesis phase.
-->


## SDK update

#### Mgmt.CallContract(contract_name, method_name, input_arguments) : call_result, output_arguments

* Output: 
    * call_result indicating successful call or not. 
    * output_arguments of a remote contract call on the ManagementVC.

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


<!-- #### Mgmt.CallContractAtBlock(block_height, contract_name, method_name, input_arguments) : call_result, output_arguments
* Output: 
    * call_result indicating successful call or not. 
    * output_arguments of a remote contract call on the ManagementVC.
* (call_result, output_arguments) = ManagementConnector.CallContract(block_height, contract_name, method_name, input_arguments)
* Log the call output into the transaction receipt - inside the offchain_data section.
* Usage: query ManagementVC state. -->

 ### Update Elected Validators Flow Diagram
 
 ![alt text][update_elected_validators_flow] <br/><br/>
 
 [update_elected_validators_flow]: ../../behaviors/_img/update_elected_validators_consumer_vc.png "Update Flow"



 ### Get Committee Flow Diagram
 
 ![alt text][get_committee_flow] <br/><br/>
 
 [get_committee_flow]: ../../behaviors/_img/get_committee_consumer_vc.png "Get Flow"
