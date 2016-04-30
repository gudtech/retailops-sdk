FROM golang:1.6

ADD . /go/src/github.com/gudtech/retailops-sdk/verify-service

ENV GOBIN=/artifacts
ENTRYPOINT ["/go/src/github.com/gudtech/retailops-sdk/verify-service/docker_install_artifacts.sh"]