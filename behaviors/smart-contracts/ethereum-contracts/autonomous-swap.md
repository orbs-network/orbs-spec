# Autonomous Swap Contarct

> The Autonomous Swap Contract is responsible to transfering tokens to and from Orbs.

> Implemented on the `Native` processor.

## General 

#### Permissions
* This is a pulic contract that any client can call its public functions.

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
* `TransferredOut(uint256 indexed tuid, address indexed from_ethereum_address, bytes20 indexed to_orbs_address, uint256 value)`
* `TransferredIn(uint256 indexed tuid, bytes20 indexed from_orbs_address, address indexed to_ethereum_address, uint256 value)`

&nbsp;
## `_Init_` (method)

#### Permissions
* `Internal` (System call)
* `ReadWrite` (potentially changes state).

#### Behavior
* Input arguments:
  * o
#### 

&nbsp;
## `TransferOut` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite` (potentially changes state).

#### Behavior
* Input arguments:
  * orbs_address : bytes20
  * amount : uint256
* Call the ERC20 account and transfer tokens from caller to the ASB.
* Emit `TransferredOut`

&nbsp;
## `TransferIn` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadWrite` (potentially changes state).

#### Behavior
* Input arguments:
  * proof - packed proof
  * receipt - transaction receipt
* Get the `OrbsProofValidator` address based on the proof revision.
  * On an unknown/invalid proof reveision fail.
  * Validate the proof by calling `OrbsProofValidator.ValifateReceiptProof(proof)`
* Check proof.receipt_hash = SHA(Receipt)
* Extract the event data from the receipt
  * Check Event.contract_name matches the configured ASB contract.
  * Unpack the event data
    * Check event_id = TRANSFERED_OUT (0x1)
    * Check that the etehreum_address is a valid etehreum_address.
* Check the event's tuid wasn't already spent.
* Call the ERC20 account and transfer tokens to the etehreum_address.
* Emit `TransferredIn`

&nbsp;
## `SetOrbsASBContract` (method)

#### Permissions
* `Owner` (external, priviliged to the owner).
* `ReadWrite` (does not change state).

#### Behavior
* Set the Orbs corresponding ASB contract name.

&nbsp;
## `Deprecate` (method)

#### Permissions
* `Owner` (external, priviliged to the owner).
* `ReadWrite` (does not change state).

#### Behavior
* Input
  * New contract address
* Check that the address is a contarct (Sanity)
* Transfers the tokens to the new contract

## TBDs
* Open/Close bridge
* Add contracts registery support