# Lean Helix Spec
> PBFT absed algorithm with rotating leader randomly selected for each term based on a verifiable random function.

## Assumptions
* PRE_PREPARE and COMMIT messages are broadcasted all nodes.
* All nodes particiapte in every committee, no `OnCommitBlockOutsideCommittee` flow.
* A random seed for a block / round is calcualted by a hash on the aggregated threshold signature of the previous block random seed.
  * The threshold signatrues are passed as part of the commit messaage.
  * The threshold is set to 2f+1.
* The consensus algo only keeps PBFT logs of the present {block_height, view}, a sync of the blockchain history is perfromed by node sync. 
* View Change of prepared node incldus the candidate block
  * May add a request / response emssage as optimization
* New View includes all the view change proofs
  * May add an optimization to avoid n^2 signatures in new view
* New View includes the block => PrePrepare is only received in View 0.
* COMMIT messages are accepted even if not in Prepared state, Commit_locally requires to be in Prepared state.
* COMMIT of earlier views are not be accepted. - TBD accepted only of committed database valid
* The threshold signature is included in the COMMIT message
* Log COMMIT messages of any view, participate in the highest valid view.

## Messages
* PRE_PREPARE - sent by the leader, incldues a block proposal, view, term, sig{message type, hash, view, term}
* PREPARE - sent by the validators, incldues the block hash, view, term, sig{message type, hash, view, term}
* COMMIT - sent by all nodes, incldues the block hash, view, term, sig{message type, hash, view, term}, sig_share{message type, hash, view, term}
* VIEW_CHANGE - sent upon timeout to the next elader candidate
* NEW_VIEW - sent by the newly elected leader (upon view change)

## Configuration
> Held by each node consensus algorithm, read from configuration file upon init
* Committee_size
* f_byzantine - max number byzantine nodes (default - 2/3 Committee_size + 1)
* Cryptographic keys

## Databases

#### Received Messages Cache
> Stores messages of future block_height pending stateful processing until block_height update. Used to reduces the chance for costly syncs.
* Accessed by (Block_height, View, Sender)
* No need to be persistent
* Stores messages with valid signature, i.e. an honest sender may only send one valid message type per {Block_height, View}
  * Should avoid storing duplciates which may be sent as part of an attack.

#### Messages Log
> Stores the current term messages that were accepted.
* Accessed by (View, Sender, Message_type)
* Persistent

#### State
> Stores the current term state variables, persistent
* Persistent

* State variables:
  * Block_height (Term)
  * Last block headers pair cache
  * Block_height 
  * View
  * Candidate_block
  * Candidate_block_hash
  * Prepared_proof 
    * Block_Height
    * View
    * Block_hash
    * PP_proof
    * Prepare_proofs
  * Prepared_block

&nbsp;
## `OnInit`
TODO


&nbsp;
## `OnNewConsensusRound` 
* Get the previously committed block pair data by calling `BlockStorage.GetTransactionsBlockHeader` and `BlockStorage.GetResultsBlockHeader`.
  * It is recommanded to cache the previously committed block data in order to prevent fetching the data.
  * If fails or times out, skip this round.
* Calculate the `random_seed` for the upcoming block:
  * Aggregate the threshold signatrues of the previous results block
  * `random_seed` = SHA256(aggregated signatrue).
* Get a sorted list of committee members for the upcoming results block (execution validation phase) by calling `ConsensusBuilder.RequestValidationCommittee`.
* Determine the current leader using `CurrentLeader(ordered_committee, view)`.
* Update the my_state.block_height to the next block_height.

#### `Init my_state for the next round`
* my_state.Block_height = my_state.Block_height + 1
* my_state.view = 0
* my_state.Prepared_proof = {}
* my_state.Candidate_block = {}

#### `OnNewConsensusRound - Leader Only`
* Request new transactions block proposal (ordering phase) by calling `ConsensusContext.RequestNewTransactionsBlock`.
* Request new results block proposal (execution phase) by calling `ConsensusContext.RequestNewResultsBlock`.
* Generate PRE_PREPARE message and broadcast to all committee nodes
  * Type = PRE_PREPARE
  * Sender = Node public key
  * View = my_state.view
  * Block_height = my_state.block_height
  * BlockPair = constructed block (without proofs)
  * Block_hash = SHA256(TransactionBlockHeader) XOR SHA256(ResultsBlockHeader).
  * Signature {Type, Block_height, View, Hash(Block pair)}
