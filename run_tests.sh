#!/bin/bash

$GOPATH/bin/ginkgo pegomock/mockgen
$GOPATH/bin/ginkgo -r --randomizeAllSpecs --randomizeSuites --race --trace
