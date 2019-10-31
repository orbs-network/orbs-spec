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
        * NextSync (Date)

* constants:
    * ELECTED_VALIDATORS_MAX_NEXT_SYNC : time cap of next sync step - relative to current timestamp. 

&nbsp;
## `_init` (method)

#### Permissions
* `Internal` (caller must be the same service).
* `ReadWrite` (does not changes state).

#### Behavior
* Empty.



## `sync` (method)
> Support calls only from the `_Triggers` system contract.
> Update the state Validators Record if necessary (try to fetch newer data).
> The elected validators record is expired if "NextSync" has passed (based on the block.timestamp).
> Dependent on `management.CallContract()` SDK call.

#### Permissions
* `External` (caller can only be `_Triggers` system.contract call).
* `ReadWrite` (might change state).

#### Behavior ConsumerVC
* Check if should update the record:
    * Check if the elected validators record from state is expired:
        * The `block.timestamp` - current time reference > `NextSync`.
* If should update the record:
    * Try to retrieve an updated record by calling `management.CallContract('_ElectedValidators', 'getElectedValidators')` SDK call.
    * If call successfully resulted with a valid record:
        * If the record holds a newer next_sync timestamp
            * Cap `NextSync` based on ELECTED_VALIDATORS_MAX_NEXT_SYNC.
            * Store record in state.

#### Behavior ManagementVC
> If should update (election event occurred), process new election data.
> At the end of the processing stores a new elected validators record in state with next_sync := next election time.

#### Behavior PrivateVC
> Does not update state record.



## `getElectedValidators()` (method)
> Returns the elected validators record from state.
#### Permissions
* `External & Internal` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* Input: none.
* Returns from state the elected validators record := list of Validators, election counter, and next sync date.
    * (node_address[], election_number, NextSync)
    
    

## `getElectedValidatorsNextSync()` (method)
> Returns the elected validators next sync timestamp from state.
#### Permissions
* `External & Internal` (caller can be anyone).
* `ReadOnly` (does not change state).
#### Behavior
* Input: none.
* Returns `NextSync`
