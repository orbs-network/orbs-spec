# Serialization Format

* All data types passed in the system must be serializable to binary format.
  * Data types are often signed or hashed which means the binary representation is signed or hashed. 
  * Data types are stored in persistent storage (eg. block storage) in their original signed form.

* Serialization format should be as canonical and deterministic as possible.
  * Different nodes may generate the same complex field separately and compare it using hash only.
  
* Serialization format is [MemBuffers](https://github.com/orbs-network/membuffers).