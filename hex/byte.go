package hex

import (
	"sourcecode.social/reiver/go-rfc2396/digit"
)

func ByteIs(value byte) bool {
	switch {
	case digit.ByteIs(value):
		return true
	case 'A' <= value && value <= 'F':
		return true
	case 'a' <= value && value <= 'f':
		return true
	default:
		return false
	}
}
