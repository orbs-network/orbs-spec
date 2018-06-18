# Continuous Block Creation Flow

> First half: before commit (steps 1-3 out of 5)

*Since the flow is long, this is the first half of the flow: all the way until the block is committed.*

## Flow

#### Step 1: round starts

* `ConsensusAlgo` of all nodes:
  * Starts a new consensus round (its [block height](../../terminology.md) is known).
  * The previously committed block is known from the previous round or read from `BlockStorage`.
  * Gets the block's committee from `ConsensusContext`.
  * Learns if they are a leader, non-leader committee member or not a committee member.
    * The way roles are chosen is an internal implementation detail of the specific `ConsensusAlgo`.
    * Roles (like leader) may change mid-round as part of `ConsensusAlgo` (due to errors, timeouts and such).
    * The flow presented here is a happy flow where roles aren't required to change mid-round.

#### Step 2: leader proposes a new block

* `ConsensusAlgo` of leader node:
  * Asks `ConsensusContext` to build new [Transactions and Results block](../../terminology.md) proposals.

  * [Ordering](../../terminology.md) phase:

  * `ConsensusContext` of leader node:
    * Requests pending transactions for the Transactions block from `TransactionPool`.

    * `TransactionPool` of leader node:
      * Executes pre order checks by calling `VirtualMachine`.
        * Note: Similar to the checks the gateway node did when adding each transaction to the network.

      * `VirtualMachine` of leader node:
        * Executes the subscription check smart contract on the native `Processor`.
        * Depending on contract code may reads state from `StateStorage` or `CrosschainConnector`.

  * [Validation](../../terminology.md) phase:

  * `ConsensusContext` of leader node:
    * Requests root hash of the state prior to execution for the Results block from `StateStorage`.
    * Executes all transactions for the Results block by calling `VirtualMachine`.

      * `VirtualMachine` of leader node:
        * Depending on contract code may reads state from `StateStorage` or `CrosschainConnector`.
        * Generates transaction receipts and state diff.

  * Sends the block proposals to all committee members for validation with `Gossip`.

#### Step 3: committee members (without leader) validate the proposal

* `ConsensusAlgo` of non-leader committee member:
  * Asks `ConsensusContext` to validate the content of the Transactions and Results block proposal.

  * Ordering phase:

  * `ConsensusContext` of non-leader committee member:
    * Validates the Transactions block proposal with `TransactionPool`.

    * `TransactionPool` of non-leader committee member:
      * Executes pre order checks by calling `VirtualMachine`.
        * Note: Similar to the checks the gateway node did when adding each transaction to the network.

      * `VirtualMachine` of non-leader committee member:
        * Executes the subscription check smart contract on the native `Processor`.
        * Depending on contract code may reads state from `StateStorage` or `CrosschainConnector`.

  * Validation phase:

  * `ConsensusContext` of non-leader committee member:
    * Validates root hash of the state prior to execution in the Results block from `StateStorage`.
    * Validates execution of all transactions in the Results block by calling `VirtualMachine`.

      * `VirtualMachine` of non-leader committee member:
        * Depending on contract code may reads state from `StateStorage` or `CrosschainConnector`.
        * Generates transaction receipts and state diff for comparison.


&nbsp;
## Block Creation Flow Before Commit, Leader Node

![alt text][block_creation_before_commit_leader] <br/><br/>

[block_creation_before_commit_leader]: ../_img/block_creation_before_commit_leader.png "block_creation_before_commit_leader"

&nbsp;
## Block Creation Flow Before Commit, Non-leader Node

![alt text][block_creation_before_commit_non_leader] <br/><br/>

[block_creation_before_commit_non_leader]: ../_img/block_creation_before_commit_non_leader.png "block_creation_before_commit_non_leader"
