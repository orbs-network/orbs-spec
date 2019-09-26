# Management Virtual Chain
> Provide a system wide source of truth across virtual chains.
> Initial feature abstracts the governance process results (occur on the Ethereum network) on an Orbs Validator Node level - across its virtual chains' instances.

### Design principles
* Assumptions:
    * "Trust" holds inside the Node (across its VCs instances)

* Hierarchical VCs design: ManagementVC and dependent virtual chains (coined "ConsumerVC").
* ManagementVC
    * Dedicated VC for cross system source of truth
    * Should be "pure": (V1) dedicated for elected validators - system config (future - subscription, and other cross VCs interaction).
    * Holds the last X states close to tip for state queries by block height (facilitates the inner node sync across its VC instances and in turn the underlying agreement).

* ConsumerVC
    * Continuously sync system config.
    * Close to real-time information retrieval, with small time phase shift (to support agreement - only for small inner node syncing issues (including its other VCs instances))
    * Maintain cache per block.timestamp (reference time -> offChain data). This implies the same context for the entire block of txs.
    * ManagementVC availability:
        * Does not respond - ConsumerVC halts.
        * Responds but does not progress - ConsumerVC progresses for a bounded allowed time diff (avoid running on stale config for too long).

* Cross VCs communication:
    * Direct HTTP Synchronous rpc calls.

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
  
 
## VC Interop Flows:

### ConsumerVC Update elected validators

#### Update elected validators
* The elected validators set is recorded in the [_ElectedValidators](../smart-contracts/system/_ElectedValidators.md) system contract.
* The state is updated using the Triggers mechanism (if cached record has expired), based on information obtained from the MgmtVC.
* The updated elected validators set information is retrieved using the SDK call `Mgmt.GetElectedValidators()`.
* Notes:
   * `Mgmt.GetElectedValidators()` acquires the new record using two consecutive queries to the ManagementVC (through the `CrossChainConnector`):
    * Get ManagementVC BlockHeight based on "finality reference time".
    * Get elected validators using the proposed MgmtVC BlockHeight.
   * Block Proposer logs the queries results in the transaction receipt.
   * Validator verifies this logged information against its own queries.
   * The "finality reference time" is deterministically deduced from the block time, aims to overcome inter-node sync across VCs.
   * The ManagementVC BlockHeight is non-deterministic (up to 1 block height difference when using the same time reference) - requires a softer verification rule (similar to block time stamp), and leaves some room for manipulation.
   * Validator first verifies the proposed ManagementVC BlockHeight (in tx receipt) using its own query. Subsequently, it retrieves the elected validators for this height, which is implicitly verified using the block state diffs.
   * If the proposed ManagementVC BlockHeight is invalid the Validator fails the entire block.

#### Participant Services
##### ConsumerVC
* `ConsensusContext` - Flag the VM with the execution mode: either as proposer or validator.
* `VirtualMachine` - Expose and verify the block proposer's MgmtVC BlockHeight. Fail block validation (ProcessTxSet - return Error) if verification fails.  Write\read offChain data to tx receipt. Provide Mgmt.SDK calls using `CrossChainConnector`. 
* `CrossChainConnector` - Get MgmtVC BlockHeight under finality time shift. Get elected validators from MgmtVC (based on block height).

##### ManagementVC
* `CrosschainAPI` - Expose "queries API" and serve the queries' results as obtained from the VM.
* `VirtualMachine` - Obtain Elected validators from state by block-height (support historical state query).
* `BlockStorage` -  Obtain blockInfo (height, timestamp) by time reference or by height.
 
### ConsumerVC RequestOrderingCommittee
* At the beginning of each consensus round (block height), the node queries the `ConsensusContext` for a sorted list of nodes (ordered committee).
* The `ConsensusContext` retrieves an ordered validators list from state by calling the `_Committee.getOrderedCommittee` (using the `VirtualMachine`).
* The `_Committee.getOrderedCommittee` relies on a call to `_ElectedValidators.fetchElectedValidators(now)`.
* `_ElectedValidators.fetchElectedValidators` retrieves either the elected validator set from cache or an updated record from ManagementVC.
* Elected validator set is expired if one of the following holds:
    * No elected validator set is yet stored.
    * The stored elected validator set dates range (valid from - to) has passed (based on the provided reference timestamp). 
