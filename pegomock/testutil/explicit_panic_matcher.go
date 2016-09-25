package testutil

import (
	"fmt"
	"reflect"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

type ExplicitPanicMatcher struct {
	object interface{}
}

func PanicButReport() types.GomegaMatcher {
	return &ExplicitPanicMatcher{}
}

func (matcher *ExplicitPanicMatcher) Match(actual interface{}) (success bool, err error) {
	if actual == nil {
		return false, fmt.Errorf("PegomockPanicMatcher expects a non-nil actual.")
	}

	actualType := reflect.TypeOf(actual)
	if actualType.Kind() != reflect.Func {
		return false, fmt.Errorf("PegomockPanicMatcher expects a function.  Got:\n%s", format.Object(actual, 1))
	}
	if !(actualType.NumIn() == 0 && actualType.NumOut() == 0) {
		return false, fmt.Errorf("PegomockPanicMatcher expects a function with no arguments and no return value.  Got:\n%s", format.Object(actual, 1))
	}

	success = false
	defer func() {
		if e := recover(); e != nil {
			matcher.object = e
			success = true
		}
	}()

	reflect.ValueOf(actual).Call([]reflect.Value{})

	return
}

func (matcher *ExplicitPanicMatcher) FailureMessage(actual interface{}) (message string) {
	return format.Message(actual, "to panic, but didn't.")
}

func (matcher *ExplicitPanicMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return format.Message(actual, fmt.Sprintf("not to panic, but panicked with <%T>: %v", matcher.object, matcher.object))
}
