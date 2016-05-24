#!/bin/bash

time $GOPATH/bin/ginkgo pegomock/mockgen
time $GOPATH/bin/ginkgo -r --randomizeAllSpecs --randomizeSuites --race --trace
