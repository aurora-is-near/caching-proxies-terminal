# Build stage
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache openssh-client git

COPY . /app

WORKDIR /app

RUN go build -o app .

# Start with a base image that includes the necessary runtime
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app /app/app

RUN chmod +x /app/app