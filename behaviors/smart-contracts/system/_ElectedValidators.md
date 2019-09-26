# `_ElectedValidators` System Service

> The system service charged with maintaining the current elected Orbs validators - matching the results of the Ethereum elections (bearing time phase shift).
> Implemented on the `Native` processor.

#### Permissions
* This is a system contract that runs under `System` permissions.


#### State
> Maintain the current set of elected Orbs validators, election number and dates range where this election holds.
* State variables:
    * Validators Record:
        * ValidatorsList:
            * MemberID (OrbsAddress)
        * ElectionNumber
        * HoldsFrom (Date)
        * HoldsUntil (Date)

&nbsp;
## `_init` (method)

#### Permissions
* `Internal` (caller must be the same service).
* `ReadWrite` (does not changes state).

#### Behavior
* Empty.



## `updateElectedValidators` (method)
> Dependent on `fetchElectedValidators()`.
> Support calls only from the `_Triggers` system contract.

#### Permissions
* `External` (caller can only be `_Triggers` system.contract call).
* `ReadWrite` (might change state).

#### Behavior
* Get the elected validators record by calling `fetchElectedValidators(block.timestamp)`
* Update the state Validators Record if necessary (fetched a newer data).


## `getElectedValidators()` (method)
> Returns the elected validators record stored in state.
#### Permissions
* `External & Internal` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* Input: none.
* Returns from state the elected validators record := list of Validators, election counter, and the dates range they hold.
    * (node_address[], election_number, holds_from, holds_until)
    


## `fetchElectedValidators()` (method
> Returns an up to date elected validators record.
> The elected validators record is expired if no such record is yet stored or the dates range (valid from - to) has passed (based on the provided reference timestamp).
> Dependent on `management.GetElectedValidators()` SDK call.
#### Permissions
* `External & Internal` (caller can be anyone).
* `ReadOnly` (does not change state).
#### Behavior
* Input: time reference.
* Returns an up to date elected validators record := list of Validators, election counter, and the dates range they hold.
    * (node_address[], election_number,holds_from, holds_until)
* Get the elected validators record from state by calling `getElectedValidators()`.
* Check if this record is expired:
    * The list of elected validators is nil.
    * The input time reference is not in dates range (holds_from, holds_until).
* If the record is expired fetch a new record by calling `management.GetElectedValidators(time_ref)` SDK call.
* Return the resulting record.