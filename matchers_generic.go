package pegomock

import "reflect"

func Eq[T any](value T) T {
	RegisterMatcher(&EqMatcher{Value: value})
	var t T
	return t
}

func NotEq[T any](value T) T {
	RegisterMatcher(&NotEqMatcher{Value: value})
	var t T
	return t
}

func Any[T any]() T {
	var t T
	RegisterMatcher(NewAnyMatcher(reflect.TypeOf(&t).Elem()))
	return t
}

func ArgThat[T any](matcher ArgumentMatcher) T {
	RegisterMatcher(matcher)
	var t T
	return t
}
