# TODO

* Decide on naming of `CallContract` and `SendTransaction` as our basic actions

### Call Contract Flow

* Why do we need to even sign calls?

### Processor (Native)

* Interface

* Who checks permissions, is it the processor or the virtual machine?

## Send Transaction Flow

* Verify that the transaction pool can add transactions even when not synchronized

* Are transactions signed by clients in their textual form? If so we need to store them like so on the blockchain or be able to translate back and forth. If we have a binary transaction format, what happens if we update the format? Clients may not update, do we need to support the old format forever? how is this versioned?

## Versioning

* Transaction format version

* Protocol version

## Init
