#!/bin/bash
set -ex

GOOS=windows GOARCH=amd64 go install /go/src/github.com/gudtech/retailops-sdk/verify-service/bin/verify-service.go && mv /artifacts/verify-service.exe /artifacts/verify-windows.exe
GOOS=linux   GOARCH=amd64 go install /go/src/github.com/gudtech/retailops-sdk/verify-service/bin/verify-service.go && mv /artifacts/verify-service /artifacts/verify-linux
GOOS=darwin  GOARCH=amd64 go install /go/src/github.com/gudtech/retailops-sdk/verify-service/bin/verify-service.go && mv /artifacts/verify-service /artifacts/verify-darwin