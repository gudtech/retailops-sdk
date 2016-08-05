FROM golang:1.6

RUN useradd -m -d /opt/gt gt \
    && rm -rf /src/patch \
    && mkdir /backplane /private \
    && chown gt:gt /backplane /private \
    && ln -sf /backplane/etc /etc/GT \
    && ln -sf /backplane/etc /etc/GTSOA \
    && ln -sf /backplane/log /var/log/GTSOA \
    && ln -sf /backplane/DBR_gt.conf /etc/DBR_gt.conf \
    && mkdir -p /var/local/spool \
    && ln -sf /backplane/spool /var/local/spool/gt-eventbus \
    && ln -sf /backplane/var /var/run/GTSOA \
    && ln -sf /private/etc /etc/GT_private \
    && apt-get update && apt-get install -y strace vim \
    ;
VOLUME /backplane
VOLUME /private

ENV GOBIN=/usr/local/bin

COPY verify/docker-entrypoint.sh /
COPY verify/provision-soa-service /usr/local/bin
COPY . /go/src/github.com/gudtech/retailops-sdk

RUN go get -d github.com/gudtech/scamp-go/scamp && \
    go install /go/src/github.com/gudtech/retailops-sdk/verify/bin/service.go && \
    go install /go/src/github.com/gudtech/scamp-go/bin/scamp.go

    # && \
    #rm -rf /go/src/*

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["service", "-config", "/backplane/etc/soa.conf"]
