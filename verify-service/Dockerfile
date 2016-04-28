FROM golang:1.6

ADD . /go/src/github.com/gudtech/retailops-sdk/verify-service

ENV GOBIN=/artifacts
RUN mkdir /artifacts && \
    GOOS=windows GOARCH=amd64 go install /go/src/github.com/gudtech/retailops-sdk/verify-service/bin/verify-service.go && mv /artifacts/verify-service.exe /artifacts/verify-service-windows-x86_64.exe \
    GOOS=linux GOARCH=amd64  go install /go/src/github.com/gudtech/retailops-sdk/verify-service/bin/verify-service.go && mv /artifacts/verify-service /artifacts/verify-service-linux-x86_64 \
    GOOS=darwin GOARCH=amd64  go install /go/src/github.com/gudtech/retailops-sdk/verify-service/bin/verify-service.go && mv /artifacts/verify-service /artifacts/verify-service-darwin-x86_64