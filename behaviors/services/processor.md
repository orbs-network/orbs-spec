# Processor
> Executes a stream of trasnactions in an isolated enviroment. 
* Permission is managed by the VM oblivious to the the processor.
* State access read / write is done calls to the VM Sdk.
  * No cahcing of local or state variables between transactions' processing.
* It is recommanded to provide the Processor with a stream of a set of transactions in order to avoid the processor overhead for each transaction.

&nbsp;
![alt text][processor_interface] <br/><br/>

[processor_interface]: ../_img/processor_interface.png "VM - Processor Interface"

## Native Processor
* Hard coded in code base and updated through node version updates.
* Implemented in native language of the node for efficiency and performance.
* Name space code is stored in a registry which is loaded on initialization and cached by the processor. //TODO.
* See API [specification](../smart-contracts/native.md).

&nbsp;
## `ProcessCallSet` (method)

> Executes an ordered stream of transactions of a single context and generates output arguments for each transaction.

* Receives a name space and a method to execute along witha list of input arguments.
* Reterives the name space code for execution by calling `VM.GetNameSpaceCode`.

#### Calling SDK API functions
* Calls `VM.SdkCall` providing a list of input arguments and receiving a list of output arguments.
* Holds progress until receiving a response.

#### Calling another service or library method
* Calls `VM.MethodCall` providing a list of input arguments and receiving a list of output arguments.
* Holds progress until receiving a response.
* The service can callback a function on the service / library that called it by calling `VM.MethodCall` with name 