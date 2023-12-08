package mark

func ByteIs(value byte) bool {
	switch value {
	case '-' , '_' , '.' , '!' , '~' , '*' , '\'' , '(' , ')':
		return true
	default:
		return false
	}
}
