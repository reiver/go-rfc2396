package decnum

import (
	"sourcecode.social/reiver/go-rfc2396/digit"
)

// Bytes return a 'decnum'.
//
//	decnum        = 1*digit
//
// 'decnum' doesn't appear in IETF RFC-2396 by name. But instead is part of `IPv4address`.
//
//	IPv4address   = 1*digit "." 1*digit "." 1*digit "." 1*digit
//
// I called this a 'decnum' (i.e., "decimal number").
func Bytes(p []byte) (result []byte, rest []byte, ok bool) {

	length := len(p)

	if length <= 0 {
		return nil, nil, false
	}

	{
		p0 := p[0]

		if !digit.ByteIs(p0) {
			return nil, p, false
		}
	}

	{
		for i,b := range p {
			if !digit.ByteIs(b) {
				return p[:i], p[i:], true
			}
		}

		return p, nil, true
	}
}
