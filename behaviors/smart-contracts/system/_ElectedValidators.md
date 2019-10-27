# `_ElectedValidators` System Service

> The system service charged with maintaining the current elected Orbs validators - matching the results of the Ethereum elections (bearing time phase shift).
> Implemented on the `Native` processor.

#### Permissions
* This is a system contract that runs under `System` permissions.


#### State
> Maintain the current set of elected Orbs validators, election number and dates range where this election holds.
* State variables:
    * Validators Record:
        * ValidatorsSet:
            * MemberID (OrbsAddress)
        * ElectionNumber
        * ValidUntil (Date)

&nbsp;
## `_init` (method)

#### Permissions
* `Internal` (caller must be the same service).
* `ReadWrite` (does not changes state).

#### Behavior
* Empty.



## `updateElectedValidators` (method)
> Support calls only from the `_Triggers` system contract.
> Update the state Validators Record if necessary (try to fetch newer data).
> The elected validators record is expired if no such record is yet stored or "ValidUntil" has passed (based on the block.timestamp).
> Dependent on `management.CallContract()` SDK call.

#### Permissions
* `External` (caller can only be `_Triggers` system.contract call).
* `ReadWrite` (might change state).

#### Behavior
* Check if should update the record:
    * Check if the elected validators record from state is expired:
        * The list of elected validators is nil.
        * The `block.timestamp` - current time reference > valid_until.
* If should update the record:
    * Try to retrieve an updated record by calling `management.CallContract('_ElectedValidators', 'getElectedValidators')` SDK call.
    * If call successfully resulted with a valid record, Store record in state.


## `getElectedValidators()` (method)
> Returns the elected validators record from state.
#### Permissions
* `External & Internal` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* Input: none.
* Returns from state the elected validators record := list of Validators, election counter, and date limit - they are still valid.
    * (node_address[], election_number, valid_until)
    
