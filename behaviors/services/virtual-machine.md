# Virtual Machine

Executes [deployed service](../../terminology.md) methods (smart contracts) using various processors (languages) and produces state difference as a result.

Currently a single instance per virtual chain per node.

#### Interacts with services

* `Processor` - Uses it to actually execute the smart contract method calls.
* `StateStorage` - Reads deployments from state and provides API for contracts to read from state.
* `CrosschainConnector` - Uses it to perform cross-chain contract calls.

&nbsp;
## `Data Structures`

#### Transient state
* Stores temporary state changes during transaction set execution (until state diff is committed).
* Relevant for `ReadWrite` methods only.
* Contains state diff (possibly for multiple services).
* Format is identical to the state store data structure in `StateStorage`.

#### Execution context
* Allocated for each Transaction or LocalMethod execution in the `VirtualMachine`.
* Specified as `ReadOnly` or `ReadWrite`.
  * `ReadOnly` context cannot modify state, allocated for LocalMethod execution.
  * `ReadWrite` context can modify state, includes a transient state cache.
* Service stack.
  * Top of the stack indicates the current service address space for the context.
  * When we nest service calls, the address space changes.
* Transaction transient state.
  * Every transaction must maintain its own temporary transient state since it can fail (and then rollback its writes).
  * Relevant for `ReadWrite` execution contexts only.
* Batch transient state pointer (points to the batch transient state which is defined outside the execution context).
  * The combined transient state of the entire batch (normally an entire block of transactions).
  * Relevant for `ReadWrite` execution contexts only.
* Events logs
  * Logs events emitted by the SDK `Log.Emit` function and are relevant only for `ReadWrite` execution contexts.
  * Each event is associated with the contract that triggered it.
  * The events functionality is initially reserved for crosschain operations. Only a single event may be emitted per transaction, a second emit fails the transaction execution.

&nbsp;
## `Init` (flow)

* Initialize the [configuration](../config/services.md).
* For each processor, register to handle the processor's SDK calls by calling `Processor.ContractSdkCallHandler`.

&nbsp;
## `RunLocalMethod` (method)

> Executes a read only method of a deployed service and returns its result (not under consensus).

#### Prepare for execution
* Get the block height and timestamp for the local method processing by calling `StateStorage.GetStateStorageBlockHeight`.
  * Note that method calls are asynchronous to block creation so execution may end up a few blocks behind.
  * Note that the reference block height and timestamp are returned to the caller on successful execution and on failure.
* If signed, validate the call signature according to the signature scheme (see transaction format for structure).
  * Currently `PublicApi.CallMethod` calls are not signed.
* Retrieve the service processor by calling system contract `_Deployments.getInfo` and fail if not deployed.
  * See `_Deployments` contract [specification](../smart-contracts/system/_Deployments.md).
* Allocate an execution context:
  * `ReadOnly` (cannot update state since not under consensus).
  * No transient state (no transaction transient state and no batch transient state).

#### Execute method call
* Push service to the execution context's service stack.
* Execute the service method on the correct processor by calling `Processor.ProcessCall`.
  * Note: Execution permissions are checked by the processor.
* Pop service from the execution context's service stack.
* Return result along with the reference block height and timestamp.

&nbsp;
## `ProcessTransactionSet` (method)

> Processes a batch of transactions on deployed services together. The transactions may update state so returns the combined state diff.

### Prepare for the batch
* Allocate a batch transient state that will hold updated state (across all transactions in the batch).

### Execute all transactions
* Go over all transactions in the set (in order) and for each one:  

#### Prepare for execution (each transaction)
* Retrieve the service processor by calling system contract `_Deployments.getInfo`.
  * If the service is not found, try to auto deploy it (only relevant for native services):
    * Check if it's a native service by calling the `Native` processor's `Processor.GetContractInfo`.
    * If so, auto deploy the service by calling system contract `_Deployments.deployService`.
    * See `_Deployments` contract [specification](../smart-contracts/system/_Deployments.md). 
* Allocate an execution context:
  * `ReadWrite` (can update state since under consensus).
  * New transaction transient state and pointer to the batch transient state.

#### Execute method call (each transaction)
* Push service to the execution context's service stack.
* Execute the service method on the correct processor by calling `Processor.ProcessCall`.
  * Note: Execution permissions are checked by the processor.
* Pop service from the execution context's service stack.
* If the transaction was successful:
  * Apply the transaction transient state to the batch transient state.
  * Add the event logs to the transaction receipt.
* Remember the result of the method call and generate a transaction receipt.

#### Prepare combined state diff
* Retrieve the changes between the batch transient state and the original state.
* Encode changes as state diff.

&nbsp;
## `TransactionSetPreOrder` (method)

> Approves transactions before allowing them to go through ordering (the virtual chain subscription is checked here for example).

#### Check transaction signatures
* Check the transaction signatures according to the supported signature schemes (see transaction format for list).
* Fail if unsupported signature scheme.
* Successful pre order check should return the transaction status `TRANSACTION_STATUS_PRE_ORDER_VALID`.

#### Run system contract
* Approval of a transaction for ordering and execution consists of 3 layers:
  * Approval on a global system level (level 1/3).
    * Run system smart contract `_GlobalPreOrder.Approve` by calling the `Native` processor's `Processor.ProcessCall`.
      * Performed once per transaction set, does not depend on the transactions content.
      * See `_GlobalPreOrder` contract [specification](../smart-contracts/system/_GlobalPreOrder.md).
  * Approval on the virtual chain level (level 2/3) not supported yet.
  * Approval on the smart contract level (level 3/3) not supported yet.
  
#### Transaction set status
* If one of the trasnactions in the set fails its pre-order check, return error

&nbsp;
## `SdkCall` (method)

> Implements a smart contract SDK method. Called by the processor whenever it is unable to implement the SDK method itself and requires data from the system. Supported SDK calls are described [here](../smart-contracts/sdk/api.md).

#### `Service.CallMethod`

> Calls a method of another service on the virtual chain.

* Push service to the execution context's service stack.
* Execute the service method on the correct processor by calling `Processor.ProcessCall`.
  * Note: Execution permissions are checked by the processor.
* Pop service from the execution context's service stack.
* Return result.

#### `State.Read`

> Reads a variable from the state of the service.

* Identify the service we're reading from, it's the top of the execution context's service stack.
* Try to read the variable from the transaction transient state (if found there).
* If not found, try to read the variable from the batch transient state (if found there).
* If not found, read the variable from state storage by calling `StateStorage.ReadKeys`.
  * Called with block_height equals to the last committed block (current block height - 1)

#### `State.Write`

> Writes a variable to (transient) state of the service.

* Make sure the execution context is `ReadWrite` and we have a transient state.
  * Otherwise, terminate the execution and return ERROR_STATE_WRITE_IN_READONLY_CALL
* Identify the service we're writing to, it's the top of the execution context's service stack.
* Write the variable to the transaction transient state.

&nbsp;
## SDK Call Handler

> Handles the processor's SDK calls.

#### `HandleSdkCall`
* Handle by calling `SdkCall`.

