FROM golang:1.6

ADD . /go/src/github.com/gudtech/retailops-sdk/verify-service

ENV GOBIN=/artifacts
ENTRYPOINT ["/go/src/github.com/gudtech/retailops-sdk/verify-service/go_install_all_platforms.sh"]