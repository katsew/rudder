FROM golang:1.12.9-stretch as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN groupadd -g 10001 anon \
    && useradd -u 10001 -g anon anon

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

FROM debian:stretch-slim

RUN apt-get update && apt-get install -y iptables procps tcpdump netcat

COPY --from=builder /go/bin/app /go/bin/app
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

USER anon

ENTRYPOINT ["/go/bin/app"]