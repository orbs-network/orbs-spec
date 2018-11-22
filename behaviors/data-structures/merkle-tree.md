# Merkle Trees / Tries

## Merkle trees in Orbs protocol
> Merkle trees or tries are used in Orbs protocol in for multiple purposes:

#### State Merkle Trie
> Maintains authenticated state storage data. Provides both inclusions and exclusion proofs.
* Tree type: Binary Merkle Trie
* Key: SHA256({contract_name, state_key})
* Key size: 32B
* Value: state_value
* Hash: SHA256 (32B)

#### Transaction Merkle Tree
> Maintains the transactions within a block, calculated on block creation and validation.
* Tree type: Binary Merkle tree
* Key: tx_index (the index of the transaction in the block)
* Key size: 4B
* Value: txhash
* Hash: SHA256 (32B)

#### Receipt Merkle Tree
> Maintains the receipts within a block, calculated on block creation and validation.
* Tree type: Binary Merkle tree
* Key: tx_index (the index of the corresponding transaction in the block)
* Key size: 4B
* Value: SHA256(receipt)
* Hash: SHA256 (32B)

## Binary Merkle Trie
> Provides inclusion / exclusion authentication for arbitrary keys.
* Leaf node serialization: {Value, prefix_size, masked_prefix}
* Core node serialization: {left_child_hash, right_child_hash, prefix_size, masked_prefix}

## Binary Merkle tree
> Provides inclusion authentication for sequential keys (0 - max_index).
* Leaf node serialization: {Value}
* Core node serialization: {left_child_hash, right_child_hash}

#### Binary Merkle Proof:
* Structure:
  * List of log(max_index) core nodes' hash.

* Proof validation:
  * hash_state = the Leaf node hash
  * key_bit = 0
  * For each node in the proof starting from the bottom
    * if key[key_bit] = 0
      * node_hash = hash(hash_state, node)
    * Else
      * node_hash = hash(node, hash_state)
    * key_bit++
    * Pop node
  * Compare the hash_state with the tree root.


