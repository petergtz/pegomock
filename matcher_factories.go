package pegomock

import (
	"reflect"

	"github.com/petergtz/pegomock/internal/matcher"
)

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
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.String})
	return ""
}

func AnyStringSlice() []string {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Slice})
	return nil
}

// AnyInt .
func AnyInt() int {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Int})
	return 0
}

// AnyFloat32 .
func AnyFloat32() float32 {
	RegisterMatcher(&matcher.AnyMatcher{Type: reflect.Float32})
	return 0
}
