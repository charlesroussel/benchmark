FROM golang:alpine as builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
COPY . /go-http
WORKDIR /go-http
ENV GO111MODULE=on
RUN go get -d -v
RUN CGO_ENABLED=0 go build -o /go/bin/fiber_http
EXPOSE 8080:8080
ENV CORE_COUNT 0


FROM golang:alpine
COPY --from=builder /go/bin/fiber_http /go/bin/fiber_http
CMD [ "/go/bin/fiber_http" ]