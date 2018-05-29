# Processor

> Executes a stream of trasnactions in an isolated enviroment.

* Permission is managed by the VM oblivious to the the processor.
* State access read / write is done calls to the VM Sdk.
  * No cahcing of local or state variables between transactions' processing.
* It is recommanded to provide the Processor with a stream of a set of transactions in order to avoid the processor overhead for each transaction.

&nbsp;
![alt text][processor_interface] <br/><br/>

[processor_interface]: ../_img/processor_interface.png "VM - Processor Interface"

&nbsp;
## `Processors`

#### `Native`
* Hard coded in code base and updated through node version updates.
* Implemented in native language of the node for efficiency and performance.
* Name space code is stored in a registry which is loaded on initialization and cached by the processor. //TODO.
* See API [specification](../smart-contracts/native.md).

&nbsp;
## `ProcessCallSet` (method)

> Executes an ordered stream of transactions of a single context and generates output arguments for each transaction.

* Receives a namespace and a method to execute along with a list of input arguments.
* Retrieves the method code and permissions, if not found, fail.
  * TODO If the code is deployable, do this by calling a dedicated function of VirtualMachine.
* Verify that the method has appropriate permissions for execution:
  * Read/Write permissions match the `access_scope` argument.
  * Internal/External permissions match the `calling_namespace` argument.
  * On mismatch fail.
* Execute the code on the processor.
  * SDK calls are done by blocking execution and calling `VirtualMachine.SdkCall`.
