FROM golang:1.8
MAINTAINER Hantao Wang

EXPOSE 8081

RUN mkdir -p /go/src/gitlab.com/will.wang1
RUN mkdir -p /go/bin
RUN go get github.com/go-redis/redis
RUN go get github.com/lib/pq
RUN go get github.com/sirupsen/logrus

WORKDIR /go/src/gitlab.com/will.wang1

ARG git_pass
ARG build_time=1

RUN git clone https://user:$git_pass@gitlab.com/will.wang1/hotrod-base
RUN git clone https://user:$git_pass@gitlab.com/will.wang1/hotrod-customer

WORKDIR /go/src/gitlab.com/will.wang1/hotrod-customer

RUN go build -o hotrod main.go
RUN mv hotrod /go/bin/

ENTRYPOINT ["/go/bin/hotrod", "customer"]
