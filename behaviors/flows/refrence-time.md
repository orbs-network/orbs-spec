# Refrence Time

## Services handling of Reference Time

* `Consensus Algo` - Provides previous block data for next block consensus, including Reference Time
* `Consensus Context` - Checks block protocol version as a function of the Reference Time
                        Gets the block committee as a function by calling the eelection smart contract
* `TransactionPool` - Proposes and checks current refrence time on new block
* `VirtualMachine` - Executes PerOrder checks, checks subscription validaity as a function of the Reference Time
                     Gets the block committee as a function of the Reference Time
* `StateStorage` - stores the last block data for RunQuery use including Reference Time


The refrence time (or management reference time) indicates the update time of which the virtual chain management data reffers to. The management data includes: the valdiators commiteee, virual chain subscription and properties and protocol version.

Each validator node holds a local reference time and a consensus refernce time that represnts an agreed reference time to which the validators are synched to and determines the virtual chain properties. The consensus refernce time is writen in the block header.

## Refernce Time Usage
* Reference Time for Committee - determined by the last agreed reference time i.e. the reference time of the previous block.
  The committee requires a reference time under consensus in order to esure consistency of the current committee and leader and prevent forks.

* Block Reference Time - used in order to query management properties determined by the reference time such as protocol version,   subscription status and virtual chain subscription properties.

## Reference Time Source per Flow
* `RequestOrderingCommitee` (Consensus committee): 
  * Reference Time for Committee: `prev_block_reference_time`
  * Block Reference Time: N/A
* `RequestNewBlock`: 
  * Reference Time for Committee: `prev_block_reference_time`
  * Block Reference Time: `ManagementService.GetReferenceTime`
* `ValidateBlock`: 
  * Reference Time for Committee: `prev_block_reference_time`
  * Block Reference Time: `block.reference_time` (Validated comapred to `ManagementService.GetReferenceTime`)
* `RunQuery`: 
  * Reference Time for Committee: N/A (`StateStorage.GetLastCommittedBlockInfo`)
  * Block Reference Time: `StateStorage.GetLastCommittedBlockInfo`
* `AddNewTransaction`: 
  * Reference Time for Committee: N/A
  * Block Reference Time: `ManagementService.GetReferenceTime` 
    * Acts as if it was a leader in a new block.

## Genesis Block, Block #1 Creation 
* Each virtual chain is created with a refrence time to be used for its first block. As the commitee must be determined in consensos the genesis reference time is determined in advance, alowing the validators to set up the resources and agree on the first block
* `RequestOrderingCommitee` for Block #1 
  * Reference Time for Committee: `ManagementService.GetGenesisReferenceTime`
  * Check `ManagementService.GetReferenceTime` is larger or equal to `ManagementService.GetGenesisReferenceTime` is larger or equal to `ManagementService.GetReferenceTime`, otherwise fail the request 
* `RequestNewBlock` / `ValidateBlock`
  * Reference Time for Committee: `ManagementService.GetGenesisReferenceTime`

## Reference Time Proposal and Consensus

#### Reference Time Proposal
* Reference Time = Max[`GetCurrentReference()`, block[n-1].`reference_time`]

#### Reference Time Proposal Validation
* Check that Block Reference Time is larger or equal than the previous block Reference Time
* Compare to the Validator Reference Time. A variation (grace) is allowed between the two values:
  * Check that Block Reference_time + `REFERENCE_TIME_CONSENSUS_GRACE` (default 10 min) < Validator Reference Time

## Reference Time Liveness
* The network reference time livness determines the maximum time of which the netwerk can operate without an updaetd management data. 
* Check Block Reference_time > Block timestamp - `REFERENCE_TIME_LIVENESS_GRACE` (default 12 hours)
* Upon experiation, no user trasnactions are executed (PreOrder fails) 


## `RequestNewCommittee` Flow
#### Particiapnts
  * `ConsensusAlgo`
  * `ConsensusContext`
  * `VirtualMachine`
  * `ManagementService`

#### Flow
* `ConsensusAlgo` request for new committee providing the previous block Reference Time
  * `ConsensusContext` - calls the Committee system contarct, providing the previous block Reference Time to the `Virtual Machine`
    * `VirtualMachine` - executes the smart contract, exposing an SDK for `GetCommittee` that queries the `ManagementService`. Committee system contarct calculates the committee order.


## `RequestNewOrderingBlock` / `ValidateOrderingBlock` Flow

#### Particiapnts
  * `ConsensusAlgo`
  * `ConsensusContext`
  * `TransactionPool`
  * `VirtualMachine`
  * `ManagementService`

#### Flow

* `ConsensusAlgo` 
  * Requests new block providing the previous block Reference Time
  
  * `ConsensusContext`
    * Requests transactions from the `TransactionPool`

    * `TransactionPool`
      * Calcualtes (leader) / validates (Validator) a Reference Tine Proposal
      * Calls PreOrder checks in the `VirtualMachine` with the Reference Time Proposal

      * `VirtualMachine` executes PreOrder checks. Checks susbcription status by querying the `ManagementService` and checks Reference Time liveness
  
  * `ConsensusContext`
    * Calcualtes / validates the protocol version based on Reference Time proposal calling the `ManagementService`

## `RequestNewResultsBlock` / `ValidateResultsBlock` Flow

#### Particiapnts
  * `ConsensusAlgo`
  * `ConsensusContext`
  * `VirtualMachine`
  * `ManagementService`

#### Flow

* `ConsensusAlgo` 
  * Requests new block providing the previous block Reference Time, current Refrece Time and protocol version.
  
  * `ConsensusContext`
    * Add a trigger transaction to update the reputation.
    * Send the transactions to the `VirtualMachine` for execution.

    * `VirtualMachine`
      * Reputation contract calls `GetCommitee` SDK the retrives the committee from the `ManagementService` using the previous block Reference Time.


