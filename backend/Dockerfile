FROM golang:1.21.5-alpine3.18
WORKDIR /app/server

COPY go.sum go.mod ./
RUN go mod download
