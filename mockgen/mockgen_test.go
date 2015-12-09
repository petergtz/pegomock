package main_test

import (
	mockgen "github.com/petergtz/pegomock/mockgen"

	. "github.com/onsi/ginkgo"
	_ "github.com/onsi/gomega"
)

var _ = Describe("mockgen", func() {
	Context("When run in reflect mode", func() {
		It("Generates mock", func() {
			mockgen.Run("",
				"../pegomock/mock_display_test.go", "pegomock_test",
				"",
				false,
				"github.com/petergtz/pegomock/test_interface", "Display")
		})
	})
})
