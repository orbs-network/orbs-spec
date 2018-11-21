# Smart Contract SDK API

API specification for [smart contract](../../../terminology.md) developers. The API is provides by each `Processor`. Some of the functionality is implemented natively by the processor itself and some of the functionality (especially calls that integrate with the system around it) is provided by `VirtualMachine`.

&nbsp;
## Environment and arguments

#### Environment

```ts
Environment.GetVirtualChain(): Uint32
```

#### Execution arguments

#### Block properties

```ts
Block.GetTime(): Uint64
```

&nbsp;
## Persistent state

Persistent state variables are available to the contract via the `State` object. Every [deployed service](../../../terminology.md) has its own isolated variable namespace (address space) that only it can write to.

Raw state addresses are of type `Ripmd160Sha256` which is a double hash of `RIPEMD160(SHA256(data))`. From an API perspective, this is just a typedef to an array of bytes with the expected length of 20.

#### Read

```ts
State.ReadBytesByAddress(address: Ripmd160Sha256): Bytes
State.ReadBytesByKey(key: String): Bytes
State.ReadStringByAddress(address: Ripmd160Sha256): String
State.ReadStringByKey(key: String): String
State.ReadUint64ByAddress(address: Ripmd160Sha256): Uint64
State.ReadUint64ByKey(key: String): Uint64
```

#### Write

```ts
State.WriteBytesByAddress(address: Ripmd160Sha256, value: Bytes): Error
State.WriteBytesByKey(key: String, value: Bytes): Error
State.WriteStringByAddress(address: Ripmd160Sha256, value: String): Error
State.WriteStringByKey(key: String, value: String): Error
State.WriteUint64ByAddress(address: Ripmd160Sha256, value: Uint64): Error
State.WriteUint64ByKey(key: String, value: Uint64): Error
```

#### Clear

```ts
State.ClearByAddress(address: Ripmd160Sha256): Error
State.ClearByKey(key: String): Error
```

#### Permissions
* Read does not require special permissions.
* Write and Clear can only be performed with `ReadWrite` permissions.

&nbsp;
## Calling other contracts

Multiple smart contracts can be deployed on one [virtual chain](../../../terminology.md). They can call each other. Calling a different [service](../../../terminology.md) will change the address space to that of the called service for the duration of the call (it will be able to access its own state variables). Calling a [library](../../../terminology.md) will keep the address space of the caller since libraries don't have their own address space.

#### Services

```ts
Service.CallMethod(service: String, method: String, arguments: MethodCallArguments): MethodCallResult
```

#### Permissions
* Methods with `External` permissions can be called by anyone.
* Methods with `Internal` permissions can only be called by the service itself.
  * `System` permissions allow calling `Internal` methods of other services.

#### Libraries

```ts
Library.CallMethod(library: String, method: String, arguments: MethodCallArguments): MethodCallResult
```

&nbsp;
## Addresses

System addresses are of type `Ripmd160Sha256` which is a double hash of `RIPEMD160(SHA256(data))`. From an API perspective, this is just a typedef to an array of bytes with the expected length of 20. These addresses may be used to represent accounts and can be derived from transaction signers. The raw state database also uses this addressing format for its keys.

#### Validating addresses

```ts
Address.ValidateAddress(address: Ripmd160Sha256): Error
```

#### Querying address of currently running contract

```ts
Address.GetSignerAddress(): Ripmd160Sha256
Address.GetCallerAddress(): Ripmd160Sha256
```

The separation between the two methods is due to contracts calling other contracts. **Signer** address returns the address of the signer of the original transaction accross all contract calls triggered by this transaction. **Caller** address returns the address of the calling contract (which equals to the signer address for the contract specified on the transaction).

Token contracts for example should only use **Caller** address to prevent other contracts calling them on behalf of users and transferring unauthorized amounts.

&nbsp;
## Crypto operations

&nbsp;
## Cross-chain access

Traditionally, smart contracts can only access state variables from the blockchain they're running on. Cross-chain allows a smart contract to access smart contracts and variables on a different blockchain. These calls are still performed under consensus.

#### Ethereum

```ts
Ethereum.CallContract(...)
```

&nbsp;
## System operations

Special operations that can only be called from system contracts (a special permission).

&nbsp;
## Event logs

The event logs are logged as part of the transaction receipt. 

```ts
Log.EmitEvent(arguments: EventArguments)
```
