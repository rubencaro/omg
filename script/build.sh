#!/bin/bash
go build -o omg -ldflags "-X main.version=$(git rev-list -1 HEAD)" main.go "$@"