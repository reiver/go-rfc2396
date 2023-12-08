package alpha

import (
	"sourcecode.social/reiver/go-rfc2396/lowalpha"
	"sourcecode.social/reiver/go-rfc2396/upalpha"
)

func RuneIs(value rune) bool {
	return lowalpha.RuneIs(value) || upalpha.RuneIs(value)
}
