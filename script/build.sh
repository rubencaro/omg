#!/bin/bash
go build -v -o omg -ldflags "-X main.version=$(git rev-parse --short HEAD)" main.go "$@"