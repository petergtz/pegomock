package mock

import "github.com/petergtz/pegomock/v3"

func Eq[T any](value T) T    { return pegomock.Eq(value) }
func NotEq[T any](value T) T { return pegomock.NotEq(value) }
func Any[T any]() T          { return pegomock.Any[T]() }

func ArgThat[T any](matcher pegomock.ArgumentMatcher) T { return pegomock.ArgThat[T](matcher) }
