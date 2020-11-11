# Merkle Trie 

* A Merkle Trie is a hash tree - each node has a hashcode computed over the nodes state and it’s child node’s hashcodes.
* A Merkle Proof is a Hash Chain comprising of a single path within the Merkle Trie (Hash Tree). As such, it is merely a representation of a string of interconnected nodes contained in the Merkle Trie. 
    * The proof corresponds to a path in the trie.
    * The proof corresponds to a state key prefix, which may be a full state key or a partial state key.
    * In a Merkle Proof array, every pair of adjacent nodes are parent and child the Merkle Tree 
    * A Merkle Proof is valid when: 
        * The first node’s hashcode is identical to the requested block’s Merkle Root 
        * Each child node’s hashcode is identical to the parent reference to it.
        * The key prefix assembled from all nodes up until the one before last in the proof is a prefix of the requested path
    * A merkle proof which is valid may indicate a missing key or indicate the validated hashcode of the value associated with the key  
        * Missing Key indication - if the proof ends with a branch node and the digit in the requested path does not reference any child node, or, if the final node indicates a divergence from the requested key in a way that negates any possibility that the requested key may be present in the trie.
        * Present Key with value hash indication - if the proof ends with a lead node and the assembled key prefix is identical to the requested path
* Prefixes are represented as hexadecimal byte arry.
* Zero values are never represented in the trie because they are not part of the state snapthot
* Each node in a Merkle Trie is addressable via it’s hash-code (cryptographic digest) in a __content-addressed__ scheme - addresses of nodes are derived from their content as a cryptographic hashcode digest.  
* Each block height represents a (possibly) distinct state in it’s virtual chain. For any block height  with a full Merkle trie and a Merkle Trie root corresponding to the matching state image.

## Hash Function 
TBD
## Serialization 
TBD - decide between
 * Option 1 - custom - specify optimized serialization… ?
 * Option 2 - Membuffers - possible issues with size
 * Option 3 - RLP - for code simplicity and gas savings on EVM
 
## Trie Key limitations
### Uniform Key Length 
All trie keys are of a uniform length. This means branch nodes and value nodes are mutually exclusive as there may not be two keys where one is contained in the other.
   * Simplify trie structure by forcing values to be only in leaf nodes
   * This means clients interpreting proofs must be aware of the uniform key length and enforce this knowledge when evaluating proofs
   * This requirement depends on our state addressing scheme. See state storage requirements
## Node types and structure 
### Common properties
All nodes have a prefix part. A prefix part is the current node’s contribution to the key prefix represented by a path in the trie.
* For child nodes (of any type - branch or leaf) if the parent node had an “Even” parity bit, we encode the prefix part with an offset of 4 bits (Even and not Odd because the parent’s branch digit reverses the parity once it is appended to the parent prefix part).
* It's possible for a prefix part to be an empty string.

### Branch nodes
In addition to the prefix part attribute, a branch node has up to 16 children.
Each child reference corresponds to a possible prefix extension. Prefixes are represented as hexadecimal digits in base64 encoding. This is why any possible extension of the current path's prefix may be any one of 16 
* Parity bit always refers to the number of digits from the beginning of the prefix up to and including the current node’s prefix part. 

### Leaf nodes 
Leaf nodes represent values or hashes of values associated with the state key which is represetned by the assembled prefix parts of the entire trie path leading to the leaf node from the root of the trie. 
* Value representation in leaf nodes - decide between (TBD):
    * Keep the entire value in each value node
    * Keep the hash of the value in each value node
    * Keep the hash of the value if the value size is bigger than the hashcode size, and the value otherwise (as is in Ethereum)
    
### Example
```
h0 is the root node in this case. 
in this example all keys are of length 4 bytes. 
Paths are written here in base64 but are in fact byte arrays  
This trie represets these entries:

abc11111 := 1 (trie path: [h0, h1])
ab0bcaff := 2 (trie path: [h0, h2, h3]) 
ab0bcdff := 3 (trie path: [h0, h2, h4])
ab0bceff := 3 (trie path: [h0, h2, h4]) 
        
h0 >> {type: "branch", prefix: "ab", parity: "even", branches: {c: h1, 0: h2}}
    this is the root node.

h1 >> {type: "leaf", prefix: "011111", value: "1"} 
    notice the padding `0` on the left as our parent had "even" parity

h2 >> {type: "branch", prefix: "0bc0", parity: "odd", branches: {a: h3, d: h4, e: h4}} 
    notice the padding `0` on the right as we have 
    only double 'digit' bytes to work with.

h3 >> {type: "leaf", prefix: "ff", value "2"}
    no padding as parent was odd.
    notice this node differs from h4 only in the value. 
    it's enough to give it a distinct hash code. 

h4 >> {type: "leaf", prefix: "ff", value "3"}
    no padding as parent was odd.
    notice how this node is reused in two paths of two keys. 
    

```
## Proof
A Merkle Proof (for a state key/value) is a hash list of node representations. 
* A Proof comprises a single path in the Merkle Trie leading from the root to the last existing node along the requested state entry key.
* Nodes in a proof must represent references to their child nodes as hash codes.  
* The hash of the first node in the list corresponds to the block state hash root
* The value key collectively represented by the trie nodes included in the proof, and constructed by traversing all nodes from the first to the last    
