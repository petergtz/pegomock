package util_test

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock/v3/mockgen/util"
)

func TestUtil(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Util Suite")
}

var _ = Describe("util", func() {
	Context("sort", func() {
		It("SortedKeys returns map keys in alphabetical order", func() {
			importPaths := map[string]bool{
				"github.com/b/mypackage": true,
				"github.com/c/mypackage": true,
				"github.com/a/mypackage": true,
			}

			sortedImportPaths := util.SortedKeys(importPaths)
			Expect(sortedImportPaths).To(HaveLen(3))
			Expect(sortedImportPaths[0]).To(Equal("github.com/a/mypackage"))
			Expect(sortedImportPaths[1]).To(Equal("github.com/b/mypackage"))
			Expect(sortedImportPaths[2]).To(Equal("github.com/c/mypackage"))
		})
	})
})
