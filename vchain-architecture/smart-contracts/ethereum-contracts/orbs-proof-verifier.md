# Orbs Receipt Proof Validator

> The Orbs Recceipt proof validator receives a packed Orbs proof amd a valdiates it. 

## General 

#### Owner
* The owner of the contract is the Federation contract

#### Dependencies
* Federation contract

#### Owner and upgrade methodology
* 


&nbsp;
## `_Init_` (method)

#### Permissions
* `Internal` (System call)
* `ReadWrite` (potentially changes state).

#### Behavior
* Inputs:
  * Orbs Federation Contract address
  * Proof revision (0x1)
  
#### 

&nbsp;
## `ValifateReceiptProof` (method)

#### Permissions
* `External` (caller can be anyone).
* `ReadOnly` (does not change state).

#### Behavior
* Validates the proof based on [ethereum-receipt-proof](../data-structures/ethereum-receipt-proof.md)


&nbsp;
## `SetFederationContarct` (method)

#### Permissions
* `Owner` (external, priviliged to the owner).
* `ReadWrite` (does not change state).

#### Behavior
* Set the Federation contract address.
* Check that the address is a contarct (Sanity)
