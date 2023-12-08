package mark

func RuneIs(value rune) bool {
	switch value {
	case '-' , '_' , '.' , '!' , '~' , '*' , '\'' , '(' , ')':
		return true
	default:
		return false
	}
}
