package gomock_test

import (
	"io"
	"os"
	"testing"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock/v3/modelgen/gomock"
)

func TestGomock(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "gomock Suite")
}

var _ = XDescribe("reflect", func() {
	AfterEach(func() {
		os.RemoveAll("../../vendor")
	})

	It("can generate mocks for interfaces taken from vendored packages", func() {
		e := os.MkdirAll("../../vendor/github.com/petergtz/vendored_package/", 0755)
		Expect(e).NotTo(HaveOccurred())

		file, e := os.Create("../../vendor/github.com/petergtz/vendored_package/iface.go")
		Expect(e).NotTo(HaveOccurred())

		_, e = io.WriteString(file, "package vendored_package\n\ntype Interface interface{}")
		Expect(e).NotTo(HaveOccurred())

		_, e = gomock.Reflect("github.com/petergtz/vendored_package", []string{"Interface"})
		Expect(e).NotTo(HaveOccurred())
	})
})
