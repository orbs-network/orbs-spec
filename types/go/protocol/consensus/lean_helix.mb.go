// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package consensus

import ()

/////////////////////////////////////////////////////////////////////////////
// enums

type LeanHelixMessageType uint16

const (
	LEAN_HELIX_MESSAGE_TYPE_RESERVED   LeanHelixMessageType = 0
	LEAN_HELIX_MESSAGE_TYPE_LEAN_HELIX LeanHelixMessageType = 1
)

func (n LeanHelixMessageType) String() string {
	switch n {
	case LEAN_HELIX_MESSAGE_TYPE_RESERVED:
		return "LEAN_HELIX_MESSAGE_TYPE_RESERVED"
	case LEAN_HELIX_MESSAGE_TYPE_LEAN_HELIX:
		return "LEAN_HELIX_MESSAGE_TYPE_LEAN_HELIX"
	}
	return "UNKNOWN"
}
