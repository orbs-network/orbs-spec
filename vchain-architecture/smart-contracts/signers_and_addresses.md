# Signers and Addresses
> A signer scheme determines the way a transaction is signed and the mapping to a public address.

## Addresses and accounts
Orbs platform does not define native accounts only public address. Accounts or other signature based databases can be implemented using the public address as the key for the database. 

&nbsp;
## SDK Functions

#### `GetSignerAddress()`
> Returns the address of transaction signer, based on the Signer scheme.

#### `GetCallerAddress()`
> Returns the address of the function caller.
* If the caller is a transaction, return `GetSignerAddress()`.
* If the caller is a smart contract, return the `SmartContractCaller` address.
* If the caller is a system call, return the `SmartContractCaller` address, with contract_name = "system".

#### `CheckValidAddresFormat(Address)`
> Performs static checks on an address argument to validate that it has a valid format based on the scheme. 
* Note: A valid format does not indicate that the address corresponds to a valid public key.
* Verifies address length = 20 bytes.

&nbsp;
## Usage Example
#### Transfer(target, tokens)
```
If `tokens` <= BalancesDB[`GetCallerAddress()`] then
  BalancesDB[`GetCallerAddress()`] -= `tokens`
  BalancesDB[target] += `tokens`
```

&nbsp;
## Signature schemes
> Determines the signature validation and addressing scheme
* Address = RIPEMD160(SHA256(serialized signer))
  
#### `EdDSA01Signer`
> Ed25519 based signature scheme
* Signature = Ed25519Signature(private_key, txhash)
  * txhash = SHA256(Transaction)  

#### `SmartContractCaller`
> Set when a function is called by a smart contract, can't be sent in a transaction.
* The Signer includes the smart contract name and scheme.
<!-- TODO caller argument -->

&nbsp;
## Text encoding scheme
1. Start with a 20-BYTE address.
2. Calculate the CRC32 checksum of the address. Construct the 24-byte binary representation of the address {Address, checksum}
3. Encode the raw public address to BASE58.
