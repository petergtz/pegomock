// Copyright 2015 Peter Goetz
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pegomock_test

import (
	"errors"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/petergtz/pegomock/v4"
	. "github.com/petergtz/pegomock/v4"
	"github.com/petergtz/pegomock/v4/test_interface"
	"github.com/samber/lo"
)

var (
	BeforeEach       = ginkgo.BeforeEach
	It               = ginkgo.It
	FIt              = ginkgo.FIt
	Describe         = ginkgo.Describe
	Context          = ginkgo.Context
	BeNil            = gomega.BeNil
	BeTrue           = gomega.BeTrue
	ConsistOf        = gomega.ConsistOf
	ContainSubstring = gomega.ContainSubstring
	MatchError       = gomega.MatchError
	Equal            = gomega.Equal
	Expect           = gomega.Expect
	HaveLen          = gomega.HaveLen
	HavePrefix       = gomega.HavePrefix
	Panic            = gomega.Panic
	SatisfyAll       = gomega.SatisfyAll
)

var _ test_interface.Display = NewMockDisplay() //check that interface is implemented

func TestDSL(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	pegomock.RegisterMockFailHandler(func(message string, callerSkip ...int) { panic(message) })
	ginkgo.RunSpecs(t, "DSL Suite")
}

type NeverMatcher struct{}

func (matcher *NeverMatcher) Matches(Param) bool { return false }
func (matcher *NeverMatcher) FailureMessage() string {
	return "This matcher never matches (and is only for testing purposes)"
}
func (matcher *NeverMatcher) String() string { return "NeverMatching" }

func NeverMatchingRequest() http.Request {
	RegisterMatcher(&NeverMatcher{})
	return http.Request{}
}

