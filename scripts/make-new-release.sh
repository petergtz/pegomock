#!/bin/bash

set -ex
cd $(dirname $0)/..

echo "This will push the current branch to GitHub and make a new release. "
echo "Make sure you're on master."
echo "Make sure you've tagged the current commit with correct version."
echo "Make sure you've exported the GITHUB_TOKEN"
echo "Only proceed if everything is correct. ctrl+c to abort"

read

git push
goreleaser release --rm-dist