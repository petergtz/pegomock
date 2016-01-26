package main_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock/pegomock"

	"testing"
)

func TestPegomock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pegomock CLI Suite")
}

var _ = Describe("Test pegomock CLI", func() {

	var (
		packageDir, subPackageDir string
		app                       *kingpin.Application
	)

	BeforeEach(func() {
		packageDir = filepath.Join(os.Getenv("GOPATH"), "src", "pegomocktest")
		Expect(os.MkdirAll(packageDir, 0755)).To(Succeed())
		subPackageDir = filepath.Join(packageDir, "subpackage")
		Expect(os.MkdirAll(subPackageDir, 0755)).To(Succeed())
		os.Chdir(packageDir)

		Expect(ioutil.WriteFile(
			filepath.Join(packageDir, "mydisplay.go"),
			[]byte("package pegomocktest; type MyDisplay interface {  Show() }"),
			0644)).
			To(Succeed())
		Expect(ioutil.WriteFile(
			filepath.Join(subPackageDir, "subdisplay.go"),
			[]byte("package pegomocktest; type SubDisplay interface {  ShowMe() }"),
			0644)).
			To(Succeed())
		app = kingpin.New("pegomock", "Generates mocks based on interfaces.")
		app.Terminate(func(int) { panic("Unexpected terminate") })

	})

	AfterEach(func() {
		Expect(os.RemoveAll(packageDir)).To(Succeed())
	})

	Context("Executing pegomock generate MyDisplay", func() {

		It(`generates a file mock_mydisplay_test.go that contains "package pegomocktest_test"`, func() {
			main.Run(cmd("pegomock generate MyDisplay"), os.Stdout, app)

			mockFile := filepath.Join(packageDir, "mock_mydisplay_test.go")
			Expect(mockFile).To(BeAnExistingFile())
			Expect(ioutil.ReadFile(mockFile)).To(ContainSubstring("package pegomocktest_test"))
		})
	})

	Context("Executing pegomock generate pegomocktest/subpackage SubDisplay", func() {

		It(`generates a file mock_subdisplay_test.go in "pegomocktest" that contains "package pegomocktest_test"`, func() {
			main.Run(cmd("pegomock generate pegomocktest/subpackage SubDisplay"), os.Stdout, app)

			mockFile := filepath.Join(packageDir, "mock_subdisplay_test.go")
			Expect(mockFile).To(BeAnExistingFile())
			Expect(ioutil.ReadFile(mockFile)).To(ContainSubstring("package pegomocktest_test"))
		})
	})

	Context("Executing pegomock generate with too many args", func() {

		It(`reports an error and the usage`, func() {
			var buf bytes.Buffer
			app.Terminate(nil)
			main.Run(cmd("pegomock generate with too many args"), &buf, app)
			Expect(buf.String()).To(ContainSubstring("Please provide exactly 1 interface or 1 package + 1 interface"))
			Expect(buf.String()).To(ContainSubstring("usage"))
		})
	})

	Context("Executing pegomock some unknown command", func() {

		It(`reports an error and the usage`, func() {
			var buf bytes.Buffer
			kingpin.CommandLine.Terminate(nil)
			kingpin.CommandLine.Writer(&buf)

			main.Run(cmd("pegomock some unknown command"), &buf, app)
			Expect(buf.String()).To(ContainSubstring("error"))
		})
	})

})

func cmd(line string) []string {
	return strings.Split(line, " ")
}
