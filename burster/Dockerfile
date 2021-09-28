FROM golang:alpine
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
WORKDIR $GOPATH/src/
ENV GHZ_VERSION=v0.103.0
RUN git clone --depth 1 --branch ${GHZ_VERSION} https://github.com/bojand/ghz
WORKDIR $GOPATH/src/ghz/cmd/ghz
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /go/bin/ghz

COPY ghz /grpc/ghz
WORKDIR /grpc/ghz

COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

COPY ../proto /grpc/proto

ENTRYPOINT ["./entrypoint.sh"]