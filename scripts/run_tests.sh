#!/bin/bash

set -ex
cd $(dirname $0)/..
export PATH=$PATH:~/go/bin
go install github.com/onsi/ginkgo/ginkgo

PACKAGES_TO_SKIP='generate_test_mocks/xtools_go_loader,generate_test_mocks/gomock_reflect,generate_test_mocks/gomock_source'
rm -f mock_display_test.go
rm -rf matchers
ginkgo -succinct generate_test_mocks/xtools_go_loader
ginkgo -r -skipPackage=$PACKAGES_TO_SKIP --randomizeAllSpecs --randomizeSuites --race --trace -cover

rm -f mock_display_test.go
rm -rf matchers
ginkgo -succinct generate_test_mocks/gomock_reflect
ginkgo --randomizeAllSpecs --randomizeSuites --race --trace -cover

rm -f mock_display_test.go
rm -rf matchers
ginkgo -succinct generate_test_mocks/gomock_source
ginkgo --randomizeAllSpecs --randomizeSuites --race --trace -cover
