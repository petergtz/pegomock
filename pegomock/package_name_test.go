package main_test

import (
	"go/build"
	"os"

	. "github.com/petergtz/pegomock/pegomock/testutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	main "github.com/petergtz/pegomock/pegomock"
)

var _ = Describe("DetermineBla", func() {
	var (
		packageDir string
	)
	BeforeEach(func() {
		packageDir = joinPath(build.Default.GOPATH, "src", "package_dir")
	})

	JustBeforeEach(func() {
		Expect(os.MkdirAll(packageDir, 0755)).To(Succeed())
	})

	AfterEach(func() {
		Expect(os.RemoveAll(packageDir)).To(Succeed())
	})

	XContext("only one go file", func() {
		It("names the package name after package name + _test suffix", func() {
			WriteFile(joinPath(packageDir, "mydisplay.go"), "package package_name")

			Expect(main.DeterminePackageNameIn(packageDir)).To(Equal("package_name_test"))
		})
	})

	XContext("multiple go files with different package names", func() {
		It("fails", func() {
			WriteFile(joinPath(packageDir, "mydisplay.go"), "package package_name")
			WriteFile(joinPath(packageDir, "other.go"), "package other_package_name")

			_, e := main.DeterminePackageNameIn(packageDir)
			Expect(e).To(MatchError(ContainSubstring("Error while determining Go package")))
		})
	})

	XContext("go file and go test file", func() {
		It("determines the package name from the test file", func() {
			WriteFile(joinPath(packageDir, "mydisplay.go"), "package package_name")
			WriteFile(joinPath(packageDir, "mydisplay_test.go"), "package non_conventional_package_name_test")

			Expect(main.DeterminePackageNameIn(packageDir)).To(Equal("non_conventional_package_name_test"))
		})
	})

	Context("no files", func() {
		It("names the package after the directory name base", func() {
			Expect(main.DeterminePackageNameIn(packageDir)).To(Equal("package_dir_test"))
		})

		Context("current dir with dashes in name", func() {
			BeforeEach(func() {
				packageDir = joinPath(build.Default.GOPATH, "src", "package-dir-with-dashes")
			})

			It("names the package after the directory name base, but replace dashes with underscores", func() {
				Expect(main.DeterminePackageNameIn(packageDir)).To(Equal("package_dir_with_dashes_test"))
			})
		})
	})

})
