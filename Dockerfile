# Dockerfile for building the image
FROM golang:1.22.4-alpine3.20 AS builder

ARG CYAMLI_VERSION=latest
RUN go install "github.com/Jumpaku/cyamli@$CYAMLI_VERSION"


FROM busybox:stable-musl

COPY --from=builder /go/bin/cyamli /bin/cyamli
COPY docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh
WORKDIR /workspace
ENTRYPOINT ["/docker-entrypoint.sh"]
