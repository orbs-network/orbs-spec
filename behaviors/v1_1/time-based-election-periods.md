# Time based election period
> Elections are set every 20,000 Ethereum blocks (targeting approximately 3 days) and a Guardian vote is valid for 45,500 (targeting approximately 7 days).
> Recently the average Ethereum block time varies and continues increase is expected due to the difficulty bomb.
> As a result, election times and rewards are less predictable.
> The specification update changes the parameters to time based instead of number of blocks based.

* Election time 
  * 3 days starting from the last election time. 
  * Election block = the first Ethereum block with timestamp > Election time.
* Delegation validity
  * Delegations are considered for the election if the Delegation block timestamp <= election time.
* Election block height:
  * The `election block number` is set to the first Ethereum block with timestamp > election time.
* Vote validity:
  * A vote is valid for an election if the vote block timestamp + 7 days <= election time.
* Rewards:
  * Calculated based on a fixed 3 days between elections, i.e. 365.24 / 3 Elections per year.

# Feature impact
## Elections contract impact
* mirrorDelegation
  * Check validity by time instead of block height.
    * May maintain database by block number or timestamp.
* Processing state machine
  * Processing start determined by time - 2 hours after election time.
* Vote recording 
  * Check validity by time instead of block height.
* Rewards
  * Replace the factor with 365.24/3
* Election schedule / first election
  * The next election is the first election in the election schedule following the first time based election.
  * Get the current Ethereum block time using `Ethereum.GetBlockTime()`
  * If current Ethereum block time < `FIRST_TIME_BASED_ELECTION` 
    * next election time = `FIRST_TIME_BASED_ELECTION`
  * Else
    * next election time = `FIRST_TIME_BASED_ELECTION` + N x `ELECTION_PERIOD` such that next election time is after the current Ethereum block time.
  * Note: will be deprecated after the management chain implementation.  

*General*

  * getEffectiveElectionBlockNumber - as is
  * *getCurrentElectionBlockNumber()* - deprecated, panic("Current election time: ...")
  * *getNextElectionBlockNumber()* - deprecated, panic("Next election time: ...")
  * *getElectionPeriod()* - deprecated, panic("Election period is % days")
  * getNumberOfElections() - as is
  * getElectedValidatorsOrbsAddress() - as is
  * getElectedValidatorsEthereumAddress() - as is
  * getCurrentEthereumBlockNumber() - as is
  * *getProcessingStartBlockNumber()* - deprecated, panic("Processing start time: ...")
  * *getMirroringEndBlockNumber()* - deprecated, panic("Mirroring end time: ...")
  * *getCurrentElectionTime()* - New function
    * Returns the time of the election in process
  * *getNextElectionTime()* - New function
    * Returns the time of the next election
  * *getProcessingTime()* - New function
    * Returns next the processing time

*Extended Last Election Results* - as is

*Historical data*
  * getElectedValidatorsOrbsAddressByBlockHeight(orbs_block_height) - as is
  * getElectedValidatorsEthereumAddressByBlockNumber(block_number) - as is
  * getElectedValidatorsOrbsAddressByIndex(index) - as is
  * getElectedValidatorsEthereumAddressByIndex(index) - as is
  * getElectedValidatorsOrbsAddressByIndex(index) - as is
  * getElectionEventBlockNumberByIndex(index) - as is
  * getElectionResultsBlockHeightByIndex(index) - as is

*Rewards* - as is
       
##  SDK / Crosschain connector new function

### Ethereum.GetBlockTimeByNumber(uint32 blockNumber) : uint64 (epoch time)
* Input: blockNumber
  * If blockNumber > block number for finality, fail the call.
* Output: the epoch time of the block (in seconds), for example GetEthereumBlockTime(7430000) = 1553411888
* Usage: validity of mirrored delegation.

### Ethereum.GetBlockNumberByTime(epoch_time : uint32) : uint32
* Input: epoch time in sec
  * If epoch_time < current time + ETHEREUM_FINALITY_TIME_COMPONENT, fail the call.
* Output: 
  * Returns the maximum Ethereum block number such that its timestamp < epoch_time.
  * If result < block number for finality, fail the call.
* Usage: retrieve the election block number.

### Ethereum.GetBlockTime() : uint64 (epoch time)
* Output: the epoch time of the current Ethereum block (the block for finality)
* Usage: processing start (may be replaced by GetBlockNumber + GetBlockTimeByNumber, provided for completeness)

## Voting UIs
* Replace next election block with next election time.
* Update vote validity for next election logic
