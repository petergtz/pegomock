// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"github.com/petergtz/pegomock"
	"reflect"

	io "io"
)

func AnyIoReadCloser() io.ReadCloser {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(io.ReadCloser))(nil)).Elem()))
	var nullValue io.ReadCloser
	return nullValue
}

func EqIoReadCloser(value io.ReadCloser) io.ReadCloser {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue io.ReadCloser
	return nullValue
}

func NotEqIoReadCloser(value io.ReadCloser) io.ReadCloser {
	pegomock.RegisterMatcher(&pegomock.NotEqMatcher{Value: value})
	var nullValue io.ReadCloser
	return nullValue
}

func IoReadCloserThat(matcher pegomock.Matcher) io.ReadCloser {
	pegomock.RegisterMatcher(matcher)
	var nullValue io.ReadCloser
	return nullValue
}
