/*
Functions that replace the first character of a string with "X".

Text strings are UTF-8 encoded sequences of Unicode code points (runes). len()
returns the number of bytes (not runes!) in a string. The i-th byte of a string
is not necessarily the i-th character! String values are immutable, and so to
change first character can't simply do s[0] = "X" (same as in Python).
*/
package replace

// Byte slice, will not work with multi-byte characters, see tests
func Replace1(s string) string {
	// will not work with multi byte characters; see test
	// in Python, this would work:
	// >>> "X" + u'\u4e16xx'[1:]
	// u'Xxx'
	return "X" + s[1:]
}

// Convert to []rune
func Replace2(s string) string {
	// rune is a synonym for int32 ; unicode code points are 4 byte
	// []rune(s) converts s to a slice of runes
	return "X" + string([]rune(s)[1:])
}

// Use range's implicit decoding of UTF-8
func Replace3(s string) string {
	// range loop, when applied to a string, decodes UTF-8
	// here i is byte offset of each rubsequent rune (so if there is a 3-byte
	// rune, i will advance by 3 in a single loop)
	var i int
	for i, _ = range s {
		if i > 0 {
			break
		}
	}
	return "X" + s[i:]
}
