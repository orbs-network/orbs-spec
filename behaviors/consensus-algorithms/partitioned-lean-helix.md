# Helix Consensus Algo
> PBFT based algorithm with rotating leader, randomly selected for each block height based on a verifiable random function.

## Archietcture - components nad inetrfaces

#### Consensu Algo
* Interfaces:
  * `GossipMessageReceived` - triggered by the gossip service upon message received with Conensus topic.
  * `RequestNewBlock() : block` - called by the One hieght PBFT, returns a block proposal.
  * `ValidateBlock(block) : valid` - called by the One hieght PBFT, valdiates a block proposal.
  * `CalcBlockHash (block) : block_hash` - called by the One hieght PBFT, calculates the hash on a block based on the hashing scheme.
  * `CommitBlock(block, pbft_proof, random_seed_database)` - called by the One hieght PBFT, generates a block proof and commits the block
  * `ValidateRandomSeedData(random_seed_data, sender)`
  * `AcknowledgeTransactionsBlockConsensus (block headers + proof)` - called by the Block storage s part of block sync.
  * `AcknowledgeResultsBlockConsensus ((block headers + proof)` - called by the Block storage s part of block sync.
  * `BroadcastMessage`
  * `UnicastMessage`

* Internal methods:
  * `OnNewConsensusRound` - internal, triggered upon block n-1 commit.

#### One Height PBFT
* Interfaces:
  * `PerformConsensus(block_height, node_list)` - perfroms a single block height consensus.
  * PBFT Message processing:
    * `OnPrePrepareReceived (PrePrepare)`
    * `OnPrepareReceived (Prepare)`
    * `OnCommitReceived (Commit)`
    * `OnViewChangeReceived (ViewChange)`
    * `OnNewViewReceived (NewView)`

* Internal methods:
  * `GetCurrentLeader (nodes_list, view)` - internal, calcualtes the leader for the view based on the ordered node list.
  * `OnPrepared`
  * `OnCommitReceived` 

  
