# syntax=docker/dockerfile:1

FROM golang:1.21.4-alpine AS build
RUN apk add --no-cache bash
WORKDIR /workspace
COPY go.mod .
RUN go mod download

CMD ["bash"]
