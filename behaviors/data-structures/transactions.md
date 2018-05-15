# Transactions
## Transaction Structure

- **Header**
  - `Contract Address` - The universal contract address (see [Orbs Addressing Scheme](./addresses.md) for more details).
  - `Sender Address` - The universal sender address (see [Orbs Addressing Scheme](./addresses.md) for more details).
    > _Note_: the `sender address` and `contract address` should have the exact same Network ID and Virtual Chain ID. Otherwise the transaction is considered as invalid.
  - `Timestamp` - The number of milliseconds between midnight of January 1, 1970 and the transaction creation time
  - `Version` - The transaction version (Currently = 1)
- **Payload** - The payload specifies the contract's method to be called, and the arguments to be passed with the call. Structured as a JSON-serialized string containing:
  - `methodName` - The contract's method name
  - `methodArgs` - An ordered array of string or number values to be passed as arguments to the method call

  _E.g._: ```{"methodName": "transfer": "methodArgs":["200","John"]}```
- **Signature Data**
  - `Public Key` - The public key of the sender
  - `Signature` - ED25519 signature of the transaction, signed by the sender's private key


## How to calculate transaction hash and transaction ID (aka TXID)?

- A transaction is hashed using SHA-256.
- The message to be hashed is a canonical JSON representation of the transaction header + payload.
  - The JSON should be structured as follows:
    - **`header`**
      - `contractAddressBase58` (String) -Base58 encoding of the contract address, [as described here](./addresses.md).
      - `senderAddressBase58` (String) - Base58 encoding of the sender address, [as described here](./addresses.md).
      - `timestamp` (String) - String representation of the timestamp
      - `version` (Number) - The transaction version
    - **`payload`** (String)**
  - It should be formatted **without any line breaks or whitespaces**.
  - The fields should be **ordered alphabetically**.
  - It should use **ASCII** encoding.

  _E.g.:_ ```{"header":{"contractAddressBase58":"T00EXMPjG82qZxHM5Kz1ghZFizRDktbVVuQmndci","senderAddressBase58":"T00EXMPYkRQRgcccBZcr2sWihYroPs5Ni4HjzuAH","timestamp":"1526298045636","version":0},"payload":"{\"method\":\"transfer\",\"args\":[\"T00EXMPYkRQRgcccBZcr2sWihYroPs5Ni4HjzuAH\",10]}"}```
 - **TXID** is the hexadecimal representation of the transaction hash.


## How to sign a transaction?
Every transaction should contain a Signature Data part, containing:

- Signature - ED22519 signature of the transaction hash.
- Public key - The public key part of the sender's ED22519 key-pair used for signing the transaction. The public key is used for verifying the transaction, whereas the private key, used for signing the transaction, **should be kept secret**.
  > _Note:_ The `Account ID`, which is part of the `Sender Address`, is the hash (SHA-256) of the `Public Key`. If  **Account ID != SHA256(PublicKey)** the transaction is considered as invalid.
