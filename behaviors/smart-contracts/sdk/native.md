# Native Smart Contracts

API specification for native smart contracts implemented as part of node core for efficiency and performance.

&nbsp;
## `state.load(key: string): Promise<string>` (method)
> Reads a value from the smart contract state by key

* Enforce permissions to make sure read is allowed (todo).
* State read is implemented by calling `StateStorage.ReadKeys`.
* Reads are cached in the transient state cache to improve performance and observe previous writes.

&nbsp;
## `state.store(key: string, value: string): Promise<string>` (method)
> Writes a value to the smart contract state by key

* Enforce permissions to make sure write is allowed (todo).
* State write is implemented over the transient state cache only.
* State diff is only actually committed to state storage after consensus.

&nbsp;
## `senderAddressBase58` (property)
> Address of the caller of the smart contract
