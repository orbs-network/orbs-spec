# OneHeight PBFT
> PBFT based algorithm, perfroms a single block_height consensus. 

## Design notes
* PREPARE and COMMIT messages are broadcasted to all nodes.
* The sync flow is performed outside the scope of the OneHeight PBFT
* View Change message sent by a prepared node includes the candidate block
  * May add a request / response message as optimization
* New View includes all the view change proofs and a signed NV_PRE_PREPARE
  * May add an optimization to avoid n^2 signatures in new view
* A block can be committed (Commit_locally) even if not in Prepared state. (The block was received in PRE_PREPARE or NV_PRE_PREPARE).
* COMMIT messages of earlier views are accepted.

## Messages
* PRE_PREPARE - sent by the leader
* PREPARE - sent by the validators
* COMMIT - sent by all nodes, incldues also the random seed signature.
* VIEW_CHANGE - sent upon timeout to the next leader candidate
* NEW_VIEW - sent by the newly elected leader (upon view change)


## Databases

#### Messages Log
> Stores the current block_height messages that were accepted.
* Accessed by (View, Sender, Message_type)
* Persistent

#### State
> Stores the current block_height state variables, persistent
* Persistent

* State variables:
  * Block_height (Term)
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
## `PerformConsensus(block_height, ordered_committee)`
* Reset the state:
  * Clear Messages Log
  * my_state.View = 0
  * my_state.Candidate_block = {}
  * my_state.Prepared_proof = {}
  * my_state.Prepared_block = {}

#### Determine if leader
* Determine the current leader using `CurrentLeader`.

#### `PerformConsensus - Leader Only`
* Request new candidate block by calling `ConsensusAlgo.RequestNewBlock`
* Calculate the candidate_block_hash, by calling `ConsensusAlgo.CalcBlockHash`
* Generate PRE_PREPARE message and broadcast to all committee nodes
  * Type = PRE_PREPARE
  * Sender = Node public key
  * View = my_state.View
  * Block_height = my_state.Block_height
  * Block = candidate_block
  * Block_hash = candidate_block_hash
  * Signature {Type, Block_height, View, Block_hash}
* Update the state and log
  * Log the PRE_PREPARE message
  * Candidate_block_hash = PRE_PREPARE message.Block_hash
  * Candidate_block = PRE_PREPARE message.Block

#### `PerformConsensus - All nodes`
* Set timer with timeout = configurable base_round_timeout.


&nbsp;
## `OnPrePrepareReceived` 

#### Check message sender
* Discard if the sender is not a valid participant node.
* Discard if the sender isn't the leader for the view based on `GetCurrentLeader`.

#### Check the message view
* Discard if message.view != 0
  * Note: NEW_VIEW messages incldue the NV_PRE_PREPARE, no PRE_PREPARE messages should be received not in view 0.

#### Check no duplicate PrePrepare message was logged
* Discard if a PRE_PREPARE message was already logged for the same view.

#### Check message content
* Get the hash of the PRE_PREPARE message.block by calling `ConsensusAlgo.CalcBlockHash` and compare it to the message.block_hash.
* Validate the block content by calling `ConsensusAlgo.ValidateBlock`.
* If one of the checks fails, discard message.

#### Generate PREPARE message
> Performed if all the message checks and content checks have passed
* Generate PREPARE message and broadcast to all nodes.
  * Type = PREPARE
  * Sender = Node public key.
  * Block_height = my_state.Block_height
  * View = my_state.View
  * Block_hash = PRE_PREPARE message.Block_hash
  * Signature {Type, Block_height, View, Block_hash}

#### Log the massages and update the state
> Performed if all the message checks and content checks have passed
* Log the received PRE_PREPARE message
* Log the sent PREPARE message
* Update the state
  * Candidate_block_hash = PRE_PREPARE message.Block_hash
  * Candidate_block = PRE_PREPARE message.Block

