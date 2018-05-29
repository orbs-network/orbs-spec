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
* Get the previously committed block pair data by calling `BlockStorage.GetTransactionsBlockHeader` and `BlockStorage.GetResultsBlockHeader`.
  * It is recommanded to cache the previously committed block data in order to prevent fetching the data.
  * If fails or times out, skip this round.
* Calculate the `random_seed` for the upcoming block.
* Get a sorted list of committee members for the upcoming transaction block (ordering phase) by calling `ConsensusBuilder.RequestOrderingCommittee`.
* Get a sorted list of committee members for the upcoming results block (execution validation phase) by calling `ConsensusBuilder.RequestValidationCommittee`.
* Note: If the consensus algorithm relies on a single committee for both, call `ConsensusBuilder.RequestValidationCommittee` only based on the results block random seed.

#### If Leader
* Request new transactions block proposal (ordering phase) by calling `ConsensusBuilder.RequestNewTransactionsBlock`.
* Request new results block proposal (execution phase) by calling `ConsensusBuilder.RequestNewResultsBlock`.
* Sign both proposals (according to the algo spec) - on hash of the header.
* Send signed proposals with gossip to the algo of all committee nodes (according to the algo spec).

&nbsp;
## `OnBlockProposalReceived` (method) <!-- tal can finish -->

> Through the consensus algorithm, a non-leader node in the committee receives a proposal for a new block by the leader.

* Wait until deceided if participating in this round or not (`OnNewConsensusRound`).
* Perform algorithm related checks on the proposal
  * Signature, block height, previous block pair hash pointers.
* Validate the transactions block (ordering phase) by calling `ConsensusBuilder.ValidateTransactionsBlock`.
* Validate the results block (execution phase) by calling `ConsensusBuilder.ValidateResultsBlock`.
* Sign both approvals (according to the algo spec) - on hash of the header.
* Send signed approvals with gossip to the algo of all committee nodes (according to the algo spec).

&nbsp;
## `OnCommitBlockInsideCommittee` (method) <!-- tal can finish -->

> The consensus algorithm has decided a block proposal can be committed (performed by committee members)

* Pass the committed block (including the block proofs) to the block storage by calling `BlockStorage.CommitBlock`.


## `OnCommitBlockOutsideCommittee` (method) <!-- tal can finish -->
> Performed by non-committee members that received a block in order to validate that it was committed under consensus.

#### Check the Transactions Block Header (stateless)
* Check the block protocol version.
* Check the virtual chain.
* Check block height
  * If the block already exist (block height != last_commited_block block + 1) discard.
* Check transactions_root_hash 
  * Calculate the merkle root hash of the block's transactions verify the hash in the header.
* Check metadata hash
  * Calculate the hash of the block's metadata and verify the hash in the header.

#### Check the Results Block Header (stateless)
* Check the block protocol version.
* Check the virtual chain.
* Check block height
  * If the block already exist (block height != last_commited_block block + 1) discard.
* Check receipts_root_hash
  * Calculate the merkle root hash of the block's transactions verify the hash in the header.
* Check state_diff_hash
  * Calculate the hash of the block's metadata and verify the hash in the header.

#### Get previously committed block data
* Get the previously committed block pair data by calling `BlockStorage.GetTransactionsBlockHeader` and `BlockStorage.GetResultsBlockHeader`.
  * It is recommanded to cache the previously committed block data in order to prevent fetching the data.
  * If fails or times out, skip this round. - TODOD

#### Validate that the block is under consensus and can be committed.
* Validate that the block can be commited under consensus by calling: `AcknowledgeTransactionsBlockConsensus` and `AcknowledgeResultsBlockConsensus`.

* Pass the committed block (including the block proofs) to the block storage by calling `BlockStorage.CommitBlock`.


&nbsp;
## `GossipMessageReceived` (method)

> Handles messages received from another node, expect to receive only Consensus messages that cosnesus-algo have subscribed to.

* Depends on consensus algorithm.

&nbsp;
## `AcknowledgeTransactionsBlockConsensus` (method)
> Validates that a tarnsactions block header can be committed assuming that the block body (content) is valid and its relevant hash values match the ones in the header and that the previous block was successfully committed.

#### Verify the block committee
* Calualte the previous block random seed.
* Get the committee members by calling `ConsensusBuilder.RequestTransactionCommittee`.
* Note: If the consensus algorithm relies on a single committee for both, call `ConsensusBuilder.RequestValidationCommittee` only based on the results block random seed.

#### Verify the block proof
* Verify the blook proof based on the committee members.

#### Verify previous block hash pointer
* Calcualte the hash of the previously commited block header.

* If All are valid:
  * Update the consensus algorithm of the block commit (Committed block height, consensus dependent data).
  * return VALID_FOR_COMMIT 
* else return NOT_VALID_FOR_COMMIT.

&nbsp;
## `AcknowledgeResultsBlockConsensus` (method)
> Validates that a results block header can be committed assuming that the block body (content) is valid and its relevant hash values match the ones in the header and that the previous block was successfully committed.

#### Verify the block committee
* Calualte the previous block random seed.
* Get the committee members by calling `ConsensusBuilder.RequestValidationCommittee`.

#### Verify the block proof
* Verify the blook proof based on the committee members.

#### Verify previous block hash pointer
* Calcualte the hash of the previously commited block header.

* If All are valid:
  * Update the consensus algorithm of the block commit (Committed block height, consensus dependent data).
  * return VALID_FOR_COMMIT 
* else return NOT_VALID_FOR_COMMIT.