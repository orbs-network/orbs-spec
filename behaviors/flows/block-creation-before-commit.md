# Continuous Block Creation Flow

> First half: before commit (steps 1-3 out of 5)

*Since the flow is long, this is the first half of the flow: all the way until the block is committed.*

## Flow

#### Step 1: round starts

* `ConsensusAlgo` of all nodes starts a new consensus round (its [block height](../../terminology.md) is known).
  * The previously committed block is known from the previous round or read from `BlockStorage`.
  * Gets the block's committee from `ConsensusContext`.
  * Learns if they are a leader, non-leader committee member or not a committee member.

#### Step 2: leader proposes a new block

* `ConsensusAlgo` of leader node:
  * Asks `ConsensusContext` to build new Transactions and Results block proposals.

  * `ConsensusContext` of leader node:
    * Requests pending transactions for the Transactions block from `TransactionPool`.

    * `TransactionPool` of leader node:
      * Executes pre order checks by calling `VirtualMachine`.
      * Similar to the checks the gateway node did when adding each transaction to the network.

       * `VirtualMachine` of leader node:
         * Executes the subscription check smart contract on the native `Processor`.
         * Depending on contract code may reads state from `StateStorage` or `CrosschainConnector`.

    * Requests root hash of the state prior to execution for the Results block from `StateStorage`.
    * Executes all transactions for the Results block by calling `VirtualMachine`.

      * `VirtualMachine` of leader node:
        * Depending on contract code may reads state from `StateStorage` or `CrosschainConnector`.
        * Generates transaction receipts and state diff.

  * Sends the block proposals to all committee members for validation with `Gossip`.

#### Step 3: committee members (without leader) validate the proposal

* `ConsensusAlgo` of non-leader committee member:
  * Asks `ConsensusContext` to validate the content of the block proposal.

  * `ConsensusContext` of non-leader committee member:
    * Validates the Transactions block proposal with `TransactionPool`.

    * `TransactionPool` of non-leader committee member:
      * Executes pre order checks by calling `VirtualMachine`.
      * Similar to the checks the gateway node did when adding each transaction to the network.

       * `VirtualMachine` of non-leader committee member:
         * Executes the subscription check smart contract on the native `Processor`.
         * Depending on contract code may reads state from `StateStorage` or `CrosschainConnector`.

    * Validates root hash of the state prior to execution in the Results block from `StateStorage`.
    * Validates execution of all transactions in the Results block by calling `VirtualMachine`.

      * `VirtualMachine` of non-leader committee member:
        * Depending on contract code may reads state from `StateStorage` or `CrosschainConnector`.
        * Generates transaction receipts and state diff for comparison.
