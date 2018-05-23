# Virtual Machine

Executes smart contracts using various processors (languages) and produces state difference as a result.

Currently a single instance per virtual chain per node.

&nbsp;
## `Processors`

* Native Go - set of native smart contracts included in node core.

&nbsp;
## `Data Structures`

#### Transient state cache
* Stores temporary state changes during transaction set execution (until state diff is committed).
* Format is identical to the state store data structure in `StateStorage`.
* No need to be persistent since a new instance is created per execution.
* No limit on max size.

&nbsp;
## `Signature Schemes`

#### Address Scheme 01
* Check Ed25519 signature over the transaction header.

&nbsp;
## `CallContract` (method)
> Execute a read only method in a native contract and return its result

#### Prepare for execution
* Retrieve the contract from the native contract registry.
* Validate that the method exists and has appropriate permissions for execution.

#### Execute
* Validate the transaction signature by performing: `Sender and Signature Validation`.
* Execute the contract method and return its result.
* Contract code cannot update state (state writes must throw an error).

&nbsp;
## `ProcessTransactionSet` (method)
> Process a group of transactions together that update state and return the combined state diff

#### Prepare for execution
* Prepare transient state cache that will hold updated state temporarily (across all transactions).
* Retrieve contracts from the native contract registry.
* Validate that the methods exist and have appropriate permissions for execution.

#### Execute
* Contract code can only update state within the transient state cache.
* Execute the contract methods and return their results.

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
