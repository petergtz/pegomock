#!/bin/bash

set -ex
cd $(dirname $0)/..

rm -f mock_display_test.go
rm -f mock_generic_display_test.go
ginkgo -succinct generate_test_mocks/xtools_go_loader
ginkgo -r --skip-package=generate_test_mocks/xtools_go_loader,pegomock/watch --randomize-all --randomize-suites --race --trace -cover