* Update the state and log
  * Log the PRE_PREPARE message
  * Candidate_block_hash = PRE_PREPARE message.Block_hash
  * Candidate_block = PRE_PREPARE message.BlockPair

#### `OnNewConsensusRound - All nodes`
* Set timer with timeout = configurable base_round_timeout x 2^(my_state.view).


&nbsp;
## `OnPrePrepareReceived` 

#### Check message signature and sender
* Discard if signature mismatch
* Discard if the sender is not a valid participant in the block's committee.

#### Check {block_height, view}
* Discard if message.block_height > my_state.block_height + configurable future_block_height_window.
* If message.block_height > my_state.block_height, store in Received Messages Cache.
* Discard if message.block_height < my_state.block_height.
* Discard if message.view != 0.
  * Note: NEW_VIEW messages incldue the NV_PRE_PREPARE, no PRE_PREPARE messages should be received not in view 0.

#### Check that the sender is the leader for the view
* Discard if the sender isn't the leader for the view based on `GetCurrentLeader(ordered_committee, message.view)`.

#### Check no duplicate PrePrepare message was logged
* Discard if a PRE_PREPARE message was already logged for the same view and sender.

#### Check message content
* Check PRE_PREPARE message.Block_hash matches the block pair headers hash. (Base on the hash scheme)
* Check the TransactionBlockHeader.prev_block_hash_ptr and the ResultsBlockHeader.prev_block_hash_ptr.
* Validate the transactions block (ordering phase) by calling `ConsensusContext.ValidateTransactionsBlock`.
* Validate the results block (execution phase) by calling `ConsensusContext.ValidateResultsBlock`.
* Check the ResultsBlockHeader.Metadata.RandomSeed <!--  TODO, place metadata in both blocks ?-->
* If one of the checks fails, discard message.

#### Generate PREPARE message
> Performed if all the message checks and content checks have passed
* Generate PREPARE message and broadcast to all nodes.
  * Type = PREPARE
  * Sender = Node public key.
  * Block_height = my_state.block_height
  * View = my_state.view
  * Block_hash = PRE_PREPARE message.Block_hash
  * Signature {Type, Block_height, View, Block_hash}

#### Log the massages and update the state
> Performed if all the message checks and content checks have passed
* Log the received PRE_PREPARE message
* Log the sent PREPARE message
* Update the state
  * Candidate_block_hash = PRE_PREPARE message.Block_hash
  * Candidate_block = PRE_PREPARE message.BlockPair

#### Check if Prepared
* Check the log(View = message.view)
  * If a PRE_PREPARE message AND 2xf PRPARE messages are logged with Block_hash equal to the PRE_PREPARE message.Block_hash. 
    * Set my_state.Prepared
    * Call `OnPrepared`

&nbsp;
## `OnPrepareReceived` 

#### Check message signature and sender
* Discard if signature mismatch
* Discard if the sender is not a valid participant in the block's committee.

#### Check Block height and view
* Discard if message.block_height > my_state.block_height + configurable future_block_height_window.
* If message.block_height > my_state.block_height, store in Received Messages Cache.
* Discard if message.block_height < my_state.block_height.
* Discard if message.view < my_state.view

#### Check that the sender isn't the leader for the view
* Discard if the sender is the leader for the view based on `GetCurrentLeader(ordered_committee, message.view)`.

#### Check no duplicate Prepare message was logged
* Discard if a PREPARE message was already logged for the same view and sender.

#### Log message 
* Log the PREPARE message in message log.

#### Check if Prepared
* Check the log(View = message.view)
  * If a PRE_PREPARE message AND 2xf PRPARE messages are logged with Block_hash equal to the PRE_PREPARE message.Block_hash. 
    * Set my_state.Prepared
    * Call `OnPrepared`


&nbsp;
## `OnPrepared` 

#### Generate COMMIT message
* Generate COMMIT message and broadcast to all nodes:
  * Type = COMMIT 
  * Sender = Node public key.
  * Block_height = my_state.block_height
  * View = my_state.view
  * Block_hash = logged PRE_PREPARE (View = my_state.view).Block_hash
  * Signature {Type, Block_height, View, Block_hash}
