# Consensus Algo

Implements a raw consensus algorithm. The system supports pluggable consensus, meaning multiple consensus algorithms running side by side on different virtual chains. The ConsensusAlgo abstracts away the details of how the consensus algorithm works and provides the system with a single unified interface.

Communicates with the system through the ConsensusCore service. Controls all the timing of consensus, for example when new blocks are committed or when new blocks need to be proposed and populated with transaction from TransactionPool.

Currently a single instance per virtual chain per node.

&nbsp;
## `Init Flow`
* Subscribe to consensus messages.
* Init view, term to 0.

&nbsp;
## `OnNewConsensusRound` (method) <!-- tal can finish -->

> The consensus algorithm decides a new term (block height) starts, probably after the last block was committed.

* Assumes the block height for the upcoming round is known.
* Calculate the shared `random_seed` for the upcoming block.
* Get a sorted list of committee members for the upcoming block by calling `ConsensusCore.RequestCommittee`.

#### If Leader
* Request new transactions block proposal (ordering phase) by calling `ConsensusCore.RequestNewTransactionsBlock`.
* Request new results block proposal (execution phase) by calling `ConsensusCore.RequestNewResultsBlock`.
* Sign both proposals (according to the algo spec) - on hash of the header.
* Send signed proposals with gossip to the algo of all committee nodes (according to the algo spec).

&nbsp;
## `OnBlockProposalReceived` (method) <!-- tal can finish -->

> Through the consensus algorithm, a non-leader node receives a proposal for a new block by the leader.

* Make sure the block height matches the current round, if not, don't actively participate.
* Perform algorithm related checks on the proposal, signature included.
* Validate the transactions block (ordering phase) by calling `ConsenusCore.ValidateTransactionsBlock`.
* Validate the results block (execution phase) by calling `ConsenusCore.ValidateResultsBlock`.
* Sign both approvals (according to the algo spec) - on hash of the header.
* Send signed approvals with gossip to the algo of all committee nodes (according to the algo spec).

&nbsp;
## `OnCommitBlock` (method) <!-- tal can finish -->

> The consensus algorithm has decided a block proposal can be committed (all nodes, not just committee).

* Ignore whether the block height matches the current round, follow the steps anyways.
* Assumes the transactions block and results block were already propagated.
* Pass the committed block's proofs by calling `ConsensusCore.CommitBlock`.

&nbsp;
## `GossipMessageReceived` (method)

> Handles messages received from another node, expect to receive only Consensus messages that cosnesus-algo have subscribed to.
