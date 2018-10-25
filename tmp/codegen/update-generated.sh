#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

vendor/k8s.io/code-generator/generate-groups.sh \
deepcopy \
github.com/integr8ly/stuff/pkg/generated \
github.com/integr8ly/stuff/pkg/apis \
example:stuff \
--go-header-file "./tmp/codegen/boilerplate.go.txt"
