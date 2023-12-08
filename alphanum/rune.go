package alphanum

import (
	"sourcecode.social/reiver/go-rfc2396/alpha"
	"sourcecode.social/reiver/go-rfc2396/digit"
)

func RuneIs(value rune) bool {
	return alpha.RuneIs(value) || digit.RuneIs(value)
}
