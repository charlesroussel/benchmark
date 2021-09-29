FROM golang:alpine
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
COPY . /go-http
WORKDIR /go-http
RUN go get -d -v
RUN go build -o /go/bin/fiber_http
EXPOSE 8080:8080
ENV CORE_COUNT 0
CMD [ "/go/bin/fiber_http" ]
