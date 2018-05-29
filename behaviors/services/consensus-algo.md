# Consensus Algo

Implements a raw consensus algorithm. The system supports pluggable consensus, meaning multiple consensus algorithms running side by side on different virtual chains. Abstracts away the details of how the consensus algorithm works and provides the system with a single unified interface.

Controls and drives the timing of consensus, for example when new blocks are committed or when new blocks need to be proposed and populated with transaction from the pool.

Currently a single instance per virtual chain per node (per supported algorithm).

&nbsp;
## `Init` <!-- oded will finish -->

TODO

&nbsp;
## `OnNewConsensusRound` (method) <!-- tal can finish -->

> The consensus algorithm decides a new term (block height) starts, probably after the last block was committed.

* Assumes the block height for the upcoming round is known.
* Get the previous committed block pair data by calling `BlockStorage.GetTransactionsBlockHeader` and `BlockStorage.GetResultsBlockHeader`.
  * It is recommanded to cache the previously committed block data in order to prevent fetching the data.
* Calculate the shared `random_seed` for the upcoming block. Use the same `random_seed` for both committees.
* Get a sorted list of committee members for the upcoming transaction block (ordering phase) by calling `ConsensusBuilder.RequestOrderingCommittee`.
* Get a sorted list of committee members for the upcoming results block (execution validation phase) by calling `ConsensusBuilder.RequestValidationCommittee`.
* Note: If the consensus algorithm relies on a single committee for both, call `RequestValidationCommittee` only.

#### If Leader
* Request new transactions block proposal (ordering phase) by calling `ConsensusBuilder.RequestNewTransactionsBlock`.
* Request new results block proposal (execution phase) by calling `ConsensusBuilder.RequestNewResultsBlock`.
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
## `Verifying that a block can commmitted`

#### Validate that the block pair can be committed
* Call `ValidateTransactionsBlockConsensus`
* Call `ValidateResultsBlockConsensus`

&nbsp;
## `OnCommitBlock` (method) <!-- tal can finish -->

> The consensus algorithm has decided a block proposal can be committed (all nodes, not just committee).

* Pass the committed block to the block storage by calling `BlockStorage.CommitBlock`.

&nbsp;
## `GossipMessageReceived` (method)

> Handles messages received from another node, expect to receive only Consensus messages that cosnesus-algo have subscribed to.

* Depends on consensus algorithm.

&nbsp;
## `ValidateTransactionsBlockConsensus` (method)
> Validates that a tarnsactions block header can be committed assuming that the block body (content) is valid and its relevant hash values match the ones in the header and that the previous block was successfully committed.

#### Verify the block committee
* Calualte the previous block random seed (based on the results block proof - for both transactions and results blocks)
* Get the committee members by calling `ConsensusBuilder.RequestValidationCommittee`.

#### Verify the block proof
* Verify the blook proof based on the committee members.

#### Verify previous block hash pointer
* Calcualte the hash of the previously commited block header.

If All are valid return VALID_FOR_COMMIT else return NOT_VALID_FOR_COMMIT. (TODO timeout)

&nbsp;
## `ValidateResultsBlockConsensus` (method)
> Validates that a results block header can be committed assuming that the block body (content) is valid and its relevant hash values match the ones in the header and that the previous block was successfully committed.

#### Verify the block committee
* Calualte the previous block random seed (based on the results block proof)
* Get the committee members by calling `ConsensusBuilder.RequestValidationCommittee`.

#### Verify the block proof
* Verify the blook proof based on the committee members.

#### Verify previous block hash pointer
* Calcualte the hash of the previously commited block header. (TODO transactions hash pointer)

If All are valid return VALID_FOR_COMMIT else return NOT_VALID_FOR_COMMIT. (TBD timeout)
