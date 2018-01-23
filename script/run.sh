#!/bin/bash
go run -ldflags "-X main.version=$(git rev-parse --short HEAD)" main.go "$@"