#### Check if Prepared
* Check the log(View = message.View)
  * If a PRE_PREPARE message AND exactly (trigger once) 2xf PRPARE messages are logged with Block_hash equal to the PRE_PREPARE message.Block_hash. 
    * Set my_state.Prepared
    * Call `OnPrepared`

&nbsp;
## `OnPrepareReceived` 

#### Check message sender
* Discard if the sender is not a valid participant node.
* Discard if the sender is the leader for the view based on `GetCurrentLeader`.

#### Check the message view
* Discard if message.view < my_state.view

#### Check no duplicate Prepare message was logged
* Discard if a PREPARE message was already logged for the same view and sender.

#### Log message 
* Log the PREPARE message in message log.

#### Check if Prepared
* Check the log(View = message.view)
  * If a PRE_PREPARE message AND exactly (trigger once) 2xf PRPARE messages are logged with Block_hash equal to the PRE_PREPARE message.Block_hash. 
    * Set my_state.Prepared
    * Call `OnPrepared`


&nbsp;
## `OnPrepared` 

#### Generate COMMIT message
* Generate COMMIT message and broadcast to all nodes:
  * Type = COMMIT 
  * Sender = Node public key.
  * Block_height = my_state.Block_height
  * View = my_state.View
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
* Prepared_block = PRE_PREPARE message.Block

#### Check if Commited_localy
* Check the log(View = message.view)
  * If a PRE_PREPARE message AND exactly (trigger once) 2xf+1 COMMIT messages are logged with Block_hash equal to the PRE_PREPARE message.Block_hash. 
    * Call `OnCommittedLocally`


&nbsp;
## `OnCommitReceived` 

#### Check message sender
* Discard if the sender is not a valid participant node.

#### Check no duplicate COMMIT message was logged
* Discard if a COMMIT message was already logged for the same view and sender.

#### Check random seed data
* Validates the message random seed data by calling `ConsensusAlgo.ValifateRandomSeedData`, discard if invalid.

#### Log message
* Log the COMMIT message in message log.

#### Check if Commited_localy
* Check the log(View = message.View)
  * If a PRE_PREPARE message AND exactly (trigger once) 2xf+1 COMMIT messages are logged with Block_hash equal to the PRE_PREPARE message.Block_hash. 
    * Call `OnCommittedLocally`


&nbsp;
## `OnCommittedLocally` 
* Generate pbft_proof based on Log(View = message.view):
  * message_type = COMMIT
  * block_height = my_state.Block_height
  * View = my_state.View
  * block_hash = Log(View = message.view, PRE_PREPARE).block_hash.
  * For each COMMIT in (View = message.View)
    * block_signatures.add({COMMIT message.Sender, COMMIT message.Signature})
* Generate random_seed_database:
  * For each COMMIT in (View = message.View)
  * random_seed_data.add({COMMIT message.Sender, COMMIT message.random_seed_share})
* Call `ConsensusAlgo.CommitBlock`. 


&nbsp;
## `OnTimeOut`
> Timeout of the PBFT timer.
> Reset conditions: on `PerformConsensus`, on timeout, on `LocalNewView`.

#### Init State for next view
* my_state.view = my_state.view + 1.
* Reset the timer to configurable base_round_timeout x 2^(my_state.view).
* Determine the current leader using `GetCurrentLeader`.

#### Generate a VIEW_CHANGE message
* Generate VIEW_CHANGE message
  * Type = VIEW_CHANGE
  * Sender = Node public key.
  * Block_height = my_state.block_height
  * View = my_state.view
  * If my_state.Prepared then 
    * Prepared_proof = my_state.Prepared_proof
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
* Verify a PP_proof and 2f Prepare_proofs, from different senders.
* For each proof in (PP_proof, Prepare_proofs):
  * Check signature
  * Check that the sender is a valid participant node.
* If all pass return valid., else return invalid.


&nbsp;
## `OnViewChangeReceived` 

#### Check message sender
* Discard if the sender is not a valid participant node.

#### Check the message view
* Discard if message.view < my_state.view 
  * Note: VIEW_CHANGE of View + 1 is common if the node timeout is behind.

