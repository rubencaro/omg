#!/bin/bash
go build -v -o omg -ldflags "-X cmd.version=$(git rev-list -1 HEAD)" main.go "$@"