package alphanum

import (
	"sourcecode.social/reiver/go-rfc2396/alpha"
	"sourcecode.social/reiver/go-rfc2396/digit"
)

func ByteIs(value byte) bool {
	return alpha.ByteIs(value) || digit.ByteIs(value)
}
