#!/bin/sh
GOOS=linux GOARCH=386 go build github.com/fission/fission-workflows/cmd/fission-workflows-bundle/
