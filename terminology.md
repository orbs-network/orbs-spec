# Terminology

* `node` - Self contained entity in the network. Usually belongs to a single company or organization. Shares a private key and represented by a public key. The number of active nodes in the network is the number of active known public keys in the network. A node is made of multiple services potentially running on multiple machines. A node is a trust boundary. Different nodes do not trust each other. Different components inside a node do trust each other.

* `virtual chain` - An isolated virtual blockchain that can hold multiple deployed services (smart contracts) and their entire persistent state. Different virtual chains are isolated from one another, rely on different resources and thus can be scaled horizontally. A single node participates in the consensus process of multiple virtual chains (ideally all of them).

* `committee` - In the context of consensus, a committee is a group of nodes chosen to work together on closing a specific block in the chain. The committee is normally smaller than the total number of nodes in the network. Consensus algorithms such as Helix change committees on every block. Helix even randomizes the committee in an unpredictable way.

* `service` - This term is used in two contexts:

  * `micro service` - An architectural unit of software used to split a node to its basic building blocks. Each service has a clear responsibility over some part of node operation. The node code base is divided into independent sub projects implementing each of the services. Different services inside a node trust each other. If the node is compromised, we can assume all of its services are compromised. The public methods of a micro service can only be called by other micro services inside the node, they are not exposed externally to entities outside the node.

  * `deployed service` - A smart contract (deployed code which contains multiple methods) combined with its address space of state variables. Each service has its isolated state. A service represents an app deployed to the network.

* `smart contract` - Code that is deployed by an app developer and executes on the virtual machine of the node. The main purpose of the network is to execute these smart contracts under consensus. In other words, a smart contract is code that is agreed upon in advance that controls the engagement between parties. A smart contract is made of multiple methods. Smart contract can be a service (self contained app) or a library (utility code that has no address space of its own).

* `library` - A smart contract (deployed code which contains multiple methods) that does not have an address space of its own and provides shared utilities to other smart contracts who can execute it. If it manipulates state, it will change the state of the deployed service calling it.

* `block height` - The block number in the blockchain. Since the blockchain is made of a linked list of blocks, each block has a sequence number. The genesis block on virtual chains is empty, its block height is `0`. This means the first meaningful block is numbered `1`.

* `ordering and validation` - The consensus process is separated into two phases - ordering and validation. During the ordering phase, transactions are chosen for the block and ordered. During the validation phase, transactions are executed and marked as valid/invalid with their results.
  * `Transactions block` - The output of the ordering phase, a block of ordered transactions that is committed to the blockchain in a specific height.
  * `Results block` - The output of the validation phase, a block of transaction execution results that is committed to the blockchain in a specific height (matching its Transactions block, but a different block nevertheless).

* `gateway node` - The node in the network that a client communicated with. If the client sends a transaction, this is the node that first introduces this transaction to the network (all other nodes). Gateway nodes are accountable for their clients and should prevent them from spamming the virtual chain.
