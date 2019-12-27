FROM golang:1.12.14-alpine3.11

RUN apk add git
WORKDIR /go/src/github.com/kurenkoff/bstu_backend
COPY . .
ENV GO111MODULE=on
RUN ["go", "build", "-o", "srv"]

ENTRYPOINT ["./srv"]