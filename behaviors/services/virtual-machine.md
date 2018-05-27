# Virtual Machine

Executes service methods (smart contracts) using various processors (languages) and produces state difference as a result.

Currently a single instance per virtual chain per node.

&nbsp;
## `Data Structures`

### Context
> Allocated for each TransactionSet or Method execution. Holds a transient state cache & address space.
* ReadOnly / ReadWrite
  * RO context can not modify state, allocated for callMethod execution.
  * RW context incldues a transient state cache.
* Address space
  * Indicates the current address space for the context. The context address space is a function of the service.

#### Transient state diff cache
> Maintains the transient state diff
* Allocated only for RW process trasnaction operations.
* Stores temporary state changes during transaction set execution (until state diff is committed).
* Format is identical to the state store data structure in `StateStorage`.
* No need to be persistent since a new instance is created per execution.
* No limit on max size.


&nbsp;
## `Signature Schemes`

#### Address Scheme 01
* Check Ed25519 signature over the transaction header.

&nbsp;
## `RunLocalMethod` (method)
> Execute a read only method its result.

#### Prepare for execution
* Validate the transaction signature by performing: `Sender and Signature Validation`.
* Retrieve the service (NameSpace) code and metadata by calling `StateStorage.ReadKeys`
  * TODO cache/pre-fetch state variables (part of the contarct meta-data)
* Validate that the method exists and has appropriate permissions for execution: Read only, external.
* Allocate a RO context. (CallMethod cannot update state, state writes must throw an error).

#### Execute
* Execute the service method by calling `Processor.ProcessCallSet`
* Provide the processor with the requried services (`GetNameSpaceCode`,`SdkCall`,`ServiceMethodCall`,`LibraryMethodCall`)
* Return result arguments to `Public API`.

&nbsp;
## `ProcessTransactionSet` (method)
> Process a group of transactions together that update state and return the combined state diff

#### Prepare for execution
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
* Validate that the method has the appropriate permissions for execution.
* Execute and return the output arguments.

&nbsp;
## `ServiceMethodCall` (method)
> Calls a method of another service for execution.
* Retrieve the services (NameSpace) code and metadata by calling `StateStorage.ReadKeys`
* Validate that the called methods exist and have appropriate permissions for execution.
* If the called method is executed by the same processor
  * Return same_processor_execution = TRUE, output_argument = NULL.
* If the called method is executed by a different processor
  * Call a `Processor.ProcessCallSet` with the relevant method and retrun the output arguments.

&nbsp;
## `LibraryMethodCall` (method)
> Calls a method of another service for execution.
* Retrieve the services (NameSpace) code and metadata by calling `StateStorage.ReadKeys`
* Validate that the called methods exist and have appropriate permissions for execution.
* If the called method is executed by the same processor
  * Return same_processor_execution = TRUE, output_argument = NULL.
* If the called method is executed by a different processor
  * Call a `Processor.ProcessCallSet` with the relevant method and retrun the output arguments.
