#!/bin/bash
go test -v -ldflags "-X main.commit=$(git rev-parse --short HEAD)" ./test/... "$@"