* Log the COMMIT message in message log. 

#### Generate Prepared Proof
* Generate Prepared_proof based on Log(View = message.view):
  * Block_Height = my_state.Block_height
  * View = my_state.View
  * Block_hash = Candidate_block_hash
  * PP_proof = {PRE_PREPARE message.Sender, PRE_PREPARE message.Signature}
  * For each PREPARE
    * Prepare_proofs.add({PREPARE message.Sender, PREPARE message.Signature})

#### Check if Commited_localy
* Check the log(View = message.view)
  * If a PRE_PREPARE message AND 2xf+1 COMMIT messages are logged with Block_hash equal to the PRE_PREPARE message.Block_hash. 
    * Set my_state.Committed_locally
    * Call `OnCommittedLocally`


&nbsp;
## `OnCommitReceived` 

#### Check message signature and sender
* Discard if signature mismatch
* Discard if the sender is not a valid participant in the block's committee.

#### Check Block height and view
* Discard if message.block_height > my_state.block_height + configurable future_block_height_window.
* If message.block_height > my_state.block_height, store in Received Messages Cache.
* Discard if message.block_height < my_state.block_height.
Note: a node may receive COMMIT messages of earlier views.

#### Check no duplicate COMMIT message was logged
* Discard if a COMMIT message was already logged for the same view and sender.

#### Check threshold signature
* Discard if the threshold signatrue of the sender on previous block random seed is invalid.

#### Log message
* Log the COMMIT message in message log.

#### Check if Commited_localy
* Check the log(View = message.view)
  * If a PRE_PREPARE message AND 2xf+1 COMMIT messages are logged with Block_hash equal to the PRE_PREPARE message.Block_hash. 
    * Set my_state.Committed_locally
    * Call `OnCommittedLocally`


&nbsp;
## `OnCommittedLocally` 

#### Generate block proof and commit block
* Aggregate the threshold signatrues of the logged COMMIT messages in (View = message.view) to generate an aggregated threshold signatrue.
* Generate a LeanHelixBlockProof for the TransactionsBlock based on Log(View = message.view):
  * opaque_message_type = COMMIT
  * block_height = my_state.Block_height
  * View = my_state.View
  * block_hash_mask = SHA256(ResultsBlockHeader)
  * block_hash = SHA256(TransactionBlockHeader)
  * For each COMMIT in (View = message.View)
    * block_signatures.add({COMMIT message.Sender, COMMIT message.Signature})
  * random_seed_signature = aggregated threshold signatrue

* Generate a LeanHelixBlockProof for the ResultsBlockHeader based on Log(View = message.view):
  * opaque_message_type = COMMIT
  * block_height = my_state.Block_height
  * View = my_state.View
  * block_hash_mask = SHA256(TransactionBlockHeader)
  * block_hash = SHA256(ResultsBlockHeader)
  * For each COMMIT in (View = message.View)
    * block_signatures.add({COMMIT message.Sender, COMMIT message.Signature})
  * random_seed_signature = aggregated threshold signatrue

* Append the corresponding LeanHelixBlockProof to the TransactionsBlock and ResultsBlockHeader.
* Commit the BlockPair by calling `BlockStorage.CommitBlock`.

#### Triger the next block height round
* Cache the required fields from the block headers for the next round.
* Clear all messages with the block_height from the log.
* Initiate the next block height round by triggering `OnNewConsensusRound`.


&nbsp;
## `OnTimeOut`
> Timeout of the PBFT timer.
> Reset conditions: on new consensus round, on timeout.

#### Init State for a View Change
* my_state.view = my_state.view + 1.
* Reset the timer to configurable base_round_timeout x 2^(my_state.view).
* Determine the current leader using `GetCurrentLeader(ordered_committee, message.view)`.

#### Generate a VIEW_CHANGE message
* Generate VIEW_CHANGE message
  * Type = VIEW_CHANGE
  * Sender = Node public key.
  * Block_height = my_state.block_height
  * View = my_state.view
  * If my_state.Prepared then 
    * Prepared_proof = my_state.Prepared_proof <!-- TODO CHANGE METHODOLOGY>
    * Prepared_block = my_state.Candidate_block
  * Else 
    * Prepared_proof = {}
    * Prepared_block = {}
  * Signature {Type, Block_height, View, Prepared_proof}
