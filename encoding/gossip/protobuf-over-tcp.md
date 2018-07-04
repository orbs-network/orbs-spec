# Gossip Protobuf over TCP Encoding

* See example implementation for inspiration: https://github.com/StabbyCutyou/buffstreams
  * Rely on the common system Protobuf serialization format (MemBuffers).
  
* As an optimization, messages should be sent over the socket in two separate write operations:
  * First write operation for the serialized message header.
  * Second write operation for the serialized message data.
  
* Every serialized write operation is prefaced with 4 bytes (little endian `uint32`) of the serialized data size.