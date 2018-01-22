#!/bin/bash
go run -ldflags "-X cmd.version=$(git rev-list -1 HEAD)" main.go "$@"