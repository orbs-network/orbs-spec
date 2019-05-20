# Signing service

Signing service is a service that contains node private key and allows to sign certain payload with this key.

It preferably runs as a separate process in a separate container. The idea behind it is secure isolation of private key from other parts of the system.

## Interacts with services

* `Conseusus Algo` - uses it to sign blocks.
* `Transaction Pool` - uses it to sign transaction forwarding.

## APIs

V1:

* `Sign` - signs sha256 hash.

V2:

* `SignTransactionBatch` - signs transactions batch hash.
* `SignBlock` - signs a block hash.
* `GetPublicKey` - returns public key.

## Provisioning

Preferred way of provisioning would be decoupling of signer service from the main process and running a single instance per node (X vchain containers + 1 signer container).

Considering low throughput (we only sign hashes) this is both possible and advisable.

Separate versioning of the signing service might be a good idea (or different API versions can be supported by the same binary).

## Implementation

Phase 1:

* implementation of Signer interface (in-process and separate process)
* no authentication
* stateless

Phase 2:

* add authentication per chain via public/private keys (key pair per chain)

Phase 3 (?):

* update authentication keys on the fly without reloading the service
* stateful