#### Check no duplicate VIEW_CHANGE message was logged
* Discard if a VIEW_CHANGE message was already logged for the same view and sender.

#### Check that node is the leader for the view.
* Discard if the node isn't the leader for the view based on `GetCurrentLeader`.

#### Check the VIEW_CHANGE Prepared_proof
* Check the received Prepared_proof is valid using `ValidatePreparedProof(View_change_view)`
* Get the hash of the block by calling `ConsensusAlgo.CalcBlockHash` and compare it to the Prepared_proof.Block_hash.
* If one of the checks fails, discard message.

#### Log message
* Log the VIEW_CHANGE message in message log.

#### Check if new view
* Check the log(View = message.view)
  * If exactly (trigger once) 2xf+1 VIEW_CHANGE messages are logged:
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
  * Request new candidate block by calling `ConsensusAlgo.RequestNewBlock`
  * Calculate the candidate_block_hash, by calling `ConsensusAlgo.CalcBlockHash`

#### Generate New_view_proof
* New_view_proof = All logged VIEW_CHANGE messages in (View = my_state.view) without the Prepared_block

#### Generate New View PRE_PREPARE
> The New View PRE_PREPARE message is a regular PRE_PREPARE message encapsulated in NEW_VIEW. Using the same message format enables a single PREPARE flow.
* Generate PRE_PREPARE
  * Type = PRE_PREPARE
  * Sender = Node public key.
  * View = my_state.view
  * Block_height = my_state.Block_height
  * BlockPair = Candidate_block
  * Block_hash = Candidate_block_hash
  * Signature {Type, Block_height, View, Candidate_block_hash}

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

#### Check message sender
* Discard if the sender is not a valid participant node.
* Discard if the sender is the leader for the view based on `GetCurrentLeader`.

#### Check the message view
* Discard if message.view < my_state.view.

#### Check no duplicate PrePrepare message was logged
* Discard if a PRE_PREPARE message was already logged for the same view and sender.

#### Check New_view_proof
* Verify 2f+1 VIEW_CHANGE messages, from different senders.
* For each VIEW_CHANGE message verify:
  * Type = VIEW_CHANGE
  * Sender is a valid participant node
  * Block_height = NEW_VIEW message.Block_height
  * View = NEW_VIEW message.View
  * Prepared_proof is valid using `ValidatePreparedProof(View_change_view)`
  * Valid signature
* Discard if one of the checks fails.

#### Check encapsulated New View PRE_PREPARE (NVPP) message 
* Check the New View PRE_PREPARE message fields
  * Check Type = PRE_PREPARE
  * Check Sender = NEW_VIEW.Sender
  * Check View = NEW_VIEW.View
  * Check Block_height = NEW_VIEW.Block_height
  * Compre the Block_hash to the hash of teh NewView message.block calculated by calling `ConsensusAlgo.CalcBlockHash`.
  * Check the NVPP signature <!-- TBD move to ConsensusAlgo -->

#### Check encapsulated New View PRE_PREPARE message block
* From all VIEW_CHANGE messages in New_view_proof with Prepared_proof != {} in , find the one with the highest Prepared_proof.View.
* If a VIEW_CHANGE message with Prepared_proof != {} was found: 
  * Check NVPP.Block_hash = highest prepraed view VIEW_CHANGE.Block_hash.
* Else (no VIEW_CHANGE was Prepared)
  * Validate the block by calling `ConsensusAlgo.ValidateBlock`
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
  * Candidate_block_hash = NV_PRE_PREPARE message.Block_hash
  * Candidate_block = NV_PRE_PREPARE message.Block

#### Check if Prepared
* Check the log(View = message.view)
  * If a PRE_PREPARE message AND exactly (trigger once) 2xf PRPARE messages are logged with Block_hash equal to the PRE_PREPARE message.Block_hash. 
    * Set my_state.Prepared
    * Call `OnPrepared`


&nbsp;
## `CheckPBFTProof` 
