FROM golang:alpine
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
COPY . /go-http
WORKDIR /go-http
RUN go get -d -v
RUN CGO_ENABLED=1 go build -o /go/bin/fiber_http
EXPOSE 8080:8080
CMD [ "/go/bin/fiber_http" ]
