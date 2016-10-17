#!/bin/bash

set -ex

rm -f mock_display_test.go
time $GOPATH/bin/ginkgo pegomock/mockgen_test_from_reflect
time $GOPATH/bin/ginkgo -r -skipPackage=pegomock/mockgen_test_from_source,pegomock/mockgen_test_from_reflect --randomizeAllSpecs --randomizeSuites --race --trace -cover
time $GOPATH/bin/ginkgo pegomock/mockgen_test_from_source
time $GOPATH/bin/ginkgo -r -skipPackage=pegomock/mockgen_test_from_source,pegomock/mockgen_test_from_reflect --randomizeAllSpecs --randomizeSuites --race --trace -cover
