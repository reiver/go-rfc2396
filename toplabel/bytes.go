package toplabel

import (
	"sourcecode.social/reiver/go-rfc2396/alpha"
	"sourcecode.social/reiver/go-rfc2396/alphanum"
)

func Bytes(p []byte) (result []byte, rest []byte, ok bool) {

	length := len(p)

	if length < 1 {
		return nil, p, false
	}

	{
		p0 := p[0]

		if !alpha.ByteIs(p0) {
			return nil, p, false
		}
	}

	{
		for i,b := range p {
			if !alphanum.ByteIs(b) && '-' != b {

				if !alphanum.ByteIs(p[i-1]) {
					return p[:i-1], p[i-1:], true
				}

				return p[:i], p[i:], true
			}
		}


		if !alphanum.ByteIs(p[length-1]) {
			return p[:length-1], p[length-1:], true
		}


		return p, nil, true
	}
}
