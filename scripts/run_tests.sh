#!/bin/bash

set -ex
cd $(dirname $0)/..

PACKAGES_TO_SKIP='generate_test_mocks/xtools_go_loader,generate_test_mocks/gomock_reflect,generate_test_mocks/gomock_source'
rm -f mock_display_test.go
rm -rf matchers
ginkgo -succinct generate_test_mocks/xtools_go_loader
ginkgo -r --skip-package=$PACKAGES_TO_SKIP,pegomock/watch --randomize-all --randomize-suites --race --trace -cover

rm -f mock_display_test.go
rm -rf matchers
ginkgo -succinct generate_test_mocks/gomock_reflect
ginkgo --skip-package=pegomock/watch --randomize-all --randomize-suites --race --trace -cover

# DEPRECATED: gomock_source is deprecated and will be removed in a future release.
#rm -f mock_display_test.go
#rm -rf matchers
#ginkgo -succinct generate_test_mocks/gomock_source
#ginkgo --skip-package=pegomock/watch --randomize-all --randomize-suites --race --trace -cover