* Log the VIEW_CHANGE message in message log. 
 
#### OnTimeOut - Validator Only
* Send the VIEW_CHANGE as unicast to the new view's leader


&nbsp;
## ValidatePreparedProof(View_change_view)
> Used by `OnViewChangeReceived` and `OnNewViewReceived` 

* Check Prepared_proof.Block_height = my_state.Block_height
* Check Prepared_proof.View < View_change_view
* Check Prepared_proof.Block_hash matches the block
* Verify a PP_proof and 2f Prepare_proofs, from different senders.
* For each proof in (PP_proof, Prepare_proofs):
  * Check signature
  * Check that the sender is a valid participant in the round's committee.
* If all pass return valid., else return invalid.


&nbsp;
## `OnViewChangeReceived` 

#### Check message signature and sender
* Discard if signature mismatch
* Discard if the sender is not a valid participant in the block's committee.

#### Check Block height and view
* Discard if message.block_height > my_state.block_height + configurable future_block_height_window.
* If message.block_height > my_state.block_height, store in Received Messages Cache.
* Discard if message.block_height < my_state.block_height.
* Discard if message.view < my_state.view
* Note: VIEW_CHANGE of View + 1 is common if the node timeout is behind.

#### Check no duplicate VIEW_CHANGE message was logged
* Discard if a VIEW_CHANGE message was already logged for the same view and sender.

#### Check that node is the leader for the view.
* Discard if the node isn't the leader for the view based on `GetCurrentLeader(ordered_committee, message.view)`.

#### Check the VIEW_CHANGE Prepared_proof
* Check the received Prepared_proof is valid using `ValidatePreparedProof(View_change_view)`
* Check that the received Prepared_proof.Block_hash matches the received Prepared_proof.Prepared_block.
* If one of the checks fails, discard message.

#### Log message
* Log the VIEW_CHANGE message in message log.

#### Check if new view
* Check the log(View = message.view)
  * If 2xf+1 VIEW_CHANGE messages are logged:
    * Call `LocalNewView`


&nbsp;
## `LocalNewView` 

#### `Init State for a New View`
* my_state.view = message.View.
* Reset the timer to configurable base_round_timeout x 2^(my_state.view).

#### Determine the next candidate block
* From all VIEW_CHANGE messages in (View = my_state.view) with Prepared_proof != {}, find the one with the highest Prepared_proof.View.
* If a VIEW_CHANGE message with Prepared_proof != {} was found: 
  * Candidate_block = highest prepraed view VIEW_CHANGE.BlockPair
  * Candidate_block_hash = highest prepraed view VIEW_CHANGE.Block_hash
* Else (no VIEW_CHANGE was Prepared)
  * Construct a new Candidate_block
    * Request new transactions block proposal (ordering phase) by calling `ConsensusBuilder.RequestNewTransactionsBlock`.
    * Request new results block proposal (execution phase) by calling `ConsensusBuilder.RequestNewResultsBlock`.
  * Candidate_block_hash = SHA256(TransactionBlockHeader) XOR SHA256(ResultsBlockHeader).

#### Generate New_view_proof
* New_view_proof = All logged VIEW_CHANGE messages in (View = my_state.view) without the Prepared_block

#### Generate New View PRE_PREPARE
> The New View PRE_PREPARE message is a regular PRE_PREPARE message encapsulated in NEW_VIEW. Using the same message format enables a single PREPARE flow.
* Generate PRE_PREPARE
  * Type = PRE_PREPARE
  * Sender = Node public key.
  * View = my_state.view
  * Block_height = my_state.block_height
  * BlockPair = Candidate_block
  * Block_hash = Candidate_block_hash
  * Signature {Type, Block_height, View, Hash(Block pair)}

#### Generate NEW_VIEW message
* Generate NEW_VIEW message and broadcast to all nodes:
  * Type = NEW_VIEW 
  * Sender = Node public key.
  * Block_height = my_state.block_height
  * View = my_state.view
  * New_view_proof = New_view_proof
  * NVPP = New View PRE_PREPARE message
  * Signature {Type, Block_height, View, New_view_proof, NVPP}

