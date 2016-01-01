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
		It("Returns zero value", func() {
			Expect(display.SomeValue()).To(Equal(""))
		})
	})

	Context("Calling Flash() without matchers", func() {
		It("Matches correctly", func() {
			display.Flash("Hello", 333)
			Expect(func() { display.VerifyWasCalled().Flash("Hello", 333) }).NotTo(Panic())
			Expect(func() { display.VerifyWasCalled().Flash("Hello", 666) }).To(Panic())
		})
	})

	Context("Calling Flash() with matchers", func() {
		It("Matches correctly", func() {
			When(display.MultipleParamsAndReturnValue(EqString("Hello"), EqInt(333))).ThenReturn("Bla")
			Expect(func() { display.VerifyWasCalled().MultipleParamsAndReturnValue("Hello", 333) }).To(Panic())
			display.MultipleParamsAndReturnValue("Hello", 333)
			Expect(func() { display.VerifyWasCalled().MultipleParamsAndReturnValue("Hello", 333) }).NotTo(Panic())
		})
	})

	Context("Calling Flash() only with partial matchers", func() {
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

		It("panics if verify is called on mock that was not invoked.", func() {
			Expect(func() { display.VerifyWasCalled().Show("param") }).To(Panic())

		})
	})

	Context("Stubbing with invalid return type", func() {
		It("panics", func() {
			Expect(func() { When(display.SomeValue()).ThenReturn("Hello").ThenReturn(0) }).To(Panic())
		})
	})

	Context("No invocation happened", func() {
		It("Panics when trying to verify", func() {
			When(display.SomeValue()).ThenReturn("Hello")
			Expect(func() { display.VerifyWasCalled().SomeValue() }).To(Panic())
		})
	})
})
