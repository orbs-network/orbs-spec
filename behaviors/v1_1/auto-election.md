# Auto-election (remove Good Samaritan)

## High Level Flow
* Processing and mirroring by a single function to be called by the Election trigger.
* No mirroring grace period (0 blocks, processing may start once all the transactions were mirrored.
* Stateful mirroring, maintaining the last processed Ethereum block. 
* Deprecating the external mirroring functions.
* Note: election calculation still take multiple blocks, once Ethereum proxy is implemented, only the calculation part is needed.

## ProcessAndMirror 
* Check if processing has started. 
  * If so call processing PROCESSING_CALLS_PER_BLOCK times
  * Else mirror Ethereum Data

#### Mirror Ethereum Data
* Activate every x seconds based on Orbs block time (`env.GetBlockTimestamp`):
* Check if the election has passed, if so
  * call processing, stop.
* Get the delegation by contract events starting from the last_processed_block_number. + 1.
  * If there are relevant events process them using mirrorDelegationData.
* Read the ERC20 events starting from the last_processed_block + 1.
  * If amount == 0.07 ORBS, process them using mirrorDelegationData.
* Store the last_processed_block_number.  

## New SDK
* Ethereum.FilterLogs

