# Receipt Proof for Ethereum Interoperability

&nbsp;
## Proof data
* protocol.ResultsBlockHeader header
* protocol.ResultsBlockProof block_proof
* protocol.TransactionReceipt receipt
* primitives.merkle_tree_proof receipt_proof

&nbsp;
### ResultsBlockHeader 

| Feild         | Offset        | Size          | Encoding      | Notes         |
|:------------- |:-------------:|:-------------:|:-------------:|:--------------|
| protocol_version | 0 | 4 | uint32 | |
| virtual_chain_id | 4 | 8 | uint64 | |
| network_type | TBD | 2 | enum (4 bytes) | |
<!--
| block_height | 8 | 8 | uint64 | To be replaced with version |
-->
| timestamp | 16 | 8 | uint64 | unix 64b time |
| receipt_merkle_root | 64 | 32 | bytes (32B) | |

* total size: 204B

#### ResultsBlockHeader Validation
* Calcualte `result_header_hash` = SHA256 hash of the `ResultsBlockHeader`.
* Check that the `protocol_version` matches the configuration.
* Check that the `virtual_chain_id` matches the configuration.
* Check that the `network_type` matches the configuration.
<!--
  * Check that the `timestamp` >`time` - `timeout_value`. 
-->
* Extract the `receipt_merkle_root`.

&nbsp;
### ResultsBlockProof

| Feild         | Offset        | Size          | Encoding      | Notes         |
|:------------- |:-------------:|:-------------:|:-------------:|:--------------|
| block_proof_version | 0 | 4 | uint32 | |
| transactions_block_hash length | 4 | always 4 | reserved |
| transactions_block_hash | 8 | 32 | bytes (32B)| |
| oneof + nesting | 40 | 12 | reserved | oneof + proof + blockref | 
| blockref_message | 52 | 52 | bytes (52B) |  |
| block_hash | 72 | 32 | bytes (32B)|  |
| node_pk_sig nesting | 104 + 100n | reserved | |
| node_pk_length | 108 + 100n | 4 | always 20 | reserved |
| node_pk | 112 + 100n | 20 | bytes (20B) | Ethereum address |
| node_sig_length | 132 + 100n | 4 | always 65 | reserved |
| node_sig | 136 + 100n | 65 | bytes (65B) |

#### ResultsBlockProof Validation
* Calcualte `calcualted_block_hash` = SHA256(`transactions_block_hash`,`result_header_hash`).
* Match `block_hash` to `calcualted_block_hash`.
* Calcualate `blockref_hash` = SHA256(`blockref_message`).
* For each node:
  * Extract `node_pk` and `node_sig`
  * Check that `node_pk` is part of the federation by querying the `federation contract`
	* based on the given `block_proof_version`.
  * Check that the signature matches the `blockref_hash`.
* Get the proof threshold from the `federation contract` and check that the number of valid nodes mathes the threshold.

&nbsp;
### Receipt

| Feild         | Offset        | Size          | Encoding      | Notes         |
|:------------- |:-------------:|:-------------:|:-------------:|:--------------|
| execution_result | 36 | 4 | enum | 0x1 indicates success |
| event length  | 40 | 4 | uint32 | |
| event data  | 44 | variable | bytes | |



* Total size: variable size. 

#### MerkleProof Validation
* Calculate the `receipt_hash` = SHA256(Receipt)
* Extract the `event`

&nbsp;
### Receipt MerkleProof

| Feild         | Offset        | Size          | Encoding      | Notes         |
|:------------- |:-------------:|:-------------:|:-------------:|:--------------|
| total_length  | 4 | 8 | uint32 | |
| merkle_node   | 8 + 32n | 32 | bytes (32B)| |

#### MerkleProof Validation
* Binary merkel tree validation:
* Extract the key = `receipt_index` and the `merkle_node`s
* Init: 
  * `node_bit` = 0.
  * `hash_state` = `receipt_hash`.
* For each node in the `merkle_node`s:
  * If `receipt_index[node_bit:node_bit] = 0`
    * `hash_state` = SHA256(`hash_state`,`merkle_node`).
  * Else
	* `hash_state` = SHA256(`merkle_node`,`hash_state`).
  * `node_bit` += 1.
* Check `hash_state` = `receipt_merkle_root`.

&nbsp;
### Autonomous Swap Event Data

| Feild         | Offset        | Size          | Encoding      | Notes         |
|:------------- |:-------------:|:-------------:|:-------------:|:--------------|
| contract name length (N)| 0 | 4 | uint32 | |
| contract name | 4 | N | string | |
| event_id | 4+N | 4 | enum | 0x1 indicates TRANSFERED_OUT|
| tuid | 8+N | 8 | uint64 | |
| ethereum_address length| N+16 | 4 | always 20 | reserved |
| ethereum_address | N+20 | 20 | bytes (20B) | |
| tokens length | N+40 | 4 | always 32 | reserved |
| tokens | N+44 | 32 | uint256 | |

&nbsp;
#### Event Data Validation
* Check that the `contract name` matches the configuration.
* Check `event_id` = TRANSFERED_OUT (0x1)
* Check that the`ethereum_address` is valid - TBD.

