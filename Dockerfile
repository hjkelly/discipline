FROM golang:1.7-alpine

ADD . /go/src/github.com/hjkelly/discipline

RUN go install github.com/hjkelly/discipline

ENTRYPOINT /go/bin/discipline

EXPOSE 8080
