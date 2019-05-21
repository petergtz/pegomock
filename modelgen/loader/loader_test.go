package loader_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/petergtz/pegomock/modelgen/loader"
)

var _ = Describe("Loader", func() {
	Describe("GenerateModel", func() {
		It("finds all methods within interface", func() {
			pkg, e := GenerateModel("io", "Reader")
			Expect(e).NotTo(HaveOccurred())
			Expect(pkg.Interfaces).To(HaveLen(1))
			Expect(pkg.Interfaces[0].Name).To(Equal("Reader"))
			Expect(pkg.Interfaces[0].Methods).To(HaveLen(1))
			Expect(pkg.Interfaces[0].Methods[0].Name).To(Equal("Read"))
		})

		Context("using an interface with embedded interfaces", func() {
			It("finds all methods", func() {
				pkg, e := GenerateModel("io", "ReadCloser")
				Expect(e).NotTo(HaveOccurred())
				Expect(pkg.Interfaces).To(HaveLen(1))
				Expect(pkg.Interfaces[0].Name).To(Equal("ReadCloser"))
				Expect(pkg.Interfaces[0].Methods).To(HaveLen(2))
				Expect(pkg.Interfaces[0].Methods[0].Name).To(Equal("Read"))
				Expect(pkg.Interfaces[0].Methods[1].Name).To(Equal("Close"))

			})
		})
	})
})
