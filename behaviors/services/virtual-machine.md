# Virtual Machine

Executes service methods (smart contracts) using various processors (languages) and produces state difference as a result.

Currently a single instance per virtual chain per node.

&nbsp;
## `Data Structures` <!-- tal will finish -->

#### Context
* Allocated for each TransactionSet or LocalMethod execution.
* ReadOnly / ReadWrite
  * RO context can not modify state, allocated for LocalMethod execution.
  * RW context incldues a transient state cache.
* Transient state
  * Stores temporary state changes during transaction set execution (until state diff is committed).
  * Relevant for ReadWrite methods only.
  * Contains state diff (possibly for multiple services).
  * Format is identical to the state store data structure in `StateStorage`.
* Service stack
  * Top of the stack indicates the current service address space for the context.

&nbsp;
## `Signature Schemes`

#### Address Scheme 01
* Check Ed25519 signature over the transaction header.

&nbsp;
## `RunLocalMethod` (method)

> Execute a read only method and return its result (not under consensus).

#### Prepare for execution
* Validate the call signature according to signature scheme.
* Retrieve the service processor by calling `StateStorage.ReadKeys` on the `_Deployments` service.
  * The key is hash(`<service-name>.Processor`).
  * If the service processor is not found, fail.
* Allocate a context:
  * ReadOnly (RunLocalMethod cannot update state)
  * No transient state.

#### Execute
* Push service to the context service stack.
* Execute the service method on the correct processor by calling `Processor.ProcessCallSet`.
  * Execution permissions are checked by the processor.
* Pop service from the content service stack.
* Return result.

&nbsp;
## `ProcessTransactionSet` (method)

> Process a group of transactions together that update state and return the combined state diff.

#### Prepare for execution
* Allocate a context:
  * ReadWrite
  * Initialize transient state.
* Go over all transactions in the set and for each one:
  * Retrieve the service processor by calling `StateStorage.ReadKeys` on the `_Deployments` service.
    * The key is hash(`<service-name>.Processor`).
    * If the service is not found, try to deploy it (only relevant for native services):
      * Check if it's a native service by calling the `Native` processor's `Processor.DeployService`.
  *




* Group transactions and allcoate a context:
  * Trasnactions of the same TransactionSet must be exectuted on the same context (executed in order, same state cache)
    * Allocate a RW context with a transient state cache that will hold updated state temporarily (across all transactions).
* Retrieve the services (NameSpace) code and metadata by calling `StateStorage.ReadKeys`
  * TODO cache/pre-fetch state variables (part of the contarct meta-data)
* Validate that the called methods exist and have appropriate permissions for execution.

#### Execute
* Group in order transactions that ar eexecuted on seperate processors and send for execution
* Execute the service methods by calling `Processor.ProcessCallSet` (streaming a set of transactions)
* Maintain the result of the call methods.

#### Prepare combined state diff
* Retrieve the changes between the transient state cache and the original state.
* Encode changes as state diff.

&nbsp;
## `TransactionSetPreOrder` (method)
> Approve transactions before allowing them to go through ordering

#### Check transaction signature
* Check the signature according to supported signature schemes, return `INVALID_SIGNATURE` on mismatch.
* Return `INVALID_ADDRESS_SCHEME` for unsupported signature schemes.

#### Check subscription
* Execute the platform pre order smart contract on the native processor.
* If the contract fails, return `NOT_APPROVED` and the return string from the contract (like `SUBSCRIPTION_ERROR`).

&nbsp;
## `SdkCall` (method)

> Calls an Sdk method executed by the VM. The supported Sdk calls are described in [VM Sdk calls](native.md)

#### `CallServiceMethod`

> Calls a method of another service for execution.

* Retrieve the services (NameSpace) code and metadata by calling `StateStorage.ReadKeys`
* Validate that the called methods exist and have appropriate permissions for execution.
* If the called method is executed by the same processor
  * Return same_processor_execution = TRUE, output_argument = NULL.
* If the called method is executed by a different processor
  * Call a `Processor.ProcessCallSet` with the relevant method and retrun the output arguments.

#### `CallLibraryMethod`
> Calls a method of another service for execution.
* Retrieve the services (NameSpace) code and metadata by calling `StateStorage.ReadKeys`
* Validate that the called methods exist and have appropriate permissions for execution.
* If the called method is executed by the same processor
  * Return same_processor_execution = TRUE, output_argument = NULL.
* If the called method is executed by a different processor
  * Call a `Processor.ProcessCallSet` with the relevant method and retrun the output arguments.

#### `StateReadKeys`
TODO

#### `StateWriteKeys`
TODO
