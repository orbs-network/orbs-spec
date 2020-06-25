# Vault
> Note that the specification updates the virtual chain signature scheme hence is not backward compatible and requires protocol number change.

The vault is a service that holds the node private key and allows to sign certain payload with this key.

It runs as a separate process in a separate container. The idea behind it is secure isolation of private key from other parts of the system. The vault verifies a virtual chains identity and signs a virtual chain data only by the right virtual chain.

## Interacts with services

* `Consensus Algo` - uses it to sign blocks.
* `Transaction Pool` - uses it to sign transaction forwarding.
* `Block Storage` - uses it to sign block sync messages.

## APIs

#### Signing service

* `NodeSign` - transforms byte array payload into SHA256 hash and signs it with EcdsaSecp256K1.
* `EthSign` - transforms byte array payload into SHA3 hash and signs it with EcdsaSecp256K1.

#### Secure private key generation

* `NodeGenerateNewKeyPair` - generates a new key pair
* `NodeGetPublicKey` - returns the node's public key

#### Virtual chain identity management

* `RegisterVirtualChain` - (privileged : vault owner) registers a new virtual chain


## Provisioning

Decoupling of signer service from the main process and running a single instance per node (X Vchain containers + 1 signer container).

The vault should be as independent as possible in the protocol revision, preventing need to modify it on protocol upgrade.

&nbsp;
## `NodeSign`
> Sign virtual chain data using the node key pair

#### Check virtual chain local signature
* Verify that the local signature matches local_address(virtual_chain_id)
  * Panic on mismatch or virtual_chain_id not registered.

#### Sign data
* Signature(data_hash, virtual_chain_id) = Ecdsa_signature(H({data_hash, virtual_chain_id}))

&nbsp;
## `NodeGetNodeAddress`
> Returns the node public address

&nbsp;
## `NodeGenerateNewKeyPair`
> Generates a new key pair for the node, overwriting the current key-pair.
> Should be performed only upon node setup, requires to update the Validators registration contract.

* The private key should be stored in a secure, isolated manner. 
* The private key can not be read by any external service / entity.

## `RegisterVirtualChain`
> Registers a new virtual, setting its local public address. 

&nbsp;
#### Check the owner signature
* Verify that the owner signature
  * Panic on mismatch

#### Register virtual chain
* Stores the virtual_chain_id local public address, update if already exist

&nbsp;
## Public Virtual chain signature verification scheme.

#### Protocol version 0 scheme
* If protocol version = 0
  * Verify(data_hash, public_key, virtual_chain_id) = Ecdsa_verify(data_hash)

* If protocol version = New protocol verion
  * Verify(data_hash, public_key, virtual_chain_id) = Ecdsa_verify(H({data_hash, virtual_chain_id}))

## Implementation phase 0 - `NodeSign`
> In implementation phase 0, virtual chain ID is not inspected maintaining protocol version 1 signature scheme.

#### Sign data
* Signature(sha256_data_hash) = Ecdsa_signature(sha256_data_hash)

## Implementation phase 1 - `EthSign`

#### EthSign data
* Signature(sha3_data_hash) = Ecdsa_signature(sha3_data_hash)