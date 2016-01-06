package matcher

import "reflect"
import . "github.com/petergtz/pegomock/internal/types"

type EqMatcher struct {
	Value Param
}

func (matcher *EqMatcher) Matches(param Param) bool {
	return matcher.Value == param
}

type AtLeastIntMatcher struct {
	Value int
}

func (matcher *AtLeastIntMatcher) Matches(param Param) bool {
	return param.(int) >= matcher.Value
}

type AtMostIntMatcher struct {
	Value int
}

func (matcher *AtMostIntMatcher) Matches(param Param) bool {
	return param.(int) <= matcher.Value
}

type AnyMatcher struct{}

func (matcher *AnyMatcher) Matches(param Param) bool {
	return true
}

type AnyIntMatcher struct{}

func (matcher *AnyIntMatcher) Matches(param Param) bool {
	return reflect.TypeOf(param).Kind() == reflect.Int
}

type AnyStringMatcher struct{}

func (matcher *AnyStringMatcher) Matches(param Param) bool {
	return reflect.TypeOf(param).Kind() == reflect.String
}

type AnyFloat32Matcher struct{}
