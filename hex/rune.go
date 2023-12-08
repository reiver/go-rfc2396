package hex

import (
	"sourcecode.social/reiver/go-rfc2396/digit"
)

func RuneIs(value rune) bool {
	switch {
	case digit.RuneIs(value):
		return true
	case 'A' <= value && value <= 'F':
		return true
	case 'a' <= value && value <= 'f':
		return true
	default:
		return false
	}
}
