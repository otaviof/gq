FROM golang:alpine AS builder

RUN apk update && \
    apk add --no-cache make

COPY . /src/
WORKDIR /src

RUN make vendor install

FROM alpine:latest

COPY --from=builder /go/bin/gq /go/bin/gq

ENTRYPOINT ["/go/bin/gq"]
