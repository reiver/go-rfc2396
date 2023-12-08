package regname

import (
	"sourcecode.social/reiver/go-rfc2396/escaped"
	"sourcecode.social/reiver/go-rfc2396/unreserved"
)

// Bytes returns the 'reg_name', as defined by IETF RFC-2396.
//
//	reg_name = 1*( unreserved | escaped | "$" | "," | ";" | ":" | "@" | "&" | "=" | "+" )
func Bytes(p []byte) (regName []byte, rest []byte, ok bool) {

	var length int = len(p)

	if length <= 0  {
		return nil, p, false
	}

	// Make sure there is at least one valid "character".
	{
		p0 := p[0]

		switch p0 {
		case '%':
			_, _, ok := escaped.Bytes(p)
			if !ok {
				return nil, p, false
			}

			// Nothingh (else) here.
		case '$' , ',' , ';' , ':' , '@' , '&' , '=' , '+':
			// Nothing here.
		default:
			switch {
			case unreserved.ByteIs(p0):
				// Nothing here.
			default:
				return nil, p, false
			}
		}
	}

	{
		var skip uint

		for i,b := range p {
			if 0 < skip {
				skip--
				continue
			}

			switch b {
			case '%':
				_, _, ok := escaped.Bytes(p[i:])
				if !ok {
					return p[:i], p[i:], true
				}

				skip = 2

				// Nothingh (else) here.
			case '$' , ',' , ';' , ':' , '@' , '&' , '=' , '+':
				// Nothing here.
			default:

				switch {
				case unreserved.ByteIs(b):
					// Nothing here.
				default:
					return p[:i], p[i:], true
				}
			}
		}

		return p, nil, true
	}
}
