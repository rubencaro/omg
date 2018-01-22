#!/bin/bash
go build -v -o omg -ldflags "-X main.version=$(git rev-list -1 HEAD)" main.go "$@"