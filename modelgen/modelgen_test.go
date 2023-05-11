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

package modelgen_test

import (
	"testing"

	"github.com/petergtz/pegomock/v4/model"
	"github.com/petergtz/pegomock/v4/modelgen/xtools_packages"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
)

func TestDSL(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "modelgen Suite")
}

type alphabetically []*model.Method

func (a alphabetically) Len() int           { return len(a) }
func (a alphabetically) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a alphabetically) Less(i, j int) bool { return a[i].Name < a[j].Name }

var _ = Describe("xtools_packages", func() {

	It("generates a model with the basic properties", func() {
		pkg, e := xtools_packages.GenerateModel("github.com/petergtz/pegomock/v4/modelgen/test_data/default_test_interface", "Display")
		Expect(e).NotTo(HaveOccurred())

		Expect(pkg.Name).To(Equal("test_interface"))
		Expect(pkg.Interfaces).To(HaveLen(1))
		Expect(pkg.Interfaces[0].Name).To(Equal("Display"))

		Expect(pkg.Interfaces[0].Methods).To(ContainElement(
			&model.Method{
				Name: "Show",
				In: []*model.Parameter{
					{
						Name: "_param0",
						Type: model.PredeclaredType("string"),
					},
				},
			},
		))

		// TODO add more test cases
	})
})
