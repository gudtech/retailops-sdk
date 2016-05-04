#!/bin/bash
set -ex

GOOS=windows GOARCH=amd64 go install /go/src/github.com/gudtech/retailops-sdk/verify/bin/verify.go && mv /artifacts/verify.exe /artifacts/verify-windows.exe
GOOS=linux   GOARCH=amd64 go install /go/src/github.com/gudtech/retailops-sdk/verify/bin/verify.go && mv /artifacts/verify /artifacts/verify-linux
GOOS=darwin  GOARCH=amd64 go install /go/src/github.com/gudtech/retailops-sdk/verify/bin/verify.go && mv /artifacts/verify /artifacts/verify-darwin