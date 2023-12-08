package unreserved

import (
	"sourcecode.social/reiver/go-rfc2396/alphanum"
	"sourcecode.social/reiver/go-rfc2396/mark"
)

func RuneIs(value rune) bool {
	return alphanum.RuneIs(value) || mark.RuneIs(value)
}
