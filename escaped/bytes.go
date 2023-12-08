package escaped

import (
	"sourcecode.social/reiver/go-rfc2396/hex"
)

func Bytes(p []byte) (result []byte, rest []byte, ok bool) {
	length := len(p)

	if length < 3 {
		return nil, p, false
	}

	p0, p1, p2 := p[0], p[1], p[2]

	if '%' != p0 {
		return nil, p, false
	}
	if !hex.ByteIs(p1) {
		return nil, p, false
	}
	if !hex.ByteIs(p2) {
		return nil, p, false
	}

	return p[:3], p[3:], true
}
