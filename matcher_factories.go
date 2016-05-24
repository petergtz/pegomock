package pegomock

import "github.com/petergtz/pegomock/internal/matcher"

// EqInt .
func EqInt(value int) int {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return 0
}

// EqString .
func EqString(value string) string {
	RegisterMatcher(&matcher.EqMatcher{Value: value})
	return ""
}

// AnyString .
func AnyString() string {
	RegisterMatcher(&matcher.AnyStringMatcher{})
	return ""
}

// AnyInt .
func AnyInt() int {
	RegisterMatcher(&matcher.AnyIntMatcher{})
	return 0
}
