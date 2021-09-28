FROM golang:alpine
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
COPY . /grpc-server
WORKDIR /grpc-server
RUN go get -d -v
RUN CGO_ENABLED=1 go build -o /go/bin/grpc
EXPOSE 8080:8080
CMD [ "/go/bin/grpc" ]
