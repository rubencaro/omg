#!/bin/bash
go build -v -o omg -ldflags "-X main.commit=$(git rev-parse --short HEAD)" main.go "$@"