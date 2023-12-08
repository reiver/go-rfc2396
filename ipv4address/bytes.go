package ipv4address

import (
	"sourcecode.social/reiver/go-rfc2396/decnum"
)

func Bytes(p []byte) (result []byte, rest []byte, ok bool) {

	length := len(p)

	if length < 7 {
		return nil, p, false
	}

	var pp []byte = p
	var sum int

	// ex: 192
	{
		result, rest, ok := decnum.Bytes(pp)
		if !ok {
			return nil, p, false
		}

		sum += len(result)
		pp = rest

		if len(pp) < 6 {
			return nil, p, false
		}
	}

	// ex: 192.
	{
		b := pp[0]
		if '.' != b {
			return nil, p, false
		}

		sum += 1
		pp = pp[1:]
	}

	// ex: 192.168
	{
		result, rest, ok := decnum.Bytes(pp)
		if !ok {
			return nil, p, false
		}

		sum += len(result)
		pp = rest

		if len(pp) < 4 {
			return nil, p, false
		}
	}

	// ex: 192.168.
	{
		b := pp[0]
		if '.' != b {
			return nil, p, false
		}

		sum += 1
		pp = pp[1:]
	}

	// ex: 192.168.0
	{
		result, rest, ok := decnum.Bytes(pp)
		if !ok {
			return nil, p, false
		}

		sum += len(result)
		pp = rest

		if len(pp) < 2 {
			return nil, p, false
		}
	}

	// ex: 192.168.0.
	{
		b := pp[0]
		if '.' != b {
			return nil, p, false
		}

		sum += 1
		pp = pp[1:]
	}

	// ex: 192.168.0.1
	{
		result, rest, ok := decnum.Bytes(pp)
		if !ok {
			return nil, p, false
		}

		sum += len(result)
		pp = rest
	}

	if length < sum {
		return nil, p, false
	}
	return p[:sum], p[sum:], true
}
