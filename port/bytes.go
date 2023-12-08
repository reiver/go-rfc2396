package port

import (
	"sourcecode.social/reiver/go-rfc2396/digit"
)

func Bytes(p []byte) ([]byte, []byte, bool) {
	length := len(p)

	if length <= 0 {
		return nil, nil, true
	}

	{
		p0 := p[0]

		if !digit.ByteIs(p0) {
			return nil, p, true
		}
	}

	for i,b := range p {
		if !digit.ByteIs(b) {
			return p[:i], p[i:], true
		}
	}

	return p, nil, true
}