var _ = Describe("MockDisplay", func() {
	var display *MockDisplay

	BeforeEach(func() {
		display = NewMockDisplay()
	})

	Context("Calling SomeValue() with no stubbing", func() {
		It("returns zero value", func() {
			Expect(display.SomeValue()).To(Equal(""))
		})
	})

	Context("Stubbing MultipleParamsAndReturnValue() with matchers", func() {
		BeforeEach(func() {
			When(display.MultipleParamsAndReturnValue(Eq("Hello"), Eq(333))).ThenReturn("Bla")
		})

		It("fails during verification when mock was not called", func() {
			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello", 333) }).To(PanicWithMessageTo(HavePrefix(
				"Mock invocation count for MultipleParamsAndReturnValue(\"Hello\", 333) does not match expectation.\n\n\tExpected: 1; but got: 0",
			)))
		})

		It("succeeds verification when mock was called", func() {
			display.MultipleParamsAndReturnValue("Hello", 333)
			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello", 333) }).NotTo(Panic())
		})

		It("succeeds verification when verification and invocation are mixed", func() {
			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello", 333) }).To(PanicWithMessageTo(HavePrefix(
				expectation{method: "MultipleParamsAndReturnValue(\"Hello\", 333)", expected: "1", actual: "0"}.string(),
			)))
			display.MultipleParamsAndReturnValue("Hello", 333)
			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello", 333) }).NotTo(Panic())
		})
	})

	Context("Calling MultipleParamsAndReturnValue() with \"Any\"-matchers", func() {
		It("succeeds all verifications that match", func() {
			When(display.MultipleParamsAndReturnValue(Any[string](), Eq(333))).ThenReturn("Bla")

			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello", 333) }).To(PanicWithMessageTo(HavePrefix(
				expectation{method: "MultipleParamsAndReturnValue(\"Hello\", 333)", expected: "1", actual: "0"}.string(),
			)))

			display.MultipleParamsAndReturnValue("Hello", 333)
			display.MultipleParamsAndReturnValue("Hello again", 333)
			display.MultipleParamsAndReturnValue("And again", 333)

			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello", 333) }).NotTo(Panic())
			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello again", 333) }).NotTo(Panic())
			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("And again", 333) }).NotTo(Panic())

			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("And again", 444) }).To(PanicWithMessageTo(HavePrefix(
				expectation{method: "MultipleParamsAndReturnValue(\"And again\", 444)", expected: "1", actual: "0"}.string(),
			)))

		})
	})

	Context("Calling MultipleParamsAndReturnValue() only with matchers on some parameters", func() {
		It("panics", func() {
			Expect(func() { When(display.MultipleParamsAndReturnValue(Eq("Hello"), 333)) }).To(PanicWithMessageTo(HavePrefix(
				"Invalid use of matchers!\n\n 2 matchers expected, 1 recorded.\n\n" +
					"This error may occur if matchers are combined with raw values:\n" +
					"    //incorrect:\n" +
					"    someFunc(Any[int](), \"raw String\")\n" +
					"When using matchers, all arguments have to be provided by matchers.\n" +
					"For example:\n" +
					"    //correct:\n" +
					"    someFunc(Any[int](), Eq(\"String by matcher\"))",
			)))
		})
	})

	Context("Stubbing with consecutive return values", func() {
		BeforeEach(func() {
			When(display.SomeValue()).ThenReturn("Hello").ThenReturn("again")
		})

		It("returns stubbed values when calling mock", func() {
			Expect(display.SomeValue()).To(Equal("Hello"))
			Expect(display.SomeValue()).To(Equal("again"))
		})

		It("returns last stubbed value repeatedly", func() {
			Expect(display.SomeValue()).To(Equal("Hello"))
			Expect(display.SomeValue()).To(Equal("again"))
			Expect(display.SomeValue()).To(Equal("again"))
			Expect(display.SomeValue()).To(Equal("again"))
			Expect(display.SomeValue()).To(Equal("again"))
			Expect(display.SomeValue()).To(Equal("again"))
		})

		It("can be verified that mock was called", func() {
			display.SomeValue()
			Expect(func() { display.VerifyWasCalledOnce().SomeValue() }).NotTo(Panic())
		})

		It("fails if verify is called on mock that was not invoked.", func() {
			Expect(func() { display.VerifyWasCalledOnce().Show("Some parameter") }).To(PanicWithMessageTo(HavePrefix(
				expectation{method: "Show(\"Some parameter\")", expected: "1", actual: "0"}.string(),
			)))
		})

		It("fails if verify is called on mock that was invoked more than once.", func() {
			display.Show("param")
			display.Show("param")
			Expect(func() { display.VerifyWasCalledOnce().Show("param") }).To(PanicWithMessageTo(HavePrefix(
				expectation{method: "Show(\"param\")", expected: "1", actual: "2"}.string(),
			)))

		})
	})

	Context("Stubbing with invalid return type", func() {
		It("panics", func() {
			Expect(func() { When(display.SomeValue()).ThenReturn("Hello").ThenReturn(0) }).To(PanicWithMessageTo(HavePrefix(
				"Return value of type int not assignable to return type string",
			)))
		})
	})

	Describe("https://github.com/petergtz/pegomock/v4/issues/24", func() {
		Context("Stubbing with nil value", func() {
			It("does not panic when return type is interface{}", func() {
				When(display.InterfaceReturnValue()).ThenReturn(nil)
				Expect(display.InterfaceReturnValue()).To(BeNil())
			})

			It("does not panic when return type is error interface", func() {
				When(display.ErrorReturnValue()).ThenReturn(nil)
				Expect(display.ErrorReturnValue()).To(BeNil())
			})
		})

		Context("Stubbing with value that implements interface{}", func() {
			It("does not panic", func() {
				When(display.InterfaceReturnValue()).ThenReturn("Hello")
				Expect(display.InterfaceReturnValue()).To(Equal("Hello"))
			})
		})

		Context("Stubbing with value that implements error interface", func() {
			It("does not panic", func() {
				When(display.ErrorReturnValue()).ThenReturn(errors.New("ouch"))
				Expect(display.ErrorReturnValue()).To(Equal(errors.New("ouch")))
			})
		})

		Context("Stubbing with value that does not implement error interface", func() {
			It("panics", func() {
				Expect(func() { When(display.ErrorReturnValue()).ThenReturn("Blub") }).To(PanicWithMessageTo(HavePrefix(
					"Return value of type string not assignable to return type error",
				)))
			})
		})

		Context("Stubbing string return type with nil value", func() {
			It("panics", func() {
				Expect(func() { When(display.SomeValue()).ThenReturn(nil) }).To(PanicWith(
					"Return value 'nil' not assignable to return type string",
				))
			})
		})

	})

	Context("Stubbed method, but no invocation takes place", func() {
		It("fails during verification", func() {
			When(display.SomeValue()).ThenReturn("Hello")
			Expect(func() { display.VerifyWasCalledOnce().SomeValue() }).To(PanicWithMessageTo(HavePrefix(
				expectation{method: "SomeValue()", expected: "1", actual: "0"}.string(),
			)))
		})
	})

	Context("Calling Flash() with specific arguments", func() {

		BeforeEach(func() { display.Flash("Hello", 333) })

		It("succeeds verification if values are matching", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash("Hello", 333) }).NotTo(Panic())
		})

		It("fails during verification if values are not matching", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash("Hello", 666) }).To(PanicWithMessageTo(HavePrefix(
				expectation{method: "Flash(\"Hello\", 666)", expected: "1", actual: "0"}.string(),
			)))
		})

		It("succeeds during verification when using Any-matchers ", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash(Any[string](), Any[int]()) }).NotTo(Panic())
		})

		It("succeeds during verification when using valid Eq-matchers ", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash(Eq("Hello"), Eq(333)) }).NotTo(Panic())
		})

		It("fails during verification when using invalid Eq-matchers ", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash(Eq("Invalid"), Eq(-1)) }).To(PanicWithMessageTo(HavePrefix(
				expectation{method: `Flash(Eq("Invalid"), Eq(-1))`, expected: "1", actual: "0"}.string(),
			)))
		})

		It("fails when not using matchers for all params", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash("Hello", Any[int]()) }).To(PanicWith(
				"Invalid use of matchers!\n\n 2 matchers expected, 1 recorded.\n\n" +
					"This error may occur if matchers are combined with raw values:\n" +
					"    //incorrect:\n" +
					"    someFunc(Any[int](), \"raw String\")\n" +
					"When using matchers, all arguments have to be provided by matchers.\n" +
					"For example:\n" +
					"    //correct:\n" +
					"    someFunc(Any[int](), Eq(\"String by matcher\"))",
			))
		})
	})

	Describe("Invocation count matching", func() {

		Context("Calling Flash() twice", func() {

			BeforeEach(func() {
				display.Flash("Hello", 333)
				display.Flash("Hello", 333)
			})

			It("succeeds verification if verifying with Times(2)", func() {
				Expect(func() { display.VerifyWasCalled(Times(2)).Flash("Hello", 333) }).NotTo(Panic())
			})

			It("fails during verification if verifying with VerifyWasCalledOnce", func() {
				Expect(func() { display.VerifyWasCalledOnce().Flash("Hello", 333) }).To(PanicWithMessageTo(HavePrefix(
					expectation{method: "Flash(\"Hello\", 333)", expected: "1", actual: "2"}.string(),
				)))
			})

			It("fails during verification if verifying with Times(1)", func() {
				Expect(func() { display.VerifyWasCalled(Times(1)).Flash("Hello", 333) }).To(PanicWithMessageTo(HavePrefix(
					expectation{method: "Flash(\"Hello\", 333)", expected: "1", actual: "2"}.string(),
				)))
			})

			It("succeeds during verification when using AtLeast(1)", func() {
				Expect(func() { display.VerifyWasCalled(AtLeast(1)).Flash("Hello", 333) }).NotTo(Panic())
			})

			It("succeeds during verification when using AtLeast(2)", func() {
				Expect(func() { display.VerifyWasCalled(AtLeast(2)).Flash("Hello", 333) }).NotTo(Panic())
			})

			It("fails during verification when using AtLeast(3)", func() {
				Expect(func() { display.VerifyWasCalled(AtLeast(3)).Flash("Hello", 333) }).To(PanicWithMessageTo(HavePrefix(
					expectation{method: "Flash(\"Hello\", 333)", expected: "at least 3", actual: "2"}.string(),
				)))
			})

			It("succeeds during verification when using Never()", func() {
				Expect(func() { display.VerifyWasCalled(Never()).Flash("Other value", 333) }).NotTo(Panic())
			})

			It("fails during verification when using Never()", func() {
				Expect(func() { display.VerifyWasCalled(Never()).Flash("Hello", 333) }).To(PanicWithMessageTo(HavePrefix(
					expectation{method: "Flash(\"Hello\", 333)", expected: "0", actual: "2"}.string(),
				)))
			})
		})

		Context("Never calling Flash", func() {
			It("succeeds during verification when using Never() and argument matchers", func() {
				// https://github.com/petergtz/pegomock/v4/issues/34
				Expect(func() { display.VerifyWasCalled(Never()).Flash(Any[string](), Any[int]()) }).NotTo(Panic())
			})
		})
	})

	Context("Calling MultipleParamsAndReturnValue()", func() {

		It("panics when stubbed to panic", func() {
			When(display.MultipleParamsAndReturnValue(Any[string](), Any[int]())).
				ThenPanic("I'm panicking")
			Expect(func() {
				display.MultipleParamsAndReturnValue("Some string", 123)
			}).To(PanicWith("I'm panicking"))
		})

		It("calls back when stubbed to call back", func() {
			When(display.MultipleParamsAndReturnValue(Any[string](), Any[int]())).Then(
				func(params []Param) ReturnValues {
					return []ReturnValue{fmt.Sprintf("%v%v", params[0], params[1])}
				},
			)
			Expect(display.MultipleParamsAndReturnValue("string and ", 123)).
				To(Equal("string and 123"))
		})

	})

	Context("Making calls in a specific order", func() {

		BeforeEach(func() {
			display.Flash("Hello", 111)
			display.Flash("again", 222)
			display.Flash("and again", 333)
		})

		It("succeeds during InOrder verification when order is correct", func() {
			Expect(func() {
				inOrderContext := new(InOrderContext)
				display.VerifyWasCalledInOrder(Once(), inOrderContext).Flash("Hello", 111)
				display.VerifyWasCalledInOrder(Once(), inOrderContext).Flash("again", 222)
				display.VerifyWasCalledInOrder(Once(), inOrderContext).Flash("and again", 333)
			}).NotTo(Panic())
		})

		It("succeeds during InOrder verification when order is correct, but not all invocations are verified", func() {
			Expect(func() {
				inOrder := new(InOrderContext)
				display.VerifyWasCalledInOrder(Once(), inOrder).Flash("Hello", 111)
				// not checking for the 2nd call here
				display.VerifyWasCalledInOrder(Once(), inOrder).Flash("and again", 333)
			}).NotTo(Panic())
		})

		It("fails during InOrder verification when order is not correct", func() {
			Expect(func() {
				inOrder := new(InOrderContext)
				display.VerifyWasCalledInOrder(Once(), inOrder).Flash("again", 222)
				display.VerifyWasCalledInOrder(Once(), inOrder).Flash("Hello", 111)
				display.VerifyWasCalledInOrder(Once(), inOrder).Flash("and again", 333)
			}).To(PanicWithMessageTo(HavePrefix(
				"Expected function call Flash(\"Hello\", 111) before function call Flash(\"again\", 222)",
			)))
		})

	})

	Context("Capturing arguments", func() {
		It("Returns arguments when verifying with argument capture", func() {
			display.Flash("Hello", 111)

			arg1, arg2 := display.VerifyWasCalledOnce().Flash(Any[string](), Any[int]()).GetCapturedArguments()

			Expect(arg1).To(Equal("Hello"))
			Expect(arg2).To(Equal(111))
		})

		It("Returns arguments of last invocation when verifying with argument capture", func() {
			display.Flash("Hello", 111)
			display.Flash("Again", 222)

			arg1, arg2 := display.VerifyWasCalled(AtLeast(1)).Flash(Any[string](), Any[int]()).GetCapturedArguments()

			Expect(arg1).To(Equal("Again"))
			Expect(arg2).To(Equal(222))
		})

		It("Returns arguments of all invocations when verifying with \"all\" argument capture", func() {
			display.Flash("Hello", 111)
			display.Flash("Again", 222)

			args1, args2 := display.VerifyWasCalled(AtLeast(1)).Flash(Any[string](), Any[int]()).GetAllCapturedArguments()

			Expect(args1).To(ConsistOf("Hello", "Again"))
			Expect(args2).To(ConsistOf(111, 222))
		})

		It("Returns *array* arguments of all invocations when verifying with \"all\" argument capture", func() {
			display.ArrayParam([]string{"one", "two"})
			display.ArrayParam([]string{"4", "5", "3"})

			args := display.VerifyWasCalled(AtLeast(1)).ArrayParam(Any[[]string]()).GetAllCapturedArguments()

			Expect(lo.Flatten(args)).To(ConsistOf("one", "two", "3", "4", "5"))
		})

	})

	Context("Stubbing using string slice", func() {
		It("does not panic when comparing the slices in the matcher", func() {
			When(func() { display.ArrayParam([]string{"one", "two"}) }).Then(func([]Param) ReturnValues {
				// do nothing, because that's not our focus here.
				return nil
			})
			display.ArrayParam([]string{"one", "two"})
		})
	})

	Describe("Different \"Any\" matcher scenarios", func() {
		It("Succeeds when int-parameter is passed as int but veryfied as float", func() {
			display.FloatParam(1)
			display.VerifyWasCalledOnce().FloatParam(Any[float32]())
		})

		It("Panics when interface{}-parameter is passed as int, but verified as float", func() {
			Expect(func() {
				display.InterfaceParam(3)
				display.VerifyWasCalledOnce().InterfaceParam(Any[float32]())
			}).To(PanicWithMessageTo(HavePrefix(
				expectation{method: "InterfaceParam(Any(float32))", expected: "1", actual: "0"}.string(),
			)))
		})

		It("Panics when interface{}-parameter is passed as float, but verified as int", func() {
			Expect(func() {
				display.InterfaceParam(3.141)
				display.VerifyWasCalledOnce().InterfaceParam(Any[int]())
			}).To(PanicWithMessageTo(HavePrefix(
				expectation{method: "InterfaceParam(Any(int))", expected: "1", actual: "0"}.string(),
			)))
		})

		It("Succeeds when interface{}-parameter is passed as int and verified as int", func() {
			display.InterfaceParam(3)
			display.VerifyWasCalledOnce().InterfaceParam(Any[int]())
		})

		It("Succeeds when interface{}-parameter is passed as nil and verified as int slice", func() {
			display.InterfaceParam(nil)
			display.VerifyWasCalledOnce().InterfaceParam(Any[[]int]())
		})

		It("Panics when interface{}-parameter is passed as nil, but verified as int", func() {
			Expect(func() {
				display.InterfaceParam(nil)
				display.VerifyWasCalledOnce().InterfaceParam(Any[int]())
			}).To(PanicWithMessageTo(HavePrefix(
				expectation{method: "InterfaceParam(Any(int))", expected: "1", actual: "0"}.string(),
			)))
		})

		It("Succeeds when error-parameter is passed as nil and verified as any error", func() {
			display.ErrorParam(nil)
			display.VerifyWasCalledOnce().ErrorParam(Any[error]())
		})

		It("Succeeds when error-parameter is passed as string error and verified as any error", func() {
			display.ErrorParam(errors.New("some error"))
			display.VerifyWasCalledOnce().ErrorParam(Any[error]())
		})

		It("Succeeds when http.Request-parameter is passed as null value and verified as any http.Request", func() {
			display.NetHttpRequestParam(http.Request{})
			display.VerifyWasCalledOnce().NetHttpRequestParam(Any[http.Request]())
		})

		It("Succeeds when http.Request-parameter is passed as null value to interface{} and verified as any http.Request", func() {
			display.InterfaceParam(http.Request{})
			display.VerifyWasCalledOnce().InterfaceParam(Any[http.Request]())
		})

		It("Fails when *pointer* to http.Request-parameter is passed to interface{} and verified as any http.Request", func() {
			display.InterfaceParam(&http.Request{})
			Expect(func() { display.VerifyWasCalledOnce().InterfaceParam(Any[http.Request]()) }).To(PanicWithMessageTo(SatisfyAll(
				ContainSubstring("InterfaceParam(Any(http.Request))"),
				ContainSubstring("InterfaceParam(&http.Request{Method"),
			)))
		})

		It("Succeeds when http.Request-Pointer-parameter is passed as nil and verified as any *http.Request", func() {
			display.NetHttpRequestPtrParam(nil)
			display.VerifyWasCalledOnce().NetHttpRequestPtrParam(Any[*http.Request]())
		})

		It("Succeeds when http.Request-Pointer-parameter is passed as null value and verified as any *http.Request", func() {
			display.NetHttpRequestPtrParam(&http.Request{})
			display.VerifyWasCalledOnce().NetHttpRequestPtrParam(Any[*http.Request]())
		})
	})

	Describe("Generated matchers", func() {
		It("Succeeds when map-parameter is passed to interface{} and verified as any map", func() {
			display.InterfaceParam(map[string]http.Request{"foo": {}})
			display.VerifyWasCalledOnce().InterfaceParam(Any[map[string]http.Request]())
		})

		It("Fails when string parameter is passed to interface{} and verified as any map", func() {
			display.InterfaceParam("This will not match")
			Expect(func() { display.VerifyWasCalledOnce().InterfaceParam(Any[map[string]http.Request]()) }).To(PanicWithMessageTo(SatisfyAll(
				ContainSubstring("InterfaceParam(Any(map[string]http.Request))"),
				ContainSubstring("InterfaceParam(\"This will not match\")"),
			)))
		})

		It("Succeeds when map-parameter is passed to interface{} and verified as eq map", func() {
			display.InterfaceParam(map[string]http.Request{"foo": {}})
			display.VerifyWasCalledOnce().InterfaceParam(Eq(map[string]http.Request{"foo": {}}))
		})

		It("Fails when map-parameter is passed to interface{} and verified as not eq map", func() {
			display.InterfaceParam(map[int]int{1: 2})
			Expect(func() { display.VerifyWasCalledOnce().InterfaceParam(NotEq(map[int]int{1: 2})) }).
				To(PanicWithMessageTo(ContainSubstring("InterfaceParam(NotEq(map[int]int{1:2}))")))
		})

		It("Succeeds when map-parameter is passed to interface{} and verified as not eq a different map", func() {
			display.InterfaceParam(map[int]int{1: 2})
			display.VerifyWasCalledOnce().InterfaceParam(NotEq(map[int]int{1: 3}))
		})

		It("Succeeds when map-parameter is passed to interface{} and verified as not eq a time", func() {
			display.InterfaceParam(map[int]int{1: 2})
			display.VerifyWasCalledOnce().InterfaceParam(NotEq(time.Now()))
		})

		It("Fails when map-parameter is passed to interface{} and verified as different map that satisfies equals matcher", func() {
			display.InterfaceParam(map[int]int{1: 2})
			Expect(func() {
				display.VerifyWasCalledOnce().InterfaceParam(ArgThat[map[int]int](&EqMatcher{Value: map[int]int{1: 3}}))
			}).
				To(PanicWithMessageTo(SatisfyAll(
					ContainSubstring("InterfaceParam(Eq(map[int]int{1:3}))"),
					ContainSubstring("InterfaceParam(map[int]int{1:2})"),
				)))
		})
		It("Succeeds when map-parameter is passed to interface{} and verified as map that satisfies equals matcher", func() {
			display.InterfaceParam(map[int]int{1: 2})
			display.VerifyWasCalledOnce().InterfaceParam(ArgThat[map[int]int](&EqMatcher{Value: map[int]int{1: 2}}))
		})
	})

	Describe("Logic around matchers and verification", func() {
		// TODO maybe this should go somewhere else
		It("Fails when http.Request-parameter is passed as null value and verified as never matching http.Request", func() {
			display.NetHttpRequestParam(http.Request{})
			Expect(func() { display.VerifyWasCalledOnce().NetHttpRequestParam(NeverMatchingRequest()) }).
				To(PanicWithMessageTo(Equal(`Mock invocation count for NetHttpRequestParam(NeverMatching) does not match expectation.

	Expected: 1; but got: 0

	Actual interactions with this mock were:
	NetHttpRequestParam(http.Request{Method:"", URL:(*url.URL)(nil), Proto:"", ProtoMajor:0, ProtoMinor:0, Header:http.Header(nil), Body:io.ReadCloser(nil), GetBody:(func() (io.ReadCloser, error))(nil), ContentLength:0, TransferEncoding:[]string(nil), Close:false, Host:"", Form:url.Values(nil), PostForm:url.Values(nil), MultipartForm:(*multipart.Form)(nil), Trailer:http.Header(nil), RemoteAddr:"", RequestURI:"", TLS:(*tls.ConnectionState)(nil), Cancel:(<-chan struct {})(nil), Response:(*http.Response)(nil), ctx:context.Context(nil)})
`)))
		})
	})

	Describe("Stubbing with multiple ThenReturns versus multiple stubbings with same parameters", func() {
		Context("One stubbing with multiple ThenReturns", func() {
			It("returns the values in the order of the ThenReturns", func() {
				When(display.MultipleParamsAndReturnValue("one", 1)).ThenReturn("first").ThenReturn("second")

				Expect(display.MultipleParamsAndReturnValue("one", 1)).To(Equal("first"))
				Expect(display.MultipleParamsAndReturnValue("one", 1)).To(Equal("second"))
			})
		})

		Context("Multiple stubbings with same parameters", func() {
			It("overrides previous stubbings with last one", func() {
				When(display.MultipleParamsAndReturnValue("one", 1)).ThenReturn("first")
				When(display.MultipleParamsAndReturnValue("one", 1)).ThenReturn("second")

				Expect(display.MultipleParamsAndReturnValue("one", 1)).To(Equal("second"))
				Expect(display.MultipleParamsAndReturnValue("one", 1)).To(Equal("second"))
			})
		})
	})

	Describe("Verifying gives hints about actual invocations in failure messages", func() {
		It("shows actual interactions with same methods", func() {
			display.Flash("Hello", 123)
			display.Flash("Again", 456)

			Expect(func() { display.VerifyWasCalledOnce().Flash("wrong string", -987) }).To(PanicWith(
				"Mock invocation count for Flash(\"wrong string\", -987) " +
					"does not match expectation.\n\n\tExpected: 1; but got: 0\n\n" +
					"\tActual interactions with this mock were:\n" +
					"\tFlash(\"Hello\", 123)\n" +
					"\tFlash(\"Again\", 456)\n",
			))
		})

		It("shows actual interactions with all methods", func() {
			display.Show("Again")
			display.Flash("Hello", 123)

			Expect(func() { display.VerifyWasCalledOnce().Flash("wrong string", -987) }).To(PanicWith(
				"Mock invocation count for Flash(\"wrong string\", -987) " +
					"does not match expectation.\n\n\tExpected: 1; but got: 0\n\n" +
					"\tActual interactions with this mock were:\n" +
					"\tFlash(\"Hello\", 123)\n" +
					"\tShow(\"Again\")\n"),
			)
		})

		It("formats params in interactions with Go syntax for better readability", func() {
			display.NetHttpRequestParam(http.Request{Host: "x.com"})
			Expect(func() { display.VerifyWasCalledOnce().NetHttpRequestParam(http.Request{Host: "y.com"}) }).To(PanicWith(
				`Mock invocation count for NetHttpRequestParam(http.Request{Method:"", URL:(*url.URL)(nil), Proto:"", ProtoMajor:0, ProtoMinor:0, Header:http.Header(nil), Body:io.ReadCloser(nil), GetBody:(func() (io.ReadCloser, error))(nil), ContentLength:0, TransferEncoding:[]string(nil), Close:false, Host:"y.com", Form:url.Values(nil), PostForm:url.Values(nil), MultipartForm:(*multipart.Form)(nil), Trailer:http.Header(nil), RemoteAddr:"", RequestURI:"", TLS:(*tls.ConnectionState)(nil), Cancel:(<-chan struct {})(nil), Response:(*http.Response)(nil), ctx:context.Context(nil)}) does not match expectation.

	Expected: 1; but got: 0

	Actual interactions with this mock were:
	NetHttpRequestParam(http.Request{Method:"", URL:(*url.URL)(nil), Proto:"", ProtoMajor:0, ProtoMinor:0, Header:http.Header(nil), Body:io.ReadCloser(nil), GetBody:(func() (io.ReadCloser, error))(nil), ContentLength:0, TransferEncoding:[]string(nil), Close:false, Host:"x.com", Form:url.Values(nil), PostForm:url.Values(nil), MultipartForm:(*multipart.Form)(nil), Trailer:http.Header(nil), RemoteAddr:"", RequestURI:"", TLS:(*tls.ConnectionState)(nil), Cancel:(<-chan struct {})(nil), Response:(*http.Response)(nil), ctx:context.Context(nil)})
`,
			))
		})

		It("shows no interactions if there were none", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash("wrong string", -987) }).To(PanicWith(
				"Mock invocation count for Flash(\"wrong string\", -987) " +
					"does not match expectation.\n\n\tExpected: 1; but got: 0\n\n" +
					"\tThere were no other interactions with this mock",
			))
		})
	})

	Describe("Stubbing methods that have no return value", func() {
		It("Can be stubbed with Panic", func() {
			When(func() { display.Show(Any[string]()) }).ThenPanic("bla")
			Expect(func() { display.Show("Hello") }).To(PanicWith("bla"))
		})

		It("Can still work with methods returning a func", func() {
			When(display.FuncReturnValue()).ThenReturn(func() { panic("It's actually a success") })
			Expect(func() { display.FuncReturnValue()() }).To(PanicWith("It's actually a success"))
		})

		It("Panics when not using a func with no params", func() {
			Expect(func() {
				When(func(invalid int) { display.Show(Any[string]()) })
			}).To(PanicWith("When using 'When' with function that does not return a value, it expects a function with no arguments and no return value."))
		})
	})

	Describe("Verifying methods that have variadic arguments", func() {
		Context("One single variadic argument", func() {

			It("succeeds when verifying one invocation with same parameters", func() {
				display.VariadicParam("one", "two")
				display.VerifyWasCalledOnce().VariadicParam("one", "two")
			})

			It("succeeds when verifying two different invocations with same parameters", func() {
				display.VariadicParam("one", "two")
				display.VariadicParam("three", "four", "five")
				display.VerifyWasCalledOnce().VariadicParam("three", "four", "five")
				display.VerifyWasCalledOnce().VariadicParam("one", "two")
			})

			It("succeeds when verifying one invocation with arg matchers", func() {
				display.VariadicParam("one", "two")
				display.VerifyWasCalledOnce().VariadicParam(Any[string](), Any[string]())
			})

			It("succeeds when verifying two different invocations with arg matchers", func() {
				display.VariadicParam("one", "two")
				display.VariadicParam("three", "four", "five")
				display.VerifyWasCalledOnce().VariadicParam(Any[string](), Any[string](), Any[string]())
				display.VerifyWasCalledOnce().VariadicParam(Any[string](), Any[string]())
			})

			It("succeeds when verifying captured arguments", func() {
				display.VariadicParam("one", "two")
				args := display.VerifyWasCalledOnce().VariadicParam(Any[string](), Any[string]()).GetCapturedArguments()
				Expect(args[0]).To(Equal("one"))
				Expect(args[1]).To(Equal("two"))
			})

			It("succeeds when verifying all captured arguments", func() {
				display.VariadicParam("one", "two")
				display.VariadicParam("three", "four", "five")
				args := display.VerifyWasCalledOnce().VariadicParam(Any[string](), Any[string](), Any[string]()).GetCapturedArguments()
				Expect(args[0]).To(Equal("three"))
				Expect(args[1]).To(Equal("four"))
				Expect(args[2]).To(Equal("five"))
			})

		})

		Context("2 normal arguments and one variadic", func() {
			It("succeeds when verifying all captured arguments (one invocation match)", func() {
				display.NormalAndVariadicParam("one", 2, "three", "four")
				display.NormalAndVariadicParam("five", 6, "seven", "eight", "nine")

				stringArg, intArg, varArgs := display.VerifyWasCalled(AtLeast(1)).NormalAndVariadicParam(Any[string](), Any[int](), Any[string](), Any[string]()).GetAllCapturedArguments()
				Expect(stringArg[0]).To(Equal("one"))
				Expect(intArg[0]).To(Equal(2))
				Expect(varArgs[0][0]).To(Equal("three"))
				Expect(varArgs[0][1]).To(Equal("four"))

				stringArg, intArg, varArgs = display.VerifyWasCalled(AtLeast(1)).NormalAndVariadicParam(Any[string](), Any[int](), Any[string](), Any[string](), Any[string]()).GetAllCapturedArguments()
				Expect(stringArg[0]).To(Equal("five"))
				Expect(intArg[0]).To(Equal(6))
				Expect(varArgs[0][0]).To(Equal("seven"))
				Expect(varArgs[0][1]).To(Equal("eight"))
				Expect(varArgs[0][2]).To(Equal("nine"))
			})

			It("succeeds when verifying all captured arguments (multiple invocation matches)", func() {
				display.NormalAndVariadicParam("one", 2, "three", "four")
				display.NormalAndVariadicParam("five", 6, "seven", "eight", "nine")
				display.NormalAndVariadicParam("ten", 11, "twelf", "thirteen", "fourteen")

				stringArg, intArg, varArgs := display.VerifyWasCalled(AtLeast(1)).NormalAndVariadicParam(Any[string](), Any[int](), Any[string](), Any[string]()).GetAllCapturedArguments()
				Expect(stringArg[0]).To(Equal("one"))
				Expect(intArg[0]).To(Equal(2))
				Expect(varArgs[0][0]).To(Equal("three"))
				Expect(varArgs[0][1]).To(Equal("four"))

				stringArg, intArg, varArgs = display.VerifyWasCalled(AtLeast(1)).NormalAndVariadicParam(Any[string](), Any[int](), Any[string](), Any[string](), Any[string]()).GetAllCapturedArguments()
				Expect(stringArg[0]).To(Equal("five"))
				Expect(intArg[0]).To(Equal(6))
				Expect(varArgs[0][0]).To(Equal("seven"))
				Expect(varArgs[0][1]).To(Equal("eight"))
				Expect(varArgs[0][2]).To(Equal("nine"))

				Expect(stringArg[1]).To(Equal("ten"))
				Expect(intArg[1]).To(Equal(11))
				Expect(varArgs[1][0]).To(Equal("twelf"))
				Expect(varArgs[1][1]).To(Equal("thirteen"))
				Expect(varArgs[1][2]).To(Equal("fourteen"))
			})

			It("does not panic when variadic arg has 0 params", func() {
				display.VerifyWasCalled(Never()).NormalAndVariadicParam(Any[string](), Any[int]()).GetAllCapturedArguments()

				display.NormalAndVariadicParam("one", 2)
				display.VerifyWasCalledOnce().NormalAndVariadicParam(Any[string](), Any[int]()).GetAllCapturedArguments()
			})
		})

		Context("Concurrent access to mock", func() {
			It("does not panic", func() {
				Expect(func() {
					wg := sync.WaitGroup{}
					for i := 0; i < 10; i++ {
						wg.Add(1)

						go func() {
							display.SomeValue()
							wg.Done()
						}()
					}
					wg.Wait()
				}).ToNot(Panic())
			})

			Context("Concurrent access due to one mock calling the other", func() {
				It("does not deadlock", func() {
					When(display.SomeValue()).Then(func(params []Param) ReturnValues {
						display.Show("Some irrelevant string")
						return []ReturnValue{}
					})
					display.SomeValue()

					display.VerifyWasCalledOnce().Show(Any[string]())
				})
			})

			Context("Concurrent access with multiple stubbing and validation", func() {
				It("does not panic", func() {
					pegomock.
						When(display.MultipleValues()).
						Then(func(params []pegomock.Param) pegomock.ReturnValues {
							return pegomock.ReturnValues{"MultipleValues", 42, float32(3.14)}
						})

					pegomock.
						When(display.MultipleParamsAndReturnValue(Any[string](), Any[int]())).
						Then(func(params []pegomock.Param) pegomock.ReturnValues {
							return pegomock.ReturnValues{"MultipleParamsAndReturnValue" + params[0].(string)}
						})

					Expect(func() {
						wg := sync.WaitGroup{}

						for i := 0; i < 10; i++ {
							wg.Add(1)

							go func() {
								display.MultipleValues()
								display.MultipleParamsAndReturnValue("TestString", 42)
								wg.Done()
							}()
						}

						wg.Wait()

						display.VerifyWasCalled(Times(10)).MultipleValues()
						display.VerifyWasCalled(Times(10)).MultipleParamsAndReturnValue(Any[string](), Any[int]())
						display.VerifyWasCalled(Never()).SomeValue()
					}).ToNot(Panic())
				})
			})
		})

	})

	Describe("Using VerifyWasCalledEventually when object under test calls goroutine", func() {
		It("correctly fails when timeout is shorter than mock invocation, and succeeds, when timeout is longer", func() {
			go func() {
				time.Sleep(1 * time.Second)
				display.Show("hello")
			}()
			Expect(func() { display.VerifyWasCalledEventually(Once(), 100*time.Millisecond).Show("hello") }).
				To(PanicWithMessageTo(SatisfyAll(
					ContainSubstring("Mock invocation count for Show(\"hello\") does not match expectation"),
					ContainSubstring("after timeout of 100ms"),
					ContainSubstring("Expected: 1; but got: 0"),
				)))

			Expect(func() { display.VerifyWasCalledEventually(Once(), 2*time.Second).Show("hello") }).NotTo(Panic())
		})

	})

	Describe("Manipulating out args (using pointers) in Then blocks", func() {
		It("correctly manipulates the out args", func() {
			type Entity struct{ i int }
			var input []Entity
			When(func() { display.InterfaceParam(Any[interface{}]()) }).Then(func(params []pegomock.Param) pegomock.ReturnValues {
				*params[0].(*[]Entity) = append(*params[0].(*[]Entity), Entity{3})
				return nil
			})

			display.InterfaceParam(&input)

			Expect(input).To(HaveLen(1))
			Expect(input[0].i).To(Equal(3))
		})
	})

	Context("Mock created with custom fail handler", func() {
		It("calls custom fail handler instead of global one", func() {
			failHandlerCalled := false
			display := NewMockDisplay(WithFailHandler(func(message string, callerSkip ...int) {
				failHandlerCalled = true
			}))

			display.VerifyWasCalledOnce().Show("This was never called")

			Expect(failHandlerCalled).To(BeTrue())
		})
	})

	Context("channels", func() {

		Context("using send-/receive-only channels in return types", func() {
			It("allows to return non-direction channels from callbacks", func() {
				When(display.ChanReturnValues()).Then(func([]pegomock.Param) pegomock.ReturnValues {
					return []ReturnValue{make(chan string), make(chan error)}
				})
				display.ChanReturnValues()
			})

			It("allows to return directed channels from callbacks", func() {
				When(display.ChanReturnValues()).Then(func([]pegomock.Param) pegomock.ReturnValues {
					return []ReturnValue{make(<-chan string), make(chan<- error)}
				})
				display.ChanReturnValues()
			})

			It("does not allow to return directed channels from callbacks with wrong direction", func() {
				When(display.ChanReturnValues()).Then(func([]pegomock.Param) pegomock.ReturnValues {
					return []ReturnValue{make(chan<- string), make(chan<- error)}
				})

				Expect(func() { display.ChanReturnValues() }).To(PanicWithMessageTo(MatchError(
					"interface conversion: pegomock.ReturnValue is chan<- string, not <-chan string",
				)))

			})
		})

		Context("using send-/receive-only channels", func() {
			It("generates the mock method with correct channel directions", func() {
				var stringReadChan <-chan string
				var errorWriteChan chan<- error
				display.ChanParams(stringReadChan, errorWriteChan)
			})
		})
	})
})

var _ = Describe("GenericMockDisplay", func() {
	var display *MockGenericDisplay[string, int64]

	BeforeEach(func() {
		display = NewMockGenericDisplay[string, int64]()
	})

	Context("Stubbing generic method with generic value", func() {
		BeforeEach(func() {
			When(display.GenericParams(map[string]int64{"Hello": 333})).ThenReturn(int64(666))
		})

		It("returns stubbed generic value", func() {
			Expect(display.GenericParams(map[string]int64{"Hello": 333})).To(Equal(int64(666)))
		})
	})
})

type expectation struct {
	method   string
	expected string
	actual   string
}

func (e expectation) string() string {
	return fmt.Sprintf("Mock invocation count for %v does not match expectation.\n\n\tExpected: %v; but got: %v",
		e.method, e.expected, e.actual)
}
