# Voting Mirroring Agent 
> The Voting Mirroring Agent is responsible for mirroring voting and delegation transactions and to initiate voting processing. Anyone may operate an agent. The network does not trust the agent or assumes an honest behaviour and only requires the existence of at least one agent that follows the protocol.

## Parameters
* `VOTE_MIRRORING_PERIOD_LENGTH_IN_BLOCKS` - the time after an election event when processing can start. Mirroring may continue until the end of the mirroring period.
* `VOTING_VALIDITY_TIME` - The number of Ethereum blocks during which a Guardian's voting is valid
  * Default: 40320 (~ 7 days)
`PROCESSING_TRANSACTIONS_BATCH_SIZE` - The maximum number of `Processing` transactions that can be sent concurrently (in the same block).
  * Default: 20 (TBD)
* `ERC20_CONTRACT_ADDRESS` - The address of Orbs ERC20 contract
* `VOTING_CONTRACT_ADDRESS` - The address of the voting and delegation contract

## Delegation and Voting mirroring

### Relevant transactions:
* Delegation transactions (triggered a delegation event): 
  * `VOTING_CONTRACT_ADDRESS` contract event `Delegate(address indexed delegator, address indexed to, uint delegationCounter)`.
  * `ERC20_CONTRACT_ADDRESS` contract event `Transfer(address indexed from, address indexed to, uint value)` **with value = 7**. 
* Voting transactions (triggered a vote out event):
  * `VOTING_CONTRACT_ADDRESS` contract event `VoteOut(address indexed voter, bytes20[] validators, uint voteCounter)`.

### Mirroring schedule 
* Transactions that were recorded on Ethereum prior to the election event (inclusive) may be mirrored until the end of the mirroring period. (Until processing has started)
* Get from Orbs:
  * current Ethereum block number (final) by calling `Election.getCurrentEthereumBlockNumber()`
  * current election block number by calling `Election.getCurrentElectionBlockNumber()`
  * first processing block number by calling `Election.getProcessingStartBlockNumber()`
* If current Ethereum block number > current election block number, start mirroring delegation and voting transactions.
  * Mirror only transactions with block number <= current election block number
* Notes: 
  * Transactions with block number <= current Ethereum block number AND block number <= current election block number may be mirrored without waiting to the current election block number.

### Mirror - **To All Active Virtual Chains**
> Note: mirror transactions may be processed out of order.
* Mirror delegations since the last block that delegations were mirrored from on this virtual chain (locally) until the current Ethereum block number (inclusive).
  * Store locally the last block that was covered by the delegation mirroring. 
  * If no delegation stored, there's no harm in collecting past delegation data.
    * The Election contract will return error on the duplicate mirrored trasnactions.
  * Use `Election.mirrorDelegation(hexEncodedEthTxHash)` or `Election.mirrorDelegationByTransfer(hexEncodedEthTxHash)`.
* Mirror all VoteOut transactions with block_height > current election block number - `VOTING_VALIDITY_TIME`. Note: no harm in mirror older votes.
  * Use `Election.mirrorVote(hexEncodedEthTxHash)`

## Processing

### Processing schedule 
* Get from Orbs:
  * current Ethereum block number (final) by calling `Election.getCurrentEthereumBlockNumber()`
  * first processing block number by calling `Election.getProcessingStartBlockNumber()`
* If current Ethereum block number > first processing block number, start sending processing transactions.
* Notes: 
  * When the current block number is further from the first processing block number, its ok to sample every few min. When getting close its recommended to increase to the sampling time to allow the processing to complete as soon as possible after the elections.
  
### Processing - **To All Active Virtual Chains concurrently** 
> Note processing transactions may be batched for up to `PROCESSING_TRANSACTIONS_BATCH_SIZE`
* Send a batch of `Election.processVoting()` transactions
  * If output argument = `1` - Election completed wait for next mirroring.
  * If output argument = `0` - Continue with next batch.
* Note: there's no harm in sending unnecessary `processVoting()` transactions. The additional transactions will return error as the processing for the next elections has not started yet.
* Alternatively may check that processing ended by querying `Election.getEffectiveElectionBlockNumber()` and see that the processed elections took effect.

## System requirements 
* At least 2 processors should be running at all time.
* New Virtual Chain deployment 
  * Upon addition of a new virtual chain, the virtual chain should first be added to the processors before deploying it.
  * Once the virtual chain is up, send `Election.initFirstElection()` to set the first election.
    * Will also be set automatically upon the first vote / delegation mirroring.


 

