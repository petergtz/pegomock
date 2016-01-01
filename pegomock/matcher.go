package pegomock

import "reflect"

type Matcher interface {
	matches(param Param) bool
}

type EqMatcher struct {
	value Param
}

func (matcher *EqMatcher) matches(param Param) bool {
	return matcher.value == param
}

type AnyMatcher struct{}

func (matcher *AnyMatcher) matches(param Param) bool {
	return true
}

type AnyIntMatcher struct{}

func (matcher *AnyIntMatcher) matches(param Param) bool {
	return reflect.TypeOf(param).Kind() == reflect.Int
}

type AnyStringMatcher struct{}

func (matcher *AnyStringMatcher) matches(param Param) bool {
	return reflect.TypeOf(param).Kind() == reflect.String
}

type AnyFloat32Matcher struct{}
