// AUTO GENERATED FILE (by membufc proto compiler v0.0.19)
package primitives

import (
	"fmt"
	"bytes"
)

type Sha256 []byte

func (x Sha256) String() string {
  return fmt.Sprintf("%x", []byte(x))
}

func (x Sha256) Equal(y Sha256) bool {
  return bytes.Equal(x, y)
}

func (x Sha256) KeyForMap() string {
  return string(x)
}

type MerkleSha256 []byte

func (x MerkleSha256) String() string {
  return fmt.Sprintf("%x", []byte(x))
}

func (x MerkleSha256) Equal(y MerkleSha256) bool {
  return bytes.Equal(x, y)
}

func (x MerkleSha256) KeyForMap() string {
  return string(x)
}

type Sha3256 []byte

func (x Sha3256) String() string {
  return fmt.Sprintf("%x", []byte(x))
}

func (x Sha3256) Equal(y Sha3256) bool {
  return bytes.Equal(x, y)
}

func (x Sha3256) KeyForMap() string {
  return string(x)
}

type Ed25519Sig []byte

func (x Ed25519Sig) String() string {
  return fmt.Sprintf("%x", []byte(x))
}

func (x Ed25519Sig) Equal(y Ed25519Sig) bool {
  return bytes.Equal(x, y)
}

func (x Ed25519Sig) KeyForMap() string {
  return string(x)
}

type Ed25519PublicKey []byte

func (x Ed25519PublicKey) String() string {
  return fmt.Sprintf("%x", []byte(x))
}

func (x Ed25519PublicKey) Equal(y Ed25519PublicKey) bool {
  return bytes.Equal(x, y)
}

func (x Ed25519PublicKey) KeyForMap() string {
  return string(x)
}

type Ed25519PrivateKey []byte

func (x Ed25519PrivateKey) String() string {
  return fmt.Sprintf("%x", []byte(x))
}

func (x Ed25519PrivateKey) Equal(y Ed25519PrivateKey) bool {
  return bytes.Equal(x, y)
}

func (x Ed25519PrivateKey) KeyForMap() string {
  return string(x)
}

type Ripmd160Sha256 []byte

func (x Ripmd160Sha256) String() string {
  return fmt.Sprintf("%x", []byte(x))
}

func (x Ripmd160Sha256) Equal(y Ripmd160Sha256) bool {
  return bytes.Equal(x, y)
}

func (x Ripmd160Sha256) KeyForMap() string {
  return string(x)
}

type Bls1Sig []byte

func (x Bls1Sig) String() string {
  return fmt.Sprintf("%x", []byte(x))
}

func (x Bls1Sig) Equal(y Bls1Sig) bool {
  return bytes.Equal(x, y)
}

func (x Bls1Sig) KeyForMap() string {
  return string(x)
}

type Bls1PublicKey []byte

func (x Bls1PublicKey) String() string {
  return fmt.Sprintf("%x", []byte(x))
}

func (x Bls1PublicKey) Equal(y Bls1PublicKey) bool {
  return bytes.Equal(x, y)
}

func (x Bls1PublicKey) KeyForMap() string {
  return string(x)
}

type Bls1PrivateKey []byte

func (x Bls1PrivateKey) String() string {
  return fmt.Sprintf("%x", []byte(x))
}

func (x Bls1PrivateKey) Equal(y Bls1PrivateKey) bool {
  return bytes.Equal(x, y)
}

func (x Bls1PrivateKey) KeyForMap() string {
  return string(x)
}

type BloomFilter []byte

func (x BloomFilter) String() string {
  return fmt.Sprintf("%x", []byte(x))
}

func (x BloomFilter) Equal(y BloomFilter) bool {
  return bytes.Equal(x, y)
}

func (x BloomFilter) KeyForMap() string {
  return string(x)
}


