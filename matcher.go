package pegomock

import (
	"fmt"
	"reflect"

	"github.com/petergtz/pegomock/internal/verify"
)

type EqMatcher struct {
	Value  Param
	actual Param
}

func (matcher *EqMatcher) Matches(param Param) bool {
	matcher.actual = param
	return matcher.Value == param
}

func (matcher *EqMatcher) Equals(other interface{}) bool {
	if other == nil {
		return false
	}
	otherMatcher, ok := other.(*EqMatcher)
	if !ok {
		return false
	}
	return otherMatcher.Value == matcher.Value
}

func (matcher *EqMatcher) FailureMessage() string {
	return fmt.Sprintf("Expected: %v; but got: %v", matcher.Value, matcher.actual)
}

type AnyMatcher struct {
	Type   reflect.Type
	actual reflect.Type
}

func NewAnyMatcher(typ reflect.Type) *AnyMatcher {
	verify.NotNil(typ, "Must provide a non-nil type")
	return &AnyMatcher{Type: typ}
}

func (matcher *AnyMatcher) Matches(param Param) bool {
	matcher.actual = reflect.TypeOf(param)
	if matcher.actual == nil {
		switch matcher.Type.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
			return true
		default:
			return false
		}
	}
	return matcher.actual.AssignableTo(matcher.Type)
}

func (matcher *AnyMatcher) Equals(other interface{}) bool {
	if other == nil {
		return false
	}
	otherMatcher, ok := other.(*AnyMatcher)
	if !ok {
		return false
	}
	return otherMatcher.Type == matcher.Type
}

func (matcher *AnyMatcher) FailureMessage() string {
	return fmt.Sprintf("Expected: %v; but got: %v", matcher.Type, matcher.actual)
}

type AtLeastIntMatcher struct {
	Value  int
	actual int
}

func (matcher *AtLeastIntMatcher) Matches(param Param) bool {
	matcher.actual = param.(int)
	return param.(int) >= matcher.Value
}

func (matcher *AtLeastIntMatcher) Equals(other interface{}) bool {
	if other == nil {
		return false
	}
	otherMatcher, ok := other.(*AtLeastIntMatcher)
	if !ok {
		return false
	}
	return otherMatcher.Value == matcher.Value
}

func (matcher *AtLeastIntMatcher) FailureMessage() string {
	return fmt.Sprintf("Expected: %v; but got: %v", matcher.Value, matcher.actual)
}

type AtMostIntMatcher struct {
	Value  int
	actual int
}

func (matcher *AtMostIntMatcher) Matches(param Param) bool {
	matcher.actual = param.(int)
	return param.(int) <= matcher.Value
}

func (matcher *AtMostIntMatcher) Equals(other interface{}) bool {
	if other == nil {
		return false
	}
	otherMatcher, ok := other.(*AtMostIntMatcher)
	if !ok {
		return false
	}
	return otherMatcher.Value == matcher.Value
}

func (matcher *AtMostIntMatcher) FailureMessage() string {
	return fmt.Sprintf("Expected: at most %v; but got: %v", matcher.Value, matcher.actual)
}
