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

type AnyMatcher struct{}

func (matcher *AnyMatcher) Matches(param Param) bool { return true }
func (matcher *AnyMatcher) FailureMessage() string   { return "Unused" }

type AnyIntMatcher struct{}

func (matcher *AnyIntMatcher) Matches(param Param) bool {
	return reflect.TypeOf(param).Kind() == reflect.Int
}
func (matcher *AnyIntMatcher) FailureMessage() string { return "Unused" }

type AnyStringMatcher struct{}

func (matcher *AnyStringMatcher) Matches(param Param) bool {
	return reflect.TypeOf(param).Kind() == reflect.String
}
func (matcher *AnyStringMatcher) FailureMessage() string { return "Unused" }

type AnyFloat32Matcher struct{}
