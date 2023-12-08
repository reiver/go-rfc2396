package unreserved

import (
	"sourcecode.social/reiver/go-rfc2396/alphanum"
	"sourcecode.social/reiver/go-rfc2396/mark"
)

func ByteIs(value byte) bool {
	return alphanum.ByteIs(value) || mark.ByteIs(value)
}