#### Log NVPP Message and update state
* Log the NV_PRE_PREPARE message
  * Note: there's no need to log the NEW_VIEW message.
  

&nbsp;
## `OnNewViewReceived` 

#### Check message signature and sender
* Discard if signature mismatch
* Discard if the sender is not a valid participant in the block's committee.

#### Check {block_height, view}
* Discard if message.block_height > my_state.block_height + configurable future_block_height_window.
* If message.block_height > my_state.block_height, store in Received Messages Cache.
* Discard if message.block_height < my_state.block_height.
* Discard if message.view < my_state.view.

#### Check that the sender is the leader for the view
* Discard if the sender isn't the leader for the view based on `GetCurrentLeader(ordered_committee, message.view)`.

#### Check no duplicate PrePrepare message was logged
* Discard if a PRE_PREPARE message was already logged for the same view and sender.

#### Check New_view_proof
* Verify 2f+1 VIEW_CHANGE messages, from different senders.
* For each VIEW_CHANGE message verify:
  * Type = VIEW_CHANGE
  * Sender is a valid participant in the block's committee.
  * Block_height = NEW_VIEW message.Block_height
  * View = NEW_VIEW message.View
  * Prepared_proof is valid using `ValidatePreparedProof(View_change_view)`
  * Valid signature
* Discard if one of the checks fails.

<!-- TODO consider unify with OnPrePrepareRecevied -->
#### Check encapsulated New View PRE_PREPARE (NVPP) message 
* Check the New View PRE_PREPARE message fields
  * Check Type = PRE_PREPARE
  * Check Sender = NEW_VIEW.Sender
  * Check View = NEW_VIEW.View
  * Check Block_height = NEW_VIEW.Block_height
  * Check Block_hash matches the NVPP block pair headers hash.
  * Check signature

#### Check encapsulated New View PRE_PREPARE message block
* From all VIEW_CHANGE messages in New_view_proof with Prepared_proof != {} in , find the one with the highest Prepared_proof.View.
* If a VIEW_CHANGE message with Prepared_proof != {} was found: 
  * Check NVPP.Block_hash = highest prepraed view VIEW_CHANGE.Block_hash.
* Else (no VIEW_CHANGE was Prepared)
  * Check the TransactionBlockHeader.prev_block_hash_ptr and the ResultsBlockHeader.prev_block_hash_ptr.
  * Validate the transactions block (ordering phase) by calling `ConsensusBuilder.ValidateTransactionsBlock`.
  * Validate the results block (execution phase) by calling `ConsensusBuilder.ValidateResultsBlock`.
  * Check the ResultsBlockHeader.Metadata.RandomSeed <!-- Oded TODO, place metadata in both blocks ?-->
* Discard the NEW_VIEW message if one of the checks fails.

#### Init State for a New View
* my_state.view = message.View.
* Reset the timer to configurable base_round_timeout x 2^(my_state.view).

#### Generate PREPARE message
> Performed if all the message checks and content checks have passed
* Generate PREPARE message and broadcast to all nodes.
  * Type = PREPARE
  * Sender = Node public key.
  * Block_height = my_state.block_height
  * View = my_state.view
  * Block_hash = NV_PRE_PREPARE message.Block_hash
  * Signature {Type, Block_height, View, Block_hash}

#### Log the massages and update the state
> Performed if all the message checks and content checks have passed
* Log the received NV_PRE_PREPARE message
  * Note: there's no need to log the NEW_VIEW message.
* Log the sent PREPARE message
* Update the state
  * Candidate_block_hash = PRE_PREPARE message.Block_hash
  * Candidate_block = BlockPair

#### Check if Prepared
* Check the log(View = message.view)
  * If a PRE_PREPARE message AND 2xf PRPARE messages are logged with Block_hash equal to the PRE_PREPARE message.Block_hash. 
    * Set my_state.Prepared
    * Call `OnPrepared`


##


## Good Flow

#### View Change


## Oded Questions
1. What happens if my view < received view, should I drop or cache and process later ?
2. "rpc" `GetCurrentLeader(ordered_committee, message.view)`.
3. avoid sending blocks in VIEW_CHANGE, for example by V
4. TODO - Remove pointers validation from Consensus Context