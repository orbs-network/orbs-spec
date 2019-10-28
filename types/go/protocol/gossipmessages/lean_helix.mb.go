// AUTO GENERATED FILE (by membufc proto compiler v0.0.32)
package gossipmessages

import (
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixMessage (non serializable)

type LeanHelixMessage struct {
	Content   primitives.LeanHelixMessageContent
	BlockPair *protocol.BlockPairContainer
}

func (x *LeanHelixMessage) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Content:%s,BlockPair:%s,}", x.StringContent(), x.StringBlockPair())
}

func (x *LeanHelixMessage) StringContent() (res string) {
	res = fmt.Sprintf("%s", x.Content)
	return
}

func (x *LeanHelixMessage) StringBlockPair() (res string) {
	res = x.BlockPair.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
