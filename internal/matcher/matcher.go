package matcher

import (
	"fmt"
	"reflect"

	. "github.com/petergtz/pegomock/internal/types"
)

type EqMatcher struct {
	Value  Param
	actual Param
}

func (matcher *EqMatcher) Matches(param Param) bool {
	matcher.actual = param
	return matcher.Value == param
}

func (matcher *EqMatcher) FailureMessage() string {
	return fmt.Sprintf("Expected: %v; but got: %v", matcher.Value, matcher.actual)
}

type AnyMatcher struct {
	Type   reflect.Kind
	actual reflect.Kind
}

func (matcher *AnyMatcher) Matches(param Param) bool {
	matcher.actual = reflect.TypeOf(param).Kind()
	return reflect.TypeOf(param).Kind() == matcher.Type
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
func (matcher *AtMostIntMatcher) FailureMessage() string {
	return fmt.Sprintf("Expected: at most %v; but got: %v", matcher.Value, matcher.actual)
}
