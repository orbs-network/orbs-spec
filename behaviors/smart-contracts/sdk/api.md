# Smart Contract SDK API

API specification for [smart contract](../../../terminology.md) developers. The API is provides by each `Processor`. Some of the functionality is implemented natively by the processor itself and some of the functionality (especially calls that integrate with the system around it) is provided by `VirtualMachine`.

&nbsp;
## Environment and arguments

#### Environment

```ts
Environment.GetVirtualChain(): Uint32
```

#### Execution arguments

&nbsp;
## Persistent state

Persistent state variables are available to the contract via the `State` object. Every [deployed service](../../../terminology.md) has its own isolated variable namespace (address space) that only it can write to.

#### Read

```ts
State.ReadBlobAddress(address: SHA256): Blob
State.ReadBlobKey(key: String): Blob
State.ReadStringAddress(address: SHA256): String
State.ReadStringKey(key: String): String
State.ReadUint64Address(address: SHA256): Uint64
State.ReadUint64Key(key: String): Uint64
```

#### Write

```ts
State.WriteBlobAddress(address: SHA256, value: Blob): Error
State.WriteBlobKey(key: String, value: Blob): Error
State.WriteStringAddress(address: SHA256, value: String): Error
State.WriteStringKey(key: String, value: String): Error
State.WriteUint64Address(address: SHA256, value: Uint64): Error
State.WriteUint64Key(key: String, value: Uint64): Error
```

#### Clear

```ts
State.ClearAddress(address: SHA256): Error
State.ClearKey(key: String): Error
```

#### Permissions
* Read does not require special permissions.
* Write and Clear can only be performed with `ReadWrite` permissions.

&nbsp;
## Calling other contracts

Multiple smart contracts can be deployed on one [virtual chain](../../../terminology.md). They can call each other. Calling a different [service](../../../terminology.md) will change the address space to that of the called service for the duration of the call (it will be able to access its own state variables). Calling a [library](../../../terminology.md) will keep the address space of the caller since libraries don't have their own address space.

#### Services

```ts
Service.IsDeployed(service: String): Boolean
Service.CallMethod(service: String, method: String, arguments: MethodCallArguments): MethodCallResult
```

#### Permissions
* Methods with `External` permissions can be called by anyone.
* Methods with `Internal` permissions can only be called by the service itself.
  * `System` permissions allow calling `Internal` methods of other services.

#### Libraries

```ts
Library.IsDeployed(library: String): Boolean
Library.CallMethod(library: String, method: String, arguments: MethodCallArguments): MethodCallResult
```

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