* If the record has expired fetch the new record by calling `Mgmt.GetElectedValidators()`.
* Note: implicit consensus occurs when the cached record has expired (for example on system upload). This could potentially close a new block with some of the nodes holding a "forked view" of validator set.
 "Forked view" is only temporary due to 1 block height potential difference and converges when the agreement over the new block has been reached.
 * TODO: Special care? verify safety. Maximal diff between consecutive elected validators is less than 'f'.


#### Participant Services
##### ConsumerVC
* `ConsensusAlgo` - ConsensusAlgo requests committee for new consensus round.
* `ConsensusContext` - Polls until success from `_Committee` system contract which relies on `_ElectedValidators` contract (pass in reference time := now). Flag the VM with the proposer execution mode on `CallSystemContract`.
* `VirtualMachine` - Provide the `ConsensusContext` the state queries access.  Note if the cached elected validators set has expired (see below), it uses the `CrossChainConnector` to retrieve the elected validators from ManagementVC.
* `CrossChainConnector` - Get ManagementVC BlockHeight under finality time shift. Get elected validators from ManagementVC (based on block height).

##### ManagementVC
* Same as the update flow specified above.


### VC system init:
#### Participant Services
* `CrossChainConnector` - Init the appropriate connector based on the Node Config - either Ethereum or Management.

## Node config update
* MANAGEMENT_FINALITY_TIME_COMPONENT - reduce current block time by this constant (~ 10 seconds).
* MANAGEMENT_VIRTUAL_CHAIN_ID - the ManagementVC virtual chain id.
* MANAGEMENT_PROTOCOL_VERSION - //TODO: upgrade
* MANAGEMENT_ENDPOINT - IP:Port.
* MANAGEMENT_MAX_SYNC_REJECT_TIME - allowed time drift of config - continue running with old information (~ 10 minutes).
* MANAGEMENT_RETRY_INTERVAL - polling param.







## SDK update
#### Mgmt.GetElectedValidators(time_reference) : (node_address[], election_number, election_start_time, election_end_time)
* Output: the elected validators set in ManagementVC (at a deduced block height), as a list of Orbs addresses along with the corresponding election information (number, start time, end time).
* Get the ManagementVC-BlockHeight according to reference timestamp (non-deterministic result)
    * queried_result := MgmtGetBlockInfoByTime(block.timestamp)
* BlockProposer logs the non-deterministic data (ManagementVC-BlockHeight) in transaction receipt
    * If ExecutionMode == EXECUTION_PROPOSE, write queried_result into the transaction receipt.
* Validator verifies non-deterministic data (ManagementVC-BlockHeight)
    * If ExecutionMode == EXECUTION_VERIFY, read proposed ManagementVC-BlockHeight from transaction receipt and verify against queried result.
        * Verify proposed == queried_result OR proposed == queried_result - 1.
    * If verification fails, fail the block.
* Use proposed ManagementVC-BlockHeight to get and return the elected validators
    * (execution_result, output_arguments) = MgmtCallContract(proposed_blockHeight, "_ElectedValidators", "getElectedValidators" )
    * output_arguments := node_address[], election_number, election_start_time, election_end_time
* Usage: mirror current elected validators set from MgmtVC. 






 ### Update Elected Validators Flow Diagram
 
 ![alt text][update_elected_validators_flow] <br/><br/>
 
 [update_elected_validators_flow]: ../../behaviors/_img/update_elected_validators_consumer_vc.png "Update Flow"



 ### Get Committee Flow Diagram
 
 ![alt text][get_committee_flow] <br/><br/>
 
 [get_committee_flow]: ../../behaviors/_img/get_committee_consumer_vc.png "Get Flow"
