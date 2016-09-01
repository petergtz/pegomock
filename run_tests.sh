#!/bin/bash

rm -f mock_display_test.go
time $GOPATH/bin/ginkgo pegomock/mockgen
time $GOPATH/bin/ginkgo -r -skipPackage=pegomock/mockgen_test_from_source,pegomock/mockgen --randomizeAllSpecs --randomizeSuites --race --trace
time $GOPATH/bin/ginkgo pegomock/mockgen_test_from_source
time $GOPATH/bin/ginkgo -r -skipPackage=pegomock/mockgen_test_from_source,pegomock/mockgen --randomizeAllSpecs --randomizeSuites --race --trace
