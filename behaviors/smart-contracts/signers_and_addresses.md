# Signers and Addresses
> A signer scheme determines the way a transaction issigned and the mapping to a public address.

## Addresses and accounts
In Orbs platform there is no notation of a native account only of an address. Accounts or other signature based databases can be implemented by using the public address as the key for the database. 

## SDK Functions

#### `GetSignerAddress()`
> Returns the address of transaction signer. 

#### `GetCallerAddress()`
> Returns the address of the function caller.
* If the caller is a transaction, then CallerAddress = SignerAddress
* If the caller is a smart contract then CallerAddress = `SmartContractCaller` address.
* If the caller is a system call then CallerAddress = `SmartContractCaller` address, with contract_name = "system"

#### `CheckValidAddresFormat(Address)`
> Performs static checks on an address argument to validate that it has a valid format based on the scheme. 
* Note: A valid format does not indicate that the address corresponds to a valid public key.
* Verifies network_type matches teh network.

#### `GetAddressScheme(Address)`
> Returns the address scheme used by the address.

<!-- TBD 
#### `GetSigner()`
#### `IsSignerValid()`
#### `VerifyNetworkType(Address)`
---> 

## Usage Examples
#### Token get_my_balance()
Return BalancesDB[`GetCallerAddress()`]

#### Token transfer
If `token` <= BalancesDB[`GetCallerAddress()`] then
  BalancesDB[target] += tokens
  BalancesDB[`GetCallerAddress()`] -= tokens


## Signature schemes
> Determines the signature validation and addressing scheme
* Address = {scheme, network_type, RIPEMD160(SHA256(signer))}

#### `EdDSA01Signer`
> Ed25519 based signature scheme
* Signature = Ed25519Signature(private_key, txhash)
  * txhash = SHA256(Transaction) <!-- TBD - should we align txhash not to include the signature?  
* Signer checks
  * network_type matches the network
  * Valid signature

#### `SmartContractCaller`
> Set when a function is called by a smart contract, can't be sent in a transaction

#### `SystemCaller`
> Set when a function is called by a system, can't be sent in a transaction
* When a call was initiated by a system call, `GetSignerAddress()` return a `SystemCaller` scheme.

## Text encoding scheme
1. Start with a 22-byte address:
  * Network ID: Address[21]
  * Account ID: Address[20-0]
   
2. Calculate the CRC32 checksum of the address. Construct the 26-byte binary representation of the address {Network ID, Account ID, checksum}
   
3. Encode the raw public address to Base58:
    Each of the following address parts is BASE58 encoded separately:
    a. Network ID
    | Network  | Value | Value (hex) |
    | :------: | :---: | :---------: |
    | Main net | M     | 4d          |
    | Test net | T     | 54          |
    
    b. {Account ID, Checksum}
    
