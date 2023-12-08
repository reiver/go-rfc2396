package alpha

import (
	"sourcecode.social/reiver/go-rfc2396/lowalpha"
	"sourcecode.social/reiver/go-rfc2396/upalpha"
)

func ByteIs(value byte) bool {
	return lowalpha.ByteIs(value) || upalpha.ByteIs(value)
}
