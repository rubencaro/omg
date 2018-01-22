#!/bin/bash
go run -ldflags "-X main.version=$(git rev-list -1 HEAD)" main.go "$@"