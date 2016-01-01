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
	. "github.com/petergtz/pegomock/pegomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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
			When(display.MultipleParamsAndReturnValue(EqString("Hello"), EqInt(333))).ThenReturn("Bla")
		})

		It("fails during verification when mock was not called", func() {
			Expect(func() { display.VerifyWasCalled().MultipleParamsAndReturnValue("Hello", 333) }).To(Panic())
		})

		It("succeeds verification when mock was called", func() {
			display.MultipleParamsAndReturnValue("Hello", 333)
			Expect(func() { display.VerifyWasCalled().MultipleParamsAndReturnValue("Hello", 333) }).NotTo(Panic())
		})

		It("succeeds verification when verification and invocation are mixed", func() {
			Expect(func() { display.VerifyWasCalled().MultipleParamsAndReturnValue("Hello", 333) }).To(Panic())
			display.MultipleParamsAndReturnValue("Hello", 333)
			Expect(func() { display.VerifyWasCalled().MultipleParamsAndReturnValue("Hello", 333) }).NotTo(Panic())
		})
	})

	Context("Calling MultipleParamsAndReturnValue() with \"Any\"-matchers", func() {
		It("succeeds all verifications that match", func() {
			When(display.MultipleParamsAndReturnValue(AnyString(), EqInt(333))).ThenReturn("Bla")

			Expect(func() { display.VerifyWasCalled().MultipleParamsAndReturnValue("Hello", 333) }).To(Panic())

			display.MultipleParamsAndReturnValue("Hello", 333)
			display.MultipleParamsAndReturnValue("Hello again", 333)
			display.MultipleParamsAndReturnValue("And again", 333)

			Expect(func() { display.VerifyWasCalled().MultipleParamsAndReturnValue("Hello", 333) }).NotTo(Panic())
			Expect(func() { display.VerifyWasCalled().MultipleParamsAndReturnValue("Hello again", 333) }).NotTo(Panic())
			Expect(func() { display.VerifyWasCalled().MultipleParamsAndReturnValue("And again", 333) }).NotTo(Panic())

			Expect(func() { display.VerifyWasCalled().MultipleParamsAndReturnValue("And again", 444) }).To(Panic())

		})
	})

	Context("Calling MultipleParamsAndReturnValue() only with matchers on some parameters", func() {
		It("panics", func() {
			Expect(func() { When(display.MultipleParamsAndReturnValue(EqString("Hello"), 333)) }).To(Panic())
		})
	})

	Context("Stubbing with consecutive return values", func() {
		BeforeEach(func() {
			Expect(display.SomeValue()).To(Equal(""))
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
			Expect(func() { display.VerifyWasCalled().SomeValue() }).NotTo(Panic())
		})

		It("fails if verify is called on mock that was not invoked.", func() {
			Expect(func() { display.VerifyWasCalled().Show("param") }).To(Panic())

		})
	})

	Context("Stubbing with invalid return type", func() {
		It("panics", func() {
			Expect(func() { When(display.SomeValue()).ThenReturn("Hello").ThenReturn(0) }).To(Panic())
		})
	})

	Context("Stubbed method, but no invocation takes place", func() {
		It("fails during verification", func() {
			When(display.SomeValue()).ThenReturn("Hello")
			Expect(func() { display.VerifyWasCalled().SomeValue() }).To(Panic())
		})
	})

	Context("Calling Flash() with specific arguments", func() {

		BeforeEach(func() { display.Flash("Hello", 333) })

		It("succeeds verification if values are matching", func() {
			Expect(func() { display.VerifyWasCalled().Flash("Hello", 333) }).NotTo(Panic())
		})

		It("fails during verification if values are not matching", func() {
			Expect(func() { display.VerifyWasCalled().Flash("Hello", 666) }).To(Panic())
		})

		It("succeeds during verification when using Any-matchers ", func() {
			Expect(func() { display.VerifyWasCalled().Flash(AnyString(), AnyInt()) }).NotTo(Panic())
		})

		It("succeeds during verification when using valid Eq-matchers ", func() {
			Expect(func() { display.VerifyWasCalled().Flash(EqString("Hello"), EqInt(333)) }).NotTo(Panic())
		})

		It("fails during verification when using invalid Eq-matchers ", func() {
			Expect(func() { display.VerifyWasCalled().Flash(EqString("Invalid"), EqInt(-1)) }).To(Panic())
		})
	})

})
