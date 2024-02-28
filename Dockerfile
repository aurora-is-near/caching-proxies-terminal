FROM golang:latest

LABEL version="1.0"

RUN mkdir /go/src/caching-proxies-terminal
COPY . /go/src/caching-proxies-terminal
WORKDIR /go/src/caching-proxies-terminal

RUN go mod tidy

ENTRYPOINT ["./entrypoint.sh"]
