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
* Calculate the shared `random_seed` for the upcoming block.
* Get a sorted list of committee members for the upcoming block by calling `ConsensusCore.RequestCommittee`. 
  * It is recommanded to cache the committee members based up to the valid block as indicated in the response.


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

* Depends on consensus algorithm.

&nbsp;
## `ValidateTransactionsBlockCommit` (method)
> Validates that a tarnsactions block header can be committed assuming that the block body (content) is valid and its relevant hash values match the ones in the header and that the previous block was successfully committed.

#### Get the previously commited block data
* Fetch the previous committed tarnsactions block data by calling `BlockStorage.GetTransactionsHeaderAndProof`.
  * It is recommanded to cache the previous block data in order to prevent fetching teh data when possible.

#### Verify the block committee
* Calualte the previous block random seed.
* Fetch the committee members by calling `ConsensusCore.RequestOrderingCommittee`.
  * It is recommanded to cache the committee members based up to the valid block as indicated in the response.

#### Verify the block proof
* Verify the blook proof based on the committee members.

#### Verify previous block hash pointer
* Calcualte the hash of the previously commited block header.

If All are valid return VALID_FOR_COMMIT else return NOT_VALID_FOR_COMMIT. (TODO timeout)

&nbsp;
## `ValidateResultsBlockCommit` (method)
> Validates that a results block header can be committed assuming that the block body (content) is valid and its relevant hash values match the ones in the header and that the previous block was successfully committed.

#### Get the previously commited block data
* Fetch the previous committed tarnsactions block data by calling `BlockStorage.GetResultsHeaderAndProof`.
  * It is recommanded to cache the previous block data in order to prevent fetching teh data when possible.

#### Verify the block committee
* Calualte the previous block random seed.
* Fetch the committee members by calling `ConsensusCore.RequestValidationCommittee`.
  * It is recommanded to cache the committee members based up to the valid block as indicated in the response.

#### Verify the block proof
* Verify the blook proof based on the committee members.

#### Verify previous block hash pointer
* Calcualte the hash of the previously commited block header. (TODO transactions hash pointer)

If All are valid return VALID_FOR_COMMIT else return NOT_VALID_FOR_COMMIT. (TBD timeout)
