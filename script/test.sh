#!/bin/bash
go test -v -ldflags "-X main.version=$(git rev-parse --short HEAD)" ./test/... "$@"