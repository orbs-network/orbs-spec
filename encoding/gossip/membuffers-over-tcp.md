# Gossip MemBuffers over TCP Encoding

* TCP socket between peers with raw data transmitted over it according to the wire format.

* The gossip protocol is a stream of MemBuffers encoded payloads. Each gossip message can be comprised of several payloads that are streamed in succession.
  * Each chunk (MemBuffer) in the stream is always padded to 4 bytes.

* Wire format (for a single message transmission over the socket):
    * `gossipmessages.Header` size in bytes (4 bytes, little endian).
    * `gossipmessages.Header` content (serialized with MemBuffers, [format](../../interfaces/protocol/gossipmessages/all.proto)).
        * Specifies the specific message type under `gossipmessages.Header.topic`.
        * Specifies the number of payloads under `gossipmessages.Header.num_payloads`.
    * Payload 1 size in bytes (4 bytes, little endian).
    * Payload 1 content (serialized with MemBuffers, format per specific message).
    * Payload 2 size ...
    * Payload 2 content ...

    Example with `BLOCK_SYNC_RESPONSE` message streaming 1000 block pairs:
    > Note: The payloads can be seen in [`BlockSyncResponseInput`](../../interfaces/services/gossiptopics/block_sync.proto).
     
    * `gossipmessages.Header` size
    * `gossipmessages.Header` content (`type`: `BLOCK_SYNC_RESPONSE`, `num_payloads`: 1002)
    * `gossipmessages.BlockSyncResponseHeader` size
    * `gossipmessages.BlockSyncResponseHeader` content
    * `gossipmessages.BlockSyncResponse` size
    * `gossipmessages.BlockSyncResponse` content
    * `protocol.BlockPair` #1 size
    * `protocol.BlockPair` #1 content
    * `protocol.BlockPair` #2 size
    * `protocol.BlockPair` #2 content
    * ...
    * `protocol.BlockPair` #1000 size
    * `protocol.BlockPair` #1000 content

* As an optimization to reduce buffer copies, messages should be sent over the socket in multiple separate write operations:
  * Each payload should be sent as a separate write.