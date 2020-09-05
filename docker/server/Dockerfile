FROM golang:1.14

WORKDIR /go/src/github.com/ughvj/goguild-ergo-sum

RUN go get -u google.golang.org/grpc && \
    go get -u github.com/grpc-ecosystem/go-grpc-middleware && \
    go get -u github.com/oxequa/realize && \
    go get -u github.com/gomodule/redigo/redis