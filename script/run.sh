#!/bin/bash
go run -ldflags "-X main.commit=$(git rev-parse --short HEAD)" main.go "$@"