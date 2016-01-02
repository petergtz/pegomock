package pegomock

// EqInt .
func EqInt(value int) int {
	argMatchers.append(&EqMatcher{value})
	return value
}

// EqString .
func EqString(value string) string {
	argMatchers.append(&EqMatcher{value})
	return value
}

// AnyString .
func AnyString() string {
	argMatchers.append(&AnyStringMatcher{})
	return ""
}

// AnyInt .
func AnyInt() int {
	argMatchers.append(&AnyIntMatcher{})
	return 0
}
