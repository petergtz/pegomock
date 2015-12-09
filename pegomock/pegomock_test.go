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
		It("Returns default value", func() {
			Expect(display.SomeValue()).To(Equal(""))
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
