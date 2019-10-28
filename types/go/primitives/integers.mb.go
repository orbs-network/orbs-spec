// AUTO GENERATED FILE (by membufc proto compiler v0.0.32)
package primitives

import (
	"bytes"
	"fmt"
)

type Uint8 uint8

func (x Uint8) String() string {
	return fmt.Sprintf("%x", uint8(x))
}

func (x Uint8) Equal(y Uint8) bool {
	return x == y
}

func (x Uint8) KeyForMap() uint8 {
	return uint8(x)
}

type Uint16 uint16

func (x Uint16) String() string {
	return fmt.Sprintf("%x", uint16(x))
}

func (x Uint16) Equal(y Uint16) bool {
	return x == y
}

func (x Uint16) KeyForMap() uint16 {
	return uint16(x)
}

type Uint128 []byte

func (x Uint128) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x Uint128) Equal(y Uint128) bool {
	return bytes.Equal(x, y)
}

func (x Uint128) KeyForMap() string {
	return string(x)
}

type Uint256 []byte

func (x Uint256) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x Uint256) Equal(y Uint256) bool {
	return bytes.Equal(x, y)
}

func (x Uint256) KeyForMap() string {
	return string(x)
}
