# Orbs Federation Contract
> The fedetaion contract holds the federarion members list and the contracts registery.

## Federation Members

## Contracts Registery
<!-- TBD swithc to a more generic format of {"name", address} pair so the federation contract does not need to be upgraded -->

## General 

#### Owner
* The owner of the contract is the Federation contract

#### Dependencies
* Proof validation contract
* ERC20 contract
* Implicit: Orbs ASB contract

#### Owner and upgrade methodology
* 

&nbsp;
## Events
* `MemberAdded(address indexed member)`
* `MemberRemoved(address indexed member)`

&nbsp;
## `_Init_` (method)

#### Permissions
* `Internal` (System call)
* `ReadWrite` (potentially changes state).

#### Behavior
* Input arguments:
  * Federation members list (addresses)
* Verifies and assigns an initial memebrs list 
#### 

&nbsp;
## Query members functions


### `isMember` (method)
<!-- TODO revisions -->
#### Permissions
* `External` (caller can be anyone).
* `ReadOnly`

#### Behavior
* Input arguments:
  * Ethereum address of the a potential member
* Returns a bool indication on whether the address belongs to a federation member.


### `getMembers` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly`

#### Behavior
* Input arguments: N/A
* Returns a list of the federation members addresses.
  * Based on the latest revision.


### `getConsensusThreshold` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly`

#### Behavior
* Input arguments: N/A
* Returns a the threshold for members signatures for consensus. (used by the proof validiator)
  * Based on the latest revision.


&nbsp;
## Query members by revision functions


### `getFederationRevision` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly`

#### Behavior
* Input arguments: N/A
* Returns the federation revision.


### `getMembersByRevision` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly`

#### Behavior
* Input arguments: 
  * federation_revision
* Returns a list of the federation members addresses based on the requested revision.


### `getConsensusThresholdByRevision` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly`

#### Behavior
* Input arguments: 
  * federation_revision
* Returns a the threshold for members signatures for consensus based on the requested revision. (used by the proof validiator)


### `isMemberByRevision` (method) 
<!-- not implemented -->

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly`

#### Behavior
* Input arguments:
  * Ethereum address of the a potential member
* Returns a bool indication on whether the address belongs to a federation member based on the reveision.


&nbsp;
## Members removal and addition

### `addMember` (method) 

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite`

#### Behavior
* Input arguments:
  * Ethereum address of a new member
* Validates not already a member
* Stores a new revision of members
* Emits `MemberAdded`


### `removeMember` (method) 

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite`

#### Behavior
* Input arguments:
  * Ethereum address of a removed member
* Validates that the address is a member
* Stores a new revision of members
* Emits `MemberRemoved`


&nbsp;
## Contracts registery


