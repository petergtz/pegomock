package pegomock_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPegomock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pegomock Suite")
}
