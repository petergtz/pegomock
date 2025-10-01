package mock_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/petergtz/pegomock/v4/ginkgo_compatible"
)

var _ = Describe("Compatible", func() {
	It("can import Ginkgo compatible package", func() {
		_ = Whenever
	})
})
