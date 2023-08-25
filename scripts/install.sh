#!/bin/bash

set -eux

env GO111MODULE=on CGO_ENABLED=0 go build -ldflags "-X github.com/khulnasoft-lab/tfsecurity/version.Version=${1}" ./cmd/tfsecurity
