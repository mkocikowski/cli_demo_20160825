package replace

// text strings are UTF-8 encoded sequences of Unicode code points (runes)
// len() returns the number of bytes (not runes!) in a string
// the i-th byte of a string is not necessarily the i-th character
func Replace1(s string) string {
	return "X" + s[1:]
}

// rune is a synonym for int32 ; unicode code points are 4 byte
// []rune(s) converts s to a slice of runes
func Replace2(s string) string {
	return "X" + string([]rune(s)[1:])
}

func Replace3(s string) string {
	var i int
	for i, _ = range s {
		if i > 0 {
			break
		}
	}
	return "X" + s[i:]
}
