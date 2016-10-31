// Copyright 2015 Peter Goetz
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pegomock_test

import (
	"testing"

	. "github.com/petergtz/pegomock"
	"github.com/petergtz/pegomock/modelgen/gomock"
	"github.com/petergtz/pegomock/modelgen/loader"

	"github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
	// . "github.com/petergtz/pegomock/pegomock/testutil"
)

func TestDSL(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	pegomock.RegisterMockFailHandler(func(message string, callerSkip ...int) { panic(message) })
	ginkgo.RunSpecs(t, "modelgen Suite")
}

var _ = Describe("MockDisplay", func() {
	It("Bla", func() {
		pkgFromReflect, e := gomock.Reflect("github.com/petergtz/pegomock/test_interface", []string{"Display"})
		Expect(e).NotTo(HaveOccurred())
		pkgFromLoader, e := loader.GenerateModel("github.com/petergtz/pegomock/test_interface", "Display")
		Expect(e).NotTo(HaveOccurred())
		// spew.Dump(pkg)
		Expect(pkgFromLoader).To(Equal(pkgFromReflect))
	})
})
