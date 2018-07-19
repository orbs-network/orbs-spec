# Gossip MemBuffers over TCP Encoding

* TCP socket between peers with raw data transmitted over it according to the wire format.

* The gossip protocol is a stream of MemBuffers encoded payloads. Each gossip message can be comprised of several payloads that are streamed in succession.
  * The number of payloads in the steam is encoded in the beginning.
  * Each chunk (MemBuffer) in the stream is always padded to 4 bytes.

* Wire format (for a single message stream over the socket):
    * Number of payloads (4 bytes, little endian).
    * Payload 1 size in bytes (4 bytes, little endian).
    * Payload 1 content (serialized with MemBuffers, format per specific message).
    * Payload 2 size ...
    * Payload 2 content ...

    Example with `BLOCK_SYNC_RESPONSE` message streaming 100 transaction blocks (each with 3 transactions):
    > Note: The message format can be seen in [`BlockSyncResponseMessage`](../../interfaces/protocol/gossipmessages/block_sync.proto).
     
    * `num_payloads` (603) (3 payloads before blocks + 100 * 6 payloads per block)
    * `gossipmessages.Header` size
    * `gossipmessages.Header` content (`type`: `BLOCK_SYNC_RESPONSE`)
    * `gossipmessages.BlockSyncRange` size (`block_type`: `BLOCK_TYPE_TRANSACTIONS_BLOCK`)
    * `gossipmessages.BlockSyncRange` content
    * `gossipmessages.SenderSignature` size
    * `gossipmessages.SenderSignature` content
    * `protocol.TransactionsBlockHeader` size (block #1)
    * `protocol.TransactionsBlockHeader` content (block #1)
    * `protocol.TransactionsBlockMetadata` size (block #1)
    * `protocol.TransactionsBlockMetadata` content (block #1)
    * `protocol.SignedTransaction` size (block #1)
    * `protocol.SignedTransaction` content (block #1)
    * `protocol.SignedTransaction` size (block #1)
    * `protocol.SignedTransaction` content (block #1)
    * `protocol.SignedTransaction` size (block #1)
    * `protocol.SignedTransaction` content (block #1)
    * `protocol.TransactionsBlockProof` size (block #1)
    * `protocol.TransactionsBlockProof` content (block #1) (total 6 payloads for this block)
    * `protocol.TransactionsBlockHeader` size (block #2)
    * `protocol.TransactionsBlockHeader` content (block #2)
    * `protocol.TransactionsBlockMetadata` size (block #2)
    * `protocol.TransactionsBlockMetadata` content (block #2)
    * `protocol.SignedTransaction` size (block #2)
    * `protocol.SignedTransaction` content (block #2)
    * `protocol.SignedTransaction` size (block #2)
    * ...

* As an optimization to reduce buffer copies, messages should be sent over the socket in multiple separate write operations:
  * Each payload should be sent as a separate write.