FROM golang:1.21-alpine3.18
WORKDIR /app

COPY . .
RUN go mod download

EXPOSE 5000
