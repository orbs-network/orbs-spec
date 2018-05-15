
## Address Scheme

Addressing in the Orbs platform are based on a universal signature and addressing scheme. This method allows applications and users to use a range of addressing schemes side by side, specifying the type of the address next to the address itself.

Every public address in Orbs platform has the form of {Network ID, Address Scheme, Address}. *Network ID* is a 1-byte field that determines that network type. *Address Scheme* is a 1-byte that determines the signature scheme and the address format.

### Network ID Encoding

| Network  | Value | Value (hex) |
| :------: | :---: | :---------: |
| Main net | M     | 4d          |
| Test net | T     | 54          |

### Address Scheme Encoding

| Address Scheme | Value (hex) |
| :------------: | :---------- |
| Rev1           | 00          |

### Address Scheme:

1. Start with a 32-byte Ed25519 public key:

    Public key: `8d41d055d00459be37f749da2caf87bd4ced6fafa335b1f2142e0f44501b2c65`

2. Set the 3-byte Virtual Chain ID:

    Virtual Chain ID: `640ed3`

    Note: The Virtual Chain ID MSB value should be > `07` in order to obtain a BASE58 encoded address with a leading non-zero value.

3. Calculate the Account ID by calculating the RIPEMD160 hash of the SHA256 of the public key:

    SHA256 of the public key: `40784b5b15e6bb364263dbb598f262bc5c5b4c18a34806ca70be180c3d995e0d`

    RIPEMD160 of the SHA256: `c13052d8208230a58ab363708c08e78f1125f488`

4. Prepend the Network ID, Address Scheme and Virtual Chain ID to the Account ID:

    Network ID + Address Scheme + Virtual Chain ID + Account ID: `4d00640ed3c13052d8208230a58ab363708c08e78f1125f488`

5. Calculate the CRC32 checksum of the result:

    Checksum: `8113211c`

    Raw public address: `4d00640ed3c13052d8208230a58ab363708c08e78f1125f4888113211c`

6. Encode the raw public address to Base58:

    Each of the following address parts is encoded separately:

    a. Network ID: `M`

    b. Address Scheme: `00`

    c. Virtual Chain ID + Account ID + Checksum: `EXMPnnaWFqRyVxWdhYCgGzpnaL4qBy4R42dE3`

    Public address: `M00EXMPnnaWFqRyVxWdhYCgGzpnaL4qBy4R42dE3`