#### PBFT Verifier
* `Interfaces: 
  * `CheckPBFTProof (pbft_proof, block_hash)` - validates a pbft proof for a specific block hash.

&nbsp;
#### Diagram <!-- TBD - Oded to add -->


## Design Notes
* All nodes particiapte in every committee, no `OnCommitBlockOutsideCommittee` flow.
* A random seed for a block / round is calcualted by a hash on the aggregated threshold signature of the previous block random seed.
  * The threshold signatrues are passed as part of the COMMIT messaage.
  * The threshold is set to 2f+1.
* Node sync is perfromed by the Block storage. The consensus algorithm does not store past block_height data.


## Configuration <!-- TBD - Oded to add -->

## Databases

#### Future Messages Cache
> Stores messages of future block_heights until block_height update.
* Accessed by (Block_height, View, Sender)
* Persistnet
* Stores messages with valid signature, i.e. an honest sender may only send one valid message type per {Block_height, View}
  * Should avoid storing duplciates which may be sent as part of an attack.

#### Previous block cache
* Stores required data from the previous block_height.


&nbsp;
## `OnNewConsensusRound` 

#### Get previous block data
* If the the previous block data is not available, get the previously committed block pair data by calling `BlockStorage.GetTransactionsBlockHeader` and `BlockStorage.GetResultsBlockHeader`.

#### Get Committee
* Calculate the `random_seed` for the upcoming block:
  * `random_seed` = SHA256(Previous block.random_seed_signature).
* Get an ordered list fo committee members for the block_height by calling `ConsensusContext.RequestValidationCommittee`.

#### Init a One Height PBFT
* Init One Height PBFT with the current block_height and the committee by calling `OneHeightPBFT.PerformConsensus`.
* Send all messages from the Future Message Cache with matching block height to the OneHeightPBFT
  * For each message in Future Message Cache with with matching block height
    * Determine the message type
    * Call the corresponding `OneHeightPBFT.On<XXX>`

&nbsp;
## `GossipMessageReceived`

#### Check message signature
* Discard if invalid signature

#### Check block height
* If message.block_height > current_block_height + configurable future_block_height_window - discard.
* If message.block_height > my_state.block_height - store in Future Messages Cache.
* If message.block_height < current_block_height - discard.

#### Demux message and pass to One Height PBFT
* Determine the message type
* Call the corresponding `OneHeightPBFT.On<XXX>`


&nbsp;
## `RequestNewBlock`
* Request new transactions block proposal (ordering phase) by calling `ConsensusContext.RequestNewTransactionsBlock`.
* Request new results block proposal (execution phase) by calling `ConsensusContext.RequestNewResultsBlock`.


&nbsp;
## `CalcBlockHash`
* Block_hash = SHA256({SHA256(TransactionsBlockHeader) XOR SHA256(ResultsBlockHeader)})


&nbsp;
## `ValidateBlock`

#### Check the hash pointers
* Check the TransactionBlockHeader.prev_block_hash_ptr matches the previous block
* Check the ResultsBlockHeader.prev_block_hash_ptr matches the previous block
* Check the ResultsBlockHeader.transactions_block_hash_ptr matches the TransactionBlockHeader

#### Check the previous block hash pointers
* Validate the transactions block (ordering phase) by calling `ConsensusContext.ValidateTransactionsBlock`.
* Validate the results block (execution phase) by calling `ConsensusContext.ValidateResultsBlock`.

&nbsp;
## `ValidateRandomSeedData`
* Get the the random seed signature share public key that corresponds to the sender.
* Check the signature on the random seed is valid.



&nbsp;
## `CommitBlock` 

#### Aggregate the threshold signatrue
* For each Sender in the random_seed_database, get the corresponding random seed signature share public key.
* Aggregate the signatures.

#### Generate LeanHelixBlockProof for the TransactionsBlock 
* Generate a LeanHelixBlockProof for the TransactionsBlock
  * pbft_proof = pbft_roof
  * block_hash_mask = SHA256(ResultsBlockHeader)
  * random_seed_signature = aggregated threshold signatrue

#### Generate LeanHelixBlockProof for the TransactionsBlock 
* Generate a LeanHelixBlockProof for the TransactionsBlock
  * pbft_proof = pbft_roof
  * block_hash_mask = SHA256(TransactionsBlock)
  * random_seed_signature = aggregated threshold signatrue

* Append the corresponding LeanHelixBlockProof to the TransactionsBlock and ResultsBlockHeader.
* Commit the BlockPair by calling `BlockStorage.CommitBlock`.

#### Triger the next block height
* Cache the required fields from the block headers for the next block_height.
* Increment current_block_height.
* Initiate the next block height round by calling `OneHeightPBFT.PerformConsensus`.


&nbsp;
## `AcknowledgeTransactionsBlockConsensus` (method)

> Validates the consensus for an untrusted block header. Ignores whether the block body (content) is valid and its relevant hash values match the ones in the header. Called internally and by block storage during block sync.

#### Check previous block pointer
* Verify the previous block hash pointer matches the hash of the previous block (given).

#### Get the block committee
* Calculate the random seed from the previous block (given).
* Get the desired committee size from [configuration](../config/services.md).
* Get a sorted list of committee members for the block by calling `ConsensusContext.RequestOrderingCommittee`.
* Note: If the consensus algorithm relies on a single committee for both, call `ConsensusContext.RequestValidationCommittee` only based on the Results block random seed.

#### Verify the block proof
* Verify the block proof based on the committee members.
* If all valid, update the consensus algorithm about the block commit (with block height and consensus dependent data).

&nbsp;
## `AcknowledgeResultsBlockConsensus` (method)

> Validates the consensus for an untrusted block header. Ignores whether the block body (content) is valid and its relevant hash values match the ones in the header. Called internally and by block storage during block sync.

#### Check previous block pointer
* Verify the previous block hash pointer matches the hash of the previous block (given).

#### Get the block committee
* Calculate the random seed from the previous block (given).
* Get the desired committee size from [configuration](../config/services.md).
* Get a sorted list of committee members for the block by calling `ConsensusContext.RequestValidationCommittee`.

#### Verify the block proof
* Verify the block proof based on the committee members.
* If all valid, update the consensus algorithm about the block commit (with block height and consensus dependent data).












&nbsp;
## `AcknowledgeTransactionsBlockConsensus` and `AcknowledgeResultsBlockConsensus`
<!-- TODO Consider to unify to a single function to prevent races -->
> See consensus-algo.md, upon valid block

#### Check Block_height
* Ignore if recevied block_height <= my_state.block_height

#### Triger the next block height round
* Update my_state.Block_height = recevied block_height.
* Cache the required fields from the block headers for the next round.
* Clear all messages with block_height <= my_state.block_height from the log.
* Initiate the next block height round by triggering `OnNewConsensusRound`.

&nbsp;
## `GetCurrentLeader(ordered_committee, message.view)`
> Returns the leader for the view
* Return ordered_committee[View MOD ordered_committee size]







## Configuration <!-- TODO --->
> Held by each node consensus algorithm, read from configuration file upon init
* Committee_size
* f_byzantine - max number byzantine nodes (default - 2/3 Committee_size + 1)
* Cryptographic keys

  * Last block headers pair cache


&nbsp;
## `OnInit`
> Perfromed upon the service init
* Read persistent data:
  * my_state.Block_height
  * my_state.View
  * Candidate_block
  * Candidate_block_hash
  * Prepared_proof