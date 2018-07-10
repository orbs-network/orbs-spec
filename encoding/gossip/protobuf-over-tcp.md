# Gossip Protobuf over TCP Encoding

* TCP socket between peers with raw data transmitted over it according to the wire format.

* The gossip protocol is a stream of MemBuffers encoded payloads. Each gossip message can be comprised of several payloads that are streamed in succession.

* Wire format (for a single message transmission over the socket):
    * `protocol.MessageHeader` size in bytes (4 bytes, little endian).
    * `protocol.MessageHeader` content (serialized with MemBuffers, [format](../../interfaces/protocol/messages.proto)).
        * This specifies the specific message type.
    * Number of payloads in this specific message (4 bytes, little endian).
    * Payload 1 size in bytes (4 bytes, little endian).
    * Payload 1 content (serialized with MemBuffers, format per specific message).
    * Payload 2 size ...
    * Payload 2 content ...

    Example with `BLOCK_SYNC_RESPONSE` message streaming 1000 block pairs:
    > Note: The payloads can be seen in [`BlockSyncResponseInput`](../../interfaces/services/gossip/block_sync.proto).
     
    * `protocol.MessageHeader` size
    * `protocol.MessageHeader` content
    * number of payloads: 1002
    * `messages.BlockSyncResponseHeader` size
    * `messages.BlockSyncResponseHeader` content
    * `messages.BlockSyncResponse` size
    * `messages.BlockSyncResponse` content
    * `protocol.BlockPair` #1 size
    * `protocol.BlockPair` #1 content
    * `protocol.BlockPair` #2 size
    * `protocol.BlockPair` #2 content
    * ...
    * `protocol.BlockPair` #1000 size
    * `protocol.BlockPair` #1000 content

* As an optimization to reduce buffer copies, messages should be sent over the socket in multiple separate write operations:
  * Each payload should be sent as a separate write.