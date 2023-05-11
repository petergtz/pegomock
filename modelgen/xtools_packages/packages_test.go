package xtools_packages_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock/v4/model"

	. "github.com/petergtz/pegomock/v4/modelgen/xtools_packages"
)

var _ = Describe("Packages", func() {
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
				Expect([]string{pkg.Interfaces[0].Methods[0].Name, pkg.Interfaces[0].Methods[1].Name}).To(
					ConsistOf("Read", "Close"))
			})
		})

		It("finds correct generic parameters in an interface", func() {
			pkg, e := GenerateModel("io", "Reader")
			Expect(e).NotTo(HaveOccurred())
			Expect(pkg.Interfaces).To(HaveLen(1))
			Expect(pkg.Interfaces[0].Name).To(Equal("Reader"))
			Expect(pkg.Interfaces[0].Methods).To(HaveLen(1))
			Expect(pkg.Interfaces[0].Methods[0].Name).To(Equal("Read"))

			pkg, e = GenerateModel("github.com/petergtz/pegomock/v4/modelgen/xtools_packages", "Bla")
			Expect(e).NotTo(HaveOccurred())
			Expect(pkg.Interfaces).To(HaveLen(1))
			Expect(pkg.Interfaces[0].Name).To(Equal("Bla"))
			Expect(pkg.Interfaces[0].TypeParams).To(ConsistOf(
				&model.Parameter{
					Name: "K",
					Type: model.PredeclaredType("comparable"),
				},
				&model.Parameter{
					Name: "V",
					Type: &model.NamedType{
						Package: "github.com/petergtz/pegomock/v4/modelgen/xtools_packages",
						Type:    "Number",
					},
				},
			))
			Expect(pkg.Interfaces[0].Methods).To(HaveLen(1))
			Expect(pkg.Interfaces[0].Methods[0].Name).To(Equal("SumNumbers"))
		})

	})
})
