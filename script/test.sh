#!/bin/bash
go test -v -ldflags "-X main.version=$(git rev-list -1 HEAD)" ./test/... "$@"