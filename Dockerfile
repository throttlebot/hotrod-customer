FROM golang:1.8
MAINTAINER Hantao Wang

EXPOSE 8081

RUN mkdir -p /go/src/github.com/kelda-inc
RUN mkdir -p /go/bin

ADD . /go/src/github.com/kelda-inc/hotrod-customer

WORKDIR /go/src/github.com/kelda-inc/hotrod-customer

RUN go build -o hotrod main.go
RUN mv hotrod /go/bin/

ENTRYPOINT ["/go/bin/hotrod", "customer"]
