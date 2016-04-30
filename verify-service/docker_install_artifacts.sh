#!/bin/bash
set -ex

if [[ $1 == "" ]]; then
  echo "usage: docker_install_artifacts.sh VERSION"
  exit 1
fi

GOOS=windows GOARCH=amd64 go install /go/src/github.com/gudtech/retailops-sdk/verify-service/bin/verify-service.go && mv /artifacts/verify-service.exe /artifacts/verify-service-windows-amd64-$1.exe
GOOS=linux   GOARCH=amd64 go install /go/src/github.com/gudtech/retailops-sdk/verify-service/bin/verify-service.go && mv /artifacts/verify-service /artifacts/verify-service-linux-amd64-$1
GOOS=darwin  GOARCH=amd64 go install /go/src/github.com/gudtech/retailops-sdk/verify-service/bin/verify-service.go && mv /artifacts/verify-service /artifacts/verify-service-darwin-amd64-$1