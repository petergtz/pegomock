package pegomock

import (
	"reflect"

	"github.com/davecgh/go-spew/spew"
	"github.com/onsi/gomega/types"
	"github.com/petergtz/goextract/util"
)

func IntTo(gomegaMatcher types.GomegaMatcher) int {
	RegisterMatcher(NewGomegaMatcherAdapter(gomegaMatcher))
	return 0
}

func StringTo(gomegaMatcher types.GomegaMatcher) string {
	RegisterMatcher(NewGomegaMatcherAdapter(gomegaMatcher))
	return ""
}

type GomegaMatcherAdapter struct {
	gomegaMatcher types.GomegaMatcher
	actual        interface{}
}

func NewGomegaMatcherAdapter(gomegaMatcher types.GomegaMatcher) *GomegaMatcherAdapter {
	return &GomegaMatcherAdapter{gomegaMatcher: gomegaMatcher}
}

func (matcher *GomegaMatcherAdapter) Matches(param Param) bool {
	matcher.actual = param
	matches, e := matcher.gomegaMatcher.Match(param)
	util.PanicOnError(e)
	return matches
}

func (matcher *GomegaMatcherAdapter) Equals(other interface{}) bool {
	if other == nil {
		return false
	}
	otherMatcher, ok := other.(*GomegaMatcherAdapter)
	if !ok {
		return false
	}
	return reflect.DeepEqual(otherMatcher.gomegaMatcher, matcher.gomegaMatcher)
}

func (matcher *GomegaMatcherAdapter) FailureMessage() string {
	return matcher.gomegaMatcher.FailureMessage(matcher.actual)
}

func (matcher *GomegaMatcherAdapter) String() string {
	return spew.Sprintf("ArgThatMatches(%#v)", matcher.gomegaMatcher)
}
